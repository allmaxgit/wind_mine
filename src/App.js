import 'bootstrap/dist/css/bootstrap.css';
import 'react-bootstrap-table/dist/react-bootstrap-table.min.css';

import React, { Component } from 'react';
import { Container, Row, Col } from 'reactstrap';
import contract from 'truffle-contract';
import './App.css';

import CrowdsaleContract from './contracts/Crowdsale.json';
import TokenContract from './contracts/WindMineToken.json';
import getWeb3 from './utils/getWeb3';

import Token from './components/Token';
import Crowdsale from './components/Crowdsale';
import Statistics from './components/Statistics';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      my: {
        address: ''
      },
      divider: 1,
      decimals: 0
    };
  }

  async componentWillMount() {
    const { web3 } = await getWeb3;

    // Crowdsale initialization
    const crowdsale = contract(CrowdsaleContract);
    crowdsale.setProvider(web3.currentProvider);
    const instanceCrowdsale = await crowdsale.deployed();
    const tokenAddress = await instanceCrowdsale.token.call();

      // Token initialization
    const XBXToken = contract({ abi: TokenContract.abi });
    XBXToken.setProvider(web3.currentProvider);
    const instanceToken = XBXToken.at(tokenAddress);

    // Token Info
    const [
      name, symbol, decimals
    ] = await Promise.all([
      instanceToken.name.call(), instanceToken.symbol.call(), instanceToken.decimals.call()
    ]);
    const decimalsNum = decimals.toNumber();
    const divider = 10 ** decimalsNum;

    // My Info
    const tokens = await instanceToken.balanceOf(web3.eth.accounts[0], { from: web3.eth.accounts[0] });

    this.setState({
      web3,
      instanceToken,
      instanceCrowdsale,

      my: {
        address: web3.eth.accounts[0],
        tokens,
      },

      token: {
        address: instanceToken.address,
        name,
        symbol,
        decimals: decimalsNum
      },

      divider
    });
  }

  render() {
    const {
      my, web3, token, instanceToken, instanceCrowdsale, divider, decimals
    } = this.state;

    if (!web3) return (<div className="wait">Web3 not initialized. Wait pls...</div>);
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">WindMine crowdsale</h1>
        </header>
        <Container>
          <Row>
            <Col>
              <Row style={{ marginTop: 50 }}><h3>My Info</h3></Row>
              <Row>Address: {my.address}</Row>
              <Row>Tokens: {(my.tokens / divider).toLocaleString(undefined, { maximumFractionDigits: decimals })} WMD</Row>
            </Col>
            <Col>
              <Token
                instanceToken={instanceToken}
                token={token}
                divider={divider}
              />
            </Col>
          </Row>
          <Row>
            <Col>
                <Crowdsale
                    web3={web3}
                    instanceCrowdsale={instanceCrowdsale}
                    divider={divider}
                    decimals={token.decimals}
                />
            </Col>
          </Row>
        </Container>
        <footer>Allmax © 2018</footer>
      </div>
    );
  }
}
/*
          <Statistics
              instanceCrowdsale={instanceCrowdsale}
              divider={divider}
          />

 */
export default App;
