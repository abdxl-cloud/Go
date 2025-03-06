package main

import (
	"fmt"
	// "encoding/json"
)

// Defining a struct
type Person struct {
	Name string
	Age  int
}

// Variables in Go

// Using var with explicit type
var name string = "Abdulrahman"
var age int = 22

// Using var with inferred type
var city = "Lagos"
var population = 200000000.00

// Zero Values in Go
var last_name string

// Constants in Go
const pi float64 = 3.14159

const (
	Appname string = "GoStore"
	Version        = "1.0.0"
	MaxUser        = 100
)

func main() {
	// Using shorthand declartation (:=), only in functions
	country := "Nigeria"
	gdp := 500.75
	var p1 Person
	p1.Name = "Latifat"

	fmt.Println(country, "is a country with a GDP of", gdp)
}
