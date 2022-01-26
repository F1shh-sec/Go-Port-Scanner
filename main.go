package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
Author: F1shh-sec
SOURCE: https://github.com/F1shh-sec/Go-Port-Scanner
USAGE go run main.go protocol target_ip Starting_port Ending_port
*/

func main() {
	argsnum := len(os.Args)
	if argsnum < 2 {
		// if there are no arguments provided
		noArgs()
	} else if argsnum >= 2 {
		// USAGE go run main.go protocol target_ip Starting_port Ending_port
		withArgs(argsnum)
	}

}
func withArgs(argsnum int) {
	// if there are
	prot := os.Args[1]
	ipa := os.Args[2]
	startport, _ := strconv.Atoi(os.Args[3])
	endport, _ := strconv.Atoi(os.Args[4])
	
	runScan(prot, ipa, startport, endport)

}
func runScan(prot string, ipa string, startport int, endport int) {
	var wg sync.WaitGroup

	if endport > startport {
		for i := startport; i <= endport; i++ {
			wg.Add(1)
			go connect(ipa, prot, i, &wg)
		}
	} else {
		fmt.Println("You have a larger starting port then end port...")
	}
	wg.Wait()
}
func noArgs() {
	var ipa, prot string
	var startport, endport int
	// Gets the protocol
	fmt.Print("Enter a Protocol: ")
	fmt.Scanf("%s\n", &prot)
	prot = strings.Replace(prot, "\n", "", -1)

	// Gets input Ip address
	fmt.Print("Enter a host IP: ")
	fmt.Scanf("%s\n", &ipa)
	ipa = strings.Replace(ipa, "\n", "", -1)

	// Grabs starting port
	fmt.Print("Enter a starting port: ")
	fmt.Scanf("%d\n", &startport)

	// grabs ending port
	fmt.Print("Enter a ending port: ")
	fmt.Scanf("%d\n", &endport)

	runScan(prot, ipa, startport, endport)
}

/**
Attempts to connect to a given port.
*/
func connect(ipaddr string, prot string, port int, wg *sync.WaitGroup) {
	// Converts the int to a string
	portstr := strconv.Itoa(port)
	// Connects the port number and the ip address
	ipandport := ipaddr + ":" + portstr
	// Attempts to connect to the port.
	conn, err := net.DialTimeout(prot, ipandport, 1*time.Second)
	// If we can not connect to the port

	if err != nil {
		if prot == "tcp" {
			//tpc
			fmt.Println(port, ": closed/filtered")
		} else {
			//udp
			fmt.Println(port, ": closed/filtered")
		}

	} else {
		// If we can Connect to the port
		conn.Close()
		if prot == "tcp" {
			//tcp
			fmt.Println(port, ": open")
		} else {
			//udp
			fmt.Println(port, ": open/Filtered")
		}

	}
	defer wg.Done()
}
