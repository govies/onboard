package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/govies/framework/logger"
	"github.com/govies/onboard/entity"
	"github.com/govies/onboard/handler"
	"github.com/govies/onboard/repository"
)

func Initialize(e *gin.Engine, l *logger.Logger) {
	g := e.Group("/onboard/api/v1/")
	handler.NewUsersHandler(g, l, entity.NewUserEntity(repository.NewUserRepository(l)))
}