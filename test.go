package main

import (
	"fmt"

	_ "gorm.io/driver/sqlite"
)

func TestFunc() {
	fmt.Println("TestFunc")
}
