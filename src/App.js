import 'bootstrap/dist/css/bootstrap.css';
import 'react-bootstrap-table/dist/react-bootstrap-table.min.css';

import React, { Component } from 'react';
import { Container, Row, Col } from 'reactstrap';
import contract from 'truffle-contract';
import './App.css';

import { TabContent, TabPane, Nav, NavItem, NavLink } from 'reactstrap';
import classnames from 'classnames';

import CrowdsaleContract from './contracts/Crowdsale.json';
import TokenContract from './contracts/WindMineToken.json';
import getWeb3 from './utils/getWeb3';

import BTCSupport from './components/BTCSupport';
import Token from './components/Token';
import Crowdsale from './components/Crowdsale';
import Statistics from './components/Statistics';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      my: {
        address: '',
        tokens: 0,
      },
      divider: 1,
      decimals: 0,
      owner: '',
      web3: null, // will be set after async componentWillMount
    };

    this.toggle = this.toggle.bind(this);
    this.state = {
      activeTab: '1'
    };
  }

  toggle(tab) {
    if (this.state.activeTab !== tab) {
      this.setState({
          activeTab: tab
      });
    }
  }

  async componentWillMount() {
    const { web3 } = await getWeb3;
    // Crowdsale initialization
    if (web3 && web3.isConnected()) {
      const crowdsale = contract(CrowdsaleContract);
      crowdsale.setProvider(web3.currentProvider);
      const instanceCrowdsale = await crowdsale.deployed();
      const tokenAddress = await instanceCrowdsale.token.call();
      // Token initialization
      const WMDToken = contract({ abi: TokenContract.abi });
      WMDToken.setProvider(web3.currentProvider);
      const instanceToken = WMDToken.at(tokenAddress);
      // Token Info
      const [
        name, symbol, decimals
      ] = await Promise.all([
        instanceToken.name.call(), instanceToken.symbol.call(), instanceToken.decimals.call()
      ]);
      const decimalsNum = decimals.toNumber();
      const divider = 10 ** decimalsNum;
      const owner = await instanceCrowdsale.owner.call();
      // My Info

      let address = '';
      let tokens = 0;
      if (web3.eth.accounts.length !== 0) {
        address = web3.eth.accounts[0]
        tokens = await instanceToken.balanceOf(web3.eth.accounts[0]);
      }
      this.setState({
        web3,
        instanceToken,
        instanceCrowdsale,

        my: {
          address,
          tokens
        },

        token: {
          address: instanceToken.address,
          name,
          symbol,
          decimals: decimalsNum
        },

        divider,
        owner
      });
    } else {
      this.setState({
        web3: null,
        instanceToken: null,
        instanceCrowdsale: null,

        my: {
          address: "",
          tokens: 0,
        },

        token: {
          address: "",
          name: "",
          symbol: "",
          decimals: 0
        },

        divider: 1,
        owner: ""
      });
    }
  }
  // getBTCWallet

  render() {
    const {
      my, web3, token, instanceToken, instanceCrowdsale, divider, decimals, owner
    } = this.state;
    let tab2;
    if (web3 && web3.eth.accounts.length > 0) {
      tab2 = (web3.eth.accounts[0] === owner);
    } else {
      tab2 = false;
    }

    let navItem2 =
      <Row style={{ marginTop: 10 }}>
        <Nav tabs>
        <NavItem>
          <NavLink className={classnames({ active: this.state.activeTab === '1' })} onClick={()=> { this.toggle('1'); }} > Buy tokens
          </NavLink>
        </NavItem>
        <NavItem>
          <NavLink className={classnames({ active: this.state.activeTab === '2' })} onClick={()=> { this.toggle('2'); }} > Manager Dashboard
          </NavLink>
        </NavItem>
        </Nav>
      </Row>;

    let tabPane2 =
      <TabPane tabId="2">
        <div className="wait">Web3 not initialized. Wait pls...</div>
      </TabPane>;
    if (web3 && web3.isConnected()) {
      tabPane2 =
        <TabPane tabId="2">
          <Row>
            <Col>
              <Row>
                <h3>My Info</h3>
              </Row>
              <Row>Address: {my.address}</Row>
              <Row>Tokens: {(my.tokens / divider).toLocaleString(undefined, { maximumFractionDigits: decimals })} WMD</Row>
            </Col>
            <Col>
              <Token instanceToken={instanceToken} token={token} divider={divider} />
            </Col>
          </Row>
          <Row>
            <Col>
              <Crowdsale web3={web3} instanceCrowdsale={instanceCrowdsale} divider={divider} decimals={token.decimals} />
              <Statistics instanceCrowdsale={instanceCrowdsale} instanceToken={instanceToken} divider={divider} />
            </Col>
          </Row>
        </TabPane>;
    }

    if (!tab2) {
      navItem2 = "";
      tabPane2 = "";
    }

    return (
      <div className="App">
        <header className="App-header">
            <h1 className="App-title">WindMine crowdsale</h1>
        </header>
        <Container>
          <div>
            {navItem2}
            <TabContent activeTab={this.state.activeTab} style={{ marginTop: 50 }}>
              <TabPane tabId="1">
                <Row>
                  <Col>
                    <BTCSupport web3={web3} />
                  </Col>
                </Row>
              </TabPane>
              {tabPane2}
            </TabContent>
          </div>
        </Container>
        <footer>Allmax © 2018</footer>
      </div>
    );
  }
}
export default App;
