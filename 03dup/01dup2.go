package main
import(
    "bufio"
    "fmt"
    "os"
)

func main(){
    counts := make(map[string]int)
    flies := os.Args[1:]

    if len(flies) == 0 {
        countLines(os.Stdin, counts)//此函数说明在下面
    }else{
        for _,arg := range flies{
            f, err := os.Open(arg)//此处两个参数，一个是错误
            if err != nil{//和C++中NULL是一样的
                fmt.Fprintf(os.Stderr,"dup2: %v",err)
                continue 
            }
            countLines(f,counts)//下面
            f.Close()
        }
    }
    for line, n := range counts{
        if n > 0{
            fmt.Printf("%d\t%s",n,line)
        }
    }
}

func countLines(f *os.File, counts map[string]int){
    input := bufio.NewScanner(f)//读取下一行并且去掉换行符
    for input.Scan(){
        counts[input.Text()]++//Text是取文本
    }

}
