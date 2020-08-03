package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/nakagami/firebirdsql"
)

type Response1 struct {
	COUNTRY  int
	CURRENCY string
}

func main() {

	var xsql string
	var resp Response1

	conn, err := sql.Open("firebirdsql", "sysdba:masterkey@127.0.0.1/EMPLOYEE.fdb")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	err = conn.Ping()

	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	xsql = "SELECT a.COUNTRY, a.CURRENCY FROM COUNTRY a"

	rows, err := conn.Query(xsql)

	if err != nil {
		fmt.Println("query")
		fmt.Println(err)
		os.Exit(3)
	}

	var res2B []byte

	for rows.Next() {

		rows.Scan(&resp.COUNTRY, &resp.CURRENCY)

		fmt.Println(&resp.COUNTRY)

		res2B, _ = json.Marshal(resp)
		fmt.Println(string(res2B))
	}

	fmt.Println("OK")
}
