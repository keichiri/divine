package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferToken{}

func NewMsgTransferToken(creator string, index string, buyer string, amount uint64) *MsgTransferToken {
	return &MsgTransferToken{
		Creator: creator,
		Index:   index,
		Buyer:   buyer,
		Amount:  amount,
	}
}

func (msg *MsgTransferToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	return nil
}
