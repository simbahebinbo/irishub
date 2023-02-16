package keeper

import (
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	"github.com/evmos/ethermint/x/evm/types"
	evm "github.com/evmos/ethermint/x/evm/vm"
)

type Keeper struct {
	evmkeeper.Keeper
	// chain ID number obtained from the context's chain id
	eip155ChainID *big.Int
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey, transientKey storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
	ak types.AccountKeeper,
	bankKeeper types.BankKeeper,
	sk types.StakingKeeper,
	fmk types.FeeMarketKeeper,
	customPrecompiles evm.PrecompiledContracts,
	evmConstructor evm.Constructor,
	tracer string,
	eip155ChainID *big.Int,
) *Keeper {
	return &Keeper{
		Keeper:        *evmkeeper.NewKeeper(cdc, storeKey, transientKey, paramSpace, ak, bankKeeper, sk, fmk, customPrecompiles, evmConstructor, tracer),
		eip155ChainID: eip155ChainID,
	}
}

// WithChainID sets the chain id to the local variable in the keeper
func (k *Keeper) WithChainID(ctx sdk.Context) {
	//nothing to do here
}

func (k Keeper) ChainID() *big.Int {
	return k.eip155ChainID
}
