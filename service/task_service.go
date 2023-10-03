package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-operations/mapper"
	"my-operations/model/apiModel"
	"my-operations/model/dbModel"
	"my-operations/repository"
)

type ToDoTaskService interface {
	GetAllTasks(ctx *gin.Context) (apiModel.GetTasksResponse, error)
	CreateTask(ctx *gin.Context, task apiModel.Task) error
	UpdateTask(ctx *gin.Context, task apiModel.Task) error
	DeleteTask(ctx *gin.Context, taskId int) error
	GetAllAssignees(ctx *gin.Context) (apiModel.GetAssigneesResponse, error)
	GetAllTaskStatus(ctx *gin.Context) (apiModel.GetAllTaskStatusResponse, error)
}

type toDoTaskService struct {
	toDoTaskRepository repository.ToDoTaskRepository
}

func NewToDoTaskService(toDoTaskRepository repository.ToDoTaskRepository) ToDoTaskService {
	return toDoTaskService{toDoTaskRepository: toDoTaskRepository}
}

func (t toDoTaskService) GetAllTasks(ctx *gin.Context) (apiModel.GetTasksResponse, error) {
	tasks, err := t.toDoTaskRepository.GetAllTasks(ctx)
	if err != nil {
		fmt.Println("Service -> Failed to fetch all tasks from database")
		return apiModel.GetTasksResponse{}, err
	}
	return mapper.ToTasksApiModel(tasks), nil
}

func (t toDoTaskService) GetAllAssignees(ctx *gin.Context) (apiModel.GetAssigneesResponse, error) {
	assignees, err := t.toDoTaskRepository.GetAllAssignees(ctx)
	if err != nil {
		fmt.Println("Service -> Failed to fetch all assignees from database")
		return apiModel.GetAssigneesResponse{}, err
	}
	return mapper.ToAssigneesApiModel(assignees), nil
}

func (t toDoTaskService) GetAllTaskStatus(ctx *gin.Context) (apiModel.GetAllTaskStatusResponse, error) {
	allTaskStatus, err := t.toDoTaskRepository.GetAllTaskStatus(ctx)
	if err != nil {
		fmt.Println("Service -> Failed to fetch all task status from database")
		return apiModel.GetAllTaskStatusResponse{}, err
	}
	return mapper.ToAllTaskStatusApiModel(allTaskStatus), nil
}

func (t toDoTaskService) CreateTask(ctx *gin.Context, task apiModel.Task) error {
	fmt.Println("Service -> Creating task")

	taskDbModel := dbModel.Task{TaskName: task.TaskName, Description: task.Description, AssigneeId: task.AssigneeId, StatusId: task.StatusId, DueDate: task.DueDate}

	err := t.toDoTaskRepository.CreateTask(ctx, taskDbModel)
	if err != nil {
		fmt.Println("Service -> Failed to create task")
		return err
	}
	return nil
}

func (t toDoTaskService) UpdateTask(ctx *gin.Context, task apiModel.Task) error {
	fmt.Println("Service -> Updating task for Task ID : ", task.TaskId)

	taskDbModel := dbModel.Task{TaskId: task.TaskId, TaskName: task.TaskName, Description: task.Description, AssigneeId: task.AssigneeId, StatusId: task.StatusId, DueDate: task.DueDate}

	err := t.toDoTaskRepository.UpdateTask(ctx, taskDbModel)
	if err != nil {
		fmt.Println("Service -> Failed to update task")
		return err
	}
	return nil
}

func (t toDoTaskService) DeleteTask(ctx *gin.Context, taskId int) error {
	fmt.Println("Service -> Deleting task for Task ID : ", taskId)

	err := t.toDoTaskRepository.DeleteTask(ctx, taskId)
	if err != nil {
		fmt.Println("Service -> Failed to delete task")
		return err
	}
	return nil
}
