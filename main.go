package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/api/reboot", rebootRaspberryPi)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	log.Info().Msg("Starting server on port 3001...")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		panic(err)
	}
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
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
