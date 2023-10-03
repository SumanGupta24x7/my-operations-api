package apiModel

import (
	null "gopkg.in/guregu/null.v3"
	"time"
)

type Task struct {
	TaskId       int         `json:"TaskId"`
	TaskName     string      `json:"TaskName" binding:"required"`
	Description  null.String `json:"Description"`
	AssigneeId   int         `json:"AssigneeId"`
	AssigneeName string      `json:"AssigneeName"`
	StatusId     int         `json:"StatusId"`
	Status       string      `json:"Status"`
	DueDate      null.String `json:"DueDate"`
	CreatedAt    time.Time   `json:"CreatedAt"`
	UpdatedAt    null.Time   `json:"UpdatedAt"`
}

type TaskStatus struct {
	StatusId int    `json:"StatusId"`
	Status   string `json:"Status" binding:"required"`
}

type Assignee struct {
	AssigneeId   int       `json:"AssigneeId"`
	AssigneeName string    `json:"AssigneeName" binding:"required"`
	RegisteredAt time.Time `json:"RegisteredAt"`
}

type Tasks []Task

type GetTasksResponse struct {
	Tasks []Task `json:"Tasks"`
}

type Assignees []Assignee

type GetAssigneesResponse struct {
	Assignees []Assignee `json:"Assignees"`
}

type AllTaskStatus []TaskStatus

type GetAllTaskStatusResponse struct {
	AllTaskStatus []TaskStatus `json:"AllTaskStatus"`
}
