package main
import "fmt"

func main() {
  var (
    score,sum,total float64
    stus,clas,pass int
  )
  score,sum = 0.0,0.0
  stus,clas,pass = 2,2,0
  fmt.Println("please input a student number")
  fmt.Scanf("%d", &stus)
  
  fmt.Println("please input a class number")
  fmt.Scanf("%d", &clas)
  
  for j:=1 ;j <= clas;j++ {
    sum = 0.0
    for i := 1;i <= stus;i++ {
      fmt.Printf("please input the score of sutdent which the number %d student in class %d \n", i,j)
      fmt.Scanf("%f", &score)
      if score >= 60 {
        pass++
      }
      sum += score
    }
    fmt.Printf("the average is %f who students are score of total \n",sum / float64(stus))
    total += sum; 
  }
  fmt.Printf("total average are %f who students of the %d classes \n", total / float64(stus*clas), clas)
  fmt.Printf("total pass are %d who students of all \n", pass)

}

