import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Row, Col, Input, Button, Card, CardTitle, CardText, Alert } from 'reactstrap';

class BTCSupport extends Component {
    constructor(props) {
        super(props);

        this.state = {
            ETHAddress: '',
            BTCAddress: '',
            message: '',
        };

        this.changeValue = this.changeValue.bind(this);
        this.buyTokens = this.buyTokens.bind(this);

    }

    async componentWillMount() {

        const { web3 } = this.props;

        if (web3) {
            const ETHAddress = web3.eth.accounts[0];

            this.setState({ETHAddress});

        }

    }

    async changeValue(name, value) {
        this.setState({
            [name]: value
        });
    }

    async buyTokens() {
        const { web3 } = this.props;
        let { ETHAddress, BTCAddress, message } = this.state;
        let WAValidator = require('wallet-address-validator');

        const isETHAddress = web3.isAddress(ETHAddress);
        const isBTCAddress = WAValidator.validate(BTCAddress, 'BTC');

        if (isETHAddress && isBTCAddress) {

            try {
                let header = new Headers({
                    'Access-Control-Allow-Origin': '*',
                    'Content-Type': 'application/json'
                });

                console.log("------------");
                const resp = await fetch('http://162.213.252.104:9091/v1/getBTCWallet', {
                    method: 'POST',
                    header,
                    body: JSON.stringify({
                        ethAddress: ETHAddress,
//                        btcAddress: BTCAddress
                    }),
                });
                const data = await resp.json();
                console.log('Resp: ', resp);
                console.log('Data: ', data);
                if (data.btcAddress) {
                    message =
                        <div>
                            <Alert color="success">
                                <h4 className="alert-heading">Well done!</h4>
                                <p>Please send BTC on this address:</p>
                                <p>{data.btcAddress}</p>
                            </Alert>
                        </div>;
                }
                if (data.error) {
                    message =
                        <div>
                            <Alert color="danger">
                                <h4 className="alert-heading">Error!</h4>
                                <p>{data.error}</p>
                            </Alert>
                        </div>;
                }
                this.setState({message});
            } catch (err) {
                console.log("failed to fetch BTC address");
            }

        } else {

            if (!isETHAddress) {
                alert(`${ETHAddress} is not ETH address!`)
            }
            if (!isBTCAddress) {
                alert(`${BTCAddress} is not BTC address!`)
            }
        }

    }

    render() {
        const {
            ETHAddress, BTCAddress, message
        } = this.state;

        return (
            <Row className="funcRow" style={{justifyContent: 'center'}}>
                <Col sm="6">
                    {message}
                    <Card body>
                        <CardTitle>Buy Tokens</CardTitle>
                        <CardText>
                            <Input
                                value={ETHAddress}
                                onChange={e => this.changeValue("ETHAddress", e.target.value)}
                                onKeyDown={this.handleSubmit}
                                placeholder="ETH address"
                            />
                            <Input
                                value={BTCAddress}
                                onChange={e => this.changeValue("BTCAddress", e.target.value)}
                                onKeyDown={this.handleSubmit}
                                placeholder="BTC address"
                            />
                        </CardText>
                        <Row style={{justifyContent: 'center'}}>
                            <Button
                            className="funcButton"
                            color="info"
                            onClick={() => this.buyTokens()}
                            >Get My Wallets
                            </Button>
                        </Row>
                    </Card>
                </Col>
            </Row>
        );
    }
}

BTCSupport.propTypes = {
    web3: PropTypes.shape({
        toWei: PropTypes.func.isRequired,
        eth: PropTypes.shape({
            accounts: PropTypes.array.isRequired
        })
    }).isRequired,
};

export default BTCSupport;
/*
            <Row className="funcRow">
                <Col md={{ size: 3 }}>
                    <Row className="funcLabel">Buy Tokens</Row>
                </Col>
                <Col md={{ size: 6 }}>
                    <Input
                        value={ETHAddress}
                        onChange={e => this.changeValue("ETHAddress", e.target.value)}
                        onKeyDown={this.handleSubmit}
                        placeholder="ETH address"
                    />
                    <Input
                        value={BTCAddress}
                        onChange={e => this.changeValue("BTCAddress", e.target.value)}
                        onKeyDown={this.handleSubmit}
                        placeholder="BTC address"
                    />
                </Col>
                <Col md={{ size: 3 }} style={{ display: 'flex', justifyContent: 'flex-end' }}>
                    <Row>
                        <Button
                            className="funcButton"
                            color="info"
                            onClick={() => this.buyTokens()}
                        >Get My Wallets
                        </Button>
                    </Row>
                </Col>
            </Row>
 */
