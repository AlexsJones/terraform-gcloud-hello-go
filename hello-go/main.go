package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var port string

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
func sayHi(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello from "+getLocalIP())
}

func main() {
	port = "9090"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	http.HandleFunc("/", sayHi) // set router
	fmt.Println("Running on " + getLocalIP() + ":" + port)
	err := http.ListenAndServe(":"+port, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
