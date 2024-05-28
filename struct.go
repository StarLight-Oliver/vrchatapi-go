package vrcapi

type AuthTotpVerify struct {
	// This seems to be unused
	Verified bool `json:"verified"`
}

type User struct {
	Id                             string   `json:"id"`
	DisplayName                    string   `json:"displayName"`
	UserIcon                       string   `json:"userIcon"`
	Bio                            string   `json:"bio"`
	BioLinks                       []string `json:"bioLinks"`
	ProfilePicOverride             string   `json:"profilePicOverride"`
	StatusDescription              string   `json:"statusDescription"`
	Username                       string   `json:"username"`
	PastDisplayNames               []string `json:"pastDisplayNames"`
	HasEmail                       bool     `json:"hasEmail"`
	HasPendingEmail                bool     `json:"hasPendingEmail"`
	ObfuscatedEmail                string   `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail         string   `json:"obfuscatedPendingEmail"`
	EmailVerified                  bool     `json:"emailVerified"`
	HasBirthday                    bool     `json:"hasBirthday"`
	Unsubscribe                    bool     `json:"unsubscribe"`
	StatusHistory                  []string `json:"statusHistory"`
	StatusFirstTime                bool     `json:"statusFirstTime"`
	Friends                        []string `json:"friends"`
	FriendGroupNames               []string `json:"friendGroupNames"`
	CurrentAvatarImageUrl          string   `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageUrl string   `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatar                  string   `json:"currentAvatar"`
	CurrentAvatarAssetUrl          string   `json:"currentAvatarAssetUrl"`
	FallbackAvatar                 string   `json:"fallbackAvatar"`
	AccountDeletionDate            string   `json:"accountDeletionDate"`
	AccountDeletionLog             string   `json:"accountDeletionLog"`
	AcceptedTOSVersion             int      `json:"acceptedTOSVersion"`
	AcceptedPrivacyVersion         int      `json:"acceptedPrivacyVersion"`
	SteamId                        string   `json:"steamId"`
	HassLoggedInFromClient         bool     `json:"hasLoggedInFromClient"`
	HomeLocation                   string   `json:"homeLocation"`
	TwoFactorAuthEnabled           bool     `json:"twoFactorAuthEnabled"`
	TwoFactorAuthEnabledDate       string   `json:"twoFactorAuthEnabledDate"`
	UpdatedAt                      string   `json:"updated_at"`
	State                          string   `json:"state"`
	Tags                           []string `json:"tags"`
	DeveloperType                  string   `json:"developerType"`
	LastLogin                      string   `json:"last_login"`
	LastPlatform                   string   `json:"last_platform"`
	AllowAvatarCopying             bool     `json:"allowAvatarCopying"`
	Status                         string   `json:"status"`
	DateJoined                     string   `json:"date_joined"`
	IsFriend                       bool     `json:"isFriend"`
	FriendKey                      string   `json:"friendKey"`
	LastActivity                   string   `json:"last_activity"`
	OnlineFriends                  []string `json:"onlineFriends"`
	ActiveFriends                  []string `json:"activeFriends"`
	// Presence Presence `json:"presence"`
	OfflineFriends []string `json:"offlineFriends"`
}

type Group struct {
	Id          string `json:"groupId"`
	Name        string `json:"name"`
	OwnerId     string `json:"ownerId"`
	MemberCount int    `json:"memberCount"`
}

type APIError struct {
	Error MsgAndStatusCode `json:"error"`
}

type SuccessResponse struct {
	Success MsgAndStatusCode `json:"success"`
}

type MsgAndStatusCode struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

/*
"assetUrl": "A",
"assetUrlObject": { },
"authorId": "usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469",
"authorName": "A",
"created_at": "1970-01-01T00:00:00.000Z",
"description": "string",
"featured": false,
"id": "avtr_912d66a4-4714-43b8-8407-7de2cafbf55b",
"imageUrl": "A",
"name": "A",
"releaseStatus": "public",
"tags": [
"A"
],
"thumbnailImageUrl": "A",
"unityPackageUrl": "string",
"unityPackages": [
{
"assetUrl": "https://vrchat.com/api/1/file/file_cd0caa7b-69ba-4715-8dfe-7d667a9d2537/65/file",
"assetUrlObject": { },
"assetVersion": 4,
"created_at": "2020-09-10T06:13:27.777Z",
"id": "unp_52b12c39-4163-457d-a4a9-630e7aff1bff",
"platform": "standalonewindows",
"pluginUrl": "",
"pluginUrlObject": { },
"unitySortNumber": 20180414000,
"unityVersion": "2022.3.6f1"
}],
"updated_at": "1970-01-01T00:00:00.000Z",
"version": 68
}
*/

type Avatar struct {
	AssetUrl          string                 `json:"assetUrl"`
	AssetUrlObject    map[string]interface{} `json:"assetUrlObject"`
	AuthorId          string                 `json:"authorId"`
	AuthorName        string                 `json:"authorName"`
	CreatedAt         string                 `json:"created_at"`
	Description       string                 `json:"description"`
	Featured          bool                   `json:"featured"`
	Id                string                 `json:"id"`
	ImageUrl          string                 `json:"imageUrl"`
	Name              string                 `json:"name"`
	ReleaseStatus     string                 `json:"releaseStatus"`
	Tags              []string               `json:"tags"`
	ThumbnailImageUrl string                 `json:"thumbnailImageUrl"`
	UnityPackageUrl   string                 `json:"unityPackageUrl"`
	UnityPackages     []UnityPackage         `json:"unityPackages"`
	UpdatedAt         string                 `json:"updated_at"`
	Version           int                    `json:"version"`
}

type UnityPackage struct {
	AssetUrl        string                 `json:"assetUrl"`
	AssetUrlObject  map[string]interface{} `json:"assetUrlObject"`
	AssetVersion    int                    `json:"assetVersion"`
	CreatedAt       string                 `json:"created_at"`
	Id              string                 `json:"id"`
	Platform        string                 `json:"platform"`
	PluginUrl       string                 `json:"pluginUrl"`
	PluginUrlObject map[string]interface{} `json:"pluginUrlObject"`
	UnitySortNumber int                    `json:"unitySortNumber"`
	UnityVersion    string                 `json:"unityVersion"`
}
