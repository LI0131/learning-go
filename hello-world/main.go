/* This module defines a simple HTTP server for responding with
 * info about the user request
 */

package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Recieved request data %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/test", handle)
	http.ListenAndServe(":80", nil)
}
