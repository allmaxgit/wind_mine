import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Row, Col, Input, Button } from 'reactstrap';

import moment from 'moment';
import '../css/Crowdsale.css';

import SmartContractFunction from './SmartContractFunction';

const dateFormat = 'D.MM.YYYY, HH:mm:ss';

class Crowdsale extends Component {
    constructor(props) {
        super(props);

        this.state = {
            owner: '',
            token: '',

            privateSale: {
                priceInFiatFracture: '',
                hardCap: '',
                startDate: '',
                duration: '',
            },

            preIco: {
                priceInFiatFracture: '',
                hardCap: '',
                startDate: '',
                duration: '',
            },

            ico: {
                priceInFiatFracture: '',
                hardCap: '',
                startDate: '',
                duration: '',
            },

            generalHardCap: '',
            currentHardCap: '',

            tokensSold: '',
            weiRaised: '',
            wallet: '',
            crowdsaleState: '',

            newPrivateParticipant: '',
            newWallet: '',
            newStartDate: '',

            reserveReceiver: '',
            reserveAmount: '',

        };

        this.changeValue = this.changeValue.bind(this);
        this.addPrivateParticipant = this.addPrivateParticipant.bind(this);
        this.setWallet = this.setWallet.bind(this);
        this.setNewStartDate = this.setNewStartDate.bind(this);
        this.manualReserve = this.manualReserve.bind(this);
        this.withdraw = this.withdraw.bind(this);

    }

    async componentWillMount() {
        const { instanceCrowdsale } = this.props;

        const [
            owner,
            token,
            privatePriceInFiatFracture,
            preIcoPriceInFiatFracture,
            icoPriceInFiatFracture,
            generalHardCap,
            currentHardCap,
            privateSaleHardCap,
            preIcoHardCap,
            icoHardCap,
            privateSaleStartDate,
            privateSaleDuration,
            preIcoStartDate,
            preIcoDuration,
            icoStartDate,
            icoDuration,
            tokensSold,
            weiRaised,
            wallet,
            crowdsaleState,
        ] = await Promise.all([
            instanceCrowdsale.owner.call(),
            instanceCrowdsale.token.call(),
            instanceCrowdsale.privatePriceInFiatFracture.call(),
            instanceCrowdsale.preIcoPriceInFiatFracture.call(),
            instanceCrowdsale.icoPriceInFiatFracture.call(),
            instanceCrowdsale.generalHardCap.call(),
            instanceCrowdsale.currentHardCap.call(),
            instanceCrowdsale.privateSaleHardCap.call(),
            instanceCrowdsale.preIcoHardCap.call(),
            instanceCrowdsale.icoHardCap.call(),
            instanceCrowdsale.privateSaleStartDate.call(),
            instanceCrowdsale.privateSaleDuration.call(),
            instanceCrowdsale.preIcoStartDate.call(),
            instanceCrowdsale.preIcoDuration.call(),
            instanceCrowdsale.icoStartDate.call(),
            instanceCrowdsale.icoDuration.call(),
            instanceCrowdsale.tokensSold.call(),
            instanceCrowdsale.weiRaised.call(),
            instanceCrowdsale.wallet.call(),
            instanceCrowdsale.crowdsaleState.call(),
        ]);


        const privateSale = {
            priceInFiatFracture: privatePriceInFiatFracture,
            hardCap: privateSaleHardCap,
            startDate: privateSaleStartDate,
            duration: privateSaleDuration,
        };

        const preIco = {
            priceInFiatFracture: preIcoPriceInFiatFracture,
            hardCap: preIcoHardCap,
            startDate: preIcoStartDate,
            duration: preIcoDuration,
        };

        const ico = {
            priceInFiatFracture: icoPriceInFiatFracture,
            hardCap: icoHardCap,
            startDate: icoStartDate,
            duration: icoDuration,
        };


        this.setState({
            owner,
            token,
            privateSale,
            preIco,
            ico,
            generalHardCap,
            currentHardCap,
            tokensSold,
            weiRaised,
            wallet,
            crowdsaleState,
        });
    }

    splitString = str => str.replace(/\s/g, '').split(',');

    async feasibility(callback, args) {
        try {
            await callback.apply(this, args);
            return true;
        } catch (error) {
            console.log(error);
            return false;
        }
    }

    async changeValue(name, value) {
        this.setState({
            [name]: value
        });
    }

    async addPrivateParticipant() {
        const { web3, instanceCrowdsale } = this.props;
        const { newPrivateParticipant } = this.state;

        const isAddress = web3.isAddress(newPrivateParticipant);

        if (isAddress) {
            await this.feasibility(instanceCrowdsale.addPrivateParticipant, [
                newPrivateParticipant, { from: web3.eth.accounts[0], gas: 300000 }
            ]);
        } else {
            alert(`${newPrivateParticipant} is not address!`);
        }

    }

    async setWallet() {
        const { web3, instanceCrowdsale } = this.props;
        const { newWallet } = this.state;

        const isAddress = web3.isAddress(newWallet);

        if (isAddress) {
            if (await this.feasibility(instanceCrowdsale.setWallet, [
                    newWallet, { from: web3.eth.accounts[0], gas: 300000 }
                ])) {
                const wallet = await instanceCrowdsale.wallet.call();

                this.setState({ wallet });
            }
        } else {
            alert(`${newWallet} is not address!`);
        }

    }

    async setNewStartDate() {
        const { web3, instanceCrowdsale } = this.props;
        const { privateSale, preIco, ico, newStartDate } = this.state;

        let date = moment.unix(newStartDate);

        if (date.isValid()) {
            if (await this.feasibility(instanceCrowdsale.setNewStartDate, [
                    newStartDate, { from: web3.eth.accounts[0], gas: 300000 }
                ])) {
                privateSale.startDate = await instanceCrowdsale.privateSaleStartDate();
                privateSale.duration = await instanceCrowdsale.privateSaleDuration();
                preIco.startDate = await instanceCrowdsale.preIcoStartDate();
                preIco.duration = await instanceCrowdsale.preIcoDuration();
                ico.startDate = await instanceCrowdsale.icoStartDate();
                ico.duration = await instanceCrowdsale.icoDuration();

                this.setState({ privateSale, preIco, ico });
            }
        } else {
            alert(`${newStartDate} is not Unix Timestamp!`);
        }
    }

    async manualReserve() {
        const { web3, instanceCrowdsale } = this.props;
        const { reserveReceiver, reserveAmount } = this.state;

        const isAddress = web3.isAddress(reserveReceiver);
        const isNumber = !isNaN(parseFloat(reserveAmount)) && isFinite(reserveAmount);

        if (isAddress && isNumber) {
            if (await this.feasibility(instanceCrowdsale.manualReserve, [
                    reserveReceiver, reserveAmount, { from: web3.eth.accounts[0], gas: 300000 }
                ])) {
                const tokensSold = await instanceCrowdsale.tokensSold();

                this.setState({ tokensSold });
            }
        } else {
            let errorMsg = '';
            if (!isAddress) {
                errorMsg += `${reserveReceiver} is not address!\n`;
            }
            if (!isNumber) {
                errorMsg += `${reserveAmount} is not number!\n`;
            }
            alert(`Error!\n${errorMsg}`);
        }

    }

    async withdraw() {
        const { web3, instanceCrowdsale } = this.props;

        await this.feasibility(instanceCrowdsale.withdraw, [
            { from: web3.eth.accounts[0], gas: 300000 }
        ]);

    }



    render() {
        const { instanceCrowdsale } = this.props;
        const {
            owner, token, crowdsaleState, wallet, weiRaised,
            privateSale, preIco, ico,
            generalHardCap, currentHardCap, tokensSold,
            newWallet, newPrivateParticipant, newStartDate,
            reserveReceiver, reserveAmount

        } = this.state;
        return (
            <Row style={{ marginTop: 50 }}>
                <Col>
                    <Row><h3>Crowdsale</h3></Row>
                    <Row>Address: {instanceCrowdsale.address}</Row>
                    <Row>Owner: {owner}</Row>
                    <Row>Token: {token}</Row>
                    <Row>Crowdsale State: {crowdsaleState.toString()}</Row>
                    <Row>Wallet: {wallet}</Row>
                    <Row>General Hard Cap: {(generalHardCap / this.props.divider).toLocaleString(undefined, { maximumFractionDigits: this.props.decimals })} WMD</Row>
                    <Row>Current Hard Cap: {(currentHardCap / this.props.divider).toLocaleString(undefined, { maximumFractionDigits: this.props.decimals })} WMD</Row>
                    <Row>Tokens Sold: {(tokensSold / this.props.divider).toLocaleString(undefined, { maximumFractionDigits: this.props.decimals })} WMD</Row>
                    <Row>Wei Raised: {weiRaised.toString()}</Row>
                    <hr className="my-2" />
                    <Row>
                        <Col>
                            <Row><h5>Private Sale</h5></Row>
                            <Row>Start: {moment.unix(privateSale.startDate).format(dateFormat)}</Row>
                            <Row>Duration: {Math.floor(privateSale.duration / 86400).toString()} days</Row>
                            <Row>Price In Fiat Fracture: {privateSale.priceInFiatFracture.toString()}</Row>
                            <Row>Hard Cap: {(privateSale.hardCap / this.props.divider).toLocaleString(undefined, { maximumFractionDigits: this.props.decimals })} WMD</Row>
                        </Col>
                        <Col>
                            <Row><h5>Pre-ICO</h5></Row>
                            <Row>Start: {moment.unix(preIco.startDate).format(dateFormat)}</Row>
                            <Row>Duration: {Math.floor(preIco.duration / 86400).toString()} days</Row>
                            <Row>Price In Fiat Fracture: {preIco.priceInFiatFracture.toString()}</Row>
                            <Row>Hard Cap: {(preIco.hardCap / this.props.divider).toLocaleString(undefined, { maximumFractionDigits: this.props.decimals })} WMD</Row>
                        </Col>
                        <Col>
                            <Row><h5>ICO</h5></Row>
                            <Row>Start: {moment.unix(ico.startDate).format(dateFormat)}</Row>
                            <Row>Duration: {Math.floor(ico.duration / 86400).toString()} days</Row>
                            <Row>Price In Fiat Fracture: {ico.priceInFiatFracture.toString()}</Row>
                            <Row>Hard Cap: {(ico.hardCap / this.props.divider).toLocaleString(undefined, { maximumFractionDigits: this.props.decimals })} WMD</Row>
                        </Col>
                    </Row>
                    <hr className="my-2" />
                    <Row className="funcRow" style={{justifyContent: 'center'}}>
                        <Button className="funcButton" color="info" onClick={() => this.withdraw()}>withdraw</Button>
                    </Row>
                    <hr className="my-2" />
                    <Row><h5>Public functions</h5></Row>
                    <SmartContractFunction
                        label="Add new Private Participant"
                        variable={newPrivateParticipant}
                        valueChange="newPrivateParticipant"
                        funcChange={this.changeValue}
                        buttonTitle="addPrivateParticipant"
                        funcButton={this.addPrivateParticipant}
                        placeholder="Address new Private Participant"
                    />
                    <SmartContractFunction
                        label="Set new wallet"
                        variable={newWallet}
                        valueChange="newWallet"
                        funcChange={this.changeValue}
                        buttonTitle="setWallet"
                        funcButton={this.setWallet}
                        placeholder={wallet}
                    />
                    <SmartContractFunction
                        label="Set new Start Date"
                        variable={newStartDate}
                        valueChange="newStartDate"
                        funcChange={this.changeValue}
                        buttonTitle="setNewStartDate"
                        funcButton={this.setNewStartDate}
                        placeholder={privateSale.startDate.toString()}
                    />
                    <Row className="funcRow">
                        <Col md={{ size: 3 }}>
                            <Row className="funcLabel">Manual Reserve</Row>
                        </Col>
                        <Col md={{ size: 6 }}>
                            <Input
                                value={reserveReceiver}
                                onChange={e => this.changeValue("reserveReceiver", e.target.value)}
                                onKeyDown={this.handleSubmit}
                                placeholder="reserveReceiver"
                            />
                            <Input
                                value={reserveAmount}
                                onChange={e => this.changeValue("reserveAmount", e.target.value)}
                                onKeyDown={this.handleSubmit}
                                placeholder="reserveAmount"
                            />
                        </Col>
                        <Col md={{ size: 3 }} style={{ display: 'flex', justifyContent: 'flex-end' }}>
                            <Row>
                                <Button
                                    className="funcButton"
                                    color="info"
                                    onClick={() => this.manualReserve()}
                                >manualReserve
                                </Button>
                            </Row>
                        </Col>
                    </Row>
                </Col>
            </Row>
        );
    }
}

Crowdsale.propTypes = {
    web3: PropTypes.shape({
        toWei: PropTypes.func.isRequired,
        eth: PropTypes.shape({
            accounts: PropTypes.array.isRequired
        })
    }).isRequired,
    instanceCrowdsale: PropTypes.shape({
        address: PropTypes.string.isRequired,

        // Info
        owner: PropTypes.func.isRequired,
        token: PropTypes.func.isRequired,

        privatePriceInFiatFracture: PropTypes.func.isRequired,
        preIcoPriceInFiatFracture: PropTypes.func.isRequired,
        icoPriceInFiatFracture: PropTypes.func.isRequired,

        generalHardCap: PropTypes.func.isRequired,
        currentHardCap: PropTypes.func.isRequired,
        privateSaleHardCap: PropTypes.func.isRequired,
        preIcoHardCap: PropTypes.func.isRequired,
        icoHardCap: PropTypes.func.isRequired,

        privateSaleStartDate: PropTypes.func.isRequired,
        privateSaleDuration: PropTypes.func.isRequired,
        preIcoStartDate: PropTypes.func.isRequired,
        preIcoDuration: PropTypes.func.isRequired,
        icoStartDate: PropTypes.func.isRequired,
        icoDuration: PropTypes.func.isRequired,

        tokensSold: PropTypes.func.isRequired,
        weiRaised: PropTypes.func.isRequired,
        wallet: PropTypes.func.isRequired,
        crowdsaleState: PropTypes.func.isRequired,

        addPrivateParticipant: PropTypes.func.isRequired,
        setWallet: PropTypes.func.isRequired,
        setNewStartDate: PropTypes.func.isRequired,
        manualReserve: PropTypes.func.isRequired,
        withdraw: PropTypes.func.isRequired,

    }).isRequired,
    divider: PropTypes.number.isRequired,
    decimals: PropTypes.number.isRequired
};

export default Crowdsale;
