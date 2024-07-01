package btree

import (
	"errors"
	"fmt"
	"strings"
)

const (
	DefaultMaxDegree uint = 4
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type BTree interface {
	Set(key int, value interface{}) error
	Get(key int) (interface{}, error)

	String() string
}

type btree struct {
	maxDegree uint
	root      *node
}

func New(maxDegree uint) *btree {
	return &btree{
		root:      nil,
		maxDegree: maxDegree,
	}
}

func (b *btree) Set(key int, value interface{}) error {
	item := &item{
		key:   key,
		value: value,
	}

	if b.root == nil {
		b.root = newNode(b.maxDegree)
		b.root.items = append(b.root.items, item)
		return nil
	}

	b.insert(b.root, item, nil)
	return nil
}

func (b *btree) insert(node *node, item *item, parent *node) {
	// preemptive split if node is root and is full
	if node == b.root && len(node.items) == int(b.maxDegree)-1 {
		left, pivot, right := b.split(node)
		b.root = newNode(b.maxDegree)
		b.root.items = append(b.root.items, pivot)
		b.root.children = append(b.root.children, left, right)

		b.insert(b.root, item, nil) // should be ok to insert now
		return
	}

	// preemptive split-merge if non-root is full
	if len(node.items) == int(b.maxDegree)-1 {
		b.splitAndMerge(node, parent)
		b.insert(b.root, item, nil)
		return
	}

	// insert item if leaf node has the capacity
	if node.isLeaf() && len(node.items) < int(b.maxDegree)-1 {
		var idx int
		for idx = 0; idx < len(node.items); idx++ {
			visitingItem := node.items[idx]
			if item.key <= visitingItem.key {
				break
			}
		}
		node.items = insertAt(node.items, idx, item)
		return
	}

	// traverse to next node
	var idx int
	for idx = 0; idx < len(node.items); idx++ {
		visitingItem := node.items[idx]
		if item.key <= visitingItem.key {
			break
		}
	}
	b.insert(node.children[idx], item, node)
}

func (b *btree) splitAndMerge(node *node, parent *node) {
	left, pivot, right := b.split(node)
	b.merge(parent, left, pivot, right)
}

func (b *btree) merge(node *node, left *node, pivot *item, right *node) {
	// check which index to insert within the node
	var idx int
	for idx = 0; idx < len(node.items); idx++ {
		item := node.items[idx]
		if pivot.key <= item.key {
			break
		}
	}

	node.items = insertAt(node.items, idx, pivot)
	node.children = deleteRange(node.children, idx, idx+1)
	node.children = insertAt(node.children, idx, left, right)
}

func (b *btree) split(node *node) (left *node, pivot *item, right *node) {
	idx := len(node.items) / 2
	pivot = node.items[idx]

	left = newNode(b.maxDegree)
	left.items = node.items[:idx]

	right = newNode(b.maxDegree)
	right.items = node.items[idx+1:]

	// copy children, if any
	if idx+1 < len(node.children) {
		left.children = node.children[:idx+1]
		right.children = node.children[idx+1:]
	}

	return left, pivot, right
}

func (b btree) Get(key int) (interface{}, error) {
	item := searchKey(key, b.root)
	if item == nil {
		return nil, ErrKeyNotFound
	}
	return item.value, nil
}

func (b btree) String() string {
	return b.root.String()
}

func insertAt[T any](slice []T, idx int, values ...T) []T {
	result := make([]T, 0, len(slice)+len(values))
	result = append(result, slice[:idx]...)
	result = append(result, values...)
	result = append(result, slice[idx:]...)
	return result
}

func deleteRange[T any](slice []T, start int, end int) []T {
	if len(slice) == 0 {
		return make([]T, 0)
	}

	result := make([]T, 0, len(slice)-1)
	result = append(result, slice[:start]...)
	result = append(result, slice[end:]...)
	return result
}

func searchKey(key int, node *node) *item {
	if node == nil {
		return nil
	}

	items := node.items
	// iterate over key-values to find best index
	var idx int
	for idx = 0; idx < len(items); idx++ {
		kv := items[idx]
		if key <= kv.key {
			break
		}
	}

	// check if item was actually found
	if idx < len(items) && key == items[idx].key {
		return items[idx]
	}

	// leaf node
	if len(node.children) == 0 {
		return nil
	}

	return searchKey(key, node.children[idx])
}

type node struct {
	items    []*item
	children []*node
}

func (n node) isLeaf() bool {
	return len(n.children) == 0
}

func (n *node) String() string {
	if n == nil {
		return ""
	}

	items := make([]string, 0, len(n.items)+len(n.children))
	var i, j int
	var printItem bool
	for {
		if i >= len(n.items) && j >= len(n.children) {
			break
		}
		if printItem {
			if i < len(n.items) {
				if item := n.items[i].String(); item != "" {
					items = append(items, item)
				}
				i++
			}
		} else {
			if j < len(n.children) {
				if item := n.children[i].String(); item != "" {
					items = append(items, item)
				}
				j++
			}
		}
		printItem = !printItem
	}
	return strings.Join(items, " ")
}

type item struct {
	key   int
	value interface{}
}

func (i *item) String() string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("%v:%v", i.key, i.value)
}

func newNode(maxDegree uint) *node {
	if maxDegree < 1 {
		panic("max degree should not be lesser than 1")
	}

	node := &node{
		items:    make([]*item, 0, maxDegree-1),
		children: make([]*node, 0, maxDegree),
	}
	return node
}
