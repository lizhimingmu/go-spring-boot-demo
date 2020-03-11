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

package example

import (
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(RedisController)).Init(func(c *RedisController) {
		SpringBoot.GetMapping("/get", c.Get)
		SpringBoot.PostMapping("/set", c.Set)
	})
}

type RedisController struct {
	RedisClient redis.Cmdable `autowire:""`
}

func (c *RedisController) Get(ctx SpringWeb.WebContext) {
	key := ctx.FormValue("key")
	val := c.RedisClient.Get(key)
	if err := val.Err(); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, val.Val())
}

func (c *RedisController) Set(ctx SpringWeb.WebContext) {
	key := ctx.FormValue("key")
	val := ctx.FormValue("val")
	c.RedisClient.Set(key, val, time.Second*5)
	ctx.NoContent(http.StatusOK)
}
