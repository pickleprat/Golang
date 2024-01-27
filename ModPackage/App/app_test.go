package main; 

import "testing" 

func TestRandomAdjectives(t *testing.T){
	names := []string {
		"Anam", "Bianca", 
		"Farheena", "Gautam", 
		"Jiya", "Camillie", 
	} 

	adjectives := [] string {
		"Gay", "Ugly", 
		"Pretty", "Bitchass", 
		"Hot", "Disgusting", 
		"Passive Agrressive", 
		"Sucks at sarcasm", 
		"Dumb", "Beautiful", 
	}

	messages, err := GenerateRandomAdjectives(names, adjectives);
	count := len(messages); 
	if err != nil {
		t.Fatal("Error raised ! Test failed because empty string was encountered"); 
	}
	
	if count < 1{
		t.Fatal("Error raised: Keys are repeated"); 
	} 
} 


