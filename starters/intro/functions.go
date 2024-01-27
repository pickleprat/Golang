package intro 

// declaring functions in golang 
import "fmt"

// shorthand syntax 
func GeneralFunction(name, date string, salary, taxes float32) string {
	details := fmt.Sprintf("Name   : %v\nDate   : %v\nSalary : %v\nTaxes  : %v\n", name, date, salary, taxes)
	return details; 

}

// closure functions 
func SalaryDeterminer(salary, multiplier float32) func(float32) float32 {
	base := salary * multiplier 
	fmt.Println("Updating dawg"); 
	
    // the returned function is allowed to use the local existing variables 
	return func(profit_margin float32) float32 {
		salary := base + salary * profit_margin; 
		return salary 
	}

}

// variadic functions
func Product(vector ...float32) float32{

	product := float32(1.0)
	for _, num := range vector {
		product = product * num; 
	}
	return product 
}

// init function runs before the main function 
func init(){
	fmt.Println("I am an ultra virgin")
}

// init function can be declared more than once 
func init(){
	fmt.Println("I am bigger virgin than you")
}

