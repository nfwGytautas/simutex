package simutex

type StateEntry struct {
	ComponentName string
	Value         ElectricalValue
}

// A structure for holding the model state at a given tick
type FrozenState struct {
	Tick    uint64
	Entries []StateEntry
}
