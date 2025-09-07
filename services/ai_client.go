package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"auspire/models"
)

type AIClient struct {
	apiKey   string
	baseURL  string
	client   *http.Client
}

type AIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

func NewAIClient() *AIClient {
	// 优先使用 DeepSeek API
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	baseURL := "https://api.deepseek.com/v1/chat/completions"
	
	if apiKey == "" {
		// 备用 OpenAI API
		apiKey = os.Getenv("OPENAI_API_KEY")
		baseURL = "https://api.openai.com/v1/chat/completions"
	}

	return &AIClient{
		apiKey:  apiKey,
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *AIClient) AnalyzeFortune(req models.FortuneRequest) (*models.FortuneResponse, error) {
	if c.apiKey == "" {
		// 如果没有API密钥，返回模拟数据
		return c.getMockFortune(req), nil
	}

	baziStr := fmt.Sprintf("年柱：%s%s（%s%s），月柱：%s%s（%s%s），日柱：%s%s（%s%s），时柱：%s%s（%s%s）",
		req.Bazi[0].Gan, req.Bazi[0].Zhi, req.Bazi[0].GanWuXing, req.Bazi[0].ZhiWuXing,
		req.Bazi[1].Gan, req.Bazi[1].Zhi, req.Bazi[1].GanWuXing, req.Bazi[1].ZhiWuXing,
		req.Bazi[2].Gan, req.Bazi[2].Zhi, req.Bazi[2].GanWuXing, req.Bazi[2].ZhiWuXing,
		req.Bazi[3].Gan, req.Bazi[3].Zhi, req.Bazi[3].GanWuXing, req.Bazi[3].ZhiWuXing)

	prompt := fmt.Sprintf(`你是一位资深的命理师，请根据以下信息为客户分析2024年的流年运势：

姓名：%s
出生日期：%s
生辰八字：%s

请从以下几个方面进行详细分析，每个方面用50-80字：
1. 整体运势
2. 事业运势
3. 财运
4. 健康运势
5. 感情运势
6. 建议

请用温和、积极的语调，避免过于消极的预测。返回格式为JSON，包含以下字段：
{
  "overallFortune": "整体运势分析",
  "career": "事业运势分析",
  "wealth": "财运分析",
  "health": "健康运势分析",
  "relationship": "感情运势分析",
  "advice": "建议"
}`, req.Name, req.BirthDate, baziStr)

	aiReq := AIRequest{
		Model: c.getModel(),
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(aiReq)
	if err != nil {
		return nil, fmt.Errorf("编码请求失败: %v", err)
	}

	httpReq, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API请求失败，状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}

	var aiResp AIResponse
	if err := json.Unmarshal(body, &aiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	if len(aiResp.Choices) == 0 {
		return nil, fmt.Errorf("AI响应为空")
	}

	content := aiResp.Choices[0].Message.Content
	
	// 尝试解析JSON响应
	var fortuneData struct {
		OverallFortune string `json:"overallFortune"`
		Career         string `json:"career"`
		Wealth         string `json:"wealth"`
		Health         string `json:"health"`
		Relationship   string `json:"relationship"`
		Advice         string `json:"advice"`
	}

	if err := json.Unmarshal([]byte(content), &fortuneData); err != nil {
		// 如果JSON解析失败，使用纯文本作为整体运势
		fortuneData.OverallFortune = content
		fortuneData.Career = "事业方面需要稳扎稳打，保持谨慎态度。"
		fortuneData.Wealth = "财运平稳，适合保守理财。"
		fortuneData.Health = "注意身体健康，保持良好生活习惯。"
		fortuneData.Relationship = "感情运势良好，真诚待人。"
		fortuneData.Advice = "保持积极心态，把握机遇。"
	}

	return &models.FortuneResponse{
		Name:           req.Name,
		CurrentYear:    "2024",
		OverallFortune: fortuneData.OverallFortune,
		Career:         fortuneData.Career,
		Wealth:         fortuneData.Wealth,
		Health:         fortuneData.Health,
		Relationship:   fortuneData.Relationship,
		Advice:         fortuneData.Advice,
	}, nil
}

// 新增十二长生专项分析
func (c *AIClient) AnalyzeChangSheng(name string, shiErChangSheng map[string]string, bazi []models.BaziColumn) (map[string]string, error) {
	if c.apiKey == "" {
		// 如果没有API密钥，返回模拟数据
		return c.getMockChangShengAnalysis(shiErChangSheng), nil
	}

	// 构建十二长生信息字符串
	changShengStr := ""
	for position, state := range shiErChangSheng {
		changShengStr += fmt.Sprintf("%s：%s；", position, state)
	}

	// 构建八字信息
	baziStr := fmt.Sprintf("年柱：%s%s（%s%s），月柱：%s%s（%s%s），日柱：%s%s（%s%s），时柱：%s%s（%s%s）",
		bazi[0].Gan, bazi[0].Zhi, bazi[0].GanWuXing, bazi[0].ZhiWuXing,
		bazi[1].Gan, bazi[1].Zhi, bazi[1].GanWuXing, bazi[1].ZhiWuXing,
		bazi[2].Gan, bazi[2].Zhi, bazi[2].GanWuXing, bazi[2].ZhiWuXing,
		bazi[3].Gan, bazi[3].Zhi, bazi[3].GanWuXing, bazi[3].ZhiWuXing)

	prompt := fmt.Sprintf(`你是一位精通十二长生理论的命理师，请根据以下信息为客户分析人生各个阶段的运势特点：

姓名：%s
生辰八字：%s
十二长生状态：%s

请从十二长生的角度，分析每个柱位（年、月、日、时）所代表的人生阶段特征：

1. 年柱代表童年至少年时期（0-15岁）
2. 月柱代表青年时期（16-30岁）  
3. 日柱代表中年时期（31-45岁）
4. 时柱代表中老年至晚年时期（46岁以后）

请针对每个时期的长生状态，给出该阶段的特点、机遇、挑战和建议。每个阶段用60-100字详细分析。

返回格式为JSON，包含以下字段：
{
  "childhood": "童年至少年时期分析（0-15岁）",
  "youth": "青年时期分析（16-30岁）",
  "middle": "中年时期分析（31-45岁）", 
  "later": "中老年至晚年时期分析（46岁以后）"
}

注意：要结合具体的十二长生状态（长生、沐浴、冠带、临官、帝旺、衰、病、死、墓、绝、胎、养）来分析每个人生阶段的特征。`, 
		name, baziStr, changShengStr)

	aiReq := AIRequest{
		Model: c.getModel(),
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(aiReq)
	if err != nil {
		return nil, fmt.Errorf("编码请求失败: %v", err)
	}

	httpReq, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API请求失败，状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}

	var aiResp AIResponse
	if err := json.Unmarshal(body, &aiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	if len(aiResp.Choices) == 0 {
		return nil, fmt.Errorf("AI响应为空")
	}

	content := aiResp.Choices[0].Message.Content
	
	// 尝试解析JSON响应
	var changShengData struct {
		Childhood string `json:"childhood"`
		Youth     string `json:"youth"`
		Middle    string `json:"middle"`
		Later     string `json:"later"`
	}

	if err := json.Unmarshal([]byte(content), &changShengData); err != nil {
		// 如果JSON解析失败，返回默认分析
		return c.getMockChangShengAnalysis(shiErChangSheng), nil
	}

	result := map[string]string{
		"childhood": changShengData.Childhood,
		"youth":     changShengData.Youth,
		"middle":    changShengData.Middle,
		"later":     changShengData.Later,
	}

	return result, nil
}

func (c *AIClient) getModel() string {
	if c.baseURL == "https://api.deepseek.com/v1/chat/completions" {
		return "deepseek-chat"
	}
	return "gpt-3.5-turbo"
}

func (c *AIClient) getMockFortune(req models.FortuneRequest) *models.FortuneResponse {
	return &models.FortuneResponse{
		Name:        req.Name,
		CurrentYear: "2024",
		OverallFortune: "2024年对您来说是平稳中带有机遇的一年。您的生辰八字显示出稳重的特质，这一年适合踏实前行，把握每一个细节。虽然不会有惊天动地的变化，但通过努力积累，会有不错的收获。建议保持开放的心态，迎接新的挑战。",
		Career: "事业运势整体向好，上半年可能会遇到一些挑战，但这些都是成长的机会。您的八字显示出很强的学习能力，适合在专业领域深耕。下半年有升职或转型的机会，建议提前做好准备，抓住关键时刻。",
		Wealth: "财运方面呈现稳中有升的趋势。正财运较强，通过正当途径获得的收入会比较稳定。投资方面建议谨慎，不宜冒险。年中可能会有一笔意外收入，但也要注意理财规划，避免不必要的支出。",
		Health: "健康运势良好，但需要注意作息规律。您的八字中显示容易因工作压力影响睡眠质量，建议多做运动，保持身心平衡。秋季需特别注意呼吸道健康，及时调养身体。",
		Relationship: "感情运势温和稳定。已有伴侣的人感情会更加深厚，适合考虑长远规划。单身者有机会在工作或学习环境中遇到合适的人，但不要急于求成，让感情自然发展。家庭关系和谐，是您重要的精神支撑。",
		Advice: "2024年的关键词是「稳中求进」。建议您保持谦逊的态度，多学习新知识，提升自己的竞争力。在人际交往中要真诚待人，建立良好的人际关系网络。遇到困难时不要急躁，相信时间会给您最好的答案。",
	}
}

// 十二长生模拟分析数据
func (c *AIClient) getMockChangShengAnalysis(shiErChangSheng map[string]string) map[string]string {
	// 获取各柱位的长生状态
	yearState := shiErChangSheng["年支"]
	monthState := shiErChangSheng["月支"]
	dayState := shiErChangSheng["日支"]
	timeState := shiErChangSheng["时支"]
	
	return map[string]string{
		"childhood": fmt.Sprintf("童年至少年时期（0-15岁）处于%s状态，%s", yearState, getStageDescription("童年", yearState)),
		"youth":     fmt.Sprintf("青年时期（16-30岁）处于%s状态，%s", monthState, getStageDescription("青年", monthState)),
		"middle":    fmt.Sprintf("中年时期（31-45岁）处于%s状态，%s", dayState, getStageDescription("中年", dayState)),
		"later":     fmt.Sprintf("中老年至晚年时期（46岁以后）处于%s状态，%s", timeState, getStageDescription("晚年", timeState)),
	}
}

// 根据人生阶段和长生状态生成描述
func getStageDescription(stage, state string) string {
	descriptions := map[string]map[string]string{
		"长生": {
			"童年": "这个阶段充满生机和希望，学习能力强，身体健康，家庭关爱有加。建议多培养兴趣爱好，打好基础。",
			"青年": "事业起步顺利，有贵人相助，适合学习新技能和建立人脉。感情方面也容易遇到合适的对象。",
			"中年": "事业迎来新的发展机遇，创新能力强，适合开拓新领域。家庭和谐，子女有成就。",
			"晚年": "身体健康，精神充实，有新的兴趣和目标。子孙满堂，晚年生活丰富多彩。",
		},
		"沐浴": {
			"童年": "性格活泼，但需要正确引导。学习中可能遇到一些波折，需要家长耐心教导。",
			"青年": "感情丰富，但要注意选择。工作中需要磨练，避免急躁情绪。多听取长辈意见。",
			"中年": "事业可能有变动，需要调整心态。感情生活需要更多沟通和理解。",
			"晚年": "保持开放心态，学习新事物。注意身体保养，避免过度劳累。",
		},
		"冠带": {
			"童年": "聪明懂事，学习成绩优秀，深受师长喜爱。适合培养领导才能和责任感。",
			"青年": "事业发展稳定，有一定的社会地位。适合承担更多责任，展现个人能力。",
			"中年": "事业达到一定高度，社会影响力增强。家庭责任重，需要平衡工作与家庭。",
			"晚年": "德高望重，受人尊敬。适合传承经验，指导后辈。身体健康需要关注。",
		},
		"临官": {
			"童年": "表现出色，有领导潜质。学习能力强，但要注意不要过于争强好胜。",
			"青年": "事业蒸蒸日上，有升职机会。人际关系良好，但要保持谦逊态度。",
			"中年": "事业达到高峰，权威地位稳固。但要注意身体健康，避免过度操劳。",
			"晚年": "影响力深远，但要学会放手。享受天伦之乐，保持平和心态。",
		},
		"帝旺": {
			"童年": "天资聪颖，各方面表现突出。但要培养谦逊品格，避免骄傲自满。",
			"青年": "事业如日中天，成就卓越。但要注意处理好人际关系，保持低调。",
			"中年": "达到人生巅峰，影响力广泛。要善用权力，回馈社会，关注家庭和谐。",
			"晚年": "成就辉煌，但要学会享受生活。传承智慧给后代，保持身心健康。",
		},
		"衰": {
			"童年": "可能遇到一些挫折，需要家庭更多关爱。培养坚韧品格，为将来奠定基础。",
			"青年": "事业发展缓慢，需要耐心积累。避免急于求成，注重能力提升。",
			"中年": "事业可能遇到瓶颈，需要调整策略。保持积极心态，寻找新的机会。",
			"晚年": "要注意身体保养，保持规律生活。享受平淡的幸福，不强求太多。",
		},
		"病": {
			"童年": "可能身体较弱，需要特别关注健康。学习上可能有困难，需要耐心引导。",
			"青年": "工作中可能遇到困难，需要坚持不懈。注意身心健康，避免过度压力。",
			"中年": "事业和健康都需要特别关注。调整生活方式，寻求专业帮助。",
			"晚年": "健康是最重要的，要积极治疗。保持乐观心态，享受简单的快乐。",
		},
		"死": {
			"童年": "可能经历一些困难，但这是成长的机会。培养坚强意志，为将来做准备。",
			"青年": "事业可能陷入低谷，但不要放弃希望。这是转机前的考验。",
			"中年": "人生可能遇到重大转折，需要勇敢面对。相信困难只是暂时的。",
			"晚年": "回顾人生，珍惜当下。保持内心平静，享受精神上的富足。",
		},
		"墓": {
			"童年": "性格内向，需要多鼓励。潜力深厚，适合深度学习和思考。",
			"青年": "事业发展需要时间，不要急躁。潜心修炼，为将来的爆发做准备。",
			"中年": "经历沉淀后，智慧增长。适合从事研究或传承类工作。",
			"晚年": "内心富足，精神世界丰富。适合静修养性，传授人生智慧。",
		},
		"绝": {
			"童年": "可能面临困境，但这是重新开始的机会。培养独立自强的品格。",
			"青年": "事业可能需要重新规划，不要畏惧改变。这是转型的最佳时期。",
			"中年": "人生可能有重大变化，勇敢拥抱新的开始。放下过去，展望未来。",
			"晚年": "经历过人生起伏，心境更加豁达。珍惜简单的幸福。",
		},
		"胎": {
			"童年": "潜力无限，需要悉心培养。保护好天真本性，培养创造力。",
			"青年": "事业处于孕育期，需要积累和准备。保持学习心态，等待时机。",
			"中年": "新的机遇正在孕育，保持敏锐嗅觉。适合投资未来，布局长远。",
			"晚年": "虽然年岁已高，但仍有新的可能。保持年轻心态，拥抱变化。",
		},
		"养": {
			"童年": "需要充足的爱和关怀，培养良好的品格。为将来的成长积蓄力量。",
			"青年": "事业需要时间培养，不要急于求成。投资自己，提升能力。",
			"中年": "经过多年积累，实力渐显。适合稳扎稳打，逐步发展。",
			"晚年": "颐养天年，享受生活。注重身心保养，传承人生经验。",
		},
	}
	
	if stageDesc, exists := descriptions[state]; exists {
		if desc, exists := stageDesc[stage]; exists {
			return desc
		}
	}
	
	return "这个阶段需要根据具体情况调整策略，保持积极心态，相信未来会更好。"
}