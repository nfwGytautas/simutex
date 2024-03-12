package core_components

import (
	"fmt"
	"math"

	simutex "github.com/nfwGytautas/Simutex/Simutex"
)

// Create a new AND gate
func NewOrGate(threshold simutex.ElectricalValue) simutex.BaseComponent {
	return simutex.NewComponent(simutex.BaseComponentArgs{
		Name:       fmt.Sprintf("OrGate_%v", getIndex()),
		NumInputs:  2,
		NumOutputs: 1,
		Tick:       orGateTick(threshold),
	})
}

func orGateTick(threshold simutex.ElectricalValue) simutex.TickFn {
	return func(Inputs []simutex.InputChannel, Outputs []simutex.OutputChannel) {
		aVal := Inputs[0].Read()
		bVal := Inputs[1].Read()

		if aVal < threshold && bVal < threshold {
			Outputs[0].Write(0)
			return
		}

		outVoltage := math.Max(float64(aVal), float64(bVal))
		Outputs[0].Write(simutex.ElectricalValue(outVoltage))
	}
}
