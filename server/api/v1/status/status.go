package status

import (
	"github.com/Shachindra/altlife/server/api/types"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/status")
	{
		g.GET("", status)
	}
}

func status(c *gin.Context) {
	status := types.ApiResponse{
		Status:  200,
		Result: "ONLINE",
	}
	c.JSON(200, status)
}
