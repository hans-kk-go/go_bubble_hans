package snowflake

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func GetSnowflakeId() int64 {

	//将包导入到项目中，然后使用 唯一的节点号。默认设置允许节点号范围从 0 到 1023。
	//如果您设置了自定义节点位值，则需要计算您的 节点编号范围将是。使用节点对象调用 Generate（） 方法来 生成并返回唯一的雪花 ID。
	//请记住，您创建的每个节点都必须具有唯一的节点编号，即使 跨多个服务器。
	//如果不保持节点号唯一，生成器 无法保证所有节点的唯一 ID。

	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		//return
	}

	// Generate a snowflake ID.
	id := node.Generate()
	return id.Int64()

	//// Print out the ID in a few different ways.
	//fmt.Printf("Int64  ID: %d\n", id)
	//fmt.Printf("String ID: %s\n", id)
	//fmt.Printf("Base2  ID: %s\n", id.Base2())
	//fmt.Printf("Base64 ID: %s\n", id.Base64())
	//
	//// Print out the ID's timestamp
	//fmt.Printf("ID Time  : %d\n", id.Time())
	//
	//// Print out the ID's node number
	//fmt.Printf("ID Node  : %d\n", id.Node())
	//
	//// Print out the ID's sequence number
	//fmt.Printf("ID Step  : %d\n", id.Step())
	//
	//// Generate and print, all in one.
	//fmt.Printf("ID       : %d\n", node.Generate().Int64())
}
