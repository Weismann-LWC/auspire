package models

type BaziRequest struct {
	Name      string `json:"name" binding:"required"`
	BirthDate string `json:"birthDate" binding:"required"`
	BirthTime string `json:"birthTime" binding:"required"`
}

type BaziColumn struct {
	Gan       string            `json:"gan"`
	Zhi       string            `json:"zhi"`
	GanWuXing string            `json:"ganWuXing"`
	ZhiWuXing string            `json:"zhiWuXing"`
	ZhuXing   string            `json:"zhuXing,omitempty"`    // 主星
	CangGan   []string          `json:"cangGan,omitempty"`   // 藏干
	FuXing    []string          `json:"fuXing,omitempty"`    // 副星
	NaYin     string            `json:"naYin,omitempty"`     // 纳音
	XingYun   string            `json:"xingYun,omitempty"`   // 星运
	ZiZuo     string            `json:"ziZuo,omitempty"`     // 自坐
	KongWang  bool              `json:"kongWang,omitempty"`  // 空亡
	ShenSha   map[string]string `json:"shenSha,omitempty"`   // 神煞
}

type BaziResponse struct {
	Name            string            `json:"name"`
	Bazi            []BaziColumn      `json:"bazi"`
	ShiErChangSheng map[string]string `json:"shiErChangSheng,omitempty"`
	Error           string            `json:"error,omitempty"`
}

type FortuneRequest struct {
	Name      string       `json:"name" binding:"required"`
	Bazi      []BaziColumn `json:"bazi" binding:"required"`
	BirthDate string       `json:"birthDate" binding:"required"`
}

type FortuneResponse struct {
	Name           string `json:"name"`
	CurrentYear    string `json:"currentYear"`
	OverallFortune string `json:"overallFortune"`
	Career         string `json:"career"`
	Wealth         string `json:"wealth"`
	Health         string `json:"health"`
	Relationship   string `json:"relationship"`
	Advice         string `json:"advice"`
	Error          string `json:"error,omitempty"`
}

type LifeStageRequest struct {
	Name            string            `json:"name" binding:"required"`
	Bazi            []BaziColumn      `json:"bazi" binding:"required"`
	ShiErChangSheng map[string]string `json:"shiErChangSheng" binding:"required"`
}

type LifeStageResponse struct {
	Name      string            `json:"name"`
	Analysis  map[string]string `json:"analysis"`
	Error     string            `json:"error,omitempty"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

// XiYongShenRequest 喜用神请求
type XiYongShenRequest struct {
	Name string       `json:"name" binding:"required"`
	Bazi []BaziColumn `json:"bazi" binding:"required"`
}

// XiYongShenResult 喜用神计算结果
type XiYongShenResult struct {
	Name          string            `json:"name"`
	RiZhu         string            `json:"riZhu"`                 // 日主
	RiZhuStrength string            `json:"riZhuStrength"`         // 日主强弱
	WuXingScores  map[string]int    `json:"wuXingScores"`          // 五行得分
	XiYongShen    string            `json:"xiYongShen"`            // 喜用神
	Logic         []string          `json:"logic"`                 // 计算逻辑
	Error         string            `json:"error,omitempty"`
}

// BaziyuceRequest 四柱八字综合分析请求
type BaziyuceRequest struct {
	Name string       `json:"name" binding:"required"`
	Bazi []BaziColumn `json:"bazi" binding:"required"`
}

// AnalysisStep 分析步骤
type AnalysisStep struct {
	Title   string   `json:"title"`
	Content []string `json:"content"`
}

// BaziyuceResult 四柱八字综合分析结果
type BaziyuceResult struct {
	Name  string         `json:"name"`
	Steps []AnalysisStep `json:"steps"`
	Error string         `json:"error,omitempty"`
}