package config

// 应用程序相关常量
const (
	AppName         = "NexTest"
	AppVersion      = "1.0.0"
	DefaultHTTPPort = "8080"
)

// 数据库相关常量
const (
	DBDriverName       = "mysql"
	DBConnectionString = "username:password@tcp(localhost:3306)/dbname?parseTime=true"
)

// API 路径相关常量
const (
	APIBasePath      = "/api/v1"
	APIHealthCheck   = APIBasePath + "/health"
	APIUserRoute     = APIBasePath + "/users"
	APITestCaseRoute = APIBasePath + "/testcases"
	APITestRunRoute  = APIBasePath + "/testruns"
)

// 认证和授权相关常量
const (
	JWTSecretKey = "your_jwt_secret_key" // 应从安全的配置源获取
	JWTTTL       = 3600                  // JWT 有效期（秒）
)

// 日志和监控相关常量
const (
	LogFilePath      = "/var/log/nextest.log"
	MonitoringSystem = "prometheus"
)

// DefaultTestTimeout 测试执行相关常量
const (
	DefaultTestTimeout = 120 // 测试默认超时时间（秒）
)
const (
	SessionSecret = "your_session_secret" // 应从安全的配置源获取
)

// 其他可能的常量
// ...

// 注意：对于涉及安全或敏感信息的常量（如JWTSecretKey），应从环境变量或安全的配置源动态获取，而不是硬编码在代码中。
