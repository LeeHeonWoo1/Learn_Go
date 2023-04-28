### Go 시작하기
1. 다운로드  
`https://go.dev/dl/`에서 운영체제에 맞게 다운로드. 이 때 설치 경로는 `C:\Go\`로 설정한다.

2. GOROOT PATH 설정하기  
시스템변수에 GOROOT라는 이름으로 환경변수를 생성하고, 그 경로는 `C:\Go\`로 생성한다. Path항목 내에 `C:\Go\bin`이 존재하는지 확인한다.

3. Go 작업폴더 생성  
Go를 위한 작업파일을 만들자. 나는 `C:\Users\OWNER\Desktop\GO`에 설치했는데, 해당 폴더 내에는 반드시 존재해야하는 세 가지 폴더가 있다.
- bin : ~.go 꼴의 소스코드를 컴파일하면 실행 가능한 바이너리 파일이 저장된다.
- pkg : 프로젝트에 필요한 패키지가 컴파일 되어 라이브러리 파일이 저장된다.
- src : 사용자가 작성한 소스코드나 사용하려는 오픈소스를 저장하는 곳이다.

4. GOPATH등록하기  
시스템변수에 작업 폴더의 경로를 GOPATH라는 이름으로 등록한다.

5. 설치 확인  
cmd창에서 `go version`을 입력해 버전을 확인할 수 있고, `go env`를 입력해 설정한 환경변수들이 잘 적용되었는지 확인할 수 있다.

### Go 실행 테스트
우선 작업 폴더 내에 아래와 같은 파일을 하나 만들자.
`[C:\Users\OWNER\Desktop\GO\test.go]`
```go
package main

import "fmt"

func main(){
	fmt.Printf("hello world!")
}
```
이후 cmd창에서 아래 명령어들을 순차적으로 입력해준다.
```
~\GO>go build test.go
~\GO>go run test.go
```
터미널 창에 `hello world!`가 출력되는 모습을 확인할 수 있을 것이다.  
또한 `test.go`파일이 위치하는 곳에 `test.exe`라는 바이너리 파일이 생성됨을 확인할 수 있을것이다.