package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := id // Perform some work
	ch <- result // Send the result to the channel
}

func main() {
	numWorkers := 10
	ch := make(chan int, numWorkers)
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	// for result := range ch {
	// 	fmt.Println(result)
	// }

}

// // completed := make(chan int)

// for i := 0; i < 5; i++ {
// 	wg.Add(1)
// 	go func(runningRoutine int) {
// 		defer wg.Done()
// 		fmt.Printf("Go routine %d started\n", runningRoutine)

// 	}(i)
// }
// wg.Wait()
// fmt.Println("Completed")

// go func() {

// 	// for range 3 {
// 	// 	completed <- "ping"
// 	// 	completed <- "pong"
// 	// }
// 	// close(completed)

// 	for i := 0; i < 10; i++ {
// 		completed <- i + 1

// 	}
// 	close(completed)
// }()

// for i := range completed {

// 	fmt.Println(i)

// }

// sigChan := make(chan bool)

// go dunc
// go func() {
// 	var c = 3
// 	for c <3 {
// 		sig := chan
// 		fmt.Printf("ping")

// 	}
// }()

// go func() {
// 	var c = 3
// 	for c <3 {
// 		sig := chan
// 		fmt.Printf("ping")
// 	}
// }()
// go func() {
// 	for range 3 {
// 		sigChan <- <-completed
// 		sigChan <- "pong"
// 	}
// }()

// }

// sigChan := make(chan int)
// go func() {
// 	for i := 0; i < 10; i++ {

// 		sigChan <- i

// 	}
// 	close(sigChan)
// }()

// for i := range sigChan {
// 	fmt.Println(i)
// }
// func ping() {
// 	fmt.Println("Ping")
// }

// func pong() {
// 	fmt.Println("Pong")
// }
// func expensive() {
//     querry
// }
// func main() {
//     sigchan := make(chan struct{})
//     go func() {
//         expensive()
//         sigChan <- struct{}{}
//     }()
//     // Progress tracker.
//     go func(){
//         ticker := time.NewTicker(time.Second)
// 		bol = false
//         for !bol{
//             select {
//             case <-sigchan:
//                 fmt.Println("completed")
//                 return
//             case <-ticker.C
//                 fmt.Println("not completed")
//             }
//         }
//     }()
// }
