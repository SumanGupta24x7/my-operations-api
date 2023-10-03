package dbModel

import (
	null "gopkg.in/guregu/null.v3"
	"time"
)

type Task struct {
	TaskId       int         `db:"ID"`
	TaskName     string      `db:"TASK_NAME"`
	Description  null.String `db:"DESCRIPTION"`
	AssigneeId   int         `db:"ASSIGNEE_ID"`
	AssigneeName string      `db:"NAME"`
	StatusId     int         `db:"STATUS_ID"`
	Status       string      `db:"STATUS"`
	DueDate      null.String `db:"DUE_DATE"`
	CreatedAt    time.Time   `db:"CREATED_AT"`
	UpdatedAt    null.Time   `db:"UPDATED_AT"`
}

type TaskStatus struct {
	StatusId int    `db:"ID"`
	Status   string `db:"STATUS"`
}

type Assignee struct {
	AssigneeId   int       `db:"ID"`
	AssigneeName string    `db:"NAME"`
	RegisteredAt time.Time `db:"REGISTERED_AT"`
}

type Tasks []Task

type Assignees []Assignee

type AllTaskStatus []TaskStatus
