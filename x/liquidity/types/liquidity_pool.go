package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"ollo/x/liquidity/amm"
)

var (
	_ amm.Orderer = (*PoolOrderer)(nil)

	poolCoinDenomRegexp = regexp.MustCompile(`^pool([1-9]\d*)$`)
)

// PoolReserveAddress returns a unique pool reserve account address for each pool.
func PoolReserveAddress(poolId uint64) sdk.AccAddress {
	return DeriveAddress(
		LiquidityAddressType,
		ModuleName,
		strings.Join([]string{PoolReserveAddressPrefix, strconv.FormatUint(poolId, 10)}, ModuleAddressNameSplitter),
	)
}

// PoolCoinDenom returns a unique pool coin denom for a pool.
func PoolCoinDenom(poolId uint64) string {
	return fmt.Sprintf("pool%d", poolId)
}

// ParsePoolCoinDenom parses a pool coin denom and returns the pool id.
func ParsePoolCoinDenom(denom string) (poolId uint64, err error) {
	chunks := poolCoinDenomRegexp.FindStringSubmatch(denom)
	if len(chunks) == 0 {
		return 0, fmt.Errorf("%s is not a pool coin denom", denom)
	}
	poolId, err = strconv.ParseUint(chunks[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse pool id: %w", err)
	}
	return poolId, nil
}

// NewBasicPool returns a new basic pool object.
func NewBasicPool(id, pairId uint64, creator sdk.AccAddress) Pool {
	return Pool{
		TypeId:                1,
		Id:                    id,
		PairId:                pairId,
		Creator:               creator.String(),
		ReserveAccountAddress: PoolReserveAddress(id).String(),
		PoolCoinDenom:         PoolCoinDenom(id),
		PrevDepositRequestId:  0,
		PrevWithdrawRequestId: 0,
		MinPrice:              nil,
		MaxPrice:              nil,
		Inactive:              false,
	}
}

// NewIntelligentPool returns a new Intelligent pool object.
func NewIntelligentPool(id, pairId uint64, creator sdk.AccAddress, minPrice, maxPrice sdk.Dec) Pool {
	return Pool{
		Id:                    id,
		TypeId:                2,
		PairId:                pairId,
		Creator:               creator.String(),
		ReserveAccountAddress: PoolReserveAddress(id).String(),
		PoolCoinDenom:         PoolCoinDenom(id),
		MinPrice:              &minPrice,
		MaxPrice:              &maxPrice,
		PrevDepositRequestId:  0,
		PrevWithdrawRequestId: 0,
		Inactive:              false,
	}
}

func (pool Pool) GetCreator() sdk.AccAddress {
	if pool.Creator == "" {
		return nil
	}
	addr, err := sdk.AccAddressFromBech32(pool.Creator)
	if err != nil {
		panic(err)
	}
	return addr
}

func (pool Pool) GetReserveAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(pool.ReserveAccountAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// Validate validates Pool for genesis.
func (pool Pool) Validate() error {
	if pool.Id == 0 {
		return fmt.Errorf("pool id must not be 0")
	}
	if pool.PairId == 0 {
		return fmt.Errorf("pair id must not be 0")
	}
	if _, err := sdk.AccAddressFromBech32(pool.ReserveAccountAddress); err != nil {
		return fmt.Errorf("invalid reserve address %s: %w", pool.ReserveAccountAddress, err)
	}
	if err := sdk.ValidateDenom(pool.PoolCoinDenom); err != nil {
		return fmt.Errorf("invalid pool coin denom: %w", err)
	}
	return nil
}

// AMMPool constructs amm.Pool interface from Pool.
func (pool Pool) AMMPool(rx, ry, ps sdk.Int) amm.Pool {
	switch pool.TypeId {
	case 1:
		return amm.NewBasicPool(rx, ry, ps)
	case 2:
		return amm.NewIntelligentPool(rx, ry, ps, *pool.MinPrice, *pool.MaxPrice)
	default:
		panic(fmt.Errorf("invalid pool type: %s", pool.TypeId))
	}
}

type PoolOrderer struct {
	amm.Pool
	Id                            uint64
	ReserveAddress                sdk.AccAddress
	BaseCoinDenom, QuoteCoinDenom string
}

func NewPoolOrderer(pool amm.Pool, id uint64, reserveAddr sdk.AccAddress, baseCoinDenom, quoteCoinDenom string) *PoolOrderer {
	return &PoolOrderer{
		Pool:           pool,
		Id:             id,
		ReserveAddress: reserveAddr,
		BaseCoinDenom:  baseCoinDenom,
		QuoteCoinDenom: quoteCoinDenom,
	}
}

func (orderer *PoolOrderer) Order(dir amm.OrderDirection, price sdk.Dec, amt sdk.Int) amm.Order {
	var offerCoinDenom, demandCoinDenom string
	switch dir {
	case amm.Buy:
		offerCoinDenom, demandCoinDenom = orderer.QuoteCoinDenom, orderer.BaseCoinDenom
	case amm.Sell:
		offerCoinDenom, demandCoinDenom = orderer.BaseCoinDenom, orderer.QuoteCoinDenom
	}
	return NewPoolOrder(orderer.Id, orderer.ReserveAddress, dir, price, amt, offerCoinDenom, demandCoinDenom)
}

// MustMarshalPool returns the pool bytes.
// It throws panic if it fails.
func MustMarshalPool(cdc codec.BinaryCodec, pool Pool) []byte {
	return cdc.MustMarshal(&pool)
}

// MustUnmarshalPool return the unmarshalled pool from bytes.
// It throws panic if it fails.
func MustUnmarshalPool(cdc codec.BinaryCodec, value []byte) Pool {
	pool, err := UnmarshalPool(cdc, value)
	if err != nil {
		panic(err)
	}

	return pool
}

// UnmarshalPool returns the pool from bytes.
func UnmarshalPool(cdc codec.BinaryCodec, value []byte) (pool Pool, err error) {
	err = cdc.Unmarshal(value, &pool)
	return pool, err
}
