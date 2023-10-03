package mapper

import (
	"my-operations/model/apiModel"
	"my-operations/model/dbModel"
)

func ToTasksApiModel(tasks dbModel.Tasks) apiModel.GetTasksResponse {
	var tasksApiModel apiModel.GetTasksResponse
	for _, task := range tasks {
		tasksApiModel.Tasks = append(tasksApiModel.Tasks, ToTaskApiModel(task))
	}
	return tasksApiModel
}

func ToTaskApiModel(task dbModel.Task) apiModel.Task {

	return apiModel.Task{
		TaskId:       task.TaskId,
		TaskName:     task.TaskName,
		Description:  task.Description,
		AssigneeId:   task.AssigneeId,
		AssigneeName: task.AssigneeName,
		StatusId:     task.StatusId,
		Status:       task.Status,
		DueDate:      task.DueDate,
		CreatedAt:    task.CreatedAt,
		UpdatedAt:    task.UpdatedAt,
	}
}

func ToAssigneesApiModel(assignees dbModel.Assignees) apiModel.GetAssigneesResponse {
	var assigneesApiModel apiModel.GetAssigneesResponse
	for _, assignee := range assignees {
		assigneesApiModel.Assignees = append(assigneesApiModel.Assignees, ToAssigneeApiModel(assignee))
	}
	return assigneesApiModel
}

func ToAssigneeApiModel(assignee dbModel.Assignee) apiModel.Assignee {

	return apiModel.Assignee{
		AssigneeId:   assignee.AssigneeId,
		AssigneeName: assignee.AssigneeName,
		RegisteredAt: assignee.RegisteredAt,
	}
}

func ToAllTaskStatusApiModel(allTaskStatus dbModel.AllTaskStatus) apiModel.GetAllTaskStatusResponse {
	var allTaskStatusApiModel apiModel.GetAllTaskStatusResponse
	for _, taskStatus := range allTaskStatus {
		allTaskStatusApiModel.AllTaskStatus = append(allTaskStatusApiModel.AllTaskStatus, ToTaskStatusApiModel(taskStatus))
	}
	return allTaskStatusApiModel
}

func ToTaskStatusApiModel(taskStatus dbModel.TaskStatus) apiModel.TaskStatus {

	return apiModel.TaskStatus{
		StatusId: taskStatus.StatusId,
		Status:   taskStatus.Status,
	}
}
