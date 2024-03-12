package simutex

type Model struct {
	Components []BaseComponent
}

func NewModel() Model {
	return Model{}
}

func (m *Model) AddComponent(c BaseComponent) {
	m.Components = append(m.Components, c)
}

func (m *Model) Tick() {
	for _, c := range m.Components {
		c.Tick()
	}
}

func (m *Model) GetState() FrozenState {
	state := FrozenState{}

	for _, component := range m.Components {
		state.Entries = append(state.Entries, component.GetStates()...)
	}

	return state
}
