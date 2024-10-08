Что выведет программа? Объяснить вывод программы.

```go
package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func asChan(vs ...int) <-chan int {
   c := make(chan int)
 
   go func() {
       for _, v := range vs {
           c <- v
           time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
      }
 
      close(c)
  }()
  return c
}
 
func merge(a, b <-chan int) <-chan int {
   c := make(chan int)
   go func() {
       for {
           select {
               case v := <-a:
                   c <- v
              case v := <-b:
                   c <- v
           }
      }
   }()
 return c
}
 
func main() {
 
   a := asChan(1, 3, 5, 7)
   b := asChan(2, 4 ,6, 8)
   c := merge(a, b )
   for v := range c {
       fmt.Println(v)
   }
}


```

Ответ:
```
Выведутся числа 1, 3, 5, 7 из первого вызова asChan и 2, 4 ,6, 8 из второго asChan, но в каждой функции будет выведено согласно своей последовательности чисел, но в рандомный момент, поэтому могут выводиться хаотично относительно первого и второго вызова функции, а затем будет выводиться бесконечно 0.
После вызова первого  asChan() вернули канал a , а после вызова второго вернули канал b. Теперь мы мерджим эти каналы в один общий c канал. После того как все числа были закинуты в канал мы каждый закрываем, но так как мы в бесконечном цикле, то после закрытия канала мы всегда будем получать значение 0 из канала.
```
