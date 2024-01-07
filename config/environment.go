package config

import (
	"log"
	"os"
	"strconv"
)

// GetEnv 用于从环境变量获取字符串值，如果未设置则返回默认值
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetEnvAsInt 读取一个环境变量作为整数，如果未设置则返回默认值
func GetEnvAsInt(key string, defaultValue int) int {
	valueStr := GetEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("警告: 环境变量 %s 的值 '%s' 不是一个有效的整数\n", key, valueStr)
		return defaultValue
	}
	return value
}

// GetEnvAsBool 读取一个环境变量作为布尔值，如果未设置则返回默认值
func GetEnvAsBool(key string, defaultValue bool) bool {
	valueStr := GetEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		log.Printf("警告: 环境变量 %s 的值 '%s' 不是一个有效的布尔值\n", key, valueStr)
		return defaultValue
	}
	return value
}

// GetEnvAsFloat64 读取一个环境变量作为float64，如果未设置则返回默认值
func GetEnvAsFloat64(key string, defaultValue float64) float64 {
	valueStr := GetEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Printf("警告: 环境变量 %s 的值 '%s' 不是一个有效的float64\n", key, valueStr)
		return defaultValue
	}
	return value
}

// 可以根据需要添加更多的辅助函数来读取特定类型的环境变量。
