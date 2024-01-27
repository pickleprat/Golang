package main

import (
	"fmt"
	"getting-started/intro"
)

func main() {
	fmt.Println("Me: Hello, can I access your pp prat? ")		
	
	// the defer keyword 
	updateSalary := intro.SalaryDeterminer(1000.0, -0.3)
	defer fmt.Println(updateSalary(1.0)); 
	defer fmt.Println("Function called"); 
	defer fmt.Printf("%v updated..!\n", updateSalary(0.9)) ; 
    newSalary := updateSalary(0.1); 
	fmt.Printf("The new salary is : %v\n", newSalary)
}
