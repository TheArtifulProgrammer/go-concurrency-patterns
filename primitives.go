package main

import "fmt"

// basic goroutine

/*

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

*/

// selector func
/*

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

*/
//  Buffered vs Unbuffered channels
// 1. for select loop func
/*
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

// 2. for range func

	func main() {
		go func() {
			for {
				select {
				default:
					fmt.Println("Doing some work...")
				}
			}
		}()
		time.Sleep(time.Second * 10)
	}

*/

// the done channel

/*
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
*/
// pipelines
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
	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	finalChannel := sq(dataChannel)
	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}
