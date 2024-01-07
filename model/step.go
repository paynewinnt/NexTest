package model

import (
	"gorm.io/gorm"
	"time"
)

// Step 定义了一个测试步骤或工作流程步骤的基本属性，如步骤名称、描述、执行顺序、所属场景或工作流程的ID、持续时间和当前状态。这个结构体旨在被用于场景和工作流程中的每个步骤的定义。
type Step struct {
	gorm.Model
	Name        string        `json:"name" gorm:"size:255"`
	Description string        `json:"description" gorm:"size:500"`
	Order       int           `json:"order"`                  // Execution order
	ScenarioID  uint          `json:"scenario_id"`            // Associated Scenario ID
	WorkflowID  uint          `json:"workflow_id"`            // Associated Workflow ID
	Duration    time.Duration `json:"duration"`               // Estimated or actual duration
	Status      string        `json:"status" gorm:"size:100"` // Current status of the step
}
