package main

import (
	"net/http"
	"fmt"
	"time"
	"log"
	"lib"
  )
  func aboutHandler(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	fmt.Fprintf(w, "You are on the about page.")
	t2 := time.Now()
	log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
  }
  
  func indexHandler(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	fmt.Fprintf(w, "Welcome!")
	t2 := time.Now()
	log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
  }

  

  func testHandler(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	a := [5]int{1, 2, 3, 4, 5}
	arr := a[:]
	i := lib.BinarySearch2(arr,0, len(arr) - 1, 1)
	if(i != -1) {
		fmt.Fprintf(w, "找到了index:%d -- value:%d",i,arr[i])
	}else {
		fmt.Fprintf(w, "没有找到")
	}
	
	t2 := time.Now()
	log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
  }

  func run() {
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
  }
  
  func main() {
	run()
  }