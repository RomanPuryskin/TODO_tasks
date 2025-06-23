package models

import "time"

// Task описывает полную модель задачи
// @Description Полная модель задачи
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Created_At  time.Time `json:"created_at" example:"2023-05-15T10:00:00Z" format:"date-time"`
	Updated_At  time.Time `json:"updated_at" example:"2023-05-15T10:00:00Z" format:"date-time"`
}

// CreateTaskRequst описывает упрощенную модель задачи для добавления
// @Description Модель задачи для запроса на создание
type CreateTaskRequst struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UpdateTaskRequst описывает упрощенную модель задачи для обновления статуса
// @Description Модель задачи для запроса на обновление
type UpdateTaskRequst struct {
	Status string `json:"status"`
}
