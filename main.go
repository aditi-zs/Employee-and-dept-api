package main

import (
	"example.com/empProject/emp"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http"
)

func main() {
	var err error

	emp.DB, err = emp.DbConnection("mysql", "root:Aditi#2#@tcp(127.0.0.1:3306)/employee")
	if err != nil {
		log.Println(err)
		return
	}

	defer emp.DB.Close()

	//err = emp.DB.Ping()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	router := mux.NewRouter()
	router.HandleFunc("/emp", emp.GetEmpData).Methods("GET")
	router.HandleFunc("/emp/{id}", emp.GetOneEmpData).Methods("GET")
	router.HandleFunc("/postempdata", emp.PostEmployeeData).Methods("POST")
	router.HandleFunc("/postdepdata", emp.PostDepartmentData).Methods("POST")
	router.HandleFunc("/dept", emp.GetDepData).Methods("GET")
	fmt.Println(("server at port 8000"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
