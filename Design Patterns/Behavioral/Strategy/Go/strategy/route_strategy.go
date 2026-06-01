package strategy

type RouteStrategy interface {
	BuildRoute(source, destination string)
}
