
package main

import "fmt"
import "math"

const s string = "This is a constant"
func main() {

    fmt.Println("***Values***")
    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)

    fmt.Println("***Variables and constatnts***")
    var name = "ashiqu"
    fmt.Println(name)

    // You can declare multiple variables at once.
    var a, b int = 10, 2
    fmt.Printf("a=%d b=%d\n",a, b)
    fmt.Println("a+b = ",a+b)

    // Go will infer the type of initialized variables.
    var d = true
    fmt.Println(d)

    // Variables declared without a corresponding
    // initialization are _zero-valued_. For example, the
    // zero value for an `int` is `0`.
    var e int
    fmt.Println(e)

    // The `:=` syntax is shorthand for declaring and
    // initializing a variable, e.g. for
    // `var f string = "apple"` in this case.
    f := "apple"
    fmt.Println(f)

    fmt.Println(s)

    const n = 500000000

    const m = 3e20 / n
    fmt.Println(m)

    fmt.Println(int64(m))

    fmt.Println(math.Sin(n))
}
