package main 

import (
	"omnitrix/heatblast"
	"omnitrix/fourarms"
	"fmt"
)

func main(){
	fmt.Print(heatblast.Run()); 
	fmt.Println(fourarms.Run()); 
	fmt.Println(fourarms.Blink("Nope")); 

}