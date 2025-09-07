package services

import (
	"fmt"
	"auspire/models"
)

type FortuneService struct {
	aiClient *AIClient
}

func NewFortuneService() *FortuneService {
	return &FortuneService{
		aiClient: NewAIClient(),
	}
}

func (s *FortuneService) AnalyzeFortune(req models.FortuneRequest) (*models.FortuneResponse, error) {
	if len(req.Bazi) != 4 {
		return &models.FortuneResponse{
			Name:  req.Name,
			Error: "生辰八字信息不完整",
		}, fmt.Errorf("生辰八字信息不完整")
	}

	// 调用AI接口分析运势
	response, err := s.aiClient.AnalyzeFortune(req)
	if err != nil {
		return &models.FortuneResponse{
			Name:  req.Name,
			Error: "运势分析失败: " + err.Error(),
		}, err
	}

	return response, nil
}

// 新增：十二长生人生阶段分析
func (s *FortuneService) AnalyzeLifeStages(name string, shiErChangSheng map[string]string, bazi []models.BaziColumn) (map[string]string, error) {
	if len(shiErChangSheng) == 0 {
		return nil, fmt.Errorf("十二长生信息不完整")
	}

	// 调用AI接口分析人生各阶段
	analysis, err := s.aiClient.AnalyzeChangSheng(name, shiErChangSheng, bazi)
	if err != nil {
		return nil, fmt.Errorf("人生阶段分析失败: %v", err)
	}

	return analysis, nil
}