package router

import (
	"yushu/opt/utility/singleton"
)

type Router struct {
	groupName string
	Route
}

var routerLazySingleton *singleton.Lazy

func New(opt ...interface{}) *Router {
	ins := routerLazySingleton.Instance(new(Router))

	return (*ins).(*Router)
}

func (r *Router) Group(name string) (res *Router) {
	res.groupName = name
	return
}
