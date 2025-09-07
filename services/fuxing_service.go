package services

import "auspire/models"

// FuXingService 副星服务
type FuXingService struct {
	zhuXingService *ZhuXingService
	cangGanService *CangGanService
}

func NewFuXingService() *FuXingService {
	return &FuXingService{
		zhuXingService: NewZhuXingService(),
		cangGanService: NewCangGanService(),
	}
}

// Calculate 计算副星（基于地支藏干的十神关系）
func (s *FuXingService) Calculate(dayGan string, column models.BaziColumn) []string {
	var fuXing []string

	// 分析地支藏干的十神关系
	cangGan := s.cangGanService.Calculate(column.Zhi)
	for _, cGan := range cangGan {
		zhuXing := s.zhuXingService.Calculate(dayGan, cGan)
		if zhuXing != "" {
			fuXing = append(fuXing, zhuXing)
		}
	}

	return fuXing
}