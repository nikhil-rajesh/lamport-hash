package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"l-hash-frontend/models"
	"log"
	"net/http"

	"github.com/ttacon/chalk"
)

func Login() {
	// Get username and password from the user
	var details = models.LoginDTO{}
	fmt.Printf("%sUsername :: %s", chalk.Blue, chalk.Reset)
	fmt.Scanf("%s", &details.Username)
	fmt.Printf("%sPassword :: %s", chalk.Blue, chalk.Reset)
	fmt.Scanf("%s", &details.Hash)
	fmt.Println()

	// Fetch N from backend using the username
	ok, n, salt := GetN(details.Username)
	if !ok {
		return
	}

	// Hash the password and salt (if present) n times
	details.Hash = Hash(details.Hash+salt, n-1)
	data, err := json.Marshal(details)

	// Send the hash to the backend for authentication
	resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err.Error)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	// Print success or error message. HTTP status code 200 is success
	if resp.StatusCode == 200 {
		fmt.Printf("%sSuccess :: %s%s\n", chalk.Green, chalk.Reset, body)
	} else {
		fmt.Printf("%sError :: %s%s\n", chalk.Red, chalk.Reset, body)
	}
}

func GetN(username string) (bool, int, string) {
	var details = models.GetNDTO{username}
	data, err := json.Marshal(details)
	fmt.Printf("%sSending request for N ...%s\n", chalk.Green, chalk.Reset)

	// Send request to backend for user details
	resp, err := http.Post("http://localhost:8080/login/n", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err.Error)
	}
	defer resp.Body.Close()

	// Return details if request was success. HTTP status code 200 is success
	// Else print Error
	if resp.StatusCode == 200 {
		var res models.GetNResponse
		json.NewDecoder(resp.Body).Decode(&res)
		if res.Salt == "" {
			fmt.Printf("%sServer Response :: %sN = %d\n", chalk.Green, chalk.Reset, res.N)
		} else {
			fmt.Printf("%sServer Response :: %sN = %d, Salt = %s\n\n", chalk.Green, chalk.Reset, res.N, res.Salt)
		}
		return true, res.N, res.Salt
	} else {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("%sError :: %s%s\n", chalk.Red, chalk.Reset, body)
		return false, 0, ""
	}
}
