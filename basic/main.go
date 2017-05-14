package basic

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	NWorkers = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen for HTTP request on")
)

func StartService() {

	flag.Parse()

	// Start Dispatcher
	fmt.Println("Starting Dispatcher")
	StartDispatcher(*NWorkers)

	// Register Collector
	fmt.Println("Starting the dispatcher")
	http.HandleFunc("/work", Collector)

	// Start the HTTP server
	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
