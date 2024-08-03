package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AskList:   []Ask{},
		ReplyList: []Reply{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in ask
	askIdMap := make(map[uint64]bool)
	askCount := gs.GetAskCount()
	for _, elem := range gs.AskList {
		if _, ok := askIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for ask")
		}
		if elem.Id >= askCount {
			return fmt.Errorf("ask id should be lower or equal than the last id")
		}
		askIdMap[elem.Id] = true
	}
	// Check for duplicated ID in reply
	replyIdMap := make(map[uint64]bool)
	replyCount := gs.GetReplyCount()
	for _, elem := range gs.ReplyList {
		if _, ok := replyIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for reply")
		}
		if elem.Id >= replyCount {
			return fmt.Errorf("reply id should be lower or equal than the last id")
		}
		replyIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
