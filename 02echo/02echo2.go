package main
import(
    "fmt"
    "os"
)

func main(){
    s, set := "",""
    for _, arg := range os.Args[1:]{
        s += set + arg
        set = " "
    }
    fmt.Println(s)
}
