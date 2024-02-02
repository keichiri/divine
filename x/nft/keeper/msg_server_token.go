package keeper

import (
	"context"
	"cosmossdk.io/math"

	errorsmod "cosmossdk.io/errors"
	"divine/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetToken(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	//Determining the fee - if the message has the fee, that one is used
	var feeAmount uint64
	if msg.Fee > 0 {
		feeAmount = msg.Fee
	} else {
		feeAmount = k.defaultFeeAmount
	}

	fee := sdk.NewCoin("div", math.NewIntFromUint64(feeAmount))

	// Checking the balance
	creatorAddr, err := k.accountKeeper.AddressCodec().StringToBytes(msg.Creator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	balance := k.bankKeeper.GetBalance(ctx, creatorAddr, "div")
	if balance.IsLT(fee) {
		return nil, sdkerrors.ErrInsufficientFee.Wrapf("Not enough balance for account: %s", creatorAddr)
	}

	// Burning the tokens
	// First sending tokens from account to module as only modules can burn
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, sdk.NewCoins(fee)); err != nil {
		return nil, sdkerrors.ErrLogic.Wrap("Failed to send from module account to user account")
	}

	// Second - actually burning the tokens
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(fee)); err != nil {
		return nil, err
	}

	// Storing the token
	var token = types.Token{
		Creator:        msg.Creator,
		Index:          msg.Index,
		Owner:          msg.Owner,
		DataDescriptor: msg.DataDescriptor,
	}

	k.SetToken(
		ctx,
		token,
	)
	return &types.MsgCreateTokenResponse{}, nil
}

//func (k msgServer) UpdateToken(goCtx context.Context, msg *types.MsgUpdateToken) (*types.MsgUpdateTokenResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//	// Check if the value exists
//	valFound, isFound := k.GetToken(
//		ctx,
//		msg.Index,
//	)
//	if !isFound {
//		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
//	}
//
//	// Checks if the msg creator is the same as the current owner
//	if msg.Creator != valFound.Creator {
//		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
//	}
//
//	var token = types.Token{
//		Creator:        msg.Creator,
//		Index:          msg.Index,
//		Owner:          msg.Owner,
//		DataDescriptor: msg.DataDescriptor,
//	}
//
//	k.SetToken(ctx, token)
//
//	return &types.MsgUpdateTokenResponse{}, nil
//}
//
//func (k msgServer) DeleteToken(goCtx context.Context, msg *types.MsgDeleteToken) (*types.MsgDeleteTokenResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//	// Check if the value exists
//	valFound, isFound := k.GetToken(
//		ctx,
//		msg.Index,
//	)
//	if !isFound {
//		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
//	}
//
//	// Checks if the msg creator is the same as the current owner
//	if msg.Creator != valFound.Creator {
//		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
//	}
//
//	k.RemoveToken(
//		ctx,
//		msg.Index,
//	)
//
//	return &types.MsgDeleteTokenResponse{}, nil
//}
