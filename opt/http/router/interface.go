package router

type HandlerFunc func(ctx interface{})

type Route interface {
	Run(addr ...string) error
	Group(name string) Route
	Use(...HandlerFunc)
}
