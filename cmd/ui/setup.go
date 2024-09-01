package main

import (
	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"

	"github.com/blinlol/bookworm/cmd/ui/vugu/root"
)


func VuguSetup(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {
	router := vgrouter.New(eventEnv)
	buildEnv.SetWireFunc(func(b vugu.Builder){
		if c, ok := b.(vgrouter.NavigatorSetter); ok {
			c.NavigatorSet(router)
		}
	})

	component := &root.Root{}
	buildEnv.WireComponent(component)

	router.MustAddRouteExact(
		"/",
		vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch){
			component.Body = &root.MainPage{}
		}),
	)
	router.MustAddRouteExact(
		"/page1",
		vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch){
			component.Body = &root.Page1{}
		}),
	)
	router.SetNotFound(
		vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch){
			component.Body = &root.PageNotFound{}
		}),
	)

	err := router.ListenForPopState()
	if err != nil {
		panic(err)
	}

	err = router.Pull()
	if err != nil {
		panic(err)
	}

	return component
}