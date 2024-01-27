package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
) 

func main(){

	var name string = "Anam Malik"; 
	var integer_i int = 12345; 
	name_reverse := reverse.String(name);
	integer_rev := reverse.Int(integer_i); 
	fmt.Printf("The reversed name is : %v\n", name_reverse);
	fmt.Printf("The reversed integer is : %v\n", integer_rev); 

} 


