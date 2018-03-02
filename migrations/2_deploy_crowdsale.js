var Crowdsale = artifacts.require("./Crowdsale.sol");

module.exports = function(deployer) {

    let now = new Date();

    deployer.deploy(
        Crowdsale,
        now.getTime() / 1000, // _privateSaleStartDate
        "1209600", // _privateSaleDuration
        "1209600", // _preIcoDuration
        "1209600", // _icoDuration
        web3.eth.accounts[0], // _wallet
        web3.eth.accounts[1], // _foundersWallet
    ).then(() => Crowdsale.deployed()).then((instance) => {
        instance.prepareCrowdsale({ from: web3.eth.accounts[0] });
    });
};
