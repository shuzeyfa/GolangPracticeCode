package data

import "TaskManagement/models"

// Global slice to persist tasks
var tasks = []models.Task{
	{ID: 1, Title: "Task 1", Description: "Description for Task 1", DueDate: "2024-06-30", Status: "Pending"},
	{ID: 2, Title: "Task 2", Description: "Description for Task 2", DueDate: "2024-07-15", Status: "In Progress"},
}

func GetTasks() *[]models.Task {
	return &tasks
}

func GetTaskByID(id int) *models.Task {
	for _, task := range tasks {
		if task.ID == id {
			return &task
		}
	}
	return nil
}

func CreateTask(newTask *models.Task) models.Task {
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, *newTask)
	return *newTask
}

func UpdateTask(id int, updatedTask models.Task) bool {
	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id // keep ID consistent
			tasks[i] = updatedTask
			return true
		}
	}
	return false
}

func DeleteTask(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
