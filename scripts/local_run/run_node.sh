#!/bin/bash

# This script is used to run a node locally.
# 1. This script will generate a genesis,
# 2. Replace all denoms at genesis.json "stake, tiov, uiov, iov and tvoi" with the given token name. (default: stake)
# 3. Start a node with the given name. (default: node0)

# Notes:
# This script will create three accounts:
# - Validator account: the account used to sign blocks and transactions.
# - Faucet account: the account used to send tokens to other accounts.
# - User account: the account used to receive tokens from faucet account.
# The default password for all accounts is "12345678".
# This chain also has a second token called "tusd"

# Usage:

# bash ./scripts/local_run/run_node.sh [folder_path] [token_name] [node_name]

# Variables:

FOLDER_PATH=${1:-"./tmp/node0"}
DEFAULT_DENOM=${2:-"stake"}
DEFAULT_DENOM_SECONDARY="tusd"
NODE_NAME=${3:-"node0"}
CHAIN_ID="localnet"

# Generate genesis:

echo "Generating genesis..."
rm -rf $FOLDER_PATH
mkdir -p $FOLDER_PATH
./build/starnamed init $NODE_NAME --chain-id $CHAIN_ID --home $FOLDER_PATH

# Generate keys:

for account in validator faucet user; do
  echo "Generating $account key..."
  ./build/starnamed keys add $account --home $FOLDER_PATH --keyring-backend test
done

# Add genesis accounts:

echo "Adding genesis accounts..."
./build/starnamed add-genesis-account $(./build/starnamed keys show validator --home $FOLDER_PATH -a --keyring-backend test) 10000000000000$DEFAULT_DENOM,100000000000$DEFAULT_DENOM_SECONDARY --home $FOLDER_PATH

for account in faucet user; do
  ./build/starnamed add-genesis-account $(./build/starnamed keys show $account --home $FOLDER_PATH -a --keyring-backend test) 100000000000$DEFAULT_DENOM,100000000000$DEFAULT_DENOM_SECONDARY --home $FOLDER_PATH
done

# Create gentx:

echo "Creating gentx..."
./build/starnamed gentx validator 1000000000$DEFAULT_DENOM --home $FOLDER_PATH --keyring-backend test --chain-id $CHAIN_ID

# Collect gentx:

echo "Collecting gentx..."
./build/starnamed collect-gentxs --home $FOLDER_PATH

# Replace all denoms:

echo "Replacing all denoms..."
for denom in stake tiov uiov iov tvoi; do
  sed -i "s/$denom/$DEFAULT_DENOM/g" $FOLDER_PATH/config/genesis.json
done

# Enable custom denom: replave "custom_denom_accepted": [] with "custom_denom_accepted": ["$DEFAULT_DENOM_SECONDARY"]

echo "Enabling custom denom..."
sed -i "s/\"custom_denom_accepted\": \[\]/\"custom_denom_accepted\": \[\"$DEFAULT_DENOM_SECONDARY\"\]/g" $FOLDER_PATH/config/genesis.json

# Start node:

echo "Starting node..."
./build/starnamed start --home $FOLDER_PATH

