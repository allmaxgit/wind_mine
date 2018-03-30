/*
    FOR TESTERS:
    I used with ganache-cli -e 1000000000
    TODO: whitelist, withdraw, dates
    Deploy -> PrepareCrowdsale
 */
/* eslint-disable */
import ether from './helpers/ether';
import { advanceBlock } from './helpers/advanceToBlock';
import { increaseTimeTo, duration } from './helpers/increaseTime';
import latestTime from './helpers/latestTime';
import EVMRevert from './helpers/EVMRevert';
import moment from 'moment';

const BigNumber = web3.BigNumber;
const should = require('chai')
    .use(require('chai-as-promised'))
    .use(require('chai-bignumber')(BigNumber))
    .should();

const Crowdsale = artifacts.require('./Crowdsale.sol');
const Token = artifacts.require('./WindMineToken.sol');

contract('Crowdsale', function ([owner, wallet, foundersWallet, ...accounts]) {
    /*
    (uint256 , uint256 , uint256 , uint256 , address _wallet, address _foundersWallet)
     */
    let privateSaleStartDate, preIcoStartDate, icoStartDate, icoFinishDate;

    before(async function () {
        // Advance to the next block to correctly read time in the solidity "now" function interpreted by testrpc
        await advanceBlock();
    });

    beforeEach(async function () {
        privateSaleStartDate = latestTime();
        preIcoStartDate = privateSaleStartDate + 1209600;
        icoStartDate = preIcoStartDate + 1209600;
        icoFinishDate = icoStartDate + 1209600;
        this.crowdsale = await Crowdsale.new(
            privateSaleStartDate,
            preIcoStartDate,
            icoStartDate,
            icoFinishDate,
            web3.eth.accounts[0],
            web3.eth.accounts[1],
        );
        const tokenAddress = await this.crowdsale.token();
        this.token = await Token.at(tokenAddress);
        this.decimals = 10 ** (await this.token.decimals());
        await this.crowdsale.prepareCrowdsale();
        await this.crowdsale.updateWeiInFiat(1e16);
    });

    describe('start', function () {
        beforeEach(async function () {
        });

        it('Default settings', async function () {

            const [
                deployer,
                privateSaleHardCap,
                preIcoHardCap,
                icoHardCap,
                generalHardCap,
                privatePriceInFiatFracture,
                preIcoPriceInFiatFracture,
                icoPriceInFiatFracture,
                totalSupply,
            ] = await Promise.all([
                await this.crowdsale.owner(),
                await this.crowdsale.privateSaleHardCap(),
                await this.crowdsale.preIcoHardCap(),
                await this.crowdsale.icoHardCap(),
                await this.crowdsale.generalHardCap(),
                await this.crowdsale.privatePriceInFiatFracture(),
                await this.crowdsale.preIcoPriceInFiatFracture(),
                await this.crowdsale.icoPriceInFiatFracture(),
                await this.token.totalSupply(),
            ]);

            let [
                fiatDecimals,
            ] = await Promise.all([
                await this.crowdsale.fiatDecimals(),
            ]);

            fiatDecimals = 10 ** +fiatDecimals;

            deployer.should.be.equal(owner);
            privateSaleHardCap.should.be.bignumber.equal(10000000 * this.decimals);
            preIcoHardCap.should.be.bignumber.equal(privateSaleHardCap.add(20000000 * this.decimals));
            icoHardCap.should.be.bignumber.equal(preIcoHardCap.add(30000000 * this.decimals));
            generalHardCap.should.be.bignumber.equal(totalSupply);
            privatePriceInFiatFracture.should.be.bignumber.equal(new BigNumber(25 / 100 * +fiatDecimals));
            preIcoPriceInFiatFracture.should.be.bignumber.equal(new BigNumber(50 / 100 * +fiatDecimals));
            icoPriceInFiatFracture.should.be.bignumber.equal(new BigNumber(+fiatDecimals));

        });
    });

    describe('Whitelist checking', function () {
        beforeEach(async function () {
        });

        it('Not in whitelist, withdraw', async function () {
            let buyer = accounts[0];
            await this.crowdsale.checkState();
            let crowdsaleState = await this.crowdsale.crowdsaleState();
            let value = ether(1);
            crowdsaleState.should.be.bignumber.equal(new BigNumber(1));
            await this.crowdsale.sendTransaction({ value: value, from: buyer }).should.be.rejected;
            await this.crowdsale.addToWhiteList(buyer, { from: owner });
            await this.crowdsale.sendTransaction({ value: value, from: buyer }).should.be.rejected;
            await this.crowdsale.addPrivateParticipant(buyer, { from: owner });
            await this.crowdsale.sendTransaction({ value: value, from: buyer }).should.be.fulfilled;

            await increaseTimeTo(preIcoStartDate + duration.minutes(1));
            buyer = accounts[1];
            await this.crowdsale.checkState();
            crowdsaleState = await this.crowdsale.crowdsaleState();

            let withdrawLogs = (await this.crowdsale.withdraw()).logs;
            if (!withdrawLogs) {
                console.log("Is no withdraw logs here!");
            }
            let event = withdrawLogs.find(e => e.event === 'FundsWithdrawn');
            event.args._amount.should.be.bignumber.equal(value);
            await this.crowdsale.withdraw().should.be.rejected;

            value = ether(2);
            crowdsaleState.should.be.bignumber.equal(new BigNumber(2));
            await this.crowdsale.sendTransaction({ value: value, from: buyer }).should.be.rejected;
            await this.crowdsale.addToWhiteList(buyer, { from: owner });
            await this.crowdsale.sendTransaction({ value: value, from: buyer }).should.be.fulfilled;

            await increaseTimeTo(icoStartDate + duration.minutes(1));
            buyer = accounts[2];
            await this.crowdsale.checkState();
            crowdsaleState = await this.crowdsale.crowdsaleState();

            withdrawLogs = (await this.crowdsale.withdraw()).logs;
            if (!withdrawLogs) {
                console.log("Is no withdraw logs here!");
            }
            event = withdrawLogs.find(e => e.event === 'FundsWithdrawn');
            event.args._amount.should.be.bignumber.equal(value);
            await this.crowdsale.withdraw().should.be.rejected;

            value = ether(3);
            crowdsaleState.should.be.bignumber.equal(new BigNumber(3));
            await this.crowdsale.sendTransaction({ value: value, from: buyer }).should.be.rejected;
            await this.crowdsale.addToWhiteList(buyer, { from: owner });
            await this.crowdsale.sendTransaction({ value: value, from: buyer }).should.be.fulfilled;

            await increaseTimeTo(icoFinishDate + duration.minutes(1));
            buyer = accounts[3];
            await this.crowdsale.checkState();
            crowdsaleState = await this.crowdsale.crowdsaleState();

            withdrawLogs = (await this.crowdsale.withdraw()).logs;
            if (!withdrawLogs) {
                console.log("Is no withdraw logs here!");
            }
            event = withdrawLogs.find(e => e.event === 'FundsWithdrawn');
            event.args._amount.should.be.bignumber.equal(value);
            await this.crowdsale.withdraw().should.be.rejected;

            value = ether(4);
            crowdsaleState.should.be.bignumber.equal(new BigNumber(4));
            await this.crowdsale.sendTransaction({ value: value, from: buyer }).should.be.rejected;
            await this.crowdsale.addToWhiteList(buyer, { from: owner });
            await this.crowdsale.sendTransaction({ value: value, from: buyer }).should.be.rejected;
        });
    });

    describe('Dates checking', function () {
        beforeEach(async function () {
        });

        it('Date moving', async function () {
            await this.crowdsale.checkState();
            let crowdsaleState = await this.crowdsale.crowdsaleState();
            crowdsaleState.should.be.bignumber.equal(new BigNumber(1));
            await this.crowdsale.setPrivateSaleStartDate(latestTime() + duration.minutes(1)).should.be.rejected;
            await this.crowdsale.setPreIcoStartDate(privateSaleStartDate - duration.minutes(1)).should.be.rejected;
            await this.crowdsale.setIcoStartDate(preIcoStartDate - duration.minutes(1)).should.be.rejected;
            await this.crowdsale.setPreIcoStartDate(preIcoStartDate - duration.minutes(1));
            await this.crowdsale.setIcoStartDate(icoStartDate - duration.minutes(1));
            await this.crowdsale.setIcoFinishDate(icoFinishDate - duration.minutes(1));

            await increaseTimeTo(preIcoStartDate + duration.minutes(1));
            await this.crowdsale.checkState();
            crowdsaleState = await this.crowdsale.crowdsaleState();
            crowdsaleState.should.be.bignumber.equal(new BigNumber(2));

            await increaseTimeTo(icoStartDate + duration.minutes(1));
            await this.crowdsale.checkState();
            crowdsaleState = await this.crowdsale.crowdsaleState();
            crowdsaleState.should.be.bignumber.equal(new BigNumber(3));

            await increaseTimeTo(icoFinishDate + duration.minutes(1));
            await this.crowdsale.checkState();
            crowdsaleState = await this.crowdsale.crowdsaleState();
            crowdsaleState.should.be.bignumber.equal(new BigNumber(4));
        });
    });

});
