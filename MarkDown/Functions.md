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
이제 오류 없이 4가 출력되는 것을 확인할 수 있다. 또한, a와 b의 타입을 명시하는 과정에서 a와 b 둘 다 정수형이라면 `a, b int`라고 적어도 무방하다.

### Multi Return
Go는 다른 언어들과는 다르게 한 함수가 여러개의 값을 리턴할 수도 있다. 우선 아래처럼 함수를 작성하자.
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
  repeatMe('a', 'b', 'c', 'd', 'e')
}
```
우선 위 코드를 실행하면 배열을 출력하는 모습을 볼 수 있는데, 여러개의 인자를 함수에 받으려면 타입 앞에 점 세개 `...` 를 찍어주면 된다.