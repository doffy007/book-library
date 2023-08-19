package handler

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type handler struct {
	logger *logrus.Logger
	router *mux.Router
}

func NewHandler(lg *logrus.Logger, db *sqlx.DB) handler {
	return handler{
		logger: lg,
	}
}
