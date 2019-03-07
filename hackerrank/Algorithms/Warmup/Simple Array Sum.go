package main
import "fmt"

func main(){
    var n, temp,sum int
    fmt.Scan(&n)
    for i:=0;i<n;i++{
        fmt.Scan(&temp)
        sum += temp
    }
    fmt.Println(sum)
}
