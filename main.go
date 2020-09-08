package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// HandleFunc resister a function on a path
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		log.Printf("> base Root")
		d, err := ioutil.ReadAll(r.Body) // ioutil is Golang package to readAll , readFile , readDir , TempDir , TempFile ,  WriteFile , NopCloser , Discard.write
		log.Printf("Data is %s", d)
		log.Printf("Error is %s", err)

		fmt.Fprintf(rw, "Hello %s", d)

		// d, err := ioutil.ReadAll(r.Body)

		// if err != nil {
		// 	http.Error(rw, "Oops", http.StatusBadRequest)
		// 	return
		// }

		// fmt.Fprintf(rw, "Hello %s", d)

		// log.Printf("Data %s\n", d)
	})

	http.HandleFunc("/goodby", func(http.ResponseWriter, *http.Request) {
		log.Println("goodby World")
	})

	http.ListenAndServe(":9090", nil)

	//HTTP.LISTENandServe
	//it is a methord which construct a httpserver and registeres a default handler to it
	// first parameter is port and second parameter is handler .
	//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handler request on incoming connections
}

// http.HandlrFunc is a methord on go HTTP package , It registers a
//function to a path on a thing called default serve mux

//Serve Mux is a http handler

//DEFAULT HTTP MUX :
// It is a http handler and everything related to server in go is http handler

//**********************************************************************************************
//1. Default serve mux
// It is a httphandler which is  responsible for redirecting path , you resister a function at a path and servemux will determine which
// function will get executed
// HandleFunc registers the handler function for the given path in the DefaultServeMux

//***********************************************************************************************
//2. How to read and write to that request  ?
// We do it from using ResponseWritter and http.Request

// ResponseWritter : It is a interface which is used by the response handler or http handler  to construct the httpResponse
//ResponseWritter has a number of methords on it
// It can do like , writing headers  , response body , starus code  ,

// http.Request : A Request represents an HTTP request received
