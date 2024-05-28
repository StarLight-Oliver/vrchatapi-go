package vrcapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type SearchAvatarRequest struct {
	IsFeatured        bool   `json:"featured"`
	Sort              string `json:"sort,omitempty"`
	User              string `json:"user"`
	AvatarCount       int    `json:"n"`
	OrderBy           string `json:"order"`
	Offset            int    `json:"offset"`
	TagsString        string `json:"tags"`
	TagsExcludeString string `json:"tagsExclude"`
	ReleaseStatus     string `json:"releaseStatus"`
	Platform          string `json:"platform"`
	MaxUnityVersion   string `json:"maxUnityVersion,omitempty"`
	MinUnityVersion   string `json:"minUnityVersion,omitempty"`
}

type SearchAvatarsInput struct {
	IsFeatured      bool
	Tags            []string
	TagsExclude     []string
	IsAscending     bool
	MaxUnityVersion string
	MinUnityVersion string
	AvatarCount     int
	Offset          int
	IsPageOffset    bool
	Sort            string

	ReleaseStatus string
	Platform      string
}

func (api *VRChatAPI) SearchAvatars(searchInfo SearchAvatarsInput) ([]Avatar, error) {

	requestStruct := SearchAvatarRequest{}

	queryParams := url.Values{}

	if searchInfo.IsFeatured {
		queryParams.Add("featured", "true")
	} else {
		queryParams.Add("user", "me")
	}

	if searchInfo.Tags != nil {
		requestStruct.TagsString = strings.Join(searchInfo.Tags, ",")
		queryParams.Add("tags", requestStruct.TagsString)
	}

	if searchInfo.TagsExclude != nil {
		requestStruct.TagsExcludeString = strings.Join(searchInfo.TagsExclude, ",")
		queryParams.Add("notag", requestStruct.TagsExcludeString)
	}

	if searchInfo.IsAscending {
		requestStruct.OrderBy = "ascending"
		queryParams.Add("order", "ascending")
	} else {
		requestStruct.OrderBy = "descending"
		queryParams.Add("order", "descending")
	}

	if searchInfo.AvatarCount > 0 {
		requestStruct.AvatarCount = searchInfo.AvatarCount
		requestStruct.Offset = searchInfo.Offset
		if searchInfo.IsPageOffset {
			requestStruct.Offset = searchInfo.Offset * searchInfo.AvatarCount
		}
		queryParams.Add("n", fmt.Sprintf("%d", searchInfo.AvatarCount))
		queryParams.Add("offset", fmt.Sprintf("%d", requestStruct.Offset))
	}

	if searchInfo.Sort != "" {
		queryParams.Add("sort", searchInfo.Sort)
	}

	if searchInfo.ReleaseStatus != "" {
		queryParams.Add("releaseStatus", searchInfo.ReleaseStatus)
	}

	if searchInfo.Platform != "" {
		queryParams.Add("platform", searchInfo.Platform)
	}

	if searchInfo.MaxUnityVersion != "" {
		queryParams.Add("maxUnityVersion", searchInfo.MaxUnityVersion)
	}

	if searchInfo.MinUnityVersion != "" {
		queryParams.Add("minUnityVersion", searchInfo.MinUnityVersion)
	}

	resp, err := api.SendRequest("GET", "avatars?"+queryParams.Encode(), "")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, api.HandleResponeError(resp)
	}

	avatars := []Avatar{}

	err = json.NewDecoder(resp.Body).Decode(&avatars)

	if err != nil {
		return nil, err
	}

	return avatars, nil
}
