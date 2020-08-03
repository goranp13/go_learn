package main

import (
	"database/sql"
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

func main() {
	var n int
	//conn, _ := sql.Open("firebirdsql", "user:password@servername/foo/bar.fdb")
	conn, _ := sql.Open("firebirdsql", "SYSDBA:SACSDBA@127.0.0.1/AURA30DB.fdb")
	defer conn.Close()
	conn.QueryRow("SELECT Count(*) FROM rdb$relations").Scan(&n)
	fmt.Println("Relations count=", n)

}
