package main
import "fmt"

func main() {
  var (
    row,op int
  )
  row,op = 1,1

  fmt.Println("please input a type")
  fmt.Scanf("%d",&op)
  
  fmt.Println("please input a row")
  fmt.Scanf("%d", &row)
  
  switch op {
    case 1:
      run1(row)
    case 2:
      up(row)  
    case 3:
      up(row)
      down(row-1)
    case 4:
      up_empty(row)
      down_empty(row-1)


  }
  

}

func run1(row int) {
  for i:=0 ;i <= row;i++ {
    for j := 0;j <= i;j++ {
      fmt.Print("*")
    }
    fmt.Println()
  }
}

func up(row int) {
  for i:=1 ;i <= row;i++ {
    for k:=1;k <= row-i;k++ {
      fmt.Print(" ")
    }
    for j := 1; j <= 2*i-1; j++ {
      fmt.Print("*")
    }
    fmt.Println()
  }
}

func down(row int) {
  for i:=row ;i >=1;i-- {
    for k:=0;k <= row-i;k++ {
      fmt.Print(" ")
    }
    for j := 2*i-1;j >=1;j-- {
      fmt.Print("*")
    }
    fmt.Println()
  }
}

func up_empty(row int) {
  for i:=1 ;i <= row;i++ {
    for k:=1;k <= row-i;k++ {
      fmt.Print(" ")
    }
    for j := 1; j <= 2*i-1; j++ {
      if 1 == j || j == 2*i-1 {
        fmt.Print("*")
      }else{
        fmt.Print(" ")
      }
      
    }
    fmt.Println()
  }
}

func down_empty(row int) {
  for i:=row ;i >=1;i-- {
    for k:=0;k <= row-i;k++ {
      fmt.Print(" ")
    }
    for j := 2*i-1;j >=1;j-- {
      if 2*i-1 == j || 1 == j {
        fmt.Print("*")
      }else{
        fmt.Print(" ")
      }
      
    }
    fmt.Println()
  }
}
