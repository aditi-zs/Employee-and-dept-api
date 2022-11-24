package emp

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var DB *sql.DB

type Department struct {
	DeptID   string `json:"dept_id"`
	DeptName string `json:"dept_name"`
}

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	PhoneNo string     `json:"phoneNo"`
	Dept    Department `json:"deptId"`
	//DeptID  Department `json:"deptID"`
}

func GetEmployeeData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := DB.Query("Select e.ID,e.Name,e.PhoneNo,e.DeptID,d.Name from emp e join dept d on e.DeptID=d.DeptID	")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var employees []Employee

	// 10 records
	for rows.Next() {
		var e Employee
		err = rows.Scan(&e.ID, &e.Name, &e.PhoneNo, &e.Dept.DeptID, &e.Dept.DeptName)
		if err != nil {
			panic(err.Error())
		}

		employees = append(employees, e)
	}
	w.WriteHeader(http.StatusOK)
	// to be discussed finally
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
	respBody, _ := json.Marshal(employees)
	w.Write(respBody)
	//return employees, nil
}

//	func GetEmpData(w http.ResponseWriter, r *http.Request) {
//		val, err := GetEmployeeData(DB)
//		if err != nil {
//			log.Println(err)
//		}
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//		err = json.NewEncoder(w).Encode(val)
//		if err != nil {
//			log.Println(err)
//		}
//		//respBody, _ := json.Marshal(val)
//		//w.Write(respBody)
//	}
func GetOneEmployeeData(db *sql.DB, id string) (Employee, error) {
	var e Employee
	row := db.QueryRow("Select e.ID,e.Name,e.PhoneNo,e.DeptID,d.Name from emp e join dept d on e.DeptID=d.DeptID WHERE id=?", id)

	err := row.Scan(&e.ID, &e.Name, &e.PhoneNo, &e.Dept.DeptID, &e.Dept.DeptName)
	if err != nil {
		return Employee{}, err
	}

	return e, nil
}

func GetOneEmpData(w http.ResponseWriter, r *http.Request) {
	empID := mux.Vars(r)["id"]
	val, err := GetOneEmployeeData(DB, empID)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")

	//for _, val := range employees {
	//	if (val.ID) == empID {
	//		json.NewEncoder(w).Encode(val)
	//
	//	}
	//}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(val)
	if err != nil {
		log.Println(err)
	}
}
