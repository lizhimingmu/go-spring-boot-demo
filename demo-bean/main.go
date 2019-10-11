package main

import (
	_ "github.com/go-spring/demo-bean/pkg"
	"github.com/go-spring/go-spring-boot/spring-boot"
)

func main() {
	SpringBoot.RunApplication("config/")
}
