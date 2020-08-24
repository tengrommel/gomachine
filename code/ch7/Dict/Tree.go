package main

import (
	"fmt"
	"sync"
)

type KeyWordTreeNode struct {
	// 1,2,3
	KeyWordIDs map[int64]bool
	// 百，度，一，下
	Char string
	// 父亲节点
	ParentKeyWordTreeNode *KeyWordTreeNode
	// 子节点集合
	SubKeyWordTreeNodes map[string]*KeyWordTreeNode
}

// 初始化节点
func NewKeyWordTreeNode() *KeyWordTreeNode {
	return &KeyWordTreeNode{
		make(map[int64]bool, 0),
		"",
		nil,
		make(map[string]*KeyWordTreeNode, 0),
	}
}

// 初始化节点设置内容以及父亲节点
func NewKeyWordTreeNodeWithParams(ch string, parent *KeyWordTreeNode) *KeyWordTreeNode {
	return &KeyWordTreeNode{
		make(map[int64]bool, 0),
		ch,
		parent,
		make(map[string]*KeyWordTreeNode, 0),
	}
}

type KeyWordTree struct {
	root        *KeyWordTreeNode // 根节点
	kv          KeyWordKV        // 映射关系
	charBeginKV CharBeginKV      // 开始映射
	rw          *sync.RWMutex    // 线程安全
}

// 开启字典树
func NewKeyWordTree() *KeyWordTree {
	return &KeyWordTree{
		NewKeyWordTreeNode(),
		KeyWordKV{},
		CharBeginKV{},
		new(sync.RWMutex),
	}
}

func (sTree *KeyWordTree) DebugOut() {
	fmt.Println("s.kv", sTree.kv) // 输出节点
	tempRoot := sTree.root
	dfs(tempRoot)
}

// 遍历字典树
func dfs(root *KeyWordTreeNode) {
	if root == nil {
		return
	} else {
		fmt.Println("s.root=", root.Char)
		fmt.Println("s.KeywordIds=", root.KeyWordIDs)
		for _, v := range root.SubKeyWordTreeNodes {
			dfs(v)
		}
	}
}

func (sTree *KeyWordTree) Put(id int64, keyWord string) {

}

// 百度搜索提示，返回字符串，limit限制深度
func (sTree *KeyWordTree) Sugg(keyWord string, limit int) []string {
	return make([]string, 0)
}
