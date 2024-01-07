package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// ExampleModel 示例模型，表示数据库中的一个表
type ExampleModel struct {
	gorm.Model
	Name string
}

func main() {
	// 数据库连接字符串：用户名:密码@tcp(地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接到数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v\n", err)
	}

	// 执行迁移
	if err := db.AutoMigrate(&ExampleModel{}); err != nil {
		fmt.Fprintf(os.Stderr, "数据库迁移失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("数据库迁移成功")
}
