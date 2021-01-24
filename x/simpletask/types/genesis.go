package types

import "fmt"

// GenesisState - all simpletask state that must be provided at genesis
type GenesisState struct {
	Tasks []Task `json:"tasks"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState() GenesisState {
	return GenesisState{
		Tasks: nil,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Tasks: []Task{},
	}
}

// ValidateGenesis validates the simpletask genesis parameters
func ValidateGenesis(data GenesisState) error {
	for _, task := range data.Tasks {
		if task.Creator == nil {
			return fmt.Errorf("Invalid Task : Creator missing")
		}
		if task.ID == "" {
			return fmt.Errorf("Invalid Task : ID missing")
		}
		if task.Name == "" {
			return fmt.Errorf("Invalid Task : Name missing")
		}
		if task.Bond == nil {
			return fmt.Errorf("Invalid Task : Bond missing")
		}
	}
	return nil
}
