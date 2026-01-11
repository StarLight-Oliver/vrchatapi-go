# Go API client for vrchatapi

![VRChat API Banner](https://vrchatapi.github.io/assets/img/api_banner_1500x400.png)

# Welcome to the VRChat API

Before we begin, we would like to state this is a **COMMUNITY DRIVEN PROJECT**.
This means that everything you read on here was written by the community itself and is **not** officially supported by VRChat.
The documentation is provided \"AS IS\", and any action you take towards VRChat is completely your own responsibility.

The documentation and additional libraries SHALL ONLY be used for applications interacting with VRChat's API in accordance
with their [Terms of Service](https://hello.vrchat.com/legal) and [Community Guidelines](https://hello.vrchat.com/community-guidelines), and MUST NOT be used for modifying the client, \"avatar ripping\", or other illegal activities.
Malicious usage or spamming the API may result in account termination.
Certain parts of the API are also more sensitive than others, for example moderation, so please tread extra carefully and read the warnings when present.

![Tupper Policy on API](https://i.imgur.com/yLlW7Ok.png)

Finally, use of the API using applications other than the approved methods (website, VRChat application, Unity SDK) is not officially supported.
VRChat provides no guarantee or support for external applications using the API. Access to API endpoints may break **at any time, without notice**.
Therefore, please **do not ping** VRChat Staff in the VRChat Discord if you are having API problems, as they do not provide API support.
We will make a best effort in keeping this documentation and associated language libraries up to date, but things might be outdated or missing.
If you find that something is no longer valid, please contact us on Discord or [create an issue](https://github.com/vrchatapi/specification/issues) and tell us so we can fix it.

# Getting Started

The VRChat API can be used to programmatically retrieve or update information regarding your profile, friends, avatars, worlds and more.
The API consists of two parts, \"Photon\" which is only used in-game, and the \"Web API\" which is used by both the game and the website.
This documentation focuses only on the Web API.

The API is designed around the REST ideology, providing semi-simple and usually predictable URIs to access and modify objects.
Requests support standard HTTP methods like GET, PUT, POST, and DELETE and standard status codes.
Response bodies are always UTF-8 encoded JSON objects, unless explicitly documented otherwise.

<div class=\"callout callout-error\">
  <strong>🛑 Warning! Do not touch Photon!</strong><br>
  Photon is only used by the in-game client and should <b>not</b> be touched. Doing so may result in permanent account termination.
</div>

<div class=\"callout callout-info\">
  <strong>ℹ️ Authentication</strong><br>
  Read <a href=\"#tag--authentication\">Authentication</a> for how to log in.
</div>

# Using the API

For simply exploring what the API can do it is strongly recommended to download [Insomnia](https://insomnia.rest/download), a free and open-source
API client that's great for sending requests to the API in an orderly fashion.
Insomnia allows you to send data in the format that's required for VRChat's API.
It is also possible to try out the API in your browser, by first logging in at [vrchat.com/home](https://vrchat.com/home/) and then going to
[vrchat.com/api/1/auth/user](https://vrchat.com/api/1/auth/user), but the information will be much harder to work with.

For more permanent operation such as software development it is instead recommended to use one of the existing language SDKs.
This community project maintains API libraries in several languages, which allows you to interact with the API with simple function calls
rather than having to implement the HTTP protocol yourself. Most of these libraries are automatically generated from the API specification,
sometimes with additional helpful wrapper code to make usage easier. This allows them to be almost automatically updated and expanded upon
as soon as a new feature is introduced in the specification itself. The libraries can be found on [GitHub](https://github.com/vrchatapi) or following:

* [NodeJS (JavaScript)](https://www.npmjs.com/package/vrchat)
* [Dart](https://pub.dev/packages/vrchat_dart)
* [Rust](https://crates.io/crates/vrchatapi)
* [C#](https://github.com/vrchatapi/vrchatapi-csharp)
* [Python](https://github.com/vrchatapi/vrchatapi-python)

# Pagination

Most endpoints enforce pagination, meaning they will only return 10 entries by default, and never more than 100.<br>
Using both the limit and offset parameters allows you to easily paginate through a large number of objects.

| Query Parameter | Type | Description |
| ----------|--|------- |
| `n` | integer  | The number of objects to return. This value often defaults to 10. Highest limit is always 100.|
| `offset` | integer  | A zero-based offset from the default object sorting.|

If a request returns fewer objects than the `limit` parameter, there are no more items available to return.

# Contribution

Do you want to get involved in the documentation effort? Do you want to help improve one of the language API libraries?
This project is an [OPEN Open Source Project](https://openopensource.org)! This means that individuals making significant and valuable contributions are given
commit-access to the project. It also means we are very open and welcoming of new people making contributions, unlike some more guarded open-source projects.

[![Discord](https://img.shields.io/static/v1?label=vrchatapi&message=discord&color=blueviolet&style=for-the-badge)](https://discord.gg/qjZE9C9fkB)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.20.7
- Package version: 1.0.0
- Generator version: 7.18.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/VRChatAPI](https://github.com/VRChatAPI)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import vrchatapi "github.com/StarLight-Oliver/vrchatapi-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `vrchatapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), vrchatapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `vrchatapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), vrchatapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `vrchatapi.ContextOperationServerIndices` and `vrchatapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), vrchatapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), vrchatapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.vrchat.cloud/api/1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthenticationAPI* | [**CancelPending2FA**](docs/AuthenticationAPI.md#cancelpending2fa) | **Delete** /auth/twofactorauth/totp/pending | Cancel pending enabling of time-based 2FA codes
*AuthenticationAPI* | [**CheckUserExists**](docs/AuthenticationAPI.md#checkuserexists) | **Get** /auth/exists | Check User Exists
*AuthenticationAPI* | [**ConfirmEmail**](docs/AuthenticationAPI.md#confirmemail) | **Get** /auth/confirmEmail | Confirm Email
*AuthenticationAPI* | [**CreateGlobalAvatarModeration**](docs/AuthenticationAPI.md#createglobalavatarmoderation) | **Post** /auth/user/avatarmoderations | Create Global Avatar Moderation
*AuthenticationAPI* | [**DeleteGlobalAvatarModeration**](docs/AuthenticationAPI.md#deleteglobalavatarmoderation) | **Delete** /auth/user/avatarmoderations | Delete Global Avatar Moderation
*AuthenticationAPI* | [**DeleteModerationReport**](docs/AuthenticationAPI.md#deletemoderationreport) | **Delete** /moderationReports/{moderationReportId} | Delete Moderation Report
*AuthenticationAPI* | [**DeleteUser**](docs/AuthenticationAPI.md#deleteuser) | **Put** /users/{userId}/delete | Delete User
*AuthenticationAPI* | [**Disable2FA**](docs/AuthenticationAPI.md#disable2fa) | **Delete** /auth/twofactorauth | Disable 2FA
*AuthenticationAPI* | [**Enable2FA**](docs/AuthenticationAPI.md#enable2fa) | **Post** /auth/twofactorauth/totp/pending | Enable time-based 2FA codes
*AuthenticationAPI* | [**GetCurrentUser**](docs/AuthenticationAPI.md#getcurrentuser) | **Get** /auth/user | Login and/or Get Current User Info
*AuthenticationAPI* | [**GetGlobalAvatarModerations**](docs/AuthenticationAPI.md#getglobalavatarmoderations) | **Get** /auth/user/avatarmoderations | Get Global Avatar Moderations
*AuthenticationAPI* | [**GetModerationReports**](docs/AuthenticationAPI.md#getmoderationreports) | **Get** /moderationReports | Get Moderation Reports
*AuthenticationAPI* | [**GetRecoveryCodes**](docs/AuthenticationAPI.md#getrecoverycodes) | **Get** /auth/user/twofactorauth/otp | Get 2FA Recovery codes
*AuthenticationAPI* | [**Logout**](docs/AuthenticationAPI.md#logout) | **Put** /logout | Logout
*AuthenticationAPI* | [**RegisterUserAccount**](docs/AuthenticationAPI.md#registeruseraccount) | **Post** /auth/register | Register User Account
*AuthenticationAPI* | [**ResendEmailConfirmation**](docs/AuthenticationAPI.md#resendemailconfirmation) | **Post** /auth/user/resendEmail | Resend Email Confirmation
*AuthenticationAPI* | [**SubmitModerationReport**](docs/AuthenticationAPI.md#submitmoderationreport) | **Post** /moderationReports | Submit Moderation Report
*AuthenticationAPI* | [**Verify2FA**](docs/AuthenticationAPI.md#verify2fa) | **Post** /auth/twofactorauth/totp/verify | Verify 2FA code
*AuthenticationAPI* | [**Verify2FAEmailCode**](docs/AuthenticationAPI.md#verify2faemailcode) | **Post** /auth/twofactorauth/emailotp/verify | Verify 2FA email code
*AuthenticationAPI* | [**VerifyAuthToken**](docs/AuthenticationAPI.md#verifyauthtoken) | **Get** /auth | Verify Auth Token
*AuthenticationAPI* | [**VerifyLoginPlace**](docs/AuthenticationAPI.md#verifyloginplace) | **Get** /auth/verifyLoginPlace | Verify Login Place
*AuthenticationAPI* | [**VerifyPending2FA**](docs/AuthenticationAPI.md#verifypending2fa) | **Post** /auth/twofactorauth/totp/pending/verify | Verify Pending 2FA code
*AuthenticationAPI* | [**VerifyRecoveryCode**](docs/AuthenticationAPI.md#verifyrecoverycode) | **Post** /auth/twofactorauth/otp/verify | Verify 2FA code with Recovery code
*AvatarsAPI* | [**CreateAvatar**](docs/AvatarsAPI.md#createavatar) | **Post** /avatars | Create Avatar
*AvatarsAPI* | [**DeleteAvatar**](docs/AvatarsAPI.md#deleteavatar) | **Delete** /avatars/{avatarId} | Delete Avatar
*AvatarsAPI* | [**DeleteImpostor**](docs/AvatarsAPI.md#deleteimpostor) | **Delete** /avatars/{avatarId}/impostor | Delete generated Impostor
*AvatarsAPI* | [**EnqueueImpostor**](docs/AvatarsAPI.md#enqueueimpostor) | **Post** /avatars/{avatarId}/impostor/enqueue | Enqueue Impostor generation
*AvatarsAPI* | [**GetAvatar**](docs/AvatarsAPI.md#getavatar) | **Get** /avatars/{avatarId} | Get Avatar
*AvatarsAPI* | [**GetAvatarStyles**](docs/AvatarsAPI.md#getavatarstyles) | **Get** /avatarStyles | Get Avatar Styles
*AvatarsAPI* | [**GetFavoritedAvatars**](docs/AvatarsAPI.md#getfavoritedavatars) | **Get** /avatars/favorites | List Favorited Avatars
*AvatarsAPI* | [**GetImpostorQueueStats**](docs/AvatarsAPI.md#getimpostorqueuestats) | **Get** /avatars/impostor/queue/stats | Get Impostor Queue Stats
*AvatarsAPI* | [**GetLicensedAvatars**](docs/AvatarsAPI.md#getlicensedavatars) | **Get** /avatars/licensed | List Licensed Avatars
*AvatarsAPI* | [**GetOwnAvatar**](docs/AvatarsAPI.md#getownavatar) | **Get** /users/{userId}/avatar | Get Own Avatar
*AvatarsAPI* | [**SearchAvatars**](docs/AvatarsAPI.md#searchavatars) | **Get** /avatars | Search Avatars
*AvatarsAPI* | [**SelectAvatar**](docs/AvatarsAPI.md#selectavatar) | **Put** /avatars/{avatarId}/select | Select Avatar
*AvatarsAPI* | [**SelectFallbackAvatar**](docs/AvatarsAPI.md#selectfallbackavatar) | **Put** /avatars/{avatarId}/selectFallback | Select Fallback Avatar
*AvatarsAPI* | [**UpdateAvatar**](docs/AvatarsAPI.md#updateavatar) | **Put** /avatars/{avatarId} | Update Avatar
*CalendarAPI* | [**CreateGroupCalendarEvent**](docs/CalendarAPI.md#creategroupcalendarevent) | **Post** /calendar/{groupId}/event | Create a calendar event
*CalendarAPI* | [**DeleteGroupCalendarEvent**](docs/CalendarAPI.md#deletegroupcalendarevent) | **Delete** /calendar/{groupId}/{calendarId} | Delete a calendar event
*CalendarAPI* | [**DiscoverCalendarEvents**](docs/CalendarAPI.md#discovercalendarevents) | **Get** /calendar/discover | Discover calendar events
*CalendarAPI* | [**FollowGroupCalendarEvent**](docs/CalendarAPI.md#followgroupcalendarevent) | **Post** /calendar/{groupId}/{calendarId}/follow | Follow a calendar event
*CalendarAPI* | [**GetCalendarEvents**](docs/CalendarAPI.md#getcalendarevents) | **Get** /calendar | List calendar events
*CalendarAPI* | [**GetFeaturedCalendarEvents**](docs/CalendarAPI.md#getfeaturedcalendarevents) | **Get** /calendar/featured | List featured calendar events
*CalendarAPI* | [**GetFollowedCalendarEvents**](docs/CalendarAPI.md#getfollowedcalendarevents) | **Get** /calendar/following | List followed calendar events
*CalendarAPI* | [**GetGroupCalendarEvent**](docs/CalendarAPI.md#getgroupcalendarevent) | **Get** /calendar/{groupId}/{calendarId} | Get a calendar event
*CalendarAPI* | [**GetGroupCalendarEventICS**](docs/CalendarAPI.md#getgroupcalendareventics) | **Get** /calendar/{groupId}/{calendarId}.ics | Download calendar event as ICS
*CalendarAPI* | [**GetGroupCalendarEvents**](docs/CalendarAPI.md#getgroupcalendarevents) | **Get** /calendar/{groupId} | List a group&#39;s calendar events
*CalendarAPI* | [**GetGroupNextCalendarEvent**](docs/CalendarAPI.md#getgroupnextcalendarevent) | **Get** /calendar/{groupId}/next | Get next calendar event
*CalendarAPI* | [**SearchCalendarEvents**](docs/CalendarAPI.md#searchcalendarevents) | **Get** /calendar/search | Search for calendar events
*CalendarAPI* | [**UpdateGroupCalendarEvent**](docs/CalendarAPI.md#updategroupcalendarevent) | **Put** /calendar/{groupId}/{calendarId}/event | Update a calendar event
*EconomyAPI* | [**GetActiveLicenses**](docs/EconomyAPI.md#getactivelicenses) | **Get** /economy/licenses/active | Get Active Licenses
*EconomyAPI* | [**GetBalance**](docs/EconomyAPI.md#getbalance) | **Get** /user/{userId}/balance | Get Balance
*EconomyAPI* | [**GetBalanceEarnings**](docs/EconomyAPI.md#getbalanceearnings) | **Get** /user/{userId}/balance/earnings | Get Balance Earnings
*EconomyAPI* | [**GetBulkGiftPurchases**](docs/EconomyAPI.md#getbulkgiftpurchases) | **Get** /user/bulk/gift/purchases | Get Bulk Gift Purchases
*EconomyAPI* | [**GetCurrentSubscriptions**](docs/EconomyAPI.md#getcurrentsubscriptions) | **Get** /auth/user/subscription | Get Current Subscriptions
*EconomyAPI* | [**GetEconomyAccount**](docs/EconomyAPI.md#geteconomyaccount) | **Get** /user/{userId}/economy/account | Get Economy Account
*EconomyAPI* | [**GetLicenseGroup**](docs/EconomyAPI.md#getlicensegroup) | **Get** /licenseGroups/{licenseGroupId} | Get License Group
*EconomyAPI* | [**GetProductListing**](docs/EconomyAPI.md#getproductlisting) | **Get** /listing/{productId} | Get Product Listing
*EconomyAPI* | [**GetProductListingAlternate**](docs/EconomyAPI.md#getproductlistingalternate) | **Get** /products/{productId} | Get Product Listing (alternate)
*EconomyAPI* | [**GetProductListings**](docs/EconomyAPI.md#getproductlistings) | **Get** /user/{userId}/listings | Get User Product Listings
*EconomyAPI* | [**GetProductPurchases**](docs/EconomyAPI.md#getproductpurchases) | **Get** /economy/purchases | Get Product Purchases
*EconomyAPI* | [**GetRecentSubscription**](docs/EconomyAPI.md#getrecentsubscription) | **Get** /user/subscription/recent | Get Recent Subscription
*EconomyAPI* | [**GetSteamTransaction**](docs/EconomyAPI.md#getsteamtransaction) | **Get** /Steam/transactions/{transactionId} | Get Steam Transaction
*EconomyAPI* | [**GetSteamTransactions**](docs/EconomyAPI.md#getsteamtransactions) | **Get** /Steam/transactions | List Steam Transactions
*EconomyAPI* | [**GetStore**](docs/EconomyAPI.md#getstore) | **Get** /economy/store | Get Store
*EconomyAPI* | [**GetStoreShelves**](docs/EconomyAPI.md#getstoreshelves) | **Get** /economy/store/shelves | Get Store Shelves
*EconomyAPI* | [**GetSubscriptions**](docs/EconomyAPI.md#getsubscriptions) | **Get** /subscriptions | List Subscriptions
*EconomyAPI* | [**GetTiliaStatus**](docs/EconomyAPI.md#gettiliastatus) | **Get** /tilia/status | Get Tilia Status
*EconomyAPI* | [**GetTiliaTos**](docs/EconomyAPI.md#gettiliatos) | **Get** /user/{userId}/tilia/tos | Get Tilia TOS Agreement Status
*EconomyAPI* | [**GetTokenBundles**](docs/EconomyAPI.md#gettokenbundles) | **Get** /tokenBundles | List Token Bundles
*EconomyAPI* | [**GetUserCreditsEligible**](docs/EconomyAPI.md#getusercreditseligible) | **Get** /users/{userId}/credits/eligible | Get User Credits Eligiblity
*EconomyAPI* | [**GetUserSubscriptionEligible**](docs/EconomyAPI.md#getusersubscriptioneligible) | **Get** /users/{userId}/subscription/eligible | Get User Subscription Eligiblity
*EconomyAPI* | [**PurchaseProductListing**](docs/EconomyAPI.md#purchaseproductlisting) | **Post** /economy/purchase/listing | Purchase Product Listing
*EconomyAPI* | [**UpdateTiliaTos**](docs/EconomyAPI.md#updatetiliatos) | **Put** /user/{userId}/tilia/tos | Update Tilia TOS Agreement Status
*FavoritesAPI* | [**AddFavorite**](docs/FavoritesAPI.md#addfavorite) | **Post** /favorites | Add Favorite
*FavoritesAPI* | [**ClearFavoriteGroup**](docs/FavoritesAPI.md#clearfavoritegroup) | **Delete** /favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId} | Clear Favorite Group
*FavoritesAPI* | [**GetFavoriteGroup**](docs/FavoritesAPI.md#getfavoritegroup) | **Get** /favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId} | Show Favorite Group
*FavoritesAPI* | [**GetFavoriteGroups**](docs/FavoritesAPI.md#getfavoritegroups) | **Get** /favorite/groups | List Favorite Groups
*FavoritesAPI* | [**GetFavoriteLimits**](docs/FavoritesAPI.md#getfavoritelimits) | **Get** /auth/user/favoritelimits | Get Favorite Limits
*FavoritesAPI* | [**GetFavorites**](docs/FavoritesAPI.md#getfavorites) | **Get** /favorites | List Favorites
*FavoritesAPI* | [**RemoveFavorite**](docs/FavoritesAPI.md#removefavorite) | **Delete** /favorites/{favoriteId} | Remove Favorite
*FavoritesAPI* | [**UpdateFavoriteGroup**](docs/FavoritesAPI.md#updatefavoritegroup) | **Put** /favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId} | Update Favorite Group
*FilesAPI* | [**CreateFile**](docs/FilesAPI.md#createfile) | **Post** /file | Create File
*FilesAPI* | [**CreateFileVersion**](docs/FilesAPI.md#createfileversion) | **Post** /file/{fileId} | Create File Version
*FilesAPI* | [**DeleteFile**](docs/FilesAPI.md#deletefile) | **Delete** /file/{fileId} | Delete File
*FilesAPI* | [**DeleteFileVersion**](docs/FilesAPI.md#deletefileversion) | **Delete** /file/{fileId}/{versionId} | Delete File Version
*FilesAPI* | [**DownloadFileVersion**](docs/FilesAPI.md#downloadfileversion) | **Get** /file/{fileId}/{versionId} | Download File Version
*FilesAPI* | [**FinishFileDataUpload**](docs/FilesAPI.md#finishfiledataupload) | **Put** /file/{fileId}/{versionId}/{fileType}/finish | Finish FileData Upload
*FilesAPI* | [**GetAdminAssetBundle**](docs/FilesAPI.md#getadminassetbundle) | **Get** /adminassetbundles/{adminAssetBundleId} | Get AdminAssetBundle
*FilesAPI* | [**GetContentAgreementStatus**](docs/FilesAPI.md#getcontentagreementstatus) | **Get** /agreement | Get Content Agreement Status
*FilesAPI* | [**GetFile**](docs/FilesAPI.md#getfile) | **Get** /file/{fileId} | Show File
*FilesAPI* | [**GetFileAnalysis**](docs/FilesAPI.md#getfileanalysis) | **Get** /analysis/{fileId}/{versionId} | Get File Version Analysis
*FilesAPI* | [**GetFileAnalysisSecurity**](docs/FilesAPI.md#getfileanalysissecurity) | **Get** /analysis/{fileId}/{versionId}/security | Get File Version Analysis Security
*FilesAPI* | [**GetFileAnalysisStandard**](docs/FilesAPI.md#getfileanalysisstandard) | **Get** /analysis/{fileId}/{versionId}/standard | Get File Version Analysis Standard
*FilesAPI* | [**GetFileDataUploadStatus**](docs/FilesAPI.md#getfiledatauploadstatus) | **Get** /file/{fileId}/{versionId}/{fileType}/status | Check FileData Upload Status
*FilesAPI* | [**GetFiles**](docs/FilesAPI.md#getfiles) | **Get** /files | List Files
*FilesAPI* | [**SetGroupGalleryFileOrder**](docs/FilesAPI.md#setgroupgalleryfileorder) | **Put** /files/order | Set Group Gallery File Order
*FilesAPI* | [**StartFileDataUpload**](docs/FilesAPI.md#startfiledataupload) | **Put** /file/{fileId}/{versionId}/{fileType}/start | Start FileData Upload
*FilesAPI* | [**SubmitContentAgreement**](docs/FilesAPI.md#submitcontentagreement) | **Post** /agreement | Submit Content Agreement
*FilesAPI* | [**UpdateAssetReviewNotes**](docs/FilesAPI.md#updateassetreviewnotes) | **Put** /assetReview/{assetReviewId}/notes | Update Asset Review Notes
*FilesAPI* | [**UploadGalleryImage**](docs/FilesAPI.md#uploadgalleryimage) | **Post** /gallery | Upload gallery image
*FilesAPI* | [**UploadIcon**](docs/FilesAPI.md#uploadicon) | **Post** /icon | Upload icon
*FilesAPI* | [**UploadImage**](docs/FilesAPI.md#uploadimage) | **Post** /file/image | Upload gallery image, icon, emoji or sticker
*FriendsAPI* | [**Boop**](docs/FriendsAPI.md#boop) | **Post** /users/{userId}/boop | Send Boop
*FriendsAPI* | [**DeleteFriendRequest**](docs/FriendsAPI.md#deletefriendrequest) | **Delete** /user/{userId}/friendRequest | Delete Friend Request
*FriendsAPI* | [**Friend**](docs/FriendsAPI.md#friend) | **Post** /user/{userId}/friendRequest | Send Friend Request
*FriendsAPI* | [**GetFriendStatus**](docs/FriendsAPI.md#getfriendstatus) | **Get** /user/{userId}/friendStatus | Check Friend Status
*FriendsAPI* | [**GetFriends**](docs/FriendsAPI.md#getfriends) | **Get** /auth/user/friends | List Friends
*FriendsAPI* | [**Unfriend**](docs/FriendsAPI.md#unfriend) | **Delete** /auth/user/friends/{userId} | Unfriend
*GroupsAPI* | [**AddGroupGalleryImage**](docs/GroupsAPI.md#addgroupgalleryimage) | **Post** /groups/{groupId}/galleries/{groupGalleryId}/images | Add Group Gallery Image
*GroupsAPI* | [**AddGroupMemberRole**](docs/GroupsAPI.md#addgroupmemberrole) | **Put** /groups/{groupId}/members/{userId}/roles/{groupRoleId} | Add Role to GroupMember
*GroupsAPI* | [**AddGroupPost**](docs/GroupsAPI.md#addgrouppost) | **Post** /groups/{groupId}/posts | Create a post in a Group
*GroupsAPI* | [**BanGroupMember**](docs/GroupsAPI.md#bangroupmember) | **Post** /groups/{groupId}/bans | Ban Group Member
*GroupsAPI* | [**BlockGroup**](docs/GroupsAPI.md#blockgroup) | **Post** /groups/{groupId}/block | Block Group
*GroupsAPI* | [**CancelGroupRequest**](docs/GroupsAPI.md#cancelgrouprequest) | **Delete** /groups/{groupId}/requests | Cancel Group Join Request
*GroupsAPI* | [**CancelGroupTransfer**](docs/GroupsAPI.md#cancelgrouptransfer) | **Delete** /groups/{groupId}/transfer | Cancel Group Transfer
*GroupsAPI* | [**CreateGroup**](docs/GroupsAPI.md#creategroup) | **Post** /groups | Create Group
*GroupsAPI* | [**CreateGroupAnnouncement**](docs/GroupsAPI.md#creategroupannouncement) | **Post** /groups/{groupId}/announcement | Create Group Announcement
*GroupsAPI* | [**CreateGroupGallery**](docs/GroupsAPI.md#creategroupgallery) | **Post** /groups/{groupId}/galleries | Create Group Gallery
*GroupsAPI* | [**CreateGroupInvite**](docs/GroupsAPI.md#creategroupinvite) | **Post** /groups/{groupId}/invites | Invite User to Group
*GroupsAPI* | [**CreateGroupRole**](docs/GroupsAPI.md#creategrouprole) | **Post** /groups/{groupId}/roles | Create GroupRole
*GroupsAPI* | [**DeclineGroupInvite**](docs/GroupsAPI.md#declinegroupinvite) | **Put** /groups/{groupId}/invites | Decline Invite from Group
*GroupsAPI* | [**DeleteGroup**](docs/GroupsAPI.md#deletegroup) | **Delete** /groups/{groupId} | Delete Group
*GroupsAPI* | [**DeleteGroupAnnouncement**](docs/GroupsAPI.md#deletegroupannouncement) | **Delete** /groups/{groupId}/announcement | Delete Group Announcement
*GroupsAPI* | [**DeleteGroupGallery**](docs/GroupsAPI.md#deletegroupgallery) | **Delete** /groups/{groupId}/galleries/{groupGalleryId} | Delete Group Gallery
*GroupsAPI* | [**DeleteGroupGalleryImage**](docs/GroupsAPI.md#deletegroupgalleryimage) | **Delete** /groups/{groupId}/galleries/{groupGalleryId}/images/{groupGalleryImageId} | Delete Group Gallery Image
*GroupsAPI* | [**DeleteGroupInvite**](docs/GroupsAPI.md#deletegroupinvite) | **Delete** /groups/{groupId}/invites/{userId} | Delete User Invite
*GroupsAPI* | [**DeleteGroupPost**](docs/GroupsAPI.md#deletegrouppost) | **Delete** /groups/{groupId}/posts/{notificationId} | Delete a Group post
*GroupsAPI* | [**DeleteGroupRole**](docs/GroupsAPI.md#deletegrouprole) | **Delete** /groups/{groupId}/roles/{groupRoleId} | Delete Group Role
*GroupsAPI* | [**GetGroup**](docs/GroupsAPI.md#getgroup) | **Get** /groups/{groupId} | Get Group by ID
*GroupsAPI* | [**GetGroupAnnouncements**](docs/GroupsAPI.md#getgroupannouncements) | **Get** /groups/{groupId}/announcement | Get Group Announcement
*GroupsAPI* | [**GetGroupAuditLogEntryTypes**](docs/GroupsAPI.md#getgroupauditlogentrytypes) | **Get** /groups/{groupId}/auditLogTypes | Get Group Audit Log Entry Types
*GroupsAPI* | [**GetGroupAuditLogs**](docs/GroupsAPI.md#getgroupauditlogs) | **Get** /groups/{groupId}/auditLogs | Get Group Audit Logs
*GroupsAPI* | [**GetGroupBans**](docs/GroupsAPI.md#getgroupbans) | **Get** /groups/{groupId}/bans | Get Group Bans
*GroupsAPI* | [**GetGroupGalleryImages**](docs/GroupsAPI.md#getgroupgalleryimages) | **Get** /groups/{groupId}/galleries/{groupGalleryId} | Get Group Gallery Images
*GroupsAPI* | [**GetGroupInstances**](docs/GroupsAPI.md#getgroupinstances) | **Get** /groups/{groupId}/instances | Get Group Instances
*GroupsAPI* | [**GetGroupInvites**](docs/GroupsAPI.md#getgroupinvites) | **Get** /groups/{groupId}/invites | Get Group Invites Sent
*GroupsAPI* | [**GetGroupMember**](docs/GroupsAPI.md#getgroupmember) | **Get** /groups/{groupId}/members/{userId} | Get Group Member
*GroupsAPI* | [**GetGroupMembers**](docs/GroupsAPI.md#getgroupmembers) | **Get** /groups/{groupId}/members | List Group Members
*GroupsAPI* | [**GetGroupPermissions**](docs/GroupsAPI.md#getgrouppermissions) | **Get** /groups/{groupId}/permissions | List Group Permissions
*GroupsAPI* | [**GetGroupPosts**](docs/GroupsAPI.md#getgroupposts) | **Get** /groups/{groupId}/posts | Get posts from a Group
*GroupsAPI* | [**GetGroupRequests**](docs/GroupsAPI.md#getgrouprequests) | **Get** /groups/{groupId}/requests | Get Group Join Requests
*GroupsAPI* | [**GetGroupRoleTemplates**](docs/GroupsAPI.md#getgrouproletemplates) | **Get** /groups/roleTemplates | Get Group Role Templates
*GroupsAPI* | [**GetGroupRoles**](docs/GroupsAPI.md#getgrouproles) | **Get** /groups/{groupId}/roles | Get Group Roles
*GroupsAPI* | [**GetGroupTransferability**](docs/GroupsAPI.md#getgrouptransferability) | **Get** /groups/{groupId}/transfer | Get Group Transferability
*GroupsAPI* | [**InitiateOrAcceptGroupTransfer**](docs/GroupsAPI.md#initiateoracceptgrouptransfer) | **Post** /groups/{groupId}/transfer | Initiate or Accept Group Transfer
*GroupsAPI* | [**JoinGroup**](docs/GroupsAPI.md#joingroup) | **Post** /groups/{groupId}/join | Join Group
*GroupsAPI* | [**KickGroupMember**](docs/GroupsAPI.md#kickgroupmember) | **Delete** /groups/{groupId}/members/{userId} | Kick Group Member
*GroupsAPI* | [**LeaveGroup**](docs/GroupsAPI.md#leavegroup) | **Post** /groups/{groupId}/leave | Leave Group
*GroupsAPI* | [**RemoveGroupMemberRole**](docs/GroupsAPI.md#removegroupmemberrole) | **Delete** /groups/{groupId}/members/{userId}/roles/{groupRoleId} | Remove Role from GroupMember
*GroupsAPI* | [**RespondGroupJoinRequest**](docs/GroupsAPI.md#respondgroupjoinrequest) | **Put** /groups/{groupId}/requests/{userId} | Respond Group Join request
*GroupsAPI* | [**SearchGroupMembers**](docs/GroupsAPI.md#searchgroupmembers) | **Get** /groups/{groupId}/members/search | Search Group Members
*GroupsAPI* | [**SearchGroups**](docs/GroupsAPI.md#searchgroups) | **Get** /groups | Search Group
*GroupsAPI* | [**UnbanGroupMember**](docs/GroupsAPI.md#unbangroupmember) | **Delete** /groups/{groupId}/bans/{userId} | Unban Group Member
*GroupsAPI* | [**UpdateGroup**](docs/GroupsAPI.md#updategroup) | **Put** /groups/{groupId} | Update Group
*GroupsAPI* | [**UpdateGroupGallery**](docs/GroupsAPI.md#updategroupgallery) | **Put** /groups/{groupId}/galleries/{groupGalleryId} | Update Group Gallery
*GroupsAPI* | [**UpdateGroupMember**](docs/GroupsAPI.md#updategroupmember) | **Put** /groups/{groupId}/members/{userId} | Update Group Member
*GroupsAPI* | [**UpdateGroupPost**](docs/GroupsAPI.md#updategrouppost) | **Put** /groups/{groupId}/posts/{notificationId} | Edits a Group post
*GroupsAPI* | [**UpdateGroupRepresentation**](docs/GroupsAPI.md#updategrouprepresentation) | **Put** /groups/{groupId}/representation | Update Group Representation
*GroupsAPI* | [**UpdateGroupRole**](docs/GroupsAPI.md#updategrouprole) | **Put** /groups/{groupId}/roles/{groupRoleId} | Update Group Role
*InstancesAPI* | [**CloseInstance**](docs/InstancesAPI.md#closeinstance) | **Delete** /instances/{worldId}:{instanceId} | Close Instance
*InstancesAPI* | [**CreateInstance**](docs/InstancesAPI.md#createinstance) | **Post** /instances | Create Instance
*InstancesAPI* | [**GetInstance**](docs/InstancesAPI.md#getinstance) | **Get** /instances/{worldId}:{instanceId} | Get Instance
*InstancesAPI* | [**GetInstanceByShortName**](docs/InstancesAPI.md#getinstancebyshortname) | **Get** /instances/s/{shortName} | Get Instance By Short Name
*InstancesAPI* | [**GetRecentLocations**](docs/InstancesAPI.md#getrecentlocations) | **Get** /instances/recent | List Recent Locations
*InstancesAPI* | [**GetShortName**](docs/InstancesAPI.md#getshortname) | **Get** /instances/{worldId}:{instanceId}/shortName | Get Instance Short Name
*InventoryAPI* | [**ConsumeOwnInventoryItem**](docs/InventoryAPI.md#consumeowninventoryitem) | **Put** /inventory/{inventoryItemId}/consume | Consume Own Inventory Item
*InventoryAPI* | [**DeleteOwnInventoryItem**](docs/InventoryAPI.md#deleteowninventoryitem) | **Delete** /inventory/{inventoryItemId} | Delete Own Inventory Item
*InventoryAPI* | [**EquipOwnInventoryItem**](docs/InventoryAPI.md#equipowninventoryitem) | **Put** /inventory/{inventoryItemId}/equip | Equip Own Inventory Item
*InventoryAPI* | [**GetInventory**](docs/InventoryAPI.md#getinventory) | **Get** /inventory | Get Inventory
*InventoryAPI* | [**GetInventoryCollections**](docs/InventoryAPI.md#getinventorycollections) | **Get** /inventory/collections | List Inventory Collections
*InventoryAPI* | [**GetInventoryDrops**](docs/InventoryAPI.md#getinventorydrops) | **Get** /inventory/drops | List Inventory Drops
*InventoryAPI* | [**GetInventoryTemplate**](docs/InventoryAPI.md#getinventorytemplate) | **Get** /inventory/template/{inventoryTemplateId} | Get Inventory Template
*InventoryAPI* | [**GetOwnInventoryItem**](docs/InventoryAPI.md#getowninventoryitem) | **Get** /inventory/{inventoryItemId} | Get Own Inventory Item
*InventoryAPI* | [**GetUserInventoryItem**](docs/InventoryAPI.md#getuserinventoryitem) | **Get** /user/{userId}/inventory/{inventoryItemId} | Get User Inventory Item
*InventoryAPI* | [**ShareInventoryItemDirect**](docs/InventoryAPI.md#shareinventoryitemdirect) | **Post** /inventory/cloning/direct | Share Inventory Item Direct
*InventoryAPI* | [**ShareInventoryItemPedestal**](docs/InventoryAPI.md#shareinventoryitempedestal) | **Get** /inventory/cloning/pedestal | Share Inventory Item by Pedestal
*InventoryAPI* | [**SpawnInventoryItem**](docs/InventoryAPI.md#spawninventoryitem) | **Get** /inventory/spawn | Spawn Inventory Item
*InventoryAPI* | [**UnequipOwnInventorySlot**](docs/InventoryAPI.md#unequipowninventoryslot) | **Delete** /inventory/{inventoryItemId}/equip | Unequip Own Inventory Slot
*InventoryAPI* | [**UpdateOwnInventoryItem**](docs/InventoryAPI.md#updateowninventoryitem) | **Put** /inventory/{inventoryItemId} | Update Own Inventory Item
*InviteAPI* | [**GetInviteMessage**](docs/InviteAPI.md#getinvitemessage) | **Get** /message/{userId}/{messageType}/{slot} | Get Invite Message
*InviteAPI* | [**GetInviteMessages**](docs/InviteAPI.md#getinvitemessages) | **Get** /message/{userId}/{messageType} | List Invite Messages
*InviteAPI* | [**InviteMyselfTo**](docs/InviteAPI.md#invitemyselfto) | **Post** /invite/myself/to/{worldId}:{instanceId} | Invite Myself To Instance
*InviteAPI* | [**InviteUser**](docs/InviteAPI.md#inviteuser) | **Post** /invite/{userId} | Invite User
*InviteAPI* | [**InviteUserWithPhoto**](docs/InviteAPI.md#inviteuserwithphoto) | **Post** /invite/{userId}/photo | Invite User with photo
*InviteAPI* | [**RequestInvite**](docs/InviteAPI.md#requestinvite) | **Post** /requestInvite/{userId} | Request Invite
*InviteAPI* | [**RequestInviteWithPhoto**](docs/InviteAPI.md#requestinvitewithphoto) | **Post** /requestInvite/{userId}/photo | Request Invite with photo
*InviteAPI* | [**ResetInviteMessage**](docs/InviteAPI.md#resetinvitemessage) | **Delete** /message/{userId}/{messageType}/{slot} | Reset Invite Message
*InviteAPI* | [**RespondInvite**](docs/InviteAPI.md#respondinvite) | **Post** /invite/{notificationId}/response | Respond Invite
*InviteAPI* | [**RespondInviteWithPhoto**](docs/InviteAPI.md#respondinvitewithphoto) | **Post** /invite/{notificationId}/response/photo | Respond Invite with photo
*InviteAPI* | [**UpdateInviteMessage**](docs/InviteAPI.md#updateinvitemessage) | **Put** /message/{userId}/{messageType}/{slot} | Update Invite Message
*JamsAPI* | [**GetJam**](docs/JamsAPI.md#getjam) | **Get** /jams/{jamId} | Show jam information
*JamsAPI* | [**GetJamSubmissions**](docs/JamsAPI.md#getjamsubmissions) | **Get** /jams/{jamId}/submissions | Show jam submissions
*JamsAPI* | [**GetJams**](docs/JamsAPI.md#getjams) | **Get** /jams | Show jams list
*MiscellaneousAPI* | [**GetAssignedPermissions**](docs/MiscellaneousAPI.md#getassignedpermissions) | **Get** /auth/permissions | Get Assigned Permissions
*MiscellaneousAPI* | [**GetCSS**](docs/MiscellaneousAPI.md#getcss) | **Get** /css/app.css | Download CSS
*MiscellaneousAPI* | [**GetConfig**](docs/MiscellaneousAPI.md#getconfig) | **Get** /config | Fetch API Config
*MiscellaneousAPI* | [**GetCurrentOnlineUsers**](docs/MiscellaneousAPI.md#getcurrentonlineusers) | **Get** /visits | Current Online Users
*MiscellaneousAPI* | [**GetHealth**](docs/MiscellaneousAPI.md#gethealth) | **Get** /health | Check API Health
*MiscellaneousAPI* | [**GetInfoPush**](docs/MiscellaneousAPI.md#getinfopush) | **Get** /infoPush | Show Information Notices
*MiscellaneousAPI* | [**GetJavaScript**](docs/MiscellaneousAPI.md#getjavascript) | **Get** /js/app.js | Download JavaScript
*MiscellaneousAPI* | [**GetPermission**](docs/MiscellaneousAPI.md#getpermission) | **Get** /permissions/{permissionId} | Get Permission
*MiscellaneousAPI* | [**GetSystemTime**](docs/MiscellaneousAPI.md#getsystemtime) | **Get** /time | Current System Time
*NotificationsAPI* | [**AcceptFriendRequest**](docs/NotificationsAPI.md#acceptfriendrequest) | **Put** /auth/user/notifications/{notificationId}/accept | Accept Friend Request
*NotificationsAPI* | [**AcknowledgeNotificationV2**](docs/NotificationsAPI.md#acknowledgenotificationv2) | **Post** /notifications/{notificationId}/see | Acknowledge NotificationV2
*NotificationsAPI* | [**ClearNotifications**](docs/NotificationsAPI.md#clearnotifications) | **Put** /auth/user/notifications/clear | Clear All Notifications
*NotificationsAPI* | [**DeleteAllNotificationV2s**](docs/NotificationsAPI.md#deleteallnotificationv2s) | **Delete** /notifications | Delete All NotificationV2s
*NotificationsAPI* | [**DeleteNotification**](docs/NotificationsAPI.md#deletenotification) | **Put** /auth/user/notifications/{notificationId}/hide | Delete Notification
*NotificationsAPI* | [**DeleteNotificationV2**](docs/NotificationsAPI.md#deletenotificationv2) | **Delete** /notifications/{notificationId} | Delete NotificationV2
*NotificationsAPI* | [**GetNotification**](docs/NotificationsAPI.md#getnotification) | **Get** /auth/user/notifications/{notificationId} | Show notification
*NotificationsAPI* | [**GetNotificationV2**](docs/NotificationsAPI.md#getnotificationv2) | **Get** /notifications/{notificationId} | Get NotificationV2
*NotificationsAPI* | [**GetNotificationV2s**](docs/NotificationsAPI.md#getnotificationv2s) | **Get** /notifications | List NotificationV2s
*NotificationsAPI* | [**GetNotifications**](docs/NotificationsAPI.md#getnotifications) | **Get** /auth/user/notifications | List Notifications
*NotificationsAPI* | [**MarkNotificationAsRead**](docs/NotificationsAPI.md#marknotificationasread) | **Put** /auth/user/notifications/{notificationId}/see | Mark Notification As Read
*NotificationsAPI* | [**ReplyNotificationV2**](docs/NotificationsAPI.md#replynotificationv2) | **Post** /notifications/{notificationId}/reply | Reply NotificationV2
*NotificationsAPI* | [**RespondNotificationV2**](docs/NotificationsAPI.md#respondnotificationv2) | **Post** /notifications/{notificationId}/respond | Respond NotificationV2
*PlayermoderationAPI* | [**ClearAllPlayerModerations**](docs/PlayermoderationAPI.md#clearallplayermoderations) | **Delete** /auth/user/playermoderations | Clear All Player Moderations
*PlayermoderationAPI* | [**GetPlayerModerations**](docs/PlayermoderationAPI.md#getplayermoderations) | **Get** /auth/user/playermoderations | Search Player Moderations
*PlayermoderationAPI* | [**ModerateUser**](docs/PlayermoderationAPI.md#moderateuser) | **Post** /auth/user/playermoderations | Moderate User
*PlayermoderationAPI* | [**UnmoderateUser**](docs/PlayermoderationAPI.md#unmoderateuser) | **Put** /auth/user/unplayermoderate | Unmoderate User
*PrintsAPI* | [**DeletePrint**](docs/PrintsAPI.md#deleteprint) | **Delete** /prints/{printId} | Delete Print
*PrintsAPI* | [**EditPrint**](docs/PrintsAPI.md#editprint) | **Post** /prints/{printId} | Edit Print
*PrintsAPI* | [**GetPrint**](docs/PrintsAPI.md#getprint) | **Get** /prints/{printId} | Get Print
*PrintsAPI* | [**GetUserPrints**](docs/PrintsAPI.md#getuserprints) | **Get** /prints/user/{userId} | Get Own Prints
*PrintsAPI* | [**UploadPrint**](docs/PrintsAPI.md#uploadprint) | **Post** /prints | Upload Print
*PropsAPI* | [**CreateProp**](docs/PropsAPI.md#createprop) | **Post** /props | Create Prop
*PropsAPI* | [**DeleteProp**](docs/PropsAPI.md#deleteprop) | **Delete** /props/{propId} | Delete Prop
*PropsAPI* | [**GetProp**](docs/PropsAPI.md#getprop) | **Get** /props/{propId} | Get Prop
*PropsAPI* | [**GetPropPublishStatus**](docs/PropsAPI.md#getproppublishstatus) | **Get** /props/{propId}/publish | Get Prop Publish Status
*PropsAPI* | [**ListProps**](docs/PropsAPI.md#listprops) | **Get** /props | List Props
*PropsAPI* | [**PublishProp**](docs/PropsAPI.md#publishprop) | **Put** /props/{propId}/publish | Publish Prop
*PropsAPI* | [**UnpublishProp**](docs/PropsAPI.md#unpublishprop) | **Delete** /props/{propId}/publish | Unpublish Prop
*PropsAPI* | [**UpdateProp**](docs/PropsAPI.md#updateprop) | **Put** /props/{propId} | Update Prop
*UsersAPI* | [**AddTags**](docs/UsersAPI.md#addtags) | **Post** /users/{userId}/addTags | Add User Tags
*UsersAPI* | [**CheckUserPersistenceExists**](docs/UsersAPI.md#checkuserpersistenceexists) | **Get** /users/{userId}/{worldId}/persist/exists | Check User Persistence Exists
*UsersAPI* | [**DeleteAllUserPersistenceData**](docs/UsersAPI.md#deletealluserpersistencedata) | **Delete** /users/{userId}/persist | Delete All User Persistence Data
*UsersAPI* | [**DeleteUserPersistence**](docs/UsersAPI.md#deleteuserpersistence) | **Delete** /users/{userId}/{worldId}/persist | Delete User Persistence
*UsersAPI* | [**GetBlockedGroups**](docs/UsersAPI.md#getblockedgroups) | **Get** /users/{userId}/groups/userblocked | Get User Group Blocks
*UsersAPI* | [**GetInvitedGroups**](docs/UsersAPI.md#getinvitedgroups) | **Get** /users/{userId}/groups/invited | Get User Group Invited
*UsersAPI* | [**GetMutualFriends**](docs/UsersAPI.md#getmutualfriends) | **Get** /users/{userId}/mutuals/friends | Get User Mutual Friends
*UsersAPI* | [**GetMutualGroups**](docs/UsersAPI.md#getmutualgroups) | **Get** /users/{userId}/mutuals/groups | Get User Mutual Groups
*UsersAPI* | [**GetMutuals**](docs/UsersAPI.md#getmutuals) | **Get** /users/{userId}/mutuals | Get User Mutuals
*UsersAPI* | [**GetUser**](docs/UsersAPI.md#getuser) | **Get** /users/{userId} | Get User by ID
*UsersAPI* | [**GetUserAllGroupPermissions**](docs/UsersAPI.md#getuserallgrouppermissions) | **Get** /users/{userId}/groups/permissions | Get user&#39;s permissions for all joined groups.
*UsersAPI* | [**GetUserByName**](docs/UsersAPI.md#getuserbyname) | **Get** /users/{username}/name | Get User by Username
*UsersAPI* | [**GetUserFeedback**](docs/UsersAPI.md#getuserfeedback) | **Get** /users/{userId}/feedback | Get User Feedback
*UsersAPI* | [**GetUserGroupInstances**](docs/UsersAPI.md#getusergroupinstances) | **Get** /users/{userId}/instances/groups | Get User Group Instances
*UsersAPI* | [**GetUserGroupInstancesForGroup**](docs/UsersAPI.md#getusergroupinstancesforgroup) | **Get** /users/{userId}/instances/groups/{groupId} | Get User Group Instances for a specific Group
*UsersAPI* | [**GetUserGroupRequests**](docs/UsersAPI.md#getusergrouprequests) | **Get** /users/{userId}/groups/requested | Get User Group Requests
*UsersAPI* | [**GetUserGroups**](docs/UsersAPI.md#getusergroups) | **Get** /users/{userId}/groups | Get User Groups
*UsersAPI* | [**GetUserNote**](docs/UsersAPI.md#getusernote) | **Get** /userNotes/{userNoteId} | Get User Note
*UsersAPI* | [**GetUserNotes**](docs/UsersAPI.md#getusernotes) | **Get** /userNotes | Get User Notes
*UsersAPI* | [**GetUserRepresentedGroup**](docs/UsersAPI.md#getuserrepresentedgroup) | **Get** /users/{userId}/groups/represented | Get user&#39;s current represented group
*UsersAPI* | [**RemoveTags**](docs/UsersAPI.md#removetags) | **Post** /users/{userId}/removeTags | Remove User Tags
*UsersAPI* | [**SearchUsers**](docs/UsersAPI.md#searchusers) | **Get** /users | Search All Users
*UsersAPI* | [**UpdateBadge**](docs/UsersAPI.md#updatebadge) | **Put** /users/{userId}/badges/{badgeId} | Update User Badge
*UsersAPI* | [**UpdateUser**](docs/UsersAPI.md#updateuser) | **Put** /users/{userId} | Update User Info
*UsersAPI* | [**UpdateUserNote**](docs/UsersAPI.md#updateusernote) | **Post** /userNotes | Update User Note
*WorldsAPI* | [**CheckUserPersistenceExists**](docs/WorldsAPI.md#checkuserpersistenceexists) | **Get** /users/{userId}/{worldId}/persist/exists | Check User Persistence Exists
*WorldsAPI* | [**CreateWorld**](docs/WorldsAPI.md#createworld) | **Post** /worlds | Create World
*WorldsAPI* | [**DeleteAllUserPersistenceData**](docs/WorldsAPI.md#deletealluserpersistencedata) | **Delete** /users/{userId}/persist | Delete All User Persistence Data
*WorldsAPI* | [**DeleteUserPersistence**](docs/WorldsAPI.md#deleteuserpersistence) | **Delete** /users/{userId}/{worldId}/persist | Delete User Persistence
*WorldsAPI* | [**DeleteWorld**](docs/WorldsAPI.md#deleteworld) | **Delete** /worlds/{worldId} | Delete World
*WorldsAPI* | [**GetActiveWorlds**](docs/WorldsAPI.md#getactiveworlds) | **Get** /worlds/active | List Active Worlds
*WorldsAPI* | [**GetFavoritedWorlds**](docs/WorldsAPI.md#getfavoritedworlds) | **Get** /worlds/favorites | List Favorited Worlds
*WorldsAPI* | [**GetRecentWorlds**](docs/WorldsAPI.md#getrecentworlds) | **Get** /worlds/recent | List Recent Worlds
*WorldsAPI* | [**GetWorld**](docs/WorldsAPI.md#getworld) | **Get** /worlds/{worldId} | Get World by ID
*WorldsAPI* | [**GetWorldInstance**](docs/WorldsAPI.md#getworldinstance) | **Get** /worlds/{worldId}/{instanceId} | Get World Instance
*WorldsAPI* | [**GetWorldMetadata**](docs/WorldsAPI.md#getworldmetadata) | **Get** /worlds/{worldId}/metadata | Get World Metadata
*WorldsAPI* | [**GetWorldPublishStatus**](docs/WorldsAPI.md#getworldpublishstatus) | **Get** /worlds/{worldId}/publish | Get World Publish Status
*WorldsAPI* | [**PublishWorld**](docs/WorldsAPI.md#publishworld) | **Put** /worlds/{worldId}/publish | Publish World
*WorldsAPI* | [**SearchWorlds**](docs/WorldsAPI.md#searchworlds) | **Get** /worlds | Search All Worlds
*WorldsAPI* | [**UnpublishWorld**](docs/WorldsAPI.md#unpublishworld) | **Delete** /worlds/{worldId}/publish | Unpublish World
*WorldsAPI* | [**UpdateWorld**](docs/WorldsAPI.md#updateworld) | **Put** /worlds/{worldId} | Update World


## Documentation For Models

 - [APIConfig](docs/APIConfig.md)
 - [APIConfigAccessLogsUrls](docs/APIConfigAccessLogsUrls.md)
 - [APIConfigAnnouncement](docs/APIConfigAnnouncement.md)
 - [APIConfigAudioConfig](docs/APIConfigAudioConfig.md)
 - [APIConfigAvatarPerfLimiter](docs/APIConfigAvatarPerfLimiter.md)
 - [APIConfigConstants](docs/APIConfigConstants.md)
 - [APIConfigConstantsGROUPS](docs/APIConfigConstantsGROUPS.md)
 - [APIConfigConstantsINSTANCE](docs/APIConfigConstantsINSTANCE.md)
 - [APIConfigConstantsINSTANCEPOPULATIONBRACKETS](docs/APIConfigConstantsINSTANCEPOPULATIONBRACKETS.md)
 - [APIConfigConstantsINSTANCEPOPULATIONBRACKETSCROWDED](docs/APIConfigConstantsINSTANCEPOPULATIONBRACKETSCROWDED.md)
 - [APIConfigConstantsINSTANCEPOPULATIONBRACKETSFEW](docs/APIConfigConstantsINSTANCEPOPULATIONBRACKETSFEW.md)
 - [APIConfigConstantsINSTANCEPOPULATIONBRACKETSMANY](docs/APIConfigConstantsINSTANCEPOPULATIONBRACKETSMANY.md)
 - [APIConfigConstantsLANGUAGE](docs/APIConfigConstantsLANGUAGE.md)
 - [APIConfigDownloadURLList](docs/APIConfigDownloadURLList.md)
 - [APIConfigEvents](docs/APIConfigEvents.md)
 - [APIConfigIosVersion](docs/APIConfigIosVersion.md)
 - [APIConfigMinSupportedClientBuildNumber](docs/APIConfigMinSupportedClientBuildNumber.md)
 - [APIConfigOfflineAnalysis](docs/APIConfigOfflineAnalysis.md)
 - [APIHealth](docs/APIHealth.md)
 - [AccountDeletionLog](docs/AccountDeletionLog.md)
 - [AddFavoriteRequest](docs/AddFavoriteRequest.md)
 - [AddGroupGalleryImageRequest](docs/AddGroupGalleryImageRequest.md)
 - [AdminAssetBundle](docs/AdminAssetBundle.md)
 - [AdminUnityPackage](docs/AdminUnityPackage.md)
 - [AgeVerificationStatus](docs/AgeVerificationStatus.md)
 - [Agreement](docs/Agreement.md)
 - [AgreementCode](docs/AgreementCode.md)
 - [AgreementRequest](docs/AgreementRequest.md)
 - [AgreementStatus](docs/AgreementStatus.md)
 - [Avatar](docs/Avatar.md)
 - [AvatarModeration](docs/AvatarModeration.md)
 - [AvatarModerationCreated](docs/AvatarModerationCreated.md)
 - [AvatarModerationType](docs/AvatarModerationType.md)
 - [AvatarPerformance](docs/AvatarPerformance.md)
 - [AvatarPublishedListingsInner](docs/AvatarPublishedListingsInner.md)
 - [AvatarStyle](docs/AvatarStyle.md)
 - [AvatarStyles](docs/AvatarStyles.md)
 - [AvatarUnityPackageUrlObject](docs/AvatarUnityPackageUrlObject.md)
 - [Badge](docs/Badge.md)
 - [Balance](docs/Balance.md)
 - [BanGroupMemberRequest](docs/BanGroupMemberRequest.md)
 - [BoopRequest](docs/BoopRequest.md)
 - [CalendarEvent](docs/CalendarEvent.md)
 - [CalendarEventAccess](docs/CalendarEventAccess.md)
 - [CalendarEventCategory](docs/CalendarEventCategory.md)
 - [CalendarEventDiscovery](docs/CalendarEventDiscovery.md)
 - [CalendarEventDiscoveryInclusion](docs/CalendarEventDiscoveryInclusion.md)
 - [CalendarEventDiscoveryScope](docs/CalendarEventDiscoveryScope.md)
 - [CalendarEventPlatform](docs/CalendarEventPlatform.md)
 - [CalendarEventUserInterest](docs/CalendarEventUserInterest.md)
 - [ChangeUserTagsRequest](docs/ChangeUserTagsRequest.md)
 - [ContentFilter](docs/ContentFilter.md)
 - [CreateAvatarModerationRequest](docs/CreateAvatarModerationRequest.md)
 - [CreateAvatarRequest](docs/CreateAvatarRequest.md)
 - [CreateCalendarEventRequest](docs/CreateCalendarEventRequest.md)
 - [CreateFileRequest](docs/CreateFileRequest.md)
 - [CreateFileVersionRequest](docs/CreateFileVersionRequest.md)
 - [CreateGroupAnnouncementRequest](docs/CreateGroupAnnouncementRequest.md)
 - [CreateGroupGalleryRequest](docs/CreateGroupGalleryRequest.md)
 - [CreateGroupInviteRequest](docs/CreateGroupInviteRequest.md)
 - [CreateGroupPostRequest](docs/CreateGroupPostRequest.md)
 - [CreateGroupRequest](docs/CreateGroupRequest.md)
 - [CreateGroupRoleRequest](docs/CreateGroupRoleRequest.md)
 - [CreateInstanceRequest](docs/CreateInstanceRequest.md)
 - [CreatePropRequest](docs/CreatePropRequest.md)
 - [CreateWorldRequest](docs/CreateWorldRequest.md)
 - [CurrentUser](docs/CurrentUser.md)
 - [CurrentUserPlatformHistoryInner](docs/CurrentUserPlatformHistoryInner.md)
 - [CurrentUserPresence](docs/CurrentUserPresence.md)
 - [DeclineGroupInviteRequest](docs/DeclineGroupInviteRequest.md)
 - [DeveloperType](docs/DeveloperType.md)
 - [Disable2FAResult](docs/Disable2FAResult.md)
 - [DiscordDetails](docs/DiscordDetails.md)
 - [DynamicContentRow](docs/DynamicContentRow.md)
 - [EconomyAccount](docs/EconomyAccount.md)
 - [EquipInventoryItemRequest](docs/EquipInventoryItemRequest.md)
 - [Error](docs/Error.md)
 - [Favorite](docs/Favorite.md)
 - [FavoriteGroup](docs/FavoriteGroup.md)
 - [FavoriteGroupLimits](docs/FavoriteGroupLimits.md)
 - [FavoriteGroupVisibility](docs/FavoriteGroupVisibility.md)
 - [FavoriteLimits](docs/FavoriteLimits.md)
 - [FavoriteType](docs/FavoriteType.md)
 - [FavoritedWorld](docs/FavoritedWorld.md)
 - [Feedback](docs/Feedback.md)
 - [File](docs/File.md)
 - [FileAnalysis](docs/FileAnalysis.md)
 - [FileAnalysisAvatarStats](docs/FileAnalysisAvatarStats.md)
 - [FileData](docs/FileData.md)
 - [FileStatus](docs/FileStatus.md)
 - [FileUploadURL](docs/FileUploadURL.md)
 - [FileVersion](docs/FileVersion.md)
 - [FileVersionUploadStatus](docs/FileVersionUploadStatus.md)
 - [FinishFileDataUploadRequest](docs/FinishFileDataUploadRequest.md)
 - [FollowCalendarEventRequest](docs/FollowCalendarEventRequest.md)
 - [FriendStatus](docs/FriendStatus.md)
 - [GetGroupPosts200Response](docs/GetGroupPosts200Response.md)
 - [GetUserGroupInstances200Response](docs/GetUserGroupInstances200Response.md)
 - [Group](docs/Group.md)
 - [GroupAccessType](docs/GroupAccessType.md)
 - [GroupAnnouncement](docs/GroupAnnouncement.md)
 - [GroupAuditLogEntry](docs/GroupAuditLogEntry.md)
 - [GroupGallery](docs/GroupGallery.md)
 - [GroupGalleryFileOrder](docs/GroupGalleryFileOrder.md)
 - [GroupGalleryFileOrderRequest](docs/GroupGalleryFileOrderRequest.md)
 - [GroupGalleryImage](docs/GroupGalleryImage.md)
 - [GroupInstance](docs/GroupInstance.md)
 - [GroupJoinRequestAction](docs/GroupJoinRequestAction.md)
 - [GroupJoinState](docs/GroupJoinState.md)
 - [GroupLimitedMember](docs/GroupLimitedMember.md)
 - [GroupMember](docs/GroupMember.md)
 - [GroupMemberLimitedUser](docs/GroupMemberLimitedUser.md)
 - [GroupMemberStatus](docs/GroupMemberStatus.md)
 - [GroupMyMember](docs/GroupMyMember.md)
 - [GroupPermission](docs/GroupPermission.md)
 - [GroupPermissions](docs/GroupPermissions.md)
 - [GroupPost](docs/GroupPost.md)
 - [GroupPostVisibility](docs/GroupPostVisibility.md)
 - [GroupPrivacy](docs/GroupPrivacy.md)
 - [GroupRole](docs/GroupRole.md)
 - [GroupRoleTemplate](docs/GroupRoleTemplate.md)
 - [GroupRoleTemplateValues](docs/GroupRoleTemplateValues.md)
 - [GroupRoleTemplateValuesRoles](docs/GroupRoleTemplateValuesRoles.md)
 - [GroupSearchSort](docs/GroupSearchSort.md)
 - [GroupTransferable](docs/GroupTransferable.md)
 - [GroupTransferableRequirements](docs/GroupTransferableRequirements.md)
 - [GroupUserVisibility](docs/GroupUserVisibility.md)
 - [ImageAnimationStyle](docs/ImageAnimationStyle.md)
 - [ImageLoopStyle](docs/ImageLoopStyle.md)
 - [ImageMask](docs/ImageMask.md)
 - [ImagePurpose](docs/ImagePurpose.md)
 - [InfoPush](docs/InfoPush.md)
 - [InfoPushData](docs/InfoPushData.md)
 - [InfoPushDataArticle](docs/InfoPushDataArticle.md)
 - [InfoPushDataArticleContent](docs/InfoPushDataArticleContent.md)
 - [InfoPushDataClickable](docs/InfoPushDataClickable.md)
 - [Instance](docs/Instance.md)
 - [InstanceContentSettings](docs/InstanceContentSettings.md)
 - [InstancePlatforms](docs/InstancePlatforms.md)
 - [InstanceRegion](docs/InstanceRegion.md)
 - [InstanceShortNameResponse](docs/InstanceShortNameResponse.md)
 - [InstanceType](docs/InstanceType.md)
 - [Inventory](docs/Inventory.md)
 - [InventoryConsumptionResults](docs/InventoryConsumptionResults.md)
 - [InventoryDefaultAttributesValue](docs/InventoryDefaultAttributesValue.md)
 - [InventoryDefaultAttributesValueValidator](docs/InventoryDefaultAttributesValueValidator.md)
 - [InventoryDrop](docs/InventoryDrop.md)
 - [InventoryEquipSlot](docs/InventoryEquipSlot.md)
 - [InventoryFlag](docs/InventoryFlag.md)
 - [InventoryItem](docs/InventoryItem.md)
 - [InventoryItemType](docs/InventoryItemType.md)
 - [InventoryMetadata](docs/InventoryMetadata.md)
 - [InventoryNotificationDetails](docs/InventoryNotificationDetails.md)
 - [InventorySpawn](docs/InventorySpawn.md)
 - [InventoryTemplate](docs/InventoryTemplate.md)
 - [InventoryUserAttributes](docs/InventoryUserAttributes.md)
 - [InviteMessage](docs/InviteMessage.md)
 - [InviteMessageType](docs/InviteMessageType.md)
 - [InviteRequest](docs/InviteRequest.md)
 - [InviteResponse](docs/InviteResponse.md)
 - [Jam](docs/Jam.md)
 - [JamStateChangeDates](docs/JamStateChangeDates.md)
 - [JoinGroupRequest](docs/JoinGroupRequest.md)
 - [License](docs/License.md)
 - [LicenseAction](docs/LicenseAction.md)
 - [LicenseGroup](docs/LicenseGroup.md)
 - [LicenseType](docs/LicenseType.md)
 - [LimitedGroup](docs/LimitedGroup.md)
 - [LimitedUnityPackage](docs/LimitedUnityPackage.md)
 - [LimitedUserFriend](docs/LimitedUserFriend.md)
 - [LimitedUserGroups](docs/LimitedUserGroups.md)
 - [LimitedUserInstance](docs/LimitedUserInstance.md)
 - [LimitedUserSearch](docs/LimitedUserSearch.md)
 - [LimitedWorld](docs/LimitedWorld.md)
 - [MIMEType](docs/MIMEType.md)
 - [ModerateUserRequest](docs/ModerateUserRequest.md)
 - [ModerationReport](docs/ModerationReport.md)
 - [MutualFriend](docs/MutualFriend.md)
 - [Mutuals](docs/Mutuals.md)
 - [Notification](docs/Notification.md)
 - [NotificationDetailBoop](docs/NotificationDetailBoop.md)
 - [NotificationDetailInvite](docs/NotificationDetailInvite.md)
 - [NotificationDetailInviteResponse](docs/NotificationDetailInviteResponse.md)
 - [NotificationDetailRequestInvite](docs/NotificationDetailRequestInvite.md)
 - [NotificationDetailRequestInviteResponse](docs/NotificationDetailRequestInviteResponse.md)
 - [NotificationDetailVoteToKick](docs/NotificationDetailVoteToKick.md)
 - [NotificationType](docs/NotificationType.md)
 - [NotificationV2](docs/NotificationV2.md)
 - [NotificationV2Data](docs/NotificationV2Data.md)
 - [NotificationV2DataBadgeEarned](docs/NotificationV2DataBadgeEarned.md)
 - [NotificationV2DataBoop](docs/NotificationV2DataBoop.md)
 - [NotificationV2DataEventAnnouncement](docs/NotificationV2DataEventAnnouncement.md)
 - [NotificationV2DataGroupAnnouncement](docs/NotificationV2DataGroupAnnouncement.md)
 - [NotificationV2DataGroupInformative](docs/NotificationV2DataGroupInformative.md)
 - [NotificationV2DataGroupTransfer](docs/NotificationV2DataGroupTransfer.md)
 - [NotificationV2DetailsBoop](docs/NotificationV2DetailsBoop.md)
 - [NotificationV2Response](docs/NotificationV2Response.md)
 - [NotificationV2Type](docs/NotificationV2Type.md)
 - [OkStatus](docs/OkStatus.md)
 - [OkStatus2](docs/OkStatus2.md)
 - [OrderOption](docs/OrderOption.md)
 - [OrderOptionShort](docs/OrderOptionShort.md)
 - [PaginatedCalendarEventList](docs/PaginatedCalendarEventList.md)
 - [PaginatedGroupAuditLogEntryList](docs/PaginatedGroupAuditLogEntryList.md)
 - [PaginatedModerationReportList](docs/PaginatedModerationReportList.md)
 - [PastDisplayName](docs/PastDisplayName.md)
 - [Pending2FAResult](docs/Pending2FAResult.md)
 - [PerformanceLimiterInfo](docs/PerformanceLimiterInfo.md)
 - [PerformanceRatings](docs/PerformanceRatings.md)
 - [Permission](docs/Permission.md)
 - [PermissionData](docs/PermissionData.md)
 - [PlatformBuildInfo](docs/PlatformBuildInfo.md)
 - [PlayerModeration](docs/PlayerModeration.md)
 - [PlayerModerationType](docs/PlayerModerationType.md)
 - [Print](docs/Print.md)
 - [PrintFiles](docs/PrintFiles.md)
 - [Product](docs/Product.md)
 - [ProductListing](docs/ProductListing.md)
 - [ProductListingType](docs/ProductListingType.md)
 - [ProductListingVariant](docs/ProductListingVariant.md)
 - [ProductPurchase](docs/ProductPurchase.md)
 - [ProductPurchasePurchaseContext](docs/ProductPurchasePurchaseContext.md)
 - [ProductType](docs/ProductType.md)
 - [Prop](docs/Prop.md)
 - [PropPublishStatus](docs/PropPublishStatus.md)
 - [PropUnityPackage](docs/PropUnityPackage.md)
 - [PurchaseProductListingRequest](docs/PurchaseProductListingRequest.md)
 - [Region](docs/Region.md)
 - [RegisterUserAccount200Response](docs/RegisterUserAccount200Response.md)
 - [RegisterUserAccountRequest](docs/RegisterUserAccountRequest.md)
 - [ReleaseStatus](docs/ReleaseStatus.md)
 - [ReportCategory](docs/ReportCategory.md)
 - [ReportReason](docs/ReportReason.md)
 - [RepresentedGroup](docs/RepresentedGroup.md)
 - [RequestInviteRequest](docs/RequestInviteRequest.md)
 - [RequiresTwoFactorAuth](docs/RequiresTwoFactorAuth.md)
 - [RespondGroupJoinRequest](docs/RespondGroupJoinRequest.md)
 - [RespondNotificationV2Request](docs/RespondNotificationV2Request.md)
 - [Response](docs/Response.md)
 - [SearchGroupMembers200Response](docs/SearchGroupMembers200Response.md)
 - [SentNotification](docs/SentNotification.md)
 - [SentNotificationDetails](docs/SentNotificationDetails.md)
 - [ServiceQueueStats](docs/ServiceQueueStats.md)
 - [ServiceStatus](docs/ServiceStatus.md)
 - [ShareInventoryItemDirectRequest](docs/ShareInventoryItemDirectRequest.md)
 - [SortOption](docs/SortOption.md)
 - [SortOptionProductPurchase](docs/SortOptionProductPurchase.md)
 - [Store](docs/Store.md)
 - [StoreShelf](docs/StoreShelf.md)
 - [StoreType](docs/StoreType.md)
 - [StoreView](docs/StoreView.md)
 - [Submission](docs/Submission.md)
 - [SubmitModerationReportRequest](docs/SubmitModerationReportRequest.md)
 - [SubmitModerationReportRequestDetails](docs/SubmitModerationReportRequestDetails.md)
 - [Subscription](docs/Subscription.md)
 - [SubscriptionPeriod](docs/SubscriptionPeriod.md)
 - [Success](docs/Success.md)
 - [SuccessFlag](docs/SuccessFlag.md)
 - [TiliaStatus](docs/TiliaStatus.md)
 - [TiliaTOS](docs/TiliaTOS.md)
 - [TokenBundle](docs/TokenBundle.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionAgreement](docs/TransactionAgreement.md)
 - [TransactionStatus](docs/TransactionStatus.md)
 - [TransactionSteamInfo](docs/TransactionSteamInfo.md)
 - [TransactionSteamWalletInfo](docs/TransactionSteamWalletInfo.md)
 - [TransferGroupRequest](docs/TransferGroupRequest.md)
 - [TwoFactorAuthCode](docs/TwoFactorAuthCode.md)
 - [TwoFactorAuthType](docs/TwoFactorAuthType.md)
 - [TwoFactorEmailCode](docs/TwoFactorEmailCode.md)
 - [TwoFactorRecoveryCodes](docs/TwoFactorRecoveryCodes.md)
 - [TwoFactorRecoveryCodesOtpInner](docs/TwoFactorRecoveryCodesOtpInner.md)
 - [UnityPackage](docs/UnityPackage.md)
 - [UpdateAssetReviewNotesRequest](docs/UpdateAssetReviewNotesRequest.md)
 - [UpdateAvatarRequest](docs/UpdateAvatarRequest.md)
 - [UpdateCalendarEventRequest](docs/UpdateCalendarEventRequest.md)
 - [UpdateFavoriteGroupRequest](docs/UpdateFavoriteGroupRequest.md)
 - [UpdateGroupGalleryRequest](docs/UpdateGroupGalleryRequest.md)
 - [UpdateGroupMemberRequest](docs/UpdateGroupMemberRequest.md)
 - [UpdateGroupRepresentationRequest](docs/UpdateGroupRepresentationRequest.md)
 - [UpdateGroupRequest](docs/UpdateGroupRequest.md)
 - [UpdateGroupRoleRequest](docs/UpdateGroupRoleRequest.md)
 - [UpdateInventoryItemRequest](docs/UpdateInventoryItemRequest.md)
 - [UpdateInviteMessageRequest](docs/UpdateInviteMessageRequest.md)
 - [UpdatePropRequest](docs/UpdatePropRequest.md)
 - [UpdateTiliaTOSRequest](docs/UpdateTiliaTOSRequest.md)
 - [UpdateUserBadgeRequest](docs/UpdateUserBadgeRequest.md)
 - [UpdateUserNoteRequest](docs/UpdateUserNoteRequest.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [UpdateWorldRequest](docs/UpdateWorldRequest.md)
 - [User](docs/User.md)
 - [UserCreditsEligible](docs/UserCreditsEligible.md)
 - [UserExists](docs/UserExists.md)
 - [UserNote](docs/UserNote.md)
 - [UserNoteTargetUser](docs/UserNoteTargetUser.md)
 - [UserState](docs/UserState.md)
 - [UserStatus](docs/UserStatus.md)
 - [UserSubscription](docs/UserSubscription.md)
 - [UserSubscriptionEligible](docs/UserSubscriptionEligible.md)
 - [Verify2FAEmailCodeResult](docs/Verify2FAEmailCodeResult.md)
 - [Verify2FAResult](docs/Verify2FAResult.md)
 - [VerifyAuthTokenResult](docs/VerifyAuthTokenResult.md)
 - [World](docs/World.md)
 - [WorldMetadata](docs/WorldMetadata.md)
 - [WorldPublishStatus](docs/WorldPublishStatus.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### authCookie

- **Type**: API key
- **API key parameter name**: auth
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: authCookie and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		vrchatapi.ContextAPIKeys,
		map[string]vrchatapi.APIKey{
			"authCookie": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### authHeader

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), vrchatapi.ContextBasicAuth, vrchatapi.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

### twoFactorAuthCookie

- **Type**: API key
- **API key parameter name**: twoFactorAuth
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: twoFactorAuthCookie and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		vrchatapi.ContextAPIKeys,
		map[string]vrchatapi.APIKey{
			"twoFactorAuthCookie": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

vrchatapi.lpv0t@aries.fyi

