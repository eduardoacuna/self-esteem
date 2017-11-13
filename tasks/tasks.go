package tasks

import (
	"context"

	"github.com/eduardoacuna/self-esteem/database"
	"github.com/eduardoacuna/self-esteem/estimates"
	"github.com/eduardoacuna/self-esteem/log"
)

// Task denotes a time estimated task
type Task struct {
	ID         int               `json:"id"`
	UserID     int               `json:"userId"`
	Title      string            `json:"title"`
	Estimation int               `json:"estimation"`
	Done       bool              `json:"done"`
	Outcome    estimates.Outcome `json:"outcome,omitempty"`
}

// GetAllTasks returns all self-esteem tasks
func GetAllTasks(ctx context.Context, userID int) ([]*Task, error) {
	log.Debug(ctx, "get all tasks for user", "userID", userID)
	conn := database.Connection()
	q := "SELECT * FROM tasks WHERE user_id=$1"
	rows, err := conn.Query(q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*Task{}

	for rows.Next() {
		task := &Task{}
		if err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Estimation, &task.Done, &task.Outcome); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	err = rows.Err()
	if err != nil {
		return nil, rows.Err()
	}

	return tasks, nil
}

// GetTaskByID returns the task with the given ID
func GetTaskByID(ctx context.Context, userID int, taskID int) (*Task, error) {
	log.Debug(ctx, "get user task by id", "taskID", taskID, "userID", userID)
	conn := database.Connection()
	q := "SELECT * FROM tasks WHERE id=$1 AND user_id=$2"
	row := conn.QueryRow(q, taskID, userID)
	task := &Task{}
	err := row.Scan(&task.ID, &task.UserID, &task.Title, &task.Estimation, &task.Done, &task.Outcome)
	if err != nil {
		return nil, err
	}
	return task, nil
}
