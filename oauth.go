package gomuse

import (
	"fmt"
	"golang.org/x/oauth2"
	"gopkg.in/jmcvetta/napping.v2"
	"net/http"
	"time"
)

type YMuseTokenSource struct {
	ClientId     string
	ClientSecret string
}

func (self YMuseTokenSource) Token() (*oauth2.Token, error) {
	result := struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   uint32 `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}{}
	payload := struct {
		GrantType    string `json:"grant_type"`
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{
		GrantType:    "client_credentials",
		ClientId:     self.ClientId,
		ClientSecret: self.ClientSecret,
	}
	err := struct {
		Error            string
		ErrorDescription string
	}{}
	session := napping.Session{
		Log: false,
	}
	response, errorcode := session.Post(getOauthTokenURL().String(), &payload, &result, &err)
	if errorcode != nil {
		return nil, errorcode
	}
	if response.Status() != 200 {
		return nil, fmt.Errorf("Error while retrieving token. Status: %d, Response: %s", response.Status(), response.RawText())
	}
	return &oauth2.Token{AccessToken: result.AccessToken, TokenType: result.TokenType, Expiry: secondsInFuture(result.ExpiresIn)}, nil
}

func getClient(cfg *oauth2.Config) *http.Client {
	return oauth2.NewClient(oauth2.NoContext, YMuseTokenSource{ClientId: cfg.ClientID, ClientSecret: cfg.ClientSecret})
}

func secondsInFuture(seconds uint32) time.Time {
	return time.Now().Add(time.Duration(seconds) * time.Second)
}
