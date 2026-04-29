package main

import (
	"bytes"
	"embed"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"

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

	
	http.HandleFunc("/turn-on", func(w http.ResponseWriter, r *http.Request) {
		// Attempt to load .env, but don't check for 'err'
		_ = godotenv.Load()
	
		// Now pull your MAC address
		targetMAC := os.Getenv("MAC_ADDRESS")
		if targetMAC == "" {
			fmt.Println("Warning: MAC_ADDRESS environment variable is empty!")
		}

		fmt.Println("Sending magic packet to", targetMAC)
		
		command := "etherwake"
		arg1 := "-i"
		arg2 := "wlan0"
		arg3 := targetMAC

		fmt.Println("Running command:", command, arg1, arg2, arg3)

		cmd := exec.Command(command, arg1, arg2, arg3)
		err := cmd.Run()

		fmt.Println("Command ran successfully")

		if err != nil {
			fmt.Println("Error running command:", err)
		}
	})

	// listen on all interfaces for port 8080
	fmt.Println("Power Hub active on http://0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
