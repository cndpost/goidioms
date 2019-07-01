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
// update: this version now able to handle multiple numbers in a single line, as well empty lines
//          with the change made in line 64-68

package main
import (
	"fmt"
	"os"
	"bufio"
	"math"
	"strconv"
	"strings"
)

func main() {
  var arr []int
  var K, N int
  var i, j, l, number int;
  var text string;
  var err error;
  var file *os.File;

  if len(os.Args) <= 2 {
     fmt.Println("usgae: topk filename K \n")
     return
  }

  K, err = strconv.Atoi(os.Args[2])
  if err != nil {
	// handle error
	fmt.Println(err)
	os.Exit(2)
  }

  arr = make( []int, K);
  
  for i=0; i< K; i++ {
		arr[i] = math.MinInt32;
  }

  file, err = os.Open(os.Args[1])
  if err != nil {
      panic(err)
  }

  defer file.Close()
  N = 0;
  
  input  := bufio.NewScanner(file)
  for input.Scan() {   // loop through all the lines of the file
	  str := input.Text();
	  strs := strings.Split(str, " ");
	  for _, text = range strs {   // loop within a line which may have more than one integers

		if len(text) == 0 {
		  continue; 
		}
		//fmt.Println(text)   // debug

        number, err = strconv.Atoi(text)
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
  }

  fmt.Printf("read total %d integers, the top %d is:\n", N, K);

  for i=0; i <K; i++ {
		fmt.Printf("%d\n", arr[i]);
  }
}