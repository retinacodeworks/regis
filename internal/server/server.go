package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/retinacodeworks/regis/internal/storage"
	"net/http"
)

func New(db *sqlx.DB, files storage.Storage) *gin.Engine {
	r := gin.Default()

	r.GET("/.well-known/terraform.json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"modules.v1":   "/v1/modules/",
			"providers.v1": "/v1/providers/",
		})
	})
	r.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	modules := &Modules{db, files}
	providers := &Providers{db, files}

	modules.Register(r)
	providers.Register(r)

	return r
}
