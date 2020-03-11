// Trie 树为框架提供动态路由
// /:lang  /*filepath
package frame

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由  /p/:lang, 只有叶子节点中才有这个
	part     string  // 路由中的部分 :lang
	childern []*node // 子节点
	isWild   bool    // 是否精确匹配
}

// 返回第一个成功匹配的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.childern {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.childern {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 插入 Trie 树节点， 用于注册服务路径
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		//只在最后一层节点中，设置 pattern。
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)

	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.childern = append(n.childern, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	childern := n.matchChildren(part)

	for _, child := range childern {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil

}

func Output(n *node, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	fmt.Print("|--")
	fmt.Println(n.part)
	for _, item := range n.childern {
		Output(item, depth+1)
	}

}
