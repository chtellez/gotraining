//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted(m int) {
	fmt.Println("Granted", m)
}

func accessDenied(m int) {
	fmt.Println("Denied", m)
}

func checkAccess(today int, role int) {
	if role > Member && (today == Tuesday || today == Thursday || today > Friday) {
		accessDenied(1)
	} else if role > Manager && role < Guest && today > Friday {
		accessDenied(2)
	} else {
		accessGranted(3)
	}
}

func main() {
	// The day and role. Change these to check your work.
	fmt.Println("Admin -------------------------------------------------")
	checkAccess(Monday, Admin)
	checkAccess(Tuesday, Admin)
	checkAccess(Wednesday, Admin)
	checkAccess(Thursday, Admin)
	checkAccess(Friday, Admin)
	checkAccess(Saturday, Admin)
	checkAccess(Sunday, Admin)

	fmt.Println("Manager -------------------------------------------------")
	checkAccess(Monday, Manager)
	checkAccess(Tuesday, Manager)
	checkAccess(Wednesday, Manager)
	checkAccess(Thursday, Manager)
	checkAccess(Friday, Manager)
	checkAccess(Saturday, Manager)
	checkAccess(Sunday, Manager)

	fmt.Println("Contractor -------------------------------------------------")
	checkAccess(Monday, Contractor)
	checkAccess(Tuesday, Contractor)
	checkAccess(Wednesday, Contractor)
	checkAccess(Thursday, Contractor)
	checkAccess(Friday, Contractor)
	checkAccess(Saturday, Contractor)
	checkAccess(Sunday, Contractor)

	fmt.Println("Guest -------------------------------------------------")
	checkAccess(Monday, Guest)
	checkAccess(Tuesday, Guest)
	checkAccess(Wednesday, Guest)
	checkAccess(Thursday, Guest)
	checkAccess(Friday, Guest)
	checkAccess(Saturday, Guest)
	checkAccess(Sunday, Guest)

	fmt.Println("Member -------------------------------------------------")
	checkAccess(Monday, Member)
	checkAccess(Tuesday, Member)
	checkAccess(Wednesday, Member)
	checkAccess(Thursday, Member)
	checkAccess(Friday, Member)
	checkAccess(Saturday, Member)
	checkAccess(Sunday, Member)
}
