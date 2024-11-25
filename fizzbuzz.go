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
    for i := 0; i <= max; i++ {
        if i%3 == 0 && i%5 == 0 { // Check for divisibility by both 3 and 5 first
            fmt.Println("fizzbuzz")
        } else if i%3 == 0 { // Check for divisibility by 3
            fmt.Println("fizz")
        } else if i%5 == 0 { // Check for divisibility by 5
            fmt.Println("buzz")
        } else { // Default case for non-multiples of 3 or 5
            fmt.Println(i)
        }
    }
}
