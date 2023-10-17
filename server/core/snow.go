package core

import (
	"github.com/bwmarrin/snowflake"
	"github.com/sreio/s-url/define"
)

func InitSnow() {
	node, err := snowflake.NewNode(define.Config.Server.Node) // 传入节点ID
	if err != nil {
		panic(err)
	}
	define.Snow = node
}

// GenId 生成ID
func GenId() int64 {
	return define.Snow.Generate().Int64()
}

// GenIdStr 生成ID
func GenIdStr() string {
	return define.Snow.Generate().String()
}

// GenIdBase62 生成Base62
func GenIdBase62() string {
	return base62Encode(GenId())
}

// ParseBase62 解析Base62
func ParseBase62(base62 string) int64 {
	return base62Decode(base62)
}
