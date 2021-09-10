package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
	"database/sql"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
	"encoding/json"
	//"database/sql"
)

type CountryMaster struct {
    CountryId    int
    CountryName  string
  }
 
type  Location struct {
    Cityname string
    State string
    Latitude float64
    Longitude float64
}
 
type WeatherData struct {
    WeatherID int64
    Location string 
    WeatherDate string 
    Temp string
}
 

func main() {

	fmt.Println("Docker Demo")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w,"Hello world. This is after running docker-compose up")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi There. This is after runnign docker compose up")
	})

	handleRequests()

}

func handleRequests() {
 
    	http.HandleFunc("/homepage", homePage)
	http.HandleFunc("/CountryList", CountryList)
    	log.Fatal(http.ListenAndServe(":8080", nil))
}



func homePage(w http.ResponseWriter, r *http.Request){
     	fmt.Fprintf(w,"Hello.. This is my first app")
 }

func dbConn() (db *sql.DB) {
	var condb *sql.DB
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")

	fmt.Println("Endpoint Hit: In dbConn connection.. after runnign docker compose up")
	fmt.Println("Usenamd =" + db_user + "pwd :" + db_password)
    	//condb, err := sql.Open("mssql", "Data Source=host.docker.internal,1433;database=WeatherDB;User ID=sa;Password=someThingComplicated1234")
	//condb, err := sql.Open("mssql", "Data Source=	172.17.0.2,1433;database=WeatherDB;User ID=sa;Password=someThingComplicated1234")
	//condb, err := sql.Open("mssql", "Data Source=localhost;database=WeatherDB;User ID=" + db_user + ";Password=" + db_password)
	//condb, err := sql.Open("mssql", "Data Source=db;database=WeatherDB;User ID=" + db_user + ";Password=" + db_password)
	condb, err := sql.Open("mssql", "Data Source=db,1433;database=WeatherDB;User ID=" + db_user + ";Password=" + db_password)

	if err = db.Ping(); err != nil {
		db.Close()
	fmt.Println("Endpoint Hit: Error in dbcon")
	}

	// Make sure to update the Password value below from "Your_password123" to your actual password.
	// var connection = @"Server=db;Database=master;User=sa;Password=Your_password123;";

	fmt.Println("Endpoint Hit: After sql.Open")


	if err != nil {
	log.Fatal(err)
	}
	return condb    
}
func CountryList(w http.ResponseWriter, r *http.Request){
    	fmt.Println("Endpoint Hit: In CountryList.. after running docker compose up")

    	ListofCountries:=GetCountryData()	
    
    	fmt.Println("Endpoint Hit: Data obtained from GetCountryData.....now converting to json.. docker compose up ")
	json.NewEncoder(w).Encode(ListofCountries)
}



func GetCountryData()[]CountryMaster {
	var ListofCountries []CountryMaster
	var eachrow CountryMaster

	fmt.Println("Endpoint Hit: in GetCountryData ")

	db := dbConn()
	fmt.Println("Endpoint Hit: In GetCountryData again")

	selDB, err := db.Query(" SELECT countryid, countryname FROM CountryMaster")
	if err != nil {
	panic(err.Error())
	}
	fmt.Println("Endpoint Hit: Query run on Countrymaster")

	for selDB.Next() {    
	err = selDB.Scan(&eachrow.CountryId, &eachrow.CountryName )
	if err != nil {
			log.Fatal(err)
	}         
	ListofCountries = append(ListofCountries, eachrow)
    }
    return ListofCountries
}



//docker run -w /go/src/app -it --link mysql55c -d --name golangapp -v $(pwd):/go/src/app golang bash -c "go get github.com/go-sql-driver/mysql;go build main.go; go test -v --config ./config.ini"
//SQL IP address             "IPAddress": "172.17.0.2",
//                        "HostPort": "1433"
// 842c6bfaea2c  -cleandemo 
// 56beb1db7406  - sql 2910 
