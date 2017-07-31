package main

import (
	"encoding/json"
	"log"
)

type node struct {
	Name     string  `json:"name"`
	Parent   *node   `json:"-"`
	Children []*node `json:"children,omitempty"`
}

func main() {
	input := "[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]"

	root := node{}
	root.Name = ""
	root.Children = make([]*node, 0)

	var name string

	var currentNode = &root

	for _, v := range input[1 : len(input)-2] {
		currentChar := string(v)
		var newNode *node

		switch currentChar {
		case "[":
			newNode = &node{}
			newNode.Name = name
			newNode.Parent = currentNode

			currentNode.Children = append(currentNode.Children, newNode)
			currentNode = newNode

			name = ""
		case "]":
			currentNode = currentNode.Parent

		case ",":
			newNode = &node{}
			newNode.Name = name
			newNode.Parent = currentNode
			currentNode.Children = append(currentNode.Children, newNode)
			name = ""
		default:
			name += currentChar

		}
	}

	j, err := json.MarshalIndent(root, " ", " ")
	if err != nil {
		panic(err)
	}
	log.Printf(string(j))
}
