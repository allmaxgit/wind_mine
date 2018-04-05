import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {Button, Col, Input, Row} from 'reactstrap';

class SmartContractFunction extends Component {
  render() {
    const {
      label, variable, funcChange, funcButton, placeholder, buttonTitle, valueChange
    } = this.props;

    return (
      <Row className="funcRow">
        <Col md={{size: 3}}>
          <Row className="funcLabel">{label}</Row>
        </Col>
        <Col md={{size: 6}}>
          <Input
            value={variable}
            onChange={e => funcChange(valueChange, e.target.value)}
            onKeyDown={this.handleSubmit}
            placeholder={placeholder}
          />
        </Col>
        <Col md={{size: 3}} style={{display: 'flex', justifyContent: 'flex-end'}}>
          <Row>
            <Button
              className="funcButton"
              color="info"
              onClick={() => funcButton()}
            >{buttonTitle}
            </Button>
          </Row>
        </Col>
      </Row>
    );
  }
}

SmartContractFunction.propTypes = {
  label: PropTypes.string.isRequired,
  variable: PropTypes.string.isRequired,
  funcChange: PropTypes.func.isRequired,
  funcButton: PropTypes.func.isRequired,
  placeholder: PropTypes.string.isRequired,
  buttonTitle: PropTypes.string.isRequired,
  valueChange: PropTypes.string.isRequired
};

export default SmartContractFunction;
