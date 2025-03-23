package http

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Handler struct {
    // Define any dependencies your handler needs here
}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) SetupRoutes(router *gin.Engine) {
    router.GET("/example", h.ExampleHandler)
}

func (h *Handler) ExampleHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}