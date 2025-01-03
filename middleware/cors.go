package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// CORSConfig 用於配置跨域的設置
func CORSConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://your-frontend-domain.com"}, // 允許的來源域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                  // 允許的 HTTP 方法
		AllowHeaders:     []string{"Content-Type", "Authorization"},                           // 允許的自定義 Header
		ExposeHeaders:    []string{"Content-Length"},                                          // 客戶端可以訪問的 Header
		AllowCredentials: true,                                                                // 是否允許攜帶憑證（如 Cookie）
		MaxAge:           12 * time.Hour,                                                      // 預檢請求的緩存時間
	})
}
