package prototype

import "fmt"

func PrototypeMain() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []INode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []INode{file1, file2, file3},
		name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print(" ")

	cloneFolder := folder1.clone()
	fmt.Println("\nPrinting hierarchy for Folder1")
	cloneFolder.print(" ")
}
