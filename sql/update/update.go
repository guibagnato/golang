package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "docker:test123@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Testador", 1)
	stmt.Exec("Logador", 2)

	// delete
	stmt2, _ := db.Prepare("delete from  usuarios where id = ?")
	stmt2.Exec(3)
}
