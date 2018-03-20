pragma solidity 0.4.21;

import './SafeMath.sol';
import './StandardToken.sol';
import './Ownable.sol';

/**
 * @title WindMineToken
 * @dev WindMineToken contract, which extends OpenZeppelin's StandardToken
 */
contract WindMineToken is StandardToken, Ownable {
    using SafeMath for uint256;

    /**
     * @title decimals
     * @dev Number of decimal places
     */
    uint256 public constant decimals = 8;

    /**
     * @title foundersReserve
     * @dev Amount of tokens, which should be reserved to founder's address
     */
    uint256 public constant foundersReserve = 20000000 * (10 ** decimals);

    /**
     * @title symbol
     * @dev Token symbol
     */
    string public constant symbol = "WMD";

    /**
     * @title name
     * @dev Full token name
     */
    string public constant name = "WindMine";

    /**
     * @title foundersWallet
     * @dev Founder's wallet address, where to send founder's token reserve
     */
    address public foundersWallet;

    /**
     * @dev WindMineToken constructor. Sets founder's wallet from constructor parameter, reserved tokens to it, all the rest
     * are put to the sender's balance.
     * @param _foundersWallet Founder's wallet address where to send founder's reserve of 20 mln tokens
     */
    function WindMineToken(address _foundersWallet) public {
        require(_foundersWallet != address(0));
        totalSupply_ = 160000000 * (10 ** decimals);
        foundersWallet = _foundersWallet;
        balances[msg.sender] = totalSupply_.sub(foundersReserve);
        balances[foundersWallet] = foundersReserve;
    }

}
