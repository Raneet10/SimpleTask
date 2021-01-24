package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgCreateTask{}, "simpletask/CreateTask", nil)
	cdc.RegisterConcrete(MsgSetTask{}, "simpletask/SetTask", nil)
	cdc.RegisterConcrete(MsgDeleteTask{}, "simpletask/DeleteTask", nil)
	cdc.RegisterConcrete(MsgFinishTask{}, "simpletask/FinishTask", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
