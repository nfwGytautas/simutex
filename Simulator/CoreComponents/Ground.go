package core_components

import (
	"fmt"

	simutex "github.com/nfwGytautas/Simutex/Simutex"
)

// Create a new ground component
func NewGround() simutex.BaseComponent {
	return simutex.NewComponent(simutex.BaseComponentArgs{
		Name:       fmt.Sprintf("Ground_%v", getIndex()),
		NumInputs:  0,
		NumOutputs: 1,
		Tick:       groundTick(),
	})
}

func groundTick() simutex.TickFn {
	return func(Inputs []simutex.InputChannel, Outputs []simutex.OutputChannel) {
		Outputs[0].Write(0)
	}
}
