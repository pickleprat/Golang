package main

import (
	"database/sql"
	"net/http"
	// "fmt"
	"log"
	"example.com/restfulapi/databaseop"
    "github.com/gin-gonic/gin"
)

var db *sql.DB; 
var postedProducts [] databaseop.Products; 

func getProducts(db *sql.DB) ([] databaseop.Products, error) {
	// establishing a connection and getting data 
	cname := databaseop.SelectRandomCompany(db);

	// getting a products array from the database 
	products, err := databaseop.RetrieveProducts(cname, db);  
	if err != nil {
		return nil, err 
	} 

	return products,nil; 
}

// converting the products into json and storing it in the context gin 
func serveProducts(c *gin.Context){
	products, _ := getProducts(db); 
	c.IndentedJSON(http.StatusOK, products); 
}

func postProducts(c *gin.Context){ 
	var product databaseop.Products; 
	if err := c.BindJSON(&product); err != nil {
		return 
	}
	
	query := "INSERT INTO Companies(MBRD, VALUE, QUANTITY) VALUES (?, ?, ?)" 
	_ , err := db.Exec(query, &product.Name, &product.NetValue, &product.Quantity); 
	if err != nil {
		log.Fatal(err); 
	} 
    

	c.IndentedJSON(http.StatusCreated, product); 
}

func main(){
	db = databaseop.ConnectDB("Transaction"); 
	router := gin.Default(); 
	router.GET("/products", serveProducts);
	router.POST("/products", postProducts); 
	router.Run("localhost:8080"); 

}



