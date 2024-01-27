package intro

import "fmt"

func Introduction() {
	// Introduction to the fmt package 
	fmt.Println("Hello, here I am virgin prat. ")

	// difference between fmt.Println and fmt.Print
	fmt.Print("Hello", "I", "am", "virgin", "prat.")

	fmt.Println("Hello", "I", "am", "virgin", "prat.")

    //codes for formatting 
	//%v stores any variable 
	fmt.Printf("%v : Hey, I really like you. Would you like to go out with me?\n", "virgin")

	// %T stores the type of the variable 
	fmt.Printf("%T : Hello I am not your type\n", 98)

	fmt.Printf("%#v: How sure are you?\n", 32)

	// '%%' is used to show the '%' symbol 
	var percent float32 = 69.0/73 * 100
	fmt.Printf("I am a %v %% sure\n", percent)

	// few interesting Integer formatting codes 
	
    // '%b' for binary conversion 
	fmt.Printf("%v: I don't understand decimal, can you tell me in Binary?\n", "virgin")
    fmt.Printf("%T: Okay, I am %b %% sure\n", 32, uint(percent)) 

	// Here are the rest of them for reference
	
	//Integer 

	// %b	base 2
	// %c	the character represented by the corresponding Unicode code point
	// %d	base 10
	// %o	base 8
	// %O	base 8 with 0o prefix
	// %q	a single-quoted character literal safely escaped with Go syntax.
	// %x	base 16, with lower-case letters for a-f
	// %X	base 16, with upper-case letters for A-F
	// %U	Unicode format: U+1234; same as "U+%04X"
    
	// link for more 
	// https://pkg.go.dev/fmt


} 

func init(){
	fmt.Println("Who let the dogs out")
}