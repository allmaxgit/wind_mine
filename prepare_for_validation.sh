#!/bin/bash

 echo "Concatenating contracts for verification..."
 cat ./contracts/ERC20Basic.sol > verification.sol

 tail -n +2 ./contracts/ERC20.sol | sed '/import/d' >> verification.sol

 tail -n +2 ./contracts/SafeMath.sol | sed '/import/d' >> verification.sol

 tail -n +2 ./contracts/BasicToken.sol | sed '/import/d' >> verification.sol

 tail -n +2 ./contracts/StandardToken.sol | sed '/import/d' >> verification.sol

 tail -n +2 ./contracts/Ownable.sol | sed '/import/d' >> verification.sol

 tail -n +2 ./contracts/WindMineToken.sol | sed '/import/d' >> verification.sol

 tail -n +2 ./contracts/UsingFiatPrice.sol | sed '/import/d' >> verification.sol

 tail -n +2 ./contracts/Crowdsale.sol | sed '/import/d' >> verification.sol

 echo "Output: verification.sol"