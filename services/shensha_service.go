package services

import "auspire/models"

// ShenShaService 神煞服务
type ShenShaService struct{}

var (
	// 神煞对照表
	shenShaData = map[string]map[string][]string{
		"甲": {
			"天乙贵人": {"丑", "未"}, "太极贵人": {"子", "午"}, "文昌贵人": {"巳"},
			"将星": {"子"}, "华盖": {"戌"}, "咸池": {"酉"}, "驿马": {"寅"}, "灾煞": {"午"},
		},
		"乙": {
			"天乙贵人": {"子", "申"}, "太极贵人": {"卯", "酉"}, "文昌贵人": {"午"},
			"将星": {"酉"}, "华盖": {"未"}, "咸池": {"午"}, "驿马": {"亥"}, "灾煞": {"卯"},
		},
		"丙": {
			"天乙贵人": {"亥", "酉"}, "太极贵人": {"卯", "酉"}, "文昌贵人": {"申"},
			"将星": {"午"}, "华盖": {"辰"}, "咸池": {"卯"}, "驿马": {"巳"}, "灾煞": {"子"},
		},
		"丁": {
			"天乙贵人": {"亥", "酉"}, "太极贵人": {"子", "午"}, "文昌贵人": {"酉"},
			"将星": {"卯"}, "华盖": {"丑"}, "咸池": {"子"}, "驿马": {"申"}, "灾煞": {"酉"},
		},
		"戊": {
			"天乙贵人": {"丑", "未"}, "太极贵人": {"卯", "酉"}, "文昌贵人": {"申"},
			"将星": {"午"}, "华盖": {"辰"}, "咸池": {"卯"}, "驿马": {"巳"}, "灾煞": {"子"},
		},
		"己": {
			"天乙贵人": {"子", "申"}, "太极贵人": {"子", "午"}, "文昌贵人": {"酉"},
			"将星": {"卯"}, "华盖": {"丑"}, "咸池": {"子"}, "驿马": {"申"}, "灾煞": {"酉"},
		},
		"庚": {
			"天乙贵人": {"丑", "未"}, "太极贵人": {"子", "午"}, "文昌贵人": {"亥"},
			"将星": {"子"}, "华盖": {"戌"}, "咸池": {"酉"}, "驿马": {"寅"}, "灾煞": {"午"},
		},
		"辛": {
			"天乙贵人": {"寅", "午"}, "太极贵人": {"卯", "酉"}, "文昌贵人": {"子"},
			"将星": {"酉"}, "华盖": {"未"}, "咸池": {"午"}, "驿马": {"亥"}, "灾煞": {"卯"},
		},
		"壬": {
			"天乙贵人": {"卯", "巳"}, "太极贵人": {"子", "午"}, "文昌贵人": {"寅"},
			"将星": {"午"}, "华盖": {"辰"}, "咸池": {"卯"}, "驿马": {"巳"}, "灾煞": {"子"},
		},
		"癸": {
			"天乙贵人": {"卯", "巳"}, "太极贵人": {"卯", "酉"}, "文昌贵人": {"卯"},
			"将星": {"卯"}, "华盖": {"丑"}, "咸池": {"子"}, "驿马": {"申"}, "灾煞": {"酉"},
		},
	}
)

func NewShenShaService() *ShenShaService {
	return &ShenShaService{}
}

// Calculate 计算神煞
func (s *ShenShaService) Calculate(dayGan string, bazi []models.BaziColumn) map[string]string {
	result := make(map[string]string)

	if shenShaConfig, exists := shenShaData[dayGan]; exists {
		for shenShaName, zhiList := range shenShaConfig {
			for _, column := range bazi {
				for _, zhi := range zhiList {
					if column.Zhi == zhi {
						result[shenShaName] = zhi
						break
					}
				}
			}
		}
	}

	return result
}

// CalculateForColumn 计算单个柱子的神煞
func (s *ShenShaService) CalculateForColumn(dayGan string, column models.BaziColumn) map[string]string {
	result := make(map[string]string)

	if shenShaConfig, exists := shenShaData[dayGan]; exists {
		for shenShaName, zhiList := range shenShaConfig {
			for _, zhi := range zhiList {
				if column.Zhi == zhi {
					result[shenShaName] = zhi
					break
				}
			}
		}
	}

	return result
}

// GetAllShenSha 获取所有神煞数据
func (s *ShenShaService) GetAllShenSha() map[string]map[string][]string {
	return shenShaData
}