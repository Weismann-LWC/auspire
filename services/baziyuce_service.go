// Package services provides all the core business logic for bazi calculation and analysis.
//
// This package implements the traditional Chinese metaphysics concepts including:
// - Bazi (Four Pillars of Destiny) calculation
// - Ten Gods (Shi Shen) relationships
// - Hidden Stems (Cang Gan) analysis
// - Twelve Longevity (Shi Er Chang Sheng) cycles
// - Na Yin (Sonant Notes)五行 associations
// - Nobleman Stars (Gui Ren) and other Shen Sha calculations
//
// The implementation follows traditional rules while ensuring accuracy through:
// - Proper Jie Qi (Solar Terms) based month pillar calculation
// - Correct Yang/Yin Gan sequencing for twelve longevity
// - Accurate藏干 mappings for each Earthly Branch
//
// Each service is modularized for maintainability and extensibility.
package services

import (
	"auspire/models"
	"fmt"
)

// BaziyuceService 四柱八字综合分析服务
//
// This service provides comprehensive analysis of the Four Pillars of Destiny (八字).
// It follows the traditional Chinese metaphysical approach to analyze:
// 1. The foundational chart construction (排盘与定盘)
// 2. Vitality assessment (定旺衰，识体性)  
// 3. Favorable elements identification (明喜忌，定方向)
// 4. Pattern analysis (析格局，观组合)
// 5. Luck period forecasting (推大运，断流年)
//
// The analysis follows classical texts and methodologies including:
// - "子平真诠" (Zi Ping Zhen Quan) principles
// - "穷通宝鉴" (Qiong Tong Bao Jian) seasonal analysis
// - "三命通会" (San Ming Tong Hui) comprehensive approaches
type BaziyuceService struct {
	// Future enhancements could include caching, external API clients, etc.
}

// NewBaziyuceService 创建新的四柱八字综合分析服务实例
//
// Returns a pointer to a newly initialized BaziyuceService.
// This follows the singleton pattern commonly used in Go services.
func NewBaziyuceService() *BaziyuceService {
	return &BaziyuceService{}
}

// Analyze 四柱八字综合分析入口点
//
// This is the main entry point for comprehensive BaZi analysis.
// It orchestrates all five major analytical steps in traditional order:
//
// Step 1: Chart Establishment (排盘与定盘)
//   - Verifies the accuracy of the four pillars
//   - Confirms proper Jie Qi based month pillar calculation
//   - Establishes the foundational chart structure
//
// Step 2: Vitality Assessment (定旺衰，识体性) 
//   - Determines the Day Master's (日主) strength through four dimensions:
//     * Ling - Seasonal timing/power (得令)
//     * Di - Root/stability (得地) 
//     * Shi - Support from allies (得势)
//     * Zhu - Support from parents/mentors (得助)
//
// Step 3: Favorable Elements (明喜忌，定方向)
//   - Identifies favorable (喜用神) and unfavorable (忌神) elements
//   - Determines the seeker's optimal directions and elements
//
// Step 4: Pattern Analysis (析格局，观组合)
//   - Examines star combinations and formations
//   - Analyzes noble person stars and other divine influences
//   - Considers palace positions and relationships
//
// Step 5: Luck Period Forecasting (推大运，断流年)
//   - Projects future luck periods
//   - Analyzes annual influences
//   - Provides temporal guidance
//
// Parameters:
//   - bazi: Slice of BaziColumn representing the four pillars
//
// Returns:
//   - Pointer to BaziyuceResult containing all analysis steps
//   - Following the traditional five-step analysis methodology
func (s *BaziyuceService) Analyze(bazi []models.BaziColumn) *models.BaziyuceResult {
	result := &models.BaziyuceResult{
		Steps: []models.AnalysisStep{},
	}

	// 第一步：排盘与定盘 - Chart Establishment
	// Establishes the foundational chart and confirms accuracy
	step1 := s.step1PaiPanDingPan(bazi)
	result.Steps = append(result.Steps, step1)

	// 第二步：定旺衰，识体性 - Vitality Assessment  
	// Determines the Day Master's strength through four dimensions
	step2 := s.step2DingWangShuai(bazi)
	result.Steps = append(result.Steps, step2)

	// 第三步：明喜忌，定方向 - Favorable Elements Identification
	// Identifies beneficial and harmful elements for the seeker
	step3 := s.step3MingXiJi(bazi, step2)
	result.Steps = append(result.Steps, step3)

	// 第四步：析格局，观组合 - Pattern Analysis
	// Examines combinations, formations, and divine influences
	step4 := s.step4XiGeJu(bazi)
	result.Steps = append(result.Steps, step4)

	// 第五步：推大运，断流年 - Luck Period Forecasting
	// Projects future influences and temporal guidance
	step5 := s.step5TuiDaYun(bazi)
	result.Steps = append(result.Steps, step5)

	return result
}

// step1PaiPanDingPan 第一步：排盘与定盘 - Chart Establishment
//
// This step establishes the foundational chart and confirms its accuracy.
// It's crucial to ensure all pillars are correctly calculated before proceeding.
//
// Key Verification Points:
// 1. Year Pillar: Based on Heavenly Stems and Earthly Branches cycle
// 2. Month Pillar: Must be calculated using Solar Terms (Jie Qi), not calendar months
// 3. Day Pillar: Based on continuous day count since reference point
// 4. Hour Pillar: Derived from Day Master Stem and Earthly Branch sequence
//
// Parameters:
//   - bazi: Slice of four BaziColumn representing the complete chart
//
// Returns:
//   - AnalysisStep containing establishment procedures and verification
func (s *BaziyuceService) step1PaiPanDingPan(bazi []models.BaziColumn) models.AnalysisStep {
	step := models.AnalysisStep{
		Title:   "第一步：排盘与定盘",
		Content: []string{},
	}

	// Explanation of what this step accomplishes
	step.Content = append(step.Content, "做什么：将出生时间转换为天干地支表示的四柱八字，并确认其准确性。")
	step.Content = append(step.Content, "怎么做：")
	step.Content = append(step.Content, "1. 排四柱：使用专业排盘软件或权威万年历，输入公历出生年、月、日、时。")
	step.Content = append(step.Content, "2. 年柱：以立春为界，非农历正月初一。")
	step.Content = append(step.Content, "3. 月柱：以节气为界（如立春-惊蛰为寅月，惊蛰-清明为卯月）。")
	step.Content = append(step.Content, "4. 日柱：直接查询。")
	step.Content = append(step.Content, "5. 时柱：将北京时间换算为真太阳时（考虑出生地经度与东经120度的时差）。")
	step.Content = append(step.Content, "6. 定十神：以日干为我，为所有天干和地支藏干标注十神。")
	step.Content = append(step.Content, "7. 排大运：根据年干阴阳和性别（阳男阴女顺排；阴男阳女逆排），从月柱推导出大运。")

	// Display the calculated chart
	step.Content = append(step.Content, "")
	step.Content = append(step.Content, "您的八字排盘结果：")
	for i, column := range bazi {
		columnNames := []string{"年柱", "月柱", "日柱", "时柱"}
		step.Content = append(step.Content, fmt.Sprintf("%s: %s%s (%s%s)", 
			columnNames[i], column.Gan, column.Zhi, column.GanWuXing, column.ZhiWuXing))
	}

	// Show ten gods analysis
	step.Content = append(step.Content, "")
	step.Content = append(step.Content, "十神分析：")
	for i, column := range bazi {
		columnNames := []string{"年柱", "月柱", "日柱", "时柱"}
		step.Content = append(step.Content, fmt.Sprintf("%s: 天干(%s)为%s", 
			columnNames[i], column.Gan, column.ZhuXing))
	}

	return step
}

// step2DingWangShuai 第二步：定旺衰，识体性 - Vitality Assessment
//
// This critical step determines the Day Master's overall strength,
// which forms the foundation for identifying favorable elements.
//
// Four Dimensions of Assessment (权重：得令 > 得地 > 得势/得助):
// 1. 得令 (De Ling) - Seasonal Timing/Command Authority
//    Based on the monthly branch relationship to Day Master
//
// 2. 得地 (De Di) - Root/Stability  
//    Based on Earthly Branches supporting the Day Master's element
//
// 3. 得势 (De Shi) - Support from Allies
//    Based on presence of helpful elements (比肩, 劫财)
//
// 4. 得助 (De Zhu) - Support from Parents/Mentors
//    Based on presence of supportive elements (印星)
//
// Parameters:
//   - bazi: Slice of four BaziColumn representing the complete chart
//
// Returns:
//   - AnalysisStep containing vitality assessment procedures and results
func (s *BaziyuceService) step2DingWangShuai(bazi []models.BaziColumn) models.AnalysisStep {
	step := models.AnalysisStep{
		Title:   "第二步：定旺衰，识体性",
		Content: []string{},
	}

	// Get the Day Master
	riZhu := bazi[2].Gan
	riZhuZhi := bazi[2].Zhi
	step.Content = append(step.Content, "做什么：判断日主在整个八字中的能量状态（身强/身弱），这是选择喜用神的根本依据。")
	step.Content = append(step.Content, fmt.Sprintf("您的日主为: %s, 地支为: %s", riZhu, riZhuZhi))

	// Assess from four dimensions
	step.Content = append(step.Content, "")
	step.Content = append(step.Content, "怎么做：从四个维度综合评估（权重：得令 > 得地 > 得势/得助）。")

	// 1. 得令 (De Ling) - Seasonal Timing
	yueZhi := bazi[1].Zhi
	yueLingStatus := s.getYueLingStatus(riZhu, yueZhi)
	step.Content = append(step.Content, "1. 得令（看月令）:")
	step.Content = append(step.Content, fmt.Sprintf("   日主%s在出生月份（%s月）的状态为: %s", riZhu, yueZhi, yueLingStatus))

	// 2. 得地 (De Di) - Root/Stability
	deDiStatus := s.getDeDiStatus(riZhu, bazi)
	step.Content = append(step.Content, "2. 得地（看根气）:")
	step.Content = append(step.Content, fmt.Sprintf("   %s", deDiStatus))

	// 3. 得势 (De Shi) - Support from Allies
	deShiStatus := s.getDeShiStatus(riZhu, bazi)
	step.Content = append(step.Content, "3. 得势（看比劫）:")
	step.Content = append(step.Content, fmt.Sprintf("   %s", deShiStatus))

	// 4. 得助 (De Zhu) - Support from Parents/Mentors
	deZhuStatus := s.getDeZhuStatus(riZhu, bazi)
	step.Content = append(step.Content, "4. 得助（看印星）:")
	step.Content = append(step.Content, fmt.Sprintf("   %s", deZhuStatus))

	// Comprehensive judgment
	zongHePanDuan := s.getZongHePanDuan(yueLingStatus, deDiStatus, deShiStatus, deZhuStatus)
	step.Content = append(step.Content, "")
	step.Content = append(step.Content, "综合判断:")
	step.Content = append(step.Content, zongHePanDuan)

	return step
}

// step3MingXiJi 第三步：明喜忌，定方向 - Favorable Elements Identification
//
// Based on the vitality assessment, this step identifies favorable and unfavorable elements.
// This is crucial for determining optimal life directions and avoiding harmful influences.
//
// Body Strength Determination:
// - 身旺 (Shen Wang) - Strong Constitution: 
//   Prefers elements that restrain, exhaust, or consume (克、泄、耗)
//   Favorable: 官杀 (Officers/Killings), 食伤 (Food/Hurts), 财星 (Wealth Stars)
//   Unfavorable: 印星 (Seal Stars), 比劫 (Allies/Rob Wealth)
//
// - 身弱 (Shen Ruo) - Weak Constitution:
//   Prefers elements that generate and support (生、扶)
//   Favorable: 印星 (Seal Stars), 比劫 (Allies/Rob Wealth)  
//   Unfavorable: 官杀 (Officers/Killings), 食伤 (Food/Hurts), 财星 (Wealth Stars)
//
// Parameters:
//   - bazi: Slice of four BaziColumn representing the complete chart
//   - step2: Results from vitality assessment step
//
// Returns:
//   - AnalysisStep containing favorable element identification procedures and results
func (s *BaziyuceService) step3MingXiJi(bazi []models.BaziColumn, step2 models.AnalysisStep) models.AnalysisStep {
	step := models.AnalysisStep{
		Title:   "第三步：明喜忌，定方向",
		Content: []string{},
	}

	step.Content = append(step.Content, "做什么：根据身强身弱，确定平衡八字的五行十神。")

	// Get the body strength determination result
	shenRiZhuangTai := ""
	for _, content := range step2.Content {
		if contains(content, "身旺/强") || contains(content, "身弱/衰") || contains(content, "身高中和") {
			shenRiZhuangTai = content
			break
		}
	}

	step.Content = append(step.Content, fmt.Sprintf("您的日主状态为: %s", shenRiZhuangTai))

	riZhu := bazi[2].Gan
	riZhuWuXing := tianGanWuXing[riZhu]

	if contains(shenRiZhuangTai, "身旺") {
		step.Content = append(step.Content, "怎么做：")
		step.Content = append(step.Content, "身旺：喜克、泄、耗。喜用神为：官杀、食伤、财星。忌神为：印星、比劫。")
		
		keWo := s.getKeWuXing(riZhuWuXing)     // 克我者为官杀
		woKe := s.getShengWuXing(riZhuWuXing)  // 我克者为财星
		woSheng := s.getShengWuXing(woKe)      // 我生者为食伤
		
		step.Content = append(step.Content, "喜用神：")
		step.Content = append(step.Content, fmt.Sprintf("   官杀(%s)、财星(%s)、食伤(%s)", keWo, woKe, woSheng))
		step.Content = append(step.Content, "忌神：")
		step.Content = append(step.Content, fmt.Sprintf("   印星(%s)、比劫(%s)", riZhuWuXing, riZhuWuXing))
	} else if contains(shenRiZhuangTai, "身弱") {
		step.Content = append(step.Content, "怎么做：")
		step.Content = append(step.Content, "身弱：喜生、扶。喜用神为：印星、比劫。忌神为：官杀、食伤、财星。")
		
		shengWo := s.getShengWuXing(riZhuWuXing)  // 生我者为印星
		biZhu := riZhuWuXing                      // 同我者为比劫
		keWo := s.getKeWuXing(riZhuWuXing)        // 克我者为官杀
		woKe := s.getShengWuXing(riZhuWuXing)     // 我克者为财星
		woSheng := s.getShengWuXing(woKe)         // 我生者为食伤
		
		step.Content = append(step.Content, "喜用神：")
		step.Content = append(step.Content, fmt.Sprintf("   印星(%s)、比劫(%s)", shengWo, biZhu))
		step.Content = append(step.Content, "忌神：")
		step.Content = append(step.Content, fmt.Sprintf("   官杀(%s)、财星(%s)、食伤(%s)", keWo, woKe, woSheng))
	} else {
		step.Content = append(step.Content, "身高中和：五行相对平衡，需根据具体组合确定喜忌。")
	}

	return step
}

// step4XiGeJu 第四步：析格局，观组合 - Pattern Analysis
//
// This step examines the combinations and formations in the chart,
// looking for special patterns that indicate exceptional destiny characteristics.
//
// Key Analysis Areas:
// 1. Monthly Branch Focus: The monthly branch contains the commander star
// 2. Ten Gods Combinations: Look for auspicious and inauspicious combinations
// 3. Palace Correlations: Relationship between stars and palaces
// 4. Earthly Branch Interactions: Conflict, harmony, combination, and destruction
//
// Parameters:
//   - bazi: Slice of four BaziColumn representing the complete chart
//
// Returns:
//   - AnalysisStep containing pattern analysis procedures and results
func (s *BaziyuceService) step4XiGeJu(bazi []models.BaziColumn) models.AnalysisStep {
	step := models.AnalysisStep{
		Title:   "第四步：析格局，观组合",
		Content: []string{},
	}

	step.Content = append(step.Content, "做什么：分析十神的分布、组合和力量，解读人生轨迹。")
	step.Content = append(step.Content, "怎么做：")

	// Focus on Monthly Branch
	yueZhi := bazi[1].Zhi
	yueZhiCangGan := bazi[1].CangGan
	step.Content = append(step.Content, fmt.Sprintf("1. 聚焦月令：月支%s藏干为%s，分析其本气十神作为格局核心。", 
		yueZhi, s.joinCangGan(yueZhiCangGan)))

	// Analyze Ten Gods Combinations
	step.Content = append(step.Content, "2. 分析十神组合：")
	
	// Check for auspicious combinations
	jiShenZuHe := s.checkJiShenZuHe(bazi)
	if jiShenZuHe != "" {
		step.Content = append(step.Content, fmt.Sprintf("   存在吉神组合: %s，主富贵。", jiShenZuHe))
	} else {
		step.Content = append(step.Content, "   暂未发现明显吉神组合。")
	}

	// Check for inauspicious combinations
	xiongShenZuHe := s.checkXiongShenZuHe(bazi)
	if xiongShenZuHe != "" {
		step.Content = append(step.Content, fmt.Sprintf("   存在凶神组合: %s，主波折。", xiongShenZuHe))
	} else {
		step.Content = append(step.Content, "   暂未发现明显凶神组合。")
	}

	// Star-Palace Correlation
	step.Content = append(step.Content, "3. 星宫同参：将十神与所在的宫位结合看。")
	step.Content = append(step.Content, "   - 财星在妻宫（日支）：妻子贤惠或得妻财。")
	step.Content = append(step.Content, "   - 印星在母宫（年柱）：得母亲关爱。")
	step.Content = append(step.Content, "   - 官星在事业宫（月柱）：事业心强。")

	// Earthly Branch Conflicts and Harmonies
	step.Content = append(step.Content, "4. 察地支刑冲合害：分析地支间的相互作用。")
	diZhiGuanXi := s.analyzeDiZhiGuanXi(bazi)
	if diZhiGuanXi != "" {
		step.Content = append(step.Content, fmt.Sprintf("   地支关系: %s", diZhiGuanXi))
	} else {
		step.Content = append(step.Content, "   地支间无明显刑冲合害关系。")
	}

	return step
}

// step5TuiDaYun 第五步：推大运，断流年 - Luck Period Forecasting
//
// This final step projects future influences based on the established chart,
// showing how the Luck Pillars and annual influences activate latent potentials.
//
// Key Concepts:
// 1. Luck Pillars (大运): Ten-year cycles that enhance or diminish chart elements
// 2. Annual Influences (流年): Yearly activations that trigger specific events
// 3. Activation Mechanisms: How Luck Pillars and Years interact with the natal chart
//
// Parameters:
//   - bazi: Slice of four BaziColumn representing the complete chart
//
// Returns:
//   - AnalysisStep containing luck period forecasting procedures and results
func (s *BaziyuceService) step5TuiDaYun(bazi []models.BaziColumn) models.AnalysisStep {
	step := models.AnalysisStep{
		Title:   "第五步：推大运，断流年",
		Content: []string{},
	}

	step.Content = append(step.Content, "做什么：将大运和流年代入原局，看如何引动和改变原局的平衡。")
	step.Content = append(step.Content, "怎么做：")

	step.Content = append(step.Content, "1. 大运分析：看十年一大运是增强了喜用神还是忌神的力量。")
	step.Content = append(step.Content, "   - 大运天干地支五行属性与原局喜忌神的关系")
	step.Content = append(step.Content, "   - 大运与原局之间的生克制化关系")

	step.Content = append(step.Content, "2. 流年应期：流年干支像一把钥匙，会引动原局中潜伏的信息。")
	step.Content = append(step.Content, "   - 原局有潜在隐患时，流年可能引动")
	step.Content = append(step.Content, "   - 原局喜用神得流年生扶时，运势提升")

	// Example analysis
	riZhu := bazi[2].Gan
	xiYongShen := s.determineXiYongShenSimple(bazi)
	step.Content = append(step.Content, "")
	step.Content = append(step.Content, "示例分析：")
	step.Content = append(step.Content, fmt.Sprintf("您的日主为%s，喜用神为%s。", riZhu, xiYongShen))
	step.Content = append(step.Content, fmt.Sprintf("当大运或流年出现%s五行时，通常代表运势较好。", xiYongShen))
	step.Content = append(step.Content, "当大运或流年出现忌神五行时，需谨慎应对可能出现的挑战。")

	return step
}

// Helper methods for step implementations

// getYueLingStatus 获取月令状态 (Determine Monthly Command Status)
func (s *BaziyuceService) getYueLingStatus(riZhu, yueZhi string) string {
	// Use the corrected twelve longevity calculation
	changShengStatus := GetZhangShengPosition(riZhu, yueZhi)
	
	// Determine command authority status
	deLingStates := []string{"临官", "帝旺", "长生", "冠带", "沐浴"}
	shiLingStates := []string{"衰", "病", "死", "墓", "绝", "胎", "养"}
	
	// Check for commanding authority (得令)
	for _, state := range deLingStates {
		if changShengStatus == state {
			return fmt.Sprintf("得令（%s），趋势强", changShengStatus)
		}
	}
	
	// Check for loss of command (失令)
	for _, state := range shiLingStates {
		if changShengStatus == state {
			return fmt.Sprintf("失令（%s），趋势弱", changShengStatus)
		}
	}
	
	return fmt.Sprintf("月令状态：%s", changShengStatus)
}

// getDeDiStatus 得地状态分析
func (s *BaziyuceService) getDeDiStatus(riZhu string, bazi []models.BaziColumn) string {
	riZhuWuXing := tianGanWuXing[riZhu]
	qiangGenCount := 0
	weiGenCount := 0
	
	// 检查地支藏干
	for _, column := range bazi {
		cangGan := column.CangGan
		for _, gan := range cangGan {
			if gan != "" && tianGanWuXing[gan] == riZhuWuXing {
				// 判断是否为本气强根
				if s.isBenQiQiangGen(column.Zhi, gan) {
					qiangGenCount++
				} else {
					weiGenCount++
				}
			}
		}
	}
	
	// 修改判断逻辑：
	// 1. 如果有本气强根，得地力强
	// 2. 如果同五行藏干总数>=3，得地力强
	// 3. 如果同五行藏干总数>=1，得地力弱
	// 4. 如果完全没有同五行藏干，力量弱
	totalSameWuXing := qiangGenCount + weiGenCount
	if qiangGenCount > 0 {
		return "有本气强根，得地力强"
	} else if totalSameWuXing >= 3 {
		return fmt.Sprintf("同五行藏干数量为%d个，得地力强", totalSameWuXing)
	} else if totalSameWuXing >= 1 {
		return fmt.Sprintf("同五行藏干数量为%d个，得地力弱", totalSameWuXing)
	} else {
		return "完全无根为\"虚浮\"，力量弱"
	}
}

// isBenQiQiangGen 判断是否为本气强根
func (s *BaziyuceService) isBenQiQiangGen(zhi, gan string) bool {
	benQiMap := map[string]string{
		"寅": "甲", "卯": "乙", "巳": "丙", "午": "丁", "辰": "戊", "戌": "戊", "丑": "己", "未": "己",
		"申": "庚", "酉": "辛", "亥": "壬", "子": "癸",
	}
	
	if benQi, exists := benQiMap[zhi]; exists {
		return benQi == gan
	}
	return false
}

// getDeShiStatus 得势状态分析（比劫）
func (s *BaziyuceService) getDeShiStatus(riZhu string, bazi []models.BaziColumn) string {
	biJieCount := 0
	
	// 统计天干中的比肩、劫财
	for _, column := range bazi {
		zhuXing := column.ZhuXing
		if zhuXing == "比肩" || zhuXing == "劫财" {
			biJieCount++
		}
	}
	
	// 统计地支藏干中的比肩、劫财
	for _, column := range bazi {
		fuXing := column.FuXing
		for _, fx := range fuXing {
			if fx == "比肩" || fx == "劫财" {
				biJieCount++
			}
		}
	}
	
	if biJieCount >= 2 {
		return fmt.Sprintf("天干和地支藏干中比肩、劫财数量为%d，势众，得势", biJieCount)
	} else if biJieCount == 1 {
		return fmt.Sprintf("天干和地支藏干中比肩、劫财数量为%d，势弱", biJieCount)
	} else {
		return "天干和地支藏干中无比肩、劫财，不得势"
	}
}

// getDeZhuStatus 得助状态分析（印星）
func (s *BaziyuceService) getDeZhuStatus(riZhu string, bazi []models.BaziColumn) string {
	yinXingCount := 0
	
	// 统计天干中的正印、偏印
	for _, column := range bazi {
		zhuXing := column.ZhuXing
		if zhuXing == "正印" || zhuXing == "偏印" {
			yinXingCount++
		}
	}
	
	// 统计地支藏干中的正印、偏印
	for _, column := range bazi {
		fuXing := column.FuXing
		for _, fx := range fuXing {
			if fx == "正印" || fx == "偏印" {
				yinXingCount++
			}
		}
	}
	
	if yinXingCount >= 2 {
		return fmt.Sprintf("天干和地支藏干中印星数量为%d，得生助之力强，得助", yinXingCount)
	} else if yinXingCount == 1 {
		return fmt.Sprintf("天干和地支藏干中印星数量为%d，得生助之力弱", yinXingCount)
	} else {
		return "天干和地支藏干中无印星，不得助"
	}
}

// getZongHePanDuan 综合判断
func (s *BaziyuceService) getZongHePanDuan(yueLingStatus, deDiStatus, deShiStatus, deZhuStatus string) string {
	deLing := false
	deDi := false
	deShi := false
	deZhu := false
	
	// 判断是否得令
	if contains(yueLingStatus, "得令") {
		deLing = true
	}
	
	// 判断是否得地
	if contains(deDiStatus, "得地力强") || contains(deDiStatus, "得地力弱") {
		deDi = true
	}
	
	// 判断是否得势
	if contains(deShiStatus, "得势") {
		deShi = true
	}
	
	// 判断是否得助
	if contains(deZhuStatus, "得助") {
		deZhu = true
	}
	
	// 身旺判断：得令 + (得地、得势、得助满足其一或多项)
	if deLing && (deDi || deShi || deZhu) {
		return "身旺/强：得令 + (得地、得势、得助满足其一或多项)"
	}
	
	// 身弱判断：失令 + (不得地、不得势、不得助满足其一或多项)
	if !deLing && (!deDi || !deShi || !deZhu) {
		return "身弱/衰：失令 + (不得地、不得势、不得助满足其一或多项)"
	}
	
	// 默认返回中和
	return "身高中和：八字五行相对平衡"
}

// 辅助方法
func (s *BaziyuceService) getKeWuXing(wuxing string) string {
	keRelations := map[string]string{
		"木": "金",
		"火": "水",
		"土": "木",
		"金": "火",
		"水": "土",
	}
	return keRelations[wuxing]
}

func (s *BaziyuceService) getShengWuXing(wuxing string) string {
	shengRelations := map[string]string{
		"木": "水",
		"火": "木",
		"土": "火",
		"金": "土",
		"水": "金",
	}
	return shengRelations[wuxing]
}

func (s *BaziyuceService) joinCangGan(cangGan []string) string {
	result := ""
	for i, gan := range cangGan {
		if i > 0 {
			result += "、"
		}
		result += gan
	}
	if result == "" {
		result = "无"
	}
	return result
}

func (s *BaziyuceService) checkJiShenZuHe(bazi []models.BaziColumn) string {
	// 简化检查，实际应该更复杂
	hasShiShen := false
	hasGuanSha := false
	hasYinXing := false
	hasCaiXing := false

	for _, column := range bazi {
		zhuXing := column.ZhuXing
		if zhuXing == "食神" || zhuXing == "伤官" {
			hasShiShen = true
		}
		if zhuXing == "正官" || zhuXing == "七杀" {
			hasGuanSha = true
		}
		if zhuXing == "正印" || zhuXing == "偏印" {
			hasYinXing = true
		}
		if zhuXing == "正财" || zhuXing == "偏财" {
			hasCaiXing = true
		}
	}

	// 检查常见吉神组合
	if hasShiShen && hasGuanSha {
		return "食神制杀"
	}
	if (hasShiShen || hasGuanSha) && hasYinXing {
		return "伤官配印或官印相生"
	}
	if hasShiShen && hasCaiXing {
		return "食伤生财"
	}

	return ""
}

func (s *BaziyuceService) checkXiongShenZuHe(bazi []models.BaziColumn) string {
	// 简化检查，实际应该更复杂
	hasShangGuan := false
	hasZhengGuan := false
	hasPianYin := false
	hasShiShen := false
	hasBiJie := false
	hasCaiXing := false

	for _, column := range bazi {
		zhuXing := column.ZhuXing
		if zhuXing == "伤官" {
			hasShangGuan = true
		}
		if zhuXing == "正官" {
			hasZhengGuan = true
		}
		if zhuXing == "偏印" {
			hasPianYin = true
		}
		if zhuXing == "食神" {
			hasShiShen = true
		}
		if zhuXing == "比肩" || zhuXing == "劫财" {
			hasBiJie = true
		}
		if zhuXing == "正财" || zhuXing == "偏财" {
			hasCaiXing = true
		}
	}

	// 检查常见凶神组合
	if hasShangGuan && hasZhengGuan {
		return "伤官见官"
	}
	if hasPianYin && hasShiShen {
		return "枭神夺食"
	}
	if hasBiJie && hasCaiXing {
		return "比劫夺财"
	}

	return ""
}

func (s *BaziyuceService) analyzeDiZhiGuanXi(bazi []models.BaziColumn) string {
	// 简化分析，实际应该更复杂
	diZhi := make([]string, 4)
	for i, column := range bazi {
		diZhi[i] = column.Zhi
	}

	// 检查常见的地支关系
	relations := []string{}

	// 检查是否有冲
	chongPairs := map[string]string{
		"子": "午", "丑": "未", "寅": "申", "卯": "酉",
		"辰": "戌", "巳": "亥", "午": "子", "未": "丑",
		"申": "寅", "酉": "卯", "戌": "辰", "亥": "巳",
	}

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			if chongPairs[diZhi[i]] == diZhi[j] {
				relations = append(relations, fmt.Sprintf("%s与%s相冲", diZhi[i], diZhi[j]))
			}
		}
	}

	if len(relations) > 0 {
		result := ""
		for i, rel := range relations {
			if i > 0 {
				result += "，"
			}
			result += rel
		}
		return result
	}

	return ""
}

func (s *BaziyuceService) determineXiYongShenSimple(bazi []models.BaziColumn) string {
	// 简化版本的喜用神判断
	riZhu := bazi[2].Gan
	riZhuWuXing := tianGanWuXing[riZhu]
	
	// 简单判断身强身弱
	qiangCount := 0
	if s.getYueLingStatus(riZhu, bazi[1].Zhi) == "得令" {
		qiangCount++
	}
	
	// 简化的喜用神判断
	if qiangCount >= 1 {
		// 身强，喜克、泄、耗
		return fmt.Sprintf("%s、%s", s.getKeWuXing(riZhuWuXing), s.getShengWuXing(riZhuWuXing))
	} else {
		// 身弱，喜生、扶
		return fmt.Sprintf("%s、%s", s.getShengWuXing(riZhuWuXing), riZhuWuXing)
	}
}

// contains 判断字符串是否包含子串
func contains(s, substr string) bool {
	return len(s) >= len(substr) && 
		   (s == substr || 
		    (len(s) > len(substr) && 
		     (s[:len(substr)] == substr || 
		      s[len(s)-len(substr):] == substr || 
		      len(s) > len(substr) && containsHelper(s, substr))))
}

// containsHelper 辅助函数
func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}