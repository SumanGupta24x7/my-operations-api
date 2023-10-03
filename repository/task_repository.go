package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"my-operations/model/dbModel"
)

type ToDoTaskRepository interface {
	GetAllTasks(ctx *gin.Context) (dbModel.Tasks, error)
	CreateTask(ctx *gin.Context, task dbModel.Task) error
	UpdateTask(ctx *gin.Context, task dbModel.Task) error
	DeleteTask(ctx *gin.Context, taskId int) error
	GetAllAssignees(ctx *gin.Context) (dbModel.Assignees, error)
	GetAllTaskStatus(ctx *gin.Context) (dbModel.AllTaskStatus, error)
}

type toDoTaskRepository struct {
	db *sqlx.DB
}

func NewToDoTaskRepository(db *sqlx.DB) ToDoTaskRepository {
	return toDoTaskRepository{db: db}
}

func (t toDoTaskRepository) GetAllTasks(ctx *gin.Context) (dbModel.Tasks, error) {
	var tasks dbModel.Tasks

	query := `SELECT t.ID, t.TASK_NAME, t.DESCRIPTION, t.ASSIGNEE_ID, a.NAME, t.STATUS_ID, ts.STATUS, t.DUE_DATE, 
    t.CREATED_AT, t.UPDATED_AT
	FROM TASK t
	INNER JOIN TASK_STATUS ts ON t.STATUS_ID = ts.ID
	INNER JOIN ASSIGNEE a ON t.ASSIGNEE_ID = a.ID;`

	err := t.db.SelectContext(ctx.Request.Context(), &tasks, query)

	if err != nil {
		fmt.Println("REPOSITORY -> Error occurred while fetching all tasks from database: ", err)
		return dbModel.Tasks{}, err
	}

	fmt.Println("REPOSITORY -> Successfully fetched all tasks from database")
	return tasks, nil
}

func (t toDoTaskRepository) GetAllAssignees(ctx *gin.Context) (dbModel.Assignees, error) {
	var assignees dbModel.Assignees

	query := `SELECT ID, NAME, REGISTERED_AT FROM ASSIGNEE;`

	err := t.db.SelectContext(ctx.Request.Context(), &assignees, query)

	if err != nil {
		fmt.Println("REPOSITORY -> Error occurred while fetching all assignees from database: ", err)
		return dbModel.Assignees{}, err
	}

	fmt.Println("REPOSITORY -> Successfully fetched all assignees from database")
	return assignees, nil
}

func (t toDoTaskRepository) GetAllTaskStatus(ctx *gin.Context) (dbModel.AllTaskStatus, error) {
	var allTaskStatus dbModel.AllTaskStatus

	query := `SELECT ID, STATUS FROM TASK_STATUS;`

	err := t.db.SelectContext(ctx.Request.Context(), &allTaskStatus, query)

	if err != nil {
		fmt.Println("REPOSITORY -> Error occurred while fetching all task status from database: ", err)
		return dbModel.AllTaskStatus{}, err
	}

	fmt.Println("REPOSITORY -> Successfully fetched all task status from database")
	return allTaskStatus, nil
}

func (t toDoTaskRepository) CreateTask(ctx *gin.Context, task dbModel.Task) error {

	tx, err := t.db.Beginx()
	if err != nil {
		fmt.Println("Repository -> Error encountered while beginning SQL transaction: ", err)
		return err
	}

	defer func() {
		if err != nil {
			fmt.Println("Repository -> Rolling back changes due to error: ", err)
			_ = tx.Rollback()
			return
		}
		fmt.Println("Repository -> Successfully created new task, now committing the transaction: ", err)
		_ = tx.Commit()
	}()

	query := `INSERT INTO TASK(TASK_NAME, DESCRIPTION, ASSIGNEE_ID, STATUS_ID, DUE_DATE) VALUES(:TASK_NAME, :DESCRIPTION, :ASSIGNEE_ID, :STATUS_ID, :DUE_DATE);`
	rows, err := tx.NamedExecContext(ctx, query, task)
	if err != nil {
		fmt.Println("Repository -> Failed to create task due to error: ", err)
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		fmt.Println("Repository -> Failed to create task due to error. No rows affected ")
		return err
	}
	return nil
}

func (t toDoTaskRepository) UpdateTask(ctx *gin.Context, task dbModel.Task) error {

	tx, err := t.db.Beginx()
	if err != nil {
		fmt.Println("Repository -> Error encountered while beginning SQL transaction: ", err)
		return err
	}

	defer func() {
		if err != nil {
			fmt.Println("Repository -> Rolling back changes due to error: ", err)
			_ = tx.Rollback()
			return
		}
		fmt.Println("Repository -> Successfully updated task, now committing the transaction: ", err)
		_ = tx.Commit()
	}()

	updateQuery := `UPDATE TASK SET TASK_NAME=:TASK_NAME, DESCRIPTION=:DESCRIPTION,  ASSIGNEE_ID=:ASSIGNEE_ID, STATUS_ID=:STATUS_ID, DUE_DATE=:DUE_DATE, UPDATED_AT=current_timestamp WHERE ID=:ID`
	rows, err := tx.NamedExecContext(ctx, updateQuery, task)
	if err != nil {
		fmt.Println("Repository -> Failed to update task due to error: ", err)
		return err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		fmt.Println("Repository -> Failed to update task due to error. No rows affected ")
		return err
	}
	return nil
}

func (t toDoTaskRepository) DeleteTask(ctx *gin.Context, taskId int) error {

	tx, err := t.db.Beginx()
	if err != nil {
		fmt.Println("Repository -> Error encountered while beginning SQL transaction: ", err)
		return err
	}

	defer func() {
		if err != nil {
			fmt.Println("Repository -> Rolling back changes due to error: ", err)
			_ = tx.Rollback()
			return
		}
		fmt.Println("Repository -> Successfully deleted task, now committing the transaction: ", err)
		_ = tx.Commit()
	}()

	updateQuery := `DELETE FROM TASK WHERE ID=?`
	rows, err := tx.ExecContext(ctx, updateQuery, taskId)
	if err != nil {
		fmt.Println("Repository -> Failed to delete task due to error: ", err)
		return err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		fmt.Println("Repository -> Failed to delete task due to error. No rows affected ")
		return err
	}
	return nil
}
