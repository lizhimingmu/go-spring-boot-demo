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

package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-spring/go-spring-boot-demo/filter"
	"github.com/go-spring/go-spring-parent/spring-logger"
	"github.com/go-spring/go-spring-parent/spring-utils"
	"github.com/go-spring/go-spring-web/spring-echo"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {

	// 使用 "router" 名称的过滤器，需要使用 SpringBoot.FilterBean 封装，为了编译器能够进行类型检查
	r := SpringBoot.Route("/api", SpringBoot.FilterBean(
		"router", (*filter.SingleBeanFilter)(nil)),
	)

	// 接受简单函数，可以使用 SpringBoot.Filter 封装，进而增加可用条件
	r.GetMapping("/func", func(ctx SpringWeb.WebContext) {
		ctx.String(http.StatusOK, "func() return ok")
	}, SpringBoot.Filter(SpringEcho.Filter(middleware.KeyAuth(
		func(key string, context echo.Context) (bool, error) {
			return key == "key_auth", nil
		}))).ConditionOnPropertyValue("key_auth", true),
	)

	// 接受类型方法，也可以不使用 SpringBoot.Filter 封装
	r.GET("/method", (*MyController).Method, SpringEcho.Filter(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(echoCtx echo.Context) error {
				webCtx := SpringEcho.WebContext(echoCtx)
				webCtx.LogInfo("call method")
				return nil
			}
		}))

	SpringBoot.RegisterBean(new(MyController)).Init(func(c *MyController) {

		// 接受对象方法
		r.GetMapping("/ok", c.OK, SpringBoot.FilterBean("router//ok")).
			ConditionOnProfile("test").
			Swagger(). // 设置接口的信息
			WithDescription("ok")

		// 该接口不会注册，因为没有匹配的端口
		r.GetMapping("/nil", c.OK).OnPorts(9999)

		// 注意这个接口不和任何 Router 绑定
		SpringBoot.GetBinding("/echo", c.Echo, SpringBoot.FilterBean("router//echo")).
			Swagger(). // 设置接口的信息
			WithDescription("echo")
	})
}

type MyController struct {
	RedisClient redis.Cmdable                 `autowire:""`
	DB          *gorm.DB                      `autowire:""`
	AppCtx      SpringBoot.ApplicationContext `autowire:""`
}

type EchoRequest struct {
	Str string `query:"str"`
}

type EchoResponse struct {
	Echo string `json:"echo"`
}

func (c *MyController) Echo(request EchoRequest) *EchoResponse {
	return &EchoResponse{"echo " + request.Str}
}

func (c *MyController) Method(ctx SpringWeb.WebContext) {
	ctx.String(http.StatusOK, "method() return ok")
}

func (c *MyController) OK(ctx SpringWeb.WebContext) {

	err := c.RedisClient.Set("key", "ok", time.Second*10).Err()
	SpringUtils.Panic(err).When(err != nil)

	val, err := c.RedisClient.Get("key").Result()
	SpringUtils.Panic(err).When(err != nil)

	rows, err := c.DB.Table("ENGINES").Select("ENGINE").Rows()
	SpringUtils.Panic(err).When(err != nil)
	defer func() { _ = rows.Close() }()

	count := 0

	for rows.Next() {
		count++

		var engine string
		_ = rows.Scan(&engine)
		SpringLogger.Info(engine)

		if engine != "sql-mock" {
			panic(errors.New("error"))
		}
	}

	if count != 1 {
		panic(errors.New("error"))
	}

	ctx.JSONBlob(200, []byte(val))
}
