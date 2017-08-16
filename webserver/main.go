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
    log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

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
  
