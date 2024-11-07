package id

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

type IdGenerator struct {
	nodeId        int64
	snowflakeNode *snowflake.Node
}

func NewIdGenerator(machineNodeId int64) *IdGenerator {
	node := &IdGenerator{
		nodeId: machineNodeId,
	}

	// 雪花id
	snowflakeNode, err := snowflake.NewNode(machineNodeId)
	if err != nil {
		panic(err)
	}

	node.snowflakeNode = snowflakeNode
	return node
}

func (thiz *IdGenerator) UUID() string {
	return uuid.NewString()
}

func (thiz *IdGenerator) SnowflakeId() int64 {
	return thiz.snowflakeNode.Generate().Int64()
}
