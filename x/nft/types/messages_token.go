package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateToken{}

func NewMsgCreateToken(
	creator string,
	index string,
	owner string,
	dataDescriptor string,
	fee uint64,
) *MsgCreateToken {
	return &MsgCreateToken{
		Creator:        creator,
		Index:          index,
		Owner:          owner,
		DataDescriptor: dataDescriptor,
		Fee:            fee,
	}
}

func (msg *MsgCreateToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	return nil
}

//var _ sdk.Msg = &MsgUpdateToken{}
//
//func NewMsgUpdateToken(
//	creator string,
//	index string,
//	owner string,
//	dataDescriptor string,
//
//) *MsgUpdateToken {
//	return &MsgUpdateToken{
//		Creator:        creator,
//		Index:          index,
//		Owner:          owner,
//		DataDescriptor: dataDescriptor,
//	}
//}
//
//func (msg *MsgUpdateToken) ValidateBasic() error {
//	_, err := sdk.AccAddressFromBech32(msg.Creator)
//	if err != nil {
//		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
//	}
//	return nil
//}
//
//var _ sdk.Msg = &MsgDeleteToken{}
//
//func NewMsgDeleteToken(
//	creator string,
//	index string,
//
//) *MsgDeleteToken {
//	return &MsgDeleteToken{
//		Creator: creator,
//		Index:   index,
//	}
//}
//
//func (msg *MsgDeleteToken) ValidateBasic() error {
//	_, err := sdk.AccAddressFromBech32(msg.Creator)
//	if err != nil {
//		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
//	}
//	return nil
//}
