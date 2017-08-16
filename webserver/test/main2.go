//usage: this is a daomon launched as a web server of static pages

package main
import (
	"fmt"
        "log"
        "net/http"
        "bufio"
        "os"
        "path"    
)

func main() {
    http.HandleFunc("/", handler)
    http.handleFunc("/test", handle2)
    log.Fatal(http.ListenAndServe("localhost:8000",nil))
}


//serve files in root folder
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
    //get the file name from the Path, open the file and send the content over http
    filename := path.Base(r.URL.Path)
    f, err := os.Open(filename)
    if err != nil {
       return
    }
    lineText := bufio.NewScanner(f)
    for lineText.Scan() {
	fmt.Fprintf(w, "%s\n", lineText.Text())
    }
    f.Close()
}
  

//serve files in test folder
func handler2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
    //get the file name from the Path, open the file and send the content over http
    filename := path.Base(r.URL.Path)
    f, err := os.Open(filename)
    if err != nil {
       return
    }
    lineText := bufio.NewScanner(f)
    for lineText.Scan() {
        fmt.Fprintf(w, "%s\n", lineText.Text())
    }
    f.Close()
}

