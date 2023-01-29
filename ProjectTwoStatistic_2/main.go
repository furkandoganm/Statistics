package main

import (
	"net/http"
	"projects/ProjectTwoStatistic_2/service"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/addvideos", service.AddData)

	http.ListenAndServe(":8080", mux)
}
