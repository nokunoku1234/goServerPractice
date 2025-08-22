package handler

import (
	"goServerPractice/ent"
	"goServerPractice/internal/config"
)

type Handler struct {
	db  *ent.Client
	cfg config.Config
}

func New(db *ent.Client, cfg config.Config) *Handler { return &Handler{db: db, cfg: cfg} }
