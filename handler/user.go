package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/govies/framework/logger"
	"github.com/govies/framework/resp"
	"github.com/govies/onboard/domain"
)

type UserHandler struct {
	UserEntity domain.UserEntity
	Logger *logger.Logger
}

func NewUsersHandler(r *gin.RouterGroup, l *logger.Logger, us domain.UserEntity) {
	handler := &UserHandler{
		UserEntity: us,
		Logger: l,
	}
	r.GET("/users", handler.FindUsers)
	r.GET("/users/:id", handler.FindUser)
}

func (a *UserHandler) FindUsers(c *gin.Context) {
	users, err := a.UserEntity.Fetch(c.Request.Context())
	if err != nil {
		resp.Error(400, err).Send(c)
		return
	}
	resp.Success(200, users).Send(c)
}

func (a *UserHandler) FindUser(c *gin.Context) {
	users, err := a.UserEntity.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		resp.Error(400, err).Send(c)
		return
	}
	resp.Success(200, users).Send(c)
}