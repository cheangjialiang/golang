package main

import "fmt"

func checkError(err error){
    if err != nil {
        fmt.Println(err)
    }
}

func main() {
    
    var x1,v1,x2,v2 int
    _, err := fmt.Scanf("%d", &x1)
    checkError(err)
    _, err = fmt.Scanf("%d", &v1)
    checkError(err)
    _, err = fmt.Scanf("%d", &x2)
    checkError(err)
    _, err = fmt.Scanf("%d", &v2)
    checkError(err)
    
    if (v1 > v2) && ((x2-x1)%(v1-v2)) == 0 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
