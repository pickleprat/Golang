package fourarms 

import (
	"fmt"
	"omnitrix/fourarms/Face"
)

func Punch() string {
	return "I'm four arms!! Bam Bam bam"
}

func Jump() string {
	return "I am jumping! Wohoo"
}

func Run() string {
	message := "Four arms doesn't RUN! It JUMPS!"
	return fmt.Sprintf("%v, %v", message, Jump()); 

}

func Blink(question string) string {
	eyes := face.Eyes();
	message := "I am not going to blink because "
	return fmt.Sprintf("%v, %v", message, eyes);  

}

