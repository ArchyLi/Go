package main
import "fmt"

func equal(a,b []string) bool{
    if len(a) != len(b){
        return false 
    }
    for i := range a{
        if a[i] != b[i]{
            return false 
        }
    }
    return true 
}

func main(){
    s := [...]string{0:"hah"}
    m := [...]string{0:"hah"}
    ret := requal(s,m)
    fmt.Println(ret)
}
