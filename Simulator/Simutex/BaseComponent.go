package simutex

import "fmt"

type TickFn func(Inputs []InputChannel, Outputs []OutputChannel)

type BaseComponent struct {
	Name    string
	Inputs  []InputChannel
	Outputs []OutputChannel
	tickFn  TickFn
}

type BaseComponentArgs struct {
	Name       string
	NumInputs  uint8
	NumOutputs uint8
	Tick       TickFn
}

func NewComponent(args BaseComponentArgs) BaseComponent {
	return BaseComponent{
		Name:    args.Name,
		Inputs:  make([]InputChannel, args.NumInputs),
		Outputs: make([]OutputChannel, args.NumOutputs),
		tickFn:  args.Tick,
	}
}

func (c *BaseComponent) Connect(bc *BaseComponent, inputIdx int, outputIdx int) {
	c.Inputs[inputIdx].Connect(&bc.Outputs[outputIdx])
}

func (c *BaseComponent) Tick() {
	c.tickFn(c.Inputs, c.Outputs)

	// Tick output channels
	for i := range c.Outputs {
		c.Outputs[i].Tick()
	}

	// Tick input channels
	for i := range c.Inputs {
		c.Inputs[i].Tick()
	}
}

func (c *BaseComponent) GetStates() []StateEntry {
	result := make([]StateEntry, 0)

	for idx, input := range c.Inputs {
		result = append(result, StateEntry{
			ComponentName: fmt.Sprintf("%s_Input_%v", c.Name, idx),
			Value:         input.Read(),
		})
	}

	for idx, output := range c.Outputs {
		result = append(result, StateEntry{
			ComponentName: fmt.Sprintf("%s_Output_%v", c.Name, idx),
			Value:         output.Read(),
		})
	}

	return result
}
