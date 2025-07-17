package routes

import (
	"net/http"
	"strconv"
	"sync/atomic"
)

type ApiConfig struct {
	Hits atomic.Int32
}

func FtpRouter(c *ApiConfig) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/app/", c.middlewareMetricInc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/app", http.FileServer(http.Dir("client"))).ServeHTTP(w, r)
	})))

	router.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`Hits: ` + strconv.Itoa(int(c.Hits.Load()))))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	router.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		c.Hits.Store(0)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"message": "Hits counter reset"}`))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	router.HandleFunc("/app/assets/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/app/assets/", http.FileServer(http.Dir("assets"))).ServeHTTP(w, r)
	})

	router.HandleFunc("/", readinessHandler)

	return router
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}

func (c *ApiConfig) middlewareMetricInc(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Hits.Add(1)
		next.ServeHTTP(w, r)
	})
}
