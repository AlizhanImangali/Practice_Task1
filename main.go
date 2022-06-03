package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type SuccessResponse struct {
	Success bool
}

type Data struct {
	Id         int    `json:"id"`
	USER_ID    int    `json:"user_id"`
	FIRST_NAME string `json:"first_name"`
	LAST_NAME  string `json:"last_name"`
	BIIN       string `json:"biin"`
	EMAIL      string `json:"email"`
	PHONE      string `json:"phone"`
	PASSWRD    string `json:"passwrd"`
}
type PostResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func main() {
	values := map[string]string{
		"userId":    "0",
		"biin":      "string",
		"firstName": "string",
		"lastName":  "string",
		"password":  "string",
		"phone":     "string",
		"email":     "string",
	}
	json_data, err := json.Marshal((values))

	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("https://petstore.swagger.io/v2/user", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	//var res PostResponse

	var res PostResponse

	fmt.Println(res)
	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res)
	go Conn()
}
func Conn() {
	connStr := "user=postgres password=1234 dbname=Test sslmode=disabled"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("insert into s_users (userId,biin,firstName,lastName,password,phone,email) values ('0')", "string", "string", "string", "string", "string", "string")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())

	rows, err := db.Query("select * from Products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []Data{}

	for rows.Next() {
		d := Data{}
		err := rows.Scan(&d.Id, &d.BIIN, &d.FIRST_NAME, &d.LAST_NAME, &d.EMAIL, &d.PASSWRD, &d.PHONE, &d.USER_ID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, d)
	}
	for _, d := range products {
		fmt.Println(d.USER_ID, d.FIRST_NAME, d.LAST_NAME, d.EMAIL)
	}

}
