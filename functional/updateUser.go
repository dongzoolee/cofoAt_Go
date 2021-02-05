package functional

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sort"
)


// 순서 상관 없이
func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func ErrCheck(e error) {
	if e != nil {
		panic(e)
	}
}
func goDotEnvVar(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error Loading .env file")
	}
	return os.Getenv(key)
}
func AddUser(id string, channel string) {
	db, err := sql.Open("mysql", goDotEnvVar("MYSQL_ID")+":"+goDotEnvVar("MYSQL_PW")+"@tcp("+goDotEnvVar("MYSQL_HOST")+")/"+goDotEnvVar("MYSQL_DB"))
	ErrCheck(err)

	chn := new(Channels)
	chn.Channels = append(chn.Channels, channel)

	mshedChn, err := json.Marshal(chn)
	result, err := db.Query("INSERT INTO clients(handle, channels) VALUES(?, ?)", id, string(mshedChn))
	ErrCheck(err)
	defer db.Close()

	fmt.Println(result)
}
func AddChannel(id string, channel string) {
	db, err := sql.Open("mysql", goDotEnvVar("MYSQL_ID")+":"+goDotEnvVar("MYSQL_PW")+"@tcp("+goDotEnvVar("MYSQL_HOST")+")/"+goDotEnvVar("MYSQL_DB"))
	ErrCheck(err)

	row, err := db.Query(`SELECT channels FROM clients WHERE handle = ?`, id)
	ErrCheck(err)

	var data string
	for row.Next() {
		err = row.Scan(&data)
		ErrCheck(err)
	}

	unMshedData := new(Channels)
	json.Unmarshal([]byte(data), unMshedData)
	unMshedData.Channels = append(unMshedData.Channels, channel)

	mshedData, err := json.Marshal(unMshedData)
	ErrCheck(err)

	result, err := db.Query("UPDATE clients SET channels = ? WHERE handle = ?", string(mshedData), id)
	ErrCheck(err)
	defer db.Close()

	fmt.Println(result)
}
func DelChannel(id string, targetPos int) {
	db, err := sql.Open("mysql", goDotEnvVar("MYSQL_ID")+":"+goDotEnvVar("MYSQL_PW")+"@tcp("+goDotEnvVar("MYSQL_HOST")+")/"+goDotEnvVar("MYSQL_DB"))
	ErrCheck(err)
	row, err := db.Query(`SELECT channels FROM clients WHERE handle = ?`, id)
	ErrCheck(err)

	var data string
	for row.Next() {
		err = row.Scan(&data)
		ErrCheck(err)
	}
	unMshedData := new(Channels)
	json.Unmarshal([]byte(data), unMshedData)

	sort.StringSlice(unMshedData.Channels).Sort()
	unMshedData.Channels = remove(unMshedData.Channels, targetPos)

	mshedData, err := json.Marshal(unMshedData)
	ErrCheck(err)

	result, err :=db.Query(`UPDATE clients SET channels = ? WHERE handle = ?`, string(mshedData), id)
	ErrCheck(err)

	fmt.Println(result)
}
