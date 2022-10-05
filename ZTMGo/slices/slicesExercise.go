// --Summary:
//
//	Create a program to manage parts on an assembly line.

//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//  - Print out the contents of the assembly line at each step

package main

import (
	"fmt"
)

type Part string

// * Create a function to print out the contents of the assembly line
func printAssemblyLine(title string, assemblyLine []Part) {
	fmt.Println("\n---", title, "---")
	for i := 0; i < len(assemblyLine); i++ {
		fmt.Println(assemblyLine[i])
	}
}

func main() {
	//* Perform the following:
	//  - Create an assembly line having any three parts
	assemblyLine := []Part{"Cutting", "Mixing", "Pasting"}
	printAssemblyLine("assembly line 1", assemblyLine)
	//	- Add two new parts to the line
	assemblyLine = append(assemblyLine, "Cleaning", "Storing")
	printAssemblyLine("assembly line 2", assemblyLine)
	//  - Slice the assembly line, so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	printAssemblyLine("assembly line 3", assemblyLine)
	//--Notes:
	//* Your program output should list 3 parts, then 5 parts, then 2 parts
}
