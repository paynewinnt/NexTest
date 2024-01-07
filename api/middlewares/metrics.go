package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

// 定义Prometheus指标
var (
	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of get requests.",
		},
		[]string{"path"},
	)
	responseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "response_status",
			Help: "Status of HTTP response",
		},
		[]string{"status"},
	)
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_time_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
		},
		[]string{"path"},
	)
)

// 初始化
func init() {
	// 注册Prometheus指标
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(responseStatus)
	prometheus.MustRegister(httpDuration)
}

// MetricsMiddleware 创建一个记录性能指标的中间件
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)
		path := c.Request.URL.Path
		status := c.Writer.Status()

		// 更新Prometheus指标
		totalRequests.WithLabelValues(path).Inc()
		responseStatus.WithLabelValues(string(status)).Inc()
		httpDuration.WithLabelValues(path).Observe(duration.Seconds())
	}
}
