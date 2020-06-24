package singleheader

import "fmt"

type FileType int

const (
	Header FileType = iota
	Source
)

type File struct {
	Type      FileType
	Name      string
	Extension string
	Path      string
	Content   string
}

// TODO: Should we track all the functions? Dividing them up by private/public?
// Be able to check for namespace issues? Or not since we are going to namespace
// the entire library anyways?

func ParseFiles(arg string) {
	files := []*File{}
	// TODO: Glob search for all *.h files in project  (is cpp hpp? i forgot)
	// TODO: Glob search for all the *.c files in a project
	// TODO: Glob search for all the *.cpp
	for _, file := range files {
		fmt.Println("Loaded: ", file)
	}
}

func RemoveHeaderFileInclude() {
	//        if ".h" in file:
	//            str = str.replace("#include \"" + fname + "\"", "");
	//            str = str.replace("#include <" + fname + ">", "");
}
