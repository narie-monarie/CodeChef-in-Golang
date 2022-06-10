package main
import "fmt"

func main(){
    var a,b,c,y int
    fmt.Scanln(&y)
    
    for i := 0; i < y; i++{
        fmt.Scanln(&a, &b, &c)
        
        if c >= a && b > c{
            fmt.Println("YES")
        }else{
            fmt.Println("NO")
        }
    }
}
