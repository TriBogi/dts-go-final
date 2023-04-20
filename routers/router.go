package routers

import (
	"DTS-GO-FINAL/controllers"
	"DTS-GO-FINAL/middlewares"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.GET("/:id", controllers.GetOnePhoto)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.PUT("/:id", middlewares.Authorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:id", middlewares.Authorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", controllers.GetAllComment)
		commentRouter.GET("/:id", controllers.GetOneComment)
		commentRouter.POST("/:photoId", controllers.CreateComment)
		commentRouter.PUT("/:id", middlewares.Authorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:id", middlewares.Authorization(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.GET("/", controllers.GetAllSocialMedia)
		socialMediaRouter.GET("/:id", controllers.GetOneSocialMedia)
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.PUT("/:id", middlewares.Authorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", middlewares.Authorization(), controllers.DeleteSocialMedia)
	}

	return r
}
