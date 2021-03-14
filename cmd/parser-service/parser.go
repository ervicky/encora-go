package main

import (
	"encoding/json"
)

type node struct {
	Name     string  `json:"name"`
	Children []*node `json:"children,omitempty"`
}

// newNode returns the pointer instance of node.
func newNode(name string, children []*node) *node {
	return &node{
		Name:     name,
		Children: children,
	}
}

// Parse the text strings into a tree of node structs.
func parse(str string) (*node, error) {
	var root node
	// store the name current node
	var nodeName string
	// basically used to store json string
	var jsonStr string
	// to handle the jsonStr for the same level nodes
	var flag bool

	// special characters used by str.
	staringBrace := "["
	endingBrace := "]"
	separator := ","

	nameTemplate := `{"name": "`
	childrenTemplate := `", "children": [`
	nodeEndingTemplate := `]}`

	for i, v := range str {
		// to get the character from byte.
		ch := string(v)

		// when character is starting bracket.
		if ch == staringBrace {
			// to append the nodeName between name & children template.
			jsonStr += nameTemplate + nodeName + childrenTemplate

			// to clear the nodeName once we append it to string.
			nodeName = ""

			// to skip to the next loop cycle.
			continue
		}

		// when character is ending bracket.
		if ch == endingBrace {
			// to append the nodeName between name & children template.
			if flag || nodeName != "" {
				jsonStr += nameTemplate + nodeName + childrenTemplate
				flag = false
			}

			// to clear the nodeName once we append it to string.
			nodeName = ""
			jsonStr += nodeEndingTemplate

			// to append the outer closing bracket.
			if i+1 == len(str) {
				jsonStr += nodeEndingTemplate
			}

			// to skip to the next loop cycle.
			continue
		}

		// when character is separator.
		if ch == separator {
			if nodeName == "" {
				// end the previous node
				jsonStr += nodeEndingTemplate + separator
			} else {
				// to append the node with no children
				jsonStr += nameTemplate + nodeName + childrenTemplate + nodeEndingTemplate + separator
				flag = true
			}

			// to clear the nodeName once we append it to string.
			nodeName = ""

			// to skip to the next loop cycle.
			continue
		}

		// to append the character to nodeName
		nodeName += ch
	}

	// unmarshal the json string into node struct
	if err := json.Unmarshal([]byte(jsonStr), &root); err != nil {
		return nil, err
	}

	return &root, nil
}
