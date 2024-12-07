package apiv1

import (
	deploy "github.com/Shachindra/altlife/server/api/v1/deploy"
	"github.com/Shachindra/altlife/server/api/v1/status"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		status.ApplyRoutes(v1)
		deploy.ApplyRoutes(v1)
	}
}