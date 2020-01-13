package avl

import (
	"errors"
	"fmt"
	"math"
)

type avlNode struct {
	key    string
	value  interface{}
	left   *avlNode
	right  *avlNode
	height int
}

type avl struct {
	root  *avlNode
	count int
}

func New() *avl {
	return &avl{
		root:  nil,
		count: 0,
	}
}

func createNode(key string, value interface{}) *avlNode {
	return &avlNode{
		key:    key,
		value:  value,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

func (avl *avl) find(key string) (*avlNode, []*avlNode) {

	stack := make([]*avlNode, 0)

	if avl.root == nil {
		return nil, stack
	}

	node := avl.root
	for true {

		stack = append(stack, node)

		if key > node.key && node.right != nil {
			node = node.right
			continue
		}

		if key < node.key && node.left != nil {
			node = node.left
			continue
		}

		break
	}

	return node, stack
}

func (avl *avl) Insert(key string, value interface{}) {

	newNode := createNode(key, value)
	node, stack := avl.find(key)

	if node == nil {
		avl.root = newNode
		return
	}

	if key == node.key {
		node.value = value
		return
	}

	if key > node.key {
		node.right = newNode
		avl.count++
	} else {
		node.left = newNode
		avl.count++
	}
	stack = append(stack, newNode)

	var grandson *avlNode = nil
	var children *avlNode = nil
	var curr *avlNode = nil
	var parent *avlNode = nil

	for i := len(stack) - 1; i >= 0; i-- {

		grandson = children
		children = curr
		curr = stack[i]
		if i == 0 {
			parent = nil
		} else {
			parent = stack[i-1]
		}

		if int(math.Abs(float64(getHeight(curr.left)-getHeight(curr.right)))) > 1 {
			if grandson == nil || children == nil {
				panic("unknown error")
			}

			if grandson.key < children.key && children.key < curr.key {
				avl.ll(parent, curr, children, grandson)
			} else if grandson.key > children.key && children.key < curr.key {
				avl.lr(parent, curr, children, grandson)
			} else if grandson.key < children.key && children.key > curr.key {
				avl.rl(parent, curr, children, grandson)
			} else {
				avl.rr(parent, curr, children, grandson)
			}

			updateHeight(curr)
			updateHeight(children)
			updateHeight(grandson)

			break
		}

		updateHeight(curr)

	}
}

func (avl *avl) Search(key string) (value interface{}, err error) {

	if avl.root == nil {
		value = nil
		err = errors.New("key does not exist")
		return
	}

	node := avl.root
	for true {

		if key > node.key && node.right != nil {
			node = node.right
			continue
		}

		if key < node.key && node.left != nil {
			node = node.left
			continue
		}

		break
	}

	if node.key != key {
		value = nil
		err = errors.New("key does not exist")
		return
	}

	value = node.value
	err = nil
	return
}

func getHeight(node *avlNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func updateHeight(node *avlNode) {
	leftHeight := getHeight(node.left)
	rightHeight := getHeight(node.right)

	if leftHeight > rightHeight {
		node.height = leftHeight + 1
	} else {
		node.height = rightHeight + 1
	}
}

func (avl *avl) ll(parent *avlNode, curr *avlNode, children *avlNode, grandson *avlNode) {
	if parent == nil {
		avl.root = children
	} else {
		if curr == parent.left {
			parent.left = children
		} else {
			parent.right = children
		}
	}

	curr.left = children.right
	children.right = curr
}

func (avl *avl) lr(parent *avlNode, curr *avlNode, children *avlNode, grandson *avlNode) {
	if parent == nil {
		avl.root = grandson
	} else {
		if curr == parent.left {
			parent.left = grandson
		} else {
			parent.right = grandson
		}
	}

	children.right = grandson.left
	curr.left = grandson.right
	grandson.left = children
	grandson.right = curr
}

func (avl *avl) rl(parent *avlNode, curr *avlNode, children *avlNode, grandson *avlNode) {
	if parent == nil {
		avl.root = grandson
	} else {
		if curr == parent.left {
			parent.left = grandson
		} else {
			parent.right = grandson
		}
	}

	curr.right = grandson.left
	children.left = grandson.right
	grandson.left = curr
	grandson.right = children
}

func (avl *avl) rr(parent *avlNode, curr *avlNode, children *avlNode, grandson *avlNode) {
	if parent == nil {
		avl.root = children
	} else {
		if curr == parent.left {
			parent.left = children
		} else {
			parent.right = children
		}
	}

	curr.right = children.left
	children.left = curr
}

func (n *avlNode) String() string {
	return fmt.Sprintf("[key=%s, value=%v, h=%d]", n.key, n.value, n.height)
}

func (n *avlNode) GetKey() string {
	return n.key
}

func (n *avlNode) GetValue() interface{} {
	return n.value
}

func (n *avlNode) GetHeight() int {
	return n.height
}

func (avl *avl) Depth() int {
	return avl.root.height
}

/**
中序遍历
*/
type traversal struct {
	List []*avlNode
}

func (it *traversal) inOrder(node *avlNode) {
	if node == nil {
		return
	}
	it.inOrder(node.left)
	it.List = append(it.List, node)
	it.inOrder(node.right)
}

func (avl *avl) InOrder() *traversal {
	it := &traversal{List: []*avlNode{}}

	if avl.root != nil {
		it.inOrder(avl.root)
	}

	return it
}
