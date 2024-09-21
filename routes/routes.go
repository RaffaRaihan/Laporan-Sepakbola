package routes

import (
    "pr/controllers"
    "pr/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()

    r.POST("/register", controllers.Register) 
    r.POST("/login", controllers.Login)
    
    protected := r.Group("/laporan")
    protected.Use(middlewares.JWTAuthMiddleware())
    {
        protected.GET("/", controllers.Index)
        protected.GET("/:id", controllers.Show)
        protected.POST("/", controllers.Create)
        protected.PUT("/:id", controllers.Update)
        protected.DELETE("/:id", controllers.Delete)
        protected.POST("/:id/upload", controllers.UploadFile)
    }

	r.Run(":3000")

    return r
}
