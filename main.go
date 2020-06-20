package main

import "github.com/wat10/finalexam/customer"

func main() {
	r := customer.SetupRouter()
	r.Run(":2019")
}
