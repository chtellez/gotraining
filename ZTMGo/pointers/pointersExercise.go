//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//

package main

import (
	"fmt"
)

//--Requirements:

// - Security tags have two states: active (true) and inactive (false)
const (
	Active   = true
	Inactive = false
)

type securityTag bool

// Item Create a structure to store items and their security tag state
type Item struct {
	name string
	tag  securityTag
}

// * Create functions to activate and deactivate security tags using pointers
func activate(tag *securityTag) {
	*tag = Active
}

func inactive(tag *securityTag) {
	*tag = Inactive
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {
	for i := 0; i < len(items); i++ {
		inactive(&items[i].tag)
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	items := []Item{
		{name: "Toy", tag: Active},
		{name: "Pencil", tag: Active},
		{name: "Notebook", tag: Active},
		{name: "Fan", tag: Active},
	}
	fmt.Println(items)

	//  - Deactivate any one security tag in the array/slice
	inactive(&items[1].tag)
	inactive(&items[3].tag)
	fmt.Println(items)

	//  - Call the checkout() function to deactivate all tags
	checkout(items)
	fmt.Println(items)

	//  - Print out the array/slice after each change
}
