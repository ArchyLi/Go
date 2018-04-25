package main
import (
    "fmt"
)
func main(){
    var x complex64 = complex(1,2)
    var y complex64 = complex(0,1)
    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(real(x))
    fmt.Println(real(y))
    fmt.Println(imag(x))
    fmt.Println(imag(y))
}
