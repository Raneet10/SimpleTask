package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateTask{}

type MsgCreateTask struct {
	ID      string
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Name    string         `json:"name" yaml:"name"`
	Bond    sdk.Coins      `json:"bond" yaml:"bond"`
}

func NewMsgCreateTask(creator sdk.AccAddress, name string, bond sdk.Coins) MsgCreateTask {
	return MsgCreateTask{
		ID:      uuid.New().String(),
		Creator: creator,
		Name:    name,
		Bond:    bond,
	}
}

func (msg MsgCreateTask) Route() string {
	return RouterKey
}

func (msg MsgCreateTask) Type() string {
	return "CreateTask"
}

func (msg MsgCreateTask) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateTask) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}

	if msg.Bond.IsZero() || msg.Bond.Empty() || msg.Bond.IsAnyNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "task bond can't be empty or zero or negative")
	}
	return nil
}
