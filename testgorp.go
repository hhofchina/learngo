package main
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
	"fmt"
	"os"
	"reflect"
)

func perr(err error){
	if err != nil {
		//panic(err.Error())
		panic(err)
	}
}

type Test struct {
	Id int64 `db:"id"`
	Name string `db:"name",size:25`
	Cont string `db:"content"`
}

func initDb() *gorp.DbMap {
	db,err := sql.Open("mysql","user:123456@tcp(localhost:3306)/laraveldb?charset=utf8")
	perr(err)
	dbmap := &gorp.DbMap{Db:db,Dialect:gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Test{},"test").SetKeys(true,"Id")

	err=dbmap.CreateTablesIfNotExists()
	perr(err)
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	//dbmap.TraceOff()
	return dbmap
}
func testIt() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	err := dbmap.TruncateTables()
	perr(err)

	t1 := Test{
		Name:"hao",
		Cont:"haocontent",
	}
	t2 := Test{
		Name:"haohao",
		Cont:"haohaocontent",
	}
	dbmap.Insert(&t1, &t2)
	count, err := dbmap.SelectInt("SELECT count(*) FROM test")
	perr(err)
	fmt.Println("count:", count)
	var tests []Test
	if _, err := dbmap.Select(&tests, "SELECT * FROM test WHERE id>?", 0); err != nil {
		perr(err)
	}
	for _,t := range tests{
		fmt.Println("Select",t)
		var tt Test
		err:= dbmap.SelectOne(&tt,"SELECT * FROM test WHERE id=?",t.Id)
		perr(err)
		fmt.Println("SelectOne:",tt)
		o,err:=dbmap.Get(Test{},t.Id)
		perr(err)
		fmt.Println("GET",o.(*Test).Id,o.(*Test).Name)
	}
	var tests2 []Test
	_,err2:=dbmap.Select(&tests2,"SELECT * FROM test WHERE name=:name",map[string]interface{}{
		"name":"hao",
	})
	perr(err2)
	expected := Test{tests2[0].Id,"hao","haocontent"}
	if reflect.DeepEqual(expected,tests2[0]) {
		fmt.Println("is expect")
	} else {
		perr(fmt.Errorf("not expect"))
	}
}

func main() {
	testIt()
}
