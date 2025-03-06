package main

import (
	"encoding/json"
	"fmt"
)

// **1. Defining a Struct**  
// A struct is a collection of related fields.
type Person struct {
	Name string `json:"name"` // JSON tag to specify field name when encoding
	Age  int
}

// **2. Method for the Struct**  
// This method increases the Age field by 1 and prints a message.
func (p *Person) HappyBirthday() {
	p.Age++
	fmt.Println("Happy Birthday", p.Name)
}

// **3. Variables in Go**  

// Using `var` with explicit type
var name string = "Abdulrahman"
var age int = 22

// Using `var` with inferred type
var city = "Lagos"
var population = 200000000.00

// **4. Zero Values in Go**  
// If a variable is not initialized, Go assigns it a default zero value.
var last_name string // Default is an empty string ""

// **5. Constants in Go**  
// Constants store values that don't change.
const pi float64 = 3.14159

// Grouping multiple constants
const (
	Appname string = "GoStore"
	Version        = "1.0.0"
	MaxUser        = 100
)

// **6. Main Function**
func main() {
	// Using shorthand declaration `:=` (only inside functions)
	country := "Nigeria"
	gdp := 500.75

	fmt.Println(country, "is a country with a GDP of", gdp)

	// **7. Working with Structs**  

	// Creating struct instances
	var p1 Person // Zero-value struct (default values)
	p2 := Person{"Latifat", 30} // Using struct literal
	p3 := Person{Age: 21}       // Partial initialization

	// Display struct values
	fmt.Println(p1)        // Prints zero values
	fmt.Println(p2.Name)   // Accessing struct field

	// Calling a method on struct
	p3.HappyBirthday()

	// **8. JSON Encoding**  
	// Converts a struct to JSON format
	p3json, _ := json.Marshal(p2) 
	fmt.Println(string(p3json))   // Prints JSON output
}
