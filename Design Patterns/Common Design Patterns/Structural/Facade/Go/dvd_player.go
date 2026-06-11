package main

type DVDPlayer struct{}

func NewDVDPlayer() *DVDPlayer {
	return &DVDPlayer{}
}

func (ss *DVDPlayer) On() {
	println("DVDPlayer is on")
}

func (ss *DVDPlayer) Off() {
	println("DVDPlayer is off")
}

func (ss *DVDPlayer) play(movie string) {
	println("Playing movie: ", movie)
}
