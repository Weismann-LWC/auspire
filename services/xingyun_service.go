package services

// XingYunService 星运（十二长生）服务
type XingYunService struct{}

var (
	// 十二长生：五行在十二地支中的长生状态
	xingYunData = map[string]map[string]string{
		"木": {"亥": "长生", "子": "沐浴", "丑": "冠带", "寅": "临官", "卯": "帝旺", "辰": "衰", "巳": "病", "午": "死", "未": "墓", "申": "绝", "酉": "胎", "戌": "养"},
		"火": {"寅": "长生", "卯": "沐浴", "辰": "冠带", "巳": "临官", "午": "帝旺", "未": "衰", "申": "病", "酉": "死", "戌": "墓", "亥": "绝", "子": "胎", "丑": "养"},
		"土": {"寅": "长生", "卯": "沐浴", "辰": "冠带", "巳": "临官", "午": "帝旺", "未": "衰", "申": "病", "酉": "死", "戌": "墓", "亥": "绝", "子": "胎", "丑": "养"},
		"金": {"巳": "长生", "午": "沐浴", "未": "冠带", "申": "临官", "酉": "帝旺", "戌": "衰", "亥": "病", "子": "死", "丑": "墓", "寅": "绝", "卯": "胎", "辰": "养"},
		"水": {"申": "长生", "酉": "沐浴", "戌": "冠带", "亥": "临官", "子": "帝旺", "丑": "衰", "寅": "病", "卯": "死", "辰": "墓", "巳": "绝", "午": "胎", "未": "养"},
	}
)

func NewXingYunService() *XingYunService {
	return &XingYunService{}
}

// Calculate 计算星运（十二运程）
func (s *XingYunService) Calculate(dayGanWuXing, zhi string) string {
	if changShengMap, exists := xingYunData[dayGanWuXing]; exists {
		if changSheng, exists := changShengMap[zhi]; exists {
			return changSheng
		}
	}
	return ""
}

// GetAllXingYun 获取所有星运数据
func (s *XingYunService) GetAllXingYun() map[string]map[string]string {
	return xingYunData
}