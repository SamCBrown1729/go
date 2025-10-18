package main
 	

import (
    "fmt"
    "os"
	"time"
	"log"
)


type node struct {
	freq int
	char string
	left *node
	right *node
}


func cutNode(cut *node) {
	
	if cut.right == nil {
		cut.left.right = nil
	} else if cut.left == nil {
		cut.right.left = nil
	} else {
		cut.right.left = cut.left
		cut.left.right = cut.right
	}

}


func insert(newNode, relativeNode *node, where string) {
	if newNode == relativeNode {
		
		return
	}
	if where == "right" {
		if relativeNode.right != nil{
			newNode.right = relativeNode.right
			newNode.right.left = newNode
		} else {
			newNode.right = nil
		}

		newNode.left = relativeNode
		relativeNode.right = newNode
	} else if where == "left" {
		if relativeNode.left != nil{
			newNode.left = relativeNode.left
		 	newNode.left.right = newNode
		} else {
			newNode.left = nil
		}

		newNode.right = relativeNode
		relativeNode.left = newNode
	} else {
		fmt.Println("That's not a direction")
	}
	
}
	

func moveToHead(newNode *node) {
	cutNode(newNode)
	largest := true
	nextNode := newNode
	
	for {
		largest = true
		if newNode.freq < nextNode.freq {
			insert(newNode, nextNode, "right")
			largest = false
			break;
		}
		if nextNode.left == nil {
			break
		}
		nextNode = nextNode.left

	}
	
	if nextNode == newNode {
		insert(newNode, newNode.right, "left")
	} else if largest {
		insert(newNode, nextNode, "left")
	}

}


func initialNodes(filename string) *node {	
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	
	head := &node{freq: 1, char : string(dat[0])}
	
	tail := head
	for _, letter := range dat[1:]{
		foundLetter := false
		newNode := head
		// Steps through the nodes from rightmost to leftmost checking if chars match	
		for {
			if newNode.char == string(letter) {
				newNode.freq += 1
				foundLetter = true
				moveToHead(newNode)
				for head.left != nil{
					head = head.left
				}
				break
			}
			if newNode.right == nil {
				break
			}
			newNode = newNode.right
			
		}


		if !foundLetter{
			newNode = &node{freq : 1, char : string(letter)}
			tail.right = newNode
			newNode.left = tail
			tail = newNode
		}
	
	}

	return tail 

}


func createTree(tail *node) (*node, *node){

	rightNode := tail
	leftNode := tail.left
	tail = leftNode.left
	newTreeNode := &node{rightNode.freq + leftNode.freq, leftNode.char + rightNode.char, leftNode, rightNode}
	newNode := &node{freq : rightNode.freq + leftNode.freq, char : leftNode.char + rightNode.char, left : leftNode.left}
	moveToHead(newNode)
	if len(rightNode.char) == 1 {
		rightNode.left = nil
		rightNode.right = nil
	}
	if len(leftNode.char) == 1  {
		leftNode.left = nil
		leftNode.right = nil
	}
 	return newTreeNode, tail
}
		

func main() {
	start := time.Now()
	startNode := initialNodes("PhdFrom.pdf")
	elapsed := time.Since(start)
	log.Printf("Initial took %s", elapsed)

	nextNode := startNode
	fmt.Println(*nextNode)
	for nextNode.left != nil {
		nextNode = nextNode.left
		fmt.Println(*nextNode)
	}
	
}


