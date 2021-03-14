package main
import ("fmt"
	"net/http"
"io/ioutil"
"encoding/xml")

// index structure at the time of writing this code 
/*
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/politics.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/opinions.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/local.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/sports.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/national.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/world.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/business.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/technology.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/lifestyle.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/entertainment.xml
</loc>
</sitemap>
<sitemap>
<loc>
https://www.washingtonpost.com/news-sitemaps/goingoutguide.xml
</loc>
</sitemap>
</sitemapindex>

*/

type SitemapIndex struct {
	// slice of type Location
	Locations []string `xml:"sitemap>loc"`
}
type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>Keywords`
	Locations []string `xml:"url>loc"`
}
// we can remove this and define above inside SitemapIndex struct with type Location itself.
//type Location struct {
//	Loc string `xml:"loc"`
//}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	var s SitemapIndex
	var n News
	resp,_ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes,_ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	//string_body := string(bytes)
	//fmt.Println(string_body)
	//resp.Body.Close()	
	//fmt.Printf("Here %s some %s","are","variables")
	for _, Location := range s.Locations {
		resp,_ := http.Get(Location)
		bytes,_ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
	}
}
