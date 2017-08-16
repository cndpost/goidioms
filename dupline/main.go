package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
  var line string
  var n int

  if (len(os.Args) <= 1) {
     fmt.Println("usgae: dupline filename \n")
     return
  }
   
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
