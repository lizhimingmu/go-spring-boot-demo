package pkg

import (
	"fmt"
	"reflect"

	"github.com/go-spring/go-spring/spring-boot"
	"github.com/go-spring/go-spring/spring-core"
)

func init() {
	SpringBoot.RegisterModule(func(ctx SpringCore.SpringContext) {
		bean := new(DemoStruct)
		fmt.Println(bean)

		t := reflect.TypeOf(bean)
		fmt.Println(t.String())

		et := t.Elem()
		fmt.Println(et.Name())
		fmt.Println(et.PkgPath())
		fmt.Println(et.String())

		v := reflect.ValueOf(bean)
		fmt.Println(v)
	})
}

type DemoStruct struct {
}

func (s *DemoStruct) String() string {
	return "d2.DemoStruct"
}
