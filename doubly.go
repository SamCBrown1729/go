package main
 	

import (
    "fmt"
    "os"
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
	nextNode := newNode.left
	for nextNode.left != nil {
		if newNode.freq < nextNode.freq {
			insert(newNode, nextNode, "right")
			largest = false
			break;
		}
		nextNode = nextNode.left
	}
	if largest {
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
		fmt.Println(string(letter), letter)
		foundLetter := false
		newNode := tail 
		for newNode.left != nil {
			if newNode.char == string(letter) {
				newNode.freq += 1
				foundLetter = true
				moveToHead(newNode)
				break
			}
			newNode = newNode.left
		}
		
		if foundLetter{continue}
			
		newNode = &node{freq : 1, char : string(letter)}
		tail.right = newNode
		newNode.left = tail
		tail = newNode
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
	if len(leftNode.char) == 1 {
		leftNode.left = nil
		leftNode.right = nil
	}
 	return newTreeNode, tail
}
		

func main() {
	startNode := initialNodes("test.txt")
	
dat, _ := os.ReadFile("test.txt")
	fmt.Println(dat)	
	treeNode, tail := createTree(startNode)
	fmt.Println(*treeNode)
	fmt.Println(*treeNode.right)
	fmt.Println(*treeNode.left)
	
	nextNode := tail
	fmt.Println(*nextNode)
	for nextNode.left != nil {
		nextNode = nextNode.left
		fmt.Println(*nextNode)
	}
	
}


