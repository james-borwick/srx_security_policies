package main

import "fmt"

var autoSecurityPolicy = make([]string, 0)

func autoPoliciesInput() {
	fmt.Print(`Steps
─────
1. Create a spreadsheet with the following columns...
    ┌───────────┬─────────┬────────┬────────┬─────────────┬─────────────┐
    │ From Zone │ To Zone │ Policy │ Source │ Destination │ Application │
    └───────────┴─────────┴────────┴────────┴─────────────┴─────────────┘
2. Fill out all of the rows
3. Copy and paste the cells below

Rules
─────
• No empty cells
• No spaces in the cells
• One item per cell
• Use # to ignore a cell

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
		// Lines
		var sourceLine string
		var destinationLine string
		var applicationLine string
		var permitLine string
		var logLine string
		// Scan
		fmt.Scanln(&from, &to, &policy, &source, &destination, &application)
		if from == "" {
			break
		}
		// Append the lines to the slice
		if source != "#" {
			sourceLine = "set security policies from-zone " + from + " to-zone " + to + " policy " + policy + " match source-address " + source
			autoSecurityPolicy = append(autoSecurityPolicy, sourceLine)
		}
		if destination != "#" {
			destinationLine = "set security policies from-zone " + from + " to-zone " + to + " policy " + policy + " match destination-address " + destination
			autoSecurityPolicy = append(autoSecurityPolicy, destinationLine)
		}
		if application != "#" {
			applicationLine = "set security policies from-zone " + from + " to-zone " + to + " policy " + policy + " match application " + application
			autoSecurityPolicy = append(autoSecurityPolicy, applicationLine)
		}
		permitLine = "set security policies from-zone " + from + " to-zone " + to + " policy " + policy + " then permit"
		autoSecurityPolicy = append(autoSecurityPolicy, permitLine)
		logLine = "set security policies from-zone " + from + " to-zone " + to + " policy " + policy + " then log session-close"
		autoSecurityPolicy = append(autoSecurityPolicy, logLine)
	}
}

func autoPoliciesOutput() {
	for _, c := range autoSecurityPolicy {
		fmt.Println(c)
	}
}
