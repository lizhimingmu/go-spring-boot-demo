package pkg

import (
	"fmt"
	"reflect"
)

func init() {
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
}

type DemoStruct struct {
}

func (s *DemoStruct) String() string {
	return "d1.DemoStruct"
}
