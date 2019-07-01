//this is a re-write of our producer consumer sample in our cppidioms collection which was written in c++
//

package main
import (
	"fmt"
	"time"
//	"os"
//	"bufio"
)

func producer(c chan int, n int) {
	var i int;
	fmt.Printf("producer got n %d\n", n);
	for i=0; i<n; i++ {
			c <- i;
			fmt.Printf("pushing %d\n", i);
			fmt.Printf("AA queue length %d\n", len(c));
		}
}

func consumer(c chan int) {
	var x int;
	for true {
		x = <-c;
		fmt.Printf("consuming %d\n", x);
		if len(c) == 0 {
		  break;
		}
		fmt.Printf("BB queue length %d\n", len(c));
		};
}

func main() {
  c := make (chan int, 300);
  go producer(c, 10);
  fmt.Printf("A queue length %d\n", len(c));
  go consumer(c );
  fmt.Printf("B queue length %d\n", len(c));
  time.Sleep( 2 * time.Second )
  fmt.Printf("C queue length %d\n", len(c));
  for len(c) > 0 {
	time.Sleep( 2 * time.Second )
  }
  fmt.Printf("finished\n");
}
