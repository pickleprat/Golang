package randomf;

import (
    "fmt"
    "math/rand"
	"errors"
) 

func Rough(name string) { 
	fmt.Println("Hello!"); 
}

func SelectRandomString(stringArray []string) (string, error) {
	if len(stringArray) == 0{
		return "", errors.New("Array is empty"); 
	}
	var index int; 
	index = rand.Intn(len(stringArray)); 
	return stringArray[index], nil 
} 


