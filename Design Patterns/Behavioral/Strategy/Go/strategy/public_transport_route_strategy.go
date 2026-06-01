package strategy

import "fmt"

type PublicTransportRouteStrategy struct {
}

func (wrs *PublicTransportRouteStrategy) BuildRoute(source, destination string) {
	fmt.Print("Building route for public transport from ", source, " to ", destination)
}
