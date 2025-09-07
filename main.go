package main

import (
	"auspire/handlers"
	"auspire/middleware"
	"auspire/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Redis client setup
	redisClient := redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       0,
	})

	// JWT configuration
	jwtSecret := getEnv("JWT_SECRET", "auspire-secret-key-2024")
	jwtIssuer := getEnv("JWT_ISSUER", "auspire")

	// Initialize handlers
	baziHandler := handlers.NewBaziHandler()
	authHandler := handlers.NewAuthHandler(redisClient, jwtSecret, jwtIssuer)

	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html")
	r.GET("/paipan", baziHandler.ServePaiPanPage)

	api := r.Group("/api")
	{
		// Authentication routes (no middleware required)
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		// Public routes (no authentication required)
		api.POST("/bazi", baziHandler.CalculateBazi)
		api.POST("/xiyongshen", baziHandler.CalculateXiYongShen)
		api.POST("/baziyuce", baziHandler.AnalyzeBaziyuce)

		// Protected routes (authentication required)
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(services.NewJWTService(jwtSecret, jwtIssuer)))
		{
			protected.GET("/profile", authHandler.GetProfile)
			protected.POST("/fortune", baziHandler.AnalyzeFortune)
			protected.POST("/lifestages", baziHandler.AnalyzeLifeStages)
		}
	}

	r.GET("/health", baziHandler.Health)

	port := ":8080"
	log.Printf("Auspire服务器启动在端口 %s", port)
	log.Println("访问 http://localhost:8080")
	log.Println("API端点:")
	log.Println("  注册: POST http://localhost:8080/api/register")
	log.Println("  登录: POST http://localhost:8080/api/login")
	log.Println("  个人资料: GET http://localhost:8080/api/profile")
	log.Println("  八字计算: POST http://localhost:8080/api/bazi")
	log.Println("  喜用神计算: POST http://localhost:8080/api/xiyongshen")
	log.Println("  八字综合分析: POST http://localhost:8080/api/baziyuce")
	log.Println("  运势分析: POST http://localhost:8080/api/fortune")
	log.Println("  人生阶段分析: POST http://localhost:8080/api/lifestages")
	log.Println("健康检查: http://localhost:8080/health")

	if err := r.Run(port); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}