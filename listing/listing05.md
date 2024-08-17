Что выведет программа? Объяснить вывод программы.

```go
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}



```

Ответ:
```
Программа выведет error, потому что мы создали интерфейс err и вернули из функции test *customError, который удовлетворяет интерфейсу error. Значение внутри err будет nil, а другой указатель будет указывать на метаинформацию о типе customError. Если хотя бы один указатель внутри интерфейса не равен nil , то интерфейс не равен nil.
```
