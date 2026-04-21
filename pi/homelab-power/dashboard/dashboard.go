package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/da-rod/wakeonlan"
)

// This tells Go to include these files in the final binary
//
//go:embed dashboard.html dashboard.css
var content embed.FS

type PageData struct {
	Status string
}

const targetMAC = "a4:ae:11:17:26:b3"

func main() {
	// Serve the CSS file
	http.Handle("/static/", http.FileServer(http.FS(content)))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status := ""
		if r.Method == http.MethodPost {
			mp, err := wakeonlan.NewMagicPacket(targetMAC)
			if err == nil {
				mp.Send()
				status = "Magic Packet Sent!"
			} else {
				status = "Error: " + err.Error()
			}
		}

		// Parse the embedded template
		tmpl, _ := template.ParseFS(content, "dashboard.html`")
		tmpl.Execute(w, PageData{Status: status})
	})

	fmt.Println("Power Hub active on http://0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
