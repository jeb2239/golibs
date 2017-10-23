package main

import "net/http"
import "fmt"

func main() {
	resp, _ := http.Get("http://example.com/")
	var docu []byte
	resp.Body.Read(docu)
	fmt.Printf(string(docu[:]))

}
