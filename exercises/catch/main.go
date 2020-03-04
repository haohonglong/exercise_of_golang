package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)
func httpRequest(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
func saveToFile(str string) {
	message := []byte(str)
	err := ioutil.WriteFile("./index.html", message, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
func run() {
	var url string
	fmt.Println("put the url")
	fmt.Scanf("%s",&url)
	html := string(httpRequest(url))
	saveToFile(html)
	fmt.Println(html)
}
func main() {
	run()

}
