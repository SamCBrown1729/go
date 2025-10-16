package main
 	

import (
    "fmt"
    "os"
)


type node struct {
	freq int
	char byte
	left *node
	right *node
}


func cutNode(cut *node) {
	
	if newNode.right == nil {
		newNode.left.right = nil
	} else if newNode.left == nil {
		newNode.right.left = nil
	} else {
		newNode.right.left = newNode.left
		newNode.left.right = newNode.right
	}

}

func insert(newNode, relativeNode *node, where string) {
	if string == "right" {
		
}
	


func moveToHead(newNode *node) {
	cutNode(newNode)
	
	nextNode := newNode.left
	for nextNode.left != nil {
		if newNode.freq < nextNode.freq {
			insert(newNode, relativeNode, right)
			break;
		}
		nextNode = nextNode.left
	}
	insert(newNode, nextNode, left)
}


		

func initialNodes(filename string) *node {	
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(dat)
	
	head := &node{freq: 1, char: dat[0]}
	
	tail := head
	for _, letter := range dat[1:len(dat)-1]{
		foundLetter := false
		newNode := tail 
		for newNode.left != nil {
			if newNode.char == letter {
				newNode.freq += 1
				foundLetter = true
				moveToHead(newNode)
				break
			}
			newNode = newNode.left
		}
		
		if foundLetter{continue}
			
		newNode = &node{freq : 1, char : letter}
		tail.right = newNode
		newNode.left = tail
		tail = newNode
	}

	return tail 

}
	


func main() {
	startNode := initialNodes("test.txt")
	
	nextNode := startNode
	for nextNode.left != nil {
		fmt.Println(*nextNode)
		nextNode = nextNode.left
	}


}



