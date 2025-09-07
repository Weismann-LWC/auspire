package services

import (
)

// BaziyuceService 四柱八字综合分析服务
type BaziyuceService struct {
}

// NewBaziyuceService 创建新的四柱八字综合分析服务实例
func NewBaziyuceService() *BaziyuceService {
	return &BaziyuceService{}
}

// Analyze 四柱八字综合分析
func (s *BaziyuceService) Analyze(bazi []models.BaziColumn) *models.BaziyuceResult {
	result := &models.BaziyuceResult{
		Steps: []models.AnalysisStep{},
	}

	// 第一步：排盘与定盘
	step1 := s.step1PaiPanDingPan(bazi)
	result.Steps = append(result.Steps, step1)

	// 第二步：定旺衰，识体性
	step2 := s.step2DingWangShuai(bazi)
	result.Steps = append(result.Steps, step2)

	// 第三步：明喜忌，定方向
	step3 := s.step3MingXiJi(bazi, step2)
	result.Steps = append(result.Steps, step3)

