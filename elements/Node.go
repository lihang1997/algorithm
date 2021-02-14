package elements

import "go/types"

// 双链表使用
const NEXT = 0
const PREV = 1

// 二叉树使用
const LEFT = 0
const RIGHT = 1

// 节点数据类型
type Node struct {
	Witch *types.Pointer
	Children *[]Node
}

// 构建一个节点结构体
func (ctx *Node) NewNode(witch *types.Pointer, children *[]Node){
	ctx.Witch = witch
	ctx.Children = children
}
