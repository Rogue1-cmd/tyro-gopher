package main

import (
	"fmt"
	"log"
	"net/http"
)


/*http
The http.ResponseWriter interface: used to write response data to the client.

The http.Request struct: represents the incoming HTTP request,
including method, URL, headers, etc


formHandler function: handles POST requests to the "/form" endpoint.
	>parses the form data from the request. 
	>Then, it retrieves the values of the "name" and "address" fields from the form data. 
	>Finally, it prints out these values as a response to the client.
	
helloHandler function: handles GET requests to the "/hello" endpoint. 
	>It checks if the request path matches "/hello" and if the HTTP method is GET. 
	>If both conditions are met, it sends a "Forrest||Forrest||Forrest"  message to the client.

*/
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
