package main

import (
	"fmt"
	"net"
	"sync"
	"flag"
	"time"
)


func worker(host string, timeout time.Duration, jobs <- chan int, results chan <- int, wg *sync.WaitGroup){
	defer wg.Done()

	for port := range jobs {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err == nil {
			conn.Close()
			results <- port
		}
	}
}

func main() {
	var wg sync.WaitGroup

	host := flag.String("host", "localhost", "Host to scan")
	start := flag.Int("start", 1, "Start of port range")
	end := flag.Int("end", 1024, "End of port range")
	workers := flag.Int("workers", 100, "Number of workers")
	timeout := flag.Duration("timeout", 500 * time.Millisecond, "Timeout duration (e.g. 50ms)")
	flag.Parse()

	fmt.Printf("Scanning %s (%d-%d) with %d workers...\n\n", *host, *start, *end, *workers)

	count := 0

	results := make(chan int)
	jobs := make(chan int)

	startTime := time.Now()

	for i:=0; i < *workers; i++{
		wg.Add(1)
		go worker(*host, *timeout, jobs, results, &wg)
	}
	
	go func(){
		for p:= *start; p<=*end; p++ {
			jobs <- p
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for port := range results {
		fmt.Printf("%d OPEN\n", port)
		count++
	}
	
	elapsed := time.Since(startTime)
	fmt.Printf("\nFound %d open ports!\n", count)
	fmt.Printf("Scan completed in %.1fs\n", elapsed.Seconds())
}
