package main
import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)


var jsonstr=`{
	"responseData": {
	"results": [
		{
			"GsearchResultClass": "GwebSearch",
			"unescapedUrl": "https://www.reddit.com/r/golang",
			"url": "https://www.reddit.com/r/golang",
			"visibleUrl": "www.reddit.com",
			"cacheUrl": "http://www.google.com/search?q=cache:W...",
			"title": "Reddit",
			"titleNoFormatting": "Golang - Reddit",
			"content": "First Open Source"
		},
		{
			"GsearchResultClass": "GwebSearch",
			"unescapedUrl": "http://tour.golang.org/","url": "http://tour.golang.org/",
			"visibleUrl": "tour.golang.org",
			"cacheUrl": "http://www.google.com/search?q=cache:O...",
			"title": "A Tour of Go",
			"titleNoFormatting": "A Tour of Go",
			"content": "Welcome to a tour of the Go programming ..."
		}]
	}
}`
func getuser(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//uid := params.Get(":uid")
	fmt.Fprint(w,jsonstr)
}
func modifyuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are modify user %s", uid)
}
func deleteuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are delete user %s", uid)
}
func adduser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprint(w, "you are add user %s", uid)
}
func main() {
	mux := routes.New()
	mux.Get("/user/:uid", getuser)
	mux.Post("/user/:uid", modifyuser)
	mux.Del("/user/:uid", deleteuser)
	mux.Put("/user/", adduser)
	http.Handle("/", mux)
	http.ListenAndServe(":8088", nil)
}