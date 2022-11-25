package main

import (
	"database/sql"
	"empProject2/emp"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"

	"log"
	_ "net/http"
)

func main() {
	var err error
	emp.DB, err = sql.Open("mysql",
		"root:Aditi#2#@tcp(127.0.0.1:3306)/employee")
	if err != nil {
		log.Println(err)
		return
	}

	defer emp.DB.Close()

	err = emp.DB.Ping()
	if err != nil {
		log.Println(err)

		return
	}

	//employees, err := emp.GetEmployeeData()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	//oneEmployee, err := emp.GetOneEmployeeData(emp.DB, "dc652fdc-6a50-11ed-90d1-64bc589051b4")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//fmt.Println(oneEmployee)
	//fmt.Println(employees)

	router := mux.NewRouter()
	router.HandleFunc("/emp", emp.GetEmpData).Methods("GET")
	router.HandleFunc("/emp/{id}", emp.GetOneEmpData).Methods("GET")
	//router.HandleFunc("/emp", PostEmployeeData).Methods("POST")
	fmt.Println(("server at port 8000"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
