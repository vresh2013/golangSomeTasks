package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

const (
	url = "https://docs.github.com/en/apps/creating-github-apps/creating-github-apps/creating-a-github-app-using-url-parameters"
)

func ping(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// 2.5 sec
func plain() {
	start := time.Now()

	for i := 0; i < 100; i++ {
		// fmt.Printf("%d ping \n", i)
		err := ping(url)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	defer fmt.Printf("took %v sec\n", time.Since(start).Seconds())
}

// 1.5 sec
// you must build main.go in root
func multiProcess() {
	start := time.Now()

	done := make(chan int)
	go func() {
		dateCmd := exec.Command("../main")
		dateOut, err := dateCmd.Output()
		if err != nil {
			fmt.Println("error =", err)
			return
		}
		fmt.Println(string(dateOut))
		done <- 1
	}()

	for i := 0; i < 50; i++ {
		// fmt.Printf("%d ping \n", i)
		err := ping(url)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	<-done

	defer fmt.Printf("took %v sec\n", time.Since(start).Seconds())
}

// 0.2 sec
func multiThread() {
	runtime.GOMAXPROCS(2)
	start := time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		// fmt.Printf("%d ping \n", i)
		go func() {
			defer wg.Done()
			err := ping(url)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
	}
	wg.Wait()

	defer fmt.Printf("took %v sec\n", time.Since(start).Seconds())
}

// 0.2 sec
func async() {
	runtime.GOMAXPROCS(1)
	start := time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		// fmt.Printf("%d ping \n", i)
		go func() {
			defer wg.Done()
			err := ping(url)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
	}
	wg.Wait()

	defer fmt.Printf("took %v sec\n", time.Since(start).Seconds())
}

func main() {
	plain()
	// multiProcess()
	// multiThread()
	// async()
}
