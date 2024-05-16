package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/retinacodeworks/regis/internal/storage"
	"time"
)

type Modules struct {
	db    *sqlx.DB
	files storage.Storage
}

func (m *Modules) Register(r *gin.Engine) {
	modules := r.Group("/v1/modules")
	modules.GET("/:namespace/:name/:provider/:version/download", m.DownloadVersion)
	modules.GET("/:namespace/:name/:provider/download", m.DownloadLatestVersion)
	modules.GET("/tarball/:namespace/:name/:provider/:version/*.tamodules.gz", m.GetArchive)
	modules.GET("/search", m.Search)
	modules.GET("/", m.List)
	modules.GET("/:namespace", m.List)
	modules.GET("/:namespace/:name/:provider/versions", m.ListVersions)
	modules.GET("/:namespace/:name", m.ListAllLatest)
	modules.POST("/:namespace/:name/:provider/:version", m.RegisterModule)
	modules.GET("/:namespace/:name/:provider/:version", m.Get)
	modules.GET("/:namespace/:name/:provider", m.LatestForProvider)
}

type Module struct {
	Id          int       `db:"id"`
	Namespace   string    `db:"namespace"`
	Name        string    `db:"name"`
	Provider    string    `db:"provider"`
	Version     string    `db:"version"`
	Owner       string    `db:"owner"`
	Location    string    `db:"location"`
	Definition  string    `db:"definition"`
	Downloads   int       `db:"downloads"`
	PublishedAt time.Time `db:"published_at"`
	Root        *any      `db:"root"`
	Submodules  *any      `db:"submodules"`
}

func (m *Modules) DownloadVersion(c *gin.Context) {
	// Get module from database
	// set x-terraform-get header to version/*.tar.gz
}

func (m *Modules) DownloadLatestVersion(c *gin.Context) {
	// get latest version of module from database
	// redirect to DownloadVersion
}

func (m *Modules) GetArchive(c *gin.Context) {
	// get module from database
	// get file from storage
	// stream attachment back to client
	// increment download counter
}

func (m *Modules) Search(c *gin.Context) {
	//check ?q=
	// check namespace, name, provider, etc
	// find all modules that match from DB
}

func (m *Modules) List(c *gin.Context) {
	// Find all modules with optional :namespace
	// as well as ?q option
}

func (m *Modules) ListVersions(c *gin.Context) {
	// Find all versions of the given module
}

func (m *Modules) ListAllLatest(c *gin.Context) {

}

func (m *Modules) RegisterModule(c *gin.Context) {

}

func (m *Modules) Get(c *gin.Context) {

}

func (m *Modules) LatestForProvider(c *gin.Context) {

}
