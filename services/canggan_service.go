package services

// CangGanService 藏干服务
type CangGanService struct{}

var (
	// 地支藏干对照表
	cangGanData = map[string][]string{
		"子": {"癸"},
		"丑": {"己", "辛", "癸"},
		"寅": {"甲", "丙", "戊"},
		"卯": {"乙"},
		"辰": {"戊", "乙", "癸"},
		"巳": {"丙", "戊", "庚"},
		"午": {"丁", "己"},
		"未": {"己", "丁", "乙"},
		"申": {"庚", "壬", "戊"},
		"酉": {"辛"},
		"戌": {"戊", "辛", "丁"},
		"亥": {"壬", "甲"},
	}
)

func NewCangGanService() *CangGanService {
	return &CangGanService{}
}

// Calculate 计算地支藏干
func (s *CangGanService) Calculate(zhi string) []string {
	if cangGan, exists := cangGanData[zhi]; exists {
		// 返回副本避免外部修改
		result := make([]string, len(cangGan))
		copy(result, cangGan)
		return result
	}
	return []string{}
}

// GetAllCangGan 获取所有藏干数据（用于其他服务）
func (s *CangGanService) GetAllCangGan() map[string][]string {
	return cangGanData
}