package main; 

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"github.com/joho/godotenv"
	"os"
	"math/rand"
)

func DisplayNames(companies [] CompanyName ){
	for idx, name := range companies{
		fmt.Printf("%v: %v\n", idx+1, name.Name); 
	}
}

func ConnectDB(dbname string) *sql.DB { 
	godotenv.Load(); 
	var db *sql.DB;
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"), 
		Passwd: os.Getenv("DBPASSWD"), 
		Net: os.Getenv("DBNET"), 
		DBName: dbname, 
		Addr: os.Getenv("DBHOST"), 
	} 

	var err error; 
	db, err = sql.Open("mysql", cfg.FormatDSN()); 
	if err != nil {
		log.SetPrefix("Configuration Error\n"); 
		log.Fatal(err); 
		log.SetFlags(1);
	} 

	ping := db.Ping(); 
	if ping != nil {
		log.SetPrefix("Ping Error\n"); 
		log.Fatal(ping); 
		log.SetFlags(1); 
	} 

	fmt.Printf("The Database %v is connected!\n", dbname); 

	return db; 
}

type CompanyName struct {
	Name string 
}


func SelectRandomCompany(companies [] CompanyName) CompanyName {
	companyLength:= len(companies); 
	return companies[rand.Intn(companyLength)]; 
}

type CompanyProduct struct {

	// Companies struct 
	Name string
	Count int32
	Quantity int32 

}

func GetCompanyArray(db *sql.DB) ([] CompanyName, error) {
	var query string = "SELECT DISTINCT MBRD FROM Products;"
	var companies [] CompanyName; 

	company_rows, err := db.Query(query);
	if  err != nil{
		return nil, fmt.Errorf("The program crashed because %q", err); 
	}

	defer company_rows.Close(); 

	for company_rows.Next(){
		var name CompanyName; 
		if err := company_rows.Scan(&name.Name); err != nil{
			return nil, fmt.Errorf("The row threw the following error: %v", err); 
		}

		companies = append(companies, name); 
	}

	if err := company_rows.Err(); err != nil{
		return nil, fmt.Errorf("The process did not successfully complete because : %v", err); 
	}

	return companies, nil; 
}

func FetchCompanyDetails(db *sql.DB, companyName CompanyName) ([] CompanyProduct, error) {
	var query string = fmt.Sprintf("SELECT BRD, QTY, VALUE FROM Products WHERE MBRD LIKE %q", companyName.Name); 
	var productList []CompanyProduct; 

	// get the query 
	queryResult, err := db.Query(query); 
	if err != nil{
		return nil, fmt.Errorf("Error occured while querying: %v", err); 
	}

	// Now iterate through the rows of the resulted query 
	defer queryResult.Close(); 

	for queryResult.Next(){
		var product CompanyProduct; 

		// Scanning takes the address of the parameter of the struct and then set in the data recieved from 
		// the table. 

		if err := queryResult.Scan(&product.Name, &product.Count, &product.Quantity); err != nil{
			return nil, fmt.Errorf("The fetching unexpectedly stopped : %v", err); 
		}

		productList = append(productList, product); 
	}

	if err := queryResult.Err(); err != nil {
		return nil, fmt.Errorf("An error occured after the fetching: %v", err); 
	}

	return productList, nil; 

}

func insertIntoCompanies(db *sql.DB, companyName string, quantity, values int32) (int64, error) {
	// creating a struct to store the data 
	product := CompanyProduct{
		Name: companyName, 
		Quantity: quantity,
		Count: values, 
	}
    
	// insertion command using db.Exec 
	result, err := db.Exec(
		"INSERT INTO Companies (MBRD, VALUE, QUANTITY) VALUES (?, ?, ?)", 
		&product.Name, &product.Count, &product.Quantity, 
	)

	if err != nil {
		return 0, fmt.Errorf("An error occured : %v", err); 
	}
    
	// returns int64 object 
	id, err := result.LastInsertId(); 

	if err != nil {
		return 0, fmt.Errorf("An error occured while getting insert id: %v", err); 
	}

	return id, nil; 



}


func LoadProducts(){

	var dbname string = "Transaction"; 
	db := ConnectDB(dbname); 
    companyNames, err := GetCompanyArray(db); 

	if err != nil {
		log.Fatalf("An error occured: %v", err); 
	}

	// select a random company 
	company := SelectRandomCompany(companyNames); 
	products, err := FetchCompanyDetails(db, company); 

	if err != nil {
		log.Fatal(err); 
	}

	// View the data now 
	fmt.Printf("Name of Product : %v\n", company.Name); 
	for idx, product := range products{
		fmt.Printf("%v: Product: %v [Quantity: %v  Value: %v]\n", idx+1, product.Name, product.Quantity, product.Count); 
	}

} 

func main(){
	var dbname string = "Transaction"; 
	db := ConnectDB(dbname); 

	id, err := insertIntoCompanies(db, "FAR SQUARE", 34, 55); 

	if err != nil {
		log.Fatal(err); 
	}

	fmt.Printf("Inserted into table Companies at id: %v\n", id)


}
