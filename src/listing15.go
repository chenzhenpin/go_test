package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
 // gResult 映射到从搜索拿到的结果文档
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL string `json:"unescapedUrl"`
		URL string `json:"url"`
		VisibleURL string `json:"visibleUrl"`
		CacheURL string `json:"cacheUrl"`
		Title string `json:"title"`
		TitleNoFormatting string `json:"titleNoFormatting"`
		Content string `json:"content"`
	}

	// gResponse 包含顶级的文档
 	gResponse struct {
		 ResponseData struct {
			 Results []gResult `json:"results"`
		 } `json:"responseData"`
	 }
 )

func main()  {
	uri := "http://localhost:8088/user/2"
	resp, err := http.Get(uri)
	if err !=nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(gr.ResponseData.Results[1].Title)
}
