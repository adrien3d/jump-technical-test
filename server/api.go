package server

import (
	"database/sql"
	"github.com/adrien3d/jump-technical-test/services"
	"github.com/adrien3d/jump-technical-test/store"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// API structure that holds various necessary services
type API struct {
	Router          *gin.Engine
	Config          *viper.Viper
	Context         *store.Context
	MongoDatabase   *mongo.Database
	PostgreDatabase *gorm.DB
	MySQLDatabase   *sql.DB
	EmailSender     services.EmailSender
	TextSender      services.TextSender
}
