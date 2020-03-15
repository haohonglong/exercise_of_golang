package main
import (
	"io/ioutil"
	"net/http"
	_"fmt"
	"log"
)
func main() {
	res, err := http.Get("http://gbus.me/2018/12/02/%e7%a7%80%e4%ba%ba-no-1136-miki%e7%b1%b3%e9%9b%aa%e5%84%bf/")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s", robots)

    ioutil.WriteFile("./a.html",robots,0777)
}