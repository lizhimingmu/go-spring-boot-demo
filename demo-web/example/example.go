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

	"github.com/go-spring/go-spring/spring-boot"
	"github.com/go-spring/go-spring/spring-core"
	"github.com/go-spring/go-spring/spring-web"
)

func init() {
	SpringBoot.RegisterModule(func(ctx SpringCore.SpringContext) {
		ctx.RegisterBean(new(Controller))
	})
}

type Controller struct {
}

func (c *Controller) InitWebBean(wc SpringWeb.WebContainer, ctx SpringCore.SpringContext) {

	var f1 SpringWeb.Filter
	ctx.GetBeanByName("f1", &f1)

	var f2 SpringWeb.Filter
	ctx.GetBeanByName("f2", &f2)

	wc.GET("/", c.Home, f2)
	wc.GET("/f1f2", c.F1F2, f1, f2)
	wc.GET("/f2f1", c.F2F1, f2, f1)
}

func (c *Controller) Home(ctx SpringWeb.WebContext) {
	ctx.String(http.StatusOK, "OK!")
}

func (c *Controller) F1F2(ctx SpringWeb.WebContext) {
	ctx.String(http.StatusOK, "f1f2!")
}

func (c *Controller) F2F1(ctx SpringWeb.WebContext) {
	ctx.String(http.StatusOK, "f2f1!")
}
