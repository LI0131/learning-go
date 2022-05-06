/*
 * This module defines a simple file server for returning static
 * files back to the requesting user. This could be helpful for
 * creating a web application and serving images.
 */

package main

import (
	"fmt"
	"net/http"

	"github.com/kardianos/osext"
)

func main() {
	fp, err := osext.ExecutableFolder()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf(fp)
		fs := http.FileServer(http.Dir(fmt.Sprintf("%s/static/", fp)))

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "This is the main endpoint.\n")
		})

		http.Handle("/static/", http.StripPrefix("/static/", fs))

		http.ListenAndServe(":8080", nil)
	}
}
