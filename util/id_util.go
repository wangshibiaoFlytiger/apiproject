package util

import "github.com/bwmarrin/snowflake"

/**
生成分布式唯一ID: 整数,自增,且不连续
*/
func GenUniqueId() snowflake.ID {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	// Generate a snowflake ID.
	id := node.Generate()

	return id
}
