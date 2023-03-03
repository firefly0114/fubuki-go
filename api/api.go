package api

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var app = gin.New()

func init() {
	app.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	app.Use(ginzap.RecoveryWithZap(zap.L(), true))
}

func Run() {
	app.Run()
}
