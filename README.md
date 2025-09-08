# Auspire 八字排盘系统

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](BUILD)

## 📖 项目概述

Auspire 是一个现代化的八字排盘系统，专为命理爱好者和专业人士设计。该系统基于传统子平八字理论，结合现代Web技术，提供专业级的四柱八字分析服务。

### 核心特性

- **精准排盘**: 基于节气的月柱计算，确保排盘准确性
- **全面分析**: 涵盖十神、藏干、纳音、十二长生等所有核心要素
- **专业界面**: 现代化深色主题设计，致敬传统命理文化
- **模块化架构**: 清晰的服务分离，便于维护和扩展
- **API驱动**: RESTful API设计，支持前后端分离

## 🏗️ 技术架构

### 后端技术栈

- **语言**: Go 1.19+
- **框架**: Gin Web Framework
- **数据库**: Redis (用户会话管理)
- **部署**: 单二进制部署，跨平台支持

### 前端技术栈

- **核心**: 原生HTML5, CSS3, JavaScript ES6+
- **样式**: 现代化深色主题，响应式设计
- **交互**: 丰富的动画效果和用户体验优化

### 目录结构

```
auspire/
├── handlers/           # HTTP请求处理器
├── middleware/         # 中间件(认证、日志等)
├── models/             # 数据模型定义
├── services/           # 核心业务逻辑服务
├── static/             # 静态资源文件
├── main.go            # 应用程序入口
├── go.mod             # Go模块定义
└── go.sum             # 依赖版本锁定
```

## 🎯 核心功能模块

### 1. 八字基础计算 (`bazi_service.go`)

负责四柱八字的基础排盘计算，包括：

- **年柱**: 基于公元年的天干地支计算
- **月柱**: 基于节气的精确月柱计算（非公历月份）
- **日柱**: 基于1900年基准的日柱计算
- **时柱**: 基于日柱和时辰的地支时柱计算

### 2. 十神关系分析 (`zhuxing_service.go`)

分析天干间的十神关系：

- 比肩、劫财（同类相助）
- 食神、伤官（我生者）
- 偏财、正财（我克者）
- 偏印、正印（生我者）
- 七杀、正官（克我者）

### 3. 藏干分析 (`canggan_service.go`)

分析地支中隐藏的天干组合：

- 子: 癸
- 丑: 己癸辛
- 寅: 甲丙戊
- 卯: 乙
- 辰: 戊乙癸
- 巳: 丙庚戊
- 午: 丁己
- 未: 己丁乙
- 申: 庚壬戊
- 酉: 辛
- 戌: 戊辛丁
- 亥: 壬甲

### 4. 纳音五行 (`nayin_service.go`)

六十甲子纳音五行对照：

- 海中金、炉中火、大林木等三十种纳音
- 基于年柱干支组合的纳音计算

### 5. 十二长生 (`shi_er_zhang_sheng.go`)

天干在十二地支中的长生状态：

- 长生、沐浴、冠带、临官、帝旺、衰、病、死、墓、绝、胎、养
- 阳干顺排，阴干逆排的正确实现

### 6. 星运分析 (`xingyun_service.go`)

地支在日主五行下的运程状态：

- 长生、沐浴、冠带、临官、帝旺、衰、病、死、墓、绝、胎、养

### 7. 自坐分析 (`zizuo_service.go`)

天干在所坐地支中的状态分析

### 8. 空亡计算 (`kongwang_service.go`)

基于日柱的旬空地支计算

### 9. 神煞分析 (`shensha_service.go`)

各类神煞星的定位计算：

- 天乙贵人、太极贵人、文昌贵人
- 将星、华盖、咸池、驿马、灾煞等

## 🚀 快速开始

### 环境要求

- Go 1.19 或更高版本
- Redis 6.0 或更高版本（可选，用于用户会话）

### 安装步骤

```bash
# 克隆项目
git clone <repository-url>
cd auspire

# 安装依赖
go mod tidy

# 编译
go build

# 运行
./auspire
```

### 配置环境变量

```bash
export REDIS_ADDR=localhost:6379
export REDIS_PASSWORD=
export JWT_SECRET=your-jwt-secret-key
export JWT_ISSUER=auspire
```

### 访问应用

应用启动后，默认监听在 `http://localhost:8080`

## 📡 API 接口

### 八字计算接口

```http
POST /api/bazi
Content-Type: application/json

{
  "name": "张三",
  "birthDate": "1990-03-15",
  "birthTime": "14:30"
}
```

### 喜用神计算接口

```http
POST /api/xiyongshen
Content-Type: application/json

{
  "name": "张三",
  "bazi": [...]
}
```

### 四柱八字综合分析接口

```http
POST /api/baziyuce
Content-Type: application/json

{
  "name": "张三",
  "bazi": [...]
}
```

## 📊 数据模型

### BaziColumn 结构

```go
type BaziColumn struct {
    Gan       string            // 天干
    Zhi       string            // 地支
    GanWuXing string            // 天干五行
    ZhiWuXing string            // 地支五行
    ZhuXing   string            // 主星(十神)
    CangGan   []string          // 藏干
    FuXing    []string          // 副星
    NaYin     string            // 纳音
    XingYun   string            // 星运
    ZiZuo     string            // 自坐
    KongWang  bool              // 空亡
    ShenSha   map[string]string // 神煞
}
```

## 🔧 开发指南

### 添加新功能模块

1. 在 `services/` 目录下创建新的服务文件
2. 实现相应的计算逻辑
3. 在 `models/bazi.go` 中添加必要的数据结构
4. 在 `handlers/bazi_handler.go` 中注册新的API端点

### 代码规范

- 遵循Go语言官方编码规范
- 每个服务文件专注于单一职责
- 函数和变量命名采用驼峰命名法
- 关键算法添加详细注释说明

## 🧪 测试策略

### 单元测试

```bash
go test ./services/ -v
```

### 集成测试

通过API接口进行端到端测试

## 📈 性能优化

- 模块化服务设计减少内存占用
- 缓存常用计算结果
- 避免重复计算

## 🤝 贡献指南

欢迎任何形式的贡献！

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 🙏 致谢

- 感谢所有传统命理学大师的智慧传承
- 感谢开源社区提供的优秀工具和库
- 感谢所有测试用户提供的宝贵反馈

---

**免责声明**: 本系统仅供学习交流使用，命理分析结果仅供参考，不作为任何决策依据。