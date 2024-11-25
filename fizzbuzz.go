/* The FizzBuzz problem  requires writing a function that prints numbers from 1 to a given limit, but with a twist:
    For multiples of 3, print "Fizz" instead of the number.
    For multiples of 5, print "Buzz" instead of the number.
    For numbers which are multiples of both 3 and 5, print "FizzBuzz".
*/
package main
import "fmt"

func main() {
    fizzbuzz(35)
}

func fizzbuzz(max int) {
    for num := 0; num <= max; num++ {

        if num%3 == 0 {
            fmt.Println("fizz")
        } else if num%5 == 0 {
            fmt.Println("buzz")
        } else if num%3 == 0 && num%5 == 0 {
            fmt.Println("fizzbuzz")
        } else {
            fmt.Println(num)
        }
	}
}