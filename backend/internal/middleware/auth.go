package middleware

import (
	"context"
	"net/http"
	"strconv"

	"backend/pkg/common"
)

type contextKey string

const UserIDKey contextKey = "userID"

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIDStr := r.Header.Get("User-ID")
		if userIDStr == "" {
			common.JSONError(w, http.StatusUnauthorized, "Authentication header missing")
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			common.JSONError(w, http.StatusUnauthorized, "Invalid user token format")
			return
		}

		if userID <= 0 {
			common.JSONError(w, http.StatusUnauthorized, "Invalid user ID value")
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
