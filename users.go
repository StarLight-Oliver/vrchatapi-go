package vrcapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

/*
	TODO API Endpoints
	_SearchAllUsers
	GetUserById
	_UpdateUserInfo
	_GetUserGroups
	GetUserGroupRequests
	GetUserCurrentRepresentedGroup
*/

func (api *VRChatAPI) SearchAllUsers(displayName string, max, offset int, findDeveloper bool) ([]User, error) {

	if !api.LoggedIn {
		return nil, fmt.Errorf("not logged in")
	}

	if max > 100 {
		return nil, fmt.Errorf("max must be less than or equal to 100")
	}

	queryParams := url.Values{
		"search": {displayName},
		"n":      {fmt.Sprintf("%d", max)},
		"offset": {fmt.Sprintf("%d", offset)},
	}

	if findDeveloper {
		queryParams.Add("developerType", "internal")
	}

	resp, err := api.SendRequest("GET", "users?"+queryParams.Encode(), "")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, api.HandleResponeError(resp)
	}

	friends := []User{}

	err = json.NewDecoder(resp.Body).Decode(&friends)

	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (api *VRChatAPI) UpdateUserInfo(userID, jsonData string) (User, error) {

	if !api.LoggedIn {
		return User{}, fmt.Errorf("not logged in")
	}

	resp, err := api.SendRequest("PUT", "users/"+userID, jsonData)
	if err != nil {
		return User{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return User{}, api.HandleResponeError(resp)
	}

	user := User{}

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (api *VRChatAPI) GetUserGroups(userId string) ([]Group, error) {
	if !api.LoggedIn {
		return nil, fmt.Errorf("not logged in")
	}

	resp, err := api.SendRequest("GET", "users/"+userId+"/groups", "")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, api.HandleResponeError(resp)
	}

	defer resp.Body.Close()

	groups := []Group{}

	err = json.NewDecoder(resp.Body).Decode(&groups)

	if err != nil {
		return nil, err
	}

	return groups, nil
}
