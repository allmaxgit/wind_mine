pragma solidity 0.4.21;

import './Ownable.sol';

/**
 * @title UsingFiatPrice
 * @dev General contract for updating ETH to fiat currency exchange rate from external source
*/
contract UsingFiatPrice is Ownable {

    /**
     * @title fiatSymbol
     * @dev Fiat currency symbol for informational purposes only
     * @notice Once it is set, it cannot be changed without redeploying contract
    */
    string public fiatSymbol;

    /**
     * @title fiatDecimals
     * @dev Number of decimal places in a used currency. Can be used for calculations or for informational purposes only
    */
    uint256 public fiatDecimals;

    /**
     * @title weiInFiat
     * @dev Amount of wei in minimum fracture of the fiat currency
     * @notice Intended to be used for calculating number of bought tokens with token price set in fiat currency
     */
    uint256 public weiInFiat;

    /**
     * @title lastUpdated
     * @dev Unix timestamp which shows last rate update time
     */
    uint256 public lastUpdated;

    /**
     * @dev This event is raised when exchange rate is updated
     * @param _oldRate Old exchange rate - before update
     * @param _newRate New exchange rate - after update
     */
    event RateUpdated(uint256 _oldRate, uint256 _newRate);

    /**
     * @dev UsingFiatPrice constructor
     * @param _fiatSymbol Fiat currency symbol for informational purposes only
     * @param _fiatDecimals Number of decimal places in a used fiat currency. E.g. _fiatDecimals is 2 for USD - minimal fracture of USD is 0.01 USD
     */
    function UsingFiatPrice(string _fiatSymbol, uint256 _fiatDecimals) public {
        require(_fiatDecimals >= 2 && _fiatDecimals <= 18);
        fiatSymbol = _fiatSymbol;
        fiatDecimals = _fiatDecimals;
    }

    /**
     * @dev Updates wei to fiat currency exchange rate
     * @param _weiInFiat Number of wei in minimum fracture of "fiatSymbol". E.g. Number of wei in 0.01 USD
     */
    function updateWeiInFiat(uint256 _weiInFiat) external onlyOwner {
        require(_weiInFiat > 0);
        emit RateUpdated(weiInFiat, _weiInFiat);
        weiInFiat = _weiInFiat;
        lastUpdated = now;
    }
}
