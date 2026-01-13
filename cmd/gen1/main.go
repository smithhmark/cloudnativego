package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"example.com/gorilla/internal/gen1"

	"github.com/gorilla/mux"
)

func helloMuxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello gorilla/mux!")
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = gen1.Put(key, string(value))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := gen1.Get(key)
	if err == gen1.ErrorNoSuchKey {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, value)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := gen1.Delete(key)
	if err == gen1.ErrorNoSuchKey {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloMuxHandler)
	r.HandleFunc("/v1/key/{key}", putHandler).Methods("PUT")
	r.HandleFunc("/v1/key/{key}", getHandler).Methods("GET")
	r.HandleFunc("/v1/key/{key}", deleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
