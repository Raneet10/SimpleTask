package simpletask

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/raneet10/simpletask/x/simpletask/types"
	"github.com/raneet10/simpletask/x/simpletask/keeper"
)

func handleMsgSetTask(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetTask) (*sdk.Result, error) {
	var task = types.Task{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Name: msg.Name,
    	Bond: msg.Bond,
	}
	if !msg.Creator.Equals(k.GetTaskOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetTask(ctx, task)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
