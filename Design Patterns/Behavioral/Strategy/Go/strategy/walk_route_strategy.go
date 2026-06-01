package strategy

import "fmt"

type WalkRouteStrategy struct {
}

func (wrs *WalkRouteStrategy) BuildRoute(source, destination string) {
	fmt.Print("Building route for walking from ", source, " to ", destination)
}
