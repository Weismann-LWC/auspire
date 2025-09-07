package services

// ZhuXingService 主星（十神）服务
type ZhuXingService struct{}

func NewZhuXingService() *ZhuXingService {
	return &ZhuXingService{}
}

// Calculate 计算主星（十神关系）
func (s *ZhuXingService) Calculate(dayGan, targetGan string) string {
	dayWuXing := tianGanWuXing[dayGan]
	targetWuXing := tianGanWuXing[targetGan]

	// 判断阴阳
	dayYinYang := s.getYinYang(dayGan)
	targetYinYang := s.getYinYang(targetGan)
	sameYinYang := dayYinYang == targetYinYang

	if dayGan == targetGan {
		return "比肩"
	}

	switch {
	case dayWuXing == targetWuXing:
		if sameYinYang {
			return "比肩"
		} else {
			return "劫财"
		}
	case s.isShengRelation(dayWuXing, targetWuXing):
		if sameYinYang {
			return "食神"
		} else {
			return "伤官"
		}
	case s.isKeRelation(dayWuXing, targetWuXing):
		if sameYinYang {
			return "偏财"
		} else {
			return "正财"
		}
	case s.isShengRelation(targetWuXing, dayWuXing):
		if sameYinYang {
			return "偏印"
		} else {
			return "正印"
		}
	case s.isKeRelation(targetWuXing, dayWuXing):
		if sameYinYang {
			return "七杀"
		} else {
			return "正官"
		}
	}

	return ""
}

// 获取天干阴阳
func (s *ZhuXingService) getYinYang(gan string) string {
	yangGan := []string{"甲", "丙", "戊", "庚", "壬"}
	for _, g := range yangGan {
		if g == gan {
			return "阳"
		}
	}
	return "阴"
}

// 判断是否为相生关系
func (s *ZhuXingService) isShengRelation(from, to string) bool {
	shengRelations := map[string]string{
		"木": "火", "火": "土", "土": "金", "金": "水", "水": "木",
	}
	return shengRelations[from] == to
}

// 判断是否为相克关系
func (s *ZhuXingService) isKeRelation(from, to string) bool {
	keRelations := map[string]string{
		"木": "土", "火": "金", "土": "水", "金": "木", "水": "火",
	}
	return keRelations[from] == to
}