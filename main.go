package main

import (
	"log"

	"github.com/turao/go-btree/btree"
)

func main() {
	tree := btree.New(btree.DefaultMaxDegree)
	tree.Set(70, 70)
	tree.Set(60, 60)
	tree.Set(50, 50)
	tree.Set(40, 40)
	tree.Set(30, 30)
	tree.Set(20, 20)
	tree.Set(10, 10)
	tree.Set(00, 00)
	tree.Set(15, 15)
	tree.Set(25, 25)
	tree.Set(35, 35)
	tree.Set(45, 45)
	tree.Set(55, 55)
	tree.Set(65, 65)
	tree.Set(75, 75)
	log.Println(tree.String())

	val, err := tree.Get(25)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(val)
}
