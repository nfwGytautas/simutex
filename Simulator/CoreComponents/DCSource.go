package core_components

import (
	"fmt"

	simutex "github.com/nfwGytautas/Simutex/Simutex"
)

// Create a new AND gate
func NewDCSource(output simutex.ElectricalValue) simutex.BaseComponent {
	return simutex.NewComponent(simutex.BaseComponentArgs{
		Name:       fmt.Sprintf("DCSource_%v", getIndex()),
		NumInputs:  0,
		NumOutputs: 1,
		Tick:       dcSourceTick(output),
	})
}

func dcSourceTick(output simutex.ElectricalValue) simutex.TickFn {
	return func(Inputs []simutex.InputChannel, Outputs []simutex.OutputChannel) {
		Outputs[0].Write(output)
	}
}
