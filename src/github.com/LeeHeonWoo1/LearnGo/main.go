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

func isAdult(age int) bool {
  switch koreanAge := age + 2; {
  case koreanAge < 18:
    return false
  case koreanAge == 18:
    return true
  case koreanAge > 50:
    return false
  }
  return true
  // koreanAge := age + 2;
  // if koreanAge < 16 {
  //   return false
  // }
  // fmt.Println(koreanAge)
  // return true
}

func main(){
  fmt.Println(isAdult(16))
  // fmt.Println(mult(2, 2))
  // totalLen, totalUpper := lenAndUpper("LeeHeonWoo1")
  // fmt.Println(totalLen, totalUpper)
  // repeatMe("a", "b", "c", "d", "e")
}


