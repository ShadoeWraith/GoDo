package tasks

import (
	"backend/internal/db"
	"backend/internal/models"
	"errors"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrTaskNotFound = errors.New("task not found")
var ErrTitleRequired = errors.New("task title cannot be empty")
var ErrTitleTooShort = errors.New("task title must be at least 3 characters")

func ListAllTasks(userID int) ([]models.Task, error) {
	var tasks []models.Task

	result := db.DB.Where("user_id = ?", userID).Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func GetSingleTask(userID int, taskID uuid.UUID) (models.Task, error) {
	var task models.Task

	result := db.DB.Where("user_id = ?", userID).First(&task, taskID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Task{}, ErrTaskNotFound
		}
		return models.Task{}, result.Error
	}

	return task, nil
}

func CreateTask(newTask models.Task) (models.Task, error) {
	title := strings.TrimSpace(newTask.Title)

	if title == "" {
		return models.Task{}, ErrTitleRequired
	}
	if len(title) < 3 {
		return models.Task{}, ErrTitleTooShort
	}
	newTask.Title = title

	if newTask.UserID <= 0 {
		return models.Task{}, errors.New("cannot create task without a valid user ID")
	}

	result := db.DB.Create(&newTask)

	if result.Error != nil {
		return models.Task{}, result.Error
	}

	return newTask, nil
}

func UpdateTask(updatedTask models.Task) (models.Task, error) {
	title := strings.TrimSpace(updatedTask.Title)

	if title == "" {
		return models.Task{}, ErrTitleRequired
	}
	if len(title) < 3 {
		return models.Task{}, ErrTitleTooShort
	}
	updatedTask.Title = title

	result := db.DB.Model(&updatedTask).
		Clauses(clause.Returning{}).
		Where("user_id = ?", updatedTask.UserID).
		Where("id = ?", updatedTask.ID).
		Updates(updatedTask)

	if result.Error != nil {
		return models.Task{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.Task{}, ErrTaskNotFound
	}

	return updatedTask, nil
}

func DeleteTask(userID int, taskID uuid.UUID) error {
	var task models.Task

	result := db.DB.
		Where("id = ? AND user_id = ?", taskID, userID).
		Delete(&task)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func ListDeletedTasks(userID int) ([]models.Task, error) {
	var tasks []models.Task

	result := db.DB.Unscoped().
		Where("user_id = ?", userID).
		Where("deleted_at IS NOT NULL").
		Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func RestoreDeletedTask(userID int, taskID uuid.UUID) error {
	var task models.Task

	result := db.DB.Unscoped().Where("id = ? AND user_id = ?", taskID, userID).First(&task, taskID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ErrTaskNotFound
	}

	if result.Error != nil {
		return result.Error
	}

	if !task.DeletedAt.Valid {
		return errors.New("task is already active and cannot be restored")
	}

	restoreResult := db.DB.Unscoped().Model(&task).
		Where("id = ? AND user_id = ?", taskID, userID).
		Update("deleted_at", nil)

	if restoreResult.Error != nil {
		return restoreResult.Error
	}

	return nil
}
