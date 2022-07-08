package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// var secret = os.Getenv("DISCORD_OAUTH_SECRET")
// var cID = os.Getenv("DISCORD_OAUTH_CLIENT_ID")

var secret = "EYYKi7zFusvMV-AcA8bbEsUUBiouBqBp"
var cID = "994229248392974366"
var router = gin.Default()

func main() {
	api := router.Group("/api")
	api.GET("auth/discord/redirect", callback)

	println(secret)
	println(cID)

	router.Use(static.Serve("/", static.LocalFile("./view/build", true)))
	router.Run(":5000")
}

type AuthRequestBody struct {
	client_id     string
	client_secret string
	grant_type    string
	code          string
	redirect_uri  string
}
type AuthTokenResponse struct {
	access_token  string
	token_type    string
	expires_in    string
	refresh_token string
	scope         string
}

func callback(c *gin.Context) {
	code := c.Request.URL.Query().Get("code")
	if code != "" {
		//authBody := AuthRequestBody{cID, secret, "authorization_code", code, "http://localhost:5000/api/auth/discord/redirect"}
		data := map[string]interface{}{}
		byte1, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		client := http.Client{}
		r := strings.NewReader(string(byte1))
		req, err := http.NewRequest("POST", "https://discord.com/api/v8/oauth2/token", r)
		if err != nil {
			log.Fatal(err)
		}
		q := req.URL.Query()
		q.Add("client_id", cID)
		q.Add("client_secret", secret)
		q.Add("grant_type", "authorization_code")
		q.Add("code", code)
		q.Add("redirect_uri", "http://localhost:5000/api/auth/discord/redirect")

		req.URL.RawQuery = q.Encode()
		println(string(req.URL.String()))

		req.Header = http.Header{
			"Content-type": {"application/x-www-form-urlencoded"},
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close() // close response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		println(string(body))
		println(string(body))

		// bodyAsByteArray, _ := ioutil.ReadAll(c.resp.Body)
		// jsonBody := string(bodyAsByteArray)
		// println(jsonBody)
		// var cringe AuthTokenResponse

		// json.Unmarshal(resp.GetBody(), &cringe)
		// println(cringe.access_token)
	}

}
