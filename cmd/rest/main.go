package main

import (
	"fmt"
	"net/http"

	"github.com/markrtucker/chorehound/chores"
)

var fs chores.Repository

func main() {

	fs = chores.NewFileSystemRepository("/Users/marktucker/go/chores/")

	// TODO Use gorilla mux router
	http.HandleFunc("/chores", doChores)
	// http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

func doChores(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		// TODO proper parsing
		c := chores.Chore{
			ID:       "something",
			Name:     "Test",
			Schedule: "some-schedule-todo",
		}

		err := c.SaveTo(req.Context(), fs)
		if err != nil {
			fmt.Printf("%v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Chores")

	} else {
		w.WriteHeader(http.StatusNotFound)

	}

}
