package main
import "fmt"

func main() {
  var y int
  
  fmt.Println("please input a number of year")
    fmt.Scanln(&y)
    if is_leay(y) {
      fmt.Printf("%d is leap year \n",y)
    }else{
      fmt.Printf("%d is not leap year \n",y)
  }
  
}

func is_leay(y int) (bool) {
  if (0 == y%4 && y%100 != 0) || 0 == y%400 {
    return true
  }else{
    return false
  }
}

