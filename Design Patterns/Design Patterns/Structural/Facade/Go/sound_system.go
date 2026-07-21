package main

type SoundSystem struct{}

func NewSoundSystem() *SoundSystem {
	return &SoundSystem{}
}

func (ss *SoundSystem) On() {
	println("Sound system is on")
}

func (ss *SoundSystem) Off() {
	println("Sound system is off")
}

func (ss *SoundSystem) SetVolume(volume int) {
	println("Setting volume to: ", volume)
}
