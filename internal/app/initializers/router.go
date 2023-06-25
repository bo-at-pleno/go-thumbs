package initializers

import (
	"github.com/bo-at-pleno/go-thumbs/internal/app/dependencies"
	"github.com/bo-at-pleno/go-thumbs/internal/gateways/web/controllers/apiv1"
	apiv1Status "github.com/bo-at-pleno/go-thumbs/internal/gateways/web/controllers/apiv1/status"
	apiv1Swagger "github.com/bo-at-pleno/go-thumbs/internal/gateways/web/controllers/apiv1/swagger"
	"github.com/bo-at-pleno/go-thumbs/internal/gateways/web/router"
	"github.com/gin-gonic/gin"
)

// InitializeRouter initializes new gin router
func InitializeRouter(container *dependencies.Container) *gin.Engine {
	r := router.NewRouter()

	ctrls := buildControllers(container)

	for i := range ctrls {
		ctrls[i].DefineRoutes(r)
	}

	return r
}

func buildControllers(container *dependencies.Container) []apiv1.Controller {
	return []apiv1.Controller{
		apiv1Status.NewController(container.BuildInfo),
		apiv1Swagger.NewController(),
	}
}
