package graph

import "xorm.io/xorm"

type Resolver struct {
	// xormエンジンのインスタンスをフィールドとして保持
	DB *xorm.Engine
}
