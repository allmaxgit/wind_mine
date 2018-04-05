import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {Alert, Button, Card, CardText, CardTitle, Col, Input, Row} from 'reactstrap';

class BTCSupport extends Component {
  constructor(props) {
    super(props);

    this.state = {
      ETHAddress: '',
      BTCAddress: '',
      message: '',

      sendBtnStateETH: false,
      sendBtnStateBTC: false,
    };

    this.buttonStateChange = this.buttonStateChange.bind(this);
    this.isETHAddress = this.isETHAddress.bind(this);
    this.isBTCAddress = this.isBTCAddress.bind(this);
    this.buyTokens = this.buyTokens.bind(this);

  }

  componentWillReceiveProps(props) {
    const {web3} = props;
    if (web3 && web3.isConnected() && web3.eth.accounts.length > 0) {
      const ETHAddress = web3.eth.accounts[0];
      const sendBtnStateETH = true;
      this.setState({
        ETHAddress,
        sendBtnStateETH
      });
    } else {
      this.setState({
        ETHAddress: '',
        sendBtnStateETH: true
      });
    }
  }

  async isETHAddress(address) {
    const {web3} = this.props;
    return web3.isAddress(address);
  }

  async isBTCAddress(address) {
    let WAValidator = require('wallet-address-validator');
    return WAValidator.validate(address, 'BTC', 'testnet'); //TODO remove when going live
  }

  async buttonStateChange(param, value, button, func) {
    const state = await func(value);

    this.setState({
      [param]: value,
      [button]: state
    });
  }

  async buyTokens() {
    let {ETHAddress, BTCAddress, message} = this.state;

    try {
      let header = new Headers({
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json'
      });

      const resp = await fetch('http://162.213.252.104:9091/v1/getBTCWallet', {
        method: 'POST',
        header,
        body: JSON.stringify({
          ethAddress: ETHAddress,
          btcAddress: BTCAddress
        }),
      });
      const data = await resp.json();
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

  }

  render() {
    const {
      ETHAddress, BTCAddress, message, sendBtnStateETH, sendBtnStateBTC
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
                onChange={e => this.buttonStateChange("ETHAddress", e.target.value, "sendBtnStateETH", this.isETHAddress)}
                onKeyDown={this.handleSubmit}
                placeholder="ETH address"
              />
              <Input
                value={BTCAddress}
                onChange={e => this.buttonStateChange("BTCAddress", e.target.value, "sendBtnStateBTC", this.isBTCAddress)}
                onKeyDown={this.handleSubmit}
                placeholder="BTC address"
              />
            </CardText>
            <Row style={{justifyContent: 'center'}}>
              <Button
                className="funcButton"
                color="info"
                onClick={() => this.buyTokens()}
                disabled={!(sendBtnStateETH && sendBtnStateBTC)}
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
  }),
};

export default BTCSupport;
