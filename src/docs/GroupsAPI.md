# \GroupsAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupGalleryImage**](GroupsAPI.md#AddGroupGalleryImage) | **Post** /groups/{groupId}/galleries/{groupGalleryId}/images | Add Group Gallery Image
[**AddGroupMemberRole**](GroupsAPI.md#AddGroupMemberRole) | **Put** /groups/{groupId}/members/{userId}/roles/{groupRoleId} | Add Role to GroupMember
[**AddGroupPost**](GroupsAPI.md#AddGroupPost) | **Post** /groups/{groupId}/posts | Create a post in a Group
[**BanGroupMember**](GroupsAPI.md#BanGroupMember) | **Post** /groups/{groupId}/bans | Ban Group Member
[**BlockGroup**](GroupsAPI.md#BlockGroup) | **Post** /groups/{groupId}/block | Block Group
[**CancelGroupRequest**](GroupsAPI.md#CancelGroupRequest) | **Delete** /groups/{groupId}/requests | Cancel Group Join Request
[**CancelGroupTransfer**](GroupsAPI.md#CancelGroupTransfer) | **Delete** /groups/{groupId}/transfer | Cancel Group Transfer
[**CreateGroup**](GroupsAPI.md#CreateGroup) | **Post** /groups | Create Group
[**CreateGroupAnnouncement**](GroupsAPI.md#CreateGroupAnnouncement) | **Post** /groups/{groupId}/announcement | Create Group Announcement
[**CreateGroupGallery**](GroupsAPI.md#CreateGroupGallery) | **Post** /groups/{groupId}/galleries | Create Group Gallery
[**CreateGroupInvite**](GroupsAPI.md#CreateGroupInvite) | **Post** /groups/{groupId}/invites | Invite User to Group
[**CreateGroupRole**](GroupsAPI.md#CreateGroupRole) | **Post** /groups/{groupId}/roles | Create GroupRole
[**DeclineGroupInvite**](GroupsAPI.md#DeclineGroupInvite) | **Put** /groups/{groupId}/invites | Decline Invite from Group
[**DeleteGroup**](GroupsAPI.md#DeleteGroup) | **Delete** /groups/{groupId} | Delete Group
[**DeleteGroupAnnouncement**](GroupsAPI.md#DeleteGroupAnnouncement) | **Delete** /groups/{groupId}/announcement | Delete Group Announcement
[**DeleteGroupGallery**](GroupsAPI.md#DeleteGroupGallery) | **Delete** /groups/{groupId}/galleries/{groupGalleryId} | Delete Group Gallery
[**DeleteGroupGalleryImage**](GroupsAPI.md#DeleteGroupGalleryImage) | **Delete** /groups/{groupId}/galleries/{groupGalleryId}/images/{groupGalleryImageId} | Delete Group Gallery Image
[**DeleteGroupInvite**](GroupsAPI.md#DeleteGroupInvite) | **Delete** /groups/{groupId}/invites/{userId} | Delete User Invite
[**DeleteGroupPost**](GroupsAPI.md#DeleteGroupPost) | **Delete** /groups/{groupId}/posts/{notificationId} | Delete a Group post
[**DeleteGroupRole**](GroupsAPI.md#DeleteGroupRole) | **Delete** /groups/{groupId}/roles/{groupRoleId} | Delete Group Role
[**GetGroup**](GroupsAPI.md#GetGroup) | **Get** /groups/{groupId} | Get Group by ID
[**GetGroupAnnouncements**](GroupsAPI.md#GetGroupAnnouncements) | **Get** /groups/{groupId}/announcement | Get Group Announcement
[**GetGroupAuditLogEntryTypes**](GroupsAPI.md#GetGroupAuditLogEntryTypes) | **Get** /groups/{groupId}/auditLogTypes | Get Group Audit Log Entry Types
[**GetGroupAuditLogs**](GroupsAPI.md#GetGroupAuditLogs) | **Get** /groups/{groupId}/auditLogs | Get Group Audit Logs
[**GetGroupBans**](GroupsAPI.md#GetGroupBans) | **Get** /groups/{groupId}/bans | Get Group Bans
[**GetGroupGalleryImages**](GroupsAPI.md#GetGroupGalleryImages) | **Get** /groups/{groupId}/galleries/{groupGalleryId} | Get Group Gallery Images
[**GetGroupInstances**](GroupsAPI.md#GetGroupInstances) | **Get** /groups/{groupId}/instances | Get Group Instances
[**GetGroupInvites**](GroupsAPI.md#GetGroupInvites) | **Get** /groups/{groupId}/invites | Get Group Invites Sent
[**GetGroupMember**](GroupsAPI.md#GetGroupMember) | **Get** /groups/{groupId}/members/{userId} | Get Group Member
[**GetGroupMembers**](GroupsAPI.md#GetGroupMembers) | **Get** /groups/{groupId}/members | List Group Members
[**GetGroupPermissions**](GroupsAPI.md#GetGroupPermissions) | **Get** /groups/{groupId}/permissions | List Group Permissions
[**GetGroupPosts**](GroupsAPI.md#GetGroupPosts) | **Get** /groups/{groupId}/posts | Get posts from a Group
[**GetGroupRequests**](GroupsAPI.md#GetGroupRequests) | **Get** /groups/{groupId}/requests | Get Group Join Requests
[**GetGroupRoleTemplates**](GroupsAPI.md#GetGroupRoleTemplates) | **Get** /groups/roleTemplates | Get Group Role Templates
[**GetGroupRoles**](GroupsAPI.md#GetGroupRoles) | **Get** /groups/{groupId}/roles | Get Group Roles
[**GetGroupTransferability**](GroupsAPI.md#GetGroupTransferability) | **Get** /groups/{groupId}/transfer | Get Group Transferability
[**InitiateOrAcceptGroupTransfer**](GroupsAPI.md#InitiateOrAcceptGroupTransfer) | **Post** /groups/{groupId}/transfer | Initiate or Accept Group Transfer
[**JoinGroup**](GroupsAPI.md#JoinGroup) | **Post** /groups/{groupId}/join | Join Group
[**KickGroupMember**](GroupsAPI.md#KickGroupMember) | **Delete** /groups/{groupId}/members/{userId} | Kick Group Member
[**LeaveGroup**](GroupsAPI.md#LeaveGroup) | **Post** /groups/{groupId}/leave | Leave Group
[**RemoveGroupMemberRole**](GroupsAPI.md#RemoveGroupMemberRole) | **Delete** /groups/{groupId}/members/{userId}/roles/{groupRoleId} | Remove Role from GroupMember
[**RespondGroupJoinRequest**](GroupsAPI.md#RespondGroupJoinRequest) | **Put** /groups/{groupId}/requests/{userId} | Respond Group Join request
[**SearchGroupMembers**](GroupsAPI.md#SearchGroupMembers) | **Get** /groups/{groupId}/members/search | Search Group Members
[**SearchGroups**](GroupsAPI.md#SearchGroups) | **Get** /groups | Search Group
[**UnbanGroupMember**](GroupsAPI.md#UnbanGroupMember) | **Delete** /groups/{groupId}/bans/{userId} | Unban Group Member
[**UpdateGroup**](GroupsAPI.md#UpdateGroup) | **Put** /groups/{groupId} | Update Group
[**UpdateGroupGallery**](GroupsAPI.md#UpdateGroupGallery) | **Put** /groups/{groupId}/galleries/{groupGalleryId} | Update Group Gallery
[**UpdateGroupMember**](GroupsAPI.md#UpdateGroupMember) | **Put** /groups/{groupId}/members/{userId} | Update Group Member
[**UpdateGroupPost**](GroupsAPI.md#UpdateGroupPost) | **Put** /groups/{groupId}/posts/{notificationId} | Edits a Group post
[**UpdateGroupRepresentation**](GroupsAPI.md#UpdateGroupRepresentation) | **Put** /groups/{groupId}/representation | Update Group Representation
[**UpdateGroupRole**](GroupsAPI.md#UpdateGroupRole) | **Put** /groups/{groupId}/roles/{groupRoleId} | Update Group Role



## AddGroupGalleryImage

> GroupGalleryImage AddGroupGalleryImage(ctx, groupId, groupGalleryId).AddGroupGalleryImageRequest(addGroupGalleryImageRequest).Execute()

Add Group Gallery Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	groupGalleryId := "ggal_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery ID.
	addGroupGalleryImageRequest := *openapiclient.NewAddGroupGalleryImageRequest("file_ce35d830-e20a-4df0-a6d4-5aaef4508044") // AddGroupGalleryImageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AddGroupGalleryImage(context.Background(), groupId, groupGalleryId).AddGroupGalleryImageRequest(addGroupGalleryImageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddGroupGalleryImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupGalleryImage`: GroupGalleryImage
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AddGroupGalleryImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupGalleryId** | **string** | Must be a valid group gallery ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupGalleryImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addGroupGalleryImageRequest** | [**AddGroupGalleryImageRequest**](AddGroupGalleryImageRequest.md) |  | 

### Return type

[**GroupGalleryImage**](GroupGalleryImage.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGroupMemberRole

> []string AddGroupMemberRole(ctx, groupId, userId, groupRoleId).Execute()

Add Role to GroupMember



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	userId := "userId_example" // string | Must be a valid user ID.
	groupRoleId := "grol_00000000-0000-0000-0000-000000000000" // string | Must be a valid group role ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AddGroupMemberRole(context.Background(), groupId, userId, groupRoleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddGroupMemberRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupMemberRole`: []string
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AddGroupMemberRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 
**groupRoleId** | **string** | Must be a valid group role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**[]string**

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGroupPost

> GroupPost AddGroupPost(ctx, groupId).CreateGroupPostRequest(createGroupPostRequest).Execute()

Create a post in a Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	createGroupPostRequest := *openapiclient.NewCreateGroupPostRequest(false, "Come join us for the event!", "Event is starting soon!", openapiclient.GroupPostVisibility("group")) // CreateGroupPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AddGroupPost(context.Background(), groupId).CreateGroupPostRequest(createGroupPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddGroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroupPost`: GroupPost
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AddGroupPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupPostRequest** | [**CreateGroupPostRequest**](CreateGroupPostRequest.md) |  | 

### Return type

[**GroupPost**](GroupPost.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BanGroupMember

> GroupMember BanGroupMember(ctx, groupId).BanGroupMemberRequest(banGroupMemberRequest).Execute()

Ban Group Member



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	banGroupMemberRequest := *openapiclient.NewBanGroupMemberRequest("usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469") // BanGroupMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.BanGroupMember(context.Background(), groupId).BanGroupMemberRequest(banGroupMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.BanGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BanGroupMember`: GroupMember
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.BanGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBanGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **banGroupMemberRequest** | [**BanGroupMemberRequest**](BanGroupMemberRequest.md) |  | 

### Return type

[**GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockGroup

> Success BlockGroup(ctx, groupId).Execute()

Block Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.BlockGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.BlockGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockGroup`: Success
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.BlockGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlockGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelGroupRequest

> CancelGroupRequest(ctx, groupId).Execute()

Cancel Group Join Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.CancelGroupRequest(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CancelGroupRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelGroupRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelGroupTransfer

> Success CancelGroupTransfer(ctx, groupId).Execute()

Cancel Group Transfer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CancelGroupTransfer(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CancelGroupTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelGroupTransfer`: Success
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CancelGroupTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelGroupTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> Group CreateGroup(ctx).CreateGroupRequest(createGroupRequest).Execute()

Create Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	createGroupRequest := *openapiclient.NewCreateGroupRequest("Name_example", openapiclient.GroupRoleTemplate("default"), "ShortCode_example") // CreateGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateGroup(context.Background()).CreateGroupRequest(createGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupRequest** | [**CreateGroupRequest**](CreateGroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupAnnouncement

> GroupAnnouncement CreateGroupAnnouncement(ctx, groupId).CreateGroupAnnouncementRequest(createGroupAnnouncementRequest).Execute()

Create Group Announcement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	createGroupAnnouncementRequest := *openapiclient.NewCreateGroupAnnouncementRequest("Event is starting soon!") // CreateGroupAnnouncementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateGroupAnnouncement(context.Background(), groupId).CreateGroupAnnouncementRequest(createGroupAnnouncementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroupAnnouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupAnnouncement`: GroupAnnouncement
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroupAnnouncement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupAnnouncementRequest** | [**CreateGroupAnnouncementRequest**](CreateGroupAnnouncementRequest.md) |  | 

### Return type

[**GroupAnnouncement**](GroupAnnouncement.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupGallery

> GroupGallery CreateGroupGallery(ctx, groupId).CreateGroupGalleryRequest(createGroupGalleryRequest).Execute()

Create Group Gallery



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	createGroupGalleryRequest := *openapiclient.NewCreateGroupGalleryRequest("Example Gallery") // CreateGroupGalleryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateGroupGallery(context.Background(), groupId).CreateGroupGalleryRequest(createGroupGalleryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroupGallery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupGallery`: GroupGallery
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroupGallery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupGalleryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupGalleryRequest** | [**CreateGroupGalleryRequest**](CreateGroupGalleryRequest.md) |  | 

### Return type

[**GroupGallery**](GroupGallery.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupInvite

> CreateGroupInvite(ctx, groupId).CreateGroupInviteRequest(createGroupInviteRequest).Execute()

Invite User to Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	createGroupInviteRequest := *openapiclient.NewCreateGroupInviteRequest("usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469") // CreateGroupInviteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.CreateGroupInvite(context.Background(), groupId).CreateGroupInviteRequest(createGroupInviteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroupInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupInviteRequest** | [**CreateGroupInviteRequest**](CreateGroupInviteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupRole

> GroupRole CreateGroupRole(ctx, groupId).CreateGroupRoleRequest(createGroupRoleRequest).Execute()

Create GroupRole



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	createGroupRoleRequest := *openapiclient.NewCreateGroupRoleRequest() // CreateGroupRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateGroupRole(context.Background(), groupId).CreateGroupRoleRequest(createGroupRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroupRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupRole`: GroupRole
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroupRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGroupRoleRequest** | [**CreateGroupRoleRequest**](CreateGroupRoleRequest.md) |  | 

### Return type

[**GroupRole**](GroupRole.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclineGroupInvite

> Success DeclineGroupInvite(ctx, groupId).DeclineGroupInviteRequest(declineGroupInviteRequest).Execute()

Decline Invite from Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	declineGroupInviteRequest := *openapiclient.NewDeclineGroupInviteRequest() // DeclineGroupInviteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeclineGroupInvite(context.Background(), groupId).DeclineGroupInviteRequest(declineGroupInviteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeclineGroupInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeclineGroupInvite`: Success
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeclineGroupInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineGroupInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **declineGroupInviteRequest** | [**DeclineGroupInviteRequest**](DeclineGroupInviteRequest.md) |  | 

### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> Success DeleteGroup(ctx, groupId).HardDelete(hardDelete).Execute()

Delete Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	hardDelete := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeleteGroup(context.Background(), groupId).HardDelete(hardDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroup`: Success
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hardDelete** | **bool** |  | 

### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupAnnouncement

> Success DeleteGroupAnnouncement(ctx, groupId).Execute()

Delete Group Announcement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeleteGroupAnnouncement(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupAnnouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupAnnouncement`: Success
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroupAnnouncement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupGallery

> Success DeleteGroupGallery(ctx, groupId, groupGalleryId).Execute()

Delete Group Gallery



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	groupGalleryId := "ggal_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeleteGroupGallery(context.Background(), groupId, groupGalleryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupGallery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupGallery`: Success
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroupGallery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupGalleryId** | **string** | Must be a valid group gallery ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupGalleryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupGalleryImage

> Success DeleteGroupGalleryImage(ctx, groupId, groupGalleryId, groupGalleryImageId).Execute()

Delete Group Gallery Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	groupGalleryId := "ggal_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery ID.
	groupGalleryImageId := "ggim_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery image ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeleteGroupGalleryImage(context.Background(), groupId, groupGalleryId, groupGalleryImageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupGalleryImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupGalleryImage`: Success
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroupGalleryImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupGalleryId** | **string** | Must be a valid group gallery ID. | 
**groupGalleryImageId** | **string** | Must be a valid group gallery image ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupGalleryImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupInvite

> DeleteGroupInvite(ctx, groupId, userId).Execute()

Delete User Invite



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	userId := "userId_example" // string | Must be a valid user ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.DeleteGroupInvite(context.Background(), groupId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupPost

> Success DeleteGroupPost(ctx, groupId, notificationId).Execute()

Delete a Group post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	notificationId := "notificationId_example" // string | Must be a valid notification ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeleteGroupPost(context.Background(), groupId, notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupPost`: Success
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroupPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupRole

> []GroupRole DeleteGroupRole(ctx, groupId, groupRoleId).Execute()

Delete Group Role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	groupRoleId := "grol_00000000-0000-0000-0000-000000000000" // string | Must be a valid group role ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeleteGroupRole(context.Background(), groupId, groupRoleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupRole`: []GroupRole
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroupRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupRoleId** | **string** | Must be a valid group role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GroupRole**](GroupRole.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroup

> Group GetGroup(ctx, groupId).IncludeRoles(includeRoles).Execute()

Get Group by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	includeRoles := true // bool | Include roles for the Group object. Defaults to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroup(context.Background(), groupId).IncludeRoles(includeRoles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeRoles** | **bool** | Include roles for the Group object. Defaults to false. | 

### Return type

[**Group**](Group.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupAnnouncements

> GroupAnnouncement GetGroupAnnouncements(ctx, groupId).Execute()

Get Group Announcement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupAnnouncements(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupAnnouncements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupAnnouncements`: GroupAnnouncement
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupAnnouncements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupAnnouncementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupAnnouncement**](GroupAnnouncement.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupAuditLogEntryTypes

> []string GetGroupAuditLogEntryTypes(ctx, groupId).Execute()

Get Group Audit Log Entry Types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupAuditLogEntryTypes(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupAuditLogEntryTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupAuditLogEntryTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupAuditLogEntryTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupAuditLogEntryTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupAuditLogs

> PaginatedGroupAuditLogEntryList GetGroupAuditLogs(ctx, groupId).N(n).Offset(offset).StartDate(startDate).EndDate(endDate).ActorIds(actorIds).EventTypes(eventTypes).TargetIds(targetIds).Execute()

Get Group Audit Logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	startDate := time.Now() // time.Time | The start date of the search range. (optional)
	endDate := time.Now() // time.Time | The end date of the search range. (optional)
	actorIds := "usr_00000000-0000-0000-0000-000000000000,usr_11111111-1111-1111-1111-111111111111" // string | The comma-separated actor ids to search for. (optional)
	eventTypes := "group.member.remove,group.instance.kick" // string | The comma-separated event types to search for. (optional)
	targetIds := "usr_00000000-0000-0000-0000-000000000000,usr_11111111-1111-1111-1111-111111111111" // string | The comma-separated target ids to search for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupAuditLogs(context.Background(), groupId).N(n).Offset(offset).StartDate(startDate).EndDate(endDate).ActorIds(actorIds).EventTypes(eventTypes).TargetIds(targetIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupAuditLogs`: PaginatedGroupAuditLogEntryList
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupAuditLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **startDate** | **time.Time** | The start date of the search range. | 
 **endDate** | **time.Time** | The end date of the search range. | 
 **actorIds** | **string** | The comma-separated actor ids to search for. | 
 **eventTypes** | **string** | The comma-separated event types to search for. | 
 **targetIds** | **string** | The comma-separated target ids to search for. | 

### Return type

[**PaginatedGroupAuditLogEntryList**](PaginatedGroupAuditLogEntryList.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupBans

> []GroupMember GetGroupBans(ctx, groupId).N(n).Offset(offset).Execute()

Get Group Bans



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupBans(context.Background(), groupId).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupBans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupBans`: []GroupMember
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupBans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupBansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupGalleryImages

> []GroupGalleryImage GetGroupGalleryImages(ctx, groupId, groupGalleryId).N(n).Offset(offset).Approved(approved).Execute()

Get Group Gallery Images



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	groupGalleryId := "ggal_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery ID.
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	approved := true // bool | If specified, only returns images that have been approved or not approved. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupGalleryImages(context.Background(), groupId, groupGalleryId).N(n).Offset(offset).Approved(approved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupGalleryImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupGalleryImages`: []GroupGalleryImage
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupGalleryImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupGalleryId** | **string** | Must be a valid group gallery ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupGalleryImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **approved** | **bool** | If specified, only returns images that have been approved or not approved. | 

### Return type

[**[]GroupGalleryImage**](GroupGalleryImage.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupInstances

> []GroupInstance GetGroupInstances(ctx, groupId).Execute()

Get Group Instances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupInstances(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupInstances`: []GroupInstance
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GroupInstance**](GroupInstance.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupInvites

> []GroupMember GetGroupInvites(ctx, groupId).N(n).Offset(offset).Execute()

Get Group Invites Sent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupInvites(context.Background(), groupId).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupInvites`: []GroupMember
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMember

> GroupLimitedMember GetGroupMember(ctx, groupId, userId).Execute()

Get Group Member



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	userId := "userId_example" // string | Must be a valid user ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupMember(context.Background(), groupId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupMember`: GroupLimitedMember
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupLimitedMember**](GroupLimitedMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMembers

> []GroupMember GetGroupMembers(ctx, groupId).N(n).Offset(offset).Sort(sort).RoleId(roleId).Execute()

List Group Members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	sort := openapiclient.GroupSearchSort("joinedAt:asc") // GroupSearchSort | The sort order of Group Member results (optional)
	roleId := "roleId_example" // string | Only returns members with a specific groupRoleId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupMembers(context.Background(), groupId).N(n).Offset(offset).Sort(sort).RoleId(roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupMembers`: []GroupMember
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **sort** | [**GroupSearchSort**](GroupSearchSort.md) | The sort order of Group Member results | 
 **roleId** | **string** | Only returns members with a specific groupRoleId | 

### Return type

[**[]GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupPermissions

> []GroupPermission GetGroupPermissions(ctx, groupId).Execute()

List Group Permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupPermissions(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupPermissions`: []GroupPermission
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GroupPermission**](GroupPermission.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupPosts

> GetGroupPosts200Response GetGroupPosts(ctx, groupId).N(n).Offset(offset).PublicOnly(publicOnly).Execute()

Get posts from a Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	publicOnly := true // bool | See public posts only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupPosts(context.Background(), groupId).N(n).Offset(offset).PublicOnly(publicOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupPosts`: GetGroupPosts200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupPosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **publicOnly** | **bool** | See public posts only. | 

### Return type

[**GetGroupPosts200Response**](GetGroupPosts200Response.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRequests

> []GroupMember GetGroupRequests(ctx, groupId).N(n).Offset(offset).Blocked(blocked).Execute()

Get Group Join Requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	blocked := true // bool | See blocked join requests (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupRequests(context.Background(), groupId).N(n).Offset(offset).Blocked(blocked).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRequests`: []GroupMember
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **blocked** | **bool** | See blocked join requests | 

### Return type

[**[]GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRoleTemplates

> map[string]GroupRoleTemplateValues GetGroupRoleTemplates(ctx).Execute()

Get Group Role Templates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupRoleTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupRoleTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRoleTemplates`: map[string]GroupRoleTemplateValues
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupRoleTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRoleTemplatesRequest struct via the builder pattern


### Return type

[**map[string]GroupRoleTemplateValues**](GroupRoleTemplateValues.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupRoles

> []GroupRole GetGroupRoles(ctx, groupId).Execute()

Get Group Roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupRoles(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRoles`: []GroupRole
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GroupRole**](GroupRole.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupTransferability

> GroupTransferable GetGroupTransferability(ctx, groupId).TransferTargetId(transferTargetId).Execute()

Get Group Transferability



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	transferTargetId := "transferTargetId_example" // string | The UserID of the prospective transferee. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupTransferability(context.Background(), groupId).TransferTargetId(transferTargetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupTransferability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupTransferability`: GroupTransferable
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupTransferability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupTransferabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferTargetId** | **string** | The UserID of the prospective transferee. | 

### Return type

[**GroupTransferable**](GroupTransferable.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateOrAcceptGroupTransfer

> Success InitiateOrAcceptGroupTransfer(ctx, groupId).TransferGroupRequest(transferGroupRequest).Execute()

Initiate or Accept Group Transfer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	transferGroupRequest := *openapiclient.NewTransferGroupRequest() // TransferGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.InitiateOrAcceptGroupTransfer(context.Background(), groupId).TransferGroupRequest(transferGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.InitiateOrAcceptGroupTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateOrAcceptGroupTransfer`: Success
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.InitiateOrAcceptGroupTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateOrAcceptGroupTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferGroupRequest** | [**TransferGroupRequest**](TransferGroupRequest.md) |  | 

### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinGroup

> GroupMember JoinGroup(ctx, groupId).ConfirmOverrideBlock(confirmOverrideBlock).JoinGroupRequest(joinGroupRequest).Execute()

Join Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	confirmOverrideBlock := true // bool | Manually override the failure that would occur if the user has blocked the group. (optional)
	joinGroupRequest := *openapiclient.NewJoinGroupRequest() // JoinGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.JoinGroup(context.Background(), groupId).ConfirmOverrideBlock(confirmOverrideBlock).JoinGroupRequest(joinGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.JoinGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JoinGroup`: GroupMember
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.JoinGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmOverrideBlock** | **bool** | Manually override the failure that would occur if the user has blocked the group. | 
 **joinGroupRequest** | [**JoinGroupRequest**](JoinGroupRequest.md) |  | 

### Return type

[**GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KickGroupMember

> Success KickGroupMember(ctx, groupId, userId).Execute()

Kick Group Member



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	userId := "userId_example" // string | Must be a valid user ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.KickGroupMember(context.Background(), groupId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.KickGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KickGroupMember`: Success
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.KickGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKickGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaveGroup

> LeaveGroup(ctx, groupId).Execute()

Leave Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.LeaveGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.LeaveGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeaveGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGroupMemberRole

> []string RemoveGroupMemberRole(ctx, groupId, userId, groupRoleId).Execute()

Remove Role from GroupMember



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	userId := "userId_example" // string | Must be a valid user ID.
	groupRoleId := "grol_00000000-0000-0000-0000-000000000000" // string | Must be a valid group role ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.RemoveGroupMemberRole(context.Background(), groupId, userId, groupRoleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RemoveGroupMemberRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveGroupMemberRole`: []string
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.RemoveGroupMemberRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 
**groupRoleId** | **string** | Must be a valid group role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**[]string**

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RespondGroupJoinRequest

> RespondGroupJoinRequest(ctx, groupId, userId).RespondGroupJoinRequest(respondGroupJoinRequest).Execute()

Respond Group Join request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	userId := "userId_example" // string | Must be a valid user ID.
	respondGroupJoinRequest := *openapiclient.NewRespondGroupJoinRequest(openapiclient.GroupJoinRequestAction("accept")) // RespondGroupJoinRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.RespondGroupJoinRequest(context.Background(), groupId, userId).RespondGroupJoinRequest(respondGroupJoinRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RespondGroupJoinRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRespondGroupJoinRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **respondGroupJoinRequest** | [**RespondGroupJoinRequest**](RespondGroupJoinRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGroupMembers

> SearchGroupMembers200Response SearchGroupMembers(ctx, groupId).Query(query).N(n).Offset(offset).Execute()

Search Group Members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	query := "query_example" // string | Filter for member displayName.
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.SearchGroupMembers(context.Background(), groupId).Query(query).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.SearchGroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGroupMembers`: SearchGroupMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.SearchGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | Filter for member displayName. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**SearchGroupMembers200Response**](SearchGroupMembers200Response.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGroups

> []LimitedGroup SearchGroups(ctx).Query(query).Offset(offset).N(n).Execute()

Search Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	query := "query_example" // string | Query to search for, can be either Group Name or Group shortCode (optional)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.SearchGroups(context.Background()).Query(query).Offset(offset).N(n).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.SearchGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGroups`: []LimitedGroup
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.SearchGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Query to search for, can be either Group Name or Group shortCode | 
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **n** | **int32** | The number of objects to return. | [default to 60]

### Return type

[**[]LimitedGroup**](LimitedGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbanGroupMember

> GroupMember UnbanGroupMember(ctx, groupId, userId).Execute()

Unban Group Member



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	userId := "userId_example" // string | Must be a valid user ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UnbanGroupMember(context.Background(), groupId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UnbanGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbanGroupMember`: GroupMember
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UnbanGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbanGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupMember**](GroupMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> Group UpdateGroup(ctx, groupId).UpdateGroupRequest(updateGroupRequest).Execute()

Update Group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	updateGroupRequest := *openapiclient.NewUpdateGroupRequest() // UpdateGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroup(context.Background(), groupId).UpdateGroupRequest(updateGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGroupRequest** | [**UpdateGroupRequest**](UpdateGroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupGallery

> GroupGallery UpdateGroupGallery(ctx, groupId, groupGalleryId).UpdateGroupGalleryRequest(updateGroupGalleryRequest).Execute()

Update Group Gallery



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	groupGalleryId := "ggal_00000000-0000-0000-0000-000000000000" // string | Must be a valid group gallery ID.
	updateGroupGalleryRequest := *openapiclient.NewUpdateGroupGalleryRequest() // UpdateGroupGalleryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroupGallery(context.Background(), groupId, groupGalleryId).UpdateGroupGalleryRequest(updateGroupGalleryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupGallery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupGallery`: GroupGallery
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroupGallery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupGalleryId** | **string** | Must be a valid group gallery ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupGalleryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGroupGalleryRequest** | [**UpdateGroupGalleryRequest**](UpdateGroupGalleryRequest.md) |  | 

### Return type

[**GroupGallery**](GroupGallery.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupMember

> GroupLimitedMember UpdateGroupMember(ctx, groupId, userId).UpdateGroupMemberRequest(updateGroupMemberRequest).Execute()

Update Group Member



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	userId := "userId_example" // string | Must be a valid user ID.
	updateGroupMemberRequest := *openapiclient.NewUpdateGroupMemberRequest() // UpdateGroupMemberRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroupMember(context.Background(), groupId, userId).UpdateGroupMemberRequest(updateGroupMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupMember`: GroupLimitedMember
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGroupMemberRequest** | [**UpdateGroupMemberRequest**](UpdateGroupMemberRequest.md) |  | 

### Return type

[**GroupLimitedMember**](GroupLimitedMember.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupPost

> GroupPost UpdateGroupPost(ctx, groupId, notificationId).CreateGroupPostRequest(createGroupPostRequest).Execute()

Edits a Group post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	notificationId := "notificationId_example" // string | Must be a valid notification ID.
	createGroupPostRequest := *openapiclient.NewCreateGroupPostRequest(false, "Come join us for the event!", "Event is starting soon!", openapiclient.GroupPostVisibility("group")) // CreateGroupPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroupPost(context.Background(), groupId, notificationId).CreateGroupPostRequest(createGroupPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupPost`: GroupPost
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroupPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createGroupPostRequest** | [**CreateGroupPostRequest**](CreateGroupPostRequest.md) |  | 

### Return type

[**GroupPost**](GroupPost.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupRepresentation

> Success UpdateGroupRepresentation(ctx, groupId).UpdateGroupRepresentationRequest(updateGroupRepresentationRequest).Execute()

Update Group Representation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	updateGroupRepresentationRequest := *openapiclient.NewUpdateGroupRepresentationRequest(false) // UpdateGroupRepresentationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroupRepresentation(context.Background(), groupId).UpdateGroupRepresentationRequest(updateGroupRepresentationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupRepresentation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupRepresentation`: Success
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroupRepresentation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRepresentationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGroupRepresentationRequest** | [**UpdateGroupRepresentationRequest**](UpdateGroupRepresentationRequest.md) |  | 

### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupRole

> []GroupRole UpdateGroupRole(ctx, groupId, groupRoleId).UpdateGroupRoleRequest(updateGroupRoleRequest).Execute()

Update Group Role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	groupRoleId := "grol_00000000-0000-0000-0000-000000000000" // string | Must be a valid group role ID.
	updateGroupRoleRequest := *openapiclient.NewUpdateGroupRoleRequest() // UpdateGroupRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroupRole(context.Background(), groupId, groupRoleId).UpdateGroupRoleRequest(updateGroupRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupRole`: []GroupRole
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroupRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**groupRoleId** | **string** | Must be a valid group role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateGroupRoleRequest** | [**UpdateGroupRoleRequest**](UpdateGroupRoleRequest.md) |  | 

### Return type

[**[]GroupRole**](GroupRole.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

