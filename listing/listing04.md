Что выведет программа? Объяснить вывод программы.

```go
package main
 
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
 
    for n := range ch {
        println(n)
    }
}


```

Ответ:
```
Программа выведет числа от 0 до 9 и произойдет deadlock. Сначала мы создадим канал типа int и вызовем анонимную функцию которая будет последовательно забрасывать в канал числа, а в цикле for range мы будем считывать значения из канала, но так как канал не будет закрыт, то и из цикла for range не выйдем так как канал будет ждать пока в него будет закинуто значение.
```
