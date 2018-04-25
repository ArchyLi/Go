package main
import(
    "unicode/utf8"
    "fmt"
)

func main(){
    s := "hello,世界"
    for i := 0; i < len(s);{
        r,size := utf8.DecodeRuneInString(s[i:])
        //此函数返回一个r和长度，r对应字符本身，
        fmt.Printf("%d\t%c\n",i,r)
        i += size
    }
    
    for i,r := range "hello,世界"{
        fmt.Printf("%d\t%q\t%d\t\n",i,r,r)
    }
   
    x := "hello,世界"
    n := 0
    for _,_ = range x {
        n++
    }
    fmt.Println(n)
}
