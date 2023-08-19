package routers

import (
	"net/http"

	"github.com/doffy007/book-library/internal/handler"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Register(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
	handlers := handler.NewHandler(lg, db)

	r.HandleFunc("/health", handlers.Health()).Methods(http.MethodGet)
}
