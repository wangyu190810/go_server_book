package main

import "fmt"

func main(){
	str := "Golang is Good!"
	strPta := &str
	fmt.Printf("str type is %T ,value is %v, address is %p\n",str,str,&str)
	fmt.Printf("str type is %T ,value is %v \n",strPta,strPta)
}