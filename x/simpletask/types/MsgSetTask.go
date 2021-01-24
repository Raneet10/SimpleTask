package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetTask{}

type MsgSetTask struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
	Bond    sdk.Coins      `json:"bond" yaml:"bond"`
}

func NewMsgSetTask(creator sdk.AccAddress, id string, name string, bond sdk.Coins) MsgSetTask {
	return MsgSetTask{
		ID:      id,
		Creator: creator,
		Name:    name,
		Bond:    bond,
	}
}

func (msg MsgSetTask) Route() string {
	return RouterKey
}

func (msg MsgSetTask) Type() string {
	return "SetTask"
}

func (msg MsgSetTask) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSetTask) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}

	if msg.Bond.IsZero() || msg.Bond.Empty() || msg.Bond.IsAnyNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "task bond can't be empty or negative or zero")
	}
	return nil
}
