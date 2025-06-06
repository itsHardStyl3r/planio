package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "hello world")
	})

	r.GET("/task", HandlerGetTasks)
	r.GET("/tasktypes", HandlerGetAllTaskTypes)
	r.POST("/task", HandlerCreateTask)
	r.PUT("/task", HandlerSummarizeTask)
	r.POST("/ai", HandlerPostAI)
	r.PUT("/ai", HandleSaveTasks)
	r.POST("/user/register", HandlerUserRegister)
	r.POST("/user/login", HandlerUserLogin)
	r.GET("/user/checkAuth", HandlerCheckIfSession)
	r.POST("/form", HandlerGetFormData)

	return r
}
