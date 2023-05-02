### Struct
Struct는 `8_Array_Slice_Map.md`파일에서 살펴본 Map객체보다 좀 더 유연한 구조라고 볼 수 있다. 아래의 코드를 보자.
```go
package main

type person struct {
  name string
  age int
  favFood []string
}

func main(){
  favFood := []string{"ramen", "icecream"}
  leeheonwoo := person{name: "leeheonwoo", age: 26, favFood: favFood}
  fmt.Println(leeheonwoo)
}
```
Map객체는 모든 key와 value에 대한 타입을 고정시키고 시작하는 반면, struct의 경우에는 각 key와 value에 대한 속성을 지정해줄 수 있다. person이라는 struct에 정의된 값들은 아래와 같다.  
|변수명|할당될 값|타입|
|---|---|---|
|name|이름|string|
|age|나이|int|
|favFood|좋아하는 음식의 배열|Slice with string|

Map처럼 모든 값에 대한 타입을 일괄적으로 지정하는게 아니라, 각각 지정해줄 수 있다. 또한, 만들어진 struct를 사용할 때는 `{"leeheonwoo", 26, favFood}`와 같이 key값을 생략하고 작성해도 무관하지만, `{name: "leeheonwoo", age: 26, favFood: favFood}`와 같이 key값을 명시해준다면 굳이 struct를 볼 필요 없이 바로 구조를 이해할 수 있을 것이다. 

이 때 주의해야 할 점으로는, `{name: "leeheonwoo", age: 26, favFood: favFood}`와 같이 key:value형태로 작성하기 시작했다면 끝까지 key:value형태로 작성해야하며, `{name: "leeheonwoo", 26, favFood}`와 같이 일부만 key:value형태로 작성하게되면 오류가 발생할 것이다.