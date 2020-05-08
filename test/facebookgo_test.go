package test

import (
	"fmt"
	"github.com/facebookgo/inject"
	"testing"
)

type C struct {
	B *B `inject:"" json:"b"`
}
type B struct {
	A *A `inject:"" json:"a"` //这里会根据注入对象的Name字段有选择的进行注入
}
type A struct {
	Name string `json:"name"`
}

// 通过injectGraph.Objects()可以获取所有设置了Name的待填充对象
func GetObject(name string) interface{} {
	for _, o := range injectGraph.Objects() {
		if o.Name == name {
			return o.Value
		}
	}
	return nil
}

var injectGraph inject.Graph

/**
测试依赖注入框架facebookgo
*/
func TestFacebookgo(t *testing.T) {
	a := A{
		Name: "hello",
	}
	a2 := A{
		Name: "hello2",
	}
	c := C{}
	err := injectGraph.Provide( //对象提供者
		&inject.Object{Value: &a},
		&inject.Object{Name: "这里可以给对象一个自定义命名", Value: &a2},
		&inject.Object{Value: &c}, //这个也需要
	)
	if err != nil {
	}
	err = injectGraph.Populate() //填充对象到使用了inject标签的结构体字段中
	if err != nil {
	}
	fmt.Println(c.B.A.Name)
}
