package model

import (
	"errors"
	"gorm.io/gorm"
	"strings"
)

// APITest 表示一个API测试的数据模型
type APITest struct {
	gorm.Model
	Name             string `gorm:"size:255;not null"` // 测试的名称
	Description      string `gorm:"size:500"`          // 测试的描述
	Endpoint         string `gorm:"not null"`          // 测试的API端点
	Method           string `gorm:"size:10;not null"`  // HTTP方法 (GET, POST, PUT, DELETE等)
	Payload          string `gorm:"type:text"`         // 请求的Payload（如果适用）
	ExpectedResponse string `gorm:"type:text"`         // 预期的响应
	// 其他字段根据需求添加
}

// Validate 检查APITest数据的有效性
func (a *APITest) Validate() error {
	if strings.TrimSpace(a.Name) == "" {
		return errors.New("测试名称不能为空")
	}
	if strings.TrimSpace(a.Endpoint) == "" {
		return errors.New("API端点不能为空")
	}
	if !isValidMethod(a.Method) {
		return errors.New("无效的HTTP方法")
	}
	// 可以添加更多的验证逻辑
	return nil
}

// isValidMethod 验证HTTP方法是否有效
func isValidMethod(method string) bool {
	validMethods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	for _, m := range validMethods {
		if strings.ToUpper(method) == m {
			return true
		}
	}
	return false
}

// PrepareForTesting 准备测试数据
func (a *APITest) PrepareForTesting() {
	a.Method = strings.ToUpper(strings.TrimSpace(a.Method))
	// 这里可以添加更多准备测试数据的逻辑
}
