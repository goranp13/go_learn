package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/nakagami/firebirdsql"
)

type Response1 struct {
	MAX_VALUE int
	ACTVALUE  int
}

func main() {
	fmt.Print("hello 123\n")
	var xsql string
	var resp Response1

	conn, err := sql.Open("firebirdsql", "sysdba:masterkey@localhost/proba")

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

	xsql = "SELECT MAX_VALUE, ACTVALUE FROM T_INSTCOUNTER  WHERE KCOUNTER = '101'"

	rows, err := conn.Query(xsql)

	if err != nil {
		fmt.Println("query")
		fmt.Println(err)
		os.Exit(3)
	}

	var res2B []byte

	for rows.Next() {

		rows.Scan(&resp.MAX_VALUE, &resp.ACTVALUE)
		fmt.Printf("actual value %v\n", resp.ACTVALUE)
		fmt.Println(resp.MAX_VALUE - resp.ACTVALUE)

		res2B, _ = json.Marshal(resp)

		fmt.Println(string(res2B))
	}

	fmt.Println("OK")
}
