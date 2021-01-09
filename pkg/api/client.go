package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mammuth/valorant-go/pkg/auth"
)

type Client struct {
	authTokens auth.AuthTokens
	region     string
}

func (client *Client) Request(method, url string, body []byte, obj interface{}) error {
	httpClient := http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("client request: %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+client.authTokens.Bearer)
	req.Header.Add("X-Riot-Entitlements-JWT", client.authTokens.Entitlements)

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("submitting request: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("fetching request: %w", err)
	}
	// fmt.Println(string(respBody))
	if err := json.Unmarshal(respBody, obj); err != nil {
		return fmt.Errorf("unmarshalling response: %w", err)
	}

	return nil
}

// NewClient creates a new client instance
func NewClient(region, username, password string) (*Client, error) {
	authTokens, err := auth.Login(username, password)
	if err != nil {
		log.Fatal(err)
	}

	client := &Client{
		authTokens: authTokens,
		region:     "eu", // todo
	}
	return client, nil
}
