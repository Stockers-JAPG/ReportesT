package Config

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

func CreateConnection() (db *sql.DB, e error) {
	db, err := sql.Open("godror", "CUDGADM/Telcel01#@10.191.209.174:1522/CUDG")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Conexión exitosa a %s:%d \n", "10.191.209.174", "1522")
	}
	return db, nil
}
