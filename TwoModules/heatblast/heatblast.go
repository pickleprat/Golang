package heatblast 

import "fmt"

func FireBall() string {
	return "Smoke and Heat Phew Phew Phew!!\n"
}

func Fly() string {
	return "Am Flying now! Woooosh!\n"
}

func Run() string {
	var message string ; 
	message = "I never Run !"
	flyMessage := Fly(); 
	return fmt.Sprintf("%v, %v", message, flyMessage); 
}