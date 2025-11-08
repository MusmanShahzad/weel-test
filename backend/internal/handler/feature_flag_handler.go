package handler
import (
	"net/http"
	"weel-backend/internal/service"
	"github.com/gin-gonic/gin"
)
type FeatureFlagHandler struct {
	flagService service.FeatureFlagService
}
func NewFeatureFlagHandler(flagService service.FeatureFlagService) *FeatureFlagHandler {
	return &FeatureFlagHandler{flagService: flagService}
}
func (h *FeatureFlagHandler) RegisterRoutes(router *gin.RouterGroup) {
	flags := router.Group("/feature-flags")
	{
		flags.GET("", h.GetAllFlags)
		flags.GET("/:name", h.GetFlagByName)
	}
}
func (h *FeatureFlagHandler) GetAllFlags(c *gin.Context) {
	flags, err := h.flagService.GetAllFlags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch feature flags"})
		return
	}
	flagMap := make(map[string]bool)
	for _, flag := range flags {
		flagMap[flag.Name] = flag.Enabled
	}
	c.JSON(http.StatusOK, gin.H{
		"flags":  flagMap,
		"details": flags,
	})
}
func (h *FeatureFlagHandler) GetFlagByName(c *gin.Context) {
	name := c.Param("name")
	flag, err := h.flagService.GetFlagByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "feature flag not found"})
		return
	}
	c.JSON(http.StatusOK, flag)
}
