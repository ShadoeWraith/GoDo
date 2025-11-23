package tasks

import (
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/pkg/common"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func getUserIDFromContext(r *http.Request) (int, error) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)

	if !ok || userID <= 0 {
		return 0, errors.New("user ID not found in context (Unauthorized)")
	}

	return userID, nil
}

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)

	if err != nil {
		common.JSONError(w, http.StatusUnauthorized, "Authentication required.")
		return
	}

	tasks, err := ListAllTasks(userID)

	if err != nil {
		common.JSONError(w, http.StatusInternalServerError, "Database error retrieving tasks")
		return
	}

	common.JSONSuccess(w, tasks, http.StatusOK)
}

func GetSingleTaskHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)

	if err != nil {
		common.JSONError(w, http.StatusUnauthorized, "Authentication required.")
		return
	}

	taskIDStr := chi.URLParam(r, "taskID")

	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		common.JSONError(w, http.StatusBadRequest, "Invalid user ID format")
		return
	}

	task, err := GetSingleTask(userID, taskID)

	if err != nil {
		switch err {
		case ErrTaskNotFound:
			common.JSONError(w, http.StatusNotFound, "Task not found")
			return

		default:
			common.JSONError(w, http.StatusInternalServerError, "Internal server error retrieving task")
			return
		}
	}

	common.JSONSuccess(w, task, http.StatusOK)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task

	userID, err := getUserIDFromContext(r)
	if err != nil {
		common.JSONError(w, http.StatusUnauthorized, "Authentication required.")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		common.JSONError(w, http.StatusBadRequest, "Invalid request body format")
		return
	}

	newTask.UserID = userID

	createdTask, err := CreateTask(newTask)

	if err != nil {
		switch err {
		case ErrTitleRequired, ErrTitleTooShort:
			common.JSONError(w, http.StatusUnprocessableEntity, err.Error())
			return

		default:
			common.JSONError(w, http.StatusInternalServerError, "Failed to create task")
			return
		}
	}

	common.JSONSuccess(w, createdTask, http.StatusCreated)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "taskID")

	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		common.JSONError(w, http.StatusBadRequest, "Invalid user ID format")
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		common.JSONError(w, http.StatusUnauthorized, "Authentication required.")
		return
	}

	var inputTask models.Task

	if err := json.NewDecoder(r.Body).Decode(&inputTask); err != nil {
		common.JSONError(w, http.StatusBadRequest, "Invalid request body format")
		return
	}

	inputTask.ID = taskID
	inputTask.UserID = userID

	updatedTask, err := UpdateTask(inputTask)

	if err != nil {
		switch err {
		case ErrTitleRequired, ErrTitleTooShort:
			common.JSONError(w, http.StatusUnprocessableEntity, err.Error())
			return
		case ErrTaskNotFound:
			common.JSONError(w, http.StatusNotFound, "Task not found")
			return
		default:
			common.JSONError(w, http.StatusInternalServerError, "Failed to update task")
			return
		}
	}

	common.JSONSuccess(w, updatedTask, http.StatusOK)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "taskID")

	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		common.JSONError(w, http.StatusBadRequest, "Invalid user ID format")
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		common.JSONError(w, http.StatusUnauthorized, "Authentication required.")
		return
	}

	err = DeleteTask(userID, taskID)

	if err != nil {
		switch err {
		case ErrTaskNotFound:
			common.JSONError(w, http.StatusNotFound, "Task not found")
			return
		default:
			common.JSONError(w, http.StatusInternalServerError, "Failed to delete task due to server error")
			return
		}
	}

	responseMessage := map[string]interface{}{
		"message":   fmt.Sprintf("Item ID %d deleted successfully", taskID),
		"deletedId": taskID,
	}

	common.JSONSuccess(w, responseMessage, http.StatusOK)
}

func ListDeletedTasksHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		common.JSONError(w, http.StatusUnauthorized, "Authentication required.")
		return
	}

	tasks, err := ListDeletedTasks(userID)

	if err != nil {
		common.JSONError(w, http.StatusInternalServerError, "Failed to retreive deleted tasks")
		return
	}

	common.JSONSuccess(w, tasks, http.StatusOK)
}

func RestoreDeletedTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "taskID")

	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		common.JSONError(w, http.StatusBadRequest, "Invalid user ID format")
		return
	}

	userID, err := getUserIDFromContext(r)
	if err != nil {
		common.JSONError(w, http.StatusUnauthorized, "Authentication required.")
		return
	}

	err = RestoreDeletedTask(userID, taskID)

	if err != nil {
		switch err {
		case ErrTaskNotFound:
			common.JSONError(w, http.StatusNotFound, "Task not found")
			return
		case errors.New("task is already active and cannot be restored"):
			common.JSONError(w, http.StatusConflict, err.Error())
			return
		default:
			common.JSONError(w, http.StatusInternalServerError, "Failed to restore task")
			return
		}
	}

	responseMessage := map[string]interface{}{
		"message":    fmt.Sprintf("Item ID %d restored successfully", taskID),
		"restoredId": taskID,
	}
	common.JSONResponse(w, responseMessage, http.StatusOK)
}
