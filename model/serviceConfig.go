package model

import (
	"github.com/gin-gonic/gin"
)

type ServiceConfig struct {
	GinRouter *gin.Engine
}
