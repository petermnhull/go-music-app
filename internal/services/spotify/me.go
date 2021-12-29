package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/petermnhull/go-music-app/internal/services"
	"gopkg.in/validator.v2"
)

// SpotifyMe holds output from getting user associated with access token
type SpotifyMe struct {
	DisplayName string `json:"display_name" validate:"nonzero"`
	Country     string `json:"country"`
	Product     string `json:"product"`
}

func createMeRequest(accessToken string) (*http.Request, error) {
	url := baseURL + "me"
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	r.Header.Add("Content-Type", "application/json")
	return r, nil
}

// GetMe returns information on user corresponding to access token
func GetMe(
	httpclient services.HTTPClientInterface,
	accessToken string,
) (*SpotifyMe, error) {
	r, err := createMeRequest(accessToken)
	if err != nil {
		return nil, NewSpotifyRequestErr("failed to create request")
	}

	resp, err := httpclient.Do(r)
	if err != nil {
		return nil, NewSpotifyRequestErr("failed to do request")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, NewSpotifyResponseErr("failed to retrieve profile", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewSpotifyRequestErr("failed to parse response body")
	}
	var body SpotifyMe
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return nil, NewSpotifyRequestErr("failed to decode response body")
	}

	errs := validator.Validate(body)
	if errs != nil {
		return nil, NewSpotifyValidationErr("failed to validate profile response")
	}

	return &body, nil
}
