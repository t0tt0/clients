package nsb

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	dbm "github.com/tendermint/tm-db"
)

func (st *NSBState) String() string {
	return fmt.Sprintf("StateRoot: %v\nHeight: %d\n", hex.EncodeToString(st.StateRoot), st.Height)
}

func (st *NSBState) Reset() {
	*st = NSBState{db: st.db}
}

func loadState(db dbm.DB) *NSBState {
	stateBytes := db.Get(stateKey)
	var state NSBState
	if len(stateBytes) != 0 {
		err := json.Unmarshal(stateBytes, &state)
		if err != nil {
			panic(err)
		}
	}
	state.db = db
	return &state
}

func saveState(state *NSBState) {
	stateBytes, err := json.Marshal(state)
	if err != nil {
		print("savestateerror!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		panic(err)
	}
	state.db.Set(stateKey, stateBytes)
}

func (st *NSBState) Close() error {
	if st.db == nil {
		return errors.New("the state db is not opened now")
	}
	st.db.Close()
	st.db = nil
	return nil
}
