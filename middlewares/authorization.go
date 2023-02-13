package middlewares

import (
	"github.com/adrien3d/jump-technical-test/config"
	"github.com/adrien3d/jump-technical-test/helpers"
	"github.com/adrien3d/jump-technical-test/models"
	"github.com/adrien3d/jump-technical-test/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

// AuthorizationMiddleware allows granting access to a resource or not
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenReader := c.Request.Header.Get("Authorization")

		authHeaderParts := strings.Split(tokenReader, " ")
		encodedKey := []byte(config.GetString(c, "rsa_private"))
		claims, _ := helpers.ValidateJwtToken(authHeaderParts[1], encodedKey, "access")
		ctx := store.AuthContext(c)
		user, _ := models.GetUser(ctx, bson.M{"_id": claims["sub"]})
		c.Set(store.CurrentUserKey, user)

		c.Next()
	}
}
