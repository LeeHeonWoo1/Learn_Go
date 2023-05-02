### Array in Go
Go언어의 배열을 만들어보자. 아래와 같이 작성한다.
```go
func main(){
  names := [5]string{"lhw", "lsh", "kms"}
  names[3] = "lsl"
  names[4] = "lji"
  names[5] = "khi"
  fmt.Println(names)
}
```
names라는 Array을 선언할 때 길이와 타입을 설정해주고, 배열의 요소를 그 옆에 작성한다. 0, 1, 2의 index에 해당하는 값을 미리 작성했고 이후에 3, 4, 5의 index에 해당하는 요소들을 작성했는데, 이 때 `names[5]`는 오류를 발생시킬 것이다. 왜냐하면 Array 최초 선언 시 작성한 길이는 5이며 index는 0, 1, 2, 3, 4까지밖에 없기 때문이다.  
그렇다면 길이 제한에 상관없이 값을 계속 추가하려면 어떻게 해야할까? 이 때 사용하는것이 바로 Slice이다.

### Slice
Go의 Slice는 Python의 Slicing과는 개념이 약간 다르다. Go에서의 Slice는 아래와 같이 사용한다.
```go
func main(){
  names := []string{"lhw", "lsh", "kms"}
  names = append(names, "lsl")
  names = append(names, "lji")
  names = append(names, "khi")
  
  for index, value := range names{
    fmt.Println(index, value)
  }
}
```
Slice는 길이 제한이 없는 타입의 배열을 의미한다. 최초 선언은 배열 선언과 비슷하지만, 배열 안에 길이를 의미하는 숫자 없이 선언한다. 이후 값을 추가하고 싶으면 `append`라는 함수를 사용한다. `append`함수는 두 가지 인자를 받는데, 첫 번째로는 값을 추가하고자 하는 Slice객체를 받고, 두 번째로는 추가하고자 하는 값을 받는다. 따라서 위의 코드는 오류 없이 모든 값들이 추가될 것이다.  

또한 Array는 itterable한 객체이기에 for문안에서 itterable한 객체들의 값에 접근하게 해주는 `range`를 이용해서 각 index나 value를 출력할 수 있다. 출력을 원치 않는 값의 경우에는 `_`으로 값을 무시하면 된다.

파이썬에서의 Slicing이라고 하면 배열 안의 특정 값을 가져오는 행위를 일컫는 말이었다. 하지만 Go에서의 Slice는 약간 다름에 주의하면서 사용하자.

### Map
이번엔 Map이라는 것을 알아보려 한다. 우선 아래의 코드를 보자.
```go
func main(){
  leeheonwoo := map[string]string{"name":"leeheonwoo", "age":"26"}
  for key, value := range leeheonwoo{
    fmt.Println(value)
  }
}
```
위 코드처럼 `key:value`쌍으로 이루어진 객체를 Map이라고 한다. 선언 시 `map[key's type]value's type`식으로 선언한다. 모두 `string`타입으로 선언되었기에 나이에 해당하는 `age`는 숫자임에도 불구하고 `""`로 감싸 문자열로 표현했다.