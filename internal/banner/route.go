package banner

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, requireAuth gin.HandlerFunc,
	requireRole func(...string) gin.HandlerFunc,) {
	banner := router.Group("/banners")
	{
		banner.GET("/", requireAuth, requireRole("admin"), handler.GetAll)
		banner.GET("/:id", requireAuth, requireRole("admin"), handler.GetByID)
		banner.POST("/", requireAuth, requireRole("admin"), handler.Create)
		banner.PUT("/:id", requireAuth, requireRole("admin"), handler.Update)
		banner.DELETE("/:id", requireAuth, requireRole("admin"), handler.Delete)
	}
}