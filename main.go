package main

import "fmt"

//并查集
//图，环检测算法
type Graph struct { //存储图的数据结构
	point int      //点的个数
	edge  [][2]int //存储边
}

func getDefaultData(graph *Graph) {
	graph.point = 6
	graph.edge = append(graph.edge,
		[2]int{0, 1},
		[2]int{1, 2},
		[2]int{1, 3},
		[2]int{2, 5},
		[2]int{3, 4},
	)
}

type Node struct {
	level int   //记录层级,使树平衡
	root  *Node //指向代表
}

func main() {
	var graph Graph
	getDefaultData(&graph)
	data := make([]Node, graph.point)
	for _, edge := range graph.edge {
		//寻找左代表
		left, right := &data[edge[0]], &data[edge[1]]
		for left.root != nil {
			left = left.root
		}
		for right.root != nil {
			right = right.root
		}
		if left == right {
			fmt.Println("检测到环")
			return
		}
		if left.level > right.level {
			right.root = left
		} else if left.level < right.level {
			left.root = right
		} else {
			left.root = right
			left.level++
		}
	}
	fmt.Println("未检测到环")
}
