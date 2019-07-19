package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func converttoint(a string) int64 {
	val, err := strconv.ParseInt(a, 10, 64)
	if err == nil {
		return val
	} else {
		return 0
	}
}

func Addproduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 7")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	category := vars["category"]
	name := vars["name"]
	mrp := converttoint(vars["mrp"])
	size := converttoint(vars["size"])
	discount := converttoint(vars["discount"])
	actualprice := float32(converttoint(vars["actualprice"]))
	sellerid := converttoint(vars["sellerid"])
	image := vars["image"]

	db.Create(&Products{Category: category, Name: name, Mrp: mrp, Size: size, Discount: discount, Actualprice: actualprice, Sellerid: sellerid, Image: image})
	defer db.Close()
}

func Getproductsdata(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 8")
	} else {
		fmt.Println("Connection Successfull")
	}

	var products []Products
	db.Find(&products)
	defer db.Close()
	json.NewEncoder(w).Encode(products)
}

func Addtocart(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 9")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	userid := converttoint(vars["userid"])
	productid := converttoint(vars["productid"])

	db.Create(&Carts{Userid: userid, Productid: productid})
	defer db.Close()
	/*db.Where("userid = ?", ).Find(&cart)*/
}

func Delcart(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 10")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	userid := converttoint(vars["userid"])
	productid := converttoint(vars["productid"])

	db.Where("userid = ? AND productid = ?", userid, productid).Unscoped().Delete(Carts{})
	defer db.Close()
}

func Getcart(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 11")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	userid := converttoint(vars["userid"])

	var cart []Carts
	var cartproducts []Products
	db.Where("userid = ?", userid).Find(&cart)

	for _, element := range cart {
		var product Products
		db.Where("ID = ?", element.Productid).Find(&product)
		if product.ID != 0 {
			cartproducts = append(cartproducts, product)
		}
	}
	defer db.Close()
	json.NewEncoder(w).Encode(cartproducts)

}

func Addtofav(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 12")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	userid := converttoint(vars["userid"])
	productid := converttoint(vars["productid"])

	db.Create(&Favs{Userid: userid, Productid: productid})
	/*db.Where("userid = ?", ).Find(&cart)*/
	defer db.Close()
}

func Delfav(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 13")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	userid := converttoint(vars["userid"])
	productid := converttoint(vars["productid"])

	db.Where("userid = ? AND productid = ?", userid, productid).Unscoped().Delete(Favs{})
	defer db.Close()
}

func Getfav(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 14")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	userid := converttoint(vars["userid"])

	var fav []Favs
	var favproducts []Products
	db.Where("userid = ?", userid).Find(&fav)

	for _, element := range fav {
		var fav Products
		db.Where("ID = ?", element.Productid).Find(&fav)
		if fav.ID != 0 {
			favproducts = append(favproducts, fav)
		}
	}
	json.NewEncoder(w).Encode(favproducts)
	defer db.Close()

}

func uploadimage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")

	r.ParseMultipartForm(5 << 20)

	file, handler, err := r.FormFile("myform")

	if err != nil {
		fmt.Println("Error uploading file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println(handler.Filename)
	fmt.Println(handler.Size)
	fmt.Println(handler.Header)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	ty := strings.Split(http.DetectContentType(fileBytes), "/")
	typ := ty[len(ty)-1:][0]
	fmt.Println(typ)
	tempfile, err := ioutil.TempFile("temp-images", "*."+typ)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tempfile.Name())

	var link = tempfile.Name()

	tempfile.Write(fileBytes)
	fmt.Fprintf(w, link)

	defer tempfile.Close()
	defer db.Close()

}

func Delproduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 15")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	sellerid := converttoint(vars["sellerid"])
	productid := converttoint(vars["productid"])

	db.Where("sellerid = ? AND id = ?", sellerid, productid).Unscoped().Delete(Products{})
	defer db.Close()
}

func Buyproduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 15")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	userid := converttoint(vars["userid"])
	sellerid := converttoint(vars["sellerid"])
	productid := converttoint(vars["productid"])

	db.Create(&Buy{Userid: userid, Sellerid: sellerid, Productid: productid})
	defer db.Close()

}

func Getallcart(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 15")
	} else {
		fmt.Println("Connection Successfull")
	}

	var cart []Carts
	db.Find(&cart)
	json.NewEncoder(w).Encode(cart)
	defer db.Close()
}
