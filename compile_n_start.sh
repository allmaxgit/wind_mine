#!/bin/bash

rm -rf ./build
rm -rf ./src/contracts
#truffle compile
truffle migrate --reset
cp -r ./build/contracts ./src
npm start
