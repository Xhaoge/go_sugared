package cmd

import (
	"fmt"
)

func Init() {
	fmt.Println("this is cmd package init")
}

func registerApiHandlers() error {
	return nil
}


gin.Default()

// 具体的default 方法
func Default() *Engine {
	debugPrintWARNINGDefault()

	// 创建一个gin 框架实例
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

