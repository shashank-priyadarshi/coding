Composite in Go
Composite is a structural design pattern that allows composing objects into a tree-like structure and work with the it as if it was a singular object.

Composite became a pretty popular solution for the most problems that require building a tree structure. Composite’s great feature is the ability to run methods recursively over the whole tree structure and sum up the results.

Conceptual Example
Let’s try to understand the Composite pattern with an example of an operating system’s file system. In the file system, there are two types of objects: files and folders. There are cases when files and folders should be treated to be the same way. This is where the Composite pattern comes in handy.

Imagine that you need to run a search for a particular keyword in your file system. This search operation applies to both files and folders. For a file, it will just look into the contents of the file; for a folder, it will go through all files of that folder to find that keyword.

### component.go: Component interface
```go
package main

type Component interface {
    search(string)
}
```
### folder.go: Composite
```go
package main

import "fmt"

type Folder struct {
    components []Component
    name       string
}

func (f *Folder) search(keyword string) {
    fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
    for _, composite := range f.components {
        composite.search(keyword)
    }
}

func (f *Folder) add(c Component) {
    f.components = append(f.components, c)
}
```
### file.go: Leaf
```go
package main

import "fmt"

type File struct {
    name string
}

func (f *File) search(keyword string) {
    fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
    return f.name
}
```
### main.go: Client code
```go
package main

func main() {
    file1 := &File{name: "File1"}
    file2 := &File{name: "File2"}
    file3 := &File{name: "File3"}

    folder1 := &Folder{
        name: "Folder1",
    }

    folder1.add(file1)

    folder2 := &Folder{
        name: "Folder2",
    }
    folder2.add(file2)
    folder2.add(file3)
    folder2.add(folder1)

    folder2.search("rose")
}
```
### output.txt: Execution result
```
Serching recursively for keyword rose in folder Folder2
Searching for keyword rose in file File2
Searching for keyword rose in file File3
Serching recursively for keyword rose in folder Folder1
Searching for keyword rose in file File1
```