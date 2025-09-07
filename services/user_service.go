package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"auspire/models"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewUserService(redisClient *redis.Client) *UserService {
	return &UserService{
		redisClient: redisClient,
		ctx:         context.Background(),
	}
}

func (s *UserService) generateID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (s *UserService) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (s *UserService) verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *UserService) Register(userReq *models.UserRegister) (*models.User, error) {
	// Check if username already exists
	exists, err := s.redisClient.HExists(s.ctx, "users:username", userReq.Username).Result()
	if err != nil {
		return nil, fmt.Errorf("检查用户名失败: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("用户名已存在")
	}

	// Check if email already exists
	exists, err = s.redisClient.HExists(s.ctx, "users:email", userReq.Email).Result()
	if err != nil {
		return nil, fmt.Errorf("检查邮箱失败: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("邮箱已存在")
	}

	// Generate user ID
	userID, err := s.generateID()
	if err != nil {
		return nil, fmt.Errorf("生成用户ID失败: %w", err)
	}

	// Hash password
	hashedPassword, err := s.hashPassword(userReq.Password)
	if err != nil {
		return nil, fmt.Errorf("密码加密失败: %w", err)
	}

	// Create user
	user := &models.User{
		ID:        userID,
		Username:  userReq.Username,
		Email:     userReq.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}

	// Store user data
	userData, err := user.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("序列化用户数据失败: %w", err)
	}

	pipe := s.redisClient.TxPipeline()
	
	// Store user by ID
	pipe.Set(s.ctx, fmt.Sprintf("user:%s", userID), userData, 0)
	
	// Store username -> userID mapping
	pipe.HSet(s.ctx, "users:username", userReq.Username, userID)
	
	// Store email -> userID mapping
	pipe.HSet(s.ctx, "users:email", userReq.Email, userID)
	
	_, err = pipe.Exec(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("保存用户数据失败: %w", err)
	}

	user.Password = "" // Don't return password
	return user, nil
}

func (s *UserService) Login(loginReq *models.UserLogin) (*models.User, error) {
	// Get user ID by username
	userID, err := s.redisClient.HGet(s.ctx, "users:username", loginReq.Username).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("用户名或密码错误")
	} else if err != nil {
		return nil, fmt.Errorf("查找用户失败: %w", err)
	}

	// Get user data
	userData, err := s.redisClient.Get(s.ctx, fmt.Sprintf("user:%s", userID)).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("用户不存在")
	} else if err != nil {
		return nil, fmt.Errorf("获取用户数据失败: %w", err)
	}

	// Parse user data
	user, err := models.UserFromJSON([]byte(userData))
	if err != nil {
		return nil, fmt.Errorf("解析用户数据失败: %w", err)
	}

	// Verify password
	if err := s.verifyPassword(user.Password, loginReq.Password); err != nil {
		return nil, fmt.Errorf("用户名或密码错误")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, fmt.Errorf("用户账户已被禁用")
	}

	user.Password = "" // Don't return password
	return user, nil
}

func (s *UserService) GetUserByID(userID string) (*models.User, error) {
	userData, err := s.redisClient.Get(s.ctx, fmt.Sprintf("user:%s", userID)).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("用户不存在")
	} else if err != nil {
		return nil, fmt.Errorf("获取用户数据失败: %w", err)
	}

	user, err := models.UserFromJSON([]byte(userData))
	if err != nil {
		return nil, fmt.Errorf("解析用户数据失败: %w", err)
	}

	user.Password = "" // Don't return password
	return user, nil
}