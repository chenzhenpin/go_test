package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

type Contact struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Contact struct {
	Home string `json:"home"`
	Cell string `json:"cell"`
	} `json:"contact"`
}

func main()  {
	var c Contact
	// 将 JSON 字符串反序列化到变量
	c_err := json.Unmarshal([]byte(JSON),&c)
	if c_err !=nil{
		log.Println("ERROR:", c_err)
		return
	}
	fmt.Println(c)

	var s map[string]interface{}
	// 将 JSON 字符串反序列化到 map 变量
	s_err := json.Unmarshal([]byte(JSON),&s)
	if s_err !=nil{
		log.Println("ERROR:", s_err)
		return
	}
	fmt.Println(s)

	data, err := json.MarshalIndent(s, "", " ")
	if err !=nil {
		log.Println("ERROR:", s_err)
		return
	}
	fmt.Println(string(data))
	
}