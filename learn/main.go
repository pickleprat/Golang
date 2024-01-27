package main; 

import "fmt"; 
import "github.com/google/shlex"; 



func main(){
	fmt.Printf("Hello World"); 
	console, something := shlex.Split("one \"two three\" four"); 
	for i := 0; i < 3; i++{
		fmt.Println(console[i]); 
	}

	fmt.Println(something); 
}