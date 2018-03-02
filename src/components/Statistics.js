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
    const { instanceCrowdsale } = this.props;

    // Buyers
    const buyers = [];
    let currentBuyer;
    let address;
    const numberOfBuyers = (await instanceCrowdsale.getNumberOfBuyers.call()).toNumber();
    if (numberOfBuyers > 0) {
      for (let index = 0; index < numberOfBuyers; index += 1) {
        // eslint-disable-next-line
        address = await instanceCrowdsale.buyersList.call(index);
        // eslint-disable-next-line
          currentBuyer = await instanceCrowdsale.salesBook.call(address);
        buyers.push({
          address,
          investments: currentBuyer[0].toNumber()  / 10 ** 18,
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
                <TableHeaderColumn dataField="investments">Total investments, ETH</TableHeaderColumn>
                <TableHeaderColumn dataField="amount">Total amount, WMD</TableHeaderColumn>
              </BootstrapTable>
            </Col>
          </Row>
        </Col>
      </Row>
    );
  }
}

Statistics.propTypes = {
  instanceCrowdsale: PropTypes.shape({
    // Buyers
    getNumberOfBuyers: PropTypes.func.isRequired,
    buyersList: PropTypes.func.isRequired,
    salesBook: PropTypes.func.isRequired,
  }).isRequired
};

export default Statistics;
