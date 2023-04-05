package main

import (
	"fmt"
	"net"
	"os/exec"
	"sync"
	"time"
)

// 21 sec
func plain(ips []string) {
	start := time.Now()

	for _, ip := range ips {
		conn, err := net.DialTimeout("tcp", ip, time.Millisecond*300)
		if err != nil {
			fmt.Println("error =", err)
			continue
		}
		fmt.Println("connected to", ip)
		conn.Close()
	}

	fmt.Printf("took %v sec\n", time.Since(start).Seconds())
}

// 0.3 sec
func async(ips []string) {
	start := time.Now()

	wg := &sync.WaitGroup{}

	wg.Add(len(ips))
	for _, ip := range ips {
		go func(ip string) {
			defer wg.Done()
			conn, err := net.DialTimeout("tcp", ip, time.Millisecond*300)
			if err != nil {
				fmt.Println("error =", err)
				return
			}
			fmt.Println("connected to", ip)
			conn.Close()
		}(ip)
	}

	wg.Wait()
	fmt.Printf("took %v sec\n", time.Since(start).Seconds())
}

// func multiThread(ips []string) {
// 	fmt.Println("cpu count =", runtime.NumCPU())

// 	start := time.Now()

// 	wg := &sync.WaitGroup{}

// 	wg.Add(len(ips))
// 	for _, ip := range ips {
// 		go func(ip string) {
// 			defer wg.Done()
// 			conn, err := net.DialTimeout("tcp", ip, time.Millisecond*300)
// 			if err != nil {
// 				fmt.Println("error =", err)
// 				return
// 			}
// fmt.Println("connected to", ip)
// 			conn.Close()
// 		}(ip)
// 	}

// 	wg.Wait()
// 	fmt.Printf("took %v sec\n", time.Since(start).Seconds())
// }

// 0.8 sec
func multiProcess(ips []string) {
	start := time.Now()

	wg := &sync.WaitGroup{}

	wg.Add(len(ips))
	for _, ip := range ips {
		go func(ip string) {
			defer wg.Done()

			cmd := exec.Command("../main", ip)

			cmdOut, err := cmd.Output()
			if err != nil {
				fmt.Println("error cmd output =", err)
				return
			}

			fmt.Println(string(cmdOut))
		}(ip)
	}

	wg.Wait()
	fmt.Printf("took %v sec\n", time.Since(start).Seconds())
}

func main() {
	ips := make([]string, 0, 100)
	for i := 1; i <= 100; i++ {
		ips = append(ips, fmt.Sprintf("94.100.180.%d:80", i))
	}

	// 21 sec
	// plain(ips)

	// 0.3 sec
	// async(ips)

	// 0.8 sec
	multiProcess(ips)
}
