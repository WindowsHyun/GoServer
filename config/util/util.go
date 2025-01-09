package util

import (
	"GoServer/config/define"
	"context"
	"strings"

	"github.com/gin-gonic/gin"
)

func RealIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		var realIP string
		if cfIP := c.GetHeader("CF-Connecting-IP"); cfIP != "" {
			realIP = cfIP
		} else if cfIP := c.GetHeader("True-Client-IP"); cfIP != "" {
			realIP = cfIP
		} else if forwardedFor := c.GetHeader("X-Forwarded-For"); forwardedFor != "" {
			ips := strings.Split(forwardedFor, ",")
			realIP = strings.TrimSpace(ips[0])
		} else if xRealIP := c.GetHeader("X-Real-IP"); xRealIP != "" {
			realIP = xRealIP
		} else {
			realIP = c.ClientIP()
		}
		newContext := context.WithValue(c.Request.Context(), define.ContextUserRealIP, realIP)
		c.Request = c.Request.WithContext(newContext)
		c.Next()
	}
}
