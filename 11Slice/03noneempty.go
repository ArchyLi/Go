package main
import "fmt"

func nonempty(sret []string) []string{
    i := 0
    for _, s := range sret{
        if s != ""{
            sret[i] = s
            i++
        }   
    }
    return sret[:i]
}

func main(){
    s := []string{"haha","","hehe"}
    fmt.Printf("%q", s)
    fmt.Printf("%q", nonempty(s))
    fmt.Printf("%q", s)
}
