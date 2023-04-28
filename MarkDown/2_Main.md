### 시작
src 내에 `github.com`이라는 디렉토리를 만들자. 이후 그 안에 내 github이름으로 된 폴더를 하나 만든다.  
`[C:\Users\OWNER\Desktop\GO\src\github.com\LeeHeonWoo1]`  
이제 이 안에는 우리가 작업할 프로젝트 폴더들을 만들것이다.

### main.go
아래 경로에 `main.go`파일을 생성하자. `LearnGo`라는 폴더는 없기에 생성해야 한다.  
`[GO\src\github.com\LeeHeonWoo1\LearnGo\main.go]`  
해당 파일은 하나의 패키지이며, 프로젝트를 컴파일한다. 목적에 따라(모든 사람들이 공유하며 같이 사용하는 경우) 필요 없는 경우도 있다.  

### main.go 작성하기
우선 `main.go`의 최상단에 아래와 같이 작성해보자.
```go
package main
```
이후 터미널창에 `go run main.go`라는 명령어를 실행하면 아래와 같은 오류 메세지를 얻을 것이다.
```
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package
```
main package에 function main이 선언되지 않았다는 뜻인데, Go는 Node.js나 Python과는 달리 특정 함수를 찾게 되는데, 그 함수가 Go의 시작점이라고 할 수 있다. 아래처럼 말이다.
```go
package main

func main(){

}
```
Go는 main package와 그 안에 있는 main function을 먼저 찾고 실행시킨다. 따라서 함수가 없으면 오류 메세지를 얻게 되고, 현재 main이라는 함수 내에 별다른 내용이 없지만, `go run main.go`를 실행하면 오류는 더이상 발생하지 않을 것이다.

### print hello world
아래와 같이 함수를 작성하여 Hello World를 출력해보자.
```go
package main

import "fmt"

func main(){
  fmt.Println("Hello World!")
}
```
이후 `go run main.go`를 터미널 창에서 실행하면 Hello World! 라는 문자열이 찍히는걸 확인할 수 있다. 코드를 하나씩 살펴보면,  
- `import "fmt"` : Go에서 제공하는 formatting관련 pakage
- `fmt.Println("Hello World!")` : fmt의 Println이라는 함수를 통해 문자열을 출력했다.  
> Println 함수의 P가 대문자인 이유
>> 일반적으로 Node.js, JS에서 함수를 export할 때 `export default ~`라는 식으로 export할 것임을 명시해주는 것을 알 것이다. Go의 경우, 함수를 export하고 싶다면 함수의 맨 앞글자를 대문자로 작성해주면 된다. 즉, 첫 글자를 소문자로 작성하는 함수는 private function, 대문자로 작성하는 함수는 export function인 것이다.