//This is a GOLANG reimplementation of the topk project in the cppidioms repo
//This exercise will warm up the integer array processing syntax of golang.
//This sample is written for the phone screen interview to be held tomorrow from Cray Inc.
//They have an out sourced C++ package written by team outside USA and now they need to
//rewrite them into GOLANG quickly using a contractor.
//Notes on 6/30/2019
// 
//
// topK - print the top K numbers from a file (stream):   
//
//           topk filename K
//
// current version able to pass test of : top3 example3.txt [1,..,5]
// In other words, all integers are supplied one at a line, no blank lines
// other examples are not working yet
//

package main
import (
	"fmt"
	"os"
	"bufio"
	"math"
	"strconv"
)

func main() {
  var arr []int
  var K, N int
  var i, j, l int;

  if len(os.Args) <= 2 {
     fmt.Println("usgae: topk filename K \n")
     return
  }

  K, err := strconv.Atoi(os.Args[2])
  if err != nil {
	// handle error
	fmt.Println(err)
	os.Exit(2)
  }

  arr = make( []int, K);
  
  for i=0; i< K; i++ {
		arr[i] = math.MinInt32;
  }

  file, err := os.Open(os.Args[1])
  if err != nil {
      panic(err)
  }

  defer file.Close()
  N = 0;

  input  := bufio.NewScanner(file)
  for input.Scan() {
      number, err := strconv.Atoi(input.Text())
      if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
	}
	
	i = K -1;
	if number < arr[i] {
	  break;
	}

	l = 0;
	
	for i=0; i <K; i++ {
		if  number > arr[i] {
			l = i;                   //this will fix the uncertainty of value i
			break;
		 }
	}
	
	j = K-1;
	for j=K-1; j>l; j-- {
		arr[j] = arr[j-1];
	}

	arr[l] = number;
	N++;
  }

  fmt.Printf("read total %d integers, the top %d is:\n", N, K);

  for i=0; i <K; i++ {
		fmt.Printf("%d\n", arr[i]);
  }
}