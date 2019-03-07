package main
import "fmt"

func main(){
    var n,temp,sum int64
    fmt.Scan(&n)
    for i:=int64(0);i<n;i++{
        fmt.Scan(&temp)
        sum += temp
    }
    fmt.Println(sum)
}
