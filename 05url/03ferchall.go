package main
import(
    "os"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "time"
)

func main(){
    start := time.Now()
    ch := make(chan string)//chan是一个FIFO队列，chan分成两种类型同步和异步
    for _,url := range os.Args[1:]{
        go ferch(url,ch)

    }
    //这个循环处理并打印channel里这个字符串
    for range os.Args[1:]{
        fmt.Println(<-ch)
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

//这个函数都会向ch这个channel里写一个字符串（感觉像是管道）
func ferch(url string, ch chan <- string){
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil{
        ch <- fmt.Sprint(err)
        return 
    }
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil{
        ch <- fmt.Sprintf("while reading %s: %v",url, err)
        return 
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s",secs,nbytes,url)

}
