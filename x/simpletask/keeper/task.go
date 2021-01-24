package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/raneet10/simpletask/x/simpletask/types"
)

// CreateTask creates a task
func (k Keeper) CreateTask(ctx sdk.Context, task types.Task) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.TaskPrefix + task.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(task)
	store.Set(key, value)
}

// GetTask returns the task information
func (k Keeper) GetTask(ctx sdk.Context, key string) (types.Task, error) {
	store := ctx.KVStore(k.storeKey)
	var task types.Task
	byteKey := []byte(types.TaskPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &task)
	if err != nil {
		return task, err
	}
	return task, nil
}

// SetTask sets a task
func (k Keeper) SetTask(ctx sdk.Context, task types.Task) {
	taskKey := task.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(task)
	key := []byte(types.TaskPrefix + taskKey)
	store.Set(key, bz)
}

// DeleteTask deletes a task
func (k Keeper) DeleteTask(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.TaskPrefix + key))
}

func (k Keeper) GetTaskIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.TaskPrefix))
}

//
// Functions used by querier
//

func listTask(ctx sdk.Context, k Keeper) ([]byte, error) {
	var taskList []types.Task
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.TaskPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var task types.Task
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &task)
		taskList = append(taskList, task)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, taskList)
	return res, nil
}

func getTask(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	task, err := k.GetTask(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, task)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetTaskOwner(ctx sdk.Context, key string) sdk.AccAddress {
	task, err := k.GetTask(ctx, key)
	if err != nil {
		return nil
	}
	return task.Creator
}

// Check if the key exists in the store
func (k Keeper) TaskExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.TaskPrefix + key))
}
