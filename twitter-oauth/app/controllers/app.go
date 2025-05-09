package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/mrjones/oauth"
	"github.com/wiselike/revel-examples/twitter-oauth/app/models"
	"github.com/wiselike/revel"
)

var TWITTER = oauth.NewConsumer(
	"VgRjky4dTA1U2Ck16MmZw",
	"l8lOLyIF3peCFEvrEoTc8h4oFwieAFgPM6eeTRo30I",
	oauth.ServiceProvider{
		AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
		RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
		AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
	},
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	user := getUser()
	if user.AccessToken == nil {
		return c.Render()
	}

	// We have a token, so look for mentions.
	resp, err := TWITTER.Get(
		"https://api.twitter.com/1.1/statuses/mentions_timeline.json",
		map[string]string{"count": "10"},
		user.AccessToken)
	if err != nil {
		c.Log.Error("mentions error", "error", err)
		return c.Render()
	}
	defer resp.Body.Close()

	// Extract the mention text.
	mentions := []struct {
		Text string `json:text`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&mentions)
	if err != nil {
		c.Log.Error("JSON error", "error", err)
	}
	c.Log.Info("Success", "mentions", mentions)
	return c.Render(mentions)
}

func (c Application) SetStatus(status string) revel.Result {
	resp, err := TWITTER.PostForm(
		"http://api.twitter.com/1.1/statuses/update.json",
		map[string]string{"status": status},
		getUser().AccessToken,
	)
	if err != nil {
		c.Log.Error(err.Error())
		return c.RenderError(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	c.Log.Info(string(bodyBytes))
	c.Response.ContentType = "application/json"
	return c.RenderText(string(bodyBytes))
}

// Twitter authentication

func (c Application) Authenticate(oauth_verifier string) revel.Result {
	user := getUser()
	if oauth_verifier != "" {
		// We got the verifier; now get the access token, store it and back to index
		accessToken, err := TWITTER.AuthorizeToken(user.RequestToken, oauth_verifier)
		if err == nil {
			user.AccessToken = accessToken
		} else {
			c.Log.Error("Error connecting to twitter:", "error", err)
		}
		return c.Redirect(Application.Index)
	}

	requestToken, url, err := TWITTER.GetRequestTokenAndUrl("http://127.0.0.1:9000/Application/Authenticate")
	if err == nil {
		// We received the unauthorized tokens in the OAuth object - store it before we proceed
		user.RequestToken = requestToken
		return c.Redirect(url)
	} else {
		c.Log.Error("Error connecting to twitter:", "error", err)
	}
	return c.Redirect(Application.Index)
}

func getUser() *models.User {
	return models.FindOrCreate("guest")
}

func init() {
	TWITTER.Debug(true)
}
