package main
    import "fmt"

    func main(){
        var n,t,max,occ int
        fmt.Scan(&n)

        a := []int{}
        for i:=0;i<n;i++{
            fmt.Scan(&t)
            a = append(a,t)
            if t > max {
                max = t
            }
        }

        for _,num := range a{
            if num == max{
                occ++
            }
        }

        fmt.Println(occ)

    }
