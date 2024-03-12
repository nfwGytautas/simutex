package core_components

import (
	"fmt"
	"math"

	simutex "github.com/nfwGytautas/Simutex/Simutex"
)

// Create a new AND gate
func NewAndGate(threshold simutex.ElectricalValue) simutex.BaseComponent {
	return simutex.NewComponent(simutex.BaseComponentArgs{
		Name:       fmt.Sprintf("AndGate_%v", getIndex()),
		NumInputs:  2,
		NumOutputs: 1,
		Tick:       andGateTick(threshold),
	})
}

func andGateTick(threshold simutex.ElectricalValue) simutex.TickFn {
	return func(Inputs []simutex.InputChannel, Outputs []simutex.OutputChannel) {
		aVal := Inputs[0].Read()
		bVal := Inputs[1].Read()

		if aVal < threshold {
			Outputs[0].Write(0)
			return
		}

		if bVal < threshold {
			Outputs[0].Write(0)
			return
		}

		outVoltage := math.Max(float64(aVal), float64(bVal))

		Outputs[0].Write(simutex.ElectricalValue(outVoltage))
	}
}
