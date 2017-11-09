package main

import "fmt"

var addressBookZoneSlice = make([]string, 0)
var addressBookNameSlice = make([]string, 0)
var addressBookIPSlice = make([]string, 0)

var applicationsNameSlice = make([]string, 0)
var applicationsProtocolSlice = make([]string, 0)
var applicationsPortSlice = make([]string, 0)

var securityPoliciesSlice = make([]string, 0)

// Address Book
func addressBookInput() {
	fmt.Println(`Address Book
────────────`)

	for {
		var zoneString string
		var nameString string
		var ipString string

		fmt.Print("Zone:        ")
		fmt.Scanln(&zoneString)

		if zoneString != "" {
			fmt.Print("Name:        ")
			fmt.Scanln(&nameString)
		} else {
			break
		}

		if nameString != "" {
			fmt.Print("IP:          ")
			fmt.Scanln(&ipString)
		} else {
			break
		}

		if ipString != "" {
			addressBookZoneSlice = append(addressBookZoneSlice, zoneString)
			addressBookNameSlice = append(addressBookNameSlice, nameString)
			addressBookIPSlice = append(addressBookIPSlice, ipString)
		} else {
			break
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
}

func addressBookOutput() {
	var zoneCount int
	var nameCount int
	var ipCount int
	for i := 0; zoneCount != len(addressBookZoneSlice) && nameCount != len(addressBookNameSlice) && ipCount != len(addressBookIPSlice); i++ {
		fmt.Println("set security zones security-zone", addressBookZoneSlice[i], "address-book address", addressBookNameSlice[i], addressBookIPSlice[i])
		zoneCount++
		nameCount++
		ipCount++
	}
}

// Applications
func applicationsInput() {
	fmt.Println(`Applications
────────────`)

	for {
		var nameString string
		var protocolString string
		var portString string

		fmt.Print("Name:        ")
		fmt.Scanln(&nameString)

		if nameString != "" {
			fmt.Print("Protocol:    ")
			fmt.Scanln(&protocolString)
		} else {
			break
		}

		if protocolString != "" {
			fmt.Print("Port:        ")
			fmt.Scanln(&portString)
		} else {
			break
		}

		if portString != "" {
			applicationsNameSlice = append(applicationsNameSlice, nameString)
			applicationsProtocolSlice = append(applicationsProtocolSlice, protocolString)
			applicationsPortSlice = append(applicationsPortSlice, portString)
		} else {
			break
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
}

func applicationsOutput() {
	var nameCount int
	var protocolCount int
	var portCount int
	for i := 0; nameCount != len(applicationsNameSlice) && protocolCount != len(applicationsProtocolSlice) && portCount != len(applicationsPortSlice); i++ {
		fmt.Println("set applications application", applicationsNameSlice[i], "protocol", applicationsProtocolSlice[i])
		fmt.Println("set applications application", applicationsNameSlice[i], "destination-port", applicationsPortSlice[i])
		nameCount++
		protocolCount++
		portCount++
	}
}

// Security Policies
func securityPoliciesInput() {
	fmt.Println(`Security Policies
─────────────────`)
	var securityPoliciesString string
	for {
		var fromZone string
		var toZone string
		var policy string
		fmt.Print("From zone:   ")
		fmt.Scanln(&fromZone)
		if fromZone == "" {
			break
		}
		fmt.Print("To zone:     ")
		fmt.Scanln(&toZone)
		if toZone == "" {
			break
		}
		fmt.Print("Policy:      ")
		fmt.Scanln(&policy)
		if policy == "" {
			break
		}
		for {
			var source string
			fmt.Print("Source:      ")
			fmt.Scanln(&source)
			if source == "" {
				break
			}
			securityPoliciesString = "set security policies from-zone " + fromZone + " to-zone " + toZone + " policy " + policy + " match source-address " + source
			securityPoliciesSlice = append(securityPoliciesSlice, securityPoliciesString)
		}
		for {
			var destination string
			fmt.Print("Destination: ")
			fmt.Scanln(&destination)
			if destination == "" {
				break
			}
			securityPoliciesString = "set security policies from-zone " + fromZone + " to-zone " + toZone + " policy " + policy + " match destination-address " + destination
			securityPoliciesSlice = append(securityPoliciesSlice, securityPoliciesString)
		}
		for {
			var application string
			fmt.Print("Application: ")
			fmt.Scanln(&application)
			if application == "" {
				break
			}
			securityPoliciesString = "set security policies from-zone " + fromZone + " to-zone " + toZone + " policy " + policy + " match application " + application
			securityPoliciesSlice = append(securityPoliciesSlice, securityPoliciesString)
		}
		securityPoliciesString = "set security policies from-zone " + fromZone + " to-zone " + toZone + " policy " + policy + " then permit"
		securityPoliciesSlice = append(securityPoliciesSlice, securityPoliciesString)
		securityPoliciesString = "set security policies from-zone " + fromZone + " to-zone " + toZone + " policy " + policy + " then log session-close"
		securityPoliciesSlice = append(securityPoliciesSlice, securityPoliciesString)
		fmt.Print("\n")
	}
}

func securityPoliciesOutput() {
	for _, d := range securityPoliciesSlice {
		fmt.Println(d)
	}
}

func main() {
	var option string
	fmt.Print(`https://github.com/james-borwick

┌───────────────────────────────┐
│ SRX Security Policies - v1.04 │
└───────────────────────────────┘

`)
	for {
		fmt.Print(`[1] Manual
[2] Auto

Option: `)
		fmt.Scanln(&option)
		fmt.Println()
		if option == "1" {
			applicationsInput()
			addressBookInput()
			securityPoliciesInput()
			fmt.Print("\n─────────────────────────────────────────────────────────────────────────────────────────────────────────────────\n\n")
			applicationsOutput()
			addressBookOutput()
			securityPoliciesOutput()
		}
		if option == "2" {
			autoPoliciesInput()
			fmt.Print("\n─────────────────────────────────────────────────────────────────────────────────────────────────────────────────\n\n")
			autoPoliciesOutput()
		}
		fmt.Print("#\n")
		fmt.Print("\n─────────────────────────────────────────────────────────────────────────────────────────────────────────────────\n\n")
		addressBookZoneSlice = nil
		addressBookNameSlice = nil
		addressBookIPSlice = nil
		applicationsNameSlice = nil
		applicationsProtocolSlice = nil
		applicationsPortSlice = nil
		securityPoliciesSlice = nil
		autoSecurityPolicy = nil
	}
}
