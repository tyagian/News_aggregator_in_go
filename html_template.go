package main
import ("fmt"
"net/http"
"html/template" )

type NewsAggPage struct {
	Title string
	News string
}

// writer type & response type
func indexHandler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"<h1>Welcome to homepage</h1>")
}

func newsAggHandler(w http.ResponseWriter,r *http.Request) {
	p := NewsAggPage{Title: "News Aggregator", News: "some news"}
	t,_ := template.ParseFiles("basictemplating.html")
	//fmt.Println(err)
	// to check error fmt.Println(t.Execute(w,p))
	t.Execute(w,p)
}

func main() {
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/agg/",newsAggHandler)
	http.ListenAndServe(":8000",nil)
}