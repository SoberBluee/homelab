package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"

	"github.com/da-rod/wakeonlan"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

// This tells Go to include these files in the final binary
//
//go:embed dashboard.html dashboard.css
var content embed.FS

type PageData struct {
	Status string
}

// runs the ssh shutdown command
func runSSHShutdown() {
	// load the .env file
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	sshUser := os.Getenv("SSH_USER")
	sshPassword := os.Getenv("SSH_PASSWORD")
	sshHost := os.Getenv("SSH_HOST")
	sshPort := os.Getenv("SSH_PORT")

	if sshUser == "" || sshPassword == "" || sshHost == "" || sshPort == "" {
		fmt.Println("Error: SSH_USER, SSH_PASSWORD, SSH_HOST, or SSH_PORT environment variables are empty")
		fmt.Println(sshUser, sshPassword, sshHost, sshPort)
	}

	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		// WARNING: For production, implement proper HostKeyCallback
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort(sshHost, sshPort), config)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}
	defer session.Close()

	var stdout bytes.Buffer
	session.Stdout = &stdout

	command := "shutdown -h now"

	// Run the command
	err = session.Run(command)
	if err != nil {
		fmt.Println("Error running command:", err)
	}

	fmt.Println("Shutdown command executed:", stdout.String())
}

func main() {
	// Serve the CSS file
	http.Handle("/static/", http.FileServer(http.FS(content)))

	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		runSSHShutdown()
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("Shutdown command sent")
	})

	// Attempt to load .env, but don't check for 'err'
	_ = godotenv.Load()

	// Now pull your MAC address
	targetMAC := os.Getenv("MAC_ADDRESS")
	if targetMAC == "" {
		fmt.Println("Warning: MAC_ADDRESS environment variable is empty!")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status := ""
		// check if its a post request
		if r.Method == http.MethodPost {
			fmt.Println("Sending magic packet")
			// create a new packet with WOL giving the target mac address
			mp, err := wakeonlan.NewMagicPacket(targetMAC)
			fmt.Println("Magic packet created")
			// nil means - representation of zero value of the type
			// so below means if no errors, send packages
			if err == nil {
				fmt.Println("Sending magic packet")
				mp.Send()
				fmt.Println("Magic packet sent")
				status = "Magic Packet Sent!"
			} else {
				status = "Error: " + err.Error()
			}
		}

		// FIX: Removed the stray backtick from the filename
		tmpl, err := template.ParseFS(content, "dashboard.html")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			http.Error(w, "Template not found", 500)
			return
		}

		// Execute and check for errors
		err = tmpl.Execute(w, PageData{Status: status})
		if err != nil {
			fmt.Println("Error executing template:", err)
		}
		fmt.Println("Response sent")
	})

	// listen on all interfaces for port 8080
	fmt.Println("Power Hub active on http://0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
