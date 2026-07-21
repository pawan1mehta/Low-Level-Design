package main

import "strategy/strategy"

func main() {

	navigator := &strategy.Navigator{}

	publicTransportStrategy := &strategy.PublicTransportRouteStrategy{}
	// roadRouteStrategy := &strategy.RoadRouteStrategy{}
	// walkRouteStrategy := &strategy.WalkRouteStrategy{}

	navigator.SetRouteStrategy(publicTransportStrategy)

	navigator.BuildRoute("Point A", "Point B")
}
