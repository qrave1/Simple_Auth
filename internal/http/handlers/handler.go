package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"rchir7/internal/auth"
	r "rchir7/internal/db"
	"rchir7/internal/model"
)

var storage []model.User

type Handler struct {
	a *auth.TokenHandler
	r *r.RedisDb
}

func NewHandler(a *auth.TokenHandler, r *r.RedisDb) *Handler {
	return &Handler{a: a, r: r}
}

func (h *Handler) SignUp(c *gin.Context) {
	var u model.User
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	if u.Role == "USER" || u.Role == "SELLER" || u.Role == "ADMINISTRATOR" {
		storage = append(storage, u)
		c.JSON(http.StatusCreated, gin.H{
			"status": "OK",
		})
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"error":  fmt.Sprintf("invalid role %s", u.Role),
	})
	return
}

func (h *Handler) SignIn(c *gin.Context) {
	var u model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	for _, v := range storage {
		if v.Email == u.Email && v.Password == u.Password {
			var token string
			if token, err = h.r.Read(u.Email); err == redis.Nil {
				token, err = h.a.GenerateToken(v)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status": "error",
						"error":  err.Error(),
					})
					return
				}
				err = h.r.Insert(u.Email, token)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status": "error",
						"error":  err.Error(),
					})
					return
				}
			}

			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
				"token":  token,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "error",
		"error":  fmt.Sprintf("user %s not found in storage", u.Email),
	})
}
