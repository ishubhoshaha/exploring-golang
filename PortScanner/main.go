package PortScanner

import (
	"exploring-golang/PortScanner/port"
	"fmt"
)

func main()  {
	open := port.ScanPort("tcp", "localhost", 8000)
	fmt.Printf("Port %t\n", open)
	results := port.InitialScan("localhost")
	fmt.Println(results)
}