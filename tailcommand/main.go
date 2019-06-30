//This is a GOLANG reimplementation of my tailcommand project originally written in C++ in my repo called cppidioms
//This implementation has taken code from the double line example. so we have a map calclation which is not needed, but
//we kept here for easy comparison with the double line example

package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
  var line string
  var n int

/*
  if (len(os.Args) <= 1) {
     fmt.Println("usgae: tail filename \n")
     return
  }
*/

  counts := make(map[string] int)
  input  := bufio.NewScanner(os.Stdin)
  for input.Scan() {
     counts[input.Text()]++
     if (len(input.Text()) <=1) {
      continue;
     }
  }
  var i, N int;
  N = len(counts);
  i = 0;

  fmt.Printf("total %d lines, print the last 10 lines\n", N);

  for line, n = range counts {
		if (n > 0) {
	    if ( i > (N - 10) ) {
		fmt.Printf("%s\n", line);
		}
		}
		i++;
  }
}