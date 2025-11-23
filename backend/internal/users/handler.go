package users

import (
	"backend/pkg/common"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			ID:        1,
			FirstName: "Shaun",
			Email:     "Shaun@email.com",
		},
		{
			ID:        2,
			FirstName: "Christina",
			Email:     "Christina@email.com",
		},
	}

	common.JSONSuccess(w, users, http.StatusOK)
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		common.JSONError(w, http.StatusBadRequest, "Invalid user ID format")
		return
	}

	userData := map[int]User{
		1: {
			ID:        1,
			FirstName: "Shaun",
			Email:     "Shaun@email.com",
		},
		2: {
			ID:        2,
			FirstName: "Christina",
			Email:     "Christina@email.com",
		},
	}

	if userData[userID].ID == 0 {
		common.JSONError(w, http.StatusNotFound, "User not found")
		return
	}

	common.JSONSuccess(w, userData[userID], http.StatusOK)
}
