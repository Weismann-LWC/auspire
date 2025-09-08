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

// ShiErZhangSheng 十二长生表 (Twelve Longevity Cycle Table)
//
// This table defines the twelve stages of existence for each Heavenly Stem.
// The concept reflects the natural cycle of growth and decline in Chinese metaphysics.
//
// Key Principles:
// 1. Yang Heavenly Stems (甲, 丙, 戊, 庚, 壬) follow顺行 (forward sequence)
// 2. Yin Heavenly Stems (乙, 丁, 己, 辛, 癸) follow逆行 (reverse sequence)
// 3. Each stem has its own unique cycle progression through the twelve Earthly Branches
//
// The twelve stages represent:
// 1. 长生 (Chang Sheng) - Longevity/Birth - Beginning of growth
// 2. 沐浴 (Mu Yu) - Bathing - Cleansing and preparation
// 3. 冠带 (Guan Dai) - Crown/Wearing - Maturation begins
// 4. 临官 (Lin Guan) - Approaching Officerdom - Near peak strength
// 5. 帝旺 (Di Wang) - Imperial Prosperity - Peak power
// 6. 衰 (Shuai) - Decline - Beginning of weakening
// 7. 病 (Bing) - Sickness - Further deterioration
// 8. 死 (Si) - Death - End of current cycle
// 9. 墓 (Mu) - Tomb/Storage - Energy stored away
// 10. 绝 (Jue) - Severance/Cutting off - Complete separation
// 11. 胎 (Tai) - Conception/Fetus - New beginning forming
// 12. 养 (Yang) - Nurturing - Preparation for rebirth
var ShiErZhangSheng = map[string][]string{
	// Yang Heavenly Stems (Forward Sequence)
	//
	// 甲木 (Jia Wood) - Begins at亥 (Hai) - Water nourishes Wood
	"甲": {"亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌"},
	
	// 丙火 (Bing Fire) - Begins at寅 (Yin) - Wood fuels Fire
	"丙": {"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"},
	
	// 戊土 (Wu Earth) - Follows same as Bing Fire
	"戊": {"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"},
	
	// 庚金 (Geng Metal) - Begins at巳 (Si) - Earth produces Metal
	"庚": {"巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰"},
	
	// 壬水 (Ren Water) - Begins at申 (Shen) - Metal generates Water
	"壬": {"申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未"},

	// Yin Heavenly Stems (Reverse Sequence)
	//
	// 乙木 (Yi Wood) - Begins at午 (Wu) - Fire nourishes Wood in reverse
	"乙": {"午", "巳", "辰", "卯", "寅", "丑", "子", "亥", "戌", "酉", "申", "未"},
	
	// 丁火 (Ding Fire) - Begins at酉 (You) - Metal generates Fire in reverse
	"丁": {"酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子", "亥", "戌"},
	
	// 己土 (Ji Earth) - Follows same as Ding Fire
	"己": {"酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑", "子", "亥", "戌"},
	
	// 辛金 (Xin Metal) - Begins at子 (Zi) - Water generates Metal in reverse
	"辛": {"子", "亥", "戌", "酉", "申", "未", "午", "巳", "辰", "卯", "寅", "丑"},
	
	// 癸水 (Gui Water) - Begins at卯 (Mao) - Wood generates Water in reverse
	"癸": {"卯", "寅", "丑", "子", "亥", "戌", "酉", "申", "未", "午", "巳", "辰"},
}

// ZhangShengNames 十二长生名称 (Twelve Longevity Stage Names)
//
// These represent the cyclical stages of energy manifestation in Chinese metaphysics.
// Each stage describes a different phase in the life cycle of elemental energy.
var ZhangShengNames = []string{
	"长生", // Birth/Beginning of growth
	"沐浴", // Cleansing/preparation 
	"冠带", // Maturation begins
	"临官", // Approaching peak strength
	"帝旺", // Peak power/maximum influence
	"衰",   // Beginning of decline
	"病",   // Further weakening
	"死",   // End of current cycle
	"墓",   // Storage/rest period
	"绝",   // Complete severance/separation
	"胎",   // New formation/conception
	"养",   // Nurturing/preparation for rebirth
}

// YangGan 阳干列表 (Yang Heavenly Stems)
//
// In Chinese metaphysics, Heavenly Stems are classified as Yang or Yin.
// Yang stems follow forward sequences, Yin stems follow reverse sequences.
var YangGan = []string{"甲", "丙", "戊", "庚", "壬"}

// YinGan 阴干列表 (Yin Heavenly Stems)
var YinGan = []string{"乙", "丁", "己", "辛", "癸"}

// GetZhangShengPosition 获取天干在地支中的长生位置
//
// This function determines what stage of the twelve longevity cycle a particular
// Heavenly Stem is in when placed in a specific Earthly Branch.
//
// Parameters:
//   - tianGan: The Heavenly Stem to analyze (e.g., "甲", "乙", "丙")
//   - diZhi: The Earthly Branch position (e.g., "子", "丑", "寅")
//
// Returns:
//   The corresponding longevity stage name (e.g., "长生", "帝旺", "墓")
//   Empty string if the stem or branch is invalid
//
// Example:
//   GetZhangShengPosition("甲", "亥") returns "长生"
//   GetZhangShengPosition("乙", "午") returns "长生"
func GetZhangShengPosition(tianGan, diZhi string) string {
	// Check if the Heavenly Stem is valid
	shengPositions, exists := ShiErZhangSheng[tianGan]
	if !exists {
		return ""
	}
	
	// Find the position of the Earthly Branch in the stem's longevity sequence
	for i, zhi := range shengPositions {
		if zhi == diZhi {
			return ZhangShengNames[i]
		}
	}
	
	return ""
}

// IsYangGan 判断是否为阳干
//
// Determines if a given Heavenly Stem is classified as Yang.
// Yang stems follow forward sequences in calculations.
//
// Parameters:
//   - gan: The Heavenly Stem to check (e.g., "甲", "乙")
//
// Returns:
//   true if the stem is Yang, false otherwise
func IsYangGan(gan string) bool {
	for _, g := range YangGan {
		if g == gan {
			return true
		}
	}
	return false
}

// IsYinGan 判断是否为阴干
//
// Determines if a given Heavenly Stem is classified as Yin.
// Yin stems follow reverse sequences in calculations.
//
// Parameters:
//   - gan: The Heavenly Stem to check (e.g., "甲", "乙")
//
// Returns:
//   true if the stem is Yin, false otherwise
func IsYinGan(gan string) bool {
	for _, g := range YinGan {
		if g == gan {
			return true
		}
	}
	return false
}