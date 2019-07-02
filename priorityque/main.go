//This is a GOLANG reimplementation of the priority class project in the cppidioms repo
//
// This exercise will help warm up the class and methods concepts available in GO as well as see its limitations
//
//
// priorityque - implement a priorityque class which has the same feature as our cpp version
//
// The push, pop, and printf seems have some timing issues that it may output 1, -1, -1, instead of 1, 2, 3
// I have to add debug printf and time delay to get it right.
//
// A more professional way to debug is to add go extension and all the tools needed to do stepping.
// Ctl-Shift-X will launch the go extension panel to install. Actually add breakpoints and do the
// stepping usinf F5 will bring up the debugger tools needed for install

package main

import (
	"fmt"
	//	"time"
	//	"os"
	//	"bufio"
	//	"math"
	//	"strconv"
	//	"strings"
)

type Node struct {
	value int
	next  *Node
}

type Priorityque struct {
	head *Node
	tail *Node
}

func (pq *Priorityque) push(data int) {
	var node = new(Node)
	node.value = data

	//	fmt.Printf("A0 %d\n", node.value);

	if pq.head == nil {
		pq.head = node
		pq.tail = node
		fmt.Printf("A1 %d\n", node.value)
		return
	}

	if data <= pq.head.value {
		node.next = pq.head
		pq.head = node
		fmt.Printf("A2 %d\n", node.value)
		return
	}

	tmp := pq.head
	if tmp.next == nil {
		node.next = nil
		tmp.next = node
		pq.tail = node
		fmt.Printf("A3 %d\n", node.value)
		return
	}

	for tmp.next != nil {
		if data > tmp.next.value {
			if tmp.next != nil {
				tmp = tmp.next
				fmt.Printf("A6 %d\n", node.value)
				continue
			} else {
				//we are at the end
				node.next = nil
				tmp.next = node
				pq.tail = node
				fmt.Printf("A4 %d\n", node.value)
			}
		} else {
			//insert node here
			node.next = tmp.next
			tmp.next = node
			fmt.Printf("A5 %d\n", node.value)
			break
		}
		fmt.Printf("A7 %d\n", node.value)

	} //for

	fmt.Printf("A8 %d\n", node.value)

}

func (pq *Priorityque) pop() {
	if pq.head != nil {
		pq.head = pq.head.next
	} else {
		pq.tail = nil
	}
}

func (pq *Priorityque) top() int {
	if pq.head != nil {
		return pq.head.value
	} else {
		return -1
	}
}

func main() {

	var pq Priorityque
	pq.push(1)
	pq.push(3)
	pq.push(2)

	fmt.Printf("%d\n", pq.top())
	//	time.Sleep( 1 * time.Second )

	pq.pop()

	// time.Sleep( 1 * time.Second )
	fmt.Printf("%d\n", pq.top())
	// time.Sleep( 1 * time.Second )
	pq.pop()
	fmt.Printf("%d\n", pq.top())
}
