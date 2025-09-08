package services

import (
	"fmt"
	"time"
	"auspire/models"
	"auspire/services/solarterm"
)

var (
	tianGan = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	diZhi   = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	
	// 天干对应的五行
	tianGanWuXing = map[string]string{
		"甲": "木", "乙": "木", "丙": "火", "丁": "火", "戊": "土",
		"己": "土", "庚": "金", "辛": "金", "壬": "水", "癸": "水",
	}
	
	// 地支对应的五行
	diZhiWuXing = map[string]string{
		"子": "水", "丑": "土", "寅": "木", "卯": "木", "辰": "土", "巳": "火",
		"午": "火", "未": "土", "申": "金", "酉": "金", "戌": "土", "亥": "水",
	}
	
	// 十二长生：五行在十二地支中的长生状态
	shiErChangSheng = map[string]map[string]string{
		"木": {"亥": "长生", "子": "沐浴", "丑": "冠带", "寅": "临官", "卯": "帝旺", "辰": "衰", "巳": "病", "午": "死", "未": "墓", "申": "绝", "酉": "胎", "戌": "养"},
		"火": {"寅": "长生", "卯": "沐浴", "辰": "冠带", "巳": "临官", "午": "帝旺", "未": "衰", "申": "病", "酉": "死", "戌": "墓", "亥": "绝", "子": "胎", "丑": "养"},
		"土": {"寅": "长生", "卯": "沐浴", "辰": "冠带", "巳": "临官", "午": "帝旺", "未": "衰", "申": "病", "酉": "死", "戌": "墓", "亥": "绝", "子": "胎", "丑": "养"},
		"金": {"巳": "长生", "午": "沐浴", "未": "冠带", "申": "临官", "酉": "帝旺", "戌": "衰", "亥": "病", "子": "死", "丑": "墓", "寅": "绝", "卯": "胎", "辰": "养"},
		"水": {"申": "长生", "酉": "沐浴", "戌": "冠带", "亥": "临官", "子": "帝旺", "丑": "衰", "寅": "病", "卯": "死", "辰": "墓", "巳": "绝", "午": "胎", "未": "养"},
	}

	// 节气与地支的对应关系（月支由节气决定）
	// 立春(2/4) 开始为寅月，惊蛰(3/6) 开始为卯月，清明(4/5) 开始为辰月
	// 立夏(5/6) 开始为巳月，芒种(6/6) 开始为午月，小暑(7/7) 开始为未月
	// 立秋(8/8) 开始为申月，白露(9/8) 开始为酉月，寒露(10/8) 开始为戌月
	// 立冬(11/8) 开始为亥月，大雪(12/7) 开始为子月，小寒(1/6) 开始为丑月
	solarTermToMonthZhi = map[string]string{
		"立春": "寅", "雨水": "寅",
		"惊蛰": "卯", "春分": "卯",
		"清明": "辰", "谷雨": "辰",
		"立夏": "巳", "小满": "巳",
		"芒种": "午", "夏至": "午",
		"小暑": "未", "大暑": "未",
		"立秋": "申", "处暑": "申",
		"白露": "酉", "秋分": "酉",
		"寒露": "戌", "霜降": "戌",
		"立冬": "亥", "小雪": "亥",
		"大雪": "子", "冬至": "子",
		"小寒": "丑", "大寒": "丑",
	}

)

type BaziService struct{
	zhuXingService  *ZhuXingService
	cangGanService  *CangGanService
	fuXingService   *FuXingService
	naYinService    *NaYinService
	xingYunService  *XingYunService
	ziZuoService    *ZiZuoService
	kongWangService *KongWangService
	shenShaService  *ShenShaService
}

func NewBaziService() *BaziService {
	return &BaziService{
		zhuXingService:  NewZhuXingService(),
		cangGanService:  NewCangGanService(),
		fuXingService:   NewFuXingService(),
		naYinService:    NewNaYinService(),
		xingYunService:  NewXingYunService(),
		ziZuoService:    NewZiZuoService(),
		kongWangService: NewKongWangService(),
		shenShaService:  NewShenShaService(),
	}
}

func (s *BaziService) CalculateBazi(req models.BaziRequest) (*models.BaziResponse, error) {
	bazi, err := s.calculateBaziColumns(req.BirthDate, req.BirthTime)
	if err != nil {
		return &models.BaziResponse{
			Name:  req.Name,
			Error: err.Error(),
		}, err
	}

	// 计算十二长生图
	shiErChangShengResult := s.calculateShiErChangSheng(bazi)

	// 增强计算：添加新功能
	bazi = s.enhanceBaziColumns(bazi)

	return &models.BaziResponse{
		Name:            req.Name,
		Bazi:            bazi,
		ShiErChangSheng: shiErChangShengResult,
	}, nil
}

func (s *BaziService) calculateBaziColumns(birthDate, birthTime string) ([]models.BaziColumn, error) {
	parsedDate, err := time.Parse("2006-01-02", birthDate)
	if err != nil {
		return nil, fmt.Errorf("日期格式错误: %v", err)
	}

	parsedTime, err := time.Parse("15:04", birthTime)
	if err != nil {
		return nil, fmt.Errorf("时间格式错误: %v", err)
	}

	year := parsedDate.Year()
	hour := parsedTime.Hour()

	yearColumn := s.calculateYearColumn(year)
	monthColumn := s.calculateMonthColumn(parsedDate)
	dayColumn := s.calculateDayColumn(parsedDate)
	hourColumn := s.calculateHourColumn(dayColumn, hour)

	return []models.BaziColumn{
		yearColumn,
		monthColumn,
		dayColumn,
		hourColumn,
	}, nil
}

func (s *BaziService) calculateYearColumn(year int) models.BaziColumn {
	yearGanIndex := (year - 4) % 10
	yearZhiIndex := (year - 4) % 12
	
	if yearGanIndex < 0 {
		yearGanIndex += 10
	}
	if yearZhiIndex < 0 {
		yearZhiIndex += 12
	}

	return models.BaziColumn{
		Gan:       tianGan[yearGanIndex],
		Zhi:       diZhi[yearZhiIndex],
		GanWuXing: tianGanWuXing[tianGan[yearGanIndex]],
		ZhiWuXing: diZhiWuXing[diZhi[yearZhiIndex]],
	}
}

// calculateMonthColumn 使用农历节气计算月柱
// 月支由节气决定，而非公历月份：
// 立春(2/4) 开始为寅月，惊蛰(3/6) 开始为卯月，清明(4/5) 开始为辰月
// 立夏(5/6) 开始为巳月，芒种(6/6) 开始为午月，小暑(7/7) 开始为未月
// 立秋(8/8) 开始为申月，白露(9/8) 开始为酉月，寒露(10/8) 开始为戌月
// 立冬(11/8) 开始为亥月，大雪(12/7) 开始为子月，小寒(1/6) 开始为丑月
func (s *BaziService) calculateMonthColumn(date time.Time) models.BaziColumn {
	year := date.Year()
	yearGanIndex := (year - 4) % 10
	if yearGanIndex < 0 {
		yearGanIndex += 10
	}
	
	// 根据日期获取节气对应的月支
	monthZhi := solarterm.GetMonthDiZhi(date)
	
	// 如果无法通过节气获取月支，则使用默认方法（仅用于测试或fallback）
	if monthZhi == "" {
		// 获取月份对应的地支索引
		// 这里使用简化的计算方式，实际应该根据节气精确计算
		month := int(date.Month())
		// 简化处理：以立春(通常在2月4日左右)为寅月起点
		// 这里需要根据具体年份的节气日期精确计算
		monthZhiIndex := (month + 1) % 12 // 简化处理
		monthZhi = diZhi[monthZhiIndex]
	}
	
	// 根据年干和月支计算月干
	// 使用五虎遁诀计算月干
	monthGan := s.calculateMonthGan(yearGanIndex, monthZhi)
	
	return models.BaziColumn{
		Gan:       monthGan,
		Zhi:       monthZhi,
		GanWuXing: tianGanWuXing[monthGan],
		ZhiWuXing: diZhiWuXing[monthZhi],
	}
}

// calculateMonthGan 使用五虎遁诀计算月干
// 甲己之年丙作首，乙庚之岁戊为头，
// 丙辛必定寻庚起，丁壬壬位顺行流，
// 更有戊癸何方觅，甲寅之上好追求。
func (s *BaziService) calculateMonthGan(yearGanIndex int, monthZhi string) string {
	// 五虎遁诀口诀对应的起始天干索引
	// 甲(0)己(5)年从丙(2)开始，乙(1)庚(6)年从戊(4)开始
	// 丙(2)辛(7)年从庚(6)开始，丁(3)壬(8)年从壬(8)开始
	// 戊(4)癸(9)年从甲(0)开始
	yueGanStartMap := []int{2, 4, 6, 8, 0, 2, 4, 6, 8, 0}
	
	// 获取月支在地支中的索引
	monthZhiIndex := -1
	for i, zhi := range diZhi {
		if zhi == monthZhi {
			monthZhiIndex = i
			break
		}
	}
	
	if monthZhiIndex == -1 {
		// fallback to default calculation
		monthZhiIndex = 2 // default to 寅
	}
	
	// 计算月干索引
	// 从寅月(2)开始计算，所以需要减去2
	startGanIndex := yueGanStartMap[yearGanIndex]
	monthGanIndex := (startGanIndex + monthZhiIndex - 2) % 10
	if monthGanIndex < 0 {
		monthGanIndex += 10
	}
	
	return tianGan[monthGanIndex]
}

func (s *BaziService) calculateDayColumn(date time.Time) models.BaziColumn {
	daysSince1900 := int(date.Sub(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)).Hours() / 24)
	dayGanIndex := (daysSince1900 + 10) % 10
	dayZhiIndex := (daysSince1900 + 10) % 12

	if dayGanIndex < 0 {
		dayGanIndex += 10
	}
	if dayZhiIndex < 0 {
		dayZhiIndex += 12
	}

	return models.BaziColumn{
		Gan:       tianGan[dayGanIndex],
		Zhi:       diZhi[dayZhiIndex],
		GanWuXing: tianGanWuXing[tianGan[dayGanIndex]],
		ZhiWuXing: diZhiWuXing[diZhi[dayZhiIndex]],
	}
}

func (s *BaziService) calculateHourColumn(dayColumn models.BaziColumn, hour int) models.BaziColumn {
	dayGanIndex := s.findGanIndex(dayColumn.Gan)
	hourZhiIndex := ((hour + 1) / 2) % 12
	hourGanIndex := (dayGanIndex*2 + hourZhiIndex) % 10

	return models.BaziColumn{
		Gan:       tianGan[hourGanIndex],
		Zhi:       diZhi[hourZhiIndex],
		GanWuXing: tianGanWuXing[tianGan[hourGanIndex]],
		ZhiWuXing: diZhiWuXing[diZhi[hourZhiIndex]],
	}
}

func (s *BaziService) findGanIndex(gan string) int {
	for i, g := range tianGan {
		if g == gan {
			return i
		}
	}
	return 0
}

// 计算十二长生图
func (s *BaziService) calculateShiErChangSheng(bazi []models.BaziColumn) map[string]string {
	result := make(map[string]string)
	
	// 以日干（日主）为准计算十二长生
	dayColumn := bazi[2] // 日柱
	dayGanWuXing := dayColumn.GanWuXing
	
	// 计算各柱地支在日干五行下的长生状态
	for i, column := range bazi {
		columnName := []string{"年", "月", "日", "时"}[i]
		if changShengMap, exists := shiErChangSheng[dayGanWuXing]; exists {
			if changSheng, exists := changShengMap[column.Zhi]; exists {
				result[columnName+"支"] = changSheng
			}
		}
	}
	
	return result
}

// 增强八字柱子计算，添加主星、藏干、副星、纳音等
func (s *BaziService) enhanceBaziColumns(bazi []models.BaziColumn) []models.BaziColumn {
	dayColumn := bazi[2] // 日柱作为主星参考
	dayGan := dayColumn.Gan

	for i := range bazi {
		// 计算藏干
		bazi[i].CangGan = s.cangGanService.Calculate(bazi[i].Zhi)

		// 计算主星（以日干为准计算每柱的关系）
		bazi[i].ZhuXing = s.zhuXingService.Calculate(dayGan, bazi[i].Gan)

		// 计算副星（十神关系）
		bazi[i].FuXing = s.fuXingService.Calculate(dayGan, bazi[i])

		// 计算纳音（年柱）
		if i == 0 { // 年柱
			bazi[i].NaYin = s.naYinService.Calculate(bazi[i].Gan, bazi[i].Zhi)
		}

		// 计算星运（十二运程）
		bazi[i].XingYun = s.xingYunService.Calculate(dayColumn.GanWuXing, bazi[i].Zhi)

		// 计算自坐（天干在地支的状态）
		bazi[i].ZiZuo = s.ziZuoService.Calculate(bazi[i].Gan, bazi[i].Zhi)

		// 计算空亡
		bazi[i].KongWang = s.kongWangService.Calculate(dayColumn.Gan+dayColumn.Zhi, bazi[i].Zhi)

		// 计算神煞
		bazi[i].ShenSha = s.shenShaService.CalculateForColumn(dayGan, bazi[i])
	}

	return bazi
}