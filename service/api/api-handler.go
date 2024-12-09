package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Login operations
	rt.router.POST("/session", rt.wrap(rt.doLogin)) // fatto
	// User operations
	rt.router.PATCH("/user/name", rt.wrap(rt.bearerAuth(rt.setMyUserName))) // da fare
	//rt.router.PATCH("/user/photo", rt.wrap(rt.bearerAuth(rt.setMyPhoto))) // da fare
	// Chat operations

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
