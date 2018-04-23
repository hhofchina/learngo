package main
import(
	"github.com/gorilla/context"
	"net/http"
)


type key int

const MyKey key = 0

func MyHandler(w http.ResponseWriter, r *http.Request) {
	// val is "bar".
	val := context.Get(r, MyKey)

	val, ok := context.GetOk(r, foo.MyKey)
}
// GetMyKey returns a value for this package from the request values.
func GetMyKey(r *http.Request) SomeType {
	if rv := context.Get(r, mykey); rv != nil {
		return rv.(SomeType)
	}
	return nil
}

// SetMyKey sets a value for this package in the request values.
func SetMyKey(r *http.Request, val SomeType) {
	context.Set(r, mykey, val)
}
func main()  {

	context.Set(r, MyKey, "bar")
}
