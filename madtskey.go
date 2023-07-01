package madtskey

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/clientcredentials"
)

// Read the .env file in the current directory
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// This function will create the API Key for use in joining new machines
// to the tailnet.
func CreateAPIKey(expirySeconds int, description string, tags []string) (*TSResponse, error) {
	// Using the Tailscale OAuth client
	var oauthConfig = &clientcredentials.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		TokenURL:     "https://api.tailscale.com/api/v2/oauth/token",
	}

	client := oauthConfig.Client(context.Background())
	// We describe the type of key that we want while taking the
	// expiry and key description from the function itself.
	//TODO: We should also take the tags as a param.
	cap := &Req{
		Capabilities{
			Devices{
				Create{
					// Adjust these settings to suit your preffered setup
					Reusable:      false,
					Ephemeral:     true,
					Preauthorized: true,
					Tags:          tags,
				},
			},
		},
		expirySeconds,
		description,
	}

	// Prepare to make the HTTP POST Request to create the key. First
	// creat the url to which we POST, then make the actual request.
	reqUrl := fmt.Sprintf("https://api.tailscale.com/api/v2/tailnet/%s/keys", os.Getenv("TAILNET"))
	req, err := http.NewRequest("POST", reqUrl, cap.AsReader())
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting keys: %v", err)
	}

	// TODO: We should probably check for a non http 200 error code
	// response.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	// Take the response body and unmarshal it back to a struct
	// so that we can present it to the caller
	var tsResponse TSResponse
	err = json.Unmarshal(body, &tsResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %v", err)
	}
	return &tsResponse, nil
}
