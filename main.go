// package main

// import (
// 	"fmt"
// 	"time"
// 	"sync"
// 	"pro7_finder/finder"
// )

// func main() {
// 	var dir, fileName string
// 	var wg sync.WaitGroup

// 	fmt.Print("Enter directory: ")
// 	fmt.Scan(&dir)

// 	fmt.Print("Enter file name: ")
// 	fmt.Scan(&fileName)

// 	// basicFinder := &finder.Finder{}
// 	// start := time.Now()
// 	// basicFinder.BasicFinder(dir, fileName)
// 	// fmt.Println("Basic time:", time.Since(start))

// 	concurrentFinder := &finder.Finder{}
// 	start := time.Now()

// 	wg.Add(1)
// 	go concurrentFinder.SemFinder(dir, fileName, &wg)
// 	wg.Wait()

// 	fmt.Println("Concurrent time:", time.Since(start))
// 	for _, ele := range concurrentFinder.Res {
// 		fmt.Printf("-> %v \n", ele)
// 	}

// 	jobFinder := &finder.Finder{}
// 	start = time.Now()

// 	// wg.Add(1)
// 	jobFinder.JobFinder(dir, fileName)
// 	// wg.Wait()
// 	fmt.Println("Jobs time : ", time.Since(start))
// 	for _, ele := range jobFinder.Res {
// 		fmt.Printf("-> %v \n", ele)
// 	}
// }

package main

import "pro7_finder/cmd"

func main() {
    cmd.Execute()
}

