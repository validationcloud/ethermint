// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/evmos/ethermint/blob/main/LICENSE
package types

import (
	"math/big"
	"strings"

	sdkmath "cosmossdk.io/math"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

// EthereumConfig returns an Ethereum ChainConfig for EVM state transitions.
// All the negative or nil values are converted to nil
func (cc ChainConfig) EthereumConfig(chainID *big.Int) *params.ChainConfig {
	return &params.ChainConfig{
		ChainID:                 chainID,
		HomesteadBlock:          getBlockValue(&cc.HomesteadBlock.Int),
		DAOForkBlock:            getBlockValue(&cc.DAOForkBlock.Int),
		DAOForkSupport:          cc.DAOForkSupport,
		EIP150Block:             getBlockValue(&cc.EIP150Block.Int),
		EIP150Hash:              common.HexToHash(cc.EIP150Hash),
		EIP155Block:             getBlockValue(&cc.EIP155Block.Int),
		EIP158Block:             getBlockValue(&cc.EIP158Block.Int),
		ByzantiumBlock:          getBlockValue(&cc.ByzantiumBlock.Int),
		ConstantinopleBlock:     getBlockValue(&cc.ConstantinopleBlock.Int),
		PetersburgBlock:         getBlockValue(&cc.PetersburgBlock.Int),
		IstanbulBlock:           getBlockValue(&cc.IstanbulBlock.Int),
		MuirGlacierBlock:        getBlockValue(&cc.MuirGlacierBlock.Int),
		BerlinBlock:             getBlockValue(&cc.BerlinBlock.Int),
		LondonBlock:             getBlockValue(&cc.LondonBlock.Int),
		ArrowGlacierBlock:       getBlockValue(&cc.ArrowGlacierBlock.Int),
		GrayGlacierBlock:        getBlockValue(&cc.GrayGlacierBlock.Int),
		MergeNetsplitBlock:      getBlockValue(&cc.MergeNetsplitBlock.Int),
		ShanghaiBlock:           getBlockValue(&cc.ShanghaiBlock.Int),
		CancunBlock:             getBlockValue(&cc.CancunBlock.Int),
		TerminalTotalDifficulty: nil,
		Ethash:                  nil,
		Clique:                  nil,
	}
}

// DefaultChainConfig returns default evm parameters.
func DefaultChainConfig() ChainConfig {
	homesteadBlock := sdk.ZeroInt()
	daoForkBlock := sdk.ZeroInt()
	eip150Block := sdk.ZeroInt()
	eip155Block := sdk.ZeroInt()
	eip158Block := sdk.ZeroInt()
	byzantiumBlock := sdk.ZeroInt()
	constantinopleBlock := sdk.ZeroInt()
	petersburgBlock := sdk.ZeroInt()
	istanbulBlock := sdk.ZeroInt()
	muirGlacierBlock := sdk.ZeroInt()
	berlinBlock := sdk.ZeroInt()
	londonBlock := sdk.ZeroInt()
	arrowGlacierBlock := sdk.ZeroInt()
	grayGlacierBlock := sdk.ZeroInt()
	mergeNetsplitBlock := sdk.ZeroInt()
	shanghaiBlock := sdk.ZeroInt()
	cancunBlock := sdk.ZeroInt()

	return ChainConfig{
		HomesteadBlock:      &homesteadBlock,
		DAOForkBlock:        &daoForkBlock,
		DAOForkSupport:      true,
		EIP150Block:         &eip150Block,
		EIP150Hash:          common.Hash{}.String(),
		EIP155Block:         &eip155Block,
		EIP158Block:         &eip158Block,
		ByzantiumBlock:      &byzantiumBlock,
		ConstantinopleBlock: &constantinopleBlock,
		PetersburgBlock:     &petersburgBlock,
		IstanbulBlock:       &istanbulBlock,
		MuirGlacierBlock:    &muirGlacierBlock,
		BerlinBlock:         &berlinBlock,
		LondonBlock:         &londonBlock,
		ArrowGlacierBlock:   &arrowGlacierBlock,
		GrayGlacierBlock:    &grayGlacierBlock,
		MergeNetsplitBlock:  &mergeNetsplitBlock,
		ShanghaiBlock:       &shanghaiBlock,
		CancunBlock:         &cancunBlock,
	}
}

func getBlockValue(block *sdkmath.Int) *big.Int {
	if block == nil || block.IsNegative() {
		return nil
	}

	return block.BigInt()
}

// Validate performs a basic validation of the ChainConfig params. The function will return an error
// if any of the block values is uninitialized (i.e nil) or if the EIP150Hash is an invalid hash.
func (cc ChainConfig) Validate() error {
	if err := validateBlock(&cc.HomesteadBlock.Int); err != nil {
		return errorsmod.Wrap(err, "homesteadBlock")
	}
	if err := validateBlock(&cc.DAOForkBlock.Int); err != nil {
		return errorsmod.Wrap(err, "daoForkBlock")
	}
	if err := validateBlock(&cc.EIP150Block.Int); err != nil {
		return errorsmod.Wrap(err, "eip150Block")
	}
	if err := validateHash(cc.EIP150Hash); err != nil {
		return err
	}
	if err := validateBlock(&cc.EIP155Block.Int); err != nil {
		return errorsmod.Wrap(err, "eip155Block")
	}
	if err := validateBlock(&cc.EIP158Block.Int); err != nil {
		return errorsmod.Wrap(err, "eip158Block")
	}
	if err := validateBlock(&cc.ByzantiumBlock.Int); err != nil {
		return errorsmod.Wrap(err, "byzantiumBlock")
	}
	if err := validateBlock(&cc.ConstantinopleBlock.Int); err != nil {
		return errorsmod.Wrap(err, "constantinopleBlock")
	}
	if err := validateBlock(&cc.PetersburgBlock.Int); err != nil {
		return errorsmod.Wrap(err, "petersburgBlock")
	}
	if err := validateBlock(&cc.IstanbulBlock.Int); err != nil {
		return errorsmod.Wrap(err, "istanbulBlock")
	}
	if err := validateBlock(&cc.MuirGlacierBlock.Int); err != nil {
		return errorsmod.Wrap(err, "muirGlacierBlock")
	}
	if err := validateBlock(&cc.BerlinBlock.Int); err != nil {
		return errorsmod.Wrap(err, "berlinBlock")
	}
	if err := validateBlock(&cc.LondonBlock.Int); err != nil {
		return errorsmod.Wrap(err, "londonBlock")
	}
	if err := validateBlock(&cc.ArrowGlacierBlock.Int); err != nil {
		return errorsmod.Wrap(err, "arrowGlacierBlock")
	}
	if err := validateBlock(&cc.GrayGlacierBlock.Int); err != nil {
		return errorsmod.Wrap(err, "GrayGlacierBlock")
	}
	if err := validateBlock(&cc.MergeNetsplitBlock.Int); err != nil {
		return errorsmod.Wrap(err, "MergeNetsplitBlock")
	}
	if err := validateBlock(&cc.ShanghaiBlock.Int); err != nil {
		return errorsmod.Wrap(err, "ShanghaiBlock")
	}
	if err := validateBlock(&cc.CancunBlock.Int); err != nil {
		return errorsmod.Wrap(err, "CancunBlock")
	}
	// NOTE: chain ID is not needed to check config order
	if err := cc.EthereumConfig(nil).CheckConfigForkOrder(); err != nil {
		return errorsmod.Wrap(err, "invalid config fork order")
	}
	return nil
}

func validateHash(hex string) error {
	if hex != "" && strings.TrimSpace(hex) == "" {
		return errorsmod.Wrap(ErrInvalidChainConfig, "hash cannot be blank")
	}

	return nil
}

func validateBlock(block *sdkmath.Int) error {
	// nil value means that the fork has not yet been applied
	if block == nil {
		return nil
	}

	if block.IsNegative() {
		return errorsmod.Wrapf(
			ErrInvalidChainConfig, "block value cannot be negative: %s", block,
		)
	}

	return nil
}
