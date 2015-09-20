package main
import(
	"net/http"
	"github.com/russross/blackfriday"
)
func main(){
	http.HandleFunc("/markdown",GenerateMarkdown)
	http.HandleFunc("/",http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080",nil)
}
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request){
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	r.Write(markdown)
}