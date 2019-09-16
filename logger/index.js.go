package logger

import "fmt"

func Println(value interface{}) {
	fmt.Println("===============> Start ")
	fmt.Println(value)
	fmt.Println("===============> End ")
}
