package router

type HandlerFunc func(ctx interface{})

type RouteInterface interface {
	Run(addr ...string) error
	Group(name string) RouteInterface
	Use(...HandlerFunc)
}
