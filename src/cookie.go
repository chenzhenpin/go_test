package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("sayhelloName-------")
	r.ParseForm();
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//设置cookie
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	fmt.Println("/sayhelloName-------")
}
func getcookie(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("getcookie-------")
	cookie, _ := r.Cookie("username")
	fmt.Println(cookie)
	fmt.Fprint(w, cookie)
	//for _, cookie := range r.Cookies() {
	//	fmt.Fprint(w, cookie.Name)
	//}
	fmt.Println("/getcookie-------")
}
func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/getcookie", getcookie) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}