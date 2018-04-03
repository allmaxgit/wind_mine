#!/bin/bash

echo "Applying this patch will add GetNetworkInfo RPC request, which exist in BTC node, but doesn't exist in btcd package"

btcd="./vendor/github.com/btcsuite/btcd"
patch=`which patch`

if [[ -e "$patch" ]]
then
    if [ -d "$btcd" ]
    then
        echo "Applying patch..."
        wd=`pwd`
        cd "$btcd"
        cp $wd/btcd.patch .
        patch -p1 < btcd.patch
        rm btcd.patch
    else
        echo "btcd package should be in vendor directory"
    fi
else
    echo "Install patch to apply patch"
fi
