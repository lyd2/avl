package main

import (
	"fmt"
	"github.com/lyd/avl"
	"strconv"
)

func main() {
	example1()
	fmt.Printf("\n\n")
	example2()
}

func example1() {
	avlTree := avl.New()

	for i := 0; i < 100; i++ {
		avlTree.Insert(strconv.FormatInt(int64(i), 10), i)
	}

	it := avlTree.InOrder()
	for i := 0; i < len(it.List); i++ {
		fmt.Println(it.List[i])
	}

	fmt.Println("depth", avlTree.Depth())
}

func example2() {
	avlTree := avl.New()

	avlTree.Insert("a", "1")
	fmt.Println(avlTree.Search("a"))

	avlTree.Insert("b", "2")
	fmt.Println(avlTree.Search("b"))

	avlTree.Insert("c", "3")
	fmt.Println(avlTree.Search("c"))

	avlTree.Insert("d", "4")
	fmt.Println(avlTree.Search("d"))

	avlTree.Insert("e", "5")
	fmt.Println(avlTree.Search("e"))

	fmt.Println("depth", avlTree.Depth())
}
