package main

import (
	"fmt"
)

func main() {
	classname := "15-440"               // Type inference

	fmt.Println("Hello,", classname)

	fmt.Println("Hello half-class: ", classname[0:3])

	fooname := classname[0:3]
	fmt.Println("Hello fooclass: ", fooname)
	
//	classname[0] = 'C'       // Strings are immutable!
//	fooname[0] = 'C'         // A slice refs the original

//	buf := []byte(classname)
//	buf[0] = byte('C')
//	fmt.Println(string(buf))
}
