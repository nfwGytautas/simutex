package simutex

// Class used to specify an electrical output, this can be a source, a wire, a component, etc.
type OutputChannel struct {
	value ElectricalValue
}

// Class used to specify an electrical input, this can be a component input, ground, etc.
type InputChannel struct {
	connected    *OutputChannel
	currentValue ElectricalValue
}

// Create a new output channel
func NewOutputChannel() OutputChannel {
	return OutputChannel{
		value: 0,
	}
}

// Write electrical value to the channel
func (oc *OutputChannel) Write(value ElectricalValue) {
	oc.value = value
}

// Read an electrical value from the channel
func (oc *OutputChannel) Read() ElectricalValue {
	return oc.value
}

// Tick the channel
func (oc *OutputChannel) Tick() {
}

// Create a new input channel
func NewInputChannel() InputChannel {
	return InputChannel{
		connected: nil,
	}
}

// Connect output channel to an input channel
func (ic *InputChannel) Connect(oc *OutputChannel) {
	ic.connected = oc
}

// Read a connect input value
func (ic *InputChannel) Read() ElectricalValue {
	return ic.currentValue
}

// Tick the channel
func (ic *InputChannel) Tick() {
	if ic.connected == nil {
		ic.currentValue = 0
		return
	}

	ic.currentValue = ic.connected.Read()
}
