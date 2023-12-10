package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"detail": "Users",
	})
}

func (h *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"detail": "User by Id",
		"id":     id,
	})
}

func (h *TestHandler) UserByUsername(c *gin.Context) {
	id := c.Param("id")
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"detail":   "User By Username",
		"username": username,
		"id":       id,
	})
}

func (h *TestHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"detail": "Created User",
	})
}
