package source

import (
	"context"
	"time"

	"github.com/TODO_task/pkg/database"
	"github.com/TODO_task/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// варианты для статуса задачи
var statusVariations = []string{"new", "in_progress", "done"}

// getTasks godoc
// @Summary Получить все задачи
// @Description Получает все имеющиеся задачи
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Failure 400 {object} map[string]interface{} "'error': 'message'"
// @Failure 500 {object}  map[string]interface{} "'error': 'message'"
// @Router /tasks [get]
func getTasks(c *fiber.Ctx) error {
	tasks := []*models.Task{}

	// получим список задач
	if err := database.GetAllTasksFromTable(context.Background(), &tasks); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(tasks)
}

// createTask godoc
// @Summary Создать новую задачу
// @Description Создает новую задачу
// @Tags task
// @Accept json
// @Produce json
// @Param task body models.CreateTaskRequst true "Данные задачи"
// @Success 200 {object} map[string]interface{} "Задача успешно создана"
// @Failure 400 {object} map[string]interface{} "'error': 'message'"
// @Failure 500 {object}  map[string]interface{} "'error': 'message'"
// @Router /tasks [post]
func createTask(c *fiber.Ctx) error {
	var newTask models.CreateTaskRequst

	// парсим json в структуру newTask , ошибку вернем так же в формате JSON
	if err := c.BodyParser(&newTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	// проверим, существует ли уже задача с таким title
	if err := database.CheckTaskExistByTitleInTable(context.Background(), &newTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// запрос к БД
	if err := database.CreateNewTaskInTable(context.Background(), &newTask); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// успешный ответ
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Задача успешно добавлена",
	})
}

// updateTask godoc
// @Summary Обновить статус задачи
// @Description Обновляет статус задачи по id, Варианты: "new" , "in_progress" , "done"
// @Tags task
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param task body models.UpdateTaskRequst true "Новые данные"
// @Success 200 {object} map[string]interface{} "Задача успешно обновлена"
// @Failure 400 {object} map[string]interface{} "'error': 'message'"
// @Failure 500 {object}  map[string]interface{} "'error': 'message'"
// @Router /tasks/{id} [put]
func updateTask(c *fiber.Ctx) error {
	// провалидируем id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный id задачи"})
	}

	//проверим,существует ли задача с таким id
	if err := database.CheckTaskExistByIdInTable(context.Background(), id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var updatedTask models.UpdateTaskRequst

	// парсим json в структуру updatedTask
	if err := c.BodyParser(&updatedTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	// проверим новый статус, он должен быть правильно указан
	correctStatus := false
	for _, status := range statusVariations {
		if status == updatedTask.Status {
			correctStatus = true
		}
	}

	if !correctStatus {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный статус"})
	}

	// обновим данные в таблице (добавим так же текущее время в качестве параметра Updated_at)
	if err := database.UpdateTaskInTable(context.Background(), id, &updatedTask, time.Now()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	// успешный ответ
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Задача успешно обновлена",
	})
}

// deleteTask godoc
// @Summary Удалить задачу
// @Description Удаляет задачу по ID
// @Tags task
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} map[string]interface{} "Задача успешно удалена"
// @Failure 400 {object} map[string]interface{} "'error': 'message'"
// @Failure 500 {object}  map[string]interface{} "'error': 'message'"
// @Router /tasks/{id} [delete]
func deleteTask(c *fiber.Ctx) error {
	// провалидируем id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Неверный id задачи"})
	}

	//проверим,существует ли задача с таким id
	if err := database.CheckTaskExistByIdInTable(context.Background(), id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// удаление
	if err := database.DeleteTaskFromTable(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// успешный ответ
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Задача успешно удалена",
	})
}
