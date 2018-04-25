package main
import (
    "fmt"
)
func main(){
    s := "hello, word"
    fmt.Println(s[0],s[7])
    fmt.Println("good," + s[7:])
}
