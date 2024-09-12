package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	clientID := "Client Id from https://dashboard.aspose.cloud/applications"
	clientSecret := "Client Secret from https://dashboard.aspose.cloud/applications"

	baseURL := "https://id.aspose.cloud/"
	endpoint := "connect/token"

	payload := []byte(fmt.Sprintf(
		"grant_type=client_credentials&client_id=%s&client_secret=%s",
		clientID, clientSecret,
	))

	req, err := http.NewRequest("POST", baseURL+endpoint, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("Error creating request: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("HTTP request error: %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("HTTP error occurred: %v\n", resp.Status)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalln("Error decoding response: %v\n", err)
	}

	fmt.Println("Token reciewed successfully")
	// To view token uncomment next line
	// fmt.Println(data["access_token"])

}
