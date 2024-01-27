package main 

import (
   "fmt"
   "example.com/randomf"
   "log"
)

func GenerateRandomAdjectives(names [] string, adjectives [] string) (map[string]string, error) {
	var messages = make(map[string] string) 
	for _, name := range names{
		description, err := randomf.SelectRandomString(adjectives); 
		messages[name] = description;
		if err != nil {
			log.SetPrefix("Omg an Error"); 
			log.SetFlags(0); 
			log.Fatal(err); 
		} 
      
	} 
	return messages, nil; 
} 

func DisplayMap(messages map[string] string){
	for name, description := range messages{
		fmt.Printf("%v : %v\n", name, description); 
	} 
} 


func main() {
   log.SetPrefix("Empty string error"); 
   log.SetFlags(0); 
   var sents = []string{"Hello", "Noob", "Idiot", "Virgin", "Loser"}; 
   var messageMap = make(map[string] string);  
   var names = []string{"Anam", "Farheena", "Bianca", "Bibi", "Biob"}; 

   messageMap, _ = GenerateRandomAdjectives(names, sents); 
   DisplayMap(messageMap); 

}

