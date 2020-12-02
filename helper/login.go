package helper

/*
import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

var OauthConfGG = &oauth2.Config{}

var OauthConfFB = &oauth2.Config{}

type ResSocialLogin struct {
	ID        uint64                            `json:"id"`
	FirstName string                            `json:"first_name"`
	LastName  string                            `json:"last_name"`
	Email     string                            `json:"email"`
	Picture   map[string]map[string]interface{} `json:"picture"`
}

var OauthStateString = utils.RandStringBytes(32)

func LoginInWithFacebook() string {
	OauthConfFB = &oauth2.Config{
		ClientID:     utils.FBClientID,
		ClientSecret: utils.FbClientSecret,
		Endpoint:     endpoints.Facebook,
		RedirectURL:  utils.BeBaseUrl + utils.FACEBOOK_REDIRECT_URL,
		Scopes:       []string{"email", "public_profile"},
	}
	Url, err := url.Parse(OauthConfFB.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", OauthConfFB.ClientID)
	parameters.Add("scope", strings.Join(OauthConfFB.Scopes, " "))
	parameters.Add("redirect_uri", OauthConfFB.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", OauthStateString)
	Url.RawQuery = parameters.Encode()
	urlRes := Url.String()
	return urlRes
}

func LoginInWithGoogle() string {
	OauthConfGG = &oauth2.Config{
		ClientID:     utils.GGClientID,
		ClientSecret: utils.GGClientSecret,
		Endpoint:     endpoints.Google,
		RedirectURL:  utils.BeBaseUrl + utils.GOOGLE_REDIRECT_URL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
	Url, err := url.Parse(OauthConfGG.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", OauthConfGG.ClientID)
	parameters.Add("scope", strings.Join(OauthConfGG.Scopes, " "))
	parameters.Add("redirect_uri", OauthConfGG.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", OauthStateString)
	Url.RawQuery = parameters.Encode()
	urlRes := Url.String()
	return urlRes
}

//Tao thang usersocial
func RedirectToRegisterPage(input models.User) string {
	//Url, err := url.Parse(utils.REGISTER_URL)
	//if err != nil {
	//	log.Fatal("Parse: ", err)
	//}
	//parameters := url.Values{}
	//parameters.Add("first_name", input.FirstName)
	//parameters.Add("last_name", input.LastName)
	//parameters.Add("email", input.Email)
	//parameters.Add("avatar", input.Avatar)
	//parameters.Add("id_social", input.IDSocial)
	//parameters.Add("type_social", input.TypeSocial)
	//Url.RawQuery = parameters.Encode()
	//urlRes := Url.String()
	a := strconv.Itoa(int(input.ID))
	return utils.FeBaseUrl + utils.REGISTER_URL + "?email_social=" + input.Email + "&id_social=" + a
}

func CheckUserSocialAndCreateToken(input models.User) (isExist bool, token *models.TokenDetails, err error) {
	u, err := models.GetCustomerSocialFromID(input.ID)
	if err != nil && err.Error() != "record not found" {
		return
	}
	if (u == models.User{}) {
		return false, nil, nil
	} else {
		token, err := models.CreateToken(u.ID)
		return true, token, err
	}
}

//User is a retrieved and authentiacted user.
// type UserApple struct {
// 	Name struct {
// 		FirstName  string `json:"firstName"`
// 		MiddleName string `json:"middleName"`
// 		LastName   string `json:"lastName"`
// 	} `json:"name"`
// 	Email string `json:"email"`
// }

func ProcessCallBackFacebook(state string, code string) (res ResSocialLogin, err error) {
	if state != OauthStateString {
		return res, fmt.Errorf("invalid oauth state, expected '%s', got '%s'\n", OauthStateString, state)
	}
	token, err := OauthConfFB.Exchange(oauth2.NoContext, code)
	if err != nil {
		return res, fmt.Errorf("oauthConf.Exchange() failed with '%s'\n", err)
	}
	resp, err := http.Get("https://graph.facebook.com/me?fields=email,first_name,last_name,id,picture&access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		return res, fmt.Errorf("Get: %s\n", err)
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("ReadAll: %s\n", err)
	}
	err = json.Unmarshal(response, &res)
	return
}

// User is a retrieved and authentiacted user.
type UserGoogle struct {
	Sub           uint64 `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

func ProcessCallBackGoogle(state string, code string) (res UserGoogle, err error) {
	if state != OauthStateString {
		return res, fmt.Errorf("invalid oauth state, expected '%s', got '%s'\n", OauthStateString, state)
	}
	// Handle the exchange code to initiate a transport.
	tok, err := OauthConfGG.Exchange(oauth2.NoContext, code)
	//if err != nil {
	//	return res, err
	//}
	// Construct the client.
	client := OauthConfGG.Client(oauth2.NoContext, tok)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("ReadAll: %s\n", err)
	}
	_ = json.Unmarshal(data, &res)
	return
}
*/
