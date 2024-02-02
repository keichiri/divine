package keeper

import (
	"cosmossdk.io/math"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"divine/x/rewards/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:           cdc,
		storeService:  storeService,
		authority:     authority,
		logger:        logger,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) IssueRewards(ctx sdk.Context, accountIssuer sdk.AccAddress, accounts []sdk.AccAddress, amounts []uint64) error {
	var totalAmount uint64

	for _, amount := range amounts {
		totalAmount += amount
	}

	// Minting coins to module account
	coins := sdk.NewCoins(sdk.NewCoin("div", math.NewIntFromUint64(totalAmount)))

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return sdkerrors.ErrLogic.Wrap("Failed to mint coins to module account")
	}

	// Sending coins to all individual accounts
	for accountIndex, accountAddress := range accounts {
		amountToSend := sdk.NewCoins(sdk.NewCoin("div", math.NewIntFromUint64(amounts[accountIndex])))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accountAddress, amountToSend); err != nil {
			return sdkerrors.ErrLogic.Wrap("Failed to send from module account to user account")
		}
	}

	return nil
}
