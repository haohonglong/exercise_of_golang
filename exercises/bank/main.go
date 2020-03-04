package main
import "fmt"

type Bank struct {
	Name string
	Password string
	Money float64
}

func (b *Bank) Deposit(name,pass string,money float64) (float64, bool) {
	if b.Name == name && b.Password == pass {
		b.Money += money
		return b.Money,true
	}else {
		return 0,false
	}
	
}

func (b *Bank) Query(name,pass string) (float64, bool) {
	if b.Name == name && b.Password == pass {
		return b.Money,true
	}else {
		return 0,false
	}
	
}

func (b *Bank) TakeOut(name,pass string,money float64) (float64, bool) {
	if b.Name == name && b.Password == pass {
		b.Money -= money
		return b.Money,true
	}else {
		return 0,false
	}
	
}

func Info() int{
	var n int
	
	fmt.Println("======================")
	fmt.Println("please choose the one")
	fmt.Println("	1 deposit")//存入
	fmt.Println("	2 query")
	fmt.Println("	3 take out")//支出
	fmt.Println("	0 exit")
	fmt.Println("======================")
	fmt.Scanf("%d",&n)
	return n
}

func InputData() (string, string, float64, int){
	var (
		name,pass string
		money float64
		n int
	)
	n = Info()
	switch n {
		case 1,3:
			fmt.Println("please input a name")
			fmt.Scanf("%s",&name)
			fmt.Println("please input a password")
			fmt.Scanf("%s",&pass)
			fmt.Println("please input the money")
			fmt.Scanf("%f",&money)
		case 2:
			fmt.Println("please input a name")
			fmt.Scanf("%s",&name)
			fmt.Println("please input a password")
			fmt.Scanf("%s",&pass)
			money = 0
		case 0:
			name,pass,money = "","",0
	}
	return name,pass,money,n
	
}

func main() {
	var (
		m float64
		err bool
	)
	bank := Bank{
		Name : "long",
		Password : "123456",
		Money : 10.00,
	} 
	
	OuterLoop:
	for {
		name,pass,money,n := InputData()
		switch n {
			case 1:
				m,err = bank.Deposit(name,pass,money)
			case 2:
				m,err = bank.Query(name,pass)
			case 3:	
				m,err = bank.TakeOut(name,pass,money)
			case 0:
				break OuterLoop
		}
		fmt.Println("显示的信息是:")
		if true == err {
			fmt.Printf("剩余余额是: %.2f \n",m)
		}else{
			fmt.Println("验证不通过")
		}
		

	}


}