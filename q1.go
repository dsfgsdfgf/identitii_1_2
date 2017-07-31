package main

import (
	"encoding/json"
	"log"
)

type node struct {
	Name     string  `json:"name"`
	Parent   *node   `json:"-"`		// Extra parent pointer placeholder - depth iteration
	Children []*node `json:"children,omitempty"`
}


var examples = []string{
	"[a,b,c]",
	"[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]",
}

func parse(v string) (*node, error) {
	root := node{}
	root.Name = ""
	root.Children = make([]*node, 0)

	var name string 	
	var currentNode = &root	

	for _, v := range v[1:] {	// skip first "["
		currentChar := string(v)
		var newNode *node
	
		switch currentChar {
		case "[":		
			if len(name) > 0 {
				newNode = &node{}
				newNode.Name = name
				newNode.Parent = currentNode	// nest it, depth+=1

				// add created newNode to Children of Parent
				currentNode.Children = append(currentNode.Children, newNode)

				currentNode = newNode	//

				name = "" // -- name reset
			}

		case "]":	
			if len(name) > 0 {
				newNode = &node{}
				newNode.Name = name
				newNode.Parent = currentNode

				// add created newNode to Children of Parent
				currentNode.Children = append(currentNode.Children, newNode)

				// end nesting, go up depth
				currentNode = currentNode.Parent

				name = "" // -- name reset
			}

		case ",":		
			if len(name) > 0 {
				newNode = &node{}
				newNode.Name = name
				newNode.Parent = currentNode

				currentNode.Children = append(currentNode.Children, newNode)

				name = "" // -- name reset
			}
		default:
			name += currentChar // add current character of name to the current name

		}
	}
	return &root, nil
}

func main() {
	for i, example := range examples {
		result, err := parse(example)
		if err != nil {
			panic(err)
		}

		j, err := json.MarshalIndent(result, " ", " ")
		if err != nil {
			panic(err)
		}
		log.Printf("Example %d: %s - %s", i, example, string(j))
	}
}

