package main
import(
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main(){
    counts := make(map[string]int)
    for _, filedata := range os.Args[1:]{
        f, err := ioutil.ReadFile(filedata)
        if err != nil{
            fmt.Fprintf(os.Stderr,"dup3 %v",err)
            continue 
        }
        for _, line := range strings.Split(string(f), "\n"){
            counts[line]++
        }
    }   
    for line,n := range counts{
        if n > 1{
            fmt.Printf("%d\t%s\n",n,line)
        }
    }
}
