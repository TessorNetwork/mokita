
#!/bin/bash
# microtick and bitcanna contributed significantly here.
# rocksdb doesn't work yet

# PRINT EVERY COMMAND
set -ux

# uncomment the three lines below to build mokita

# export GOPATH=~/go
# export PATH=$PATH:~/go/bin
# go install -./...


# MAKE HOME FOLDER AND GET GENESIS
mokitad init test 
wget -O ~/.mokitad/config/genesis.json https://cloudflare-ipfs.com/ipfs/QmXRvBT3hgoXwwPqbK6a2sXUuArGM8wPyo1ybskyyUwUxs

INTERVAL=1500

# GET TRUST HASH AND TRUST HEIGHT

LATEST_HEIGHT=$(curl -s https://mokita.validator.network/block | jq -r .result.block.header.height);
BLOCK_HEIGHT=$(($LATEST_HEIGHT-$INTERVAL))
TRUST_HASH=$(curl -s "https://mokita.validator.network/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)


# TELL USER WHAT WE ARE DOING
echo "TRUST HEIGHT: $BLOCK_HEIGHT"
echo "TRUST HASH: $TRUST_HASH"


# export state sync vars
export MOKITAD_P2P_MAX_NUM_OUTBOUND_PEERS=200
export MOKITAD_STATESYNC_ENABLE=true
export MOKITAD_STATESYNC_RPC_SERVERS="https://mokita.validator.network:443,https://rpc.mokita.notional.ventures:443,https://rpc-mokita.ecostake.com:443"
export MOKITAD_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export MOKITAD_STATESYNC_TRUST_HASH=$TRUST_HASH

# THIS WILL FAIL BECAUSE THE APP VERSION IS CORRECTLY SET IN MOKITA
mokitad start 

# THIS WILL FIX THE APP VERSION, contributed by callum and claimens
git clone https://github.com/tendermint/tendermint
cd tendermint
git checkout remotes/origin/callum/app-version
go install ./...
tendermint set-app-version 1 --home ~/.mokitad

# THERE, NOW IT'S SYNCED AND YOU CAN PLAY
mokitad start 
