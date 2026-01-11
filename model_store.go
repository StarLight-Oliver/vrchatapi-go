/*
VRChat API Documentation

![VRChat API Banner](https://vrchatapi.github.io/assets/img/api_banner_1500x400.png)  # Welcome to the VRChat API  Before we begin, we would like to state this is a **COMMUNITY DRIVEN PROJECT**. This means that everything you read on here was written by the community itself and is **not** officially supported by VRChat. The documentation is provided \"AS IS\", and any action you take towards VRChat is completely your own responsibility.  The documentation and additional libraries SHALL ONLY be used for applications interacting with VRChat's API in accordance with their [Terms of Service](https://hello.vrchat.com/legal) and [Community Guidelines](https://hello.vrchat.com/community-guidelines), and MUST NOT be used for modifying the client, \"avatar ripping\", or other illegal activities. Malicious usage or spamming the API may result in account termination. Certain parts of the API are also more sensitive than others, for example moderation, so please tread extra carefully and read the warnings when present.  ![Tupper Policy on API](https://i.imgur.com/yLlW7Ok.png)  Finally, use of the API using applications other than the approved methods (website, VRChat application, Unity SDK) is not officially supported. VRChat provides no guarantee or support for external applications using the API. Access to API endpoints may break **at any time, without notice**. Therefore, please **do not ping** VRChat Staff in the VRChat Discord if you are having API problems, as they do not provide API support. We will make a best effort in keeping this documentation and associated language libraries up to date, but things might be outdated or missing. If you find that something is no longer valid, please contact us on Discord or [create an issue](https://github.com/vrchatapi/specification/issues) and tell us so we can fix it.  # Getting Started  The VRChat API can be used to programmatically retrieve or update information regarding your profile, friends, avatars, worlds and more. The API consists of two parts, \"Photon\" which is only used in-game, and the \"Web API\" which is used by both the game and the website. This documentation focuses only on the Web API.  The API is designed around the REST ideology, providing semi-simple and usually predictable URIs to access and modify objects. Requests support standard HTTP methods like GET, PUT, POST, and DELETE and standard status codes. Response bodies are always UTF-8 encoded JSON objects, unless explicitly documented otherwise.  <div class=\"callout callout-error\">   <strong>🛑 Warning! Do not touch Photon!</strong><br>   Photon is only used by the in-game client and should <b>not</b> be touched. Doing so may result in permanent account termination. </div>  <div class=\"callout callout-info\">   <strong>ℹ️ Authentication</strong><br>   Read <a href=\"#tag--authentication\">Authentication</a> for how to log in. </div>  # Using the API  For simply exploring what the API can do it is strongly recommended to download [Insomnia](https://insomnia.rest/download), a free and open-source API client that's great for sending requests to the API in an orderly fashion. Insomnia allows you to send data in the format that's required for VRChat's API. It is also possible to try out the API in your browser, by first logging in at [vrchat.com/home](https://vrchat.com/home/) and then going to [vrchat.com/api/1/auth/user](https://vrchat.com/api/1/auth/user), but the information will be much harder to work with.  For more permanent operation such as software development it is instead recommended to use one of the existing language SDKs. This community project maintains API libraries in several languages, which allows you to interact with the API with simple function calls rather than having to implement the HTTP protocol yourself. Most of these libraries are automatically generated from the API specification, sometimes with additional helpful wrapper code to make usage easier. This allows them to be almost automatically updated and expanded upon as soon as a new feature is introduced in the specification itself. The libraries can be found on [GitHub](https://github.com/vrchatapi) or following:  * [NodeJS (JavaScript)](https://www.npmjs.com/package/vrchat) * [Dart](https://pub.dev/packages/vrchat_dart) * [Rust](https://crates.io/crates/vrchatapi) * [C#](https://github.com/vrchatapi/vrchatapi-csharp) * [Python](https://github.com/vrchatapi/vrchatapi-python)  # Pagination  Most endpoints enforce pagination, meaning they will only return 10 entries by default, and never more than 100.<br> Using both the limit and offset parameters allows you to easily paginate through a large number of objects.  | Query Parameter | Type | Description | | ----------|--|------- | | `n` | integer  | The number of objects to return. This value often defaults to 10. Highest limit is always 100.| | `offset` | integer  | A zero-based offset from the default object sorting.|  If a request returns fewer objects than the `limit` parameter, there are no more items available to return.  # Contribution  Do you want to get involved in the documentation effort? Do you want to help improve one of the language API libraries? This project is an [OPEN Open Source Project](https://openopensource.org)! This means that individuals making significant and valuable contributions are given commit-access to the project. It also means we are very open and welcoming of new people making contributions, unlike some more guarded open-source projects.  [![Discord](https://img.shields.io/static/v1?label=vrchatapi&message=discord&color=blueviolet&style=for-the-badge)](https://discord.gg/qjZE9C9fkB)

API version: 1.20.7
Contact: vrchatapi.lpv0t@aries.fyi
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Store type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Store{}

// Store struct for Store
type Store struct {
	Description string  `json:"description"`
	DisplayName string  `json:"displayName"`
	GroupId     *string `json:"groupId,omitempty"`
	Id          string  `json:"id"`
	// Only for store type world and group
	ListingIds []string `json:"listingIds,omitempty"`
	// Only for store type world and group
	Listings          []ProductListing `json:"listings,omitempty"`
	SellerDisplayName string           `json:"sellerDisplayName"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	SellerId string `json:"sellerId"`
	// Only for store type house
	ShelfIds []string `json:"shelfIds,omitempty"`
	// Only for store type house
	Shelves   []StoreShelf `json:"shelves,omitempty"`
	StoreId   string       `json:"storeId"`
	StoreType StoreType    `json:"storeType"`
	Tags      []string     `json:"tags"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	WorldId *string `json:"worldId,omitempty"`
}

type _Store Store

// NewStore instantiates a new Store object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStore(description string, displayName string, id string, sellerDisplayName string, sellerId string, storeId string, storeType StoreType, tags []string) *Store {
	this := Store{}
	this.Description = description
	this.DisplayName = displayName
	this.Id = id
	this.SellerDisplayName = sellerDisplayName
	this.SellerId = sellerId
	this.StoreId = storeId
	this.StoreType = storeType
	this.Tags = tags
	return &this
}

// NewStoreWithDefaults instantiates a new Store object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreWithDefaults() *Store {
	this := Store{}
	var storeType StoreType = StoreType_GROUP
	this.StoreType = storeType
	return &this
}

// GetDescription returns the Description field value
func (o *Store) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Store) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Store) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value
func (o *Store) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Store) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Store) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *Store) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *Store) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *Store) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value
func (o *Store) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Store) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Store) SetId(v string) {
	o.Id = v
}

// GetListingIds returns the ListingIds field value if set, zero value otherwise.
func (o *Store) GetListingIds() []string {
	if o == nil || IsNil(o.ListingIds) {
		var ret []string
		return ret
	}
	return o.ListingIds
}

// GetListingIdsOk returns a tuple with the ListingIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetListingIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ListingIds) {
		return nil, false
	}
	return o.ListingIds, true
}

// HasListingIds returns a boolean if a field has been set.
func (o *Store) HasListingIds() bool {
	if o != nil && !IsNil(o.ListingIds) {
		return true
	}

	return false
}

// SetListingIds gets a reference to the given []string and assigns it to the ListingIds field.
func (o *Store) SetListingIds(v []string) {
	o.ListingIds = v
}

// GetListings returns the Listings field value if set, zero value otherwise.
func (o *Store) GetListings() []ProductListing {
	if o == nil || IsNil(o.Listings) {
		var ret []ProductListing
		return ret
	}
	return o.Listings
}

// GetListingsOk returns a tuple with the Listings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetListingsOk() ([]ProductListing, bool) {
	if o == nil || IsNil(o.Listings) {
		return nil, false
	}
	return o.Listings, true
}

// HasListings returns a boolean if a field has been set.
func (o *Store) HasListings() bool {
	if o != nil && !IsNil(o.Listings) {
		return true
	}

	return false
}

// SetListings gets a reference to the given []ProductListing and assigns it to the Listings field.
func (o *Store) SetListings(v []ProductListing) {
	o.Listings = v
}

// GetSellerDisplayName returns the SellerDisplayName field value
func (o *Store) GetSellerDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerDisplayName
}

// GetSellerDisplayNameOk returns a tuple with the SellerDisplayName field value
// and a boolean to check if the value has been set.
func (o *Store) GetSellerDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerDisplayName, true
}

// SetSellerDisplayName sets field value
func (o *Store) SetSellerDisplayName(v string) {
	o.SellerDisplayName = v
}

// GetSellerId returns the SellerId field value
func (o *Store) GetSellerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerId
}

// GetSellerIdOk returns a tuple with the SellerId field value
// and a boolean to check if the value has been set.
func (o *Store) GetSellerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerId, true
}

// SetSellerId sets field value
func (o *Store) SetSellerId(v string) {
	o.SellerId = v
}

// GetShelfIds returns the ShelfIds field value if set, zero value otherwise.
func (o *Store) GetShelfIds() []string {
	if o == nil || IsNil(o.ShelfIds) {
		var ret []string
		return ret
	}
	return o.ShelfIds
}

// GetShelfIdsOk returns a tuple with the ShelfIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetShelfIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ShelfIds) {
		return nil, false
	}
	return o.ShelfIds, true
}

// HasShelfIds returns a boolean if a field has been set.
func (o *Store) HasShelfIds() bool {
	if o != nil && !IsNil(o.ShelfIds) {
		return true
	}

	return false
}

// SetShelfIds gets a reference to the given []string and assigns it to the ShelfIds field.
func (o *Store) SetShelfIds(v []string) {
	o.ShelfIds = v
}

// GetShelves returns the Shelves field value if set, zero value otherwise.
func (o *Store) GetShelves() []StoreShelf {
	if o == nil || IsNil(o.Shelves) {
		var ret []StoreShelf
		return ret
	}
	return o.Shelves
}

// GetShelvesOk returns a tuple with the Shelves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetShelvesOk() ([]StoreShelf, bool) {
	if o == nil || IsNil(o.Shelves) {
		return nil, false
	}
	return o.Shelves, true
}

// HasShelves returns a boolean if a field has been set.
func (o *Store) HasShelves() bool {
	if o != nil && !IsNil(o.Shelves) {
		return true
	}

	return false
}

// SetShelves gets a reference to the given []StoreShelf and assigns it to the Shelves field.
func (o *Store) SetShelves(v []StoreShelf) {
	o.Shelves = v
}

// GetStoreId returns the StoreId field value
func (o *Store) GetStoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value
// and a boolean to check if the value has been set.
func (o *Store) GetStoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoreId, true
}

// SetStoreId sets field value
func (o *Store) SetStoreId(v string) {
	o.StoreId = v
}

// GetStoreType returns the StoreType field value
func (o *Store) GetStoreType() StoreType {
	if o == nil {
		var ret StoreType
		return ret
	}

	return o.StoreType
}

// GetStoreTypeOk returns a tuple with the StoreType field value
// and a boolean to check if the value has been set.
func (o *Store) GetStoreTypeOk() (*StoreType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoreType, true
}

// SetStoreType sets field value
func (o *Store) SetStoreType(v StoreType) {
	o.StoreType = v
}

// GetTags returns the Tags field value
func (o *Store) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Store) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Store) SetTags(v []string) {
	o.Tags = v
}

// GetWorldId returns the WorldId field value if set, zero value otherwise.
func (o *Store) GetWorldId() string {
	if o == nil || IsNil(o.WorldId) {
		var ret string
		return ret
	}
	return *o.WorldId
}

// GetWorldIdOk returns a tuple with the WorldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetWorldIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorldId) {
		return nil, false
	}
	return o.WorldId, true
}

// HasWorldId returns a boolean if a field has been set.
func (o *Store) HasWorldId() bool {
	if o != nil && !IsNil(o.WorldId) {
		return true
	}

	return false
}

// SetWorldId gets a reference to the given string and assigns it to the WorldId field.
func (o *Store) SetWorldId(v string) {
	o.WorldId = &v
}

func (o Store) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Store) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.ListingIds) {
		toSerialize["listingIds"] = o.ListingIds
	}
	if !IsNil(o.Listings) {
		toSerialize["listings"] = o.Listings
	}
	toSerialize["sellerDisplayName"] = o.SellerDisplayName
	toSerialize["sellerId"] = o.SellerId
	if !IsNil(o.ShelfIds) {
		toSerialize["shelfIds"] = o.ShelfIds
	}
	if !IsNil(o.Shelves) {
		toSerialize["shelves"] = o.Shelves
	}
	toSerialize["storeId"] = o.StoreId
	toSerialize["storeType"] = o.StoreType
	toSerialize["tags"] = o.Tags
	if !IsNil(o.WorldId) {
		toSerialize["worldId"] = o.WorldId
	}
	return toSerialize, nil
}

func (o *Store) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"displayName",
		"id",
		"sellerDisplayName",
		"sellerId",
		"storeId",
		"storeType",
		"tags",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varStore := _Store{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStore)

	if err != nil {
		return err
	}

	*o = Store(varStore)

	return err
}

type NullableStore struct {
	value *Store
	isSet bool
}

func (v NullableStore) Get() *Store {
	return v.value
}

func (v *NullableStore) Set(val *Store) {
	v.value = val
	v.isSet = true
}

func (v NullableStore) IsSet() bool {
	return v.isSet
}

func (v *NullableStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStore(val *Store) *NullableStore {
	return &NullableStore{value: val, isSet: true}
}

func (v NullableStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
