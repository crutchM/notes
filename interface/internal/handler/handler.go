package handler

import (
	"github.com/crutchm/notes-core/logic"
)

type Handler struct {
	services *logic.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
