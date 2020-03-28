package main

import (
    "fmt"
)

func main() {
    var n, x int
    fmt.Scan(&n, &x)

    //文字列の作成
    alphabet := [26]string{"A","B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M","N","O", "P","Q","R","S","T","U","V","W","X","Y","Z"}
    var a = "A"
    var s = ""

    for i:=0; i<n; i++ {
        s=s+a
        a = alphabet[i+1]
    }
    
    for i:= 0; i < n; i++ {
        for j:=0; j<n-1; j++{
            fmt.Println("?", s[j], s[j+1])
            var ans string        
            fmt.Scan(&ans)
            //if ans == "<" {
                //tmp := s[j+1]
                //s[j+1] = s[j]
                //s[j] = tmp
            //}
        }
    }

    fmt.Println(s)
}