package main

func main() {
	tv := NewTV()

	basicRemote := NewBasicRemote(tv)
	basicRemote.TurnOn()
	basicRemote.TurnOff()

	advancedRemote := NewAdvancedRemote(tv)
	advancedRemote.TurnOn()
	advancedRemote.TurnOff()
	advancedRemote.SetVolume(25)
}
