package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
测试UUID
*/
func TestUUID(t *testing.T) {
	fmt.Println(uuid.NewV1().String())
	fmt.Println(uuid.NewV4().String())
}
