package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-oci8"
)

func main() {
	sRetval := ""

	fmt.Println("Hello OCI8")

	db, err := sql.Open("oci8", "xxxxx/xxx@xx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	fmt.Println("successfully connected to oracle server using dsn:")

	_, err = db.Exec("begin pk_testgo.getproducts(:retval); end;",
		sql.Named("retval", sql.Out{Dest: &sRetval}))
	println(string(sRetval))
	println(err)
}
