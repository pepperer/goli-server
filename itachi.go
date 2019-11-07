package main

import (
	"flag"
	"fmt"
)

var (
	url  string
	name string
	pwd  string
)

func init() {
	flag.StringVar(&url, "url", "", "a url")
	flag.StringVar(&name, "name", "", "a url")
	flag.StringVar(&pwd, "pwd", "", "a url")
}

func main() {
	flag.Parse()
	fmt.Println("url:", url)
	fmt.Println("name:", name)
	fmt.Println("pwd:", pwd)


}
