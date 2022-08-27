package handler

import (
	"github.com/crutchm/notes-web-interface/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const authorisationHeader = "Authorization"

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		middleware.newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		middleware.newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		middleware.newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.services.Authorisation.ParseToken(headerParts[1])
	if err != nil {
		middleware.newErrorResponse(c, http.StatusUnauthorized, err.Error)
		return
	}

	c.Set("userId", userId)
}
