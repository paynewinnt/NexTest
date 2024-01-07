package testservice

import (
	"github.com/paynewinnt/NexTest/model"
	"gorm.io/gorm"
	// 导入其他需要的包
)

type APITestService struct {
	Db *gorm.DB // 数据库连接
}

func NewAPITestService(db *gorm.DB) *APITestService {
	return &APITestService{
		Db: db,
	}
}

// CreateAPITest 创建一个新的API测试
func (s *APITestService) CreateAPITest(test *model.APITest) error {
	if err := test.Validate(); err != nil {
		return err
	}
	return s.Db.Create(test).Error
}

// GetAPITest 根据ID获取API测试
func (s *APITestService) GetAPITest(id uint) (*model.APITest, error) {
	var test model.APITest
	result := s.Db.First(&test, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &test, nil
}

// UpdateAPITest 更新一个API测试
func (s *APITestService) UpdateAPITest(test *model.APITest) error {
	if err := test.Validate(); err != nil {
		return err
	}
	return s.Db.Save(test).Error
}

// DeleteAPITest 删除一个API测试
func (s *APITestService) DeleteAPITest(id uint) error {
	return s.Db.Delete(&model.APITest{}, id).Error
}

// ListAPITests 列出所有API测试
func (s *APITestService) ListAPITests() ([]model.APITest, error) {
	var tests []model.APITest
	result := s.Db.Find(&tests)
	if result.Error != nil {
		return nil, result.Error
	}
	return tests, nil
}
