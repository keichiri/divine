package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"divine/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TransferToken(goCtx context.Context, msg *types.MsgTransferToken) (*types.MsgTransferTokenResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checking if such a token exists
	token, isFound := k.GetToken(ctx, msg.Index)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no such NFT registered")
	}

	buyerAddress, err := k.accountKeeper.AddressCodec().StringToBytes(msg.Buyer)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid buyer address: %s", err)
	}

	// Checking if the buyer exists
	buyerAccount := k.accountKeeper.GetAccount(ctx, buyerAddress)
	if buyerAccount == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("Buyer not registered on the blockchain")
	}

	// Checking if buyer has sufficient balance
	price := sdk.NewCoin("div", math.NewIntFromUint64(msg.Amount))
	buyerBalance := k.bankKeeper.GetBalance(ctx, buyerAddress, "div")

	if buyerBalance.IsLT(price) {
		return nil, sdkerrors.ErrInsufficientFee.Wrap("Buyer does not have enough divine tokens")
	}

	// Transfering tokens
	seller, _ := k.accountKeeper.AddressCodec().StringToBytes(token.Owner)

	if err := k.bankKeeper.SendCoins(ctx, buyerAddress, seller, sdk.NewCoins(price)); err != nil {
		return nil, sdkerrors.ErrConflict.Wrap("Failed to perform the transfer")
	}

	// Transferring NFT
	// Storing the token
	var updatedToken = types.Token{
		Creator:        msg.Creator,
		Index:          msg.Index,
		Owner:          msg.Buyer,
		DataDescriptor: token.DataDescriptor,
	}

	k.SetToken(
		ctx,
		updatedToken,
	)

	return &types.MsgTransferTokenResponse{}, nil
}
