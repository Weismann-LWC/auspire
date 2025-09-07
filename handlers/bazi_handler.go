package handlers

import (
	"auspire/models"
	"auspire/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaziHandler struct {
	baziService      *services.BaziService
	fortuneService   *services.FortuneService
	xiyongshenService *services.XiYongShenService
	baziyuceService   *services.BaziyuceService
}

func NewBaziHandler() *BaziHandler {
	return &BaziHandler{
		baziService:       services.NewBaziService(),
		fortuneService:    services.NewFortuneService(),
		xiyongshenService: services.NewXiYongShenService(),
		baziyuceService:   services.NewBaziyuceService(),
	}
}

func (h *BaziHandler) CalculateBazi(c *gin.Context) {
	var req models.BaziRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.BaziResponse{
			Error: "请求数据格式错误: " + err.Error(),
		})
		return
	}

	response, err := h.baziService.CalculateBazi(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *BaziHandler) AnalyzeFortune(c *gin.Context) {
	var req models.FortuneRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.FortuneResponse{
			Error: "请求数据格式错误: " + err.Error(),
		})
		return
	}

	response, err := h.fortuneService.AnalyzeFortune(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *BaziHandler) AnalyzeLifeStages(c *gin.Context) {
	var req models.LifeStageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.LifeStageResponse{
			Error: "请求数据格式错误: " + err.Error(),
		})
		return
	}

	analysis, err := h.fortuneService.AnalyzeLifeStages(req.Name, req.ShiErChangSheng, req.Bazi)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.LifeStageResponse{
			Name:  req.Name,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.LifeStageResponse{
		Name:     req.Name,
		Analysis: analysis,
	})
}

func (h *BaziHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, models.HealthResponse{
		Status:  "healthy",
		Service: "auspire-bazi",
	})
}

// ServePaiPanPage 服务排盘页面
func (h *BaziHandler) ServePaiPanPage(c *gin.Context) {
	c.File("./static/paipan.html")
}

// CalculateXiYongShen 计算喜用神
func (h *BaziHandler) CalculateXiYongShen(c *gin.Context) {
	var req models.XiYongShenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.XiYongShenResult{
			Error: "请求数据格式错误: " + err.Error(),
		})
		return
	}

	result := h.xiyongshenService.Calculate(req.Bazi)
	result.Name = req.Name

	c.JSON(http.StatusOK, result)
}

// AnalyzeBaziyuce 四柱八字综合分析
func (h *BaziHandler) AnalyzeBaziyuce(c *gin.Context) {
	var req models.BaziyuceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.BaziyuceResult{
			Error: "请求数据格式错误: " + err.Error(),
		})
		return
	}

	result := h.baziyuceService.Analyze(req.Bazi)
	result.Name = req.Name

	c.JSON(http.StatusOK, result)
}