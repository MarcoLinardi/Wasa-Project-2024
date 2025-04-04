package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Login operations
	rt.router.POST("/session", rt.wrap(rt.doLogin)) // testato

	// User operations
	rt.router.PATCH("/user/name", rt.wrap(rt.bearerAuth(rt.setMyUserName))) // testato
	rt.router.PATCH("/user/photo", rt.wrap(rt.bearerAuth(rt.setMyPhoto)))
	rt.router.GET("/users", rt.wrap(rt.bearerAuth(rt.listOfUsers))) // testato

	// Chat operations
	rt.router.GET("/chats", rt.wrap(rt.bearerAuth(rt.listOfChats))) // testato
	rt.router.POST("/chats", rt.wrap(rt.bearerAuth(rt.createChat))) // testato
	rt.router.GET("/chats/:ChatId", rt.wrap(rt.bearerAuth(rt.detailsChat)))
	rt.router.DELETE("/chats/:ChatId", rt.wrap(rt.bearerAuth(rt.deleteChat)))

	// Message Operations
	rt.router.POST("/chats/:ChatId/messages", rt.wrap(rt.bearerAuth(rt.sendMessage)))

	return rt.router
}
