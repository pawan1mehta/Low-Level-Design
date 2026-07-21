package main

type Remote struct {
	device Device
}

func NewRemote(device Device) *Remote {
	return &Remote{
		device: device,
	}
}

func (remote *Remote) TurnOn() {
	remote.device.TurnOn()
}

func (remote *Remote) TurnOff() {
	remote.device.TurnOff()
}

type BasicRemote struct {
	*Remote
}

func NewBasicRemote(device Device) *BasicRemote {
	return &BasicRemote{
		Remote: NewRemote(device),
	}
}

type AdvancedRemote struct {
	*Remote
}

func NewAdvancedRemote(device Device) *AdvancedRemote {
	return &AdvancedRemote{
		Remote: NewRemote(device),
	}
}

func (ad *AdvancedRemote) SetVolume(volume int) {
	ad.device.SetVolume(volume)
}
