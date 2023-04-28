package main

import (
	"fmt"
	"strings"
)

func mult(a int, b int) int {
  return a * b
}

func lenAndUpper(name string)(length int, uppercase string){
  defer fmt.Println("lenAndUpper함수가 실행되었습니다.")
  length = len(name)
  uppercase = strings.ToUpper(name)
  // naked return
  return 
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


