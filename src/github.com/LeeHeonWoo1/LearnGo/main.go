package main

import (
	"fmt"
	"strings"
)

func mult(a int, b int) int {
  return a * b
}

func lenAndUpper(name string)(int, string){
  return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
  fmt.Println(words)
}

func main(){
  fmt.Println(mult(2, 2))
  totalLen, totalUpper := lenAndUpper("LeeHeonWoo1")
  fmt.Println(totalLen, totalUpper)
  repeatMe("a", "b", "c", "d", "e")
}


