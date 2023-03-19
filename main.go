package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os/exec"

	"github.com/rs/zerolog/log"
)

//go:embed static
var staticFiles embed.FS

//go:embed templates/index.html
var indexTemplate embed.FS

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/api/reboot", rebootRaspberryPi)

	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		panic(err)
	}
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.FS(staticFS))))

	log.Info().Msg("Starting server on port 3001...")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		panic(err)
	}
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(indexTemplate, "templates/index.html")
	if err != nil {
		log.Error().Err(err).Msg("Failed to load template")
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to render template")
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func rebootRaspberryPi(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Error().Msg("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cmd := exec.Command("sudo", "reboot")
	if err := cmd.Run(); err != nil {
		log.Error().Err(err).Msg("Failed to reboot Pi")
		http.Error(w, "Failed to reboot Pi", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Rebooting...")
}
