package api

import (
	"github.com/Shachindra/altlife/server/api/types"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/deploy")
	{
		v1.GET("", deploy)
	}
}

func deploy(c *gin.Context) {
	status := types.ApiResponse{
		Status: 200,
		Result: "Deployment Success",
	}
	c.JSON(200, status)
}
