package util

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/bwmarrin/snowflake"
	uuid "github.com/satori/go.uuid"
	"time"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
生成分布式唯一ID: 整数,自增,且不连续
*/
func GenUniqueId() snowflake.ID {
	// Create a new Node with a Node number of 10000
	var nodeNumber int64 = 900
	node, err := snowflake.NewNode(nodeNumber + int64(randomdata.Number(99)))
	if err != nil {
		panic(err)
	}

	//等待1毫秒,避免发生ID重复的问题
	time.Sleep(1 * time.Millisecond)
	// Generate a snowflake ID.
	id := node.Generate()

	return id
}

/**
生成uuid
*/
func GenUUID() string {
	return uuid.NewV4().String()
}
