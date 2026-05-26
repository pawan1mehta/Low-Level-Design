package main

import "fmt"

type Computer struct {
	CPU                   string
	RAM                   string
	isGraphicsCardEnabled bool
	isBluetoothEnabled    bool
}

func (c Computer) String() string {
	return fmt.Sprintf(
		"Computer [CPU: %s, RAM: %s, Graphics Card Enabled: %t, Bluetooth Enabled: %t]",
		c.CPU,
		c.RAM,
		c.isGraphicsCardEnabled,
		c.isBluetoothEnabled,
	)
}

type ComputerBuilder struct {
	CPU                   string
	RAM                   string
	isGraphicsCardEnabled bool
	isBluetoothEnabled    bool
}

func (cb *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	cb.CPU = cpu
	return cb
}

func (cb *ComputerBuilder) SetRAM(ram string) *ComputerBuilder {
	cb.RAM = ram
	return cb
}

func (cb *ComputerBuilder) SetGraphicsCardEnabled(isGraphicsCardEnabled bool) *ComputerBuilder {
	cb.isGraphicsCardEnabled = isGraphicsCardEnabled
	return cb
}

func (cb *ComputerBuilder) SetBluetoothEnabled(isBluetoothEnabled bool) *ComputerBuilder {
	cb.isBluetoothEnabled = isBluetoothEnabled
	return cb
}

func (cb *ComputerBuilder) Build() Computer {
	return Computer{
		CPU:                   cb.CPU,
		RAM:                   cb.RAM,
		isGraphicsCardEnabled: cb.isGraphicsCardEnabled,
		isBluetoothEnabled:    cb.isBluetoothEnabled,
	}
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{}
}

func main() {
	computer := NewComputerBuilder().
		SetCPU("Intel i7").
		SetRAM("16GB").
		SetGraphicsCardEnabled(true).
		SetBluetoothEnabled(true).
		Build()

	println(computer.String())
}
