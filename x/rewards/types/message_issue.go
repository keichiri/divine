package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgIssue{}

func NewMsgIssue(creator string, accounts []string, amounts []uint64) *MsgIssue {
	return &MsgIssue{
		Creator:  creator,
		Accounts: accounts,
		Amounts:  amounts,
	}
}

func (msg *MsgIssue) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
