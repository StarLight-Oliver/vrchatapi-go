package vrcapi

import (
	"encoding/json"
	"fmt"
)

type CreateInstanceRequest struct {
	WorldId          string   `json:"worldId"`
	WorldType        string   `json:"type"`
	Region           string   `json:"region"`
	OwnerId          string   `json:"ownerId"`
	RoleIds          []string `json:"roleIds"`
	GroupAccessType  string   `json:"groupAccessType"`
	QueueEnabled     bool     `json:"queueEnabled"`
	ClosedAt         string   `json:"closedAt"`
	CanRequestInvite bool     `json:"canRequestInvite"`
	HardClose        bool     `json:"hardClose"`
	InviteOnly       bool     `json:"inviteOnly"`
}

type Instance struct {
	Active              bool           `json:"active"`
	CanRequestInvite    bool           `json:"canRequestInvite"`
	Capacity            int            `json:"capacity"`
	Full                bool           `json:"full"`
	Id                  string         `json:"id"`
	InstanceId          string         `json:"instanceId"`
	Location            string         `json:"location"`
	N_Users             int            `json:"n_users"`
	Name                string         `json:"name"`
	OwnerId             string         `json:"ownerId,omitempty"`
	Permanent           bool           `json:"permanent"`
	PhotonRegion        string         `json:"photonRegion"`
	Platforms           map[string]int `json:"platforms"`
	Region              string         `json:"region"`
	SecureName          string         `json:"secureName"`
	ShortName           string         `json:"shortName,omitempty"`
	Tags                []string       `json:"tags"`
	Type                string         `json:"type"`
	WorldID             string         `json:"worldId"`
	Hidden              string         `json:"hidden,omitempty"`
	Friends             string         `json:"friends,omitempty"`
	Private             string         `json:"private,omitempty"`
	QueueEnabled        bool           `json:"queueEnabled"`
	QueueSize           int            `json:"queueSize"`
	RecommendedCapacity int            `json:"recommendedCapacity"`
	RoleRestricted      bool           `json:"roleRestricted"`
	Strict              bool           `json:"strict"`
	UserCount           int            `json:"userCount"`
	// World               World `json:"world"`
	Users             []User `json:"users"`
	GroupAccessType   string `json:"groupAccessType"`
	HasCapacityForYou bool   `json:"hasCapacityForYou"`
	Nonce             string `json:"nonce"`
	CloseAt           string `json:"closeAt"`
	HardClose         bool   `json:"hardClose"`
}

type InstanceShortNameResponse struct {
	ShortName  string `json:"shortName"`
	SecureName string `json:"secureName"`
}

func (c *VRChatAPI) CreateInstance(req *CreateInstanceRequest) (*Instance, error) {
	var instance *Instance

	if !c.LoggedIn {
		return nil, fmt.Errorf("not logged in")
	}

	jsonData, err := json.Marshal(req)

	if err != nil {
		return nil, err
	}

	resp, err := c.SendRequest("POST", "/instances", string(jsonData))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, c.HandleResponeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&instance)

	if err != nil {
		return nil, err
	}
	return instance, err
}

func (c *VRChatAPI) GetInstance(worldId string, instanceId string) (*Instance, error) {
	var instance *Instance

	if !c.LoggedIn {
		return nil, fmt.Errorf("not logged in")
	}

	url := fmt.Sprintf("/instances/%s:%s", worldId, instanceId)

	resp, err := c.SendRequest("GET", url, "")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, c.HandleResponeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&instance)

	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (c *VRChatAPI) CloseInstance(worldId string, instanceId string) (*Instance, error) {
	var instance *Instance

	if !c.LoggedIn {
		return nil, fmt.Errorf("not logged in")
	}

	url := fmt.Sprintf("/instances/%s:%s", worldId, instanceId)

	resp, err := c.SendRequest("DELETE", url, "")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, c.HandleResponeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&instance)

	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (c *VRChatAPI) GetInstanceShortName(worldId string, instanceId string) (*InstanceShortNameResponse, error) {
	var instance *InstanceShortNameResponse

	if !c.LoggedIn {
		return nil, fmt.Errorf("not logged in")
	}

	url := fmt.Sprintf("/instances/%s:%s/shortname", worldId, instanceId)

	resp, err := c.SendRequest("GET", url, "")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, c.HandleResponeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&instance)

	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (c *VRChatAPI) SendSelfInvite(worldId string, instanceId string) (*SuccessResponse, error) {
	var instance *SuccessResponse

	if !c.LoggedIn {
		return nil, fmt.Errorf("not logged in")
	}

	url := fmt.Sprintf("/instances/%s:%s/invite", worldId, instanceId)

	resp, err := c.SendRequest("POST", url, "")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, c.HandleResponeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&instance)

	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (c *VRChatAPI) GetInstanceByShortName(shortName string) (*Instance, error) {
	var instance *Instance

	if !c.LoggedIn {
		return nil, fmt.Errorf("not logged in")
	}

	url := fmt.Sprintf("/instances/s/%s", shortName)

	resp, err := c.SendRequest("GET", url, "")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, c.HandleResponeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&instance)

	if err != nil {
		return nil, err
	}
	return instance, nil
}
