package middlewares

import (
	"github.com/adrien3d/jump-technical-test/config"
	"github.com/adrien3d/jump-technical-test/store"
	"github.com/adrien3d/jump-technical-test/store/mongodb"
	"github.com/adrien3d/jump-technical-test/store/postgresql"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

// StoreMongoMiddleware allows to setup MongoDB database
func StoreMongoMiddleware(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println("Store mongo middleware:", db, c, mongodb.New(db, "things", c))
		//store.ToContext(c, mongodb.New(db, config.GetString(c, "mongo_db_name"), c))
		store.ToContext(c, mongodb.New(c, db, config.GetString(c, "mongo_db_name")))
		//c.Set(store.AppKey, user)
		c.Next()
	}
}

// StorePostgreMiddleware allows to setup SQL database
func StorePostgreMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		store.ToContext(c, postgresql.New(c, db, config.GetString(c, "postgres_db_name")))
		c.Next()
	}
}
