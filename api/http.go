package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"dawg.pics/api/routes"
	"dawg.pics/modules/env"
)

func StartWebServer() {
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		err := os.Mkdir("uploads", 0755)
		if err != nil {
			log.Fatalf("Failed to create uploads directory: %v\n", err)
		}
	}

	userHandler("/", routes.Index, http.MethodGet)

	log.Printf("Starting REST API server on %s:%s\n", env.GetEnv("HOST"), env.GetEnv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", env.GetEnv("HOST"), env.GetEnv("PORT")), nil))
}

func enableCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

		// preflight -> OK
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	}
}

func userHandler(path string, handler http.HandlerFunc, method string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"message": "Method not allowed"}`))
			return
		}

		enableCORS(handler)(w, r)
	})
}
