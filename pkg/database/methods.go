package database

import (
	"context"
	"fmt"
	"time"

	"github.com/TODO_task/pkg/models"
)

func CreateNewTaskInTable(ctx context.Context, task *models.CreateTaskRequst) error {
	query := `INSERT INTO tasks (title, description) VALUES ($1, $2)`
	_, err := DB.Exec(ctx, query, task.Title, task.Description)
	if err != nil {
		return fmt.Errorf("ошибка создания задачи: %w", err)
	}
	return nil
}

func GetAllTasksFromTable(ctx context.Context, tasks *[]*models.Task) error {
	query := `SELECT title , description , status , created_at , updated_at FROM tasks`
	rows, err := DB.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}

	for rows.Next() {
		var curTask models.Task
		err := rows.Scan(&curTask.Title, &curTask.Description, &curTask.Status, &curTask.Created_At, &curTask.Updated_At)
		if err != nil {
			return fmt.Errorf("ошибка чтения задачи: %w", err)
		}

		*tasks = append(*tasks, &curTask)
	}

	return nil
}

func UpdateTaskInTable(ctx context.Context, id int, task *models.UpdateTaskRequst, time time.Time) error {
	_, err := DB.Exec(ctx, "UPDATE tasks SET status = $1 , updated_at = $2 WHERE id = $3", task.Status, time, id)
	if err != nil {
		return fmt.Errorf("ошибка обновления задачи: %w", err)
	}
	return nil
}

func DeleteTaskFromTable(ctx context.Context, id int) error {
	_, err := DB.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("ошибка удаления задачи: %w", err)
	}
	return nil
}

func CheckTaskExistByTitleInTable(ctx context.Context, task *models.CreateTaskRequst) error {
	var count int
	if err := DB.QueryRow(ctx, "SELECT COUNT(*) FROM tasks WHERE title = $1", task.Title).Scan(&count); err != nil {
		return fmt.Errorf("ошибка обращения к задаче: %w", err)
	}

	if count != 0 {
		return fmt.Errorf("задача с таким title уже существует")
	}

	return nil
}

func CheckTaskExistByIdInTable(ctx context.Context, id int) error {
	var count int
	if err := DB.QueryRow(ctx, "SELECT COUNT(*) FROM tasks WHERE id = $1", id).Scan(&count); err != nil {
		return fmt.Errorf("ошибка обращения к задаче: %w", err)
	}

	if count == 0 {
		return fmt.Errorf("задачи с таким id не существует")
	}

	return nil
}
