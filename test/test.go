package main

import "fmt"

func main()  {
	num:=10001
	pos := num & 0x07
	fmt.Println(pos,fmt.Sprintf("%b,%b,%b",num,0x07,pos))
}
