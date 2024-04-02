package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/controllers"
)

func setupUserRoutes(baseRouter *gin.RouterGroup) {
	users := baseRouter.Group("/users")

	users.POST("", controllers.CreateUser)

	users.GET("/:user_id/posts/:post_id/comments", controllers.GetAllCommentsByPostId)
	// users.POST("/:user_id/posts/:post_id", controllers.CreateComment)
	users.GET("", controllers.GetUser)
}
