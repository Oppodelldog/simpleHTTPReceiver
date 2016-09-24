package main

import (
	"flag"
	"fmt"
	"os"
	"simpleHTTPReceiver/http_receiver"
)

func main() {
	port := flag.Int("port", 80, "server port to bind to")
	response := flag.String("response", "YEAH MON!!", "string the server will return for every request")
	help := flag.Bool("h", false, "help")
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Printf("Starting HTTPReceiver on port %v\n", *port)
	http_receiver.StartHTTPReceiver(*port, *response)
}
