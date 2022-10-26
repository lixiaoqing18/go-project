package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walkIn func(t *tree.Tree)
	walkIn = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walkIn(t.Left)
		ch <- t.Value
		walkIn(t.Right)
	}
	walkIn(t)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 {
			fmt.Println("end of channel")
			return true
		}
		fmt.Printf("%v==%v \n", v1, v2)
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	/*
		t1 := tree.New(1)
		ch := make(chan int)
		go Walk(t1, ch)
		for v := range ch {
			fmt.Println(v)
		}
	*/
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(Same(t1, t2))
}
