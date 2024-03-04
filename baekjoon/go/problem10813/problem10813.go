package main

import "fmt"

func main() {
    var N,M int

    _, err := fmt.Scan(&N, &M)
    if err != nil {
        fmt.Print(err)
    }

    a := make([]int, N)

    for i := 0; i < N; i++{
        a[i] = i+1
    }

    var b,c,d,f int

    for i := 0; i < M; i++ {
         fmt.Scan(&b, &c)
         d = a[b-1]
         f = a[c-1]
         a[b-1] = f
         a[c-1] = d
    }
    
    for i,_ := range a {
        fmt.Print(a[i], " ")
    }
}