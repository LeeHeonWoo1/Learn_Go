### if, else
Go의 If, else문 형태를 살펴보자. 아래와 같이 코드를 작성한다.
```go
func isAdult(age int){
  if age < 18 {
    return false
  }
  return true
}
```
if else문의 형태는 js와 비슷하며, else를 적지 않는다. 위의 코드만으로 18세 미만이면 false가, 그가 아닌 경우에는 true가 반환된다.  

### Variable Expression
Go에서는 특정 문법을 위한 변수를 만들 수 있다. 이를 <b>variable expression</b>이라고 하는데, 아래의 `isAdult1` 함수를 보면 if문이 시작됨과 동시에 `koreanAge`라는 변수가 생성되었다. 그 변수는 인수로 입력받는 age에 2를 더한 값임을 알 수 있다.  

어떤 차이일까? 아래의 코드를 보자
```go
func isAdult1(age int){
  if koreanAge := age + 2;koreanAge < 18 {
    return false
  }
  fmt.Println(koeanAge)
  return true
}

func isAdult2(age int){
  koreanAge := age + 2
  if koreanAge < 18 {
    return false
  }
  fmt.Println(koeanAge)
  return true
}
```
위의 두 함수의 차이는 `variable expression사용여부` 단 하나이다. 어느 함수가 오류가 발생할까?  
함수를 살펴보면 if문을 지나 else문에 조건이 걸릴 때 koreanAge를 print하고 true를 return하고 있다. 결과부터 얘기하자면 `isAdult1`함수는 오류를 발생시킨다. if문이 생성됨과 동시에 생성된 `koreanAge`라는 변수는 `else`영역에는 존재하지 않는 값이기 때문이다.

### Switch
js의 switch와 동일하게 사용한다. 아래의 코드를 보자.
```go
func isAdult(age int){
  switch koreanAge := age + 2; {
  case koreanAge < 18:
    return false
  case koreanAge == 18:
    return true
  case koreanAge > 50:
    return false
  }
  return true
}
```
if -> else if 가 반복되며 코드 줄 수가 길어지는 현상을 방지해주며 variable expression도 동일하게 사용할 수 있다. 