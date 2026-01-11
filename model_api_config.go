/*
VRChat API Documentation

![VRChat API Banner](https://vrchatapi.github.io/assets/img/api_banner_1500x400.png)  # Welcome to the VRChat API  Before we begin, we would like to state this is a **COMMUNITY DRIVEN PROJECT**. This means that everything you read on here was written by the community itself and is **not** officially supported by VRChat. The documentation is provided \"AS IS\", and any action you take towards VRChat is completely your own responsibility.  The documentation and additional libraries SHALL ONLY be used for applications interacting with VRChat's API in accordance with their [Terms of Service](https://hello.vrchat.com/legal) and [Community Guidelines](https://hello.vrchat.com/community-guidelines), and MUST NOT be used for modifying the client, \"avatar ripping\", or other illegal activities. Malicious usage or spamming the API may result in account termination. Certain parts of the API are also more sensitive than others, for example moderation, so please tread extra carefully and read the warnings when present.  ![Tupper Policy on API](https://i.imgur.com/yLlW7Ok.png)  Finally, use of the API using applications other than the approved methods (website, VRChat application, Unity SDK) is not officially supported. VRChat provides no guarantee or support for external applications using the API. Access to API endpoints may break **at any time, without notice**. Therefore, please **do not ping** VRChat Staff in the VRChat Discord if you are having API problems, as they do not provide API support. We will make a best effort in keeping this documentation and associated language libraries up to date, but things might be outdated or missing. If you find that something is no longer valid, please contact us on Discord or [create an issue](https://github.com/vrchatapi/specification/issues) and tell us so we can fix it.  # Getting Started  The VRChat API can be used to programmatically retrieve or update information regarding your profile, friends, avatars, worlds and more. The API consists of two parts, \"Photon\" which is only used in-game, and the \"Web API\" which is used by both the game and the website. This documentation focuses only on the Web API.  The API is designed around the REST ideology, providing semi-simple and usually predictable URIs to access and modify objects. Requests support standard HTTP methods like GET, PUT, POST, and DELETE and standard status codes. Response bodies are always UTF-8 encoded JSON objects, unless explicitly documented otherwise.  <div class=\"callout callout-error\">   <strong>🛑 Warning! Do not touch Photon!</strong><br>   Photon is only used by the in-game client and should <b>not</b> be touched. Doing so may result in permanent account termination. </div>  <div class=\"callout callout-info\">   <strong>ℹ️ Authentication</strong><br>   Read <a href=\"#tag--authentication\">Authentication</a> for how to log in. </div>  # Using the API  For simply exploring what the API can do it is strongly recommended to download [Insomnia](https://insomnia.rest/download), a free and open-source API client that's great for sending requests to the API in an orderly fashion. Insomnia allows you to send data in the format that's required for VRChat's API. It is also possible to try out the API in your browser, by first logging in at [vrchat.com/home](https://vrchat.com/home/) and then going to [vrchat.com/api/1/auth/user](https://vrchat.com/api/1/auth/user), but the information will be much harder to work with.  For more permanent operation such as software development it is instead recommended to use one of the existing language SDKs. This community project maintains API libraries in several languages, which allows you to interact with the API with simple function calls rather than having to implement the HTTP protocol yourself. Most of these libraries are automatically generated from the API specification, sometimes with additional helpful wrapper code to make usage easier. This allows them to be almost automatically updated and expanded upon as soon as a new feature is introduced in the specification itself. The libraries can be found on [GitHub](https://github.com/vrchatapi) or following:  * [NodeJS (JavaScript)](https://www.npmjs.com/package/vrchat) * [Dart](https://pub.dev/packages/vrchat_dart) * [Rust](https://crates.io/crates/vrchatapi) * [C#](https://github.com/vrchatapi/vrchatapi-csharp) * [Python](https://github.com/vrchatapi/vrchatapi-python)  # Pagination  Most endpoints enforce pagination, meaning they will only return 10 entries by default, and never more than 100.<br> Using both the limit and offset parameters allows you to easily paginate through a large number of objects.  | Query Parameter | Type | Description | | ----------|--|------- | | `n` | integer  | The number of objects to return. This value often defaults to 10. Highest limit is always 100.| | `offset` | integer  | A zero-based offset from the default object sorting.|  If a request returns fewer objects than the `limit` parameter, there are no more items available to return.  # Contribution  Do you want to get involved in the documentation effort? Do you want to help improve one of the language API libraries? This project is an [OPEN Open Source Project](https://openopensource.org)! This means that individuals making significant and valuable contributions are given commit-access to the project. It also means we are very open and welcoming of new people making contributions, unlike some more guarded open-source projects.  [![Discord](https://img.shields.io/static/v1?label=vrchatapi&message=discord&color=blueviolet&style=for-the-badge)](https://discord.gg/qjZE9C9fkB)

API version: 1.20.7
Contact: vrchatapi.lpv0t@aries.fyi
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the APIConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIConfig{}

// APIConfig Global configuration for various features.
type APIConfig struct {
	// The current platform-wide event taking place
	CampaignStatus string `json:"CampaignStatus"`
	// Toggles if certain assets are preloaded in the background
	DisableBackgroundPreloads bool `json:"DisableBackgroundPreloads"`
	// Toggles whether users without a current VRC+ subscription are priority recipients for gift drops
	LocationGiftingNonSubPrioEnabled bool `json:"LocationGiftingNonSubPrioEnabled"`
	// Unknown, probably voice optimization testing
	VoiceEnableDegradation bool `json:"VoiceEnableDegradation"`
	// Unknown, probably voice optimization testing
	VoiceEnableReceiverLimiting bool `json:"VoiceEnableReceiverLimiting"`
	AccessLogsUrls APIConfigAccessLogsUrls `json:"accessLogsUrls"`
	// VRChat's office address
	Address string `json:"address"`
	AgeVerificationInviteVisible bool `json:"ageVerificationInviteVisible"`
	AgeVerificationP bool `json:"ageVerificationP"`
	AgeVerificationStatusVisible bool `json:"ageVerificationStatusVisible"`
	// Max retries for avatar analysis requests
	AnalysisMaxRetries int32 `json:"analysisMaxRetries"`
	// Interval between retries for avatar analysis requests
	AnalysisRetryInterval int32 `json:"analysisRetryInterval"`
	// Unknown
	AnalyticsSegmentNewUIPctOfUsers int32 `json:"analyticsSegment_NewUI_PctOfUsers"`
	// Unknown
	AnalyticsSegmentNewUISalt string `json:"analyticsSegment_NewUI_Salt"`
	// Public Announcements
	Announcements []APIConfigAnnouncement `json:"announcements"`
	AudioConfig *APIConfigAudioConfig `json:"audioConfig,omitempty"`
	// List of supported Languages
	AvailableLanguageCodes []string `json:"availableLanguageCodes"`
	// List of supported Languages
	AvailableLanguages []string `json:"availableLanguages"`
	AvatarPerfLimiter APIConfigAvatarPerfLimiter `json:"avatarPerfLimiter"`
	// Unknown
	ChatboxLogBufferSeconds int32 `json:"chatboxLogBufferSeconds"`
	// apiKey to be used for all other requests
	ClientApiKey string `json:"clientApiKey"`
	// Unknown
	ClientBPSCeiling int32 `json:"clientBPSCeiling"`
	// Unknown
	ClientDisconnectTimeout int32 `json:"clientDisconnectTimeout"`
	// Unknown
	ClientNetDispatchThread *bool `json:"clientNetDispatchThread,omitempty"`
	// Unknown
	ClientNetDispatchThreadMobile bool `json:"clientNetDispatchThreadMobile"`
	// Unknown
	ClientNetInThread *bool `json:"clientNetInThread,omitempty"`
	// Unknown
	ClientNetInThread2 *bool `json:"clientNetInThread2,omitempty"`
	// Unknown
	ClientNetInThreadMobile *bool `json:"clientNetInThreadMobile,omitempty"`
	// Unknown
	ClientNetInThreadMobile2 *bool `json:"clientNetInThreadMobile2,omitempty"`
	// Unknown
	ClientNetOutThread *bool `json:"clientNetOutThread,omitempty"`
	// Unknown
	ClientNetOutThread2 *bool `json:"clientNetOutThread2,omitempty"`
	// Unknown
	ClientNetOutThreadMobile *bool `json:"clientNetOutThreadMobile,omitempty"`
	// Unknown
	ClientNetOutThreadMobile2 *bool `json:"clientNetOutThreadMobile2,omitempty"`
	// Unknown
	ClientQR *int32 `json:"clientQR,omitempty"`
	// Unknown
	ClientReservedPlayerBPS int32 `json:"clientReservedPlayerBPS"`
	// Unknown
	ClientSentCountAllowance int32 `json:"clientSentCountAllowance"`
	Constants APIConfigConstants `json:"constants"`
	// VRChat's contact email
	ContactEmail string `json:"contactEmail"`
	// VRChat's copyright-issues-related email
	CopyrightEmail string `json:"copyrightEmail"`
	// VRChat's DMCA claim webform url
	CopyrightFormUrl string `json:"copyrightFormUrl"`
	// Current version number of the Privacy Agreement
	CurrentPrivacyVersion int32 `json:"currentPrivacyVersion"`
	// Current version number of the Terms of Service
	CurrentTOSVersion int32 `json:"currentTOSVersion"`
	DefaultAvatar string `json:"defaultAvatar"`
	DefaultStickerSet string `json:"defaultStickerSet"`
	// Unknown
	DevLanguageCodes []string `json:"devLanguageCodes,omitempty"`
	// Link to download the development SDK, use downloadUrls instead
	// Deprecated
	DevSdkUrl string `json:"devSdkUrl"`
	// Version of the development SDK
	// Deprecated
	DevSdkVersion string `json:"devSdkVersion"`
	// Unknown, \"dis\" maybe for disconnect?
	DisCountdown time.Time `json:"dis-countdown"`
	// Unknown
	DisableAVProInProton *bool `json:"disableAVProInProton,omitempty"`
	// Toggles if copying avatars should be disabled
	DisableAvatarCopying bool `json:"disableAvatarCopying"`
	// Toggles if avatar gating should be disabled. Avatar gating restricts uploading of avatars to people with the `system_avatar_access` Tag or `admin_avatar_access` Tag
	DisableAvatarGating bool `json:"disableAvatarGating"`
	// Unknown
	DisableCaptcha *bool `json:"disableCaptcha,omitempty"`
	// Toggles if the Community Labs should be disabled
	DisableCommunityLabs bool `json:"disableCommunityLabs"`
	// Toggles if promotion out of Community Labs should be disabled
	DisableCommunityLabsPromotion bool `json:"disableCommunityLabsPromotion"`
	// Unknown
	DisableEmail bool `json:"disableEmail"`
	// Toggles if Analytics should be disabled.
	DisableEventStream bool `json:"disableEventStream"`
	// Toggles if feedback gating should be disabled. Feedback gating restricts submission of feedback (reporting a World or User) to people with the `system_feedback_access` Tag.
	DisableFeedbackGating bool `json:"disableFeedbackGating"`
	// Unknown, probably toggles compilation of frontend web builds? So internal flag?
	DisableFrontendBuilds bool `json:"disableFrontendBuilds"`
	// Toggles if gift drops should be disabled
	DisableGiftDrops bool `json:"disableGiftDrops"`
	// Unknown
	DisableHello bool `json:"disableHello"`
	// Toggles if signing up for Subscriptions in Oculus is disabled or not.
	DisableOculusSubs bool `json:"disableOculusSubs"`
	// Toggles if new user account registration should be disabled.
	DisableRegistration bool `json:"disableRegistration"`
	// Toggles if Steam Networking should be disabled. VRChat these days uses Photon Unity Networking (PUN) instead.
	DisableSteamNetworking bool `json:"disableSteamNetworking"`
	// Toggles if 2FA should be disabled.
	// Deprecated
	DisableTwoFactorAuth bool `json:"disableTwoFactorAuth"`
	// Toggles if Udon should be universally disabled in-game.
	DisableUdon bool `json:"disableUdon"`
	// Toggles if account upgrading \"linking with Steam/Oculus\" should be disabled.
	DisableUpgradeAccount bool `json:"disableUpgradeAccount"`
	// Download link for game on the Oculus Rift website.
	DownloadLinkWindows string `json:"downloadLinkWindows"`
	DownloadUrls APIConfigDownloadURLList `json:"downloadUrls"`
	// Array of DynamicWorldRow objects, used by the game to display the list of world rows
	DynamicWorldRows []DynamicContentRow `json:"dynamicWorldRows"`
	// Unknown
	EconomyLedgerBackfill bool `json:"economyLedgerBackfill"`
	// Unknown
	EconomyLedgerMigrationStop string `json:"economyLedgerMigrationStop"`
	// Unknown
	EconomyLedgerMode string `json:"economyLedgerMode"`
	// Unknown
	EconomyPauseEnd time.Time `json:"economyPauseEnd"`
	// Unknown
	EconomyPauseStart time.Time `json:"economyPauseStart"`
	// Unknown
	EconomyPurchaseRepairEnabled bool `json:"economyPurchaseRepairEnabled"`
	// Unknown
	EconomyState int32 `json:"economyState"`
	Events APIConfigEvents `json:"events"`
	// Unknown
	ForceUseLatestWorld bool `json:"forceUseLatestWorld"`
	// Display type of gifts
	GiftDisplayType string `json:"giftDisplayType"`
	// Unknown
	GoogleApiClientId string `json:"googleApiClientId"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	HomeWorldId string `json:"homeWorldId"`
	// Redirect target if you try to open the base API domain in your browser
	HomepageRedirectTarget string `json:"homepageRedirectTarget"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	HubWorldId string `json:"hubWorldId"`
	// A list of explicitly allowed origins that worlds can request images from via the Udon's [VRCImageDownloader#DownloadImage](https://creators.vrchat.com/worlds/udon/image-loading/#downloadimage).
	ImageHostUrlList []string `json:"imageHostUrlList"`
	// Current app version for iOS
	IosAppVersion []string `json:"iosAppVersion"`
	IosVersion APIConfigIosVersion `json:"iosVersion"`
	// VRChat's job application email
	JobsEmail string `json:"jobsEmail"`
	// The maximum number of custom emoji each user may have at a given time.
	MaxUserEmoji int32 `json:"maxUserEmoji"`
	// The maximum number of custom stickers each user may have at a given time.
	MaxUserStickers int32 `json:"maxUserStickers"`
	MinSupportedClientBuildNumber APIConfigMinSupportedClientBuildNumber `json:"minSupportedClientBuildNumber"`
	// Minimum Unity version required for uploading assets
	MinimumUnityVersionForUploads string `json:"minimumUnityVersionForUploads"`
	// VRChat's moderation related email
	ModerationEmail string `json:"moderationEmail"`
	// Used in-game to notify a user they aren't allowed to select avatars in private worlds
	NotAllowedToSelectAvatarInPrivateWorldMessage string `json:"notAllowedToSelectAvatarInPrivateWorldMessage"`
	OfflineAnalysis APIConfigOfflineAnalysis `json:"offlineAnalysis"`
	// Unknown
	PhotonNameserverOverrides []string `json:"photonNameserverOverrides"`
	// Unknown
	PhotonPublicKeys []string `json:"photonPublicKeys"`
	// Currently used youtube-dl.exe hash in SHA1-delimited format
	PlayerUrlResolverSha1 string `json:"player-url-resolver-sha1"`
	// Currently used youtube-dl.exe version
	PlayerUrlResolverVersion string `json:"player-url-resolver-version"`
	// Public key, hex encoded
	PublicKey string `json:"publicKey"`
	// Categories available for reporting objectionable content
	ReportCategories map[string]ReportCategory `json:"reportCategories"`
	// URL to the report form
	ReportFormUrl string `json:"reportFormUrl"`
	// Options for reporting content. Select a key+value from this mapping as the `type` of the report. Select one key+value from the object at reportOptions[type] as the `category` of the report. reportCategories[category] contains user-facing text to display for all possible categories. Select one value from the array at reportOptions[type][category] as the `reason` of the report. reportReasons[reason] contains user-facing text to display for all possible categories.
	ReportOptions map[string]map[string][]string `json:"reportOptions"`
	// Reasons available for submitting a report
	ReportReasons map[string]ReportReason `json:"reportReasons"`
	RequireAgeVerificationBetaTag bool `json:"requireAgeVerificationBetaTag"`
	// Link to the developer FAQ
	SdkDeveloperFaqUrl string `json:"sdkDeveloperFaqUrl"`
	// Link to the official VRChat Discord
	SdkDiscordUrl string `json:"sdkDiscordUrl"`
	// Used in the SDK to notify a user they aren't allowed to upload avatars/worlds yet
	SdkNotAllowedToPublishMessage string `json:"sdkNotAllowedToPublishMessage"`
	// Unity version supported by the SDK
	SdkUnityVersion string `json:"sdkUnityVersion"`
	// A list of explicitly allowed origins that worlds can request strings from via the Udon's [VRCStringDownloader.LoadUrl](https://creators.vrchat.com/worlds/udon/string-loading/#ivrcstringdownload).
	StringHostUrlList []string `json:"stringHostUrlList"`
	// VRChat's support email
	SupportEmail string `json:"supportEmail"`
	// VRChat's support form
	SupportFormUrl string `json:"supportFormUrl"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	TimeOutWorldId string `json:"timeOutWorldId"`
	// Unknown
	Timekeeping bool `json:"timekeeping"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	TutorialWorldId string `json:"tutorialWorldId"`
	// Unknown
	UpdateRateMsMaximum int32 `json:"updateRateMsMaximum"`
	// Unknown
	UpdateRateMsMinimum int32 `json:"updateRateMsMinimum"`
	// Unknown
	UpdateRateMsNormal int32 `json:"updateRateMsNormal"`
	// Unknown
	UpdateRateMsUdonManual int32 `json:"updateRateMsUdonManual"`
	// Unknown
	UploadAnalysisPercent int32 `json:"uploadAnalysisPercent"`
	// List of allowed URLs that bypass the \"Allow untrusted URL's\" setting in-game
	UrlList []string `json:"urlList"`
	// Unknown
	UseReliableUdpForVoice bool `json:"useReliableUdpForVoice"`
	// Download link for game on the Steam website.
	ViveWindowsUrl string `json:"viveWindowsUrl"`
	// Unknown
	WebsocketMaxFriendsRefreshDelay int32 `json:"websocketMaxFriendsRefreshDelay"`
	// Unknown
	WebsocketQuickReconnectTime int32 `json:"websocketQuickReconnectTime"`
	// Unknown
	WebsocketReconnectMaxDelay int32 `json:"websocketReconnectMaxDelay"`
	// List of allowed URLs that are allowed to host avatar assets
	WhiteListedAssetUrls []string `json:"whiteListedAssetUrls"`
}

type _APIConfig APIConfig

// NewAPIConfig instantiates a new APIConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIConfig(campaignStatus string, disableBackgroundPreloads bool, locationGiftingNonSubPrioEnabled bool, voiceEnableDegradation bool, voiceEnableReceiverLimiting bool, accessLogsUrls APIConfigAccessLogsUrls, address string, ageVerificationInviteVisible bool, ageVerificationP bool, ageVerificationStatusVisible bool, analysisMaxRetries int32, analysisRetryInterval int32, analyticsSegmentNewUIPctOfUsers int32, analyticsSegmentNewUISalt string, announcements []APIConfigAnnouncement, availableLanguageCodes []string, availableLanguages []string, avatarPerfLimiter APIConfigAvatarPerfLimiter, chatboxLogBufferSeconds int32, clientApiKey string, clientBPSCeiling int32, clientDisconnectTimeout int32, clientNetDispatchThreadMobile bool, clientReservedPlayerBPS int32, clientSentCountAllowance int32, constants APIConfigConstants, contactEmail string, copyrightEmail string, copyrightFormUrl string, currentPrivacyVersion int32, currentTOSVersion int32, defaultAvatar string, defaultStickerSet string, devSdkUrl string, devSdkVersion string, disCountdown time.Time, disableAvatarCopying bool, disableAvatarGating bool, disableCommunityLabs bool, disableCommunityLabsPromotion bool, disableEmail bool, disableEventStream bool, disableFeedbackGating bool, disableFrontendBuilds bool, disableGiftDrops bool, disableHello bool, disableOculusSubs bool, disableRegistration bool, disableSteamNetworking bool, disableTwoFactorAuth bool, disableUdon bool, disableUpgradeAccount bool, downloadLinkWindows string, downloadUrls APIConfigDownloadURLList, dynamicWorldRows []DynamicContentRow, economyLedgerBackfill bool, economyLedgerMigrationStop string, economyLedgerMode string, economyPauseEnd time.Time, economyPauseStart time.Time, economyPurchaseRepairEnabled bool, economyState int32, events APIConfigEvents, forceUseLatestWorld bool, giftDisplayType string, googleApiClientId string, homeWorldId string, homepageRedirectTarget string, hubWorldId string, imageHostUrlList []string, iosAppVersion []string, iosVersion APIConfigIosVersion, jobsEmail string, maxUserEmoji int32, maxUserStickers int32, minSupportedClientBuildNumber APIConfigMinSupportedClientBuildNumber, minimumUnityVersionForUploads string, moderationEmail string, notAllowedToSelectAvatarInPrivateWorldMessage string, offlineAnalysis APIConfigOfflineAnalysis, photonNameserverOverrides []string, photonPublicKeys []string, playerUrlResolverSha1 string, playerUrlResolverVersion string, publicKey string, reportCategories map[string]ReportCategory, reportFormUrl string, reportOptions map[string]map[string][]string, reportReasons map[string]ReportReason, requireAgeVerificationBetaTag bool, sdkDeveloperFaqUrl string, sdkDiscordUrl string, sdkNotAllowedToPublishMessage string, sdkUnityVersion string, stringHostUrlList []string, supportEmail string, supportFormUrl string, timeOutWorldId string, timekeeping bool, tutorialWorldId string, updateRateMsMaximum int32, updateRateMsMinimum int32, updateRateMsNormal int32, updateRateMsUdonManual int32, uploadAnalysisPercent int32, urlList []string, useReliableUdpForVoice bool, viveWindowsUrl string, websocketMaxFriendsRefreshDelay int32, websocketQuickReconnectTime int32, websocketReconnectMaxDelay int32, whiteListedAssetUrls []string) *APIConfig {
	this := APIConfig{}
	this.CampaignStatus = campaignStatus
	this.DisableBackgroundPreloads = disableBackgroundPreloads
	this.LocationGiftingNonSubPrioEnabled = locationGiftingNonSubPrioEnabled
	this.VoiceEnableDegradation = voiceEnableDegradation
	this.VoiceEnableReceiverLimiting = voiceEnableReceiverLimiting
	this.AccessLogsUrls = accessLogsUrls
	this.Address = address
	this.AgeVerificationInviteVisible = ageVerificationInviteVisible
	this.AgeVerificationP = ageVerificationP
	this.AgeVerificationStatusVisible = ageVerificationStatusVisible
	this.AnalysisMaxRetries = analysisMaxRetries
	this.AnalysisRetryInterval = analysisRetryInterval
	this.AnalyticsSegmentNewUIPctOfUsers = analyticsSegmentNewUIPctOfUsers
	this.AnalyticsSegmentNewUISalt = analyticsSegmentNewUISalt
	this.Announcements = announcements
	this.AvailableLanguageCodes = availableLanguageCodes
	this.AvailableLanguages = availableLanguages
	this.AvatarPerfLimiter = avatarPerfLimiter
	this.ChatboxLogBufferSeconds = chatboxLogBufferSeconds
	this.ClientApiKey = clientApiKey
	this.ClientBPSCeiling = clientBPSCeiling
	this.ClientDisconnectTimeout = clientDisconnectTimeout
	var clientNetDispatchThread bool = false
	this.ClientNetDispatchThread = &clientNetDispatchThread
	this.ClientNetDispatchThreadMobile = clientNetDispatchThreadMobile
	var clientNetInThread bool = false
	this.ClientNetInThread = &clientNetInThread
	var clientNetInThread2 bool = false
	this.ClientNetInThread2 = &clientNetInThread2
	var clientNetInThreadMobile bool = false
	this.ClientNetInThreadMobile = &clientNetInThreadMobile
	var clientNetInThreadMobile2 bool = false
	this.ClientNetInThreadMobile2 = &clientNetInThreadMobile2
	var clientNetOutThread bool = false
	this.ClientNetOutThread = &clientNetOutThread
	var clientNetOutThread2 bool = false
	this.ClientNetOutThread2 = &clientNetOutThread2
	var clientNetOutThreadMobile bool = false
	this.ClientNetOutThreadMobile = &clientNetOutThreadMobile
	var clientNetOutThreadMobile2 bool = false
	this.ClientNetOutThreadMobile2 = &clientNetOutThreadMobile2
	var clientQR int32 = 1
	this.ClientQR = &clientQR
	this.ClientReservedPlayerBPS = clientReservedPlayerBPS
	this.ClientSentCountAllowance = clientSentCountAllowance
	this.Constants = constants
	this.ContactEmail = contactEmail
	this.CopyrightEmail = copyrightEmail
	this.CopyrightFormUrl = copyrightFormUrl
	this.CurrentPrivacyVersion = currentPrivacyVersion
	this.CurrentTOSVersion = currentTOSVersion
	this.DefaultAvatar = defaultAvatar
	this.DefaultStickerSet = defaultStickerSet
	this.DevSdkUrl = devSdkUrl
	this.DevSdkVersion = devSdkVersion
	this.DisCountdown = disCountdown
	var disableAVProInProton bool = false
	this.DisableAVProInProton = &disableAVProInProton
	this.DisableAvatarCopying = disableAvatarCopying
	this.DisableAvatarGating = disableAvatarGating
	var disableCaptcha bool = true
	this.DisableCaptcha = &disableCaptcha
	this.DisableCommunityLabs = disableCommunityLabs
	this.DisableCommunityLabsPromotion = disableCommunityLabsPromotion
	this.DisableEmail = disableEmail
	this.DisableEventStream = disableEventStream
	this.DisableFeedbackGating = disableFeedbackGating
	this.DisableFrontendBuilds = disableFrontendBuilds
	this.DisableGiftDrops = disableGiftDrops
	this.DisableHello = disableHello
	this.DisableOculusSubs = disableOculusSubs
	this.DisableRegistration = disableRegistration
	this.DisableSteamNetworking = disableSteamNetworking
	this.DisableTwoFactorAuth = disableTwoFactorAuth
	this.DisableUdon = disableUdon
	this.DisableUpgradeAccount = disableUpgradeAccount
	this.DownloadLinkWindows = downloadLinkWindows
	this.DownloadUrls = downloadUrls
	this.DynamicWorldRows = dynamicWorldRows
	this.EconomyLedgerBackfill = economyLedgerBackfill
	this.EconomyLedgerMigrationStop = economyLedgerMigrationStop
	this.EconomyLedgerMode = economyLedgerMode
	this.EconomyPauseEnd = economyPauseEnd
	this.EconomyPauseStart = economyPauseStart
	this.EconomyPurchaseRepairEnabled = economyPurchaseRepairEnabled
	this.EconomyState = economyState
	this.Events = events
	this.ForceUseLatestWorld = forceUseLatestWorld
	this.GiftDisplayType = giftDisplayType
	this.GoogleApiClientId = googleApiClientId
	this.HomeWorldId = homeWorldId
	this.HomepageRedirectTarget = homepageRedirectTarget
	this.HubWorldId = hubWorldId
	this.ImageHostUrlList = imageHostUrlList
	this.IosAppVersion = iosAppVersion
	this.IosVersion = iosVersion
	this.JobsEmail = jobsEmail
	this.MaxUserEmoji = maxUserEmoji
	this.MaxUserStickers = maxUserStickers
	this.MinSupportedClientBuildNumber = minSupportedClientBuildNumber
	this.MinimumUnityVersionForUploads = minimumUnityVersionForUploads
	this.ModerationEmail = moderationEmail
	this.NotAllowedToSelectAvatarInPrivateWorldMessage = notAllowedToSelectAvatarInPrivateWorldMessage
	this.OfflineAnalysis = offlineAnalysis
	this.PhotonNameserverOverrides = photonNameserverOverrides
	this.PhotonPublicKeys = photonPublicKeys
	this.PlayerUrlResolverSha1 = playerUrlResolverSha1
	this.PlayerUrlResolverVersion = playerUrlResolverVersion
	this.PublicKey = publicKey
	this.ReportCategories = reportCategories
	this.ReportFormUrl = reportFormUrl
	this.ReportOptions = reportOptions
	this.ReportReasons = reportReasons
	this.RequireAgeVerificationBetaTag = requireAgeVerificationBetaTag
	this.SdkDeveloperFaqUrl = sdkDeveloperFaqUrl
	this.SdkDiscordUrl = sdkDiscordUrl
	this.SdkNotAllowedToPublishMessage = sdkNotAllowedToPublishMessage
	this.SdkUnityVersion = sdkUnityVersion
	this.StringHostUrlList = stringHostUrlList
	this.SupportEmail = supportEmail
	this.SupportFormUrl = supportFormUrl
	this.TimeOutWorldId = timeOutWorldId
	this.Timekeeping = timekeeping
	this.TutorialWorldId = tutorialWorldId
	this.UpdateRateMsMaximum = updateRateMsMaximum
	this.UpdateRateMsMinimum = updateRateMsMinimum
	this.UpdateRateMsNormal = updateRateMsNormal
	this.UpdateRateMsUdonManual = updateRateMsUdonManual
	this.UploadAnalysisPercent = uploadAnalysisPercent
	this.UrlList = urlList
	this.UseReliableUdpForVoice = useReliableUdpForVoice
	this.ViveWindowsUrl = viveWindowsUrl
	this.WebsocketMaxFriendsRefreshDelay = websocketMaxFriendsRefreshDelay
	this.WebsocketQuickReconnectTime = websocketQuickReconnectTime
	this.WebsocketReconnectMaxDelay = websocketReconnectMaxDelay
	this.WhiteListedAssetUrls = whiteListedAssetUrls
	return &this
}

// NewAPIConfigWithDefaults instantiates a new APIConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIConfigWithDefaults() *APIConfig {
	this := APIConfig{}
	var disableBackgroundPreloads bool = true
	this.DisableBackgroundPreloads = disableBackgroundPreloads
	var locationGiftingNonSubPrioEnabled bool = true
	this.LocationGiftingNonSubPrioEnabled = locationGiftingNonSubPrioEnabled
	var voiceEnableDegradation bool = false
	this.VoiceEnableDegradation = voiceEnableDegradation
	var voiceEnableReceiverLimiting bool = true
	this.VoiceEnableReceiverLimiting = voiceEnableReceiverLimiting
	var chatboxLogBufferSeconds int32 = 40
	this.ChatboxLogBufferSeconds = chatboxLogBufferSeconds
	var clientBPSCeiling int32 = 18432
	this.ClientBPSCeiling = clientBPSCeiling
	var clientDisconnectTimeout int32 = 30000
	this.ClientDisconnectTimeout = clientDisconnectTimeout
	var clientNetDispatchThread bool = false
	this.ClientNetDispatchThread = &clientNetDispatchThread
	var clientNetDispatchThreadMobile bool = true
	this.ClientNetDispatchThreadMobile = clientNetDispatchThreadMobile
	var clientNetInThread bool = false
	this.ClientNetInThread = &clientNetInThread
	var clientNetInThread2 bool = false
	this.ClientNetInThread2 = &clientNetInThread2
	var clientNetInThreadMobile bool = false
	this.ClientNetInThreadMobile = &clientNetInThreadMobile
	var clientNetInThreadMobile2 bool = false
	this.ClientNetInThreadMobile2 = &clientNetInThreadMobile2
	var clientNetOutThread bool = false
	this.ClientNetOutThread = &clientNetOutThread
	var clientNetOutThread2 bool = false
	this.ClientNetOutThread2 = &clientNetOutThread2
	var clientNetOutThreadMobile bool = false
	this.ClientNetOutThreadMobile = &clientNetOutThreadMobile
	var clientNetOutThreadMobile2 bool = false
	this.ClientNetOutThreadMobile2 = &clientNetOutThreadMobile2
	var clientQR int32 = 1
	this.ClientQR = &clientQR
	var clientReservedPlayerBPS int32 = 7168
	this.ClientReservedPlayerBPS = clientReservedPlayerBPS
	var clientSentCountAllowance int32 = 15
	this.ClientSentCountAllowance = clientSentCountAllowance
	var currentPrivacyVersion int32 = 1
	this.CurrentPrivacyVersion = currentPrivacyVersion
	var disableAVProInProton bool = false
	this.DisableAVProInProton = &disableAVProInProton
	var disableAvatarCopying bool = false
	this.DisableAvatarCopying = disableAvatarCopying
	var disableAvatarGating bool = false
	this.DisableAvatarGating = disableAvatarGating
	var disableCaptcha bool = true
	this.DisableCaptcha = &disableCaptcha
	var disableCommunityLabs bool = false
	this.DisableCommunityLabs = disableCommunityLabs
	var disableCommunityLabsPromotion bool = false
	this.DisableCommunityLabsPromotion = disableCommunityLabsPromotion
	var disableEmail bool = false
	this.DisableEmail = disableEmail
	var disableEventStream bool = false
	this.DisableEventStream = disableEventStream
	var disableFeedbackGating bool = false
	this.DisableFeedbackGating = disableFeedbackGating
	var disableFrontendBuilds bool = false
	this.DisableFrontendBuilds = disableFrontendBuilds
	var disableGiftDrops bool = false
	this.DisableGiftDrops = disableGiftDrops
	var disableHello bool = false
	this.DisableHello = disableHello
	var disableOculusSubs bool = false
	this.DisableOculusSubs = disableOculusSubs
	var disableRegistration bool = false
	this.DisableRegistration = disableRegistration
	var disableSteamNetworking bool = true
	this.DisableSteamNetworking = disableSteamNetworking
	var disableTwoFactorAuth bool = false
	this.DisableTwoFactorAuth = disableTwoFactorAuth
	var disableUdon bool = false
	this.DisableUdon = disableUdon
	var disableUpgradeAccount bool = false
	this.DisableUpgradeAccount = disableUpgradeAccount
	var economyState int32 = 1
	this.EconomyState = economyState
	var forceUseLatestWorld bool = true
	this.ForceUseLatestWorld = forceUseLatestWorld
	var googleApiClientId string = "827942544393-r2ouvckvouldn9dg9uruseje575e878f.apps.googleusercontent.com"
	this.GoogleApiClientId = googleApiClientId
	var homepageRedirectTarget string = "https://hello.vrchat.com"
	this.HomepageRedirectTarget = homepageRedirectTarget
	var maxUserEmoji int32 = 18
	this.MaxUserEmoji = maxUserEmoji
	var maxUserStickers int32 = 18
	this.MaxUserStickers = maxUserStickers
	var minimumUnityVersionForUploads string = "2019.0.0f1"
	this.MinimumUnityVersionForUploads = minimumUnityVersionForUploads
	var reportFormUrl string = "https://help.vrchat.com/hc/en-us/requests/new?ticket_form_id=1500000182242&tf_360056455174=user_report&tf_360057451993={userId}&tf_1500001445142={reportedId}&tf_subject={reason} {category} By {contentType} {reportedName}&tf_description={description}"
	this.ReportFormUrl = reportFormUrl
	var timekeeping bool = true
	this.Timekeeping = timekeeping
	var useReliableUdpForVoice bool = false
	this.UseReliableUdpForVoice = useReliableUdpForVoice
	var websocketMaxFriendsRefreshDelay int32 = 900
	this.WebsocketMaxFriendsRefreshDelay = websocketMaxFriendsRefreshDelay
	var websocketQuickReconnectTime int32 = 2
	this.WebsocketQuickReconnectTime = websocketQuickReconnectTime
	var websocketReconnectMaxDelay int32 = 2
	this.WebsocketReconnectMaxDelay = websocketReconnectMaxDelay
	return &this
}

// GetCampaignStatus returns the CampaignStatus field value
func (o *APIConfig) GetCampaignStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignStatus
}

// GetCampaignStatusOk returns a tuple with the CampaignStatus field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetCampaignStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignStatus, true
}

// SetCampaignStatus sets field value
func (o *APIConfig) SetCampaignStatus(v string) {
	o.CampaignStatus = v
}

// GetDisableBackgroundPreloads returns the DisableBackgroundPreloads field value
func (o *APIConfig) GetDisableBackgroundPreloads() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableBackgroundPreloads
}

// GetDisableBackgroundPreloadsOk returns a tuple with the DisableBackgroundPreloads field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableBackgroundPreloadsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableBackgroundPreloads, true
}

// SetDisableBackgroundPreloads sets field value
func (o *APIConfig) SetDisableBackgroundPreloads(v bool) {
	o.DisableBackgroundPreloads = v
}

// GetLocationGiftingNonSubPrioEnabled returns the LocationGiftingNonSubPrioEnabled field value
func (o *APIConfig) GetLocationGiftingNonSubPrioEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LocationGiftingNonSubPrioEnabled
}

// GetLocationGiftingNonSubPrioEnabledOk returns a tuple with the LocationGiftingNonSubPrioEnabled field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetLocationGiftingNonSubPrioEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationGiftingNonSubPrioEnabled, true
}

// SetLocationGiftingNonSubPrioEnabled sets field value
func (o *APIConfig) SetLocationGiftingNonSubPrioEnabled(v bool) {
	o.LocationGiftingNonSubPrioEnabled = v
}

// GetVoiceEnableDegradation returns the VoiceEnableDegradation field value
func (o *APIConfig) GetVoiceEnableDegradation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VoiceEnableDegradation
}

// GetVoiceEnableDegradationOk returns a tuple with the VoiceEnableDegradation field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetVoiceEnableDegradationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VoiceEnableDegradation, true
}

// SetVoiceEnableDegradation sets field value
func (o *APIConfig) SetVoiceEnableDegradation(v bool) {
	o.VoiceEnableDegradation = v
}

// GetVoiceEnableReceiverLimiting returns the VoiceEnableReceiverLimiting field value
func (o *APIConfig) GetVoiceEnableReceiverLimiting() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VoiceEnableReceiverLimiting
}

// GetVoiceEnableReceiverLimitingOk returns a tuple with the VoiceEnableReceiverLimiting field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetVoiceEnableReceiverLimitingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VoiceEnableReceiverLimiting, true
}

// SetVoiceEnableReceiverLimiting sets field value
func (o *APIConfig) SetVoiceEnableReceiverLimiting(v bool) {
	o.VoiceEnableReceiverLimiting = v
}

// GetAccessLogsUrls returns the AccessLogsUrls field value
func (o *APIConfig) GetAccessLogsUrls() APIConfigAccessLogsUrls {
	if o == nil {
		var ret APIConfigAccessLogsUrls
		return ret
	}

	return o.AccessLogsUrls
}

// GetAccessLogsUrlsOk returns a tuple with the AccessLogsUrls field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAccessLogsUrlsOk() (*APIConfigAccessLogsUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessLogsUrls, true
}

// SetAccessLogsUrls sets field value
func (o *APIConfig) SetAccessLogsUrls(v APIConfigAccessLogsUrls) {
	o.AccessLogsUrls = v
}

// GetAddress returns the Address field value
func (o *APIConfig) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *APIConfig) SetAddress(v string) {
	o.Address = v
}

// GetAgeVerificationInviteVisible returns the AgeVerificationInviteVisible field value
func (o *APIConfig) GetAgeVerificationInviteVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgeVerificationInviteVisible
}

// GetAgeVerificationInviteVisibleOk returns a tuple with the AgeVerificationInviteVisible field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAgeVerificationInviteVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgeVerificationInviteVisible, true
}

// SetAgeVerificationInviteVisible sets field value
func (o *APIConfig) SetAgeVerificationInviteVisible(v bool) {
	o.AgeVerificationInviteVisible = v
}

// GetAgeVerificationP returns the AgeVerificationP field value
func (o *APIConfig) GetAgeVerificationP() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgeVerificationP
}

// GetAgeVerificationPOk returns a tuple with the AgeVerificationP field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAgeVerificationPOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgeVerificationP, true
}

// SetAgeVerificationP sets field value
func (o *APIConfig) SetAgeVerificationP(v bool) {
	o.AgeVerificationP = v
}

// GetAgeVerificationStatusVisible returns the AgeVerificationStatusVisible field value
func (o *APIConfig) GetAgeVerificationStatusVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgeVerificationStatusVisible
}

// GetAgeVerificationStatusVisibleOk returns a tuple with the AgeVerificationStatusVisible field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAgeVerificationStatusVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgeVerificationStatusVisible, true
}

// SetAgeVerificationStatusVisible sets field value
func (o *APIConfig) SetAgeVerificationStatusVisible(v bool) {
	o.AgeVerificationStatusVisible = v
}

// GetAnalysisMaxRetries returns the AnalysisMaxRetries field value
func (o *APIConfig) GetAnalysisMaxRetries() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AnalysisMaxRetries
}

// GetAnalysisMaxRetriesOk returns a tuple with the AnalysisMaxRetries field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAnalysisMaxRetriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalysisMaxRetries, true
}

// SetAnalysisMaxRetries sets field value
func (o *APIConfig) SetAnalysisMaxRetries(v int32) {
	o.AnalysisMaxRetries = v
}

// GetAnalysisRetryInterval returns the AnalysisRetryInterval field value
func (o *APIConfig) GetAnalysisRetryInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AnalysisRetryInterval
}

// GetAnalysisRetryIntervalOk returns a tuple with the AnalysisRetryInterval field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAnalysisRetryIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalysisRetryInterval, true
}

// SetAnalysisRetryInterval sets field value
func (o *APIConfig) SetAnalysisRetryInterval(v int32) {
	o.AnalysisRetryInterval = v
}

// GetAnalyticsSegmentNewUIPctOfUsers returns the AnalyticsSegmentNewUIPctOfUsers field value
func (o *APIConfig) GetAnalyticsSegmentNewUIPctOfUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AnalyticsSegmentNewUIPctOfUsers
}

// GetAnalyticsSegmentNewUIPctOfUsersOk returns a tuple with the AnalyticsSegmentNewUIPctOfUsers field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAnalyticsSegmentNewUIPctOfUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyticsSegmentNewUIPctOfUsers, true
}

// SetAnalyticsSegmentNewUIPctOfUsers sets field value
func (o *APIConfig) SetAnalyticsSegmentNewUIPctOfUsers(v int32) {
	o.AnalyticsSegmentNewUIPctOfUsers = v
}

// GetAnalyticsSegmentNewUISalt returns the AnalyticsSegmentNewUISalt field value
func (o *APIConfig) GetAnalyticsSegmentNewUISalt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnalyticsSegmentNewUISalt
}

// GetAnalyticsSegmentNewUISaltOk returns a tuple with the AnalyticsSegmentNewUISalt field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAnalyticsSegmentNewUISaltOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyticsSegmentNewUISalt, true
}

// SetAnalyticsSegmentNewUISalt sets field value
func (o *APIConfig) SetAnalyticsSegmentNewUISalt(v string) {
	o.AnalyticsSegmentNewUISalt = v
}

// GetAnnouncements returns the Announcements field value
func (o *APIConfig) GetAnnouncements() []APIConfigAnnouncement {
	if o == nil {
		var ret []APIConfigAnnouncement
		return ret
	}

	return o.Announcements
}

// GetAnnouncementsOk returns a tuple with the Announcements field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAnnouncementsOk() ([]APIConfigAnnouncement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Announcements, true
}

// SetAnnouncements sets field value
func (o *APIConfig) SetAnnouncements(v []APIConfigAnnouncement) {
	o.Announcements = v
}

// GetAudioConfig returns the AudioConfig field value if set, zero value otherwise.
func (o *APIConfig) GetAudioConfig() APIConfigAudioConfig {
	if o == nil || IsNil(o.AudioConfig) {
		var ret APIConfigAudioConfig
		return ret
	}
	return *o.AudioConfig
}

// GetAudioConfigOk returns a tuple with the AudioConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAudioConfigOk() (*APIConfigAudioConfig, bool) {
	if o == nil || IsNil(o.AudioConfig) {
		return nil, false
	}
	return o.AudioConfig, true
}

// HasAudioConfig returns a boolean if a field has been set.
func (o *APIConfig) HasAudioConfig() bool {
	if o != nil && !IsNil(o.AudioConfig) {
		return true
	}

	return false
}

// SetAudioConfig gets a reference to the given APIConfigAudioConfig and assigns it to the AudioConfig field.
func (o *APIConfig) SetAudioConfig(v APIConfigAudioConfig) {
	o.AudioConfig = &v
}

// GetAvailableLanguageCodes returns the AvailableLanguageCodes field value
func (o *APIConfig) GetAvailableLanguageCodes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableLanguageCodes
}

// GetAvailableLanguageCodesOk returns a tuple with the AvailableLanguageCodes field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAvailableLanguageCodesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableLanguageCodes, true
}

// SetAvailableLanguageCodes sets field value
func (o *APIConfig) SetAvailableLanguageCodes(v []string) {
	o.AvailableLanguageCodes = v
}

// GetAvailableLanguages returns the AvailableLanguages field value
func (o *APIConfig) GetAvailableLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableLanguages
}

// GetAvailableLanguagesOk returns a tuple with the AvailableLanguages field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAvailableLanguagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableLanguages, true
}

// SetAvailableLanguages sets field value
func (o *APIConfig) SetAvailableLanguages(v []string) {
	o.AvailableLanguages = v
}

// GetAvatarPerfLimiter returns the AvatarPerfLimiter field value
func (o *APIConfig) GetAvatarPerfLimiter() APIConfigAvatarPerfLimiter {
	if o == nil {
		var ret APIConfigAvatarPerfLimiter
		return ret
	}

	return o.AvatarPerfLimiter
}

// GetAvatarPerfLimiterOk returns a tuple with the AvatarPerfLimiter field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAvatarPerfLimiterOk() (*APIConfigAvatarPerfLimiter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarPerfLimiter, true
}

// SetAvatarPerfLimiter sets field value
func (o *APIConfig) SetAvatarPerfLimiter(v APIConfigAvatarPerfLimiter) {
	o.AvatarPerfLimiter = v
}

// GetChatboxLogBufferSeconds returns the ChatboxLogBufferSeconds field value
func (o *APIConfig) GetChatboxLogBufferSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChatboxLogBufferSeconds
}

// GetChatboxLogBufferSecondsOk returns a tuple with the ChatboxLogBufferSeconds field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetChatboxLogBufferSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatboxLogBufferSeconds, true
}

// SetChatboxLogBufferSeconds sets field value
func (o *APIConfig) SetChatboxLogBufferSeconds(v int32) {
	o.ChatboxLogBufferSeconds = v
}

// GetClientApiKey returns the ClientApiKey field value
func (o *APIConfig) GetClientApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientApiKey
}

// GetClientApiKeyOk returns a tuple with the ClientApiKey field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientApiKey, true
}

// SetClientApiKey sets field value
func (o *APIConfig) SetClientApiKey(v string) {
	o.ClientApiKey = v
}

// GetClientBPSCeiling returns the ClientBPSCeiling field value
func (o *APIConfig) GetClientBPSCeiling() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClientBPSCeiling
}

// GetClientBPSCeilingOk returns a tuple with the ClientBPSCeiling field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientBPSCeilingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientBPSCeiling, true
}

// SetClientBPSCeiling sets field value
func (o *APIConfig) SetClientBPSCeiling(v int32) {
	o.ClientBPSCeiling = v
}

// GetClientDisconnectTimeout returns the ClientDisconnectTimeout field value
func (o *APIConfig) GetClientDisconnectTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClientDisconnectTimeout
}

// GetClientDisconnectTimeoutOk returns a tuple with the ClientDisconnectTimeout field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientDisconnectTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientDisconnectTimeout, true
}

// SetClientDisconnectTimeout sets field value
func (o *APIConfig) SetClientDisconnectTimeout(v int32) {
	o.ClientDisconnectTimeout = v
}

// GetClientNetDispatchThread returns the ClientNetDispatchThread field value if set, zero value otherwise.
func (o *APIConfig) GetClientNetDispatchThread() bool {
	if o == nil || IsNil(o.ClientNetDispatchThread) {
		var ret bool
		return ret
	}
	return *o.ClientNetDispatchThread
}

// GetClientNetDispatchThreadOk returns a tuple with the ClientNetDispatchThread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientNetDispatchThreadOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientNetDispatchThread) {
		return nil, false
	}
	return o.ClientNetDispatchThread, true
}

// HasClientNetDispatchThread returns a boolean if a field has been set.
func (o *APIConfig) HasClientNetDispatchThread() bool {
	if o != nil && !IsNil(o.ClientNetDispatchThread) {
		return true
	}

	return false
}

// SetClientNetDispatchThread gets a reference to the given bool and assigns it to the ClientNetDispatchThread field.
func (o *APIConfig) SetClientNetDispatchThread(v bool) {
	o.ClientNetDispatchThread = &v
}

// GetClientNetDispatchThreadMobile returns the ClientNetDispatchThreadMobile field value
func (o *APIConfig) GetClientNetDispatchThreadMobile() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ClientNetDispatchThreadMobile
}

// GetClientNetDispatchThreadMobileOk returns a tuple with the ClientNetDispatchThreadMobile field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientNetDispatchThreadMobileOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientNetDispatchThreadMobile, true
}

// SetClientNetDispatchThreadMobile sets field value
func (o *APIConfig) SetClientNetDispatchThreadMobile(v bool) {
	o.ClientNetDispatchThreadMobile = v
}

// GetClientNetInThread returns the ClientNetInThread field value if set, zero value otherwise.
func (o *APIConfig) GetClientNetInThread() bool {
	if o == nil || IsNil(o.ClientNetInThread) {
		var ret bool
		return ret
	}
	return *o.ClientNetInThread
}

// GetClientNetInThreadOk returns a tuple with the ClientNetInThread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientNetInThreadOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientNetInThread) {
		return nil, false
	}
	return o.ClientNetInThread, true
}

// HasClientNetInThread returns a boolean if a field has been set.
func (o *APIConfig) HasClientNetInThread() bool {
	if o != nil && !IsNil(o.ClientNetInThread) {
		return true
	}

	return false
}

// SetClientNetInThread gets a reference to the given bool and assigns it to the ClientNetInThread field.
func (o *APIConfig) SetClientNetInThread(v bool) {
	o.ClientNetInThread = &v
}

// GetClientNetInThread2 returns the ClientNetInThread2 field value if set, zero value otherwise.
func (o *APIConfig) GetClientNetInThread2() bool {
	if o == nil || IsNil(o.ClientNetInThread2) {
		var ret bool
		return ret
	}
	return *o.ClientNetInThread2
}

// GetClientNetInThread2Ok returns a tuple with the ClientNetInThread2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientNetInThread2Ok() (*bool, bool) {
	if o == nil || IsNil(o.ClientNetInThread2) {
		return nil, false
	}
	return o.ClientNetInThread2, true
}

// HasClientNetInThread2 returns a boolean if a field has been set.
func (o *APIConfig) HasClientNetInThread2() bool {
	if o != nil && !IsNil(o.ClientNetInThread2) {
		return true
	}

	return false
}

// SetClientNetInThread2 gets a reference to the given bool and assigns it to the ClientNetInThread2 field.
func (o *APIConfig) SetClientNetInThread2(v bool) {
	o.ClientNetInThread2 = &v
}

// GetClientNetInThreadMobile returns the ClientNetInThreadMobile field value if set, zero value otherwise.
func (o *APIConfig) GetClientNetInThreadMobile() bool {
	if o == nil || IsNil(o.ClientNetInThreadMobile) {
		var ret bool
		return ret
	}
	return *o.ClientNetInThreadMobile
}

// GetClientNetInThreadMobileOk returns a tuple with the ClientNetInThreadMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientNetInThreadMobileOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientNetInThreadMobile) {
		return nil, false
	}
	return o.ClientNetInThreadMobile, true
}

// HasClientNetInThreadMobile returns a boolean if a field has been set.
func (o *APIConfig) HasClientNetInThreadMobile() bool {
	if o != nil && !IsNil(o.ClientNetInThreadMobile) {
		return true
	}

	return false
}

// SetClientNetInThreadMobile gets a reference to the given bool and assigns it to the ClientNetInThreadMobile field.
func (o *APIConfig) SetClientNetInThreadMobile(v bool) {
	o.ClientNetInThreadMobile = &v
}

// GetClientNetInThreadMobile2 returns the ClientNetInThreadMobile2 field value if set, zero value otherwise.
func (o *APIConfig) GetClientNetInThreadMobile2() bool {
	if o == nil || IsNil(o.ClientNetInThreadMobile2) {
		var ret bool
		return ret
	}
	return *o.ClientNetInThreadMobile2
}

// GetClientNetInThreadMobile2Ok returns a tuple with the ClientNetInThreadMobile2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientNetInThreadMobile2Ok() (*bool, bool) {
	if o == nil || IsNil(o.ClientNetInThreadMobile2) {
		return nil, false
	}
	return o.ClientNetInThreadMobile2, true
}

// HasClientNetInThreadMobile2 returns a boolean if a field has been set.
func (o *APIConfig) HasClientNetInThreadMobile2() bool {
	if o != nil && !IsNil(o.ClientNetInThreadMobile2) {
		return true
	}

	return false
}

// SetClientNetInThreadMobile2 gets a reference to the given bool and assigns it to the ClientNetInThreadMobile2 field.
func (o *APIConfig) SetClientNetInThreadMobile2(v bool) {
	o.ClientNetInThreadMobile2 = &v
}

// GetClientNetOutThread returns the ClientNetOutThread field value if set, zero value otherwise.
func (o *APIConfig) GetClientNetOutThread() bool {
	if o == nil || IsNil(o.ClientNetOutThread) {
		var ret bool
		return ret
	}
	return *o.ClientNetOutThread
}

// GetClientNetOutThreadOk returns a tuple with the ClientNetOutThread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientNetOutThreadOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientNetOutThread) {
		return nil, false
	}
	return o.ClientNetOutThread, true
}

// HasClientNetOutThread returns a boolean if a field has been set.
func (o *APIConfig) HasClientNetOutThread() bool {
	if o != nil && !IsNil(o.ClientNetOutThread) {
		return true
	}

	return false
}

// SetClientNetOutThread gets a reference to the given bool and assigns it to the ClientNetOutThread field.
func (o *APIConfig) SetClientNetOutThread(v bool) {
	o.ClientNetOutThread = &v
}

// GetClientNetOutThread2 returns the ClientNetOutThread2 field value if set, zero value otherwise.
func (o *APIConfig) GetClientNetOutThread2() bool {
	if o == nil || IsNil(o.ClientNetOutThread2) {
		var ret bool
		return ret
	}
	return *o.ClientNetOutThread2
}

// GetClientNetOutThread2Ok returns a tuple with the ClientNetOutThread2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientNetOutThread2Ok() (*bool, bool) {
	if o == nil || IsNil(o.ClientNetOutThread2) {
		return nil, false
	}
	return o.ClientNetOutThread2, true
}

// HasClientNetOutThread2 returns a boolean if a field has been set.
func (o *APIConfig) HasClientNetOutThread2() bool {
	if o != nil && !IsNil(o.ClientNetOutThread2) {
		return true
	}

	return false
}

// SetClientNetOutThread2 gets a reference to the given bool and assigns it to the ClientNetOutThread2 field.
func (o *APIConfig) SetClientNetOutThread2(v bool) {
	o.ClientNetOutThread2 = &v
}

// GetClientNetOutThreadMobile returns the ClientNetOutThreadMobile field value if set, zero value otherwise.
func (o *APIConfig) GetClientNetOutThreadMobile() bool {
	if o == nil || IsNil(o.ClientNetOutThreadMobile) {
		var ret bool
		return ret
	}
	return *o.ClientNetOutThreadMobile
}

// GetClientNetOutThreadMobileOk returns a tuple with the ClientNetOutThreadMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientNetOutThreadMobileOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientNetOutThreadMobile) {
		return nil, false
	}
	return o.ClientNetOutThreadMobile, true
}

// HasClientNetOutThreadMobile returns a boolean if a field has been set.
func (o *APIConfig) HasClientNetOutThreadMobile() bool {
	if o != nil && !IsNil(o.ClientNetOutThreadMobile) {
		return true
	}

	return false
}

// SetClientNetOutThreadMobile gets a reference to the given bool and assigns it to the ClientNetOutThreadMobile field.
func (o *APIConfig) SetClientNetOutThreadMobile(v bool) {
	o.ClientNetOutThreadMobile = &v
}

// GetClientNetOutThreadMobile2 returns the ClientNetOutThreadMobile2 field value if set, zero value otherwise.
func (o *APIConfig) GetClientNetOutThreadMobile2() bool {
	if o == nil || IsNil(o.ClientNetOutThreadMobile2) {
		var ret bool
		return ret
	}
	return *o.ClientNetOutThreadMobile2
}

// GetClientNetOutThreadMobile2Ok returns a tuple with the ClientNetOutThreadMobile2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientNetOutThreadMobile2Ok() (*bool, bool) {
	if o == nil || IsNil(o.ClientNetOutThreadMobile2) {
		return nil, false
	}
	return o.ClientNetOutThreadMobile2, true
}

// HasClientNetOutThreadMobile2 returns a boolean if a field has been set.
func (o *APIConfig) HasClientNetOutThreadMobile2() bool {
	if o != nil && !IsNil(o.ClientNetOutThreadMobile2) {
		return true
	}

	return false
}

// SetClientNetOutThreadMobile2 gets a reference to the given bool and assigns it to the ClientNetOutThreadMobile2 field.
func (o *APIConfig) SetClientNetOutThreadMobile2(v bool) {
	o.ClientNetOutThreadMobile2 = &v
}

// GetClientQR returns the ClientQR field value if set, zero value otherwise.
func (o *APIConfig) GetClientQR() int32 {
	if o == nil || IsNil(o.ClientQR) {
		var ret int32
		return ret
	}
	return *o.ClientQR
}

// GetClientQROk returns a tuple with the ClientQR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientQROk() (*int32, bool) {
	if o == nil || IsNil(o.ClientQR) {
		return nil, false
	}
	return o.ClientQR, true
}

// HasClientQR returns a boolean if a field has been set.
func (o *APIConfig) HasClientQR() bool {
	if o != nil && !IsNil(o.ClientQR) {
		return true
	}

	return false
}

// SetClientQR gets a reference to the given int32 and assigns it to the ClientQR field.
func (o *APIConfig) SetClientQR(v int32) {
	o.ClientQR = &v
}

// GetClientReservedPlayerBPS returns the ClientReservedPlayerBPS field value
func (o *APIConfig) GetClientReservedPlayerBPS() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClientReservedPlayerBPS
}

// GetClientReservedPlayerBPSOk returns a tuple with the ClientReservedPlayerBPS field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientReservedPlayerBPSOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientReservedPlayerBPS, true
}

// SetClientReservedPlayerBPS sets field value
func (o *APIConfig) SetClientReservedPlayerBPS(v int32) {
	o.ClientReservedPlayerBPS = v
}

// GetClientSentCountAllowance returns the ClientSentCountAllowance field value
func (o *APIConfig) GetClientSentCountAllowance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClientSentCountAllowance
}

// GetClientSentCountAllowanceOk returns a tuple with the ClientSentCountAllowance field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientSentCountAllowanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSentCountAllowance, true
}

// SetClientSentCountAllowance sets field value
func (o *APIConfig) SetClientSentCountAllowance(v int32) {
	o.ClientSentCountAllowance = v
}

// GetConstants returns the Constants field value
func (o *APIConfig) GetConstants() APIConfigConstants {
	if o == nil {
		var ret APIConfigConstants
		return ret
	}

	return o.Constants
}

// GetConstantsOk returns a tuple with the Constants field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetConstantsOk() (*APIConfigConstants, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Constants, true
}

// SetConstants sets field value
func (o *APIConfig) SetConstants(v APIConfigConstants) {
	o.Constants = v
}

// GetContactEmail returns the ContactEmail field value
func (o *APIConfig) GetContactEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetContactEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactEmail, true
}

// SetContactEmail sets field value
func (o *APIConfig) SetContactEmail(v string) {
	o.ContactEmail = v
}

// GetCopyrightEmail returns the CopyrightEmail field value
func (o *APIConfig) GetCopyrightEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CopyrightEmail
}

// GetCopyrightEmailOk returns a tuple with the CopyrightEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetCopyrightEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyrightEmail, true
}

// SetCopyrightEmail sets field value
func (o *APIConfig) SetCopyrightEmail(v string) {
	o.CopyrightEmail = v
}

// GetCopyrightFormUrl returns the CopyrightFormUrl field value
func (o *APIConfig) GetCopyrightFormUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CopyrightFormUrl
}

// GetCopyrightFormUrlOk returns a tuple with the CopyrightFormUrl field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetCopyrightFormUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyrightFormUrl, true
}

// SetCopyrightFormUrl sets field value
func (o *APIConfig) SetCopyrightFormUrl(v string) {
	o.CopyrightFormUrl = v
}

// GetCurrentPrivacyVersion returns the CurrentPrivacyVersion field value
func (o *APIConfig) GetCurrentPrivacyVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentPrivacyVersion
}

// GetCurrentPrivacyVersionOk returns a tuple with the CurrentPrivacyVersion field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetCurrentPrivacyVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPrivacyVersion, true
}

// SetCurrentPrivacyVersion sets field value
func (o *APIConfig) SetCurrentPrivacyVersion(v int32) {
	o.CurrentPrivacyVersion = v
}

// GetCurrentTOSVersion returns the CurrentTOSVersion field value
func (o *APIConfig) GetCurrentTOSVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentTOSVersion
}

// GetCurrentTOSVersionOk returns a tuple with the CurrentTOSVersion field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetCurrentTOSVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentTOSVersion, true
}

// SetCurrentTOSVersion sets field value
func (o *APIConfig) SetCurrentTOSVersion(v int32) {
	o.CurrentTOSVersion = v
}

// GetDefaultAvatar returns the DefaultAvatar field value
func (o *APIConfig) GetDefaultAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultAvatar
}

// GetDefaultAvatarOk returns a tuple with the DefaultAvatar field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDefaultAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultAvatar, true
}

// SetDefaultAvatar sets field value
func (o *APIConfig) SetDefaultAvatar(v string) {
	o.DefaultAvatar = v
}

// GetDefaultStickerSet returns the DefaultStickerSet field value
func (o *APIConfig) GetDefaultStickerSet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultStickerSet
}

// GetDefaultStickerSetOk returns a tuple with the DefaultStickerSet field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDefaultStickerSetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultStickerSet, true
}

// SetDefaultStickerSet sets field value
func (o *APIConfig) SetDefaultStickerSet(v string) {
	o.DefaultStickerSet = v
}

// GetDevLanguageCodes returns the DevLanguageCodes field value if set, zero value otherwise.
func (o *APIConfig) GetDevLanguageCodes() []string {
	if o == nil || IsNil(o.DevLanguageCodes) {
		var ret []string
		return ret
	}
	return o.DevLanguageCodes
}

// GetDevLanguageCodesOk returns a tuple with the DevLanguageCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDevLanguageCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.DevLanguageCodes) {
		return nil, false
	}
	return o.DevLanguageCodes, true
}

// HasDevLanguageCodes returns a boolean if a field has been set.
func (o *APIConfig) HasDevLanguageCodes() bool {
	if o != nil && !IsNil(o.DevLanguageCodes) {
		return true
	}

	return false
}

// SetDevLanguageCodes gets a reference to the given []string and assigns it to the DevLanguageCodes field.
func (o *APIConfig) SetDevLanguageCodes(v []string) {
	o.DevLanguageCodes = v
}

// GetDevSdkUrl returns the DevSdkUrl field value
// Deprecated
func (o *APIConfig) GetDevSdkUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevSdkUrl
}

// GetDevSdkUrlOk returns a tuple with the DevSdkUrl field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetDevSdkUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevSdkUrl, true
}

// SetDevSdkUrl sets field value
// Deprecated
func (o *APIConfig) SetDevSdkUrl(v string) {
	o.DevSdkUrl = v
}

// GetDevSdkVersion returns the DevSdkVersion field value
// Deprecated
func (o *APIConfig) GetDevSdkVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevSdkVersion
}

// GetDevSdkVersionOk returns a tuple with the DevSdkVersion field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetDevSdkVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevSdkVersion, true
}

// SetDevSdkVersion sets field value
// Deprecated
func (o *APIConfig) SetDevSdkVersion(v string) {
	o.DevSdkVersion = v
}

// GetDisCountdown returns the DisCountdown field value
func (o *APIConfig) GetDisCountdown() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DisCountdown
}

// GetDisCountdownOk returns a tuple with the DisCountdown field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisCountdownOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisCountdown, true
}

// SetDisCountdown sets field value
func (o *APIConfig) SetDisCountdown(v time.Time) {
	o.DisCountdown = v
}

// GetDisableAVProInProton returns the DisableAVProInProton field value if set, zero value otherwise.
func (o *APIConfig) GetDisableAVProInProton() bool {
	if o == nil || IsNil(o.DisableAVProInProton) {
		var ret bool
		return ret
	}
	return *o.DisableAVProInProton
}

// GetDisableAVProInProtonOk returns a tuple with the DisableAVProInProton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableAVProInProtonOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableAVProInProton) {
		return nil, false
	}
	return o.DisableAVProInProton, true
}

// HasDisableAVProInProton returns a boolean if a field has been set.
func (o *APIConfig) HasDisableAVProInProton() bool {
	if o != nil && !IsNil(o.DisableAVProInProton) {
		return true
	}

	return false
}

// SetDisableAVProInProton gets a reference to the given bool and assigns it to the DisableAVProInProton field.
func (o *APIConfig) SetDisableAVProInProton(v bool) {
	o.DisableAVProInProton = &v
}

// GetDisableAvatarCopying returns the DisableAvatarCopying field value
func (o *APIConfig) GetDisableAvatarCopying() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableAvatarCopying
}

// GetDisableAvatarCopyingOk returns a tuple with the DisableAvatarCopying field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableAvatarCopyingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableAvatarCopying, true
}

// SetDisableAvatarCopying sets field value
func (o *APIConfig) SetDisableAvatarCopying(v bool) {
	o.DisableAvatarCopying = v
}

// GetDisableAvatarGating returns the DisableAvatarGating field value
func (o *APIConfig) GetDisableAvatarGating() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableAvatarGating
}

// GetDisableAvatarGatingOk returns a tuple with the DisableAvatarGating field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableAvatarGatingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableAvatarGating, true
}

// SetDisableAvatarGating sets field value
func (o *APIConfig) SetDisableAvatarGating(v bool) {
	o.DisableAvatarGating = v
}

// GetDisableCaptcha returns the DisableCaptcha field value if set, zero value otherwise.
func (o *APIConfig) GetDisableCaptcha() bool {
	if o == nil || IsNil(o.DisableCaptcha) {
		var ret bool
		return ret
	}
	return *o.DisableCaptcha
}

// GetDisableCaptchaOk returns a tuple with the DisableCaptcha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableCaptchaOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableCaptcha) {
		return nil, false
	}
	return o.DisableCaptcha, true
}

// HasDisableCaptcha returns a boolean if a field has been set.
func (o *APIConfig) HasDisableCaptcha() bool {
	if o != nil && !IsNil(o.DisableCaptcha) {
		return true
	}

	return false
}

// SetDisableCaptcha gets a reference to the given bool and assigns it to the DisableCaptcha field.
func (o *APIConfig) SetDisableCaptcha(v bool) {
	o.DisableCaptcha = &v
}

// GetDisableCommunityLabs returns the DisableCommunityLabs field value
func (o *APIConfig) GetDisableCommunityLabs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableCommunityLabs
}

// GetDisableCommunityLabsOk returns a tuple with the DisableCommunityLabs field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableCommunityLabsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableCommunityLabs, true
}

// SetDisableCommunityLabs sets field value
func (o *APIConfig) SetDisableCommunityLabs(v bool) {
	o.DisableCommunityLabs = v
}

// GetDisableCommunityLabsPromotion returns the DisableCommunityLabsPromotion field value
func (o *APIConfig) GetDisableCommunityLabsPromotion() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableCommunityLabsPromotion
}

// GetDisableCommunityLabsPromotionOk returns a tuple with the DisableCommunityLabsPromotion field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableCommunityLabsPromotionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableCommunityLabsPromotion, true
}

// SetDisableCommunityLabsPromotion sets field value
func (o *APIConfig) SetDisableCommunityLabsPromotion(v bool) {
	o.DisableCommunityLabsPromotion = v
}

// GetDisableEmail returns the DisableEmail field value
func (o *APIConfig) GetDisableEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableEmail
}

// GetDisableEmailOk returns a tuple with the DisableEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableEmail, true
}

// SetDisableEmail sets field value
func (o *APIConfig) SetDisableEmail(v bool) {
	o.DisableEmail = v
}

// GetDisableEventStream returns the DisableEventStream field value
func (o *APIConfig) GetDisableEventStream() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableEventStream
}

// GetDisableEventStreamOk returns a tuple with the DisableEventStream field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableEventStreamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableEventStream, true
}

// SetDisableEventStream sets field value
func (o *APIConfig) SetDisableEventStream(v bool) {
	o.DisableEventStream = v
}

// GetDisableFeedbackGating returns the DisableFeedbackGating field value
func (o *APIConfig) GetDisableFeedbackGating() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableFeedbackGating
}

// GetDisableFeedbackGatingOk returns a tuple with the DisableFeedbackGating field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableFeedbackGatingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableFeedbackGating, true
}

// SetDisableFeedbackGating sets field value
func (o *APIConfig) SetDisableFeedbackGating(v bool) {
	o.DisableFeedbackGating = v
}

// GetDisableFrontendBuilds returns the DisableFrontendBuilds field value
func (o *APIConfig) GetDisableFrontendBuilds() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableFrontendBuilds
}

// GetDisableFrontendBuildsOk returns a tuple with the DisableFrontendBuilds field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableFrontendBuildsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableFrontendBuilds, true
}

// SetDisableFrontendBuilds sets field value
func (o *APIConfig) SetDisableFrontendBuilds(v bool) {
	o.DisableFrontendBuilds = v
}

// GetDisableGiftDrops returns the DisableGiftDrops field value
func (o *APIConfig) GetDisableGiftDrops() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableGiftDrops
}

// GetDisableGiftDropsOk returns a tuple with the DisableGiftDrops field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableGiftDropsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableGiftDrops, true
}

// SetDisableGiftDrops sets field value
func (o *APIConfig) SetDisableGiftDrops(v bool) {
	o.DisableGiftDrops = v
}

// GetDisableHello returns the DisableHello field value
func (o *APIConfig) GetDisableHello() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableHello
}

// GetDisableHelloOk returns a tuple with the DisableHello field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableHelloOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableHello, true
}

// SetDisableHello sets field value
func (o *APIConfig) SetDisableHello(v bool) {
	o.DisableHello = v
}

// GetDisableOculusSubs returns the DisableOculusSubs field value
func (o *APIConfig) GetDisableOculusSubs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableOculusSubs
}

// GetDisableOculusSubsOk returns a tuple with the DisableOculusSubs field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableOculusSubsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableOculusSubs, true
}

// SetDisableOculusSubs sets field value
func (o *APIConfig) SetDisableOculusSubs(v bool) {
	o.DisableOculusSubs = v
}

// GetDisableRegistration returns the DisableRegistration field value
func (o *APIConfig) GetDisableRegistration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableRegistration
}

// GetDisableRegistrationOk returns a tuple with the DisableRegistration field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableRegistrationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableRegistration, true
}

// SetDisableRegistration sets field value
func (o *APIConfig) SetDisableRegistration(v bool) {
	o.DisableRegistration = v
}

// GetDisableSteamNetworking returns the DisableSteamNetworking field value
func (o *APIConfig) GetDisableSteamNetworking() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableSteamNetworking
}

// GetDisableSteamNetworkingOk returns a tuple with the DisableSteamNetworking field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableSteamNetworkingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableSteamNetworking, true
}

// SetDisableSteamNetworking sets field value
func (o *APIConfig) SetDisableSteamNetworking(v bool) {
	o.DisableSteamNetworking = v
}

// GetDisableTwoFactorAuth returns the DisableTwoFactorAuth field value
// Deprecated
func (o *APIConfig) GetDisableTwoFactorAuth() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableTwoFactorAuth
}

// GetDisableTwoFactorAuthOk returns a tuple with the DisableTwoFactorAuth field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetDisableTwoFactorAuthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableTwoFactorAuth, true
}

// SetDisableTwoFactorAuth sets field value
// Deprecated
func (o *APIConfig) SetDisableTwoFactorAuth(v bool) {
	o.DisableTwoFactorAuth = v
}

// GetDisableUdon returns the DisableUdon field value
func (o *APIConfig) GetDisableUdon() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableUdon
}

// GetDisableUdonOk returns a tuple with the DisableUdon field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableUdonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableUdon, true
}

// SetDisableUdon sets field value
func (o *APIConfig) SetDisableUdon(v bool) {
	o.DisableUdon = v
}

// GetDisableUpgradeAccount returns the DisableUpgradeAccount field value
func (o *APIConfig) GetDisableUpgradeAccount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableUpgradeAccount
}

// GetDisableUpgradeAccountOk returns a tuple with the DisableUpgradeAccount field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableUpgradeAccountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableUpgradeAccount, true
}

// SetDisableUpgradeAccount sets field value
func (o *APIConfig) SetDisableUpgradeAccount(v bool) {
	o.DisableUpgradeAccount = v
}

// GetDownloadLinkWindows returns the DownloadLinkWindows field value
func (o *APIConfig) GetDownloadLinkWindows() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DownloadLinkWindows
}

// GetDownloadLinkWindowsOk returns a tuple with the DownloadLinkWindows field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDownloadLinkWindowsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadLinkWindows, true
}

// SetDownloadLinkWindows sets field value
func (o *APIConfig) SetDownloadLinkWindows(v string) {
	o.DownloadLinkWindows = v
}

// GetDownloadUrls returns the DownloadUrls field value
func (o *APIConfig) GetDownloadUrls() APIConfigDownloadURLList {
	if o == nil {
		var ret APIConfigDownloadURLList
		return ret
	}

	return o.DownloadUrls
}

// GetDownloadUrlsOk returns a tuple with the DownloadUrls field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDownloadUrlsOk() (*APIConfigDownloadURLList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadUrls, true
}

// SetDownloadUrls sets field value
func (o *APIConfig) SetDownloadUrls(v APIConfigDownloadURLList) {
	o.DownloadUrls = v
}

// GetDynamicWorldRows returns the DynamicWorldRows field value
func (o *APIConfig) GetDynamicWorldRows() []DynamicContentRow {
	if o == nil {
		var ret []DynamicContentRow
		return ret
	}

	return o.DynamicWorldRows
}

// GetDynamicWorldRowsOk returns a tuple with the DynamicWorldRows field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDynamicWorldRowsOk() ([]DynamicContentRow, bool) {
	if o == nil {
		return nil, false
	}
	return o.DynamicWorldRows, true
}

// SetDynamicWorldRows sets field value
func (o *APIConfig) SetDynamicWorldRows(v []DynamicContentRow) {
	o.DynamicWorldRows = v
}

// GetEconomyLedgerBackfill returns the EconomyLedgerBackfill field value
func (o *APIConfig) GetEconomyLedgerBackfill() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EconomyLedgerBackfill
}

// GetEconomyLedgerBackfillOk returns a tuple with the EconomyLedgerBackfill field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetEconomyLedgerBackfillOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EconomyLedgerBackfill, true
}

// SetEconomyLedgerBackfill sets field value
func (o *APIConfig) SetEconomyLedgerBackfill(v bool) {
	o.EconomyLedgerBackfill = v
}

// GetEconomyLedgerMigrationStop returns the EconomyLedgerMigrationStop field value
func (o *APIConfig) GetEconomyLedgerMigrationStop() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EconomyLedgerMigrationStop
}

// GetEconomyLedgerMigrationStopOk returns a tuple with the EconomyLedgerMigrationStop field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetEconomyLedgerMigrationStopOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EconomyLedgerMigrationStop, true
}

// SetEconomyLedgerMigrationStop sets field value
func (o *APIConfig) SetEconomyLedgerMigrationStop(v string) {
	o.EconomyLedgerMigrationStop = v
}

// GetEconomyLedgerMode returns the EconomyLedgerMode field value
func (o *APIConfig) GetEconomyLedgerMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EconomyLedgerMode
}

// GetEconomyLedgerModeOk returns a tuple with the EconomyLedgerMode field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetEconomyLedgerModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EconomyLedgerMode, true
}

// SetEconomyLedgerMode sets field value
func (o *APIConfig) SetEconomyLedgerMode(v string) {
	o.EconomyLedgerMode = v
}

// GetEconomyPauseEnd returns the EconomyPauseEnd field value
func (o *APIConfig) GetEconomyPauseEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EconomyPauseEnd
}

// GetEconomyPauseEndOk returns a tuple with the EconomyPauseEnd field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetEconomyPauseEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EconomyPauseEnd, true
}

// SetEconomyPauseEnd sets field value
func (o *APIConfig) SetEconomyPauseEnd(v time.Time) {
	o.EconomyPauseEnd = v
}

// GetEconomyPauseStart returns the EconomyPauseStart field value
func (o *APIConfig) GetEconomyPauseStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EconomyPauseStart
}

// GetEconomyPauseStartOk returns a tuple with the EconomyPauseStart field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetEconomyPauseStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EconomyPauseStart, true
}

// SetEconomyPauseStart sets field value
func (o *APIConfig) SetEconomyPauseStart(v time.Time) {
	o.EconomyPauseStart = v
}

// GetEconomyPurchaseRepairEnabled returns the EconomyPurchaseRepairEnabled field value
func (o *APIConfig) GetEconomyPurchaseRepairEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EconomyPurchaseRepairEnabled
}

// GetEconomyPurchaseRepairEnabledOk returns a tuple with the EconomyPurchaseRepairEnabled field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetEconomyPurchaseRepairEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EconomyPurchaseRepairEnabled, true
}

// SetEconomyPurchaseRepairEnabled sets field value
func (o *APIConfig) SetEconomyPurchaseRepairEnabled(v bool) {
	o.EconomyPurchaseRepairEnabled = v
}

// GetEconomyState returns the EconomyState field value
func (o *APIConfig) GetEconomyState() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EconomyState
}

// GetEconomyStateOk returns a tuple with the EconomyState field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetEconomyStateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EconomyState, true
}

// SetEconomyState sets field value
func (o *APIConfig) SetEconomyState(v int32) {
	o.EconomyState = v
}

// GetEvents returns the Events field value
func (o *APIConfig) GetEvents() APIConfigEvents {
	if o == nil {
		var ret APIConfigEvents
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetEventsOk() (*APIConfigEvents, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value
func (o *APIConfig) SetEvents(v APIConfigEvents) {
	o.Events = v
}

// GetForceUseLatestWorld returns the ForceUseLatestWorld field value
func (o *APIConfig) GetForceUseLatestWorld() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForceUseLatestWorld
}

// GetForceUseLatestWorldOk returns a tuple with the ForceUseLatestWorld field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetForceUseLatestWorldOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForceUseLatestWorld, true
}

// SetForceUseLatestWorld sets field value
func (o *APIConfig) SetForceUseLatestWorld(v bool) {
	o.ForceUseLatestWorld = v
}

// GetGiftDisplayType returns the GiftDisplayType field value
func (o *APIConfig) GetGiftDisplayType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GiftDisplayType
}

// GetGiftDisplayTypeOk returns a tuple with the GiftDisplayType field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetGiftDisplayTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GiftDisplayType, true
}

// SetGiftDisplayType sets field value
func (o *APIConfig) SetGiftDisplayType(v string) {
	o.GiftDisplayType = v
}

// GetGoogleApiClientId returns the GoogleApiClientId field value
func (o *APIConfig) GetGoogleApiClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoogleApiClientId
}

// GetGoogleApiClientIdOk returns a tuple with the GoogleApiClientId field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetGoogleApiClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoogleApiClientId, true
}

// SetGoogleApiClientId sets field value
func (o *APIConfig) SetGoogleApiClientId(v string) {
	o.GoogleApiClientId = v
}

// GetHomeWorldId returns the HomeWorldId field value
func (o *APIConfig) GetHomeWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HomeWorldId
}

// GetHomeWorldIdOk returns a tuple with the HomeWorldId field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetHomeWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomeWorldId, true
}

// SetHomeWorldId sets field value
func (o *APIConfig) SetHomeWorldId(v string) {
	o.HomeWorldId = v
}

// GetHomepageRedirectTarget returns the HomepageRedirectTarget field value
func (o *APIConfig) GetHomepageRedirectTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HomepageRedirectTarget
}

// GetHomepageRedirectTargetOk returns a tuple with the HomepageRedirectTarget field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetHomepageRedirectTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomepageRedirectTarget, true
}

// SetHomepageRedirectTarget sets field value
func (o *APIConfig) SetHomepageRedirectTarget(v string) {
	o.HomepageRedirectTarget = v
}

// GetHubWorldId returns the HubWorldId field value
func (o *APIConfig) GetHubWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HubWorldId
}

// GetHubWorldIdOk returns a tuple with the HubWorldId field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetHubWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubWorldId, true
}

// SetHubWorldId sets field value
func (o *APIConfig) SetHubWorldId(v string) {
	o.HubWorldId = v
}

// GetImageHostUrlList returns the ImageHostUrlList field value
func (o *APIConfig) GetImageHostUrlList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ImageHostUrlList
}

// GetImageHostUrlListOk returns a tuple with the ImageHostUrlList field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetImageHostUrlListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageHostUrlList, true
}

// SetImageHostUrlList sets field value
func (o *APIConfig) SetImageHostUrlList(v []string) {
	o.ImageHostUrlList = v
}

// GetIosAppVersion returns the IosAppVersion field value
func (o *APIConfig) GetIosAppVersion() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IosAppVersion
}

// GetIosAppVersionOk returns a tuple with the IosAppVersion field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetIosAppVersionOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IosAppVersion, true
}

// SetIosAppVersion sets field value
func (o *APIConfig) SetIosAppVersion(v []string) {
	o.IosAppVersion = v
}

// GetIosVersion returns the IosVersion field value
func (o *APIConfig) GetIosVersion() APIConfigIosVersion {
	if o == nil {
		var ret APIConfigIosVersion
		return ret
	}

	return o.IosVersion
}

// GetIosVersionOk returns a tuple with the IosVersion field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetIosVersionOk() (*APIConfigIosVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IosVersion, true
}

// SetIosVersion sets field value
func (o *APIConfig) SetIosVersion(v APIConfigIosVersion) {
	o.IosVersion = v
}

// GetJobsEmail returns the JobsEmail field value
func (o *APIConfig) GetJobsEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobsEmail
}

// GetJobsEmailOk returns a tuple with the JobsEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetJobsEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobsEmail, true
}

// SetJobsEmail sets field value
func (o *APIConfig) SetJobsEmail(v string) {
	o.JobsEmail = v
}

// GetMaxUserEmoji returns the MaxUserEmoji field value
func (o *APIConfig) GetMaxUserEmoji() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxUserEmoji
}

// GetMaxUserEmojiOk returns a tuple with the MaxUserEmoji field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetMaxUserEmojiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxUserEmoji, true
}

// SetMaxUserEmoji sets field value
func (o *APIConfig) SetMaxUserEmoji(v int32) {
	o.MaxUserEmoji = v
}

// GetMaxUserStickers returns the MaxUserStickers field value
func (o *APIConfig) GetMaxUserStickers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxUserStickers
}

// GetMaxUserStickersOk returns a tuple with the MaxUserStickers field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetMaxUserStickersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxUserStickers, true
}

// SetMaxUserStickers sets field value
func (o *APIConfig) SetMaxUserStickers(v int32) {
	o.MaxUserStickers = v
}

// GetMinSupportedClientBuildNumber returns the MinSupportedClientBuildNumber field value
func (o *APIConfig) GetMinSupportedClientBuildNumber() APIConfigMinSupportedClientBuildNumber {
	if o == nil {
		var ret APIConfigMinSupportedClientBuildNumber
		return ret
	}

	return o.MinSupportedClientBuildNumber
}

// GetMinSupportedClientBuildNumberOk returns a tuple with the MinSupportedClientBuildNumber field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetMinSupportedClientBuildNumberOk() (*APIConfigMinSupportedClientBuildNumber, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinSupportedClientBuildNumber, true
}

// SetMinSupportedClientBuildNumber sets field value
func (o *APIConfig) SetMinSupportedClientBuildNumber(v APIConfigMinSupportedClientBuildNumber) {
	o.MinSupportedClientBuildNumber = v
}

// GetMinimumUnityVersionForUploads returns the MinimumUnityVersionForUploads field value
func (o *APIConfig) GetMinimumUnityVersionForUploads() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumUnityVersionForUploads
}

// GetMinimumUnityVersionForUploadsOk returns a tuple with the MinimumUnityVersionForUploads field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetMinimumUnityVersionForUploadsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumUnityVersionForUploads, true
}

// SetMinimumUnityVersionForUploads sets field value
func (o *APIConfig) SetMinimumUnityVersionForUploads(v string) {
	o.MinimumUnityVersionForUploads = v
}

// GetModerationEmail returns the ModerationEmail field value
func (o *APIConfig) GetModerationEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModerationEmail
}

// GetModerationEmailOk returns a tuple with the ModerationEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetModerationEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModerationEmail, true
}

// SetModerationEmail sets field value
func (o *APIConfig) SetModerationEmail(v string) {
	o.ModerationEmail = v
}

// GetNotAllowedToSelectAvatarInPrivateWorldMessage returns the NotAllowedToSelectAvatarInPrivateWorldMessage field value
func (o *APIConfig) GetNotAllowedToSelectAvatarInPrivateWorldMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotAllowedToSelectAvatarInPrivateWorldMessage
}

// GetNotAllowedToSelectAvatarInPrivateWorldMessageOk returns a tuple with the NotAllowedToSelectAvatarInPrivateWorldMessage field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetNotAllowedToSelectAvatarInPrivateWorldMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotAllowedToSelectAvatarInPrivateWorldMessage, true
}

// SetNotAllowedToSelectAvatarInPrivateWorldMessage sets field value
func (o *APIConfig) SetNotAllowedToSelectAvatarInPrivateWorldMessage(v string) {
	o.NotAllowedToSelectAvatarInPrivateWorldMessage = v
}

// GetOfflineAnalysis returns the OfflineAnalysis field value
func (o *APIConfig) GetOfflineAnalysis() APIConfigOfflineAnalysis {
	if o == nil {
		var ret APIConfigOfflineAnalysis
		return ret
	}

	return o.OfflineAnalysis
}

// GetOfflineAnalysisOk returns a tuple with the OfflineAnalysis field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetOfflineAnalysisOk() (*APIConfigOfflineAnalysis, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfflineAnalysis, true
}

// SetOfflineAnalysis sets field value
func (o *APIConfig) SetOfflineAnalysis(v APIConfigOfflineAnalysis) {
	o.OfflineAnalysis = v
}

// GetPhotonNameserverOverrides returns the PhotonNameserverOverrides field value
func (o *APIConfig) GetPhotonNameserverOverrides() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PhotonNameserverOverrides
}

// GetPhotonNameserverOverridesOk returns a tuple with the PhotonNameserverOverrides field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetPhotonNameserverOverridesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhotonNameserverOverrides, true
}

// SetPhotonNameserverOverrides sets field value
func (o *APIConfig) SetPhotonNameserverOverrides(v []string) {
	o.PhotonNameserverOverrides = v
}

// GetPhotonPublicKeys returns the PhotonPublicKeys field value
func (o *APIConfig) GetPhotonPublicKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PhotonPublicKeys
}

// GetPhotonPublicKeysOk returns a tuple with the PhotonPublicKeys field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetPhotonPublicKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhotonPublicKeys, true
}

// SetPhotonPublicKeys sets field value
func (o *APIConfig) SetPhotonPublicKeys(v []string) {
	o.PhotonPublicKeys = v
}

// GetPlayerUrlResolverSha1 returns the PlayerUrlResolverSha1 field value
func (o *APIConfig) GetPlayerUrlResolverSha1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlayerUrlResolverSha1
}

// GetPlayerUrlResolverSha1Ok returns a tuple with the PlayerUrlResolverSha1 field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetPlayerUrlResolverSha1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayerUrlResolverSha1, true
}

// SetPlayerUrlResolverSha1 sets field value
func (o *APIConfig) SetPlayerUrlResolverSha1(v string) {
	o.PlayerUrlResolverSha1 = v
}

// GetPlayerUrlResolverVersion returns the PlayerUrlResolverVersion field value
func (o *APIConfig) GetPlayerUrlResolverVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlayerUrlResolverVersion
}

// GetPlayerUrlResolverVersionOk returns a tuple with the PlayerUrlResolverVersion field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetPlayerUrlResolverVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayerUrlResolverVersion, true
}

// SetPlayerUrlResolverVersion sets field value
func (o *APIConfig) SetPlayerUrlResolverVersion(v string) {
	o.PlayerUrlResolverVersion = v
}

// GetPublicKey returns the PublicKey field value
func (o *APIConfig) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *APIConfig) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetReportCategories returns the ReportCategories field value
func (o *APIConfig) GetReportCategories() map[string]ReportCategory {
	if o == nil {
		var ret map[string]ReportCategory
		return ret
	}

	return o.ReportCategories
}

// GetReportCategoriesOk returns a tuple with the ReportCategories field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetReportCategoriesOk() (*map[string]ReportCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportCategories, true
}

// SetReportCategories sets field value
func (o *APIConfig) SetReportCategories(v map[string]ReportCategory) {
	o.ReportCategories = v
}

// GetReportFormUrl returns the ReportFormUrl field value
func (o *APIConfig) GetReportFormUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportFormUrl
}

// GetReportFormUrlOk returns a tuple with the ReportFormUrl field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetReportFormUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportFormUrl, true
}

// SetReportFormUrl sets field value
func (o *APIConfig) SetReportFormUrl(v string) {
	o.ReportFormUrl = v
}

// GetReportOptions returns the ReportOptions field value
func (o *APIConfig) GetReportOptions() map[string]map[string][]string {
	if o == nil {
		var ret map[string]map[string][]string
		return ret
	}

	return o.ReportOptions
}

// GetReportOptionsOk returns a tuple with the ReportOptions field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetReportOptionsOk() (*map[string]map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportOptions, true
}

// SetReportOptions sets field value
func (o *APIConfig) SetReportOptions(v map[string]map[string][]string) {
	o.ReportOptions = v
}

// GetReportReasons returns the ReportReasons field value
func (o *APIConfig) GetReportReasons() map[string]ReportReason {
	if o == nil {
		var ret map[string]ReportReason
		return ret
	}

	return o.ReportReasons
}

// GetReportReasonsOk returns a tuple with the ReportReasons field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetReportReasonsOk() (*map[string]ReportReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportReasons, true
}

// SetReportReasons sets field value
func (o *APIConfig) SetReportReasons(v map[string]ReportReason) {
	o.ReportReasons = v
}

// GetRequireAgeVerificationBetaTag returns the RequireAgeVerificationBetaTag field value
func (o *APIConfig) GetRequireAgeVerificationBetaTag() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireAgeVerificationBetaTag
}

// GetRequireAgeVerificationBetaTagOk returns a tuple with the RequireAgeVerificationBetaTag field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetRequireAgeVerificationBetaTagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireAgeVerificationBetaTag, true
}

// SetRequireAgeVerificationBetaTag sets field value
func (o *APIConfig) SetRequireAgeVerificationBetaTag(v bool) {
	o.RequireAgeVerificationBetaTag = v
}

// GetSdkDeveloperFaqUrl returns the SdkDeveloperFaqUrl field value
func (o *APIConfig) GetSdkDeveloperFaqUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SdkDeveloperFaqUrl
}

// GetSdkDeveloperFaqUrlOk returns a tuple with the SdkDeveloperFaqUrl field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetSdkDeveloperFaqUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SdkDeveloperFaqUrl, true
}

// SetSdkDeveloperFaqUrl sets field value
func (o *APIConfig) SetSdkDeveloperFaqUrl(v string) {
	o.SdkDeveloperFaqUrl = v
}

// GetSdkDiscordUrl returns the SdkDiscordUrl field value
func (o *APIConfig) GetSdkDiscordUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SdkDiscordUrl
}

// GetSdkDiscordUrlOk returns a tuple with the SdkDiscordUrl field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetSdkDiscordUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SdkDiscordUrl, true
}

// SetSdkDiscordUrl sets field value
func (o *APIConfig) SetSdkDiscordUrl(v string) {
	o.SdkDiscordUrl = v
}

// GetSdkNotAllowedToPublishMessage returns the SdkNotAllowedToPublishMessage field value
func (o *APIConfig) GetSdkNotAllowedToPublishMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SdkNotAllowedToPublishMessage
}

// GetSdkNotAllowedToPublishMessageOk returns a tuple with the SdkNotAllowedToPublishMessage field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetSdkNotAllowedToPublishMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SdkNotAllowedToPublishMessage, true
}

// SetSdkNotAllowedToPublishMessage sets field value
func (o *APIConfig) SetSdkNotAllowedToPublishMessage(v string) {
	o.SdkNotAllowedToPublishMessage = v
}

// GetSdkUnityVersion returns the SdkUnityVersion field value
func (o *APIConfig) GetSdkUnityVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SdkUnityVersion
}

// GetSdkUnityVersionOk returns a tuple with the SdkUnityVersion field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetSdkUnityVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SdkUnityVersion, true
}

// SetSdkUnityVersion sets field value
func (o *APIConfig) SetSdkUnityVersion(v string) {
	o.SdkUnityVersion = v
}

// GetStringHostUrlList returns the StringHostUrlList field value
func (o *APIConfig) GetStringHostUrlList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.StringHostUrlList
}

// GetStringHostUrlListOk returns a tuple with the StringHostUrlList field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetStringHostUrlListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StringHostUrlList, true
}

// SetStringHostUrlList sets field value
func (o *APIConfig) SetStringHostUrlList(v []string) {
	o.StringHostUrlList = v
}

// GetSupportEmail returns the SupportEmail field value
func (o *APIConfig) GetSupportEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportEmail
}

// GetSupportEmailOk returns a tuple with the SupportEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetSupportEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportEmail, true
}

// SetSupportEmail sets field value
func (o *APIConfig) SetSupportEmail(v string) {
	o.SupportEmail = v
}

// GetSupportFormUrl returns the SupportFormUrl field value
func (o *APIConfig) GetSupportFormUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportFormUrl
}

// GetSupportFormUrlOk returns a tuple with the SupportFormUrl field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetSupportFormUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportFormUrl, true
}

// SetSupportFormUrl sets field value
func (o *APIConfig) SetSupportFormUrl(v string) {
	o.SupportFormUrl = v
}

// GetTimeOutWorldId returns the TimeOutWorldId field value
func (o *APIConfig) GetTimeOutWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeOutWorldId
}

// GetTimeOutWorldIdOk returns a tuple with the TimeOutWorldId field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetTimeOutWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeOutWorldId, true
}

// SetTimeOutWorldId sets field value
func (o *APIConfig) SetTimeOutWorldId(v string) {
	o.TimeOutWorldId = v
}

// GetTimekeeping returns the Timekeeping field value
func (o *APIConfig) GetTimekeeping() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Timekeeping
}

// GetTimekeepingOk returns a tuple with the Timekeeping field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetTimekeepingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timekeeping, true
}

// SetTimekeeping sets field value
func (o *APIConfig) SetTimekeeping(v bool) {
	o.Timekeeping = v
}

// GetTutorialWorldId returns the TutorialWorldId field value
func (o *APIConfig) GetTutorialWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TutorialWorldId
}

// GetTutorialWorldIdOk returns a tuple with the TutorialWorldId field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetTutorialWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TutorialWorldId, true
}

// SetTutorialWorldId sets field value
func (o *APIConfig) SetTutorialWorldId(v string) {
	o.TutorialWorldId = v
}

// GetUpdateRateMsMaximum returns the UpdateRateMsMaximum field value
func (o *APIConfig) GetUpdateRateMsMaximum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdateRateMsMaximum
}

// GetUpdateRateMsMaximumOk returns a tuple with the UpdateRateMsMaximum field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUpdateRateMsMaximumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateRateMsMaximum, true
}

// SetUpdateRateMsMaximum sets field value
func (o *APIConfig) SetUpdateRateMsMaximum(v int32) {
	o.UpdateRateMsMaximum = v
}

// GetUpdateRateMsMinimum returns the UpdateRateMsMinimum field value
func (o *APIConfig) GetUpdateRateMsMinimum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdateRateMsMinimum
}

// GetUpdateRateMsMinimumOk returns a tuple with the UpdateRateMsMinimum field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUpdateRateMsMinimumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateRateMsMinimum, true
}

// SetUpdateRateMsMinimum sets field value
func (o *APIConfig) SetUpdateRateMsMinimum(v int32) {
	o.UpdateRateMsMinimum = v
}

// GetUpdateRateMsNormal returns the UpdateRateMsNormal field value
func (o *APIConfig) GetUpdateRateMsNormal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdateRateMsNormal
}

// GetUpdateRateMsNormalOk returns a tuple with the UpdateRateMsNormal field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUpdateRateMsNormalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateRateMsNormal, true
}

// SetUpdateRateMsNormal sets field value
func (o *APIConfig) SetUpdateRateMsNormal(v int32) {
	o.UpdateRateMsNormal = v
}

// GetUpdateRateMsUdonManual returns the UpdateRateMsUdonManual field value
func (o *APIConfig) GetUpdateRateMsUdonManual() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdateRateMsUdonManual
}

// GetUpdateRateMsUdonManualOk returns a tuple with the UpdateRateMsUdonManual field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUpdateRateMsUdonManualOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateRateMsUdonManual, true
}

// SetUpdateRateMsUdonManual sets field value
func (o *APIConfig) SetUpdateRateMsUdonManual(v int32) {
	o.UpdateRateMsUdonManual = v
}

// GetUploadAnalysisPercent returns the UploadAnalysisPercent field value
func (o *APIConfig) GetUploadAnalysisPercent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UploadAnalysisPercent
}

// GetUploadAnalysisPercentOk returns a tuple with the UploadAnalysisPercent field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUploadAnalysisPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadAnalysisPercent, true
}

// SetUploadAnalysisPercent sets field value
func (o *APIConfig) SetUploadAnalysisPercent(v int32) {
	o.UploadAnalysisPercent = v
}

// GetUrlList returns the UrlList field value
func (o *APIConfig) GetUrlList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UrlList
}

// GetUrlListOk returns a tuple with the UrlList field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUrlListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UrlList, true
}

// SetUrlList sets field value
func (o *APIConfig) SetUrlList(v []string) {
	o.UrlList = v
}

// GetUseReliableUdpForVoice returns the UseReliableUdpForVoice field value
func (o *APIConfig) GetUseReliableUdpForVoice() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseReliableUdpForVoice
}

// GetUseReliableUdpForVoiceOk returns a tuple with the UseReliableUdpForVoice field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUseReliableUdpForVoiceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseReliableUdpForVoice, true
}

// SetUseReliableUdpForVoice sets field value
func (o *APIConfig) SetUseReliableUdpForVoice(v bool) {
	o.UseReliableUdpForVoice = v
}

// GetViveWindowsUrl returns the ViveWindowsUrl field value
func (o *APIConfig) GetViveWindowsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ViveWindowsUrl
}

// GetViveWindowsUrlOk returns a tuple with the ViveWindowsUrl field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetViveWindowsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViveWindowsUrl, true
}

// SetViveWindowsUrl sets field value
func (o *APIConfig) SetViveWindowsUrl(v string) {
	o.ViveWindowsUrl = v
}

// GetWebsocketMaxFriendsRefreshDelay returns the WebsocketMaxFriendsRefreshDelay field value
func (o *APIConfig) GetWebsocketMaxFriendsRefreshDelay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WebsocketMaxFriendsRefreshDelay
}

// GetWebsocketMaxFriendsRefreshDelayOk returns a tuple with the WebsocketMaxFriendsRefreshDelay field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetWebsocketMaxFriendsRefreshDelayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebsocketMaxFriendsRefreshDelay, true
}

// SetWebsocketMaxFriendsRefreshDelay sets field value
func (o *APIConfig) SetWebsocketMaxFriendsRefreshDelay(v int32) {
	o.WebsocketMaxFriendsRefreshDelay = v
}

// GetWebsocketQuickReconnectTime returns the WebsocketQuickReconnectTime field value
func (o *APIConfig) GetWebsocketQuickReconnectTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WebsocketQuickReconnectTime
}

// GetWebsocketQuickReconnectTimeOk returns a tuple with the WebsocketQuickReconnectTime field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetWebsocketQuickReconnectTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebsocketQuickReconnectTime, true
}

// SetWebsocketQuickReconnectTime sets field value
func (o *APIConfig) SetWebsocketQuickReconnectTime(v int32) {
	o.WebsocketQuickReconnectTime = v
}

// GetWebsocketReconnectMaxDelay returns the WebsocketReconnectMaxDelay field value
func (o *APIConfig) GetWebsocketReconnectMaxDelay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WebsocketReconnectMaxDelay
}

// GetWebsocketReconnectMaxDelayOk returns a tuple with the WebsocketReconnectMaxDelay field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetWebsocketReconnectMaxDelayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebsocketReconnectMaxDelay, true
}

// SetWebsocketReconnectMaxDelay sets field value
func (o *APIConfig) SetWebsocketReconnectMaxDelay(v int32) {
	o.WebsocketReconnectMaxDelay = v
}

// GetWhiteListedAssetUrls returns the WhiteListedAssetUrls field value
func (o *APIConfig) GetWhiteListedAssetUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WhiteListedAssetUrls
}

// GetWhiteListedAssetUrlsOk returns a tuple with the WhiteListedAssetUrls field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetWhiteListedAssetUrlsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WhiteListedAssetUrls, true
}

// SetWhiteListedAssetUrls sets field value
func (o *APIConfig) SetWhiteListedAssetUrls(v []string) {
	o.WhiteListedAssetUrls = v
}

func (o APIConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CampaignStatus"] = o.CampaignStatus
	toSerialize["DisableBackgroundPreloads"] = o.DisableBackgroundPreloads
	toSerialize["LocationGiftingNonSubPrioEnabled"] = o.LocationGiftingNonSubPrioEnabled
	toSerialize["VoiceEnableDegradation"] = o.VoiceEnableDegradation
	toSerialize["VoiceEnableReceiverLimiting"] = o.VoiceEnableReceiverLimiting
	toSerialize["accessLogsUrls"] = o.AccessLogsUrls
	toSerialize["address"] = o.Address
	toSerialize["ageVerificationInviteVisible"] = o.AgeVerificationInviteVisible
	toSerialize["ageVerificationP"] = o.AgeVerificationP
	toSerialize["ageVerificationStatusVisible"] = o.AgeVerificationStatusVisible
	toSerialize["analysisMaxRetries"] = o.AnalysisMaxRetries
	toSerialize["analysisRetryInterval"] = o.AnalysisRetryInterval
	toSerialize["analyticsSegment_NewUI_PctOfUsers"] = o.AnalyticsSegmentNewUIPctOfUsers
	toSerialize["analyticsSegment_NewUI_Salt"] = o.AnalyticsSegmentNewUISalt
	toSerialize["announcements"] = o.Announcements
	if !IsNil(o.AudioConfig) {
		toSerialize["audioConfig"] = o.AudioConfig
	}
	toSerialize["availableLanguageCodes"] = o.AvailableLanguageCodes
	toSerialize["availableLanguages"] = o.AvailableLanguages
	toSerialize["avatarPerfLimiter"] = o.AvatarPerfLimiter
	toSerialize["chatboxLogBufferSeconds"] = o.ChatboxLogBufferSeconds
	toSerialize["clientApiKey"] = o.ClientApiKey
	toSerialize["clientBPSCeiling"] = o.ClientBPSCeiling
	toSerialize["clientDisconnectTimeout"] = o.ClientDisconnectTimeout
	if !IsNil(o.ClientNetDispatchThread) {
		toSerialize["clientNetDispatchThread"] = o.ClientNetDispatchThread
	}
	toSerialize["clientNetDispatchThreadMobile"] = o.ClientNetDispatchThreadMobile
	if !IsNil(o.ClientNetInThread) {
		toSerialize["clientNetInThread"] = o.ClientNetInThread
	}
	if !IsNil(o.ClientNetInThread2) {
		toSerialize["clientNetInThread2"] = o.ClientNetInThread2
	}
	if !IsNil(o.ClientNetInThreadMobile) {
		toSerialize["clientNetInThreadMobile"] = o.ClientNetInThreadMobile
	}
	if !IsNil(o.ClientNetInThreadMobile2) {
		toSerialize["clientNetInThreadMobile2"] = o.ClientNetInThreadMobile2
	}
	if !IsNil(o.ClientNetOutThread) {
		toSerialize["clientNetOutThread"] = o.ClientNetOutThread
	}
	if !IsNil(o.ClientNetOutThread2) {
		toSerialize["clientNetOutThread2"] = o.ClientNetOutThread2
	}
	if !IsNil(o.ClientNetOutThreadMobile) {
		toSerialize["clientNetOutThreadMobile"] = o.ClientNetOutThreadMobile
	}
	if !IsNil(o.ClientNetOutThreadMobile2) {
		toSerialize["clientNetOutThreadMobile2"] = o.ClientNetOutThreadMobile2
	}
	if !IsNil(o.ClientQR) {
		toSerialize["clientQR"] = o.ClientQR
	}
	toSerialize["clientReservedPlayerBPS"] = o.ClientReservedPlayerBPS
	toSerialize["clientSentCountAllowance"] = o.ClientSentCountAllowance
	toSerialize["constants"] = o.Constants
	toSerialize["contactEmail"] = o.ContactEmail
	toSerialize["copyrightEmail"] = o.CopyrightEmail
	toSerialize["copyrightFormUrl"] = o.CopyrightFormUrl
	toSerialize["currentPrivacyVersion"] = o.CurrentPrivacyVersion
	toSerialize["currentTOSVersion"] = o.CurrentTOSVersion
	toSerialize["defaultAvatar"] = o.DefaultAvatar
	toSerialize["defaultStickerSet"] = o.DefaultStickerSet
	if !IsNil(o.DevLanguageCodes) {
		toSerialize["devLanguageCodes"] = o.DevLanguageCodes
	}
	toSerialize["devSdkUrl"] = o.DevSdkUrl
	toSerialize["devSdkVersion"] = o.DevSdkVersion
	toSerialize["dis-countdown"] = o.DisCountdown
	if !IsNil(o.DisableAVProInProton) {
		toSerialize["disableAVProInProton"] = o.DisableAVProInProton
	}
	toSerialize["disableAvatarCopying"] = o.DisableAvatarCopying
	toSerialize["disableAvatarGating"] = o.DisableAvatarGating
	if !IsNil(o.DisableCaptcha) {
		toSerialize["disableCaptcha"] = o.DisableCaptcha
	}
	toSerialize["disableCommunityLabs"] = o.DisableCommunityLabs
	toSerialize["disableCommunityLabsPromotion"] = o.DisableCommunityLabsPromotion
	toSerialize["disableEmail"] = o.DisableEmail
	toSerialize["disableEventStream"] = o.DisableEventStream
	toSerialize["disableFeedbackGating"] = o.DisableFeedbackGating
	toSerialize["disableFrontendBuilds"] = o.DisableFrontendBuilds
	toSerialize["disableGiftDrops"] = o.DisableGiftDrops
	toSerialize["disableHello"] = o.DisableHello
	toSerialize["disableOculusSubs"] = o.DisableOculusSubs
	toSerialize["disableRegistration"] = o.DisableRegistration
	toSerialize["disableSteamNetworking"] = o.DisableSteamNetworking
	toSerialize["disableTwoFactorAuth"] = o.DisableTwoFactorAuth
	toSerialize["disableUdon"] = o.DisableUdon
	toSerialize["disableUpgradeAccount"] = o.DisableUpgradeAccount
	toSerialize["downloadLinkWindows"] = o.DownloadLinkWindows
	toSerialize["downloadUrls"] = o.DownloadUrls
	toSerialize["dynamicWorldRows"] = o.DynamicWorldRows
	toSerialize["economyLedgerBackfill"] = o.EconomyLedgerBackfill
	toSerialize["economyLedgerMigrationStop"] = o.EconomyLedgerMigrationStop
	toSerialize["economyLedgerMode"] = o.EconomyLedgerMode
	toSerialize["economyPauseEnd"] = o.EconomyPauseEnd
	toSerialize["economyPauseStart"] = o.EconomyPauseStart
	toSerialize["economyPurchaseRepairEnabled"] = o.EconomyPurchaseRepairEnabled
	toSerialize["economyState"] = o.EconomyState
	toSerialize["events"] = o.Events
	toSerialize["forceUseLatestWorld"] = o.ForceUseLatestWorld
	toSerialize["giftDisplayType"] = o.GiftDisplayType
	toSerialize["googleApiClientId"] = o.GoogleApiClientId
	toSerialize["homeWorldId"] = o.HomeWorldId
	toSerialize["homepageRedirectTarget"] = o.HomepageRedirectTarget
	toSerialize["hubWorldId"] = o.HubWorldId
	toSerialize["imageHostUrlList"] = o.ImageHostUrlList
	toSerialize["iosAppVersion"] = o.IosAppVersion
	toSerialize["iosVersion"] = o.IosVersion
	toSerialize["jobsEmail"] = o.JobsEmail
	toSerialize["maxUserEmoji"] = o.MaxUserEmoji
	toSerialize["maxUserStickers"] = o.MaxUserStickers
	toSerialize["minSupportedClientBuildNumber"] = o.MinSupportedClientBuildNumber
	toSerialize["minimumUnityVersionForUploads"] = o.MinimumUnityVersionForUploads
	toSerialize["moderationEmail"] = o.ModerationEmail
	toSerialize["notAllowedToSelectAvatarInPrivateWorldMessage"] = o.NotAllowedToSelectAvatarInPrivateWorldMessage
	toSerialize["offlineAnalysis"] = o.OfflineAnalysis
	toSerialize["photonNameserverOverrides"] = o.PhotonNameserverOverrides
	toSerialize["photonPublicKeys"] = o.PhotonPublicKeys
	toSerialize["player-url-resolver-sha1"] = o.PlayerUrlResolverSha1
	toSerialize["player-url-resolver-version"] = o.PlayerUrlResolverVersion
	toSerialize["publicKey"] = o.PublicKey
	toSerialize["reportCategories"] = o.ReportCategories
	toSerialize["reportFormUrl"] = o.ReportFormUrl
	toSerialize["reportOptions"] = o.ReportOptions
	toSerialize["reportReasons"] = o.ReportReasons
	toSerialize["requireAgeVerificationBetaTag"] = o.RequireAgeVerificationBetaTag
	toSerialize["sdkDeveloperFaqUrl"] = o.SdkDeveloperFaqUrl
	toSerialize["sdkDiscordUrl"] = o.SdkDiscordUrl
	toSerialize["sdkNotAllowedToPublishMessage"] = o.SdkNotAllowedToPublishMessage
	toSerialize["sdkUnityVersion"] = o.SdkUnityVersion
	toSerialize["stringHostUrlList"] = o.StringHostUrlList
	toSerialize["supportEmail"] = o.SupportEmail
	toSerialize["supportFormUrl"] = o.SupportFormUrl
	toSerialize["timeOutWorldId"] = o.TimeOutWorldId
	toSerialize["timekeeping"] = o.Timekeeping
	toSerialize["tutorialWorldId"] = o.TutorialWorldId
	toSerialize["updateRateMsMaximum"] = o.UpdateRateMsMaximum
	toSerialize["updateRateMsMinimum"] = o.UpdateRateMsMinimum
	toSerialize["updateRateMsNormal"] = o.UpdateRateMsNormal
	toSerialize["updateRateMsUdonManual"] = o.UpdateRateMsUdonManual
	toSerialize["uploadAnalysisPercent"] = o.UploadAnalysisPercent
	toSerialize["urlList"] = o.UrlList
	toSerialize["useReliableUdpForVoice"] = o.UseReliableUdpForVoice
	toSerialize["viveWindowsUrl"] = o.ViveWindowsUrl
	toSerialize["websocketMaxFriendsRefreshDelay"] = o.WebsocketMaxFriendsRefreshDelay
	toSerialize["websocketQuickReconnectTime"] = o.WebsocketQuickReconnectTime
	toSerialize["websocketReconnectMaxDelay"] = o.WebsocketReconnectMaxDelay
	toSerialize["whiteListedAssetUrls"] = o.WhiteListedAssetUrls
	return toSerialize, nil
}

func (o *APIConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CampaignStatus",
		"DisableBackgroundPreloads",
		"LocationGiftingNonSubPrioEnabled",
		"VoiceEnableDegradation",
		"VoiceEnableReceiverLimiting",
		"accessLogsUrls",
		"address",
		"ageVerificationInviteVisible",
		"ageVerificationP",
		"ageVerificationStatusVisible",
		"analysisMaxRetries",
		"analysisRetryInterval",
		"analyticsSegment_NewUI_PctOfUsers",
		"analyticsSegment_NewUI_Salt",
		"announcements",
		"availableLanguageCodes",
		"availableLanguages",
		"avatarPerfLimiter",
		"chatboxLogBufferSeconds",
		"clientApiKey",
		"clientBPSCeiling",
		"clientDisconnectTimeout",
		"clientNetDispatchThreadMobile",
		"clientReservedPlayerBPS",
		"clientSentCountAllowance",
		"constants",
		"contactEmail",
		"copyrightEmail",
		"copyrightFormUrl",
		"currentPrivacyVersion",
		"currentTOSVersion",
		"defaultAvatar",
		"defaultStickerSet",
		"devSdkUrl",
		"devSdkVersion",
		"dis-countdown",
		"disableAvatarCopying",
		"disableAvatarGating",
		"disableCommunityLabs",
		"disableCommunityLabsPromotion",
		"disableEmail",
		"disableEventStream",
		"disableFeedbackGating",
		"disableFrontendBuilds",
		"disableGiftDrops",
		"disableHello",
		"disableOculusSubs",
		"disableRegistration",
		"disableSteamNetworking",
		"disableTwoFactorAuth",
		"disableUdon",
		"disableUpgradeAccount",
		"downloadLinkWindows",
		"downloadUrls",
		"dynamicWorldRows",
		"economyLedgerBackfill",
		"economyLedgerMigrationStop",
		"economyLedgerMode",
		"economyPauseEnd",
		"economyPauseStart",
		"economyPurchaseRepairEnabled",
		"economyState",
		"events",
		"forceUseLatestWorld",
		"giftDisplayType",
		"googleApiClientId",
		"homeWorldId",
		"homepageRedirectTarget",
		"hubWorldId",
		"imageHostUrlList",
		"iosAppVersion",
		"iosVersion",
		"jobsEmail",
		"maxUserEmoji",
		"maxUserStickers",
		"minSupportedClientBuildNumber",
		"minimumUnityVersionForUploads",
		"moderationEmail",
		"notAllowedToSelectAvatarInPrivateWorldMessage",
		"offlineAnalysis",
		"photonNameserverOverrides",
		"photonPublicKeys",
		"player-url-resolver-sha1",
		"player-url-resolver-version",
		"publicKey",
		"reportCategories",
		"reportFormUrl",
		"reportOptions",
		"reportReasons",
		"requireAgeVerificationBetaTag",
		"sdkDeveloperFaqUrl",
		"sdkDiscordUrl",
		"sdkNotAllowedToPublishMessage",
		"sdkUnityVersion",
		"stringHostUrlList",
		"supportEmail",
		"supportFormUrl",
		"timeOutWorldId",
		"timekeeping",
		"tutorialWorldId",
		"updateRateMsMaximum",
		"updateRateMsMinimum",
		"updateRateMsNormal",
		"updateRateMsUdonManual",
		"uploadAnalysisPercent",
		"urlList",
		"useReliableUdpForVoice",
		"viveWindowsUrl",
		"websocketMaxFriendsRefreshDelay",
		"websocketQuickReconnectTime",
		"websocketReconnectMaxDelay",
		"whiteListedAssetUrls",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAPIConfig := _APIConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAPIConfig)

	if err != nil {
		return err
	}

	*o = APIConfig(varAPIConfig)

	return err
}

type NullableAPIConfig struct {
	value *APIConfig
	isSet bool
}

func (v NullableAPIConfig) Get() *APIConfig {
	return v.value
}

func (v *NullableAPIConfig) Set(val *APIConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIConfig(val *APIConfig) *NullableAPIConfig {
	return &NullableAPIConfig{value: val, isSet: true}
}

func (v NullableAPIConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


