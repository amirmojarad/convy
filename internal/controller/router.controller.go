package controller

import "github.com/gin-gonic/gin"

func SetupUserRoutes(routes gin.IRoutes, userCtrl *User) {
	routes.POST("/user/auth/signup", userCtrl.Signup)
	routes.POST("/user/auth/login", userCtrl.Login)
}

func SetupUserFollowRoutes(routes gin.IRoutes, userFollowCtrl *UserFollow) {
	routes.POST("/user/follow/:id", userFollowCtrl.Follow)
	routes.DELETE("/user/follow/:id", userFollowCtrl.UnFollow)
}
