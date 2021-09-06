package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "gamm"

	StoreKey = ModuleName

	TStoreKey = "transient_" + ModuleName

	RouterKey = ModuleName

	QuerierRoute = ModuleName
)

var (
	// KeyNextGlobalPoolNumber defines key to store the next Pool ID to be used
	KeyNextGlobalPoolNumber = []byte{0x01}
	// KeyPrefixPools defines prefix to store pools
	KeyPrefixPools = []byte{0x02}
	// KeyTotalLiquidity defines key to store total liquidity
	KeyTotalLiquidity = []byte{0x03}
	// KeyPrefixPoolTwaps defines prefix to store pool twaps
	KeyPrefixPoolTwaps = []byte{0x04}
	// KeyNextTwapHistoryDeleteIndex defines key to store the next twap history to be deleted
	KeyNextTwapHistoryDeleteIndex = []byte{0x05}
	// KeyModifiedPool defines key to store modified pools in the block
	KeyModifiedPool = []byte{0x06}
)

func GetPoolShareDenom(poolId uint64) string {
	return fmt.Sprintf("gamm/pool/%d", poolId)
}

func GetKeyPrefixPools(poolId uint64) []byte {
	return append(KeyPrefixPools, sdk.Uint64ToBigEndian(poolId)...)
}

func GetKeyPoolTwaps(poolId uint64, timestamp int64) []byte {
	if timestamp < 0 {
		panic("Timestamp has negative value")
	}
	keyPrefix := append(KeyPrefixPoolTwaps, sdk.Uint64ToBigEndian(poolId)...)
	return append(keyPrefix, sdk.Uint64ToBigEndian(uint64(timestamp))...)
}

func GetKeyModifiedPool(poolId uint64) []byte {
	return append(KeyModifiedPool, sdk.Uint64ToBigEndian(poolId)...)
}
