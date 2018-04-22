//输出参数索引和值，每个一行
package main
import(
    "fmt"
    "os"
    "strconv"//数字转字符串
)
func main(){
    s := ""
    var i = 0
    for _, arg := range os.Args[1:]{
        d := strconv.Itoa(i)
        s = "Args[" + d+ "]" +" " + arg
        fmt.Println(s)
        i++
        s = ""
    }
}
