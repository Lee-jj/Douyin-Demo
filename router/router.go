package router

import (
	"DOUYIN-DEMO/controller"
	"DOUYIN-DEMO/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", middleware.JWTMiddleware(), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.UserRegister)
	apiRouter.POST("/user/login/", controller.UserLogin)
	apiRouter.POST("/publish/action/", middleware.JWTMiddleware(), controller.Publish)
	apiRouter.GET("/publish/list/", middleware.JWTMiddleware(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.JWTMiddleware(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.JWTMiddleware(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", middleware.JWTMiddleware(), controller.CommentAction)
	apiRouter.GET("/comment/list/", middleware.JWTMiddleware(), controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", middleware.JWTMiddleware(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", middleware.JWTMiddleware(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.JWTMiddleware(), controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", middleware.JWTMiddleware(), controller.FriendList)
	apiRouter.GET("/message/chat/", middleware.JWTMiddleware(), controller.MessageChat)
	apiRouter.POST("/message/action/", middleware.JWTMiddleware(), controller.MessageAction)
}
