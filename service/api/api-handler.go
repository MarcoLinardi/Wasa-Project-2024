package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Login operations
	rt.router.POST("/login", rt.wrap(rt.doLogin)) // testato

	// User operations
	rt.router.PUT("/user/name", rt.wrap(rt.bearerAuth(rt.setMyUserName))) // testato
	rt.router.PUT("/user/photo", rt.wrap(rt.bearerAuth(rt.setMyPhoto)))   // testato
	rt.router.GET("/users", rt.wrap(rt.bearerAuth(rt.listOfUsers)))       // testato

	// Chat operations
	rt.router.GET("/chats", rt.wrap(rt.bearerAuth(rt.listOfChats)))           // testato
	rt.router.POST("/chats", rt.wrap(rt.bearerAuth(rt.createChat)))           // testato
	rt.router.GET("/chats/:chatId", rt.wrap(rt.bearerAuth(rt.detailsChat)))   // testato
	rt.router.DELETE("/chats/:chatId", rt.wrap(rt.bearerAuth(rt.deleteChat))) // testato

	// Message Operations
	rt.router.POST("/chats/:chatId/messages", rt.wrap(rt.bearerAuth(rt.sendMessage)))                           // testato
	rt.router.DELETE("/chats/:chatId/messages/:messageId", rt.wrap(rt.bearerAuth(rt.deleteMessage)))            // testato
	rt.router.GET("/chats/:chatId/messages/status", rt.wrap(rt.bearerAuth(rt.getAllMessages)))                  // testato
	rt.router.PUT("/chats/:chatId/messages/status/update", rt.wrap(rt.bearerAuth(rt.updateMessagesStatus)))     // testato
	rt.router.POST("/chats/:chatId/messages/:messageId/forward", rt.wrap(rt.bearerAuth(rt.forwardMessage)))     // testato
	rt.router.POST("/chats/:chatId/messages/:messageId/reactions", rt.wrap(rt.bearerAuth(rt.reactToMessage)))   // testato
	rt.router.DELETE("/chats/:chatId/messages/:messageId/reactions", rt.wrap(rt.bearerAuth(rt.deleteReaction))) // testato

	// Group Operations
	rt.router.POST("/chats/:chatId/members", rt.wrap(rt.bearerAuth(rt.addMemberToGroup)))              // testato
	rt.router.DELETE("/chats/:chatId/members/:userId", rt.wrap(rt.bearerAuth(rt.removeMemberToGroup))) // testato
	rt.router.PUT("/chats/:chatId/name", rt.wrap(rt.bearerAuth(rt.setGroupName)))                      // testato
	rt.router.PUT("/chats/:chatId/photo", rt.wrap(rt.bearerAuth(rt.setGroupPhoto)))                    // testato

	return rt.router
}
