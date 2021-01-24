package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFinishTask{}

type MsgFinishTask struct {
	ID    string
	Claim sdk.AccAddress `json:"creator" yaml:"creator"`
	Name  string         `json:"name" yaml:"name"`
}

func NewMsgFinishTask(id string, claim sdk.AccAddress, name string) MsgFinishTask {
	return MsgFinishTask{
		ID:    id,
		Claim: claim,
		Name:  name,
	}
}

func (msg MsgFinishTask) Route() string {
	return RouterKey
}

func (msg MsgFinishTask) Type() string {
	return "FinishTask"
}

func (msg MsgFinishTask) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Claim)}
}

func (msg MsgFinishTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgFinishTask) ValidateBasic() error {
	if msg.Claim.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "claim address can't be empty")
	}

	return nil
}
