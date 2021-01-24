package simpletask

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/raneet10/simpletask/x/simpletask/keeper"
	"github.com/raneet10/simpletask/x/simpletask/types"
)

func handleMsgFinishTask(ctx sdk.Context, k keeper.Keeper, msg types.MsgFinishTask) (*sdk.Result, error) {

	taskId := msg.ID
	task, err := k.GetTask(ctx, taskId)
	if err != nil {
		return nil, err
	}
	err = k.CoinKeeper.SendCoins(ctx, task.Creator, msg.Claim, task.Bond)

	if err != nil {
		return nil, err
	}

	k.DeleteTask(ctx, taskId)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
