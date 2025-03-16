#!/bin/bash

KEY="dev0"
CHAINID="sacas_1317-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t sacas-datadir.XXXXX)

echo "create and add new keys"
./sacasd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Sacas with moniker=$MONIKER and chain-id=$CHAINID"
./sacasd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./sacasd add-genesis-account \
"$(./sacasd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000asacas,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./sacasd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./sacasd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./sacasd validate-genesis --home $DATA_DIR

echo "starting sacas node $i in background ..."
./sacasd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started sacas node"
tail -f /dev/null