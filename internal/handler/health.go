package handler

import (
	"net/http"

	"github.com/doffy007/book-library/internal/helper"
)

func (h handler) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helper.Respond(w, "OK", http.StatusOK)
	}
}
