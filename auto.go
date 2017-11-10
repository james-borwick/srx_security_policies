package main

import "fmt"

var autoSecurityPolicy = make([]string, 0)

func autoPoliciesInput() {
	fmt.Print(`Steps
─────
1. Create a spreadsheet with the following columns...
┌───────────────────────────────────────────────────────────────────┬──────────────────┬────────────────────────┐
│ Security Policy                                                   │ Address Book     │ Application            │
├───────────┬─────────┬────────┬────────┬─────────────┬─────────────┼──────┬──────┬────┼──────┬──────────┬──────┤
│ A         │ B       │ C      │ D      │ E           │ F           │ G    │ H    │ I  │ J    │ K        │ L    │
├───────────┼─────────┼────────┼────────┼─────────────┼─────────────┼──────┼──────┼────┼──────┼──────────┼──────┤
│ From Zone │ To Zone │ Policy │ Source │ Destination │ Application │ Zone │ Name │ IP │ Name │ Protocol │ Port │
└───────────┴─────────┴────────┴────────┴─────────────┴─────────────┴──────┴──────┴────┴──────┴──────────┴──────┘
2. Fill out all of the rows
3. Copy and paste the cells below

Rules
─────
No empty cells
No spaces in the cells
One item per cell
Use # to ignore a cell

Paste cells...
`)
	for {
		// Cells
		var from string
		var to string
		var policy string
		var source string
		var destination string
		var application string
		var addZone string
		var addName string
		var addIp string
		var appName string
		var appProtocol string
		var appPort string
		// Lines
		var sourceLine string
		var destinationLine string
		var applicationLine string
		var permitLine string
		var logLine string
		var addBookLine string
		var appProtocolLine string
		var appPortLine string
		// Scan
		fmt.Scanln(&from, &to, &policy, &source, &destination, &application, &addZone, &addName, &addIp, &appName, &appProtocol, &appPort)
		if from == "" {
			break
		}
		// Append the lines to the slice
		if source != "#" && from != "#" && to != "#" && policy != "#" {
			sourceLine = "set security policies from-zone " + from + " to-zone " + to + " policy " + policy + " match source-address " + source
			autoSecurityPolicy = append(autoSecurityPolicy, sourceLine)
		}
		if destination != "#" && from != "#" && to != "#" && policy != "#" {
			destinationLine = "set security policies from-zone " + from + " to-zone " + to + " policy " + policy + " match destination-address " + destination
			autoSecurityPolicy = append(autoSecurityPolicy, destinationLine)
		}
		if application != "#" && from != "#" && to != "#" && policy != "#" {
			applicationLine = "set security policies from-zone " + from + " to-zone " + to + " policy " + policy + " match application " + application
			autoSecurityPolicy = append(autoSecurityPolicy, applicationLine)
		}
		if from != "#" && to != "#" && policy != "#" {
			permitLine = "set security policies from-zone " + from + " to-zone " + to + " policy " + policy + " then permit"
			autoSecurityPolicy = append(autoSecurityPolicy, permitLine)
			logLine = "set security policies from-zone " + from + " to-zone " + to + " policy " + policy + " then log session-close"
			autoSecurityPolicy = append(autoSecurityPolicy, logLine)
		}
		if addZone != "#" && addName != "#" && addIp != "#" {
			addBookLine = "set security zones security-zone " + addZone + " address-book address " + addName + " " + addIp
			autoSecurityPolicy = append(autoSecurityPolicy, addBookLine)
		}
		if appName != "#" && appProtocol != "#" {
			appProtocolLine = "set applications application " + appName + " protocol " + appProtocol
			autoSecurityPolicy = append(autoSecurityPolicy, appProtocolLine)
		}
		if appName != "#" && appPort != "#" {
			appPortLine = "set applications application " + appName + " destination-port " + appPort
			autoSecurityPolicy = append(autoSecurityPolicy, appPortLine)
		}

	}
}

func autoPoliciesOutput() {
	for _, c := range autoSecurityPolicy {
		fmt.Println(c)
	}
}
