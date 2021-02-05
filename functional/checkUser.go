package functional

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sort"

	_ "github.com/go-sql-driver/mysql"
)

type Channels struct {
	Channels []string
}

func UserExists(id string) bool {
	db, err := sql.Open("mysql", goDotEnvVar("MYSQL_ID")+":"+goDotEnvVar("MYSQL_PW")+"@tcp("+goDotEnvVar("MYSQL_HOST")+")/"+goDotEnvVar("MYSQL_DB"))
	ErrCheck(err)
	rows, err := db.Query(`SELECT EXISTS (SELECT * FROM clients WHERE handle = ?) AS chk`, id)
	ErrCheck(err)
	defer db.Close()

	var ret bool
	for rows.Next() {
		err = rows.Scan(&ret)
		ErrCheck(err)
		fmt.Println(ret)
		return ret
	}
	return false
}
func ChannelExists(id string, channel string) int {
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
	json.Unmarshal([]byte(data), &unMshedData)
	sort.StringSlice(unMshedData.Channels).Sort()
	var targetPos int = sort.StringSlice(unMshedData.Channels).Search(channel)

	if targetPos < len(unMshedData.Channels) && unMshedData.Channels[targetPos] == channel {
		return targetPos
	}
	return -1
}
