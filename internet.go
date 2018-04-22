package main

import("fmt"
		"net/http"
		"io/ioutil")

func main() {
	
	/*
	* resp = response
	* _ error, if not used, go discard _ variable
	*/

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
}