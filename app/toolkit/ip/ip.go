package ip

import (
	"net"
	"net/http"
	"strings"
)

// GetClientIP 提取客户端真实 IP（兼容 IPv4 / IPv6）
func GetClientIP(r *http.Request) string {
	// 1. 优先使用 X-Forwarded-For
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// 多个 IP 情况，如 "192.168.0.1, 10.0.0.1"
		ips := strings.Split(ip, ",")
		if len(ips) > 0 {
			ip = strings.TrimSpace(ips[0])
		}
	}

	// 2. 其次尝试 X-Real-IP
	if ip == "" {
		ip = strings.TrimSpace(r.Header.Get("X-Real-IP"))
	}

	// 3. fallback 到 RemoteAddr（自动处理 IPv6 地址）
	if ip == "" {
		host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
		if err == nil {
			ip = host
		}
	}

	// 4. 最后检查合法性
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return ""
	}

	return parsedIP.String()
}
