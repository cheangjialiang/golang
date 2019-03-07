package main
import "fmt"

func main(){
    var h,m,s int
    var suffix string
    
    fmt.Scanf("%d:%d:%d%s",&h,&m,&s,&suffix)
    
    if suffix == "AM" && h == 12 {
        h = 0
    }
    if suffix == "PM" && h != 12 {
        h += 12
    }
    fmt.Printf("%02d:%02d:%02d",h,m,s)
    
}
