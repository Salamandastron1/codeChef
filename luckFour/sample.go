// We have populated the solutions for the 10 easiest problems for your support. 
// Click on the SUBMIT button to make a submission to this problem.

package main
import "fmt"

func main(){
	var t int
	var num int
	fmt.Scanln(&t)
	counter:=0
	for ;t>0;t--{
	    fmt.Scanln(&num)
	    counter = 0
	    for num>0{
	        if num%10==4{
	            counter+=1
	        }
	        num= num/10
	    }
	    fmt.Println(counter)
	}
	
}


