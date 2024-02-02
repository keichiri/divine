package keeper

import (
	"context"
	"divine/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := k.accountKeeper.AddressCodec().StringToBytes(msg.Creator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	accountStrings := msg.Accounts
	accountAddresses := make([]sdk.AccAddress, 0)

	for _, accountString := range accountStrings {
		accountAddress, err := k.accountKeeper.AddressCodec().StringToBytes(accountString)
		if err != nil {
			return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid account address: %s", err)
		}

		accountAddresses = append(accountAddresses, accountAddress)
	}

	err = k.IssueRewards(ctx, creatorAddr, accountAddresses, msg.Amounts)
	if err != nil {
		return nil, err
	}

	return &types.MsgIssueResponse{}, nil
}
