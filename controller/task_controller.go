package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-operations/model/apiModel"
	"my-operations/service"
	"net/http"
	"strconv"
)

type ToDoTaskController interface {
	GetAllTasks(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
	GetAllAssignees(ctx *gin.Context)
	GetAllTaskStatus(ctx *gin.Context)
}

type toDoTaskController struct {
	toDoTaskService service.ToDoTaskService
}

func NewToDoTaskController(toDoTaskService service.ToDoTaskService) ToDoTaskController {
	return toDoTaskController{toDoTaskService: toDoTaskService}
}

func (t toDoTaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := t.toDoTaskService.GetAllTasks(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (t toDoTaskController) GetAllAssignees(ctx *gin.Context) {
	assignees, err := t.toDoTaskService.GetAllAssignees(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, assignees)
}

func (t toDoTaskController) GetAllTaskStatus(ctx *gin.Context) {
	allTaskStatus, err := t.toDoTaskService.GetAllTaskStatus(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, allTaskStatus)
}

func (t toDoTaskController) CreateTask(ctx *gin.Context) {

	var taskApiModel apiModel.Task
	bindErr := ctx.BindJSON(&taskApiModel)
	if bindErr != nil {
		fmt.Println("Controller -> Received bad request to create task")
		errorBody := "Bad request api body"
		ctx.JSON(http.StatusBadRequest, errorBody)
		return
	}
	err := t.toDoTaskService.CreateTask(ctx, taskApiModel)

	if err != nil {
		fmt.Println("Controller -> Failed to create task")
		errorBody := "Failed to create task"
		ctx.JSON(http.StatusInternalServerError, errorBody)
		return
	}
	fmt.Println("Controller -> Task created")
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Created!"})
}

func (t toDoTaskController) UpdateTask(ctx *gin.Context) {

	var taskApiModel apiModel.Task
	bindErr := ctx.BindJSON(&taskApiModel)
	if bindErr != nil {
		fmt.Println("Controller -> Received bad request to update task")
		errorBody := "Bad request api body"
		ctx.JSON(http.StatusBadRequest, errorBody)
		return
	}
	err := t.toDoTaskService.UpdateTask(ctx, taskApiModel)

	if err != nil {
		fmt.Println("Controller -> Failed to update task")
		errorBody := "Failed to update task"
		ctx.JSON(http.StatusInternalServerError, errorBody)
		return
	}
	fmt.Println("Controller -> Task updated")
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Updated!"})
}

func (t toDoTaskController) DeleteTask(ctx *gin.Context) {

	taskId, _ := strconv.Atoi(ctx.Param("taskId"))

	err := t.toDoTaskService.DeleteTask(ctx, taskId)

	if err != nil {
		fmt.Println("Controller -> Failed to delete task")
		errorBody := "Failed to delete task"
		ctx.JSON(http.StatusInternalServerError, errorBody)
		return
	}
	fmt.Println("Controller -> Task deleted")
	ctx.JSON(http.StatusOK, gin.H{"message": "Task Deleted!"})
}
