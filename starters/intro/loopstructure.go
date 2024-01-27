package intro 

import "fmt"

func Loops(){

	// similar to other programming languages 
	for i := 0; i < 10; i++{
		fmt.Println("Anam is retardedddd!!!")
	}

	// similar use of continue and break functions 
	for i := 0; i < 10; i++ {
		if i < 2 {
			continue
		}

		fmt.Printf("%v\t", i)

		if i > 5 {
			break
		}
	}

	fmt.Println("We broke out!")




}