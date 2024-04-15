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
