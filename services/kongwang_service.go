package services

// KongWangService 空亡服务
type KongWangService struct{}

var (
	// 空亡对照表（以日柱查询）
	kongWangData = map[string][]string{
		"甲子": {"戌", "亥"}, "甲戌": {"申", "酉"}, "甲申": {"午", "未"}, "甲午": {"辰", "巳"},
		"甲辰": {"寅", "卯"}, "甲寅": {"子", "丑"}, "乙丑": {"戌", "亥"}, "乙亥": {"申", "酉"},
		"乙酉": {"午", "未"}, "乙未": {"辰", "巳"}, "乙巳": {"寅", "卯"}, "乙卯": {"子", "丑"},
		"丙寅": {"戌", "亥"}, "丙子": {"申", "酉"}, "丙戌": {"午", "未"}, "丙申": {"辰", "巳"},
		"丙午": {"寅", "卯"}, "丙辰": {"子", "丑"}, "丁卯": {"戌", "亥"}, "丁丑": {"申", "酉"},
		"丁亥": {"午", "未"}, "丁酉": {"辰", "巳"}, "丁未": {"寅", "卯"}, "丁巳": {"子", "丑"},
		"戊辰": {"戌", "亥"}, "戊寅": {"申", "酉"}, "戊子": {"午", "未"}, "戊戌": {"辰", "巳"},
		"戊申": {"寅", "卯"}, "戊午": {"子", "丑"}, "己巳": {"戌", "亥"}, "己卯": {"申", "酉"},
		"己丑": {"午", "未"}, "己亥": {"辰", "巳"}, "己酉": {"寅", "卯"}, "己未": {"子", "丑"},
		"庚午": {"戌", "亥"}, "庚辰": {"申", "酉"}, "庚寅": {"午", "未"}, "庚子": {"辰", "巳"},
		"庚戌": {"寅", "卯"}, "庚申": {"子", "丑"}, "辛未": {"戌", "亥"}, "辛巳": {"申", "酉"},
		"辛卯": {"午", "未"}, "辛丑": {"辰", "巳"}, "辛亥": {"寅", "卯"}, "辛酉": {"子", "丑"},
		"壬申": {"戌", "亥"}, "壬午": {"申", "酉"}, "壬辰": {"午", "未"}, "壬寅": {"辰", "巳"},
		"壬子": {"寅", "卯"}, "壬戌": {"子", "丑"}, "癸酉": {"戌", "亥"}, "癸未": {"申", "酉"},
		"癸巳": {"午", "未"}, "癸卯": {"辰", "巳"}, "癸丑": {"寅", "卯"}, "癸亥": {"子", "丑"},
	}
)

func NewKongWangService() *KongWangService {
	return &KongWangService{}
}

// Calculate 计算空亡
func (s *KongWangService) Calculate(dayGanZhi, targetZhi string) bool {
	if kongWangZhi, exists := kongWangData[dayGanZhi]; exists {
		for _, kwZhi := range kongWangZhi {
			if kwZhi == targetZhi {
				return true
			}
		}
	}
	return false
}

// GetKongWangZhi 获取指定日柱的空亡地支
func (s *KongWangService) GetKongWangZhi(dayGanZhi string) []string {
	if kongWangZhi, exists := kongWangData[dayGanZhi]; exists {
		result := make([]string, len(kongWangZhi))
		copy(result, kongWangZhi)
		return result
	}
	return []string{}
}

// GetAllKongWang 获取所有空亡数据
func (s *KongWangService) GetAllKongWang() map[string][]string {
	return kongWangData
}