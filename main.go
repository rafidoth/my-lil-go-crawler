package main

import "fmt"

func main() {
	res := getHeadingFromHTML("<html><body><h1>Test Title</h1></body></html>")
	fmt.Println(res)
}
