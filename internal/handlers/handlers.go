package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "../index.html")
}
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error receiving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	inputStr := string(data)
	convertedStr, err := service.ConverterToTextToMorse(inputStr)
	if err != nil {
		http.Error(w, "Conversion failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	timeStamp := time.Now().UTC().String()
	ext := filepath.Ext(handler.Filename)
	fileName := timeStamp + ext
	err = os.WriteFile(fileName, []byte(convertedStr), 0755)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(convertedStr))
}
