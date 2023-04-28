### For Loop in Go
Go의 반복문은 `for`단 하나이다. for each, for in, for if 등 아무것도 없고 오로지 `for`하나만 존재한다. 아래와 같은 형태이다.
```go
func addAll(numbers ...int) int {
  total := 0
  for _, numbers := range numbers {
    total += number
  }
  return total
}

func main(){
  result := addAll(1, 2, 3, 4, 5, 6)
  fmt.Println(result)
}
```
`range`는 배열에 loop을 적용할 수 있게 해주는 역할이다. `_`로 무시한 요소는 index이며, 배열의 각 요소들을 모두 더해 값을 리턴하고 있다. 