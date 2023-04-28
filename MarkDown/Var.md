### Variable, Constant

#### Constant(상수)
Go를 이용해서 상수를 생성해보자.
```go
func main(){
  const name="LeeHeonWoo1"
}
```
`name`이라는 변수에 커서를 올리면 `untyped`임을 알 수 있는데, Go는 개발자가 작성하는 값의 타입을 찾는 과정을 거친다. Type Language이기에 어떤 타입인지 타입을 명시해야 한다. Java에서 `double b = 12;`라고 작성하는 것처럼 말이다. 따라서 아래처럼 수정해보자.
```go
func main(){
  const name string ="LeeHeonWoo1"
  fmt.Println(name)
}
```
당연히 const로 선언된 Constant값이기에 수정은 불가하며, `go run main.go`를 통해 실행하면 name에 할당된 문자열이 출력될 것이다.

#### Variable(변수)
Go를 이용해서 변수를 생성해보자.
```go
func main(){
  var name string ="LeeHeonWoo1"
  name = "JJW"
  fmt.Println(name)
}
```
변수이기에 값을 수정할 수 있다.

#### 타입 자동 인식
위에서 잠깐 언급했던 것 처럼, Go는 Type Language이기에 타입을 지정해줘야 하지만, Go에서 자체적으로 타입을 인식하게끔 작성할 수도 있다.
```go
func main(){
  name := "LeeHeonWoo1"
}
```
이를 <b>축약형</b>이라고 하는데, 이런 축약형은 <b>함수 안에서만 사용 가능</b>하고, 함수 밖에서는 기존에 선언하던 방식대로 선언해야 한다.