package router

import (
	"yushu/box/utility/singleton"
)

type Router struct {
	groupName string
	RouteInterface
}

var routerLazySingleton *singleton.Lazy

type FrameworkInterface interface {
	Use(...HandlerFunc)
	GET(string, ...HandlerFunc)
	Run(addr ...string) error
}

func Register(opt ...interface{}) *Router {
	ins := routerLazySingleton.Instance(new(Router))
	return (*ins).(*Router)
}

func (r *Router) Group(name string) (res *Router) {
	res.groupName = name
	return
}
