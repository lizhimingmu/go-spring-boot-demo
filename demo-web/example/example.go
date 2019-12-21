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

	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(Controller)).InitFunc(func(c *Controller) {
		SpringBoot.GetMapping("/", c.Home).SetFilterNames("f2")
		SpringBoot.GetMapping("/f1f2", c.F1F2).SetFilterNames("f1", "f2")
		SpringBoot.GetMapping("/f2f1", c.F2F1).SetFilterNames("f2", "f1")
	})
}

type Controller struct {
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
