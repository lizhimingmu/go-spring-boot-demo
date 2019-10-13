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

	"github.com/go-spring/demo-web/filter"
	"github.com/go-spring/go-spring-boot/spring-boot"
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

func (controller *Controller) InitWebBean(c SpringWeb.WebContainer, ctx SpringCore.SpringContext) {

	f1, _ := ctx.FindBeanByName("f1").(*filter.NumberFilter)
	f2, _ := ctx.FindBeanByName("f2").(*filter.NumberFilter)

	c.GET("/", controller.Home, f2)
	c.GET("/f1f2", controller.F1F2, f1, f2)
	c.GET("/f2f1", controller.F2F1, f2, f1)
}

func (controller *Controller) Home(ctx SpringWeb.WebContext) {
	ctx.String(http.StatusOK, "OK!")
}

func (controller *Controller) F1F2(ctx SpringWeb.WebContext) {
	ctx.String(http.StatusOK, "f1f2!")
}

func (controller *Controller) F2F1(ctx SpringWeb.WebContext) {
	ctx.String(http.StatusOK, "f2f1!")
}
