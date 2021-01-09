package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"time"
)

type AuthTokens struct {
	Bearer       string
	Entitlements string
	UserId       string
	// Expiry // todo
}

func createSession(client *http.Client) error {
	url := "https://auth.riotgames.com/api/v1/authorization"
	body := []byte(`{
		"client_id": "play-valorant-web-prod",
		"nonce": "1", 
		"redirect_uri": "https://playvalorant.com/opt_in", 
		"response_type": "token id_token", 
		"scope": "account openid"
	}`)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	return nil // todo error handling
}

func authenticate(client *http.Client, username, password string) (string, error) {
	url := "https://auth.riotgames.com/api/v1/authorization"
	body := []byte(fmt.Sprintf(`{
		"type": "auth",
		"username": "%s",
		"password": "%s"
	}`, username, password))

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	respBody, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// todo handle "error":"auth_failure"

	// extract accessToken from response
	// resp := `{"type":"response","response":{"mode":"fragment","parameters":{"uri":"https://playvalorant.com/opt_in#access_token=eyJraWQiOiJzMSIsInR5cCI6ImF0K2p3dCIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIwY2FhZGVlYy05MmViLTVjMGQtODdiYy1hMThiMjQwN2ViZTUiLCJzY3AiOlsiYWNjb3VudCIsIm9wZW5pZCJdLCJjbG0iOlsib3BlbmlkIiwicHciLCJyZ25fRVVXMSIsImFjY3RfZ250IiwiYWNjdCIsIiFGQUFDIl0sImRhdCI6eyJwIjoiIiwiciI6IkVVVzEiLCJjIjoiZWMxIiwidSI6MjM5NjA4OTY1MzQ5MDMwNH0sImlzcyI6Imh0dHBzOlwvXC9hdXRoLnJpb3RnYW1lcy5jb20iLCJleHAiOjE2MTAyMjMyMDEsImlhdCI6MTYxMDIxOTYwMSwianRpIjoiQ3RaU1NvNGthcDQiLCJjaWQiOiJwbGF5LXZhbG9yYW50LXdlYi1wcm9kIn0.KfbENaO-aPc3eg4G6_cSvt3bAQ5ZYViCqg1Llul1jLDc5LBG8bfFhv8KsqCXMV0jFycbhPqCf9w_dBPuzAxYw0JER9rbEUdF0k40-ruSKC8ycfMNo1yWJxjWfngtaXngKVd7KYgm6ynROI9nJlK4HVtQn-ROr6tbx4xLOIZERbA&scope=account+openid&id_token=eyJraWQiOiJzMSIsImFsZyI6IlJTMjU2In0.eyJhdF9oYXNoIjoia2JpMGZ4S0piemtXWUZaQ2pqSHdSdyIsInN1YiI6IjBjYWFkZWVjLTkyZWItNWMwZC04N2JjLWExOGIyNDA3ZWJlNSIsImNvdW50cnkiOiJkZXUiLCJjb3VudHJ5X2F0IjoxNTgxMzI2Mzg2MDAwLCJhbXIiOlsicGFzc3dvcmQiXSwiaXNzIjoiaHR0cHM6XC9cL2F1dGgucmlvdGdhbWVzLmNvbSIsImxvbCI6W3siY3VpZCI6MjM5NjA4OTY1MzQ5MDMwNCwiY3BpZCI6IkVVVzEiLCJ1aWQiOjIzOTYwODk2NTM0OTAzMDQsInVuYW1lIjoibWFtbXV0aDAiLCJwdHJpZCI6bnVsbCwicGlkIjoiRVVXMSIsInN0YXRlIjoiRU5BQkxFRCJ9XSwibm9uY2UiOiIxIiwiYXVkIjoicGxheS12YWxvcmFudC13ZWItcHJvZCIsImFjciI6InVybjpyaW90OmJyb256ZSIsInBsYXllcl9sb2NhbGUiOiJlbiIsImV4cCI6MTYxMDMwNjAwMSwiaWF0IjoxNjEwMjE5NjAxLCJhY2N0Ijp7ImdhbWVfbmFtZSI6Ik1BTU1VVEgiLCJ0YWdfbGluZSI6IjAwMDAifSwianRpIjoia0Q0dWgzelR6N1kiLCJsb2dpbl9jb3VudHJ5IjoiZGV1In0.fU03fPI9mAoktaYSFtd_4jD5dtrS8pNveQgOZxGSr-KU9w_c0gXpyRkICeidd-IIO9Ub1ZyDNf_dZRzCyset_IoILJZuegsXsHS8VOG6fpUpuOxhgv-S6-FT0Nw6r_YBtmkArpaIQnHvJ_XavNRnhFgA67BCDyEG7lGbHirXLTY&token_type=Bearer&expires_in=3600"}},"country":"deu"}`
	re := regexp.MustCompile(`(?m)access_token=(.*)&scope`)

	matches := re.FindStringSubmatch(string(respBody))
	accessToken := matches[1]
	if len(matches) != 2 {
		return "", errors.New("Couldn't parse access token from response")
	}
	return accessToken, nil // todo error handling
}

type EntitlementsResponse struct {
	Token string `json:"entitlements_token"`
}

func getEntitlements(client *http.Client, accessToken string) (string, error) {
	url := "https://entitlements.auth.riotgames.com/api/token/v1"
	body := []byte(`{}`)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	respBody, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var entitlementsResponse EntitlementsResponse
	json.Unmarshal([]byte(respBody), &entitlementsResponse)
	if entitlementsResponse.Token == "" {
		return "", errors.New("Couldn't parse entitlements token")
	}

	return entitlementsResponse.Token, nil
}

func getUser(client *http.Client, accessToken string) (string, error) {
	url := "https://auth.riotgames.com/userinfo"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBody, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var respJson map[string]interface{}
	if err := json.Unmarshal(respBody, &respJson); err != nil {
		log.Fatal(err)
	}
	userId, ok := respJson["sub"].(string)
	if ok != true {
		log.Fatal("Cannot parse entitlements")
	}
	// todo error handling
	return userId, nil
}

func Login(username, password string) (AuthTokens, error) {
	cookieJar, _ := cookiejar.New(nil)
	client := http.Client{
		Timeout: 10 * time.Second,
		Jar:     cookieJar,
	}
	createSession(&client) // todo comment in
	accessToken, _ := authenticate(&client, username, password)
	// todo remove hardcoded
	// accessToken := "eyJraWQiOiJzMSIsInR5cCI6ImF0K2p3dCIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIwY2FhZGVlYy05MmViLTVjMGQtODdiYy1hMThiMjQwN2ViZTUiLCJzY3AiOlsiYWNjb3VudCIsIm9wZW5pZCJdLCJjbG0iOlsib3BlbmlkIiwicHciLCJyZ25fRVVXMSIsImFjY3RfZ250IiwiYWNjdCIsIiFGQUFDIl0sImRhdCI6eyJwIjoiIiwiciI6IkVVVzEiLCJjIjoiZWMxIiwidSI6MjM5NjA4OTY1MzQ5MDMwNH0sImlzcyI6Imh0dHBzOlwvXC9hdXRoLnJpb3RnYW1lcy5jb20iLCJleHAiOjE2MTAyMjY5NjEsImlhdCI6MTYxMDIyMzM2MSwianRpIjoiQW1EWld6R2xseWsiLCJjaWQiOiJwbGF5LXZhbG9yYW50LXdlYi1wcm9kIn0.itSGuzSZuKHQYw5rBbRC-OyxHaDr0O4os62m2wELFpvOXBVECDkPESg63aIv2FYbY1gL4rtcc2NCL0VRGsIvI1_5NrXPqtye2_m4e5A3GWdKwFgrKg3dyuZ1bvUb34ePj9wi6XtmaH9apbXHXGlwrMKVzPctXeP1yy8rPg5OYN8"

	entitlements, _ := getEntitlements(&client, accessToken)
	userId, _ := getUser(&client, accessToken)

	tokens := AuthTokens{
		Bearer:       accessToken,
		Entitlements: entitlements,
		UserId:       userId,
	}

	return tokens, nil
}
