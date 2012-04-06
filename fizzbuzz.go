package main
import "fmt"
func main () {
    i:=1
    for i<=100 {
        out := ""
        if i%3==0 {
            out = out+"Fizz"
        }
        if i%5 == 0 {
            out = out+"Buzz"
        }
        if out == "" {
            fmt.Println(i)
        } else {
            fmt.Println(out)
        }
        i++

    }
}
