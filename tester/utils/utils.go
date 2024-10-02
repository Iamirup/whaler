package utils

import (
	"math/rand/v2"
	"net"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetClientIP(request *gin.Context) string {
	ip := request.GetHeader("X-Forwarded-For")
	if ip != "" {
		// first X-Forwarded-For
		ips := strings.Split(ip, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}
	// Check the X-Real-IP header
	ip = request.GetHeader("X-Real-IP")
	if ip != "" {
		return ip
	}
	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(request.Request.RemoteAddr)
	if err != nil {
		return request.Request.RemoteAddr
	}
	return ip
}
func GetExpiryTime() time.Time {
	const expTime int8 = 60
	return time.Now().Add(time.Second * time.Duration(expTime))

}
func SmsTokenGenerate() int {
	return rand.IntN(8999) + 1000
}
