package routes

import (
	"net/http"
	"strconv"

	"golang.org/x/exp/rand"
)

var FileCache = make(map[int][]byte)

func Index(w http.ResponseWriter, r *http.Request) {
	index := rand.Intn(len(FileCache))
	selectedImage := FileCache[index]

	w.Header().Set("Content-Type", http.DetectContentType(selectedImage))
	w.Header().Set("Content-Length", strconv.Itoa(len(selectedImage)))
	w.WriteHeader(http.StatusOK)

	w.Write(selectedImage)
}
