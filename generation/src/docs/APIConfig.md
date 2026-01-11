# APIConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignStatus** | **string** | The current platform-wide event taking place | 
**DisableBackgroundPreloads** | **bool** | Toggles if certain assets are preloaded in the background | [default to true]
**LocationGiftingNonSubPrioEnabled** | **bool** | Toggles whether users without a current VRC+ subscription are priority recipients for gift drops | [default to true]
**VoiceEnableDegradation** | **bool** | Unknown, probably voice optimization testing | [default to false]
**VoiceEnableReceiverLimiting** | **bool** | Unknown, probably voice optimization testing | [default to true]
**AccessLogsUrls** | [**APIConfigAccessLogsUrls**](APIConfigAccessLogsUrls.md) |  | 
**Address** | **string** | VRChat&#39;s office address | 
**AgeVerificationInviteVisible** | **bool** |  | 
**AgeVerificationP** | **bool** |  | 
**AgeVerificationStatusVisible** | **bool** |  | 
**AnalysisMaxRetries** | **int32** | Max retries for avatar analysis requests | 
**AnalysisRetryInterval** | **int32** | Interval between retries for avatar analysis requests | 
**AnalyticsSegmentNewUIPctOfUsers** | **int32** | Unknown | 
**AnalyticsSegmentNewUISalt** | **string** | Unknown | 
**Announcements** | [**[]APIConfigAnnouncement**](APIConfigAnnouncement.md) | Public Announcements | 
**AudioConfig** | Pointer to [**APIConfigAudioConfig**](APIConfigAudioConfig.md) |  | [optional] 
**AvailableLanguageCodes** | **[]string** | List of supported Languages | 
**AvailableLanguages** | **[]string** | List of supported Languages | 
**AvatarPerfLimiter** | [**APIConfigAvatarPerfLimiter**](APIConfigAvatarPerfLimiter.md) |  | 
**ChatboxLogBufferSeconds** | **int32** | Unknown | [default to 40]
**ClientApiKey** | **string** | apiKey to be used for all other requests | 
**ClientBPSCeiling** | **int32** | Unknown | [default to 18432]
**ClientDisconnectTimeout** | **int32** | Unknown | [default to 30000]
**ClientNetDispatchThread** | Pointer to **bool** | Unknown | [optional] [default to false]
**ClientNetDispatchThreadMobile** | **bool** | Unknown | [default to true]
**ClientNetInThread** | Pointer to **bool** | Unknown | [optional] [default to false]
**ClientNetInThread2** | Pointer to **bool** | Unknown | [optional] [default to false]
**ClientNetInThreadMobile** | Pointer to **bool** | Unknown | [optional] [default to false]
**ClientNetInThreadMobile2** | Pointer to **bool** | Unknown | [optional] [default to false]
**ClientNetOutThread** | Pointer to **bool** | Unknown | [optional] [default to false]
**ClientNetOutThread2** | Pointer to **bool** | Unknown | [optional] [default to false]
**ClientNetOutThreadMobile** | Pointer to **bool** | Unknown | [optional] [default to false]
**ClientNetOutThreadMobile2** | Pointer to **bool** | Unknown | [optional] [default to false]
**ClientQR** | Pointer to **int32** | Unknown | [optional] [default to 1]
**ClientReservedPlayerBPS** | **int32** | Unknown | [default to 7168]
**ClientSentCountAllowance** | **int32** | Unknown | [default to 15]
**Constants** | [**APIConfigConstants**](APIConfigConstants.md) |  | 
**ContactEmail** | **string** | VRChat&#39;s contact email | 
**CopyrightEmail** | **string** | VRChat&#39;s copyright-issues-related email | 
**CopyrightFormUrl** | **string** | VRChat&#39;s DMCA claim webform url | 
**CurrentPrivacyVersion** | **int32** | Current version number of the Privacy Agreement | [default to 1]
**CurrentTOSVersion** | **int32** | Current version number of the Terms of Service | 
**DefaultAvatar** | **string** |  | 
**DefaultStickerSet** | **string** |  | 
**DevLanguageCodes** | Pointer to **[]string** | Unknown | [optional] 
**DevSdkUrl** | **string** | Link to download the development SDK, use downloadUrls instead | 
**DevSdkVersion** | **string** | Version of the development SDK | 
**DisCountdown** | **time.Time** | Unknown, \&quot;dis\&quot; maybe for disconnect? | 
**DisableAVProInProton** | Pointer to **bool** | Unknown | [optional] [default to false]
**DisableAvatarCopying** | **bool** | Toggles if copying avatars should be disabled | [default to false]
**DisableAvatarGating** | **bool** | Toggles if avatar gating should be disabled. Avatar gating restricts uploading of avatars to people with the &#x60;system_avatar_access&#x60; Tag or &#x60;admin_avatar_access&#x60; Tag | [default to false]
**DisableCaptcha** | Pointer to **bool** | Unknown | [optional] [default to true]
**DisableCommunityLabs** | **bool** | Toggles if the Community Labs should be disabled | [default to false]
**DisableCommunityLabsPromotion** | **bool** | Toggles if promotion out of Community Labs should be disabled | [default to false]
**DisableEmail** | **bool** | Unknown | [default to false]
**DisableEventStream** | **bool** | Toggles if Analytics should be disabled. | [default to false]
**DisableFeedbackGating** | **bool** | Toggles if feedback gating should be disabled. Feedback gating restricts submission of feedback (reporting a World or User) to people with the &#x60;system_feedback_access&#x60; Tag. | [default to false]
**DisableFrontendBuilds** | **bool** | Unknown, probably toggles compilation of frontend web builds? So internal flag? | [default to false]
**DisableGiftDrops** | **bool** | Toggles if gift drops should be disabled | [default to false]
**DisableHello** | **bool** | Unknown | [default to false]
**DisableOculusSubs** | **bool** | Toggles if signing up for Subscriptions in Oculus is disabled or not. | [default to false]
**DisableRegistration** | **bool** | Toggles if new user account registration should be disabled. | [default to false]
**DisableSteamNetworking** | **bool** | Toggles if Steam Networking should be disabled. VRChat these days uses Photon Unity Networking (PUN) instead. | [default to true]
**DisableTwoFactorAuth** | **bool** | Toggles if 2FA should be disabled. | [default to false]
**DisableUdon** | **bool** | Toggles if Udon should be universally disabled in-game. | [default to false]
**DisableUpgradeAccount** | **bool** | Toggles if account upgrading \&quot;linking with Steam/Oculus\&quot; should be disabled. | [default to false]
**DownloadLinkWindows** | **string** | Download link for game on the Oculus Rift website. | 
**DownloadUrls** | [**APIConfigDownloadURLList**](APIConfigDownloadURLList.md) |  | 
**DynamicWorldRows** | [**[]DynamicContentRow**](DynamicContentRow.md) | Array of DynamicWorldRow objects, used by the game to display the list of world rows | 
**EconomyLedgerBackfill** | **bool** | Unknown | 
**EconomyLedgerMigrationStop** | **string** | Unknown | 
**EconomyLedgerMode** | **string** | Unknown | 
**EconomyPauseEnd** | **time.Time** | Unknown | 
**EconomyPauseStart** | **time.Time** | Unknown | 
**EconomyPurchaseRepairEnabled** | **bool** | Unknown | 
**EconomyState** | **int32** | Unknown | [default to 1]
**Events** | [**APIConfigEvents**](APIConfigEvents.md) |  | 
**ForceUseLatestWorld** | **bool** | Unknown | [default to true]
**GiftDisplayType** | **string** | Display type of gifts | 
**GoogleApiClientId** | **string** | Unknown | [default to "827942544393-r2ouvckvouldn9dg9uruseje575e878f.apps.googleusercontent.com"]
**HomeWorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**HomepageRedirectTarget** | **string** | Redirect target if you try to open the base API domain in your browser | [default to "https://hello.vrchat.com"]
**HubWorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**ImageHostUrlList** | **[]string** | A list of explicitly allowed origins that worlds can request images from via the Udon&#39;s [VRCImageDownloader#DownloadImage](https://creators.vrchat.com/worlds/udon/image-loading/#downloadimage). | 
**IosAppVersion** | **[]string** | Current app version for iOS | 
**IosVersion** | [**APIConfigIosVersion**](APIConfigIosVersion.md) |  | 
**JobsEmail** | **string** | VRChat&#39;s job application email | 
**MaxUserEmoji** | **int32** | The maximum number of custom emoji each user may have at a given time. | [default to 18]
**MaxUserStickers** | **int32** | The maximum number of custom stickers each user may have at a given time. | [default to 18]
**MinSupportedClientBuildNumber** | [**APIConfigMinSupportedClientBuildNumber**](APIConfigMinSupportedClientBuildNumber.md) |  | 
**MinimumUnityVersionForUploads** | **string** | Minimum Unity version required for uploading assets | [default to "2019.0.0f1"]
**ModerationEmail** | **string** | VRChat&#39;s moderation related email | 
**NotAllowedToSelectAvatarInPrivateWorldMessage** | **string** | Used in-game to notify a user they aren&#39;t allowed to select avatars in private worlds | 
**OfflineAnalysis** | [**APIConfigOfflineAnalysis**](APIConfigOfflineAnalysis.md) |  | 
**PhotonNameserverOverrides** | **[]string** | Unknown | 
**PhotonPublicKeys** | **[]string** | Unknown | 
**PlayerUrlResolverSha1** | **string** | Currently used youtube-dl.exe hash in SHA1-delimited format | 
**PlayerUrlResolverVersion** | **string** | Currently used youtube-dl.exe version | 
**PublicKey** | **string** | Public key, hex encoded | 
**ReportCategories** | [**map[string]ReportCategory**](ReportCategory.md) | Categories available for reporting objectionable content | 
**ReportFormUrl** | **string** | URL to the report form | [default to "https://help.vrchat.com/hc/en-us/requests/new?ticket_form_id=1500000182242&tf_360056455174=user_report&tf_360057451993={userId}&tf_1500001445142={reportedId}&tf_subject={reason} {category} By {contentType} {reportedName}&tf_description={description}"]
**ReportOptions** | [**map[string]map[string][]string**](map.md) | Options for reporting content. Select a key+value from this mapping as the &#x60;type&#x60; of the report. Select one key+value from the object at reportOptions[type] as the &#x60;category&#x60; of the report. reportCategories[category] contains user-facing text to display for all possible categories. Select one value from the array at reportOptions[type][category] as the &#x60;reason&#x60; of the report. reportReasons[reason] contains user-facing text to display for all possible categories. | 
**ReportReasons** | [**map[string]ReportReason**](ReportReason.md) | Reasons available for submitting a report | 
**RequireAgeVerificationBetaTag** | **bool** |  | 
**SdkDeveloperFaqUrl** | **string** | Link to the developer FAQ | 
**SdkDiscordUrl** | **string** | Link to the official VRChat Discord | 
**SdkNotAllowedToPublishMessage** | **string** | Used in the SDK to notify a user they aren&#39;t allowed to upload avatars/worlds yet | 
**SdkUnityVersion** | **string** | Unity version supported by the SDK | 
**StringHostUrlList** | **[]string** | A list of explicitly allowed origins that worlds can request strings from via the Udon&#39;s [VRCStringDownloader.LoadUrl](https://creators.vrchat.com/worlds/udon/string-loading/#ivrcstringdownload). | 
**SupportEmail** | **string** | VRChat&#39;s support email | 
**SupportFormUrl** | **string** | VRChat&#39;s support form | 
**TimeOutWorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**Timekeeping** | **bool** | Unknown | [default to true]
**TutorialWorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**UpdateRateMsMaximum** | **int32** | Unknown | 
**UpdateRateMsMinimum** | **int32** | Unknown | 
**UpdateRateMsNormal** | **int32** | Unknown | 
**UpdateRateMsUdonManual** | **int32** | Unknown | 
**UploadAnalysisPercent** | **int32** | Unknown | 
**UrlList** | **[]string** | List of allowed URLs that bypass the \&quot;Allow untrusted URL&#39;s\&quot; setting in-game | 
**UseReliableUdpForVoice** | **bool** | Unknown | [default to false]
**ViveWindowsUrl** | **string** | Download link for game on the Steam website. | 
**WebsocketMaxFriendsRefreshDelay** | **int32** | Unknown | [default to 900]
**WebsocketQuickReconnectTime** | **int32** | Unknown | [default to 2]
**WebsocketReconnectMaxDelay** | **int32** | Unknown | [default to 2]
**WhiteListedAssetUrls** | **[]string** | List of allowed URLs that are allowed to host avatar assets | 

## Methods

### NewAPIConfig

`func NewAPIConfig(campaignStatus string, disableBackgroundPreloads bool, locationGiftingNonSubPrioEnabled bool, voiceEnableDegradation bool, voiceEnableReceiverLimiting bool, accessLogsUrls APIConfigAccessLogsUrls, address string, ageVerificationInviteVisible bool, ageVerificationP bool, ageVerificationStatusVisible bool, analysisMaxRetries int32, analysisRetryInterval int32, analyticsSegmentNewUIPctOfUsers int32, analyticsSegmentNewUISalt string, announcements []APIConfigAnnouncement, availableLanguageCodes []string, availableLanguages []string, avatarPerfLimiter APIConfigAvatarPerfLimiter, chatboxLogBufferSeconds int32, clientApiKey string, clientBPSCeiling int32, clientDisconnectTimeout int32, clientNetDispatchThreadMobile bool, clientReservedPlayerBPS int32, clientSentCountAllowance int32, constants APIConfigConstants, contactEmail string, copyrightEmail string, copyrightFormUrl string, currentPrivacyVersion int32, currentTOSVersion int32, defaultAvatar string, defaultStickerSet string, devSdkUrl string, devSdkVersion string, disCountdown time.Time, disableAvatarCopying bool, disableAvatarGating bool, disableCommunityLabs bool, disableCommunityLabsPromotion bool, disableEmail bool, disableEventStream bool, disableFeedbackGating bool, disableFrontendBuilds bool, disableGiftDrops bool, disableHello bool, disableOculusSubs bool, disableRegistration bool, disableSteamNetworking bool, disableTwoFactorAuth bool, disableUdon bool, disableUpgradeAccount bool, downloadLinkWindows string, downloadUrls APIConfigDownloadURLList, dynamicWorldRows []DynamicContentRow, economyLedgerBackfill bool, economyLedgerMigrationStop string, economyLedgerMode string, economyPauseEnd time.Time, economyPauseStart time.Time, economyPurchaseRepairEnabled bool, economyState int32, events APIConfigEvents, forceUseLatestWorld bool, giftDisplayType string, googleApiClientId string, homeWorldId string, homepageRedirectTarget string, hubWorldId string, imageHostUrlList []string, iosAppVersion []string, iosVersion APIConfigIosVersion, jobsEmail string, maxUserEmoji int32, maxUserStickers int32, minSupportedClientBuildNumber APIConfigMinSupportedClientBuildNumber, minimumUnityVersionForUploads string, moderationEmail string, notAllowedToSelectAvatarInPrivateWorldMessage string, offlineAnalysis APIConfigOfflineAnalysis, photonNameserverOverrides []string, photonPublicKeys []string, playerUrlResolverSha1 string, playerUrlResolverVersion string, publicKey string, reportCategories map[string]ReportCategory, reportFormUrl string, reportOptions map[string]map[string][]string, reportReasons map[string]ReportReason, requireAgeVerificationBetaTag bool, sdkDeveloperFaqUrl string, sdkDiscordUrl string, sdkNotAllowedToPublishMessage string, sdkUnityVersion string, stringHostUrlList []string, supportEmail string, supportFormUrl string, timeOutWorldId string, timekeeping bool, tutorialWorldId string, updateRateMsMaximum int32, updateRateMsMinimum int32, updateRateMsNormal int32, updateRateMsUdonManual int32, uploadAnalysisPercent int32, urlList []string, useReliableUdpForVoice bool, viveWindowsUrl string, websocketMaxFriendsRefreshDelay int32, websocketQuickReconnectTime int32, websocketReconnectMaxDelay int32, whiteListedAssetUrls []string, ) *APIConfig`

NewAPIConfig instantiates a new APIConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIConfigWithDefaults

`func NewAPIConfigWithDefaults() *APIConfig`

NewAPIConfigWithDefaults instantiates a new APIConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignStatus

`func (o *APIConfig) GetCampaignStatus() string`

GetCampaignStatus returns the CampaignStatus field if non-nil, zero value otherwise.

### GetCampaignStatusOk

`func (o *APIConfig) GetCampaignStatusOk() (*string, bool)`

GetCampaignStatusOk returns a tuple with the CampaignStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignStatus

`func (o *APIConfig) SetCampaignStatus(v string)`

SetCampaignStatus sets CampaignStatus field to given value.


### GetDisableBackgroundPreloads

`func (o *APIConfig) GetDisableBackgroundPreloads() bool`

GetDisableBackgroundPreloads returns the DisableBackgroundPreloads field if non-nil, zero value otherwise.

### GetDisableBackgroundPreloadsOk

`func (o *APIConfig) GetDisableBackgroundPreloadsOk() (*bool, bool)`

GetDisableBackgroundPreloadsOk returns a tuple with the DisableBackgroundPreloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBackgroundPreloads

`func (o *APIConfig) SetDisableBackgroundPreloads(v bool)`

SetDisableBackgroundPreloads sets DisableBackgroundPreloads field to given value.


### GetLocationGiftingNonSubPrioEnabled

`func (o *APIConfig) GetLocationGiftingNonSubPrioEnabled() bool`

GetLocationGiftingNonSubPrioEnabled returns the LocationGiftingNonSubPrioEnabled field if non-nil, zero value otherwise.

### GetLocationGiftingNonSubPrioEnabledOk

`func (o *APIConfig) GetLocationGiftingNonSubPrioEnabledOk() (*bool, bool)`

GetLocationGiftingNonSubPrioEnabledOk returns a tuple with the LocationGiftingNonSubPrioEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationGiftingNonSubPrioEnabled

`func (o *APIConfig) SetLocationGiftingNonSubPrioEnabled(v bool)`

SetLocationGiftingNonSubPrioEnabled sets LocationGiftingNonSubPrioEnabled field to given value.


### GetVoiceEnableDegradation

`func (o *APIConfig) GetVoiceEnableDegradation() bool`

GetVoiceEnableDegradation returns the VoiceEnableDegradation field if non-nil, zero value otherwise.

### GetVoiceEnableDegradationOk

`func (o *APIConfig) GetVoiceEnableDegradationOk() (*bool, bool)`

GetVoiceEnableDegradationOk returns a tuple with the VoiceEnableDegradation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceEnableDegradation

`func (o *APIConfig) SetVoiceEnableDegradation(v bool)`

SetVoiceEnableDegradation sets VoiceEnableDegradation field to given value.


### GetVoiceEnableReceiverLimiting

`func (o *APIConfig) GetVoiceEnableReceiverLimiting() bool`

GetVoiceEnableReceiverLimiting returns the VoiceEnableReceiverLimiting field if non-nil, zero value otherwise.

### GetVoiceEnableReceiverLimitingOk

`func (o *APIConfig) GetVoiceEnableReceiverLimitingOk() (*bool, bool)`

GetVoiceEnableReceiverLimitingOk returns a tuple with the VoiceEnableReceiverLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceEnableReceiverLimiting

`func (o *APIConfig) SetVoiceEnableReceiverLimiting(v bool)`

SetVoiceEnableReceiverLimiting sets VoiceEnableReceiverLimiting field to given value.


### GetAccessLogsUrls

`func (o *APIConfig) GetAccessLogsUrls() APIConfigAccessLogsUrls`

GetAccessLogsUrls returns the AccessLogsUrls field if non-nil, zero value otherwise.

### GetAccessLogsUrlsOk

`func (o *APIConfig) GetAccessLogsUrlsOk() (*APIConfigAccessLogsUrls, bool)`

GetAccessLogsUrlsOk returns a tuple with the AccessLogsUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLogsUrls

`func (o *APIConfig) SetAccessLogsUrls(v APIConfigAccessLogsUrls)`

SetAccessLogsUrls sets AccessLogsUrls field to given value.


### GetAddress

`func (o *APIConfig) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *APIConfig) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *APIConfig) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAgeVerificationInviteVisible

`func (o *APIConfig) GetAgeVerificationInviteVisible() bool`

GetAgeVerificationInviteVisible returns the AgeVerificationInviteVisible field if non-nil, zero value otherwise.

### GetAgeVerificationInviteVisibleOk

`func (o *APIConfig) GetAgeVerificationInviteVisibleOk() (*bool, bool)`

GetAgeVerificationInviteVisibleOk returns a tuple with the AgeVerificationInviteVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeVerificationInviteVisible

`func (o *APIConfig) SetAgeVerificationInviteVisible(v bool)`

SetAgeVerificationInviteVisible sets AgeVerificationInviteVisible field to given value.


### GetAgeVerificationP

`func (o *APIConfig) GetAgeVerificationP() bool`

GetAgeVerificationP returns the AgeVerificationP field if non-nil, zero value otherwise.

### GetAgeVerificationPOk

`func (o *APIConfig) GetAgeVerificationPOk() (*bool, bool)`

GetAgeVerificationPOk returns a tuple with the AgeVerificationP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeVerificationP

`func (o *APIConfig) SetAgeVerificationP(v bool)`

SetAgeVerificationP sets AgeVerificationP field to given value.


### GetAgeVerificationStatusVisible

`func (o *APIConfig) GetAgeVerificationStatusVisible() bool`

GetAgeVerificationStatusVisible returns the AgeVerificationStatusVisible field if non-nil, zero value otherwise.

### GetAgeVerificationStatusVisibleOk

`func (o *APIConfig) GetAgeVerificationStatusVisibleOk() (*bool, bool)`

GetAgeVerificationStatusVisibleOk returns a tuple with the AgeVerificationStatusVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeVerificationStatusVisible

`func (o *APIConfig) SetAgeVerificationStatusVisible(v bool)`

SetAgeVerificationStatusVisible sets AgeVerificationStatusVisible field to given value.


### GetAnalysisMaxRetries

`func (o *APIConfig) GetAnalysisMaxRetries() int32`

GetAnalysisMaxRetries returns the AnalysisMaxRetries field if non-nil, zero value otherwise.

### GetAnalysisMaxRetriesOk

`func (o *APIConfig) GetAnalysisMaxRetriesOk() (*int32, bool)`

GetAnalysisMaxRetriesOk returns a tuple with the AnalysisMaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisMaxRetries

`func (o *APIConfig) SetAnalysisMaxRetries(v int32)`

SetAnalysisMaxRetries sets AnalysisMaxRetries field to given value.


### GetAnalysisRetryInterval

`func (o *APIConfig) GetAnalysisRetryInterval() int32`

GetAnalysisRetryInterval returns the AnalysisRetryInterval field if non-nil, zero value otherwise.

### GetAnalysisRetryIntervalOk

`func (o *APIConfig) GetAnalysisRetryIntervalOk() (*int32, bool)`

GetAnalysisRetryIntervalOk returns a tuple with the AnalysisRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisRetryInterval

`func (o *APIConfig) SetAnalysisRetryInterval(v int32)`

SetAnalysisRetryInterval sets AnalysisRetryInterval field to given value.


### GetAnalyticsSegmentNewUIPctOfUsers

`func (o *APIConfig) GetAnalyticsSegmentNewUIPctOfUsers() int32`

GetAnalyticsSegmentNewUIPctOfUsers returns the AnalyticsSegmentNewUIPctOfUsers field if non-nil, zero value otherwise.

### GetAnalyticsSegmentNewUIPctOfUsersOk

`func (o *APIConfig) GetAnalyticsSegmentNewUIPctOfUsersOk() (*int32, bool)`

GetAnalyticsSegmentNewUIPctOfUsersOk returns a tuple with the AnalyticsSegmentNewUIPctOfUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsSegmentNewUIPctOfUsers

`func (o *APIConfig) SetAnalyticsSegmentNewUIPctOfUsers(v int32)`

SetAnalyticsSegmentNewUIPctOfUsers sets AnalyticsSegmentNewUIPctOfUsers field to given value.


### GetAnalyticsSegmentNewUISalt

`func (o *APIConfig) GetAnalyticsSegmentNewUISalt() string`

GetAnalyticsSegmentNewUISalt returns the AnalyticsSegmentNewUISalt field if non-nil, zero value otherwise.

### GetAnalyticsSegmentNewUISaltOk

`func (o *APIConfig) GetAnalyticsSegmentNewUISaltOk() (*string, bool)`

GetAnalyticsSegmentNewUISaltOk returns a tuple with the AnalyticsSegmentNewUISalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsSegmentNewUISalt

`func (o *APIConfig) SetAnalyticsSegmentNewUISalt(v string)`

SetAnalyticsSegmentNewUISalt sets AnalyticsSegmentNewUISalt field to given value.


### GetAnnouncements

`func (o *APIConfig) GetAnnouncements() []APIConfigAnnouncement`

GetAnnouncements returns the Announcements field if non-nil, zero value otherwise.

### GetAnnouncementsOk

`func (o *APIConfig) GetAnnouncementsOk() (*[]APIConfigAnnouncement, bool)`

GetAnnouncementsOk returns a tuple with the Announcements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncements

`func (o *APIConfig) SetAnnouncements(v []APIConfigAnnouncement)`

SetAnnouncements sets Announcements field to given value.


### GetAudioConfig

`func (o *APIConfig) GetAudioConfig() APIConfigAudioConfig`

GetAudioConfig returns the AudioConfig field if non-nil, zero value otherwise.

### GetAudioConfigOk

`func (o *APIConfig) GetAudioConfigOk() (*APIConfigAudioConfig, bool)`

GetAudioConfigOk returns a tuple with the AudioConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioConfig

`func (o *APIConfig) SetAudioConfig(v APIConfigAudioConfig)`

SetAudioConfig sets AudioConfig field to given value.

### HasAudioConfig

`func (o *APIConfig) HasAudioConfig() bool`

HasAudioConfig returns a boolean if a field has been set.

### GetAvailableLanguageCodes

`func (o *APIConfig) GetAvailableLanguageCodes() []string`

GetAvailableLanguageCodes returns the AvailableLanguageCodes field if non-nil, zero value otherwise.

### GetAvailableLanguageCodesOk

`func (o *APIConfig) GetAvailableLanguageCodesOk() (*[]string, bool)`

GetAvailableLanguageCodesOk returns a tuple with the AvailableLanguageCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableLanguageCodes

`func (o *APIConfig) SetAvailableLanguageCodes(v []string)`

SetAvailableLanguageCodes sets AvailableLanguageCodes field to given value.


### GetAvailableLanguages

`func (o *APIConfig) GetAvailableLanguages() []string`

GetAvailableLanguages returns the AvailableLanguages field if non-nil, zero value otherwise.

### GetAvailableLanguagesOk

`func (o *APIConfig) GetAvailableLanguagesOk() (*[]string, bool)`

GetAvailableLanguagesOk returns a tuple with the AvailableLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableLanguages

`func (o *APIConfig) SetAvailableLanguages(v []string)`

SetAvailableLanguages sets AvailableLanguages field to given value.


### GetAvatarPerfLimiter

`func (o *APIConfig) GetAvatarPerfLimiter() APIConfigAvatarPerfLimiter`

GetAvatarPerfLimiter returns the AvatarPerfLimiter field if non-nil, zero value otherwise.

### GetAvatarPerfLimiterOk

`func (o *APIConfig) GetAvatarPerfLimiterOk() (*APIConfigAvatarPerfLimiter, bool)`

GetAvatarPerfLimiterOk returns a tuple with the AvatarPerfLimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarPerfLimiter

`func (o *APIConfig) SetAvatarPerfLimiter(v APIConfigAvatarPerfLimiter)`

SetAvatarPerfLimiter sets AvatarPerfLimiter field to given value.


### GetChatboxLogBufferSeconds

`func (o *APIConfig) GetChatboxLogBufferSeconds() int32`

GetChatboxLogBufferSeconds returns the ChatboxLogBufferSeconds field if non-nil, zero value otherwise.

### GetChatboxLogBufferSecondsOk

`func (o *APIConfig) GetChatboxLogBufferSecondsOk() (*int32, bool)`

GetChatboxLogBufferSecondsOk returns a tuple with the ChatboxLogBufferSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatboxLogBufferSeconds

`func (o *APIConfig) SetChatboxLogBufferSeconds(v int32)`

SetChatboxLogBufferSeconds sets ChatboxLogBufferSeconds field to given value.


### GetClientApiKey

`func (o *APIConfig) GetClientApiKey() string`

GetClientApiKey returns the ClientApiKey field if non-nil, zero value otherwise.

### GetClientApiKeyOk

`func (o *APIConfig) GetClientApiKeyOk() (*string, bool)`

GetClientApiKeyOk returns a tuple with the ClientApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientApiKey

`func (o *APIConfig) SetClientApiKey(v string)`

SetClientApiKey sets ClientApiKey field to given value.


### GetClientBPSCeiling

`func (o *APIConfig) GetClientBPSCeiling() int32`

GetClientBPSCeiling returns the ClientBPSCeiling field if non-nil, zero value otherwise.

### GetClientBPSCeilingOk

`func (o *APIConfig) GetClientBPSCeilingOk() (*int32, bool)`

GetClientBPSCeilingOk returns a tuple with the ClientBPSCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBPSCeiling

`func (o *APIConfig) SetClientBPSCeiling(v int32)`

SetClientBPSCeiling sets ClientBPSCeiling field to given value.


### GetClientDisconnectTimeout

`func (o *APIConfig) GetClientDisconnectTimeout() int32`

GetClientDisconnectTimeout returns the ClientDisconnectTimeout field if non-nil, zero value otherwise.

### GetClientDisconnectTimeoutOk

`func (o *APIConfig) GetClientDisconnectTimeoutOk() (*int32, bool)`

GetClientDisconnectTimeoutOk returns a tuple with the ClientDisconnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDisconnectTimeout

`func (o *APIConfig) SetClientDisconnectTimeout(v int32)`

SetClientDisconnectTimeout sets ClientDisconnectTimeout field to given value.


### GetClientNetDispatchThread

`func (o *APIConfig) GetClientNetDispatchThread() bool`

GetClientNetDispatchThread returns the ClientNetDispatchThread field if non-nil, zero value otherwise.

### GetClientNetDispatchThreadOk

`func (o *APIConfig) GetClientNetDispatchThreadOk() (*bool, bool)`

GetClientNetDispatchThreadOk returns a tuple with the ClientNetDispatchThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetDispatchThread

`func (o *APIConfig) SetClientNetDispatchThread(v bool)`

SetClientNetDispatchThread sets ClientNetDispatchThread field to given value.

### HasClientNetDispatchThread

`func (o *APIConfig) HasClientNetDispatchThread() bool`

HasClientNetDispatchThread returns a boolean if a field has been set.

### GetClientNetDispatchThreadMobile

`func (o *APIConfig) GetClientNetDispatchThreadMobile() bool`

GetClientNetDispatchThreadMobile returns the ClientNetDispatchThreadMobile field if non-nil, zero value otherwise.

### GetClientNetDispatchThreadMobileOk

`func (o *APIConfig) GetClientNetDispatchThreadMobileOk() (*bool, bool)`

GetClientNetDispatchThreadMobileOk returns a tuple with the ClientNetDispatchThreadMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetDispatchThreadMobile

`func (o *APIConfig) SetClientNetDispatchThreadMobile(v bool)`

SetClientNetDispatchThreadMobile sets ClientNetDispatchThreadMobile field to given value.


### GetClientNetInThread

`func (o *APIConfig) GetClientNetInThread() bool`

GetClientNetInThread returns the ClientNetInThread field if non-nil, zero value otherwise.

### GetClientNetInThreadOk

`func (o *APIConfig) GetClientNetInThreadOk() (*bool, bool)`

GetClientNetInThreadOk returns a tuple with the ClientNetInThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetInThread

`func (o *APIConfig) SetClientNetInThread(v bool)`

SetClientNetInThread sets ClientNetInThread field to given value.

### HasClientNetInThread

`func (o *APIConfig) HasClientNetInThread() bool`

HasClientNetInThread returns a boolean if a field has been set.

### GetClientNetInThread2

`func (o *APIConfig) GetClientNetInThread2() bool`

GetClientNetInThread2 returns the ClientNetInThread2 field if non-nil, zero value otherwise.

### GetClientNetInThread2Ok

`func (o *APIConfig) GetClientNetInThread2Ok() (*bool, bool)`

GetClientNetInThread2Ok returns a tuple with the ClientNetInThread2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetInThread2

`func (o *APIConfig) SetClientNetInThread2(v bool)`

SetClientNetInThread2 sets ClientNetInThread2 field to given value.

### HasClientNetInThread2

`func (o *APIConfig) HasClientNetInThread2() bool`

HasClientNetInThread2 returns a boolean if a field has been set.

### GetClientNetInThreadMobile

`func (o *APIConfig) GetClientNetInThreadMobile() bool`

GetClientNetInThreadMobile returns the ClientNetInThreadMobile field if non-nil, zero value otherwise.

### GetClientNetInThreadMobileOk

`func (o *APIConfig) GetClientNetInThreadMobileOk() (*bool, bool)`

GetClientNetInThreadMobileOk returns a tuple with the ClientNetInThreadMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetInThreadMobile

`func (o *APIConfig) SetClientNetInThreadMobile(v bool)`

SetClientNetInThreadMobile sets ClientNetInThreadMobile field to given value.

### HasClientNetInThreadMobile

`func (o *APIConfig) HasClientNetInThreadMobile() bool`

HasClientNetInThreadMobile returns a boolean if a field has been set.

### GetClientNetInThreadMobile2

`func (o *APIConfig) GetClientNetInThreadMobile2() bool`

GetClientNetInThreadMobile2 returns the ClientNetInThreadMobile2 field if non-nil, zero value otherwise.

### GetClientNetInThreadMobile2Ok

`func (o *APIConfig) GetClientNetInThreadMobile2Ok() (*bool, bool)`

GetClientNetInThreadMobile2Ok returns a tuple with the ClientNetInThreadMobile2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetInThreadMobile2

`func (o *APIConfig) SetClientNetInThreadMobile2(v bool)`

SetClientNetInThreadMobile2 sets ClientNetInThreadMobile2 field to given value.

### HasClientNetInThreadMobile2

`func (o *APIConfig) HasClientNetInThreadMobile2() bool`

HasClientNetInThreadMobile2 returns a boolean if a field has been set.

### GetClientNetOutThread

`func (o *APIConfig) GetClientNetOutThread() bool`

GetClientNetOutThread returns the ClientNetOutThread field if non-nil, zero value otherwise.

### GetClientNetOutThreadOk

`func (o *APIConfig) GetClientNetOutThreadOk() (*bool, bool)`

GetClientNetOutThreadOk returns a tuple with the ClientNetOutThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetOutThread

`func (o *APIConfig) SetClientNetOutThread(v bool)`

SetClientNetOutThread sets ClientNetOutThread field to given value.

### HasClientNetOutThread

`func (o *APIConfig) HasClientNetOutThread() bool`

HasClientNetOutThread returns a boolean if a field has been set.

### GetClientNetOutThread2

`func (o *APIConfig) GetClientNetOutThread2() bool`

GetClientNetOutThread2 returns the ClientNetOutThread2 field if non-nil, zero value otherwise.

### GetClientNetOutThread2Ok

`func (o *APIConfig) GetClientNetOutThread2Ok() (*bool, bool)`

GetClientNetOutThread2Ok returns a tuple with the ClientNetOutThread2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetOutThread2

`func (o *APIConfig) SetClientNetOutThread2(v bool)`

SetClientNetOutThread2 sets ClientNetOutThread2 field to given value.

### HasClientNetOutThread2

`func (o *APIConfig) HasClientNetOutThread2() bool`

HasClientNetOutThread2 returns a boolean if a field has been set.

### GetClientNetOutThreadMobile

`func (o *APIConfig) GetClientNetOutThreadMobile() bool`

GetClientNetOutThreadMobile returns the ClientNetOutThreadMobile field if non-nil, zero value otherwise.

### GetClientNetOutThreadMobileOk

`func (o *APIConfig) GetClientNetOutThreadMobileOk() (*bool, bool)`

GetClientNetOutThreadMobileOk returns a tuple with the ClientNetOutThreadMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetOutThreadMobile

`func (o *APIConfig) SetClientNetOutThreadMobile(v bool)`

SetClientNetOutThreadMobile sets ClientNetOutThreadMobile field to given value.

### HasClientNetOutThreadMobile

`func (o *APIConfig) HasClientNetOutThreadMobile() bool`

HasClientNetOutThreadMobile returns a boolean if a field has been set.

### GetClientNetOutThreadMobile2

`func (o *APIConfig) GetClientNetOutThreadMobile2() bool`

GetClientNetOutThreadMobile2 returns the ClientNetOutThreadMobile2 field if non-nil, zero value otherwise.

### GetClientNetOutThreadMobile2Ok

`func (o *APIConfig) GetClientNetOutThreadMobile2Ok() (*bool, bool)`

GetClientNetOutThreadMobile2Ok returns a tuple with the ClientNetOutThreadMobile2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetOutThreadMobile2

`func (o *APIConfig) SetClientNetOutThreadMobile2(v bool)`

SetClientNetOutThreadMobile2 sets ClientNetOutThreadMobile2 field to given value.

### HasClientNetOutThreadMobile2

`func (o *APIConfig) HasClientNetOutThreadMobile2() bool`

HasClientNetOutThreadMobile2 returns a boolean if a field has been set.

### GetClientQR

`func (o *APIConfig) GetClientQR() int32`

GetClientQR returns the ClientQR field if non-nil, zero value otherwise.

### GetClientQROk

`func (o *APIConfig) GetClientQROk() (*int32, bool)`

GetClientQROk returns a tuple with the ClientQR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientQR

`func (o *APIConfig) SetClientQR(v int32)`

SetClientQR sets ClientQR field to given value.

### HasClientQR

`func (o *APIConfig) HasClientQR() bool`

HasClientQR returns a boolean if a field has been set.

### GetClientReservedPlayerBPS

`func (o *APIConfig) GetClientReservedPlayerBPS() int32`

GetClientReservedPlayerBPS returns the ClientReservedPlayerBPS field if non-nil, zero value otherwise.

### GetClientReservedPlayerBPSOk

`func (o *APIConfig) GetClientReservedPlayerBPSOk() (*int32, bool)`

GetClientReservedPlayerBPSOk returns a tuple with the ClientReservedPlayerBPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReservedPlayerBPS

`func (o *APIConfig) SetClientReservedPlayerBPS(v int32)`

SetClientReservedPlayerBPS sets ClientReservedPlayerBPS field to given value.


### GetClientSentCountAllowance

`func (o *APIConfig) GetClientSentCountAllowance() int32`

GetClientSentCountAllowance returns the ClientSentCountAllowance field if non-nil, zero value otherwise.

### GetClientSentCountAllowanceOk

`func (o *APIConfig) GetClientSentCountAllowanceOk() (*int32, bool)`

GetClientSentCountAllowanceOk returns a tuple with the ClientSentCountAllowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSentCountAllowance

`func (o *APIConfig) SetClientSentCountAllowance(v int32)`

SetClientSentCountAllowance sets ClientSentCountAllowance field to given value.


### GetConstants

`func (o *APIConfig) GetConstants() APIConfigConstants`

GetConstants returns the Constants field if non-nil, zero value otherwise.

### GetConstantsOk

`func (o *APIConfig) GetConstantsOk() (*APIConfigConstants, bool)`

GetConstantsOk returns a tuple with the Constants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstants

`func (o *APIConfig) SetConstants(v APIConfigConstants)`

SetConstants sets Constants field to given value.


### GetContactEmail

`func (o *APIConfig) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *APIConfig) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *APIConfig) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.


### GetCopyrightEmail

`func (o *APIConfig) GetCopyrightEmail() string`

GetCopyrightEmail returns the CopyrightEmail field if non-nil, zero value otherwise.

### GetCopyrightEmailOk

`func (o *APIConfig) GetCopyrightEmailOk() (*string, bool)`

GetCopyrightEmailOk returns a tuple with the CopyrightEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightEmail

`func (o *APIConfig) SetCopyrightEmail(v string)`

SetCopyrightEmail sets CopyrightEmail field to given value.


### GetCopyrightFormUrl

`func (o *APIConfig) GetCopyrightFormUrl() string`

GetCopyrightFormUrl returns the CopyrightFormUrl field if non-nil, zero value otherwise.

### GetCopyrightFormUrlOk

`func (o *APIConfig) GetCopyrightFormUrlOk() (*string, bool)`

GetCopyrightFormUrlOk returns a tuple with the CopyrightFormUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightFormUrl

`func (o *APIConfig) SetCopyrightFormUrl(v string)`

SetCopyrightFormUrl sets CopyrightFormUrl field to given value.


### GetCurrentPrivacyVersion

`func (o *APIConfig) GetCurrentPrivacyVersion() int32`

GetCurrentPrivacyVersion returns the CurrentPrivacyVersion field if non-nil, zero value otherwise.

### GetCurrentPrivacyVersionOk

`func (o *APIConfig) GetCurrentPrivacyVersionOk() (*int32, bool)`

GetCurrentPrivacyVersionOk returns a tuple with the CurrentPrivacyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPrivacyVersion

`func (o *APIConfig) SetCurrentPrivacyVersion(v int32)`

SetCurrentPrivacyVersion sets CurrentPrivacyVersion field to given value.


### GetCurrentTOSVersion

`func (o *APIConfig) GetCurrentTOSVersion() int32`

GetCurrentTOSVersion returns the CurrentTOSVersion field if non-nil, zero value otherwise.

### GetCurrentTOSVersionOk

`func (o *APIConfig) GetCurrentTOSVersionOk() (*int32, bool)`

GetCurrentTOSVersionOk returns a tuple with the CurrentTOSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTOSVersion

`func (o *APIConfig) SetCurrentTOSVersion(v int32)`

SetCurrentTOSVersion sets CurrentTOSVersion field to given value.


### GetDefaultAvatar

`func (o *APIConfig) GetDefaultAvatar() string`

GetDefaultAvatar returns the DefaultAvatar field if non-nil, zero value otherwise.

### GetDefaultAvatarOk

`func (o *APIConfig) GetDefaultAvatarOk() (*string, bool)`

GetDefaultAvatarOk returns a tuple with the DefaultAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAvatar

`func (o *APIConfig) SetDefaultAvatar(v string)`

SetDefaultAvatar sets DefaultAvatar field to given value.


### GetDefaultStickerSet

`func (o *APIConfig) GetDefaultStickerSet() string`

GetDefaultStickerSet returns the DefaultStickerSet field if non-nil, zero value otherwise.

### GetDefaultStickerSetOk

`func (o *APIConfig) GetDefaultStickerSetOk() (*string, bool)`

GetDefaultStickerSetOk returns a tuple with the DefaultStickerSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStickerSet

`func (o *APIConfig) SetDefaultStickerSet(v string)`

SetDefaultStickerSet sets DefaultStickerSet field to given value.


### GetDevLanguageCodes

`func (o *APIConfig) GetDevLanguageCodes() []string`

GetDevLanguageCodes returns the DevLanguageCodes field if non-nil, zero value otherwise.

### GetDevLanguageCodesOk

`func (o *APIConfig) GetDevLanguageCodesOk() (*[]string, bool)`

GetDevLanguageCodesOk returns a tuple with the DevLanguageCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevLanguageCodes

`func (o *APIConfig) SetDevLanguageCodes(v []string)`

SetDevLanguageCodes sets DevLanguageCodes field to given value.

### HasDevLanguageCodes

`func (o *APIConfig) HasDevLanguageCodes() bool`

HasDevLanguageCodes returns a boolean if a field has been set.

### GetDevSdkUrl

`func (o *APIConfig) GetDevSdkUrl() string`

GetDevSdkUrl returns the DevSdkUrl field if non-nil, zero value otherwise.

### GetDevSdkUrlOk

`func (o *APIConfig) GetDevSdkUrlOk() (*string, bool)`

GetDevSdkUrlOk returns a tuple with the DevSdkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevSdkUrl

`func (o *APIConfig) SetDevSdkUrl(v string)`

SetDevSdkUrl sets DevSdkUrl field to given value.


### GetDevSdkVersion

`func (o *APIConfig) GetDevSdkVersion() string`

GetDevSdkVersion returns the DevSdkVersion field if non-nil, zero value otherwise.

### GetDevSdkVersionOk

`func (o *APIConfig) GetDevSdkVersionOk() (*string, bool)`

GetDevSdkVersionOk returns a tuple with the DevSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevSdkVersion

`func (o *APIConfig) SetDevSdkVersion(v string)`

SetDevSdkVersion sets DevSdkVersion field to given value.


### GetDisCountdown

`func (o *APIConfig) GetDisCountdown() time.Time`

GetDisCountdown returns the DisCountdown field if non-nil, zero value otherwise.

### GetDisCountdownOk

`func (o *APIConfig) GetDisCountdownOk() (*time.Time, bool)`

GetDisCountdownOk returns a tuple with the DisCountdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisCountdown

`func (o *APIConfig) SetDisCountdown(v time.Time)`

SetDisCountdown sets DisCountdown field to given value.


### GetDisableAVProInProton

`func (o *APIConfig) GetDisableAVProInProton() bool`

GetDisableAVProInProton returns the DisableAVProInProton field if non-nil, zero value otherwise.

### GetDisableAVProInProtonOk

`func (o *APIConfig) GetDisableAVProInProtonOk() (*bool, bool)`

GetDisableAVProInProtonOk returns a tuple with the DisableAVProInProton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAVProInProton

`func (o *APIConfig) SetDisableAVProInProton(v bool)`

SetDisableAVProInProton sets DisableAVProInProton field to given value.

### HasDisableAVProInProton

`func (o *APIConfig) HasDisableAVProInProton() bool`

HasDisableAVProInProton returns a boolean if a field has been set.

### GetDisableAvatarCopying

`func (o *APIConfig) GetDisableAvatarCopying() bool`

GetDisableAvatarCopying returns the DisableAvatarCopying field if non-nil, zero value otherwise.

### GetDisableAvatarCopyingOk

`func (o *APIConfig) GetDisableAvatarCopyingOk() (*bool, bool)`

GetDisableAvatarCopyingOk returns a tuple with the DisableAvatarCopying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAvatarCopying

`func (o *APIConfig) SetDisableAvatarCopying(v bool)`

SetDisableAvatarCopying sets DisableAvatarCopying field to given value.


### GetDisableAvatarGating

`func (o *APIConfig) GetDisableAvatarGating() bool`

GetDisableAvatarGating returns the DisableAvatarGating field if non-nil, zero value otherwise.

### GetDisableAvatarGatingOk

`func (o *APIConfig) GetDisableAvatarGatingOk() (*bool, bool)`

GetDisableAvatarGatingOk returns a tuple with the DisableAvatarGating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAvatarGating

`func (o *APIConfig) SetDisableAvatarGating(v bool)`

SetDisableAvatarGating sets DisableAvatarGating field to given value.


### GetDisableCaptcha

`func (o *APIConfig) GetDisableCaptcha() bool`

GetDisableCaptcha returns the DisableCaptcha field if non-nil, zero value otherwise.

### GetDisableCaptchaOk

`func (o *APIConfig) GetDisableCaptchaOk() (*bool, bool)`

GetDisableCaptchaOk returns a tuple with the DisableCaptcha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCaptcha

`func (o *APIConfig) SetDisableCaptcha(v bool)`

SetDisableCaptcha sets DisableCaptcha field to given value.

### HasDisableCaptcha

`func (o *APIConfig) HasDisableCaptcha() bool`

HasDisableCaptcha returns a boolean if a field has been set.

### GetDisableCommunityLabs

`func (o *APIConfig) GetDisableCommunityLabs() bool`

GetDisableCommunityLabs returns the DisableCommunityLabs field if non-nil, zero value otherwise.

### GetDisableCommunityLabsOk

`func (o *APIConfig) GetDisableCommunityLabsOk() (*bool, bool)`

GetDisableCommunityLabsOk returns a tuple with the DisableCommunityLabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCommunityLabs

`func (o *APIConfig) SetDisableCommunityLabs(v bool)`

SetDisableCommunityLabs sets DisableCommunityLabs field to given value.


### GetDisableCommunityLabsPromotion

`func (o *APIConfig) GetDisableCommunityLabsPromotion() bool`

GetDisableCommunityLabsPromotion returns the DisableCommunityLabsPromotion field if non-nil, zero value otherwise.

### GetDisableCommunityLabsPromotionOk

`func (o *APIConfig) GetDisableCommunityLabsPromotionOk() (*bool, bool)`

GetDisableCommunityLabsPromotionOk returns a tuple with the DisableCommunityLabsPromotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCommunityLabsPromotion

`func (o *APIConfig) SetDisableCommunityLabsPromotion(v bool)`

SetDisableCommunityLabsPromotion sets DisableCommunityLabsPromotion field to given value.


### GetDisableEmail

`func (o *APIConfig) GetDisableEmail() bool`

GetDisableEmail returns the DisableEmail field if non-nil, zero value otherwise.

### GetDisableEmailOk

`func (o *APIConfig) GetDisableEmailOk() (*bool, bool)`

GetDisableEmailOk returns a tuple with the DisableEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmail

`func (o *APIConfig) SetDisableEmail(v bool)`

SetDisableEmail sets DisableEmail field to given value.


### GetDisableEventStream

`func (o *APIConfig) GetDisableEventStream() bool`

GetDisableEventStream returns the DisableEventStream field if non-nil, zero value otherwise.

### GetDisableEventStreamOk

`func (o *APIConfig) GetDisableEventStreamOk() (*bool, bool)`

GetDisableEventStreamOk returns a tuple with the DisableEventStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEventStream

`func (o *APIConfig) SetDisableEventStream(v bool)`

SetDisableEventStream sets DisableEventStream field to given value.


### GetDisableFeedbackGating

`func (o *APIConfig) GetDisableFeedbackGating() bool`

GetDisableFeedbackGating returns the DisableFeedbackGating field if non-nil, zero value otherwise.

### GetDisableFeedbackGatingOk

`func (o *APIConfig) GetDisableFeedbackGatingOk() (*bool, bool)`

GetDisableFeedbackGatingOk returns a tuple with the DisableFeedbackGating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFeedbackGating

`func (o *APIConfig) SetDisableFeedbackGating(v bool)`

SetDisableFeedbackGating sets DisableFeedbackGating field to given value.


### GetDisableFrontendBuilds

`func (o *APIConfig) GetDisableFrontendBuilds() bool`

GetDisableFrontendBuilds returns the DisableFrontendBuilds field if non-nil, zero value otherwise.

### GetDisableFrontendBuildsOk

`func (o *APIConfig) GetDisableFrontendBuildsOk() (*bool, bool)`

GetDisableFrontendBuildsOk returns a tuple with the DisableFrontendBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFrontendBuilds

`func (o *APIConfig) SetDisableFrontendBuilds(v bool)`

SetDisableFrontendBuilds sets DisableFrontendBuilds field to given value.


### GetDisableGiftDrops

`func (o *APIConfig) GetDisableGiftDrops() bool`

GetDisableGiftDrops returns the DisableGiftDrops field if non-nil, zero value otherwise.

### GetDisableGiftDropsOk

`func (o *APIConfig) GetDisableGiftDropsOk() (*bool, bool)`

GetDisableGiftDropsOk returns a tuple with the DisableGiftDrops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableGiftDrops

`func (o *APIConfig) SetDisableGiftDrops(v bool)`

SetDisableGiftDrops sets DisableGiftDrops field to given value.


### GetDisableHello

`func (o *APIConfig) GetDisableHello() bool`

GetDisableHello returns the DisableHello field if non-nil, zero value otherwise.

### GetDisableHelloOk

`func (o *APIConfig) GetDisableHelloOk() (*bool, bool)`

GetDisableHelloOk returns a tuple with the DisableHello field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHello

`func (o *APIConfig) SetDisableHello(v bool)`

SetDisableHello sets DisableHello field to given value.


### GetDisableOculusSubs

`func (o *APIConfig) GetDisableOculusSubs() bool`

GetDisableOculusSubs returns the DisableOculusSubs field if non-nil, zero value otherwise.

### GetDisableOculusSubsOk

`func (o *APIConfig) GetDisableOculusSubsOk() (*bool, bool)`

GetDisableOculusSubsOk returns a tuple with the DisableOculusSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOculusSubs

`func (o *APIConfig) SetDisableOculusSubs(v bool)`

SetDisableOculusSubs sets DisableOculusSubs field to given value.


### GetDisableRegistration

`func (o *APIConfig) GetDisableRegistration() bool`

GetDisableRegistration returns the DisableRegistration field if non-nil, zero value otherwise.

### GetDisableRegistrationOk

`func (o *APIConfig) GetDisableRegistrationOk() (*bool, bool)`

GetDisableRegistrationOk returns a tuple with the DisableRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRegistration

`func (o *APIConfig) SetDisableRegistration(v bool)`

SetDisableRegistration sets DisableRegistration field to given value.


### GetDisableSteamNetworking

`func (o *APIConfig) GetDisableSteamNetworking() bool`

GetDisableSteamNetworking returns the DisableSteamNetworking field if non-nil, zero value otherwise.

### GetDisableSteamNetworkingOk

`func (o *APIConfig) GetDisableSteamNetworkingOk() (*bool, bool)`

GetDisableSteamNetworkingOk returns a tuple with the DisableSteamNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSteamNetworking

`func (o *APIConfig) SetDisableSteamNetworking(v bool)`

SetDisableSteamNetworking sets DisableSteamNetworking field to given value.


### GetDisableTwoFactorAuth

`func (o *APIConfig) GetDisableTwoFactorAuth() bool`

GetDisableTwoFactorAuth returns the DisableTwoFactorAuth field if non-nil, zero value otherwise.

### GetDisableTwoFactorAuthOk

`func (o *APIConfig) GetDisableTwoFactorAuthOk() (*bool, bool)`

GetDisableTwoFactorAuthOk returns a tuple with the DisableTwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTwoFactorAuth

`func (o *APIConfig) SetDisableTwoFactorAuth(v bool)`

SetDisableTwoFactorAuth sets DisableTwoFactorAuth field to given value.


### GetDisableUdon

`func (o *APIConfig) GetDisableUdon() bool`

GetDisableUdon returns the DisableUdon field if non-nil, zero value otherwise.

### GetDisableUdonOk

`func (o *APIConfig) GetDisableUdonOk() (*bool, bool)`

GetDisableUdonOk returns a tuple with the DisableUdon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUdon

`func (o *APIConfig) SetDisableUdon(v bool)`

SetDisableUdon sets DisableUdon field to given value.


### GetDisableUpgradeAccount

`func (o *APIConfig) GetDisableUpgradeAccount() bool`

GetDisableUpgradeAccount returns the DisableUpgradeAccount field if non-nil, zero value otherwise.

### GetDisableUpgradeAccountOk

`func (o *APIConfig) GetDisableUpgradeAccountOk() (*bool, bool)`

GetDisableUpgradeAccountOk returns a tuple with the DisableUpgradeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUpgradeAccount

`func (o *APIConfig) SetDisableUpgradeAccount(v bool)`

SetDisableUpgradeAccount sets DisableUpgradeAccount field to given value.


### GetDownloadLinkWindows

`func (o *APIConfig) GetDownloadLinkWindows() string`

GetDownloadLinkWindows returns the DownloadLinkWindows field if non-nil, zero value otherwise.

### GetDownloadLinkWindowsOk

`func (o *APIConfig) GetDownloadLinkWindowsOk() (*string, bool)`

GetDownloadLinkWindowsOk returns a tuple with the DownloadLinkWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLinkWindows

`func (o *APIConfig) SetDownloadLinkWindows(v string)`

SetDownloadLinkWindows sets DownloadLinkWindows field to given value.


### GetDownloadUrls

`func (o *APIConfig) GetDownloadUrls() APIConfigDownloadURLList`

GetDownloadUrls returns the DownloadUrls field if non-nil, zero value otherwise.

### GetDownloadUrlsOk

`func (o *APIConfig) GetDownloadUrlsOk() (*APIConfigDownloadURLList, bool)`

GetDownloadUrlsOk returns a tuple with the DownloadUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrls

`func (o *APIConfig) SetDownloadUrls(v APIConfigDownloadURLList)`

SetDownloadUrls sets DownloadUrls field to given value.


### GetDynamicWorldRows

`func (o *APIConfig) GetDynamicWorldRows() []DynamicContentRow`

GetDynamicWorldRows returns the DynamicWorldRows field if non-nil, zero value otherwise.

### GetDynamicWorldRowsOk

`func (o *APIConfig) GetDynamicWorldRowsOk() (*[]DynamicContentRow, bool)`

GetDynamicWorldRowsOk returns a tuple with the DynamicWorldRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicWorldRows

`func (o *APIConfig) SetDynamicWorldRows(v []DynamicContentRow)`

SetDynamicWorldRows sets DynamicWorldRows field to given value.


### GetEconomyLedgerBackfill

`func (o *APIConfig) GetEconomyLedgerBackfill() bool`

GetEconomyLedgerBackfill returns the EconomyLedgerBackfill field if non-nil, zero value otherwise.

### GetEconomyLedgerBackfillOk

`func (o *APIConfig) GetEconomyLedgerBackfillOk() (*bool, bool)`

GetEconomyLedgerBackfillOk returns a tuple with the EconomyLedgerBackfill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomyLedgerBackfill

`func (o *APIConfig) SetEconomyLedgerBackfill(v bool)`

SetEconomyLedgerBackfill sets EconomyLedgerBackfill field to given value.


### GetEconomyLedgerMigrationStop

`func (o *APIConfig) GetEconomyLedgerMigrationStop() string`

GetEconomyLedgerMigrationStop returns the EconomyLedgerMigrationStop field if non-nil, zero value otherwise.

### GetEconomyLedgerMigrationStopOk

`func (o *APIConfig) GetEconomyLedgerMigrationStopOk() (*string, bool)`

GetEconomyLedgerMigrationStopOk returns a tuple with the EconomyLedgerMigrationStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomyLedgerMigrationStop

`func (o *APIConfig) SetEconomyLedgerMigrationStop(v string)`

SetEconomyLedgerMigrationStop sets EconomyLedgerMigrationStop field to given value.


### GetEconomyLedgerMode

`func (o *APIConfig) GetEconomyLedgerMode() string`

GetEconomyLedgerMode returns the EconomyLedgerMode field if non-nil, zero value otherwise.

### GetEconomyLedgerModeOk

`func (o *APIConfig) GetEconomyLedgerModeOk() (*string, bool)`

GetEconomyLedgerModeOk returns a tuple with the EconomyLedgerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomyLedgerMode

`func (o *APIConfig) SetEconomyLedgerMode(v string)`

SetEconomyLedgerMode sets EconomyLedgerMode field to given value.


### GetEconomyPauseEnd

`func (o *APIConfig) GetEconomyPauseEnd() time.Time`

GetEconomyPauseEnd returns the EconomyPauseEnd field if non-nil, zero value otherwise.

### GetEconomyPauseEndOk

`func (o *APIConfig) GetEconomyPauseEndOk() (*time.Time, bool)`

GetEconomyPauseEndOk returns a tuple with the EconomyPauseEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomyPauseEnd

`func (o *APIConfig) SetEconomyPauseEnd(v time.Time)`

SetEconomyPauseEnd sets EconomyPauseEnd field to given value.


### GetEconomyPauseStart

`func (o *APIConfig) GetEconomyPauseStart() time.Time`

GetEconomyPauseStart returns the EconomyPauseStart field if non-nil, zero value otherwise.

### GetEconomyPauseStartOk

`func (o *APIConfig) GetEconomyPauseStartOk() (*time.Time, bool)`

GetEconomyPauseStartOk returns a tuple with the EconomyPauseStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomyPauseStart

`func (o *APIConfig) SetEconomyPauseStart(v time.Time)`

SetEconomyPauseStart sets EconomyPauseStart field to given value.


### GetEconomyPurchaseRepairEnabled

`func (o *APIConfig) GetEconomyPurchaseRepairEnabled() bool`

GetEconomyPurchaseRepairEnabled returns the EconomyPurchaseRepairEnabled field if non-nil, zero value otherwise.

### GetEconomyPurchaseRepairEnabledOk

`func (o *APIConfig) GetEconomyPurchaseRepairEnabledOk() (*bool, bool)`

GetEconomyPurchaseRepairEnabledOk returns a tuple with the EconomyPurchaseRepairEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomyPurchaseRepairEnabled

`func (o *APIConfig) SetEconomyPurchaseRepairEnabled(v bool)`

SetEconomyPurchaseRepairEnabled sets EconomyPurchaseRepairEnabled field to given value.


### GetEconomyState

`func (o *APIConfig) GetEconomyState() int32`

GetEconomyState returns the EconomyState field if non-nil, zero value otherwise.

### GetEconomyStateOk

`func (o *APIConfig) GetEconomyStateOk() (*int32, bool)`

GetEconomyStateOk returns a tuple with the EconomyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomyState

`func (o *APIConfig) SetEconomyState(v int32)`

SetEconomyState sets EconomyState field to given value.


### GetEvents

`func (o *APIConfig) GetEvents() APIConfigEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *APIConfig) GetEventsOk() (*APIConfigEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *APIConfig) SetEvents(v APIConfigEvents)`

SetEvents sets Events field to given value.


### GetForceUseLatestWorld

`func (o *APIConfig) GetForceUseLatestWorld() bool`

GetForceUseLatestWorld returns the ForceUseLatestWorld field if non-nil, zero value otherwise.

### GetForceUseLatestWorldOk

`func (o *APIConfig) GetForceUseLatestWorldOk() (*bool, bool)`

GetForceUseLatestWorldOk returns a tuple with the ForceUseLatestWorld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUseLatestWorld

`func (o *APIConfig) SetForceUseLatestWorld(v bool)`

SetForceUseLatestWorld sets ForceUseLatestWorld field to given value.


### GetGiftDisplayType

`func (o *APIConfig) GetGiftDisplayType() string`

GetGiftDisplayType returns the GiftDisplayType field if non-nil, zero value otherwise.

### GetGiftDisplayTypeOk

`func (o *APIConfig) GetGiftDisplayTypeOk() (*string, bool)`

GetGiftDisplayTypeOk returns a tuple with the GiftDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftDisplayType

`func (o *APIConfig) SetGiftDisplayType(v string)`

SetGiftDisplayType sets GiftDisplayType field to given value.


### GetGoogleApiClientId

`func (o *APIConfig) GetGoogleApiClientId() string`

GetGoogleApiClientId returns the GoogleApiClientId field if non-nil, zero value otherwise.

### GetGoogleApiClientIdOk

`func (o *APIConfig) GetGoogleApiClientIdOk() (*string, bool)`

GetGoogleApiClientIdOk returns a tuple with the GoogleApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleApiClientId

`func (o *APIConfig) SetGoogleApiClientId(v string)`

SetGoogleApiClientId sets GoogleApiClientId field to given value.


### GetHomeWorldId

`func (o *APIConfig) GetHomeWorldId() string`

GetHomeWorldId returns the HomeWorldId field if non-nil, zero value otherwise.

### GetHomeWorldIdOk

`func (o *APIConfig) GetHomeWorldIdOk() (*string, bool)`

GetHomeWorldIdOk returns a tuple with the HomeWorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeWorldId

`func (o *APIConfig) SetHomeWorldId(v string)`

SetHomeWorldId sets HomeWorldId field to given value.


### GetHomepageRedirectTarget

`func (o *APIConfig) GetHomepageRedirectTarget() string`

GetHomepageRedirectTarget returns the HomepageRedirectTarget field if non-nil, zero value otherwise.

### GetHomepageRedirectTargetOk

`func (o *APIConfig) GetHomepageRedirectTargetOk() (*string, bool)`

GetHomepageRedirectTargetOk returns a tuple with the HomepageRedirectTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageRedirectTarget

`func (o *APIConfig) SetHomepageRedirectTarget(v string)`

SetHomepageRedirectTarget sets HomepageRedirectTarget field to given value.


### GetHubWorldId

`func (o *APIConfig) GetHubWorldId() string`

GetHubWorldId returns the HubWorldId field if non-nil, zero value otherwise.

### GetHubWorldIdOk

`func (o *APIConfig) GetHubWorldIdOk() (*string, bool)`

GetHubWorldIdOk returns a tuple with the HubWorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubWorldId

`func (o *APIConfig) SetHubWorldId(v string)`

SetHubWorldId sets HubWorldId field to given value.


### GetImageHostUrlList

`func (o *APIConfig) GetImageHostUrlList() []string`

GetImageHostUrlList returns the ImageHostUrlList field if non-nil, zero value otherwise.

### GetImageHostUrlListOk

`func (o *APIConfig) GetImageHostUrlListOk() (*[]string, bool)`

GetImageHostUrlListOk returns a tuple with the ImageHostUrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHostUrlList

`func (o *APIConfig) SetImageHostUrlList(v []string)`

SetImageHostUrlList sets ImageHostUrlList field to given value.


### GetIosAppVersion

`func (o *APIConfig) GetIosAppVersion() []string`

GetIosAppVersion returns the IosAppVersion field if non-nil, zero value otherwise.

### GetIosAppVersionOk

`func (o *APIConfig) GetIosAppVersionOk() (*[]string, bool)`

GetIosAppVersionOk returns a tuple with the IosAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosAppVersion

`func (o *APIConfig) SetIosAppVersion(v []string)`

SetIosAppVersion sets IosAppVersion field to given value.


### GetIosVersion

`func (o *APIConfig) GetIosVersion() APIConfigIosVersion`

GetIosVersion returns the IosVersion field if non-nil, zero value otherwise.

### GetIosVersionOk

`func (o *APIConfig) GetIosVersionOk() (*APIConfigIosVersion, bool)`

GetIosVersionOk returns a tuple with the IosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosVersion

`func (o *APIConfig) SetIosVersion(v APIConfigIosVersion)`

SetIosVersion sets IosVersion field to given value.


### GetJobsEmail

`func (o *APIConfig) GetJobsEmail() string`

GetJobsEmail returns the JobsEmail field if non-nil, zero value otherwise.

### GetJobsEmailOk

`func (o *APIConfig) GetJobsEmailOk() (*string, bool)`

GetJobsEmailOk returns a tuple with the JobsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsEmail

`func (o *APIConfig) SetJobsEmail(v string)`

SetJobsEmail sets JobsEmail field to given value.


### GetMaxUserEmoji

`func (o *APIConfig) GetMaxUserEmoji() int32`

GetMaxUserEmoji returns the MaxUserEmoji field if non-nil, zero value otherwise.

### GetMaxUserEmojiOk

`func (o *APIConfig) GetMaxUserEmojiOk() (*int32, bool)`

GetMaxUserEmojiOk returns a tuple with the MaxUserEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUserEmoji

`func (o *APIConfig) SetMaxUserEmoji(v int32)`

SetMaxUserEmoji sets MaxUserEmoji field to given value.


### GetMaxUserStickers

`func (o *APIConfig) GetMaxUserStickers() int32`

GetMaxUserStickers returns the MaxUserStickers field if non-nil, zero value otherwise.

### GetMaxUserStickersOk

`func (o *APIConfig) GetMaxUserStickersOk() (*int32, bool)`

GetMaxUserStickersOk returns a tuple with the MaxUserStickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUserStickers

`func (o *APIConfig) SetMaxUserStickers(v int32)`

SetMaxUserStickers sets MaxUserStickers field to given value.


### GetMinSupportedClientBuildNumber

`func (o *APIConfig) GetMinSupportedClientBuildNumber() APIConfigMinSupportedClientBuildNumber`

GetMinSupportedClientBuildNumber returns the MinSupportedClientBuildNumber field if non-nil, zero value otherwise.

### GetMinSupportedClientBuildNumberOk

`func (o *APIConfig) GetMinSupportedClientBuildNumberOk() (*APIConfigMinSupportedClientBuildNumber, bool)`

GetMinSupportedClientBuildNumberOk returns a tuple with the MinSupportedClientBuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSupportedClientBuildNumber

`func (o *APIConfig) SetMinSupportedClientBuildNumber(v APIConfigMinSupportedClientBuildNumber)`

SetMinSupportedClientBuildNumber sets MinSupportedClientBuildNumber field to given value.


### GetMinimumUnityVersionForUploads

`func (o *APIConfig) GetMinimumUnityVersionForUploads() string`

GetMinimumUnityVersionForUploads returns the MinimumUnityVersionForUploads field if non-nil, zero value otherwise.

### GetMinimumUnityVersionForUploadsOk

`func (o *APIConfig) GetMinimumUnityVersionForUploadsOk() (*string, bool)`

GetMinimumUnityVersionForUploadsOk returns a tuple with the MinimumUnityVersionForUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumUnityVersionForUploads

`func (o *APIConfig) SetMinimumUnityVersionForUploads(v string)`

SetMinimumUnityVersionForUploads sets MinimumUnityVersionForUploads field to given value.


### GetModerationEmail

`func (o *APIConfig) GetModerationEmail() string`

GetModerationEmail returns the ModerationEmail field if non-nil, zero value otherwise.

### GetModerationEmailOk

`func (o *APIConfig) GetModerationEmailOk() (*string, bool)`

GetModerationEmailOk returns a tuple with the ModerationEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationEmail

`func (o *APIConfig) SetModerationEmail(v string)`

SetModerationEmail sets ModerationEmail field to given value.


### GetNotAllowedToSelectAvatarInPrivateWorldMessage

`func (o *APIConfig) GetNotAllowedToSelectAvatarInPrivateWorldMessage() string`

GetNotAllowedToSelectAvatarInPrivateWorldMessage returns the NotAllowedToSelectAvatarInPrivateWorldMessage field if non-nil, zero value otherwise.

### GetNotAllowedToSelectAvatarInPrivateWorldMessageOk

`func (o *APIConfig) GetNotAllowedToSelectAvatarInPrivateWorldMessageOk() (*string, bool)`

GetNotAllowedToSelectAvatarInPrivateWorldMessageOk returns a tuple with the NotAllowedToSelectAvatarInPrivateWorldMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllowedToSelectAvatarInPrivateWorldMessage

`func (o *APIConfig) SetNotAllowedToSelectAvatarInPrivateWorldMessage(v string)`

SetNotAllowedToSelectAvatarInPrivateWorldMessage sets NotAllowedToSelectAvatarInPrivateWorldMessage field to given value.


### GetOfflineAnalysis

`func (o *APIConfig) GetOfflineAnalysis() APIConfigOfflineAnalysis`

GetOfflineAnalysis returns the OfflineAnalysis field if non-nil, zero value otherwise.

### GetOfflineAnalysisOk

`func (o *APIConfig) GetOfflineAnalysisOk() (*APIConfigOfflineAnalysis, bool)`

GetOfflineAnalysisOk returns a tuple with the OfflineAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineAnalysis

`func (o *APIConfig) SetOfflineAnalysis(v APIConfigOfflineAnalysis)`

SetOfflineAnalysis sets OfflineAnalysis field to given value.


### GetPhotonNameserverOverrides

`func (o *APIConfig) GetPhotonNameserverOverrides() []string`

GetPhotonNameserverOverrides returns the PhotonNameserverOverrides field if non-nil, zero value otherwise.

### GetPhotonNameserverOverridesOk

`func (o *APIConfig) GetPhotonNameserverOverridesOk() (*[]string, bool)`

GetPhotonNameserverOverridesOk returns a tuple with the PhotonNameserverOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotonNameserverOverrides

`func (o *APIConfig) SetPhotonNameserverOverrides(v []string)`

SetPhotonNameserverOverrides sets PhotonNameserverOverrides field to given value.


### GetPhotonPublicKeys

`func (o *APIConfig) GetPhotonPublicKeys() []string`

GetPhotonPublicKeys returns the PhotonPublicKeys field if non-nil, zero value otherwise.

### GetPhotonPublicKeysOk

`func (o *APIConfig) GetPhotonPublicKeysOk() (*[]string, bool)`

GetPhotonPublicKeysOk returns a tuple with the PhotonPublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotonPublicKeys

`func (o *APIConfig) SetPhotonPublicKeys(v []string)`

SetPhotonPublicKeys sets PhotonPublicKeys field to given value.


### GetPlayerUrlResolverSha1

`func (o *APIConfig) GetPlayerUrlResolverSha1() string`

GetPlayerUrlResolverSha1 returns the PlayerUrlResolverSha1 field if non-nil, zero value otherwise.

### GetPlayerUrlResolverSha1Ok

`func (o *APIConfig) GetPlayerUrlResolverSha1Ok() (*string, bool)`

GetPlayerUrlResolverSha1Ok returns a tuple with the PlayerUrlResolverSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerUrlResolverSha1

`func (o *APIConfig) SetPlayerUrlResolverSha1(v string)`

SetPlayerUrlResolverSha1 sets PlayerUrlResolverSha1 field to given value.


### GetPlayerUrlResolverVersion

`func (o *APIConfig) GetPlayerUrlResolverVersion() string`

GetPlayerUrlResolverVersion returns the PlayerUrlResolverVersion field if non-nil, zero value otherwise.

### GetPlayerUrlResolverVersionOk

`func (o *APIConfig) GetPlayerUrlResolverVersionOk() (*string, bool)`

GetPlayerUrlResolverVersionOk returns a tuple with the PlayerUrlResolverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerUrlResolverVersion

`func (o *APIConfig) SetPlayerUrlResolverVersion(v string)`

SetPlayerUrlResolverVersion sets PlayerUrlResolverVersion field to given value.


### GetPublicKey

`func (o *APIConfig) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *APIConfig) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *APIConfig) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetReportCategories

`func (o *APIConfig) GetReportCategories() map[string]ReportCategory`

GetReportCategories returns the ReportCategories field if non-nil, zero value otherwise.

### GetReportCategoriesOk

`func (o *APIConfig) GetReportCategoriesOk() (*map[string]ReportCategory, bool)`

GetReportCategoriesOk returns a tuple with the ReportCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCategories

`func (o *APIConfig) SetReportCategories(v map[string]ReportCategory)`

SetReportCategories sets ReportCategories field to given value.


### GetReportFormUrl

`func (o *APIConfig) GetReportFormUrl() string`

GetReportFormUrl returns the ReportFormUrl field if non-nil, zero value otherwise.

### GetReportFormUrlOk

`func (o *APIConfig) GetReportFormUrlOk() (*string, bool)`

GetReportFormUrlOk returns a tuple with the ReportFormUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormUrl

`func (o *APIConfig) SetReportFormUrl(v string)`

SetReportFormUrl sets ReportFormUrl field to given value.


### GetReportOptions

`func (o *APIConfig) GetReportOptions() map[string]map[string][]string`

GetReportOptions returns the ReportOptions field if non-nil, zero value otherwise.

### GetReportOptionsOk

`func (o *APIConfig) GetReportOptionsOk() (*map[string]map[string][]string, bool)`

GetReportOptionsOk returns a tuple with the ReportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportOptions

`func (o *APIConfig) SetReportOptions(v map[string]map[string][]string)`

SetReportOptions sets ReportOptions field to given value.


### GetReportReasons

`func (o *APIConfig) GetReportReasons() map[string]ReportReason`

GetReportReasons returns the ReportReasons field if non-nil, zero value otherwise.

### GetReportReasonsOk

`func (o *APIConfig) GetReportReasonsOk() (*map[string]ReportReason, bool)`

GetReportReasonsOk returns a tuple with the ReportReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportReasons

`func (o *APIConfig) SetReportReasons(v map[string]ReportReason)`

SetReportReasons sets ReportReasons field to given value.


### GetRequireAgeVerificationBetaTag

`func (o *APIConfig) GetRequireAgeVerificationBetaTag() bool`

GetRequireAgeVerificationBetaTag returns the RequireAgeVerificationBetaTag field if non-nil, zero value otherwise.

### GetRequireAgeVerificationBetaTagOk

`func (o *APIConfig) GetRequireAgeVerificationBetaTagOk() (*bool, bool)`

GetRequireAgeVerificationBetaTagOk returns a tuple with the RequireAgeVerificationBetaTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAgeVerificationBetaTag

`func (o *APIConfig) SetRequireAgeVerificationBetaTag(v bool)`

SetRequireAgeVerificationBetaTag sets RequireAgeVerificationBetaTag field to given value.


### GetSdkDeveloperFaqUrl

`func (o *APIConfig) GetSdkDeveloperFaqUrl() string`

GetSdkDeveloperFaqUrl returns the SdkDeveloperFaqUrl field if non-nil, zero value otherwise.

### GetSdkDeveloperFaqUrlOk

`func (o *APIConfig) GetSdkDeveloperFaqUrlOk() (*string, bool)`

GetSdkDeveloperFaqUrlOk returns a tuple with the SdkDeveloperFaqUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkDeveloperFaqUrl

`func (o *APIConfig) SetSdkDeveloperFaqUrl(v string)`

SetSdkDeveloperFaqUrl sets SdkDeveloperFaqUrl field to given value.


### GetSdkDiscordUrl

`func (o *APIConfig) GetSdkDiscordUrl() string`

GetSdkDiscordUrl returns the SdkDiscordUrl field if non-nil, zero value otherwise.

### GetSdkDiscordUrlOk

`func (o *APIConfig) GetSdkDiscordUrlOk() (*string, bool)`

GetSdkDiscordUrlOk returns a tuple with the SdkDiscordUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkDiscordUrl

`func (o *APIConfig) SetSdkDiscordUrl(v string)`

SetSdkDiscordUrl sets SdkDiscordUrl field to given value.


### GetSdkNotAllowedToPublishMessage

`func (o *APIConfig) GetSdkNotAllowedToPublishMessage() string`

GetSdkNotAllowedToPublishMessage returns the SdkNotAllowedToPublishMessage field if non-nil, zero value otherwise.

### GetSdkNotAllowedToPublishMessageOk

`func (o *APIConfig) GetSdkNotAllowedToPublishMessageOk() (*string, bool)`

GetSdkNotAllowedToPublishMessageOk returns a tuple with the SdkNotAllowedToPublishMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkNotAllowedToPublishMessage

`func (o *APIConfig) SetSdkNotAllowedToPublishMessage(v string)`

SetSdkNotAllowedToPublishMessage sets SdkNotAllowedToPublishMessage field to given value.


### GetSdkUnityVersion

`func (o *APIConfig) GetSdkUnityVersion() string`

GetSdkUnityVersion returns the SdkUnityVersion field if non-nil, zero value otherwise.

### GetSdkUnityVersionOk

`func (o *APIConfig) GetSdkUnityVersionOk() (*string, bool)`

GetSdkUnityVersionOk returns a tuple with the SdkUnityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkUnityVersion

`func (o *APIConfig) SetSdkUnityVersion(v string)`

SetSdkUnityVersion sets SdkUnityVersion field to given value.


### GetStringHostUrlList

`func (o *APIConfig) GetStringHostUrlList() []string`

GetStringHostUrlList returns the StringHostUrlList field if non-nil, zero value otherwise.

### GetStringHostUrlListOk

`func (o *APIConfig) GetStringHostUrlListOk() (*[]string, bool)`

GetStringHostUrlListOk returns a tuple with the StringHostUrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringHostUrlList

`func (o *APIConfig) SetStringHostUrlList(v []string)`

SetStringHostUrlList sets StringHostUrlList field to given value.


### GetSupportEmail

`func (o *APIConfig) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *APIConfig) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *APIConfig) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.


### GetSupportFormUrl

`func (o *APIConfig) GetSupportFormUrl() string`

GetSupportFormUrl returns the SupportFormUrl field if non-nil, zero value otherwise.

### GetSupportFormUrlOk

`func (o *APIConfig) GetSupportFormUrlOk() (*string, bool)`

GetSupportFormUrlOk returns a tuple with the SupportFormUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportFormUrl

`func (o *APIConfig) SetSupportFormUrl(v string)`

SetSupportFormUrl sets SupportFormUrl field to given value.


### GetTimeOutWorldId

`func (o *APIConfig) GetTimeOutWorldId() string`

GetTimeOutWorldId returns the TimeOutWorldId field if non-nil, zero value otherwise.

### GetTimeOutWorldIdOk

`func (o *APIConfig) GetTimeOutWorldIdOk() (*string, bool)`

GetTimeOutWorldIdOk returns a tuple with the TimeOutWorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOutWorldId

`func (o *APIConfig) SetTimeOutWorldId(v string)`

SetTimeOutWorldId sets TimeOutWorldId field to given value.


### GetTimekeeping

`func (o *APIConfig) GetTimekeeping() bool`

GetTimekeeping returns the Timekeeping field if non-nil, zero value otherwise.

### GetTimekeepingOk

`func (o *APIConfig) GetTimekeepingOk() (*bool, bool)`

GetTimekeepingOk returns a tuple with the Timekeeping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimekeeping

`func (o *APIConfig) SetTimekeeping(v bool)`

SetTimekeeping sets Timekeeping field to given value.


### GetTutorialWorldId

`func (o *APIConfig) GetTutorialWorldId() string`

GetTutorialWorldId returns the TutorialWorldId field if non-nil, zero value otherwise.

### GetTutorialWorldIdOk

`func (o *APIConfig) GetTutorialWorldIdOk() (*string, bool)`

GetTutorialWorldIdOk returns a tuple with the TutorialWorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTutorialWorldId

`func (o *APIConfig) SetTutorialWorldId(v string)`

SetTutorialWorldId sets TutorialWorldId field to given value.


### GetUpdateRateMsMaximum

`func (o *APIConfig) GetUpdateRateMsMaximum() int32`

GetUpdateRateMsMaximum returns the UpdateRateMsMaximum field if non-nil, zero value otherwise.

### GetUpdateRateMsMaximumOk

`func (o *APIConfig) GetUpdateRateMsMaximumOk() (*int32, bool)`

GetUpdateRateMsMaximumOk returns a tuple with the UpdateRateMsMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRateMsMaximum

`func (o *APIConfig) SetUpdateRateMsMaximum(v int32)`

SetUpdateRateMsMaximum sets UpdateRateMsMaximum field to given value.


### GetUpdateRateMsMinimum

`func (o *APIConfig) GetUpdateRateMsMinimum() int32`

GetUpdateRateMsMinimum returns the UpdateRateMsMinimum field if non-nil, zero value otherwise.

### GetUpdateRateMsMinimumOk

`func (o *APIConfig) GetUpdateRateMsMinimumOk() (*int32, bool)`

GetUpdateRateMsMinimumOk returns a tuple with the UpdateRateMsMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRateMsMinimum

`func (o *APIConfig) SetUpdateRateMsMinimum(v int32)`

SetUpdateRateMsMinimum sets UpdateRateMsMinimum field to given value.


### GetUpdateRateMsNormal

`func (o *APIConfig) GetUpdateRateMsNormal() int32`

GetUpdateRateMsNormal returns the UpdateRateMsNormal field if non-nil, zero value otherwise.

### GetUpdateRateMsNormalOk

`func (o *APIConfig) GetUpdateRateMsNormalOk() (*int32, bool)`

GetUpdateRateMsNormalOk returns a tuple with the UpdateRateMsNormal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRateMsNormal

`func (o *APIConfig) SetUpdateRateMsNormal(v int32)`

SetUpdateRateMsNormal sets UpdateRateMsNormal field to given value.


### GetUpdateRateMsUdonManual

`func (o *APIConfig) GetUpdateRateMsUdonManual() int32`

GetUpdateRateMsUdonManual returns the UpdateRateMsUdonManual field if non-nil, zero value otherwise.

### GetUpdateRateMsUdonManualOk

`func (o *APIConfig) GetUpdateRateMsUdonManualOk() (*int32, bool)`

GetUpdateRateMsUdonManualOk returns a tuple with the UpdateRateMsUdonManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRateMsUdonManual

`func (o *APIConfig) SetUpdateRateMsUdonManual(v int32)`

SetUpdateRateMsUdonManual sets UpdateRateMsUdonManual field to given value.


### GetUploadAnalysisPercent

`func (o *APIConfig) GetUploadAnalysisPercent() int32`

GetUploadAnalysisPercent returns the UploadAnalysisPercent field if non-nil, zero value otherwise.

### GetUploadAnalysisPercentOk

`func (o *APIConfig) GetUploadAnalysisPercentOk() (*int32, bool)`

GetUploadAnalysisPercentOk returns a tuple with the UploadAnalysisPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadAnalysisPercent

`func (o *APIConfig) SetUploadAnalysisPercent(v int32)`

SetUploadAnalysisPercent sets UploadAnalysisPercent field to given value.


### GetUrlList

`func (o *APIConfig) GetUrlList() []string`

GetUrlList returns the UrlList field if non-nil, zero value otherwise.

### GetUrlListOk

`func (o *APIConfig) GetUrlListOk() (*[]string, bool)`

GetUrlListOk returns a tuple with the UrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlList

`func (o *APIConfig) SetUrlList(v []string)`

SetUrlList sets UrlList field to given value.


### GetUseReliableUdpForVoice

`func (o *APIConfig) GetUseReliableUdpForVoice() bool`

GetUseReliableUdpForVoice returns the UseReliableUdpForVoice field if non-nil, zero value otherwise.

### GetUseReliableUdpForVoiceOk

`func (o *APIConfig) GetUseReliableUdpForVoiceOk() (*bool, bool)`

GetUseReliableUdpForVoiceOk returns a tuple with the UseReliableUdpForVoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseReliableUdpForVoice

`func (o *APIConfig) SetUseReliableUdpForVoice(v bool)`

SetUseReliableUdpForVoice sets UseReliableUdpForVoice field to given value.


### GetViveWindowsUrl

`func (o *APIConfig) GetViveWindowsUrl() string`

GetViveWindowsUrl returns the ViveWindowsUrl field if non-nil, zero value otherwise.

### GetViveWindowsUrlOk

`func (o *APIConfig) GetViveWindowsUrlOk() (*string, bool)`

GetViveWindowsUrlOk returns a tuple with the ViveWindowsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViveWindowsUrl

`func (o *APIConfig) SetViveWindowsUrl(v string)`

SetViveWindowsUrl sets ViveWindowsUrl field to given value.


### GetWebsocketMaxFriendsRefreshDelay

`func (o *APIConfig) GetWebsocketMaxFriendsRefreshDelay() int32`

GetWebsocketMaxFriendsRefreshDelay returns the WebsocketMaxFriendsRefreshDelay field if non-nil, zero value otherwise.

### GetWebsocketMaxFriendsRefreshDelayOk

`func (o *APIConfig) GetWebsocketMaxFriendsRefreshDelayOk() (*int32, bool)`

GetWebsocketMaxFriendsRefreshDelayOk returns a tuple with the WebsocketMaxFriendsRefreshDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketMaxFriendsRefreshDelay

`func (o *APIConfig) SetWebsocketMaxFriendsRefreshDelay(v int32)`

SetWebsocketMaxFriendsRefreshDelay sets WebsocketMaxFriendsRefreshDelay field to given value.


### GetWebsocketQuickReconnectTime

`func (o *APIConfig) GetWebsocketQuickReconnectTime() int32`

GetWebsocketQuickReconnectTime returns the WebsocketQuickReconnectTime field if non-nil, zero value otherwise.

### GetWebsocketQuickReconnectTimeOk

`func (o *APIConfig) GetWebsocketQuickReconnectTimeOk() (*int32, bool)`

GetWebsocketQuickReconnectTimeOk returns a tuple with the WebsocketQuickReconnectTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketQuickReconnectTime

`func (o *APIConfig) SetWebsocketQuickReconnectTime(v int32)`

SetWebsocketQuickReconnectTime sets WebsocketQuickReconnectTime field to given value.


### GetWebsocketReconnectMaxDelay

`func (o *APIConfig) GetWebsocketReconnectMaxDelay() int32`

GetWebsocketReconnectMaxDelay returns the WebsocketReconnectMaxDelay field if non-nil, zero value otherwise.

### GetWebsocketReconnectMaxDelayOk

`func (o *APIConfig) GetWebsocketReconnectMaxDelayOk() (*int32, bool)`

GetWebsocketReconnectMaxDelayOk returns a tuple with the WebsocketReconnectMaxDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketReconnectMaxDelay

`func (o *APIConfig) SetWebsocketReconnectMaxDelay(v int32)`

SetWebsocketReconnectMaxDelay sets WebsocketReconnectMaxDelay field to given value.


### GetWhiteListedAssetUrls

`func (o *APIConfig) GetWhiteListedAssetUrls() []string`

GetWhiteListedAssetUrls returns the WhiteListedAssetUrls field if non-nil, zero value otherwise.

### GetWhiteListedAssetUrlsOk

`func (o *APIConfig) GetWhiteListedAssetUrlsOk() (*[]string, bool)`

GetWhiteListedAssetUrlsOk returns a tuple with the WhiteListedAssetUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteListedAssetUrls

`func (o *APIConfig) SetWhiteListedAssetUrls(v []string)`

SetWhiteListedAssetUrls sets WhiteListedAssetUrls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


