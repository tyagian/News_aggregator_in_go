// net/http usage with handler
package main
import ("fmt"
"net/http")

// writer type & response type
func index_handler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"Welcome to homepage")
}
func about_handler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"About page")
}
func main() {
	http.HandleFunc("/",index_handler)
	http.HandleFunc("/about/",about_handler)
	http.ListenAndServe(":8000",nil)
}