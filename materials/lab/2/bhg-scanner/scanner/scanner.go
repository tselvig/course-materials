// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// First run "go build" in the terminal, and then run "go test"

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)    
		conn, err := net.DialTimeout("tcp", address, 3 * time.Second) // I settled on 3 seconds for the timeout because 1 seemed pretty small
		if err != nil { 
			results <- (p * -1)
			continue
		}
		conn.Close()
		results <- p
	}
}

// Scans ports and takes a parameter maxPort (int) that tells the scanner how many ports to scan
func PortScanner(maxPort int) (int, int) {  
	var openports []int  // notice the capitalization here. access limited!
	var closedports []int

	ports := make(chan int, 400)  
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= maxPort; i++ {
			ports <- i
		}
	}()

	for i := 0; i < maxPort; i++ {
		port := <-results
		if port < 0 {
			closedports = append(closedports, (port * -1))
		} else {
			openports = append(openports, port) 
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)

	for _, portO := range openports {
		fmt.Printf("%d, open\n", portO)
	}
	for _, portC := range closedports {
		fmt.Printf("%d, closed\n", portC)
	}

	return len(openports), len(closedports)
}
