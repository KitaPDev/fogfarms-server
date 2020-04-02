package modulegroup_management

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/gddo/httputil/header"
)

const (
	DbHost = "localhost"
	DbPort = 5432
	DbUser = "postgres"
	DbPass = "postgres"
	DbName = "fogfarms-01"
)

var db *sql.DB

func GetDB() *sql.DB {
	var err error

	if db == nil {
		connectionString := fmt.Sprintf("port=%d host=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			DbPort, DbHost, DbUser, DbPass, DbName)

		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			panic(err)
		}
	}

	return db
}

type TestData struct {
	Name string
}

func GetTestName(w http.ResponseWriter, r *http.Request) {
	db := GetDB()

	rows, err := db.Query("SELECT * FROM Test")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var testDatas []TestData
	for rows.Next() {
		var name string

		err := rows.Scan(&name)
		if err != nil {
			panic(err)
		}

		var testData = TestData{
			Name: name,
		}

		testDatas = append(testDatas, testData)
	}
	fmt.Printf("%v", testDatas)

	js, err := json.Marshal(testDatas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func PostTestName(w http.ResponseWriter, r *http.Request) {
	var testData TestData
	if r.Header.Get("Content-Type") != "" {
        value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
        if value != "application/json" {
            msg := "Content-Type header is not application/json"
            http.Error(w, msg, http.StatusUnsupportedMediaType)
            return
        }
    }

	err := json.NewDecoder(r.Body).Decode(&testData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	js, err := json.Marshal(testData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
