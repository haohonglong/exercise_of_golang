package main
import (
	f "fmt"
	"encoding/json"
	"os"
)

type City struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func test() {
	data := map[string]string{
		"001": "北京",
		"002": "天津",
		"003": "上海",
		"004": "重庆",
	}
	cities := []City{}
	city := City{}

	for k, v := range data {
		city.Id = k
		city.Name = v
		cities = append(cities,city)
	}

	b, err := json.Marshal(cities)
	if err != nil {
		f.Println("error:", err)
	}
	os.Stdout.Write(b)

	for _, v := range cities {
	  f.Printf("%s=>%s \n", v.Id, v.Name)
	}
	// println()
	a := []rune{'a','b','c'}
	f.Printf("====%v\n",string(a));
	
}