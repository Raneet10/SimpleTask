package simpletask

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/raneet10/simpletask/x/simpletask/keeper"
	"github.com/raneet10/simpletask/x/simpletask/types"
)

// Handle a message to delete name
func handleMsgDeleteTask(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteTask) (*sdk.Result, error) {
	if !k.TaskExists(ctx, msg.ID) {

		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetTaskOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteTask(ctx, msg.ID)
	return &sdk.Result{}, nil
}
