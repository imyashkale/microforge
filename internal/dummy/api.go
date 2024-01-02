package dummy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imyashkale/microforge/pkg/log"
)

type DummyResource struct {
	service Service
	logger  log.Logger
}

// NewResource
func NewResource(service Service, logger log.Logger) DummyResource {
	return DummyResource{}
}

// RegisterURLsHandler
func RegisterDummyHandler(r *gin.RouterGroup, service Service, logger log.Logger) {
	urlResource := NewResource(service, logger)

	r = r.Group("/dummy")
	{
		r.GET("/", urlResource.Get)
		r.POST("/", urlResource.Create)
	}

}

// Get
func (res *DummyResource) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	var err error
	var dm Dummy
	if dm, err = res.service.Get(ctx, id); err != nil {
		res.logger.Errorf("error getting ulr %s", err.Error())
		c.JSON(http.StatusInternalServerError, dm)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": dm,
	})
}

// Create
func (res *DummyResource) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var cr CreateDummy
	var err error

	var url Dummy
	if url, err = res.service.Create(ctx, cr); err != nil {
		res.logger.Errorf("error getting ulr %s", err.Error())
		c.JSON(http.StatusInternalServerError, url)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": url,
	})
}
