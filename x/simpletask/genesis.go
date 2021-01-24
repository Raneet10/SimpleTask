package simpletask

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/raneet10/simpletask/x/simpletask/keeper"
	"github.com/raneet10/simpletask/x/simpletask/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	for _, task := range data.Tasks {
		k.SetTask(ctx, task)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	var tasks []types.Task
	iterator := k.GetTaskIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		task, _ := k.GetTask(ctx, key)
		tasks = append(tasks, task)
	}
	return types.GenesisState{Tasks: tasks}
}
