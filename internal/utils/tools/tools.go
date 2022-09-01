package tools

import (
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

// IsContain 检查字符串是否在slice
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// GetClientIp 获取真实客户端IP
func GetClientIp(r *http.Request) string {
	// 尝试从 X-Forwarded-For 中获取
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	// 如果 X-Forwarded-For 有值，取第一个非unknown的ip
	clientIp := ""
	ipArr := strings.Split(xForwardedFor, ",")
	for _, ip := range ipArr {
		ip = strings.TrimSpace(ip)
		if ip != "" && strings.ToLower(ip) != "unknown" {
			clientIp = ip
			break
		}
	}
	if clientIp == "" {
		// 尝试从 X-Real-Ip 中获取
		clientIp = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
		if clientIp == "" {
			// 直接从 Remote Addr 中获取
			remoteIp, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
			if err != nil {
				clientIp = ""
			} else {
				clientIp = remoteIp
			}
		}
	}
	return clientIp
}

// GetOutBoundIp 获取本机出口IP
func GetOutBoundIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// GetRandomString 随机生成指定长度的字符串
// flag 0字母或数字 1 纯数字
func GetRandomString(length int, flag int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if flag == 1 {
		letters = []byte("0123456789")
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
