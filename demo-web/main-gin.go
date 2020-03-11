/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-spring/demo-web/example"
	_ "github.com/go-spring/demo-web/filter"
	"github.com/go-spring/go-spring-parent/spring-logger"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
	_ "github.com/go-spring/go-spring/starter-gin"
)

// 注册 gin 容器
func registerGinContainer() {
	gin.SetMode(gin.ReleaseMode)

	e := gin.Default()

	e.Use(func(ginCtx *gin.Context) {
		fmt.Println("use registerGinContainer()")
		ginCtx.Next()
	})

	c := SpringGin.NewContainer()
	c.SetPort(8080)

	webServer := SpringWeb.NewWebServer()
	webServer.AddWebContainer(c)

	SpringBoot.RegisterBean(webServer)
}

func main() {
	SpringLogger.SetLogger(&SpringLogger.Console{})
	if false {
		registerGinContainer()
	}
	SpringBoot.RunApplication("config/")
}
