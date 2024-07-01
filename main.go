package main

import (
	"fmt"

	"github.com/turao/go-btree/btree"
)

func main() {
	tree := btree.New(btree.DefaultMaxDegree)
	fmt.Println(tree.String())
	tree.Set(70, 70)
	fmt.Println(tree.String())
	tree.Set(60, 60)
	fmt.Println(tree.String())
	tree.Set(50, 50)
	fmt.Println(tree.String())
	tree.Set(40, 40)
	fmt.Println(tree.String())
	tree.Set(30, 30)
	fmt.Println(tree.String())
	tree.Set(20, 20)
	fmt.Println(tree.String())
	tree.Set(10, 10)
	fmt.Println(tree.String())
	tree.Set(00, 00)
	fmt.Println(tree.String())
	tree.Set(15, 15)
	fmt.Println(tree.String())
	tree.Set(25, 25)
	fmt.Println(tree.String())
	tree.Set(35, 35)
	fmt.Println(tree.String())
	tree.Set(45, 45)
	fmt.Println(tree.String())
	tree.Set(55, 55)
	fmt.Println(tree.String())
	tree.Set(65, 65)
	fmt.Println(tree.String())
	tree.Set(75, 75)
	fmt.Println(tree.String())
}
