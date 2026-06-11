package strategy

type Navigator struct {
	routeStrategy RouteStrategy
}

func (n *Navigator) SetRouteStrategy(routeStrategy RouteStrategy) {
	n.routeStrategy = routeStrategy
}

func (n *Navigator) BuildRoute(source, destination string) {
	n.routeStrategy.BuildRoute(source, destination)
}
