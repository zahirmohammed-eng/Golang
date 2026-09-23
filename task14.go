package main

// Error 1: Using Small Letters for Exported functions (Private vs Public)
// What caused the error: In Go, variables and fields starting with a lowercase letter are private to their package.

func employee() { // Lowercase 'employee' makes this private!
	var name string

}

// How to fix it: Change employee to Employee (capitalized) so it becomes public/exported across packages.

func Employee() {
	var name string

}
