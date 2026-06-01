package main

func main() {
	dvdPlayer := NewDVDPlayer()
	projector := NewProjector()
	soundSystem := NewSoundSystem()

	homeTheater := NewHomeTheaterFacade(dvdPlayer, projector, soundSystem)

	homeTheater.WatchMovie("You&Me")
}
