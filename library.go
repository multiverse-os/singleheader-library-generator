package singleheader

import (
	"strings"
)

// TODO: These should be using a io.Writer and using the buffer. So that way in
// the same way you work with graphics, you establish a buffer, write everything
// to that, then send that as a batch to its final location

// TODO: Will eventually need to pull in all non-standard libraries to ensure we
// create single headers that are actually single headers, we should merge in
// dependencies that are non-standard libraries if necessary.
type Dependency struct {
	Name            string
	StandardLibrary bool
}

type Library struct {
	Name          string
	Description   string
	About         string
	Lanuage       Language
	Path          string
	Files         []*File
	ParsedContent *Content
	//intro_files = []
	//pub_files = []
	//priv_files = []
	//outro_files = []
}

type Content struct {
	PublicHeader  string
	PrivateHeader string
	Source        string
}

// NOTE:
// About is dependent on if it is included in the configuration file, wheras the
// usage section is generated.
func (self Library) Introduction() (output string) {
	output += "/*\n"
	output += "\n"
	output += strings.ToLower(self.Name) + ".h - Single Header (or Header Only) Library \n"
	output += "\n"
	output += "ABOUT:\n"
	output += "    " + "This file is a single file header only library for: " + self.Name
	output += "    " + "and is merged with everything needed to work with a single inclusion"
	output += "    " + "call to this header file. There are no dependencies."
	output += "\n"
	output += "USAGE:\n"
	output += "    " + "Single header libraries simplify inclusion of libraries into a new project"
	output += "    " + "by merging all the code of a library into a single header. Developers who"
	output += "    " + "require this library for their project simply need to add a single include"
	output += "    " + "and the entire library is encapsulated in a namespace to further simplify"
	output += "    " + "using the library, and avoiding potential namespace pollution issues."
	output += "\n"
	output += "*/"
	return output
}

func (self Library) Banner(title string) (output string) {
	output += "/* ***********************************************************\n"
	output += "\n"
	output += "                       " + title + "\n"
	output += "\n"
	output += "*********************************************************** */\n"
	return output
}

// NOTE:
// You can include code instead of a comment using, that will not be included
// in the compiled version:
//   #if 0
//
