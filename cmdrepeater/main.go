package main
import (
        "bufio"
	"fmt"
	"os"
)
func main(){
var s, y,y2, sep, input string
sep = " "
y = "Yes sir we will do "
y2 = " Sure we do"
s = " "
for i := 0; i < len(os.Args); i++ {
	s += sep + os.Args[i];
	}
	fmt.Println(y + s)

scanner := bufio.NewScanner(os.Stdin)
for {
//wait for an enter key
	scanner.Scan()
        s = scanner.Text()
        fmt.Println(y + sep + s)
        fmt.Scanln(&input);
        fmt.Println( s + sep + y2);
}
}
