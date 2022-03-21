package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {

	hostname := flag.String("hostname", "47.99.118.218", "hostname to test")
	startPort := flag.Int("start-port", 1100, "the port on which the scanning starts")
	endPort := flag.Int("end-port", 10000, "the port from which the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond * 200, "timeout")
	flag.Parse()

	ports := []int{}

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for port := *startPort; port < *endPort; port++ {
		wg.Add(1)
		port := port
		go func() {
			open := isOpen(*hostname, port, *timeout)
			if open {
				mutex.Lock()
				ports = append(ports, port)
				println(port)
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("opened ports: %v\n", ports)

}

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port),timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}
