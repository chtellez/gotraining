//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// --Requirements:
// * Create a function to print server status displaying:
func printServerStatus(order int, servers map[string]int) {
	fmt.Println("\n--- Servers Status", order, "---")
	//  - number of servers
	fmt.Println("There are: ", len(servers), "servers.")

	//  - number of servers for each status (Online, Offline, Maintenance, Retired)
	var totalOnline, totalOffline, totalMaintenance, totalRetired int
	//or
	totalStatus := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			totalOnline += 1
			//or
			totalStatus[status] += 1
		case Offline:
			totalOffline += 1
			//or
			totalStatus[status] += 1
		case Maintenance:
			totalMaintenance += 1
			//or
			totalStatus[status] += 1
		case Retired:
			totalRetired += 1
			//or
			totalStatus[status] += 1
		default:
			panic("Server status unhandled")
		}
	}

	fmt.Println("  Online:", totalOnline, "servers.")
	fmt.Println("  Offline:", totalOffline, "servers.")
	fmt.Println("  Maintenance:", totalMaintenance, "servers.")
	fmt.Println("  Retired:", totalRetired, "servers.")

	//or
	fmt.Println("\n  Online:", totalStatus[Online], "servers.")
	fmt.Println("  Offline:", totalStatus[Offline], "servers.")
	fmt.Println("  Maintenance:", totalStatus[Maintenance], "servers.")
	fmt.Println("  Retired:", totalStatus[Retired], "servers.")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Create a map using the server names as the key and the server status
	//  as the value
	serverStatus := make(map[string]int)

	//* Set all the server statuses to `Online` when creating the map
	for _, server := range servers {
		serverStatus[server] = Online
	}

	//* After creating the map, perform the following actions:
	//  - call display server info function
	printServerStatus(1, serverStatus)

	//  - change server status of `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired

	//  - change server status of `aiur` to `Offline`
	serverStatus["aiur"] = Offline

	//  - call display server info function
	printServerStatus(2, serverStatus)

	//  - change server status of all servers to `Maintenance`
	for _, server := range servers {
		serverStatus[server] = Maintenance
	}

	//  - call display server info function
	printServerStatus(3, serverStatus)
}
