package solarterm

import (
	"time"
)

// 节气与地支的对应关系
// 每个节气对应一个地支，月支由节气决定，而非公历月份

// 节气名称常量
const (
	Lichun  = "立春"  // 2月3-5日
	Yushui  = "雨水"  // 2月18-20日
	Jingzhe  = "惊蛰"  // 3月5-7日
	Chunfen  = "春分"  // 3月20-22日
	Qingming = "清明" // 4月4-6日
	Guyu     = "谷雨" // 4月19-21日
	Lixia    = "立夏" // 5月5-7日
	Xiaoman  = "小满" // 5月20-22日
	Mangzhong = "芒种" // 6月5-7日
	Xiazhi   = "夏至" // 6月21-22日
	Xiaoshu  = "小暑" // 7月6-8日
	Dashu    = "大暑" // 7月22-24日
	Liqiu    = "立秋" // 8月7-9日
	Chushu   = "处暑" // 8月22-24日
	Bailu    = "白露" // 9月7-9日
	Qiufen   = "秋分" // 9月22-24日
	Hanlu    = "寒露" // 10月8-9日
	Shuangjiang = "霜降" // 10月23-24日
	Lidong   = "立冬" // 11月7-8日
	Xiaoxue  = "小雪" // 11月22-23日
	Daxue    = "大雪" // 12月6-8日
	Dongzhi  = "冬至" // 12月21-23日
	Xiaohan  = "小寒" // 1月5-7日
	Dahan    = "大寒" // 1月20-21日
)

// 节气与地支的对应关系
// 节气决定了月支，每两个节气对应一个月份
var SolarTermToDiZhi = map[string]string{
	Lichun:  "寅", // 立春开始为寅月
	Yushui:  "寅",
	Jingzhe:  "卯", // 惊蛰开始为卯月
	Chunfen:  "卯",
	Qingming: "辰", // 清明开始为辰月
	Guyu:     "辰",
	Lixia:    "巳", // 立夏开始为巳月
	Xiaoman:  "巳",
	Mangzhong: "午", // 芒种开始为午月
	Xiazhi:   "午",
	Xiaoshu:  "未", // 小暑开始为未月
	Dashu:    "未",
	Liqiu:    "申", // 立秋开始为申月
	Chushu:   "申", // 处暑仍为申月
	Bailu:    "酉", // 白露开始为酉月
	Qiufen:   "酉",
	Hanlu:    "戌", // 寒露开始为戌月
	Shuangjiang: "戌",
	Lidong:   "亥", // 立冬开始为亥月
	Xiaoxue:  "亥",
	Daxue:    "子", // 大雪开始为子月
	Dongzhi:  "子",
	Xiaohan:  "丑", // 小寒开始为丑月
	Dahan:    "丑", // 大寒结束丑月，立春开始新一轮
}

// 获取节气对应的地支
func GetDiZhiFromSolarTerm(solarTerm string) string {
	if dizhi, exists := SolarTermToDiZhi[solarTerm]; exists {
		return dizhi
	}
	return ""
}

// 根据日期获取节气
// 这里使用简化的计算方法，实际应用中可能需要更精确的天文算法
func GetSolarTerm(date time.Time) string {
	month := int(date.Month())
	day := date.Day()
	
	// 简化的节气判断，实际日期可能因年份而略有差异
	switch month {
	case 1:
		if day >= 5 && day <= 7 {
			return Xiaohan
		} else if day >= 20 && day <= 21 {
			return Dahan
		}
	case 2:
		if day >= 3 && day <= 5 {
			return Lichun
		} else if day >= 18 && day <= 20 {
			return Yushui
		}
	case 3:
		if day >= 5 && day <= 7 {
			return Jingzhe
		} else if day >= 20 && day <= 22 {
			return Chunfen
		}
	case 4:
		if day >= 4 && day <= 6 {
			return Qingming
		} else if day >= 19 && day <= 21 {
			return Guyu
		}
	case 5:
		if day >= 5 && day <= 7 {
			return Lixia
		} else if day >= 20 && day <= 22 {
			return Xiaoman
		}
	case 6:
		if day >= 5 && day <= 7 {
			return Mangzhong
		} else if day >= 21 && day <= 22 {
			return Xiazhi
		}
	case 7:
		if day >= 6 && day <= 8 {
			return Xiaoshu
		} else if day >= 22 && day <= 24 {
			return Dashu
		}
	case 8:
		if day >= 7 && day <= 9 {
			return Liqiu  // 立秋(8月8日左右)
		} else if day >= 22 && day <= 24 {
			return Chushu // 处暑(8月23日左右)
		} else if day > 9 && day < 22 {
			// 8月10日至8月21日期间，仍然属于立秋节气周期
			return Liqiu
		}
	case 9:
		if day >= 7 && day <= 9 {
			return Bailu
		} else if day >= 22 && day <= 24 {
			return Qiufen
		} else if day > 9 && day < 22 {
			// 9月10日至9月21日期间，仍然属于白露节气周期
			return Bailu
		}
	case 10:
		if day >= 8 && day <= 9 {
			return Hanlu
		} else if day >= 23 && day <= 24 {
			return Shuangjiang
		} else if day > 9 && day < 23 {
			// 10月10日至10月22日期间，仍然属于寒露节气周期
			return Hanlu
		}
	case 11:
		if day >= 7 && day <= 8 {
			return Lidong
		} else if day >= 22 && day <= 23 {
			return Xiaoxue
		} else if day > 8 && day < 22 {
			// 11月9日至11月21日期间，仍然属于立冬节气周期
			return Lidong
		}
	case 12:
		if day >= 6 && day <= 8 {
			return Daxue
		} else if day >= 21 && day <= 23 {
			return Dongzhi
		} else if day > 8 && day < 21 {
			// 12月9日至12月20日期间，仍然属于大雪节气周期
			return Daxue
		}
	}
	
	return ""
}

// 获取月柱地支（基于节气）
func GetMonthDiZhi(date time.Time) string {
	// 实际应用中需要更复杂的算法来精确计算节气日期
	// 这里使用简化的实现
	solarTerm := GetSolarTerm(date)
	return GetDiZhiFromSolarTerm(solarTerm)
}