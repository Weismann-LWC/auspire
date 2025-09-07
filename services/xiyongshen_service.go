package services

import (
	"auspire/models"
)

// XiYongShenService 喜用神计算服务
type XiYongShenService struct {
}

// NewXiYongShenService 创建新的喜用神服务实例
func NewXiYongShenService() *XiYongShenService {
	return &XiYongShenService{}
}

// Calculate 计算喜用神
func (s *XiYongShenService) Calculate(bazi []models.BaziColumn) *models.XiYongShenResult {
	result := &models.XiYongShenResult{
		Logic: []string{},
	}

	// 获取日主（日柱天干）
	riZhu := bazi[2].Gan
	result.RiZhu = riZhu

	// 计算五行得分
	wuXingScores := s.calculateWuXingScores(bazi)
	result.WuXingScores = wuXingScores

	// 判断日主强弱
	riZhuStrength := s.determineRiZhuStrength(riZhu, wuXingScores)
	result.RiZhuStrength = riZhuStrength

	// 确定喜用神
	xiYongShen, logic := s.determineXiYongShen(riZhu, riZhuStrength, wuXingScores)
	result.XiYongShen = xiYongShen
	result.Logic = logic

	return result
}

// calculateWuXingScores 计算五行得分
func (s *XiYongShenService) calculateWuXingScores(bazi []models.BaziColumn) map[string]int {
	scores := map[string]int{
		"木": 0,
		"火": 0,
		"土": 0,
		"金": 0,
		"水": 0,
	}

	// 统计天干地支中的五行数量
	for _, column := range bazi {
		// 天干五行
		scores[column.GanWuXing]++

		// 地支五行
		scores[column.ZhiWuXing]++

		// 考虑藏干的五行（简化处理，实际应该更复杂）
		// 这里我们暂时只考虑天干地支的五行
	}

	return scores
}

// determineRiZhuStrength 判断日主强弱
func (s *XiYongShenService) determineRiZhuStrength(riZhu string, scores map[string]int) string {
	// 获取日主五行
	riZhuWuXing := tianGanWuXing[riZhu]

	// 简化的判断逻辑：
	// 1. 如果日主五行在八字中出现次数>=2，则为偏强
	// 2. 如果日主五行在八字中出现次数<2，则为偏弱
	if scores[riZhuWuXing] >= 2 {
		return "偏强"
	}
	return "偏弱"
}

// determineXiYongShen 确定喜用神
func (s *XiYongShenService) determineXiYongShen(riZhu, riZhuStrength string, scores map[string]int) (string, []string) {
	var xiYongShen string
	logic := []string{}

	// 获取日主五行
	riZhuWuXing := tianGanWuXing[riZhu]

	logic = append(logic, "开始分析喜用神...")
	logic = append(logic, "1. 确定日主为: "+riZhu+", 五行属: "+riZhuWuXing)
	logic = append(logic, "2. 分析日主强弱: 日主"+riZhuStrength)

	// 根据日主强弱确定喜用神
	if riZhuStrength == "偏强" {
		// 日主偏强，需要克制或泄耗
		logic = append(logic, "3. 日主偏强，需要寻找能够克制或泄耗日主五行的元素")
		
		// 克制日主五行的为官杀（正官、七杀）
		keRiZhu := s.getKeWuXing(riZhuWuXing)
		logic = append(logic, "4. 能克制日主"+riZhuWuXing+"的五行是: "+keRiZhu+"（官杀）")
		
		// 泄耗日主五行的为财星（正财、偏财）和食伤（食神、伤官）
		shengKeRiZhu := s.getShengWuXing(riZhuWuXing)
		logic = append(logic, "5. 能泄耗日主"+riZhuWuXing+"的五行是: "+shengKeRiZhu+"（财星）")
		
		// 在克制和泄耗中选择更合适的作为喜用神
		// 简化处理：选择克制的五行作为喜用神
		xiYongShen = keRiZhu
		logic = append(logic, "6. 综合分析，确定喜用神为: "+xiYongShen)
	} else {
		// 日主偏弱，需要生扶
		logic = append(logic, "3. 日主偏弱，需要寻找能够生扶日主五行的元素")
		
		// 生扶日主五行的为印星（正印、偏印）
		shengRiZhu := s.getShengWuXing(riZhuWuXing)
		logic = append(logic, "4. 能生扶日主"+riZhuWuXing+"的五行是: "+shengRiZhu+"（印星）")
		
		// 比助日主五行的为比劫（比肩、劫财）
		biZhuRiZhu := riZhuWuXing
		logic = append(logic, "5. 能比助日主"+riZhuWuXing+"的五行是: "+biZhuRiZhu+"（比劫）")
		
		// 在生扶和比助中选择更合适的作为喜用神
		// 简化处理：选择生扶的五行作为喜用神
		xiYongShen = shengRiZhu
		logic = append(logic, "6. 综合分析，确定喜用神为: "+xiYongShen)
	}

	return xiYongShen, logic
}

// getKeWuXing 获取克制某五行的五行
func (s *XiYongShenService) getKeWuXing(wuxing string) string {
	keRelations := map[string]string{
		"木": "金",
		"火": "水",
		"土": "木",
		"金": "火",
		"水": "土",
	}
	return keRelations[wuxing]
}

// getShengWuXing 获取生某五行的五行
func (s *XiYongShenService) getShengWuXing(wuxing string) string {
	shengRelations := map[string]string{
		"木": "水",
		"火": "木",
		"土": "火",
		"金": "土",
		"水": "金",
	}
	return shengRelations[wuxing]
}