package vrcapi

import (
	"encoding/json"
	"fmt"
	"io"
)

/*
TODO API Endpoints
exists
verify2fa code with recovery code
verify auth token
delete user
*/

func (api *VRChatAPI) Login() error {

	resp, err := api.sendAuthRequest("GET", "auth/user", "")

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	user := User{}

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return err
	}

	if user.Id == "" {
		// this needs to be refactored
		fmt.Println("User not logged in")

		// // wait for user to input an auth code
		// fmt.Println("Enter auth code: ")

		// var code string

		// fmt.Scanln(&code)

		if !api.ShouldAutoAuth {
			return fmt.Errorf("need to auth manually")
		}

		code := api.AuthCodeFunc()

		res, err := api.sendAuthRequest("POST", "auth/twofactorauth/totp/verify", fmt.Sprintf(`{"code": "%s"}`, code))

		if err != nil {
			return err
		}

		defer res.Body.Close()

		body2, err := io.ReadAll(res.Body)

		fmt.Println(string(body2))
		// err = json.NewDecoder(res.Body).Decode(&auth)

		if err != nil {
			return err
		}

		// fmt.Println(auth)

		// if !auth.Verified {
		// 	return fmt.Errorf("auth code failed")
		// }
	}

	// save cookies
	api.Jar.Save()

	api.LoggedIn = true

	return nil
}

func (api *VRChatAPI) GetCurrentUser() (User, error) {

	if !api.LoggedIn {
		return User{}, fmt.Errorf("not logged in")
	}

	resp, err := api.SendRequest("GET", "auth/user", "")

	if err != nil {
		return User{}, err
	}

	defer resp.Body.Close()

	user := User{}

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

type Verify2FAResult struct {
	Verified bool `json:"verified"`
}

func (api *VRChatAPI) verify2FA(twoFAType, code string) (bool, error) {

	url := fmt.Sprintf("auth/twofactorauth/%s/verify", twoFAType)

	res, err := api.sendAuthRequest("POST", url, fmt.Sprintf(`{"code": "%s"}`, code))

	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return false, api.HandleResponeError(res)
	}

	var verified *Verify2FAResult

	err = json.NewDecoder(res.Body).Decode(&verified)

	if err != nil {
		return false, err
	}

	return verified.Verified, nil
}

func (api *VRChatAPI) Verify2FACode(code string) (bool, error) {

	return api.verify2FA("totp", code)
}

func (api *VRChatAPI) Verify2FARecoveryCode(code string) (bool, error) {
	// res, err := api.sendAuthRequest("POST", "auth/twofactorauth/otp/verify", fmt.Sprintf(`{"code": "%s"}`, code))

	// if err != nil {
	// 	return false, err
	// }

	// defer res.Body.Close()

	// if res.StatusCode != 200 {
	// 	return false, api.HandleResponeError(res)
	// }

	// var verified *Verify2FAResult

	// err = json.NewDecoder(res.Body).Decode(&verified)

	// if err != nil {
	// 	return false, err
	// }

	// return verified.Verified, nil

	return api.verify2FA("otp", code)
}

func (api *VRChatAPI) Verify2FAEmail(code string) (bool, error) {
	// res, err := api.sendAuthRequest("POST", "auth/twofactorauth/emailotp/verify", fmt.Sprintf(`{"code": "%s"}`, code))

	// if err != nil {
	// 	return false, err
	// }

	// defer res.Body.Close()

	// if res.StatusCode != 200 {
	// 	return false, api.HandleResponeError(res)
	// }

	// var verified *Verify2FAResult

	// err = json.NewDecoder(res.Body).Decode(&verified)

	// if err != nil {
	// 	return false, err
	// }

	// return verified.Verified, nil

	return api.verify2FA("emailotp", code)
}

type VerifyAuthTokenResult struct {
	Ok    bool   `json:"ok"`
	Token string `json:"token"`
}

func (api *VRChatAPI) VerifyAuthToken() (*VerifyAuthTokenResult, error) {
	resp, err := api.sendAuthRequest("POST", "auth/verifyauthtoken", "")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, api.HandleResponeError(resp)
	}

	var verified *VerifyAuthTokenResult

	err = json.NewDecoder(resp.Body).Decode(&verified)

	if err != nil {
		return nil, err
	}

	return verified, nil
}

func (api *VRChatAPI) Logout() error {

	resp, err := api.SendRequest("PUT", "logout", "")

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errorStruct := APIError{}

		err = json.NewDecoder(resp.Body).Decode(&errorStruct)

		if err != nil {
			// something went REALLY wrong
			return err
		}

		return fmt.Errorf("error: %s", errorStruct.Error.Message)
	}

	logout := SuccessResponse{}

	err = json.NewDecoder(resp.Body).Decode(&logout)

	if err != nil {
		return err
	}

	if logout.Success.StatusCode != 200 {
		// can this even happen?
		return fmt.Errorf("error: %s", logout.Success.Message)
	}

	api.LoggedIn = false

	return nil
}
