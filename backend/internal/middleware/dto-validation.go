package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type contextKey string

const bodyKey contextKey = "body"
const responseKey contextKey = "response"

func ValidateBody[T any]() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var incomingDto T
			if err := json.NewDecoder(r.Body).Decode(&incomingDto); err != nil {
				fmt.Println("Unable to parse incoming json to DTO")
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			ctx := context.WithValue(r.Context(), bodyKey, incomingDto)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func BodyFromContext[T any](ctx context.Context, w http.ResponseWriter) (value T, ok bool) {
	val, ok := ctx.Value(bodyKey).(T)

	if !ok {
		fmt.Println("Unable to retrieve request DTO from context")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return val, false
	}

	return val, true
}
