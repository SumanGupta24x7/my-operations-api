package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"my-operations/controller"
	"my-operations/repository"
	"my-operations/service"
)

func Init(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	registerRoutes(router, db)

	return router
}

func registerRoutes(r *gin.Engine, db *sqlx.DB) {
	taskRepository := repository.NewToDoTaskRepository(db)
	taskService := service.NewToDoTaskService(taskRepository)
	taskController := controller.NewToDoTaskController(taskService)

	routeGroup := r.Group("/api/to-do-task/")
	{
		routeGroup.GET("", taskController.GetAllTasks)
		routeGroup.POST("/create-task", taskController.CreateTask)
		routeGroup.POST("/update-task", taskController.UpdateTask)
		routeGroup.DELETE("/delete-task/:taskId", taskController.DeleteTask)
		routeGroup.GET("/get-assignees", taskController.GetAllAssignees)
		routeGroup.GET("/get-all-task-status", taskController.GetAllTaskStatus)
	}
}
