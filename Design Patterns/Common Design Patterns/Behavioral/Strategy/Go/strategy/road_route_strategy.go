package strategy

import "fmt"

type RoadRouteStrategy struct {
}

func (wrs *RoadRouteStrategy) BuildRoute(source, destination string) {
	fmt.Print("Building route for road from ", source, " to ", destination)
}
