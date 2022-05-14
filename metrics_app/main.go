package main

import (
	"fmt"
	"net/http"
	"io"
	"math/rand"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// func h2(w http.ResponseWriter, _ *http.Request) {
	// io.WriteString(w, "Hello from a HandleFunc #2!\n")
	// fmt.Println("Nothing here, try reaching /metrics")
// }

var (
	goRandomValue = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_random_value",
		Help: "randomly generated value in Go",
	})
)

func main() {

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	var METRICS_PATH = "/metrics"
	var PORT = ":8080"
	fmt.Printf("Running metrics exporter on localhost:%s%s\n", PORT, METRICS_PATH)
	http.HandleFunc("/", Default)
	http.HandleFunc("/handle", handle)
	http.HandleFunc("/h1", h1)
	http.HandleFunc("/h2", h2)
	http.Handle(METRICS_PATH, promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}

func Default(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Nothing here\n")
}

func handle(w http.ResponseWriter, r *http.Request) {
	goRandomValue.Set(rand.Float64())
	fmt.Fprint(w, "Hello")
}