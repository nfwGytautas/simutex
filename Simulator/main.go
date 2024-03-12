package main

import (
	core_components "github.com/nfwGytautas/Simutex/CoreComponents"
	debugger "github.com/nfwGytautas/Simutex/Debugger"
	simutex "github.com/nfwGytautas/Simutex/Simutex"
)

func main() {
	model := simutex.NewModel()

	// Create components
	ground := core_components.NewGround()
	dcSource := core_components.NewDCSource(5)
	andGate := core_components.NewAndGate(3)
	orGate := core_components.NewOrGate(3)

	// Link components
	andGate.Connect(&dcSource, 0, 0)

	orGate.Connect(&ground, 0, 0)
	orGate.Connect(&dcSource, 1, 0)

	andGate.Connect(&orGate, 1, 0)

	// Add to the model
	model.AddComponent(ground)
	model.AddComponent(dcSource)
	model.AddComponent(andGate)
	model.AddComponent(orGate)

	// Create a debugger
	debugger := debugger.NewDebugger(model)

	// Start the debugger, we will control out from the outside
	err := debugger.Start(":9000")
	if err != nil {
		panic(err)
	}
}
