# Go Concurrency Patterns

This repository contains examples and implementations of various concurrency patterns in Go. These patterns are designed to help you better understand how to manage concurrent processes efficiently using Go's powerful concurrency primitives.

## Overview

In Go, concurrency is handled through **goroutines** and **channels**, which enable safe communication between concurrent processes. This repository explores several concurrency patterns such as basic goroutines, select statements, buffered vs. unbuffered channels, and pipelines.

### Examples Covered

1. **Basic Goroutine**
   - A goroutine is a lightweight thread managed by the Go runtime.
   - Example code to run multiple goroutines concurrently.

   ```go
   func someFunc(num string) {
       fmt.Println(num)
   }

   func main() {
       go someFunc("1")
       go someFunc("2")
       go someFunc("3")
       time.Sleep(2 * time.Second)
       fmt.Println("Hi")
   }
   ```

2. **Select Statement**
   - The `select` statement allows you to wait on multiple channel operations. The first channel to send or receive a value is processed.

```go
func main() {
    myChannel := make(chan string)
    anotherChannel := make(chan string)

    go func() {
        myChannel <- "goat"
    }()
    go func() {
        anotherChannel <- "cow"
    }()

    select {
    case msgFromMyChannel := <-myChannel:
        fmt.Println(msgFromMyChannel)
    case msgFromAnotherChannel := <-anotherChannel:
        fmt.Println(msgFromAnotherChannel)
    }
}

```

3. **Buffered vs Unbuffered Channels**
   - Buffered channels can hold values without a receiver being ready.
   - Example showing how buffered channels work with the `select` statement.
  
 ```go
func main() {
    charChannel := make(chan string, 3)
    chars := []string{"a", "b", "c"}

    for _, s := range chars {
        select {
        case charChannel <- s:
        }
    }
    close(charChannel)

    for result := range charChannel {
        fmt.Println(result)
    }
}

```

4. **Done Channel**
   - A`done` channel is used to signal the termination of a goroutine.

```go
func doWork(done <-chan bool) {
    for {
        select {
        case <-done:
            return
        default:
            fmt.Println("Doing some work...")
        }
    }
}

func main() {
    done := make(chan bool)

    go doWork(done)

    time.Sleep(time.Second * 3)

    close(done)
}

```

5. **Pipeline**
   - Pipelines are a series of stages where the output of one stage is the input to the next, connected via channels.

```go
func sliceToChannel(num []int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range num {
            out <- n
        }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {
    nums := []int{2, 4, 6, 8, 9, 12, 15}
    dataChannel := sliceToChannel(nums)
    finalChannel := sq(dataChannel)
    for n := range finalChannel {
        fmt.Println(n)
    }
}

```

### Conclusion

This repository demonstrates fundamental Go concurrency patterns such as goroutines, channels, and more advanced patterns like pipelines. Understanding these concepts will enable you to write efficient concurrent programs in Go.

### I had much fun learning these patterns, Happy Hacking 💻
