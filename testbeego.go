package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "database/sql"
)
func init(){
	fmt.Println("in init")
}

func testBeego()  {
	beego.Run()
}

func main() {
	testBeego()
}
