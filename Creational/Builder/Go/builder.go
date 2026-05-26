package main

import "fmt"

type Computer struct {
	CPU                 string
	RAM                 string
	graphicsCardEnabled bool
	bluetoothEnabled    bool
}

func (c Computer) String() string {
	return fmt.Sprintf(
		"Computer [CPU: %s, RAM: %s, Graphics Card Enabled: %t, Bluetooth Enabled: %t]",
		c.CPU,
		c.RAM,
		c.graphicsCardEnabled,
		c.bluetoothEnabled,
	)
}

type ComputerOption func(*Computer)

func WithCPU(cpu string) ComputerOption {
	return func(c *Computer) {
		c.CPU = cpu
	}
}

func WithRAM(ram string) ComputerOption {
	return func(c *Computer) {
		c.RAM = ram
	}
}

func WithGraphicsCardEnabled(v bool) ComputerOption {
	return func(c *Computer) { c.graphicsCardEnabled = v }
}

func WithBluetoothEnabled(v bool) ComputerOption {
	return func(c *Computer) { c.bluetoothEnabled = v }
}

func NewComputer(opts ...ComputerOption) (Computer, error) {
	computer := Computer{
		CPU: "default-cpu",
		RAM: "default-ram",
	}

	for _, opt := range opts {
		opt(&computer)
	}

	return computer, nil
}

func main() {
	computer, err := NewComputer(
		WithCPU("Intel i7"),
		WithRAM("16GB"),
		WithGraphicsCardEnabled(true),
		WithBluetoothEnabled(true),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(computer.String())
}
