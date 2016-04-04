package pipe

import (
	"github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/types"

	"github.com/eris-ltd/eris-db/tmsp"
)

// The consensus struct.
type consensus struct {
	erisdbApp *tmsp.ErisDBApp
}

func newConsensus(erisdbApp *tmsp.ErisDBApp) *consensus {
	return &consensus{erisdbApp}
}

// Get the current consensus state.
func (this *consensus) State() (*ConsensusState, error) {
// <<<<<<< 6f845953eaf7fdc682d444a7018c6c0399984ba4
// 	roundState := this.consensusState.GetRoundState()
// 	peerRoundStates := []string{}
// 	for _, peer := range this.p2pSwitch.Peers().List() {
// 		// TODO: clean this up?
// 		peerState := peer.Data.Get(types.PeerStateKey).(*cm.PeerState)
// 		peerRoundState := peerState.GetRoundState()
// 		peerRoundStateStr := peer.Key + ":" + string(wire.JSONBytes(peerRoundState))
// 		peerRoundStates = append(peerRoundStates, peerRoundStateStr)
// 	}
// 	return FromRoundState(roundState), nil
// =======
	// TODO-RPC!
	return &ConsensusState{}, nil
}

// Get all validators.
func (this *consensus) Validators() (*ValidatorList, error) {
	var blockHeight int
	bondedValidators := make([]*types.Validator, 0)
	unbondingValidators := make([]*types.Validator, 0)

	s := this.erisdbApp.GetState()
	blockHeight = s.LastBlockHeight
	s.BondedValidators.Iterate(func(index int, val *types.Validator) bool {
		bondedValidators = append(bondedValidators, val)
		return false
	})
	s.UnbondingValidators.Iterate(func(index int, val *types.Validator) bool {
		unbondingValidators = append(unbondingValidators, val)
		return false
	})

	return &ValidatorList{blockHeight, bondedValidators, unbondingValidators}, nil
}
