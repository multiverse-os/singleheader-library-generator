package singleheader

import (
	"strings"
)

// #ifndef FLAGS_H_
// #define FLAGS_H_

func SingleHeaderTemplateMethod1(libraryName, source string) (output string) {
	output += "#ifndef " + strings.ToUpper(libraryName) + "\n"
	output += "#define " + strings.ToUpper(libraryName) + "\n"
	output += "\n"
	output += "namespace " + libraryName + "\n"
	output += "{\n"
	output += source
	output += "\n"
	output += "}\n"
	output += "#endif"
	return output
}

//print(os.linesep + "#ifdef " + macro + "_IMPLEMENTATION");
//print("");
func SingleHeaderTemplateMethod2(libraryName, publicHeader, privateHeader, source string) (output string) {
	//# print(os.linesep + "#ifndef " + macro + "_SINGLE_HEADER");
	//# print("#define " + macro + "_SINGLE_HEADER");

	//for f in pub_files:
	//    sys.stdout.write(open(f, 'r').read())
	//# print("#endif /* " + macro + "_SINGLE_HEADER */");
	output += "#ifndef " + libraryName + "_SINGLE_HEADER \n"
	output += "#define " + libraryName + "_SINGLE_HEADER \n"
	output += publicHeader
	output += "\n"
	output += "#endif \n"
	output += "\n"
	//print("#endif");
	//print("");

	//print("#ifndef NK_SINGLE_FILE");
	//print("  #define NK_SINGLE_FILE");

	//for f in priv_files:
	//    print(omit_includes(open(f, 'r').read(),
	//                        pub_files + priv_files))
	//print("#endif /* " + macro + "_IMPLEMENTATION */");
	output += "#ifndef " + libraryName + "_SINGLE_HEADER \n"
	output += "#define " + libraryName + "_SINGLE_HEADER \n"
	output += privateHeader
	output += "\n"
	output += source
	output += "\n"
	output += "#endif \n"

	//print(os.linesep + "/*")
	//for f in outro_files:
	//    sys.stdout.write(open(f, 'r').read())
	//print("*/" + os.linesep)

	//     /*
	//     [outro file contents]
	//     */""")
	return output
}
