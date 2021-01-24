package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteTask{}

type MsgDeleteTask struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteTask(id string, creator sdk.AccAddress) MsgDeleteTask {
  return MsgDeleteTask{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteTask) Route() string {
  return RouterKey
}

func (msg MsgDeleteTask) Type() string {
  return "DeleteTask"
}

func (msg MsgDeleteTask) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteTask) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteTask) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}