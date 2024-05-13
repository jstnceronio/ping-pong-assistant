package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")         // Your API Key
	modelEndpoint := "ping-pong-tracker/1" // Set model endpoint

	// Open file on disk.
	f, err := os.Open("test.png")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Encode as base64.
	data := base64.StdEncoding.EncodeToString(content)
	uploadURL := "https://detect.roboflow.com/" + modelEndpoint + "?api_key=" + apiKey + "&name=test.png"

	req, err := http.NewRequest("POST", uploadURL, strings.NewReader(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println(string(bytes))
}
