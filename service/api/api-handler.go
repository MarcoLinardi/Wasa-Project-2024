package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Login operations
	rt.router.POST("/login", rt.wrap(rt.doLogin)) // testato + frontend

	// User operations
	rt.router.PUT("/user/name", rt.wrap(rt.bearerAuth(rt.setMyUserName))) // testato + frontend
	rt.router.PUT("/user/photo", rt.wrap(rt.bearerAuth(rt.setMyPhoto)))   // testato + frontend
	rt.router.GET("/users", rt.wrap(rt.bearerAuth(rt.listOfUsers)))       // testato + frontend

	// Chat operations
	rt.router.GET("/chats", rt.wrap(rt.bearerAuth(rt.listOfChats)))           // testato + frontend
	rt.router.POST("/chats", rt.wrap(rt.bearerAuth(rt.createChat)))           // testato + frontend
	rt.router.GET("/chats/:chatId", rt.wrap(rt.bearerAuth(rt.detailsChat)))   // testato
	rt.router.DELETE("/chats/:chatId", rt.wrap(rt.bearerAuth(rt.deleteChat))) // testato + frontend

	// Message Operations
	rt.router.POST("/chats/:chatId/messages", rt.wrap(rt.bearerAuth(rt.sendMessage)))                           // testato + frontend (manca invio foto)
	rt.router.DELETE("/chats/:chatId/messages/:messageId", rt.wrap(rt.bearerAuth(rt.deleteMessage)))            // testato + frontend
	rt.router.GET("/chats/:chatId/messages", rt.wrap(rt.bearerAuth(rt.getAllMessages)))                         // testato + frontend
	rt.router.PUT("/chats/:chatId/messages/status", rt.wrap(rt.bearerAuth(rt.updateMessagesStatus)))            // testato
	rt.router.POST("/chats/:chatId/messages/:messageId/forward", rt.wrap(rt.bearerAuth(rt.forwardMessage)))     // testato
	rt.router.POST("/chats/:chatId/messages/:messageId/reactions", rt.wrap(rt.bearerAuth(rt.reactToMessage)))   // testato + frontend
	rt.router.DELETE("/chats/:chatId/messages/:messageId/reactions", rt.wrap(rt.bearerAuth(rt.deleteReaction))) // testato + frontend

	// Group Operations
	rt.router.POST("/chats/:chatId/members", rt.wrap(rt.bearerAuth(rt.addMemberToGroup)))      // testato + frontend
	rt.router.DELETE("/chats/:chatId/members", rt.wrap(rt.bearerAuth(rt.removeMemberToGroup))) // testato + frontend
	rt.router.PUT("/chats/:chatId/name", rt.wrap(rt.bearerAuth(rt.setGroupName)))              // testato + frontend
	rt.router.PUT("/chats/:chatId/photo", rt.wrap(rt.bearerAuth(rt.setGroupPhoto)))            // testato + frontend

	return rt.router
}
