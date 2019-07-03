package main
import "fmt"
func main() {

    if 5%2 == 0 {
        fmt.Println("5 is even")
    } else {
        fmt.Println("5 is odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    //A statement can precede conditionals; any variables declared in this statement are available in all branches.
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    var num int
    fmt.Printf("Enter a number: ")
    fmt.Scan(&num)

    if num%2 == 0 {
        fmt.Println(num, "is even number")
    } else {
        fmt.Println(num, "is odd number")
    }
}
