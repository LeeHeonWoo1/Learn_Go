### Functions
아래와 같이 곱셈을 해주는 함수를 작성해보자.
```go
func mult(a, b){
  return a*b
}

func main(){
  fmt.Println(mult(2, 2))
}
```
하지만 이는 오류를 발생시킬 것이다. Go는 기본적으로 컴파일러에게 모든 값을 정의해서 전달해야 한다. 즉, 현재 `mult`라는 함수가 어떤 타입의 인자를 받아들이는지, 어떤 형태로 return하는지를 모두 알려줘야 한다. 그렇지 않으면 정의되지 않은 값으로 인식한다. 아래와 같이 수정하자.
```go
func mult(a int, b int) int {
  return a * b
}

func main(){
  fmt.Println(mult(2, 2))
}
```
함수의 기본 양식을 살펴볼 수 있는데, 아래와 같을 것이다.
```go
func 함수명(인자 타입, 인자 타입, ...) 리턴값 타입 {
  return 리턴값
}
```
이제 오류 없이 4가 출력되는 것을 확인할 수 있다. 또한, a와 b의 타입을 명시하는 과정에서 a와 b 둘 다 정수형이라면 `a, b int`라고 적어도 무방하다.

### Multi Return
Go는 한 함수가 여러개의 값을 리턴할 수도 있다. 우선 아래처럼 함수를 작성하자.
```go
package main

import "strings"

func lenAndUpper(name string)(int, string){
  return len(name), strings.ToUpper(name)
}

func main(){
  totalLen, totalUpper := lenAndUpper("LeeHeonWoo1")
  fmt.Println(totalLen, totalUpper)
}
```
실행해보면 문자열의 길이인 11과 대문자로 바꾼 결과 두 개의 값이 return됨을 알 수 있다. 만약 대문자로 변환한 결과는 필요없다고 하면 아래와 같이 작성할 수도 있다.
```go
package main

import "strings"

func lenAndUpper(name string)(int, string){
  return len(name), strings.ToUpper(name)
}

func main(){
  totalLen, _ := lenAndUpper("LeeHeonWoo1")
  fmt.Println(totalLen)
}
```
`_`(underscore)는 컴파일러에게 <b>이 값은 무시하는 값</b>임을 명시하는것과 같다.

### 여러 인자를 출력
Go에서는 함수에 여러개의 인자를 입력하고 싶다면 아래와 같이 쓸수도 있다.
```go
func repeatMe(words ...string){
  fmt.Println(words)
}

func main(){
  repeatMe("a", "b", "c", "d", "e")
}
```
우선 위 코드를 실행하면 배열을 출력하는 모습을 볼 수 있는데, 여러개의 인자를 함수에 받으려면 타입 앞에 점 세개 `...` 를 찍어주면 된다.

### naked return
우선 아래의 코드부터 살펴본다.
```go
func lenAndUpper(name string)(length int, uppercase string){
  length = len(name)
  uppercase = strings.ToUpper(name)
  return
}

func main(){
  len, up := lenAndUpper("LeeHeonWoo1")
  fmt.Println(len, up)
}
```
위에서 만들었던 함수를 조금은 다르게 작성해보았다. `return` 이라는 문구 옆에는 아무런 리턴값이 없음에도, 실행한 결과는 `11 LEEHEONWOO1`를 얻을 수 있을 것이다. 이를 `naked return`이라고 하는데, 리턴값의 타입을 지정하는 영역에서 리턴값의 이름까지 지정해주면 굳이 아래에서 `return length, uppercase`라는 문구를 작성할 필요가 없는 것이다. 물론 작성해두어도 작동은 한다.

### defer
함수 안에 작성하는 `defer`는 함수가 어떤 값을 `return`하고 나면 실행될 코드를 작성하는 곳이다. 아래의 코드를 보자.
```go
func lenAndUpper(name string) (length int, uppercase string){
  defer fmt.Println("I'm done")
  length = len(name)
  uppercase = strings.ToUpper(name)
  return
}

func main(){
  length, upper := lenAndUpper("LeeHeonWoo1")
  fmt.Println(length, upper)
}
```
`main`함수의 `length, upper := lenAndUpper("LeeHeonWoo1")`부분에서 `I'm done`이라는 문자열이 출력될 것이다.