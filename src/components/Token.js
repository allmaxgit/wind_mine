import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { Row, Col } from 'reactstrap';

class Token extends Component {
    constructor(props) {
        super(props);

        this.state = {
            totalSupply: ''
        };
    }

    async componentWillMount() {
        const { instanceToken } = this.props;
        const [
            owner,
            totalSupply
        ] = await Promise.all([
            instanceToken.owner.call(),
            instanceToken.totalSupply.call()
        ]);

        this.setState({
            owner,
            totalSupply: totalSupply.toNumber()
        });
    }

    render() {
        const {
            address, name, symbol, decimals
        } = this.props.token;
        const { owner, totalSupply } = this.state;

        return (
            <Row>
                <Col>
                    <Row><h3>Token Info</h3></Row>
                    <Row>Address: {address}</Row>
                    <Row>Owner: {owner}</Row>
                    <Row>Name: {name}</Row>
                    <Row>Symbol: {symbol}</Row>
                    <Row>Decimals: {decimals}</Row>
                    <Row>Total Supply: {(totalSupply / this.props.divider).toLocaleString(undefined, { maximumFractionDigits: decimals })} WMD</Row>
                </Col>
            </Row>
        );
    }
}

Token.defaultProps = {
    token: {
        address: '',
        name: '',
        symbol: '',
        decimals: 0
    }
};

Token.propTypes = {
    token: PropTypes.shape({
        address: PropTypes.string,
        name: PropTypes.string,
        symbol: PropTypes.string,
        decimals: PropTypes.number
    }),
    instanceToken: PropTypes.shape({
        owner: PropTypes.func.isRequired,
        totalSupply: PropTypes.func.isRequired,
        balanceOf: PropTypes.func.isRequired
    }).isRequired,
    divider: PropTypes.number.isRequired
};

export default Token;
