package main

type HomeTheaterFacade struct {
	dvdPlayer   *DVDPlayer
	projector   *Projector
	soundSystem *SoundSystem
}

func NewHomeTheaterFacade(dvdPlayer *DVDPlayer, projector *Projector, soundSystem *SoundSystem) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		dvdPlayer:   dvdPlayer,
		projector:   projector,
		soundSystem: soundSystem,
	}
}

func (htf *HomeTheaterFacade) WatchMovie(movie string) {
	println("Get read to watch a movie...")

	htf.dvdPlayer.On()
	htf.projector.On()

	htf.projector.SetInput("DVD")

	htf.soundSystem.On()

	htf.soundSystem.SetVolume(8)

	htf.dvdPlayer.play(movie)
}
