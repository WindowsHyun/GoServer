package middleware

import (
	"GoServer/config"
	"GoServer/config/define"
	"GoServer/config/structure"
	IRedis "GoServer/database/redis"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	contextKey   = "session"
	session      IRedis.RedisInterface
	appSecretKey string
)

func Init(config *config.Config, redis map[string]IRedis.RedisInterface) error {
	if config == nil {
		return fmt.Errorf("config is nil")
	}
	session = redis[define.RedisSessionDB]
	if session == nil {
		return fmt.Errorf("session redis is nil")
	}
	appSecretKey = config.GetSecretKey().APPSecretKey
	if appSecretKey == "" {
		return fmt.Errorf("app secret key is empty")
	}

	return nil
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims := &structure.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(appSecretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check session in Redis
		sessionKey := fmt.Sprintf("session:%s", claims.Id)
		var sessionValue string
		err = session.Get(c.Request.Context(), sessionKey, &sessionValue)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session not found"})
			c.Abort()
			return
		}

		c.Set(contextKey, claims)
		c.Next()
	}
}

func GenerateToken(ctx context.Context, userId, email string) (string, error) {
	expTime := time.Now().Add(time.Hour * 24)
	claims := structure.Claims{
		Id:    userId,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(appSecretKey))
	if err != nil {
		return "", err
	}

	// Store session in Redis
	sessionKey := fmt.Sprintf("session:%s", userId)
	err = session.Set(ctx, sessionKey, tokenString, time.Hour*24)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Logout(ctx context.Context, tokenString string) error {
	sessionKey := fmt.Sprintf("session:%s", tokenString)
	return session.Del(ctx, sessionKey)
}

func GetClaims(c *gin.Context) *structure.Claims {
	value := c.Value(contextKey)
	if value == nil {
		return nil
	}

	claims, ok := value.(*structure.Claims)
	if !ok {
		return nil
	}

	return claims
}
