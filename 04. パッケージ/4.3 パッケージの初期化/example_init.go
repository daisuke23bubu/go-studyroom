package main

import "fmt"

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
	main()
}

func main() {
	// init()	これはできない
	fmt.Println("main")
}

/*
init1
init2
main
main
*/
