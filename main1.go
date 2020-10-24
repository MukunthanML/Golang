package main

import "fmt"

func main1() {
	byteVal := byte('a')
	runeVal := 'a'   //rune as default type of char is rune
	stringVal := "a" //double quotes is not a character
	fmt.Printf("%c = %d %c = %U \n\n", byteVal, byteVal, runeVal, runeVal)

	fmt.Printf("byteVal : type is %T\n", byteVal)
	fmt.Printf("runeVal : type is %T\n", runeVal)
	fmt.Printf("stringVal : type is %T\n", stringVal)
}
