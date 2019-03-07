package main
    import "fmt"

    func main(){
        var l,p,n,z,t int
        fmt.Scan(&l)

        for i:=0;i<l;i++{
            fmt.Scan(&t)
            if t > 0 {
                p++
            }else if t <0 {
                n++
            }else {
                z++
            }
        }

        pf := float64(p)/float64(l)
        nf := float64(n)/float64(l)
        zf := float64(z)/float64(l)

        fmt.Printf("%.6f\n",pf)
        fmt.Printf("%.6f\n",nf)
        fmt.Printf("%.6f\n",zf)

    }
