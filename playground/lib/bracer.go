package main

import (
	"fmt"
)

type Node struct {
	Value string
}

type Stack struct {
	nodes []*Node
}

func mainbracer() {
	flag := true
	stack := &Stack{}
	openBraces := map[string]int{"{": 1, "[": 2, "(": 3}
	closeBraces := map[string]int{"}": 1, "]": 2, ")": 3}
	strInput := "[(]"
	for _, v := range strInput {
		// fmt.Printf("%T: %s \n", v, string(v))
		if openBraces[string(v)] > 0 {
			stack.Push(&Node{string(v)})
			continue
		}
		if closeBraces[string(v)] > 0 {
			if stack.Size() == 0 {
				flag = false
				break
			}
			lastNode := stack.Last()
			fmt.Println(lastNode.Value)
			if closeBraces[string(v)] == openBraces[lastNode.Value] {
				stack.Pop()
			} else {
				flag = false
				break
			}
		}
	}
	if flag && stack.Size() == 0 {
		fmt.Println("it's balance")
	} else {
		fmt.Println("it's not balance")
	}
}

func (stack *Stack) Push(node *Node) {
	stack.nodes = append(stack.nodes, node)
}

func (stack *Stack) Pop() {
	if stack.Size() > 0 {
		stack.nodes = stack.nodes[:len(stack.nodes)-1]
	}
}

func (stack *Stack) Size() int {
	return len(stack.nodes)
}

func (stack *Stack) Last() *Node {
	return stack.nodes[stack.Size()-1]
}
