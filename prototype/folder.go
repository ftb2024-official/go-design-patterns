package prototype

import "fmt"

type Folder struct {
	children []INode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, v := range f.children {
		v.print(indentation + indentation)
	}
}

func (f *Folder) clone() INode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []INode

	for _, v := range f.children {
		copy := v.clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.children = tempChildren
	return cloneFolder
}
