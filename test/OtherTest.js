import { advanceBlock } from './helpers/advanceToBlock';
import ether from './helpers/ether';
import { increaseTimeTo, duration } from './helpers/increaseTime';
import latestTime from './helpers/latestTime';

const BigNumber = web3.BigNumber;
const should = require('chai')
    .use(require('chai-as-promised'))
    .use(require('chai-bignumber')(BigNumber))
    .should();

const Crowdsale = artifacts.require('./Crowdsale.sol');
const Token = artifacts.require('./WindMineToken.sol');

contract('OtherTest', function ([wallet, foundersWallet, frozeWallet, ...accounts]) {
  before(async function () {
    // Advance to the next block to correctly read time in the solidity "now" function interpreted by testrpc
    await advanceBlock();
  });

  beforeEach(async function () {
    this.privateSaleStartDate = latestTime() + duration.weeks(1);
    this.preIcoStartDate = this.privateSaleStartDate + duration.weeks(1);
    this.icoStartDate = this.preIcoStartDate + duration.weeks(1);
    this.icoFinishDate = this.icoStartDate + duration.weeks(1);

    this.crowdsale = await Crowdsale.new(
        this.privateSaleStartDate,
        this.preIcoStartDate,
        this.icoStartDate,
        this.icoFinishDate,
        wallet,
        foundersWallet
    );
    const tokenAddress = await this.crowdsale.token();
    this.token = await Token.at(tokenAddress);
    this.decimals = 10 ** (await this.token.decimals());

    await this.crowdsale.prepareCrowdsale();
    await this.crowdsale.updateWeiInFiat(ether(0.000018));
  });

  describe('Investors list', function () {
    it('Added after buy tokens', async function () {
      await this.crowdsale.addToWhiteList(accounts[0]);
      await this.crowdsale.addToWhiteList(accounts[1]);
      await increaseTimeTo(this.preIcoStartDate);
      await this.crowdsale.sendTransaction({ from: accounts[0], value: ether(0.1) });
      await this.crowdsale.sendTransaction({ from: accounts[0], value: ether(0.2) });
      await this.crowdsale.sendTransaction({ from: accounts[1], value: ether(0.3) });
      const investors = await this.crowdsale.getInvestorsListLength();
      investors.should.be.bignumber.equal(2);
    });
  });

  describe('Buy tokens', function () {
    it('Buying a token share', async function () {
      await this.crowdsale.addToWhiteList(accounts[0]);
      await increaseTimeTo(this.icoStartDate);
      await this.crowdsale.sendTransaction({ from: accounts[0], value: ether(0.18271818) });
      const tokens = await this.crowdsale.tokensOrdered(accounts[0]);
      tokens.should.be.bignumber.equal(10151E6);
    });

    it('Over hard cap', async function () {
      await this.crowdsale.updateWeiInFiat(ether(0.000000018));
      await this.crowdsale.addPrivateParticipant(accounts[0]);
      await this.crowdsale.addToWhiteList(accounts[0]);
      await increaseTimeTo(this.privateSaleStartDate);
      await this.crowdsale.sendTransaction({ from: accounts[0], value: ether(13.5) });
      const tokens = await this.crowdsale.tokensOrdered(accounts[0]);
      tokens.should.be.bignumber.equal(10E14);
      await this.crowdsale.sendTransaction({ from: accounts[0], value: ether(1) }).should.be.rejected;
    });
  });

  describe('Frozen reserve', function () {
    it('Receive', async function () {
      await increaseTimeTo(this.icoFinishDate + duration.years(1) + duration.days(1));
      await this.crowdsale.checkState();
      await this.crowdsale.receiveFrozenReserve(frozeWallet);
      const frozenTokens = (await this.crowdsale.frozenReserve()).toNumber();
      const tokens = await this.token.balanceOf(frozeWallet);
      tokens.should.be.bignumber.equal(frozenTokens * this.decimals);
    });
  });

  describe('Withdraw', function () {
    beforeEach(async function () {
      await this.crowdsale.addPrivateParticipant(accounts[0]);
      await this.crowdsale.addToWhiteList(accounts[0]);
      await increaseTimeTo(this.privateSaleStartDate);
    });

    it('All stages', async function () {
      await this.crowdsale.sendTransaction({ from: accounts[0], value: ether(1) });
      await increaseTimeTo(this.preIcoStartDate);
      let logs = (await this.crowdsale.withdraw()).logs;
      let event = logs.find(e => e.event === 'FundsWithdrawn');
      event.args._amount.should.be.bignumber.equal(ether(1));
      await this.crowdsale.sendTransaction({ from: accounts[0], value: ether(1.1) });
      await increaseTimeTo(this.icoStartDate);
      logs = (await this.crowdsale.withdraw()).logs;
      event = logs.find(e => e.event === 'FundsWithdrawn');
      event.args._amount.should.be.bignumber.equal(ether(1.1));
      await this.crowdsale.sendTransaction({ from: accounts[0], value: ether(0.5) });
      await increaseTimeTo(this.icoFinishDate);
      logs = (await this.crowdsale.withdraw()).logs;
      event = logs.find(e => e.event === 'FundsWithdrawn');
      event.args._amount.should.be.bignumber.equal(ether(0.5));
    });

    it('Separately', async function () {
      await this.crowdsale.sendTransaction({ from: accounts[0], value: ether(0.5) });
      await increaseTimeTo(this.preIcoStartDate);
      await this.crowdsale.sendTransaction({ from: accounts[0], value: ether(0.6) });
      await increaseTimeTo(this.icoStartDate);
      const { logs } = await this.crowdsale.withdraw();
      const events = logs.filter(e => e.event === 'FundsWithdrawn');
      events[0].args._amount.should.be.bignumber.equal(ether(0.5));
      events[1].args._amount.should.be.bignumber.equal(ether(0.6));
    });
  });
});
