package main

import (
	"net/http"

	"github.com/kxrxh/queue-master/internal/metrics"
)

func main() {
	metrics.SetupMetrics()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", nil)
}
