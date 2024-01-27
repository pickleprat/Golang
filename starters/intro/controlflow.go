package intro

import "fmt"

func Flow() {
	// Switch cases in Go
	dead := false

	switch dead {
	case !dead:
		fmt.Printf("Yoooo Party, %T is aliveeeeeee\n", dead)

	case dead:
		fmt.Printf("Yooooo Party, %T is deaaaaad\n", dead)

	default:
		fmt.Printf("We just wanna party man!")
	}

	// if you don't know who farheena i.e. Fari is, you don't deserve to read this code
	host := "Prat"

	switch pretty := "Fari"; pretty {

	case "Fari":
		fmt.Printf("%v: CORRECT ANSWER\n", host)
		fallthrough

	case "Farheena":
		fmt.Printf("%v: YOU GOT IT!!\n", host)

	case "Fari boki":
		fmt.Printf("%v: IT IS CORRECT\n", host)

	default:
		fmt.Printf("%v: WRONG ANSWER\n", host)
	}

	// fallthrough keyword allows you to fall through and evaluate the next case
	// if there is a fallthrough in the next case as well it will fall through and evaluate the next to next case as well

	if psswd := "t@TeTittiE$"; psswd == "retard" {
		fmt.Printf("Get lost.")

	} else if psswd == "t@TeTittiE$" {
		fmt.Printf("ACCEPTABLE ANSWER\n")

	} else if psswd == "I am a cunt" {
		fmt.Println("ACCEPTABLE ANSWER")
		
	} else {
		fmt.Println("Bye.")
	}

}
