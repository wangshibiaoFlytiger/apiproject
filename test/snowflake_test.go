package test

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"testing"
	"time"
)

/**
测试snowflake雪花算法
*/
func TestSnowflake(t *testing.T) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake ID.
	id := node.Generate()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

	// Generate and print, all in one.
	fmt.Printf("ID       : %d\n", node.Generate().Int64())

	//可以看出每次生成的id是大小递增的,且不连续
	for i := 1; true; i++ {
		fmt.Printf("ID       : %d\n", node.Generate())
		time.Sleep(1000000)
	}
}
