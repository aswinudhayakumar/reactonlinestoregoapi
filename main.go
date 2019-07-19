package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

var db *gorm.DB
var err error

type Users struct {
	gorm.Model

	Accounttype string
	Name        string
	Email       string
	Password    string
	Gid         string
}

type Products struct {
	gorm.Model

	Category    string
	Name        string
	Image       string
	Mrp         int64
	Size        int64
	Discount    int64
	Actualprice float32
	Sellerid    int64
}

type Carts struct {
	gorm.Model

	Userid    int64
	Productid int64
}

type Favs struct {
	gorm.Model

	Userid    int64
	Productid int64
}

type Buy struct {
	gorm.Model

	Userid    int64
	Sellerid  int64
	Productid int64
}

func InitialMigration() {

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect")
	} else {
		fmt.Println("Connected successfully")
	}

	defer db.Close()

	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Products{})
	db.AutoMigrate(&Carts{})
	db.AutoMigrate(&Favs{})
	db.AutoMigrate(&Buy{})

}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Helloworld")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", helloworld).Methods("GET")
	myRouter.HandleFunc("/users", Allusers).Methods("GET")
	myRouter.HandleFunc("/user", Newuser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", Deleteuser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", Updateuser).Methods("PUT")
	myRouter.HandleFunc("/signin/{email}/{password}", Signin).Methods("POST")
	myRouter.HandleFunc("/signinwithgoogle", Signinwithgoogle).Methods("POST")
	myRouter.HandleFunc("/sendmail", Sendamail).Methods("GET")

	myRouter.HandleFunc("/addproduct/{category}/{name}/{mrp}/{size}/{discount}/{actualprice}/{sellerid}/{image}", Addproduct).Methods("POST")
	myRouter.HandleFunc("/products", Getproductsdata).Methods("GET")
	myRouter.HandleFunc("/deleteproduct/{sellerid}/{productid}", Delproduct).Methods("POST")

	myRouter.HandleFunc("/upload", uploadimage).Methods("POST")

	myRouter.HandleFunc("/addtocart/{userid}/{productid}", Addtocart).Methods("POST")
	myRouter.HandleFunc("/getcart/{userid}", Getcart).Methods("POST")
	myRouter.HandleFunc("/getallcart", Getallcart).Methods("GET")
	myRouter.HandleFunc("/delcart/{userid}/{productid}", Delcart).Methods("POST")

	myRouter.HandleFunc("/buynow/{userid}/{sellerid}/{productid}", Buyproduct).Methods("GET")

	myRouter.HandleFunc("/addtofav/{userid}/{productid}", Addtofav).Methods("POST")
	myRouter.HandleFunc("/getfav/{userid}", Getfav).Methods("POST")
	myRouter.HandleFunc("/delfav/{userid}/{productid}", Delfav).Methods("POST")

	myRouter.PathPrefix("/temp-images/").Handler(http.StripPrefix("/temp-images/", http.FileServer(http.Dir("temp-images"))))

	log.Fatal(http.ListenAndServe(":8123", cors.Default().Handler(myRouter)))
}

func main() {
	fmt.Println("Started")
	InitialMigration()
	handleRequests()
}
