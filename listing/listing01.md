Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
Программа выведет слайс чисел от [77,78,79] т.к. в новый слайс b войдет срез a[1:4] под индексами 1-3 из 
Слайс b имеет длину 3, а cap 4.
```
