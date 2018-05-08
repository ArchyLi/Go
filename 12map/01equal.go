package main
import "fmt"
func equal(x,y map[string]int)bool{
    if len(x)!= len(y){
        return false 
    }
    for k,xv := range x{
        if yv,ok := y[x]; !ok || yv != xv{
            return false 
        }
    }
    return true 
}
