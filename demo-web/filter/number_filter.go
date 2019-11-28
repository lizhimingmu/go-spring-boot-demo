package filter

import (
	"fmt"

	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterNameBean("f1", NewNumberFilter(1))
	SpringBoot.RegisterNameBean("f2", NewNumberFilter(2))
}

type NumberFilter struct {
	n int
}

func NewNumberFilter(n int) *NumberFilter {
	return &NumberFilter{
		n: n,
	}
}

func (f *NumberFilter) Invoke(ctx SpringWeb.WebContext, chain *SpringWeb.FilterChain) {
	defer fmt.Println("::after", f.n)
	fmt.Println("::before", f.n)
	chain.Next(ctx)
}
