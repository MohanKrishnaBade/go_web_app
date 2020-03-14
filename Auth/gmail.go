package Auth

import (
	"net/http"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"os"
	"golang.org/x/oauth2/google"
	"math/rand"
	"github.com/go_web_app/db"
	"github.com/go_web_app/models"
	"github.com/go_web_app/Errors"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  string
)

func SetConfig() {
	oauthStateString = randomString(9)
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GetUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := GetUserInfo(r.FormValue("state"), r.FormValue("code"))

	fmt.Println(content)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	gUser, err := db.ParseUser(content)
	fmt.Println(gUser)
	Errors.PanicError(err)
	sUser, err := db.FindByEmail(gUser.Email)
	fmt.Println(sUser)

	user := models.User{}
	if err != nil {
		fmt.Println("create")
		db.CreateUser(gUser)
		user = gUser
	} else {
		fmt.Println("update")
		user = db.Merge(gUser, sUser)
		db.Update(user)

	}
	helper(w, r, user)
}
