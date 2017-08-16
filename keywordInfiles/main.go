//usage: keywordfiles keyword file1 [file2] ...
//output: print out the line in file1 which contains the keyword, so do file 2,...

package main
import (
	"fmt"
	"os"
	"bufio"
        "path/filepath"
        "strings"
        //"io"
)

func main() {
  var line string
  var n int

  if (len(os.Args) <= 2) {
     fmt.Println("usgae: keywordfiles keyword  file1 file2 ... \n")
     return
  }
   
  keyword := os.Args[1]
  files := os.Args[2:]

  //check if files exist, if any of them not exist, we report error and exit

  //get current path first
  exefile := os.Args[0]
  for _, file := range files {
     filepath := filepath.Join(filepath.Dir(exefile), file)
     _, err := os.Stat(filepath) 
     if os.IsNotExist(err) {
         fmt.Printf("search %s in file %s but file not exist \n", keyword, filepath)
         return
     }
     fmt.Printf("search %s in %s ...\n", keyword, filepath);     
   
     //read each line of the current file and check if it contains the keyword
     f, err := os.Open(file)
     if err != nil {
         fmt.Printf("error open file %s, skipped\n", file)
         continue
     }
     var line string  
     var lineno int   
     linetext := bufio.NewScanner(f)
     for linetext.Scan() {
       lineno++
       line = linetext.Text() 
   //    println("processing " + line);
       if strings.Contains(line, keyword) {
          fmt.Printf( "%d ***** %s \n", lineno, line)
       }
//       if len(line) <= 1  {
//          continue
//       }
    } 
  }

  
  return

  counts := make(map[string] int)
  input  := bufio.NewScanner(os.Stdin)
  for input.Scan() {
     counts[input.Text()]++
     if (len(input.Text()) <=1) {
      break
     }
  }
  for line, n = range counts {
          if n > 1 {
		fmt.Printf("%d   %s\n", n, line)
          } else {
		fmt.Printf("no duplicated lines: %s\n", line)
         }
  }
}
