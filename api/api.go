package api

import (
	"fmt"
	"fubuki-go/cmd"
	"strconv"
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

func Run() error {
	err := app.Run(":" + strconv.Itoa(cmd.ApiListingPort))
	return fmt.Errorf("api server exited, %w", err)
}
