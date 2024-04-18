package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 定义一个模型
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// 连接SQLite数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移schema
	db.AutoMigrate(&Product{})

	TestFunc()

	// 创建
	p := &Product{Code: "D42", Price: 100}
	if db.Create(p).Error != nil {
		fmt.Println("Create failed")
		return
	}

	fmt.Println("Create success", p.ID)

	writer := tabwriter.NewWriter(os.Stdout, 5, 2, 1, 'x', tabwriter.TabIndent)
	if _, err := fmt.Fprintln(writer, strings.Join([]string{"A", "B", "C", "D"}, "\t")); err != nil {
		fmt.Println("write table headers failed")
		return
	}

	if _, err := fmt.Fprintln(writer, strings.Join([]string{"100", "2", "3", "400000"}, "\t")); err != nil {
		fmt.Println("write table headers failed")
		return
	}

	writer.Flush()

}
