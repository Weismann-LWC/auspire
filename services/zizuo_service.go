package services

// ZiZuoService 自坐服务
type ZiZuoService struct {
	cangGanService *CangGanService
	xingYunService *XingYunService
}

func NewZiZuoService() *ZiZuoService {
	return &ZiZuoService{
		cangGanService: NewCangGanService(),
		xingYunService: NewXingYunService(),
	}
}

// Calculate 计算自坐（天干在地支的状态）
func (s *ZiZuoService) Calculate(gan, zhi string) string {
	ganWuXing := tianGanWuXing[gan]
	
	// 查看地支藏干中是否有本天干
	cangGan := s.cangGanService.Calculate(zhi)
	for _, cGan := range cangGan {
		if cGan == gan {
			return "自坐本气"
		}
	}

	// 计算天干在该地支的长生状态
	changSheng := s.xingYunService.Calculate(ganWuXing, zhi)
	if changSheng != "" {
		return "自坐" + changSheng
	}

	return ""
}