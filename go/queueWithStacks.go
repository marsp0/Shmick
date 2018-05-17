package main

import "fmt"

// MyQueue - represents a Q object.
type MyQueue struct {
	// https://leetcode.com/problems/implement-queue-using-stacks/description/
	frontStack []int
	backStack  []int
	length     int
}

// Constructor - constructs a new Q
func Constructor() MyQueue {
	var q = MyQueue{}
	return q
}

// Push - inserts an element in the Q
func (queue *MyQueue) Push(x int) {
	queue.frontStack = []int{}
	for i := len(queue.backStack) - 1; i >= 0; i-- {
		queue.frontStack = append(queue.frontStack, queue.backStack[i])
	}
	queue.frontStack = append(queue.frontStack, x)
	queue.backStack = []int{}
	for i := len(queue.frontStack) - 1; i >= 0; i-- {
		queue.backStack = append(queue.backStack, queue.frontStack[i])
	}
	queue.length++
	fmt.Println(queue.frontStack, queue.backStack)
}

// Pop - removes the first element from the Q
func (queue *MyQueue) Pop() int {
	var toReturn = queue.backStack[len(queue.backStack)-1]
	queue.backStack = queue.backStack[:len(queue.backStack)-1]
	queue.length--
	return toReturn
}

// Peek - check the first element in the q
func (queue *MyQueue) Peek() int {
	if queue.length > 0 {
		return queue.backStack[len(queue.backStack)-1]
	}
	return 0
}

// Empty - check if the Q is empty
func (queue *MyQueue) Empty() bool {
	return queue.length == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
