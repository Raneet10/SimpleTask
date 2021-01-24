package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers simpletask-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/simpletask/task", createTaskHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/simpletask/task", listTaskHandler(cliCtx, "simpletask")).Methods("GET")
	r.HandleFunc("/simpletask/task/{key}", getTaskHandler(cliCtx, "simpletask")).Methods("GET")
	r.HandleFunc("/simpletask/task", setTaskHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/simpletask/task", finishTaskHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/simpletask/task", deleteTaskHandler(cliCtx)).Methods("DELETE")

}
