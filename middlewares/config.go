package middlewares

import (
	"github.com/adrien3d/jump-technical-test/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// ConfigMiddleware allows to use viper configuration parameters set in .env files
func ConfigMiddleware(viper *viper.Viper) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.ToContext(c, config.New(viper))
		c.Next()
	}
}
