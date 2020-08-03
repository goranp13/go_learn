package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/arteev/firebirdsql"
)

type Response1 struct {
	ACT int
	MAX int
}

func main() {

	var xsql string
	var resp Response1

	//conn, err := sql.Open("firebirdsql", "sysdba:masterkey@127.0.0.1/EMPLOYEE.fdb")
	conn, err := sql.Open("firebirdsql", "sysdba:masterkey@127.0.0.1/proba")

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

	xsql = "SELECT ACTVALUE, MAX_VALUE FROM T_INSTCOUNTER WHERE KCOUNTER = '101'"

	rows, err := conn.Query(xsql)

	if err != nil {
		fmt.Println("query")
		fmt.Println(err)
		os.Exit(3)
	}

	//var res2B []byte

	for rows.Next() {

		rows.Scan(&resp.ACT, &resp.MAX)

		mv := &resp.MAX
		fmt.Println(*mv)
		av := &resp.ACT
		fmt.Println(*av)

		fmt.Println("Slobodno mjesta: ", *mv-*av)

		/* res2B, _ = json.Marshal(resp)
		fmt.Println(string(res2B)) */
	}

	fmt.Println("OK")
}
