package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
Author: F1shh-sec

*/

func main() {
	var ipa string
	var startport, endport int

	// Gets input Ip address
	fmt.Println("Enter a host IP: ")
	fmt.Scanf("%s\n", &ipa)
	ipa = strings.Replace(ipa, "\n", "", -1)

	// Grabs starting port
	fmt.Println("Enter a starting port: ")
	fmt.Scanf("%d\n", &startport)

	// grabs ending port
	fmt.Println("Enter a ending port: ")
	fmt.Scanf("%d\n", &endport)

	var wg sync.WaitGroup

	if endport > startport {
		for i := startport; i <= endport; i++ {
			wg.Add(1)
			go connect(ipa, i, &wg)
		}
	}
	wg.Wait()
}

/**
Attempts to connect to a given port.
*/
func connect(ipaddr string, port int, wg *sync.WaitGroup) {
	// Converts the int to a string
	portstr := strconv.Itoa(port)
	// Connects the port number and the ip address
	ipandport := ipaddr + ":" + portstr
	// Attempts to connect to the port.
	conn, err := net.DialTimeout("tcp", ipandport, 1*time.Second)
	// If we can not connect to the port
	if err != nil {
		fmt.Println(port, ": closed")
	} else {
		// If we can Connect to the port
		conn.Close()
		fmt.Println(port, ": open")
	}
	defer wg.Done()
}
