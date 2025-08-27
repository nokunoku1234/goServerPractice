package handler

import (
	"goServerPractice/internal/config"

	"github.com/uptrace/bun"
)

type Handler struct {
	db  *bun.DB
	cfg config.Config
}

func New(db *bun.DB, cfg config.Config) *Handler { return &Handler{db: db, cfg: cfg} }
