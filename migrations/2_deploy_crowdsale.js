var Crowdsale = artifacts.require("./Crowdsale.sol");

module.exports = function(deployer) {

    let now = new Date();
    now = now.getTime() / 1000;

    deployer.deploy(
        Crowdsale,
        now, // _privateSaleStartDate
        now + 1209600, // _privateSaleDuration
        now + 1209600 + 1209600, // _preIcoDuration
        now + 1209600 + 1209600 + 1209600, // _icoDuration
        web3.eth.accounts[0], // _wallet
        web3.eth.accounts[1], // _foundersWallet
    ).then(() => Crowdsale.deployed()).then((instance) => {
        instance.prepareCrowdsale({ from: web3.eth.accounts[0] });
    });
};
