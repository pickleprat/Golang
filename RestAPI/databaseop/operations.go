package databaseop

import (
	"math/rand"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// structs required
type CompanyName struct {
	Name string `json:"name"`
}

type Products struct {
	Name string `json:"name"`
	Quantity int64 `json:"quantity"`
	NetValue int64 `json:"value"`
}


// reading configurations from the .env file 
func getConfig(dbname string) mysql.Config {

	// loading the environment variables 
	godotenv.Load("/home/pickleprat/Desktop/projects/backend/golang/RestAPI/databaseop/.env"); 

	cfg := mysql.Config{
		User: os.Getenv("DBUSER"), 
		Passwd: os.Getenv("DBPASSWD"), 
		Net: os.Getenv("DBNET"), 
		DBName: dbname, 
		Addr: os.Getenv("DBHOST"), 
	} 
	return cfg 
}

// connecting to the primary database 
func ConnectDB(dbname string) *sql.DB {

	// getting configurations and initializing pointers 
	cfg := getConfig(dbname); 
	var db *sql.DB 
    
	// getting the database pointer and handling error 
	db, err := sql.Open("mysql", cfg.FormatDSN()); 
	if err != nil {
		log.Fatalf("The database connection was interrupted because of : %q", err); 
	}
    
	// getting the database ready for pinging 
	pingErr := db.Ping(); 
	if pingErr != nil {
		log.Fatalf("The database connection failed at pinging because of : %q", pingErr); 
	}

	return db; 

}


// retreiving the data from the database 
func RetrieveProducts(cname CompanyName, db *sql.DB) ([] Products, error) {
	var products  [] Products; 

	// query the results and load it into a variable 
	query := "SELECT DISTINCT BRD, QTY, VALUE FROM Products WHERE MBRD LIKE ?; "
	queryResult, err := db.Query(query, cname.Name); 

	if err != nil {
		return nil, fmt.Errorf("An Error occured while querying: %v", err); 
	}

	// defer the closing of connection until the end of the function 
	defer queryResult.Close(); 

	// iterate over the results 
	for queryResult.Next(){

		// create a temporary struct to store your data[i]
		var product Products; 
        
		// Take the data and store it in the appropriate pointers 
		scanErr := queryResult.Scan(&product.Name, &product.Quantity, &product.NetValue); 

		if scanErr != nil {
			return nil, fmt.Errorf("An error occured while scanning: %q", scanErr); 
		}

		// storing in the products array 
		products = append(products, product); 

	}

	if err := queryResult.Err() ; err != nil {
		return nil, fmt.Errorf("Error occured post querying: %q", err); 
	}

	return products, nil; 

}

func retrieveCompanies(db *sql.DB) ([] CompanyName, error) {
	query := "SELECT DISTINCT MBRD FROM Products;"
	result, err := db.Query(query); 
	if err != nil {
		return nil,fmt.Errorf("Error occured while querying : %q", err); 
	}	
    var companies [] CompanyName; 

	defer result.Close(); 
	for result.Next(){ 

		var cname CompanyName; 
		scanErr := result.Scan(&cname.Name); 

		if scanErr != nil {
			return nil, fmt.Errorf("Error occured while scannning : %q", scanErr); 
		}

		companies = append(companies, cname); 
	}

	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("Post query error : %q", err)
	}

	return companies, nil 
}

func SelectRandomCompany(db *sql.DB) CompanyName { 
	companies, err := retrieveCompanies(db); 
	if err != nil {
		log.Fatal(err); 
	}
	return companies[rand.Intn(len(companies))]; 
}
