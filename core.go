package vrcapi

import (
	"net/http"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"net/url"
	"strings"

	cookiejar "github.com/juju/persistent-cookiejar"
)

type AuthCodeFuncVar func() string

type VRChatAPI struct {
	username       string
	password       string
	useragent      string
	Jar            *cookiejar.Jar
	Client         *http.Client
	LoggedIn       bool
	Url            string
	ShouldAutoAuth bool
	AuthCodeFunc   AuthCodeFuncVar
}

func New(username, password, useragent string) *VRChatAPI {

	jar, _ := cookiejar.New(nil)

	jar.Save()
	return &VRChatAPI{
		username:  username,
		password:  password,
		useragent: useragent,
		Jar:       jar,
		Client:    &http.Client{Jar: jar},
		Url:       "https://api.vrchat.cloud/api/1/",
		AuthCodeFunc: func() string {
			fmt.Println("Enter auth code: ")

			var code string

			fmt.Scanln(&code)

			return code
		},
	}
}

func (api *VRChatAPI) SetAuthCodeFunc(f AuthCodeFuncVar) {
	api.ShouldAutoAuth = true
	api.AuthCodeFunc = f
}

func (api *VRChatAPI) sendAuthRequest(method, endpoint, body string) (*http.Response, error) {

	res, err := http.NewRequest(method, api.Url+endpoint, nil)

	if err != nil {
		return nil, err
	}

	userName := url.QueryEscape(api.username)
	passWord := url.QueryEscape(api.password)

	res.Header.Add("User-Agent", api.useragent)
	res.Header.Add("Accept", "*/*")
	res.Header.Add("Content-Type", "application/json")
	res.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(userName+":"+passWord)))
	res.Header.Add("auth", "JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26")

	res.Body = io.NopCloser(strings.NewReader(body))

	resp, err := api.Client.Do(res)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (api *VRChatAPI) SendRequest(method, endpoint, body string) (*http.Response, error) {

	fmt.Println(api.Url + endpoint)

	res, err := http.NewRequest(method, api.Url+endpoint, nil)

	if err != nil {
		return nil, err
	}

	res.Header.Add("User-Agent", api.useragent)
	res.Header.Add("Accept", "*/*")
	res.Header.Add("Content-Type", "application/json")
	res.Header.Add("auth", "JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26")

	res.Body = io.NopCloser(strings.NewReader(body))

	resp, err := api.Client.Do(res)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		err = api.Login()
		if err != nil {
			return nil, err
		}
		return api.SendRequest(method, endpoint, body)
	}

	if resp.StatusCode == 429 {
		fmt.Println("Rate limited sleeping for 60 minutes until rate limit is reset")
		time.Sleep(60 * time.Minute)
		return api.SendRequest(method, endpoint, body)
	}

	return resp, nil
}

func (api *VRChatAPI) HandleResponeError(resp *http.Response) error {
	apiErr := APIError{}

	err := json.NewDecoder(resp.Body).Decode(&apiErr)

	if err != nil {
		return err
	}

	return fmt.Errorf(apiErr.Error.Message + " " + fmt.Sprintf("%d", apiErr.Error.StatusCode) + " " + resp.Status)
}
