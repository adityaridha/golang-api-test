package main

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

func getUsers(baseURL string) (*resty.Response, error) {
	client := resty.New()
	response, err := client.R().
		Get(baseURL + "/users?page=1")
	return response, err
}

func createUser(baseURL string, userData map[string]interface{}) (*resty.Response, error) {
	client := resty.New()
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(userData).
		Post(baseURL + "/users")
	return response, err
}

func printResponseInfo(response *resty.Response) {
	fmt.Printf("Response Status Code: %v\n", response.StatusCode())

	fmt.Println("Response Body:")
	fmt.Println(string(response.Body()))
}

func TestGetUsers(t *testing.T) {
	baseURL := "https://reqres.in/api"
	response, err := getUsers(baseURL)

	if err != nil {
		t.Fatalf("Error during the request: %v", err)
	}

	printResponseInfo(response)
}

func TestCreateUser(t *testing.T) {
	baseURL := "https://reqres.in/api"
	userData := map[string]interface{}{
		"name":  "John Doe",
		"job":   "Software Engineer",
		"email": "john.doe@example.com",
	}

	response, err := createUser(baseURL, userData)

	if err != nil {
		t.Fatalf("Error during the request: %v", err)
	}

	printResponseInfo(response)
}
