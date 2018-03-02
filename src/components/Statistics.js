import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Row, Col } from 'reactstrap';
import { BootstrapTable, TableHeaderColumn } from 'react-bootstrap-table';


class Statistics extends Component {
  constructor(props) {
    super(props);

    this.state = {
      buyers: []
    };
  }
  async componentWillMount() {
    const { instanceCrowdsale, instanceToken } = this.props;

    // Buyers
    const buyers = [];
    let currentBuyer;
    let address;
    console.log(instanceCrowdsale);
    const numberOfBuyers = (await instanceCrowdsale.getInvestorsListLength.call()).toNumber();
    if (numberOfBuyers > 0) {
      for (let index = 0; index < numberOfBuyers; index += 1) {
        // eslint-disable-next-line
        address = await instanceCrowdsale.investorsList.call(index);
        // eslint-disable-next-line
        currentBuyer = await instanceToken.balanceOf.call(address);
        buyers.push({
          address,
//          investments: currentBuyer[0].toNumber()  / 10 ** 18,
          amount: currentBuyer[1].toNumber() / this.props.divider
        });
      }
    }
    this.setState({ buyers });
  }

  render() {
    const { buyers } = this.state;

    return (
      <Row style={{ marginTop: 50 }}>
        <Col>
          <Row><h3>Statistics</h3></Row>
          <Row>
            <Col>
              <BootstrapTable data={buyers} version="4" striped hover>
                <TableHeaderColumn dataField="address" isKey>Address</TableHeaderColumn>
                <TableHeaderColumn dataField="amount">Total amount, WMD</TableHeaderColumn>
              </BootstrapTable>
            </Col>
          </Row>
        </Col>
      </Row>
    );
  }
}
/*
                <TableHeaderColumn dataField="investments">Total investments, ETH</TableHeaderColumn>

 */
Statistics.propTypes = {
    instanceToken: PropTypes.shape({
//        owner: PropTypes.func.isRequired,
//        totalSupply: PropTypes.func.isRequired,
        balanceOf: PropTypes.func.isRequired
    }).isRequired,
  instanceCrowdsale: PropTypes.shape({
    // Buyers
    getInvestorsListLength: PropTypes.func.isRequired,
    investorsList: PropTypes.func.isRequired,
  }).isRequired
};

export default Statistics;
