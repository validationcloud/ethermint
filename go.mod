module github.com/validationcloud/ethermint

go 1.21

toolchain go1.21.6

require (
	cosmossdk.io/errors v1.0.1
	cosmossdk.io/math v1.2.0
	github.com/btcsuite/btcd v0.23.4
	github.com/btcsuite/btcd/btcutil v1.1.3
	github.com/cosmos/cosmos-proto v1.0.0-beta.4
	github.com/cosmos/cosmos-sdk v0.50.4
	github.com/cosmos/gogoproto v1.4.11
	github.com/ethereum/go-ethereum v1.10.26
	github.com/stretchr/testify v1.8.4
	github.com/tendermint/tendermint v0.34.27
	github.com/tidwall/gjson v1.14.4
	github.com/tidwall/sjson v1.2.5
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/text v0.14.0
)

replace (
	github.com/tendermint/tendermint => github.com/cometbft/cometbft v0.34.28
)