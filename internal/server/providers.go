package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/retinacodeworks/regis/internal/storage"
	"net/http"
	"time"
)

type Providers struct {
	db    *sqlx.DB
	files storage.Storage
}

func (p *Providers) Register(r *gin.Engine) {
	prov := r.Group("/v1/providers")
	prov.POST("/:namespace/:type/:version", p.RegisterProvider)
	prov.GET("/:namespace/:type/versions", p.ListAvailableVersions)
	prov.GET("/:namespace/:type/:version/download/:os/:arch", p.Find)
	prov.GET("/:namespace/:type/:version/download/:os/:arch/zip", p.Download)
}

type Provider struct {
	Id            int       `db:"id"`
	Namespace     string    `db:"namespace"`
	Type          string    `db:"type"`
	Version       string    `db:"version"`
	Protocols     any       `db:"protocols"`
	Platforms     any       `db:"platforms"`
	GpgPublicKeys any       `db:"gpg_public_keys"`
	PublishedAt   time.Time `db:"published_at"`
}

func (p *Providers) RegisterProvider(c *gin.Context) {

}

func (p *Providers) ListAvailableVersions(c *gin.Context) {
	providers := []Provider{}
	err := p.db.Select(
		&providers,
		"SELECT * FROM providers WHERE namespace = $1 and type = $2 ORDER BY id ASC",
		c.Param("namespace"),
		c.Param("type"),
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"versions": toProviderVersions(providers),
	})
}

func (p *Providers) Find(c *gin.Context) {

}

func (p *Providers) Download(c *gin.Context) {

}

func toProviderVersions(providers []Provider) any {
	var result []interface{}
	for _, provider := range providers {
		result = append(result, gin.H{
			"version":   provider.Version,
			"platforms": provider.Platforms,
			"protocols": provider.Protocols,
		})
	}
	return result
}
