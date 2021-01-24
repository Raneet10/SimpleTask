package simpletask

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/raneet10/simpletask/x/simpletask/types"
	"github.com/raneet10/simpletask/x/simpletask/keeper"
)

func handleMsgCreateTask(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateTask) (*sdk.Result, error) {
	var task = types.Task{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Name: msg.Name,
    	Bond: msg.Bond,
	}
	k.CreateTask(ctx, task)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
