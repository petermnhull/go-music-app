package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/petermnhull/go-music-app/pkg"
	"gopkg.in/validator.v2"
)

// SpotifyAccess holds tokens and metadata for user
type SpotifyAccess struct {
	AccessToken  string `json:"access_token" validate:"nonzero"`
	RefreshToken string `json:"refresh_token" validate:"nonzero"`
	Scope        string `json:"scope" validate:"nonzero"`
	TokenType    string `json:"token_type" validate:"nonzero"`
	ExpiresIn    int64  `json:"expires_in" validate:"nonzero"`
}

func createAccessRequest(code string, redirectURI string, clientID string, clientSecret string) (*http.Request, error) {
	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", code)
	form.Add("redirect_uri", redirectURI)
	r, err := http.NewRequest(http.MethodPost, authURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	credentials := []byte(fmt.Sprintf("%s:%s", clientID, clientSecret))
	encoded := base64.StdEncoding.EncodeToString(credentials)
	auth := fmt.Sprintf("Basic %s", encoded)
	r.Header.Add("Authorization", auth)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return r, nil
}

// GetAccess completes OAuth2 flow given an access code
func GetAccess(
	httpclient pkg.HTTPClient,
	code string,
	redirectURI string,
	clientID string,
	clientSecret string,
) (*SpotifyAccess, error) {
	r, err := createAccessRequest(code, redirectURI, clientID, clientSecret)
	if err != nil {
		return nil, NewSpotifyRequestErr("failed to create request")
	}
	resp, err := httpclient.Do(r)
	if err != nil {
		return nil, NewSpotifyRequestErr("failed to do request")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, NewSpotifyResponseErr("failed to retrieve tokens", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewSpotifyRequestErr("failed to parse response")
	}
	var body SpotifyAccess
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return nil, NewSpotifyRequestErr("failed to decode response")
	}

	errs := validator.Validate(body)
	if errs != nil {
		return nil, NewSpotifyValidationErr("failed to validate access response")
	}

	return &body, nil
}
