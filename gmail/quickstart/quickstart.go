/**
 * @license
 * Copyright Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// [START gmail_quickstart]
package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// generateState creates a secure random string for the OAuth2 state parameter.
func generateState() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// Request a token from the web, then returns the retrieved token.
// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatalf("Unable to start local listener: %v", err)
	}

	port := listener.Addr().(*net.TCPAddr).Port
	config.RedirectURL = fmt.Sprintf("http://localhost:%d", port)

	// Generate a dynamic, cryptographically secure state parameter
	state := generateState()

	authURL := config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser: \n%v\n", authURL)

	// Channels to handle successful codes and errors
	codeCh := make(chan string)
	errCh := make(chan error)

	m := http.NewServeMux()
	server := &http.Server{Handler: m}

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if errStr := r.URL.Query().Get("error"); errStr != "" {
			fmt.Fprintf(w, "Authentication error: %s. You may close this window.", errStr)
			errCh <- errors.New(errStr)
			return
		}

		returnedState := r.URL.Query().Get("state")
		if returnedState != state {
			fmt.Fprintf(w, "Security error: invalid state parameter. You may close this window.")
			errCh <- errors.New("invalid state parameter")
			return
		}

		code := r.URL.Query().Get("code")
		if code != "" {
			fmt.Fprintf(w, "Authentication successful! You may close this window.")
			codeCh <- code
		} else {
			fmt.Fprintf(w, "Failed to get authorization code. You may close this window.")
			errCh <- errors.New("authorization code missing")
		}
	})

	go func() {
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Unable to start local web server: %v", err)
		}
	}()

	var authCode string

	// Wait for a successful code, an error, or a timeout
	select {
	case authCode = <-codeCh:
		// Success case, proceed to shutdown and exchange
	case callbackErr := <-errCh:
		server.Shutdown(context.Background())
		log.Fatalf("Authorization failed during callback: %v", callbackErr)
	case <-time.After(3 * time.Minute):
		// Timeout case: user took too long or closed the browser
		server.Shutdown(context.Background())
		log.Fatalf("Authorization timed out after 3 minutes. Please try again.")
	}

	// Shutdown the server gracefully upon success
	server.Shutdown(context.Background())

	// Exchange the authorization code for an access token using context.Background()
	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	r, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	if len(r.Labels) == 0 {
		fmt.Println("No labels found.")
		return
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
		fmt.Printf("- %s\n", l.Name)
	}
}

// [END gmail_quickstart]
