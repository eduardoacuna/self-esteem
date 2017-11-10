package tasks

import (
	"fmt"

	"github.com/eduardoacuna/self-esteem/estimates"
)

// Task denotes a time estimated task
type Task struct {
	ID         string            `json:"id"`
	UserID     string            `json:"userId"`
	Title      string            `json:"title"`
	Estimation int               `json:"estimation"`
	Done       bool              `json:"done"`
	Outcome    estimates.Outcome `json:"outcome,omitempty"`
}

// GetAllTasks returns all self-esteem tasks
func GetAllTasks(userID string) ([]*Task, error) {
	tasks, ok := dummyTaskData[userID]
	if !ok {
		return nil, fmt.Errorf("user with id %v not found", userID)
	}
	return tasks, nil
}

// GetTaskByID returns the task with the given ID
func GetTaskByID(userID string, taskID string) (*Task, error) {
	tasks, ok := dummyTaskData[userID]
	if !ok {
		return nil, fmt.Errorf("user with id %v not found", userID)
	}
	for _, task := range tasks {
		if task.ID == taskID {
			return task, nil
		}
	}
	return nil, fmt.Errorf("task with id %v not found", taskID)
}
