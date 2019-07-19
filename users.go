package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"github.com/domodwyer/mailyak"
	"github.com/gbrlsnchs/jwt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Allusers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 1")
	} else {
		fmt.Println("Connection Successfull")
	}
	var users []Users
	db.Find(&users)
	defer db.Close()
	json.NewEncoder(w).Encode(users)

}

type CustomPayload struct {
	jwt.Payload
	Name    string
	Email   string
	Userid  uint
	Account string
}

func Signin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 2")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	email := vars["email"]
	password := vars["password"]

	var users Users

	db.Where("Email = ?", email).Find(&users)
	defer db.Close()
	if users.ID != 0 && users.Password == password {
		var accounttype = users.Accounttype

		var now = time.Now()
		var hs256 = jwt.NewHMAC(jwt.SHA256, []byte("secret"))
		var h = jwt.Header{
			Type: accounttype,
		}
		var p = CustomPayload{
			Payload: jwt.Payload{
				Issuer:         "gbrlsnchs",
				Subject:        "someone",
				Audience:       jwt.Audience{"https://golang.org", "https://jwt.io"},
				ExpirationTime: now.Add(24 * 30 * 12 * time.Hour).Unix(),
				NotBefore:      now.Add(30 * time.Minute).Unix(),
				IssuedAt:       now.Unix(),
				JWTID:          "foobar",
			},
			Name:    users.Name,
			Email:   users.Email,
			Userid:  users.ID,
			Account: users.Accounttype,
		}
		token, err := jwt.Sign(h, p, hs256)
		if err != nil {
			// Handle error.
		}
		fmt.Fprintf(w, string(token))
		log.Printf("token = %s", token)
	} else {
		fmt.Fprintf(w, "null")
	}

}

func Newuser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 3")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := json.NewDecoder(r.Body)
	var u Users
	err = vars.Decode(&u)
	name := u.Name
	email := u.Email
	fmt.Println("hello", email)
	accounttype := u.Accounttype
	password := u.Password
	gid := u.Gid

	var user Users
	db.Where("Email = ?", email).Find(&user)
	fmt.Println("hi", user)

	if user.Name == "" {

		db.Create(&Users{Name: name, Email: email, Password: password, Accounttype: accounttype, Gid: gid})
		defer db.Close()
		sendmail(email)
		fmt.Println(w, "New user created successfully")
		fmt.Fprintf(w, "New user created successfully")

	}

}

func Deleteuser(w http.ResponseWriter, r *http.Request) {

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 4")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	name := vars["name"]

	var user Users
	db.Where("name=?", name).Find(&user)
	db.Delete(user)
	defer db.Close()
	fmt.Fprintf(w, "User deleted successfully")

}

func Updateuser(w http.ResponseWriter, r *http.Request) {

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 5")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user Users
	db.Where("name=?", name).Find(&user)
	user.Email = email
	db.Save(&user)
	defer db.Close()
	fmt.Fprintf(w, "User Updated successfully")

}

func Getusers(w http.ResponseWriter, r *http.Request) {

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 6")
	} else {
		fmt.Println("Connection Successfull")
	}

	var user Users
	db.Find(&user)
	defer db.Close()

}

func Signinwithgoogle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")

	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=test_go password=helloworld sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect 2")
	} else {
		fmt.Println("Connection Successfull")
	}

	vars := json.NewDecoder(r.Body)
	var user Users
	err1 := vars.Decode(&user)
	fmt.Println(err1)
	fmt.Println("hello ", user.Email)

	if user.Email != "" {
		email := user.Email
		gid := user.Gid
		fmt.Println("hello ", email, gid)
		var users Users

		db.Where("Email = ?", email).Find(&users)

		fmt.Println("hi ", users)

		if users.ID != 0 && users.Gid == "" && gid != "" {
			users.Gid = gid
			db.Save(&users)
			defer db.Close()
		}

		if users.ID != 0 && users.Gid == gid {
			var accounttype = users.Accounttype

			var now = time.Now()
			var hs256 = jwt.NewHMAC(jwt.SHA256, []byte("secret"))
			var h = jwt.Header{
				Type: accounttype,
			}
			var p = CustomPayload{
				Payload: jwt.Payload{
					Issuer:         "gbrlsnchs",
					Subject:        "someone",
					Audience:       jwt.Audience{"https://golang.org", "https://jwt.io"},
					ExpirationTime: now.Add(1).Unix(),
					NotBefore:      now.Add(1 * time.Minute).Unix(),
					IssuedAt:       now.Unix(),
					JWTID:          "foobar",
				},
				Name:    users.Name,
				Email:   users.Email,
				Userid:  users.ID,
				Account: users.Accounttype,
			}
			token, err := jwt.Sign(h, p, hs256)
			if err != nil {
				// Handle error.
			}
			fmt.Fprintf(w, string(token))
			log.Printf("token = %s", token)
		} else {
			fmt.Fprintf(w, "null")
		}
	}
}

func Sendamail(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")
	sendmail("aswiniu03@gmail.com")
	return

}

func sendmail(email string) {
	fmt.Println("hi")
	mail := mailyak.New("smtp.gmail.com:587", smtp.PlainAuth("", "aswiniu03@gmail.com", "hhkkdeftahyieblg", "smtp.gmail.com"))
	fmt.Println("sending mail to ", email)
	mail.To(email)
	mail.From("aswiniu03@gmail.com")
	mail.Subject("Welcome to Reactonlinestore")
	mail.HTML().Set(
		"<h1>Welcome to reactonlinestore, you have successfully Signed up to the reactonlinestore</h1>" + "\r\n" +
			"<h2>You can sign in using this mail ID at our reactonlinestore from now</h2>" + "\r\n" +
			"<a href='http://localhost:3000/signin'><button style='padding:10px;border:none;background:#1B9CFC;border-radius:5px;' > Signin </button></a>")

	if err := mail.Send(); err != nil {
		fmt.Println(err)
		panic(" :bomb: ")
	}
}
