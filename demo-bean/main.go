package main

import (
	_ "github.com/go-spring/demo-bean/pkg"
	"github.com/go-spring/go-spring-parent/spring-logger"
	"github.com/go-spring/go-spring/spring-boot"
)

func main() {
	SpringLogger.SetLogger(&SpringLogger.Console{})
	SpringBoot.RunApplication("config/")
}
