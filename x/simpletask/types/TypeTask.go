package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Task struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Name    string         `json:"name" yaml:"name"`
	Bond    sdk.Coins      `json:"bond" yaml:"bond"`
}

func NewTask(name string, creator sdk.AccAddress, bond sdk.Coins) Task {
	return Task{
		Creator: creator,
		Name:    name,
		Bond:    bond,
	}
}
