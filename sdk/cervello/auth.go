package cervello

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Nerzal/gocloak/v11"
)

// Token User access_token model
type tokenType struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

var token tokenType

type TokenClaim struct {
	Acr               string   `json:"acr"`
	AllowedOrigins    []string `json:"allowed-origins"`
	Aud               string   `json:"aud"`
	AuthTime          int      `json:"auth_time"`
	Azp               string   `json:"azp"`
	Email             string   `json:"email"`
	EmailVerified     bool     `json:"email_verified"`
	Exp               int      `json:"exp"`
	FamilyName        string   `json:"family_name"`
	GivenName         string   `json:"given_name"`
	Iat               int      `json:"iat"`
	Iss               string   `json:"iss"`
	Jti               string   `json:"jti"`
	Name              string   `json:"name"`
	Nbf               int      `json:"nbf"`
	Nonce             string   `json:"nonce"`
	PreferredUsername string   `json:"preferred_username"`
	RealmAccess       struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
	ResourceAccess struct {
		Account struct {
			Roles []string `json:"roles"`
		} `json:"account"`
	} `json:"resource_access"`
	Scope        string `json:"scope"`
	SessionState string `json:"session_state"`
	Sub          string `json:"sub"`
	Typ          string `json:"typ"`
}

// UserLogin used to gain user access_token
func Login() {
	apiURL := envAuthURI
	resource := "/auth/realms/" + envAuthREALM + "/protocol/openid-connect/token"
	data := url.Values{}
	data.Set("client_id", envAuthClientID)
	data.Set("grant_type", envAuthGrantType)
	data.Set("username", envAuthUsername)
	data.Set("password", envAuthPassword)

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = resource
	urlStr := u.String() // "https://api.com/user/"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	r, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		internalLog("panic", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		internalLog("panic", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		internalLog("panic", err)
	}

	if resp.StatusCode != 200 {
		internalLog("panic", errors.New("error authenticate: "+string(body)))
	}

	err = json.Unmarshal(body, &token)
	if err != nil {
		internalLog("panic", err)
	}

	internalLog("info", "Logged in to cervello successfully")

	// Refresh token before it expires by 10 seconds
	time.AfterFunc(time.Second*time.Duration(token.ExpiresIn-10), Login)
}

func GetCervelloToken() string {
	return token.AccessToken
}

type User struct {
	Token      string
	TokenClaim TokenClaim
}

func ValidateToken(token string) (*User, error) {
	client := gocloak.NewClient(keycloakHost)
	restyClient := client.RestyClient()
restyClient.SetDebug(true)
restyClient.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
	ctx := context.Background()
	jwt, _, err := client.DecodeAccessToken(ctx, token, keycloakRealm)
	if err != nil {
		return nil, err
	}
	if !jwt.Valid {
		return nil, errors.New("invalid token")
	}

	claim, err := json.Marshal(jwt.Claims)
	if err != nil {
		return nil, errors.New("serializing token error")
	}
	var parsedClaim TokenClaim
	if err = json.Unmarshal(claim, &parsedClaim); err != nil {
		return nil, errors.New("parsing token error")
	}

	return &User{
		Token:      token,
		TokenClaim: parsedClaim,
	}, nil

}

func ValidateBarearToken(bearerToken string) (*User, error) {
	r, err := regexp.Compile("Bearer\\s.*")
	if err != nil {
		return nil, errors.New("invalidRegex")
	}
	if !r.MatchString(bearerToken) {
		return nil, errors.New("noToken")
	}

	s := strings.Split(bearerToken, " ")
	token := s[1]

	// 2- check if the token is valid using keycloak and parse the token to construct user object
	user, err := ValidateToken(token)
	if err != nil {
		return nil, errors.New("invalidToken")
	}

	return user, nil
}

// ValidateRoles ...
func (user *User) ValidateRoles(roles ...string) bool {
	userRolesMap := make(map[string]*User)
	for _, userRole := range user.TokenClaim.RealmAccess.Roles {
		userRolesMap[userRole] = user
	}

	for _, role := range roles {
		_, ok := userRolesMap[role]
		if !ok {
			return false
		}
	}
	return true
}
