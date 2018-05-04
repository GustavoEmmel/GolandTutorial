package main

/*
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-entertainment-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-entertainment-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-goingoutguide-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-goingoutguide-sitemap.xml</loc>
   </sitemap>
</sitemapindex>
*/

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {

	/*
	* resp = response
	* _ error, if not used, go discard _ variable
	 */

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}

	fmt.Printf("\nHere %s some %s", "are", "variables")

	//string_body := string(bytes)
	//fmt.Println(string_body)

}
