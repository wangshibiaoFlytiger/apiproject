package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

/**
测试UUID
*/
func TestUUID(t *testing.T) {
	fmt.Println(uuid.NewV1().String())
	fmt.Println(uuid.NewV4().String())
}
