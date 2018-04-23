package main
import(
    "fmt"
    "flag"
    "strings"
)

var n = flag.Bool("n", false, "omit trailing newline")//创建一个新的布尔形标志参数的变量
var tep = flag.String("s", "", "separator")//创建一个对应字符串类型标志参数的变量

func main(){
    flag.Parse()//更性标志参数对应的值
    fmt.Print(strings.Join(flag.Args(), *tep))
    if !*n{
        fmt.Println()
    }
}
