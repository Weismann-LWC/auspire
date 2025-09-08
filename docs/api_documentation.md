# Auspire API 文档

## 📡 接口概览

Auspire 提供 RESTful API 接口，支持 JSON 格式的数据交换。

### 基础URL

```
http://localhost:8080/api
```

### 认证方式

- **公共接口**: 无需认证，可直接访问
- **保护接口**: 需要 JWT Token 认证

### 响应格式

所有接口均返回 JSON 格式数据：

```json
{
  "status": "success|error",
  "data": {},
  "message": "optional message"
}
```

## 🔐 认证接口

### 用户注册

```http
POST /api/register
```

**请求参数**

| 字段名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| username | string | 是 | 用户名 |
| email | string | 是 | 邮箱地址 |
| password | string | 是 | 密码(最少6位) |

**请求示例**

```json
{
  "username": "zhangsan",
  "email": "zhangsan@example.com",
  "password": "password123"
}
```

**响应示例**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "1234567890",
    "username": "zhangsan",
    "email": "zhangsan@example.com",
    "created_at": "2023-01-01T00:00:00Z"
  }
}
```

### 用户登录

```http
POST /api/login
```

**请求参数**

| 字段名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| username | string | 是 | 用户名 |
| password | string | 是 | 密码 |

**请求示例**

```json
{
  "username": "zhangsan",
  "password": "password123"
}
```

**响应示例**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "1234567890",
    "username": "zhangsan",
    "email": "zhangsan@example.com",
    "created_at": "2023-01-01T00:00:00Z"
  }
}
```

### 获取用户资料

```http
GET /api/profile
Authorization: Bearer <token>
```

**响应示例**

```json
{
  "id": "1234567890",
  "username": "zhangsan",
  "email": "zhangsan@example.com",
  "created_at": "2023-01-01T00:00:00Z"
}
```

## 🔮 八字计算接口

### 基础八字计算

```http
POST /api/bazi
```

**请求参数**

| 字段名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| name | string | 是 | 姓名 |
| birthDate | string | 是 | 出生日期(YYYY-MM-DD) |
| birthTime | string | 是 | 出生时间(HH:MM) |

**请求示例**

```json
{
  "name": "张三",
  "birthDate": "1990-03-15",
  "birthTime": "14:30"
}
```

**响应示例**

```json
{
  "name": "张三",
  "bazi": [
    {
      "gan": "庚",
      "zhi": "午",
      "ganWuXing": "金",
      "zhiWuXing": "火",
      "zhuXing": "比肩",
      "cangGan": ["丁", "己"],
      "fuXing": ["伤官", "正印"],
      "naYin": "路旁土",
      "xingYun": "死",
      "ziZuo": "病",
      "kongWang": false,
      "shenSha": {
        "天乙贵人": "酉"
      }
    },
    {
      "gan": "己",
      "zhi": "酉",
      "ganWuXing": "土",
      "zhiWuXing": "金",
      "zhuXing": "劫财",
      "cangGan": ["辛"],
      "fuXing": ["比肩"],
      "naYin": "大驿土",
      "xingYun": "墓",
      "ziZuo": "帝旺",
      "kongWang": false,
      "shenSha": {
        "将星": "酉"
      }
    },
    {
      "gan": "乙",
      "zhi": "丑",
      "ganWuXing": "木",
      "zhiWuXing": "土",
      "zhuXing": "日主",
      "cangGan": ["己", "癸", "辛"],
      "fuXing": ["正印", "偏印", "比肩"],
      "naYin": "海中金",
      "xingYun": "养",
      "ziZuo": "衰",
      "kongWang": true,
      "shenSha": {
        "华盖": "丑"
      }
    },
    {
      "gan": "癸",
      "zhi": "酉",
      "ganWuXing": "水",
      "zhiWuXing": "金",
      "zhuXing": "偏印",
      "cangGan": ["辛"],
      "fuXing": ["比肩"],
      "naYin": "剑锋金",
      "xingYun": "长生",
      "ziZuo": "临官",
      "kongWang": false,
      "shenSha": {
        "文昌贵人": "酉"
      }
    }
  ],
  "shiErChangSheng": {
    "年支": "死",
    "月支": "墓",
    "日支": "养",
    "时支": "长生"
  }
}
```

### 喜用神计算

```http
POST /api/xiyongshen
```

**请求参数**

| 字段名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| name | string | 是 | 姓名 |
| bazi | array | 是 | 八字四柱数组 |

**请求示例**

```json
{
  "name": "张三",
  "bazi": [
    // ... 同基础八字计算响应中的 bazi 数组
  ]
}
```

**响应示例**

```json
{
  "name": "张三",
  "riZhu": "乙",
  "riZhuStrength": "身弱",
  "wuXingScores": {
    "木": 2,
    "火": 1,
    "土": 3,
    "金": 2,
    "水": 1
  },
  "xiYongShen": "水、木",
  "logic": [
    "开始分析喜用神...",
    "1. 确定日主为: 乙, 五行属: 木",
    "2. 分析日主强弱: 日主身弱",
    "3. 日主身弱，需要寻找能够生扶日主五行的元素",
    "4. 生扶日主木五行的是: 水（印星）",
    "5. 比助日主木五行的是: 木（比劫）",
    "6. 综合分析，确定喜用神为: 水、木"
  ]
}
```

### 四柱八字综合分析

```http
POST /api/baziyuce
```

**请求参数**

| 字段名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| name | string | 是 | 姓名 |
| bazi | array | 是 | 八字四柱数组 |

**请求示例**

```json
{
  "name": "张三",
  "bazi": [
    // ... 同基础八字计算响应中的 bazi 数组
  ]
}
```

**响应示例**

```json
{
  "name": "张三",
  "steps": [
    {
      "title": "第一步：排盘与定盘",
      "content": [
        "做什么：将出生时间转换为天干地支表示的四柱八字，并确认其准确性。",
        "怎么做：",
        "1. 排四柱：使用专业排盘软件或权威万年历，输入公历出生年、月、日、时。",
        "..."
      ]
    },
    {
      "title": "第二步：定旺衰，识体性",
      "content": [
        "做什么：判断日主在整个八字中的能量状态（身强/身弱），这是选择喜用神的根本依据。",
        "您的日主为: 乙, 地支为: 丑",
        "怎么做：从四个维度综合评估（权重：得令 > 得地 > 得势/得助）。",
        "..."
      ]
    }
  ]
}
```

### 运势分析

```http
POST /api/fortune
Authorization: Bearer <token>
```

**请求参数**

| 字段名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| name | string | 是 | 姓名 |
| bazi | array | 是 | 八字四柱数组 |
| birthDate | string | 是 | 出生日期 |

**响应示例**

```json
{
  "name": "张三",
  "currentYear": "2023",
  "overallFortune": "整体运势平稳，需要注意身体健康。",
  "career": "事业发展稳定，但缺乏突破性进展。",
  "wealth": "财运一般，不宜进行高风险投资。",
  "health": "注意脾胃消化系统，避免暴饮暴食。",
  "relationship": "人际关系和谐，家庭和睦。",
  "advice": "保持现状，稳扎稳打，避免冒进。"
}
```

### 人生阶段分析

```http
POST /api/lifestages
Authorization: Bearer <token>
```

**请求参数**

| 字段名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| name | string | 是 | 姓名 |
| bazi | array | 是 | 八字四柱数组 |
| shiErChangSheng | object | 是 | 十二长生图 |

**响应示例**

```json
{
  "name": "张三",
  "analysis": {
    "childhood": "童年时期得到父母关爱，性格温和。",
    "youth": "青年时期学业顺利，但感情波折较多。",
    "middle": "中年时期事业发展平稳，需注意身体健康。",
    "later": "晚年生活安逸，子孙孝顺。"
  }
}
```

## 📊 工具接口

### 健康检查

```http
GET /health
```

**响应示例**

```json
{
  "status": "healthy",
  "service": "auspire-bazi"
}
```

## 📈 错误响应格式

所有错误响应遵循统一格式：

```json
{
  "error": "错误描述信息"
}
```

### 常见错误码

| HTTP状态码 | 错误类型 | 描述 |
|------------|----------|------|
| 400 | Bad Request | 请求参数错误 |
| 401 | Unauthorized | 未授权访问 |
| 404 | Not Found | 资源不存在 |
| 500 | Internal Server Error | 服务器内部错误 |

## 🔧 请求示例

### 使用 curl

```bash
# 基础八字计算
curl -X POST http://localhost:8080/api/bazi \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "birthDate": "1990-03-15",
    "birthTime": "14:30"
  }'

# 喜用神计算
curl -X POST http://localhost:8080/api/xiyongshen \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "bazi": [...]  // 完整的八字数组
  }'
```

### 使用 JavaScript (fetch)

```javascript
// 基础八字计算
async function calculateBazi() {
  const response = await fetch('/api/bazi', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      name: '张三',
      birthDate: '1990-03-15',
      birthTime: '14:30'
    })
  });
  
  const data = await response.json();
  console.log(data);
}
```

### 使用 Python (requests)

```python
import requests

# 基础八字计算
response = requests.post('http://localhost:8080/api/bazi', json={
    'name': '张三',
    'birthDate': '1990-03-15',
    'birthTime': '14:30'
});

data = response.json();
print(data);
```

## 🔄 版本历史

### v1.0.0 (2023-09-07)

**新增功能**
- 基础八字四柱计算
- 十神关系分析
- 藏干分析
- 纳音五行计算
- 十二长生分析
- 喜用神计算
- 四柱八字综合分析
- 用户注册登录系统
- 运势分析功能
- 人生阶段分析

**技术改进**
- 模块化服务架构
- RESTful API 设计
- JWT 认证机制
- Redis 会话管理
- 响应式前端界面

## 📞 技术支持

如有技术问题，请联系项目维护团队或查阅相关文档。

---
*本文档最后更新: 2023-09-07*