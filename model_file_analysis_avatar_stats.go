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
	"bytes"
	"fmt"
)

// checks if the FileAnalysisAvatarStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileAnalysisAvatarStats{}

// FileAnalysisAvatarStats struct for FileAnalysisAvatarStats
type FileAnalysisAvatarStats struct {
	AnimatorCount int32 `json:"animatorCount"`
	AudioSourceCount int32 `json:"audioSourceCount"`
	BlendShapeCount int32 `json:"blendShapeCount"`
	BoneCount int32 `json:"boneCount"`
	Bounds []float32 `json:"bounds"`
	CameraCount *int32 `json:"cameraCount,omitempty"`
	ClothCount int32 `json:"clothCount"`
	ConstraintCount int32 `json:"constraintCount"`
	ConstraintDepth int32 `json:"constraintDepth"`
	ContactCount int32 `json:"contactCount"`
	CustomExpressions bool `json:"customExpressions"`
	CustomizeAnimationLayers bool `json:"customizeAnimationLayers"`
	EnableEyeLook bool `json:"enableEyeLook"`
	LightCount int32 `json:"lightCount"`
	LineRendererCount int32 `json:"lineRendererCount"`
	LipSync int32 `json:"lipSync"`
	MaterialCount int32 `json:"materialCount"`
	MaterialSlotsUsed int32 `json:"materialSlotsUsed"`
	MeshCount int32 `json:"meshCount"`
	MeshIndices int32 `json:"meshIndices"`
	MeshParticleMaxPolygons int32 `json:"meshParticleMaxPolygons"`
	MeshPolygons int32 `json:"meshPolygons"`
	MeshVertices int32 `json:"meshVertices"`
	ParticleCollisionEnabled bool `json:"particleCollisionEnabled"`
	ParticleSystemCount int32 `json:"particleSystemCount"`
	ParticleTrailsEnabled bool `json:"particleTrailsEnabled"`
	PhysBoneColliderCount int32 `json:"physBoneColliderCount"`
	PhysBoneCollisionCheckCount int32 `json:"physBoneCollisionCheckCount"`
	PhysBoneComponentCount int32 `json:"physBoneComponentCount"`
	PhysBoneTransformCount int32 `json:"physBoneTransformCount"`
	PhysicsColliders int32 `json:"physicsColliders"`
	PhysicsRigidbodies int32 `json:"physicsRigidbodies"`
	SkinnedMeshCount int32 `json:"skinnedMeshCount"`
	SkinnedMeshIndices int32 `json:"skinnedMeshIndices"`
	SkinnedMeshPolygons int32 `json:"skinnedMeshPolygons"`
	SkinnedMeshVertices int32 `json:"skinnedMeshVertices"`
	TotalClothVertices int32 `json:"totalClothVertices"`
	TotalIndices int32 `json:"totalIndices"`
	TotalMaxParticles int32 `json:"totalMaxParticles"`
	TotalPolygons int32 `json:"totalPolygons"`
	TotalTextureUsage int32 `json:"totalTextureUsage"`
	TotalVertices int32 `json:"totalVertices"`
	TrailRendererCount int32 `json:"trailRendererCount"`
	WriteDefaultsUsed bool `json:"writeDefaultsUsed"`
}

type _FileAnalysisAvatarStats FileAnalysisAvatarStats

// NewFileAnalysisAvatarStats instantiates a new FileAnalysisAvatarStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileAnalysisAvatarStats(animatorCount int32, audioSourceCount int32, blendShapeCount int32, boneCount int32, bounds []float32, clothCount int32, constraintCount int32, constraintDepth int32, contactCount int32, customExpressions bool, customizeAnimationLayers bool, enableEyeLook bool, lightCount int32, lineRendererCount int32, lipSync int32, materialCount int32, materialSlotsUsed int32, meshCount int32, meshIndices int32, meshParticleMaxPolygons int32, meshPolygons int32, meshVertices int32, particleCollisionEnabled bool, particleSystemCount int32, particleTrailsEnabled bool, physBoneColliderCount int32, physBoneCollisionCheckCount int32, physBoneComponentCount int32, physBoneTransformCount int32, physicsColliders int32, physicsRigidbodies int32, skinnedMeshCount int32, skinnedMeshIndices int32, skinnedMeshPolygons int32, skinnedMeshVertices int32, totalClothVertices int32, totalIndices int32, totalMaxParticles int32, totalPolygons int32, totalTextureUsage int32, totalVertices int32, trailRendererCount int32, writeDefaultsUsed bool) *FileAnalysisAvatarStats {
	this := FileAnalysisAvatarStats{}
	this.AnimatorCount = animatorCount
	this.AudioSourceCount = audioSourceCount
	this.BlendShapeCount = blendShapeCount
	this.BoneCount = boneCount
	this.Bounds = bounds
	this.ClothCount = clothCount
	this.ConstraintCount = constraintCount
	this.ConstraintDepth = constraintDepth
	this.ContactCount = contactCount
	this.CustomExpressions = customExpressions
	this.CustomizeAnimationLayers = customizeAnimationLayers
	this.EnableEyeLook = enableEyeLook
	this.LightCount = lightCount
	this.LineRendererCount = lineRendererCount
	this.LipSync = lipSync
	this.MaterialCount = materialCount
	this.MaterialSlotsUsed = materialSlotsUsed
	this.MeshCount = meshCount
	this.MeshIndices = meshIndices
	this.MeshParticleMaxPolygons = meshParticleMaxPolygons
	this.MeshPolygons = meshPolygons
	this.MeshVertices = meshVertices
	this.ParticleCollisionEnabled = particleCollisionEnabled
	this.ParticleSystemCount = particleSystemCount
	this.ParticleTrailsEnabled = particleTrailsEnabled
	this.PhysBoneColliderCount = physBoneColliderCount
	this.PhysBoneCollisionCheckCount = physBoneCollisionCheckCount
	this.PhysBoneComponentCount = physBoneComponentCount
	this.PhysBoneTransformCount = physBoneTransformCount
	this.PhysicsColliders = physicsColliders
	this.PhysicsRigidbodies = physicsRigidbodies
	this.SkinnedMeshCount = skinnedMeshCount
	this.SkinnedMeshIndices = skinnedMeshIndices
	this.SkinnedMeshPolygons = skinnedMeshPolygons
	this.SkinnedMeshVertices = skinnedMeshVertices
	this.TotalClothVertices = totalClothVertices
	this.TotalIndices = totalIndices
	this.TotalMaxParticles = totalMaxParticles
	this.TotalPolygons = totalPolygons
	this.TotalTextureUsage = totalTextureUsage
	this.TotalVertices = totalVertices
	this.TrailRendererCount = trailRendererCount
	this.WriteDefaultsUsed = writeDefaultsUsed
	return &this
}

// NewFileAnalysisAvatarStatsWithDefaults instantiates a new FileAnalysisAvatarStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileAnalysisAvatarStatsWithDefaults() *FileAnalysisAvatarStats {
	this := FileAnalysisAvatarStats{}
	return &this
}

// GetAnimatorCount returns the AnimatorCount field value
func (o *FileAnalysisAvatarStats) GetAnimatorCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AnimatorCount
}

// GetAnimatorCountOk returns a tuple with the AnimatorCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetAnimatorCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnimatorCount, true
}

// SetAnimatorCount sets field value
func (o *FileAnalysisAvatarStats) SetAnimatorCount(v int32) {
	o.AnimatorCount = v
}

// GetAudioSourceCount returns the AudioSourceCount field value
func (o *FileAnalysisAvatarStats) GetAudioSourceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AudioSourceCount
}

// GetAudioSourceCountOk returns a tuple with the AudioSourceCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetAudioSourceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudioSourceCount, true
}

// SetAudioSourceCount sets field value
func (o *FileAnalysisAvatarStats) SetAudioSourceCount(v int32) {
	o.AudioSourceCount = v
}

// GetBlendShapeCount returns the BlendShapeCount field value
func (o *FileAnalysisAvatarStats) GetBlendShapeCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlendShapeCount
}

// GetBlendShapeCountOk returns a tuple with the BlendShapeCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetBlendShapeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlendShapeCount, true
}

// SetBlendShapeCount sets field value
func (o *FileAnalysisAvatarStats) SetBlendShapeCount(v int32) {
	o.BlendShapeCount = v
}

// GetBoneCount returns the BoneCount field value
func (o *FileAnalysisAvatarStats) GetBoneCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BoneCount
}

// GetBoneCountOk returns a tuple with the BoneCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetBoneCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoneCount, true
}

// SetBoneCount sets field value
func (o *FileAnalysisAvatarStats) SetBoneCount(v int32) {
	o.BoneCount = v
}

// GetBounds returns the Bounds field value
func (o *FileAnalysisAvatarStats) GetBounds() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Bounds
}

// GetBoundsOk returns a tuple with the Bounds field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetBoundsOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bounds, true
}

// SetBounds sets field value
func (o *FileAnalysisAvatarStats) SetBounds(v []float32) {
	o.Bounds = v
}

// GetCameraCount returns the CameraCount field value if set, zero value otherwise.
func (o *FileAnalysisAvatarStats) GetCameraCount() int32 {
	if o == nil || IsNil(o.CameraCount) {
		var ret int32
		return ret
	}
	return *o.CameraCount
}

// GetCameraCountOk returns a tuple with the CameraCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetCameraCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CameraCount) {
		return nil, false
	}
	return o.CameraCount, true
}

// HasCameraCount returns a boolean if a field has been set.
func (o *FileAnalysisAvatarStats) HasCameraCount() bool {
	if o != nil && !IsNil(o.CameraCount) {
		return true
	}

	return false
}

// SetCameraCount gets a reference to the given int32 and assigns it to the CameraCount field.
func (o *FileAnalysisAvatarStats) SetCameraCount(v int32) {
	o.CameraCount = &v
}

// GetClothCount returns the ClothCount field value
func (o *FileAnalysisAvatarStats) GetClothCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClothCount
}

// GetClothCountOk returns a tuple with the ClothCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetClothCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClothCount, true
}

// SetClothCount sets field value
func (o *FileAnalysisAvatarStats) SetClothCount(v int32) {
	o.ClothCount = v
}

// GetConstraintCount returns the ConstraintCount field value
func (o *FileAnalysisAvatarStats) GetConstraintCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConstraintCount
}

// GetConstraintCountOk returns a tuple with the ConstraintCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetConstraintCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConstraintCount, true
}

// SetConstraintCount sets field value
func (o *FileAnalysisAvatarStats) SetConstraintCount(v int32) {
	o.ConstraintCount = v
}

// GetConstraintDepth returns the ConstraintDepth field value
func (o *FileAnalysisAvatarStats) GetConstraintDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConstraintDepth
}

// GetConstraintDepthOk returns a tuple with the ConstraintDepth field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetConstraintDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConstraintDepth, true
}

// SetConstraintDepth sets field value
func (o *FileAnalysisAvatarStats) SetConstraintDepth(v int32) {
	o.ConstraintDepth = v
}

// GetContactCount returns the ContactCount field value
func (o *FileAnalysisAvatarStats) GetContactCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContactCount
}

// GetContactCountOk returns a tuple with the ContactCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetContactCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactCount, true
}

// SetContactCount sets field value
func (o *FileAnalysisAvatarStats) SetContactCount(v int32) {
	o.ContactCount = v
}

// GetCustomExpressions returns the CustomExpressions field value
func (o *FileAnalysisAvatarStats) GetCustomExpressions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CustomExpressions
}

// GetCustomExpressionsOk returns a tuple with the CustomExpressions field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetCustomExpressionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomExpressions, true
}

// SetCustomExpressions sets field value
func (o *FileAnalysisAvatarStats) SetCustomExpressions(v bool) {
	o.CustomExpressions = v
}

// GetCustomizeAnimationLayers returns the CustomizeAnimationLayers field value
func (o *FileAnalysisAvatarStats) GetCustomizeAnimationLayers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CustomizeAnimationLayers
}

// GetCustomizeAnimationLayersOk returns a tuple with the CustomizeAnimationLayers field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetCustomizeAnimationLayersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomizeAnimationLayers, true
}

// SetCustomizeAnimationLayers sets field value
func (o *FileAnalysisAvatarStats) SetCustomizeAnimationLayers(v bool) {
	o.CustomizeAnimationLayers = v
}

// GetEnableEyeLook returns the EnableEyeLook field value
func (o *FileAnalysisAvatarStats) GetEnableEyeLook() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableEyeLook
}

// GetEnableEyeLookOk returns a tuple with the EnableEyeLook field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetEnableEyeLookOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableEyeLook, true
}

// SetEnableEyeLook sets field value
func (o *FileAnalysisAvatarStats) SetEnableEyeLook(v bool) {
	o.EnableEyeLook = v
}

// GetLightCount returns the LightCount field value
func (o *FileAnalysisAvatarStats) GetLightCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LightCount
}

// GetLightCountOk returns a tuple with the LightCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetLightCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LightCount, true
}

// SetLightCount sets field value
func (o *FileAnalysisAvatarStats) SetLightCount(v int32) {
	o.LightCount = v
}

// GetLineRendererCount returns the LineRendererCount field value
func (o *FileAnalysisAvatarStats) GetLineRendererCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LineRendererCount
}

// GetLineRendererCountOk returns a tuple with the LineRendererCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetLineRendererCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LineRendererCount, true
}

// SetLineRendererCount sets field value
func (o *FileAnalysisAvatarStats) SetLineRendererCount(v int32) {
	o.LineRendererCount = v
}

// GetLipSync returns the LipSync field value
func (o *FileAnalysisAvatarStats) GetLipSync() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LipSync
}

// GetLipSyncOk returns a tuple with the LipSync field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetLipSyncOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LipSync, true
}

// SetLipSync sets field value
func (o *FileAnalysisAvatarStats) SetLipSync(v int32) {
	o.LipSync = v
}

// GetMaterialCount returns the MaterialCount field value
func (o *FileAnalysisAvatarStats) GetMaterialCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaterialCount
}

// GetMaterialCountOk returns a tuple with the MaterialCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetMaterialCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaterialCount, true
}

// SetMaterialCount sets field value
func (o *FileAnalysisAvatarStats) SetMaterialCount(v int32) {
	o.MaterialCount = v
}

// GetMaterialSlotsUsed returns the MaterialSlotsUsed field value
func (o *FileAnalysisAvatarStats) GetMaterialSlotsUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaterialSlotsUsed
}

// GetMaterialSlotsUsedOk returns a tuple with the MaterialSlotsUsed field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetMaterialSlotsUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaterialSlotsUsed, true
}

// SetMaterialSlotsUsed sets field value
func (o *FileAnalysisAvatarStats) SetMaterialSlotsUsed(v int32) {
	o.MaterialSlotsUsed = v
}

// GetMeshCount returns the MeshCount field value
func (o *FileAnalysisAvatarStats) GetMeshCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MeshCount
}

// GetMeshCountOk returns a tuple with the MeshCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetMeshCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeshCount, true
}

// SetMeshCount sets field value
func (o *FileAnalysisAvatarStats) SetMeshCount(v int32) {
	o.MeshCount = v
}

// GetMeshIndices returns the MeshIndices field value
func (o *FileAnalysisAvatarStats) GetMeshIndices() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MeshIndices
}

// GetMeshIndicesOk returns a tuple with the MeshIndices field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetMeshIndicesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeshIndices, true
}

// SetMeshIndices sets field value
func (o *FileAnalysisAvatarStats) SetMeshIndices(v int32) {
	o.MeshIndices = v
}

// GetMeshParticleMaxPolygons returns the MeshParticleMaxPolygons field value
func (o *FileAnalysisAvatarStats) GetMeshParticleMaxPolygons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MeshParticleMaxPolygons
}

// GetMeshParticleMaxPolygonsOk returns a tuple with the MeshParticleMaxPolygons field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetMeshParticleMaxPolygonsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeshParticleMaxPolygons, true
}

// SetMeshParticleMaxPolygons sets field value
func (o *FileAnalysisAvatarStats) SetMeshParticleMaxPolygons(v int32) {
	o.MeshParticleMaxPolygons = v
}

// GetMeshPolygons returns the MeshPolygons field value
func (o *FileAnalysisAvatarStats) GetMeshPolygons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MeshPolygons
}

// GetMeshPolygonsOk returns a tuple with the MeshPolygons field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetMeshPolygonsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeshPolygons, true
}

// SetMeshPolygons sets field value
func (o *FileAnalysisAvatarStats) SetMeshPolygons(v int32) {
	o.MeshPolygons = v
}

// GetMeshVertices returns the MeshVertices field value
func (o *FileAnalysisAvatarStats) GetMeshVertices() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MeshVertices
}

// GetMeshVerticesOk returns a tuple with the MeshVertices field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetMeshVerticesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeshVertices, true
}

// SetMeshVertices sets field value
func (o *FileAnalysisAvatarStats) SetMeshVertices(v int32) {
	o.MeshVertices = v
}

// GetParticleCollisionEnabled returns the ParticleCollisionEnabled field value
func (o *FileAnalysisAvatarStats) GetParticleCollisionEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ParticleCollisionEnabled
}

// GetParticleCollisionEnabledOk returns a tuple with the ParticleCollisionEnabled field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetParticleCollisionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParticleCollisionEnabled, true
}

// SetParticleCollisionEnabled sets field value
func (o *FileAnalysisAvatarStats) SetParticleCollisionEnabled(v bool) {
	o.ParticleCollisionEnabled = v
}

// GetParticleSystemCount returns the ParticleSystemCount field value
func (o *FileAnalysisAvatarStats) GetParticleSystemCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParticleSystemCount
}

// GetParticleSystemCountOk returns a tuple with the ParticleSystemCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetParticleSystemCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParticleSystemCount, true
}

// SetParticleSystemCount sets field value
func (o *FileAnalysisAvatarStats) SetParticleSystemCount(v int32) {
	o.ParticleSystemCount = v
}

// GetParticleTrailsEnabled returns the ParticleTrailsEnabled field value
func (o *FileAnalysisAvatarStats) GetParticleTrailsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ParticleTrailsEnabled
}

// GetParticleTrailsEnabledOk returns a tuple with the ParticleTrailsEnabled field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetParticleTrailsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParticleTrailsEnabled, true
}

// SetParticleTrailsEnabled sets field value
func (o *FileAnalysisAvatarStats) SetParticleTrailsEnabled(v bool) {
	o.ParticleTrailsEnabled = v
}

// GetPhysBoneColliderCount returns the PhysBoneColliderCount field value
func (o *FileAnalysisAvatarStats) GetPhysBoneColliderCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PhysBoneColliderCount
}

// GetPhysBoneColliderCountOk returns a tuple with the PhysBoneColliderCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetPhysBoneColliderCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhysBoneColliderCount, true
}

// SetPhysBoneColliderCount sets field value
func (o *FileAnalysisAvatarStats) SetPhysBoneColliderCount(v int32) {
	o.PhysBoneColliderCount = v
}

// GetPhysBoneCollisionCheckCount returns the PhysBoneCollisionCheckCount field value
func (o *FileAnalysisAvatarStats) GetPhysBoneCollisionCheckCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PhysBoneCollisionCheckCount
}

// GetPhysBoneCollisionCheckCountOk returns a tuple with the PhysBoneCollisionCheckCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetPhysBoneCollisionCheckCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhysBoneCollisionCheckCount, true
}

// SetPhysBoneCollisionCheckCount sets field value
func (o *FileAnalysisAvatarStats) SetPhysBoneCollisionCheckCount(v int32) {
	o.PhysBoneCollisionCheckCount = v
}

// GetPhysBoneComponentCount returns the PhysBoneComponentCount field value
func (o *FileAnalysisAvatarStats) GetPhysBoneComponentCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PhysBoneComponentCount
}

// GetPhysBoneComponentCountOk returns a tuple with the PhysBoneComponentCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetPhysBoneComponentCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhysBoneComponentCount, true
}

// SetPhysBoneComponentCount sets field value
func (o *FileAnalysisAvatarStats) SetPhysBoneComponentCount(v int32) {
	o.PhysBoneComponentCount = v
}

// GetPhysBoneTransformCount returns the PhysBoneTransformCount field value
func (o *FileAnalysisAvatarStats) GetPhysBoneTransformCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PhysBoneTransformCount
}

// GetPhysBoneTransformCountOk returns a tuple with the PhysBoneTransformCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetPhysBoneTransformCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhysBoneTransformCount, true
}

// SetPhysBoneTransformCount sets field value
func (o *FileAnalysisAvatarStats) SetPhysBoneTransformCount(v int32) {
	o.PhysBoneTransformCount = v
}

// GetPhysicsColliders returns the PhysicsColliders field value
func (o *FileAnalysisAvatarStats) GetPhysicsColliders() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PhysicsColliders
}

// GetPhysicsCollidersOk returns a tuple with the PhysicsColliders field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetPhysicsCollidersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhysicsColliders, true
}

// SetPhysicsColliders sets field value
func (o *FileAnalysisAvatarStats) SetPhysicsColliders(v int32) {
	o.PhysicsColliders = v
}

// GetPhysicsRigidbodies returns the PhysicsRigidbodies field value
func (o *FileAnalysisAvatarStats) GetPhysicsRigidbodies() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PhysicsRigidbodies
}

// GetPhysicsRigidbodiesOk returns a tuple with the PhysicsRigidbodies field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetPhysicsRigidbodiesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhysicsRigidbodies, true
}

// SetPhysicsRigidbodies sets field value
func (o *FileAnalysisAvatarStats) SetPhysicsRigidbodies(v int32) {
	o.PhysicsRigidbodies = v
}

// GetSkinnedMeshCount returns the SkinnedMeshCount field value
func (o *FileAnalysisAvatarStats) GetSkinnedMeshCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SkinnedMeshCount
}

// GetSkinnedMeshCountOk returns a tuple with the SkinnedMeshCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetSkinnedMeshCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkinnedMeshCount, true
}

// SetSkinnedMeshCount sets field value
func (o *FileAnalysisAvatarStats) SetSkinnedMeshCount(v int32) {
	o.SkinnedMeshCount = v
}

// GetSkinnedMeshIndices returns the SkinnedMeshIndices field value
func (o *FileAnalysisAvatarStats) GetSkinnedMeshIndices() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SkinnedMeshIndices
}

// GetSkinnedMeshIndicesOk returns a tuple with the SkinnedMeshIndices field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetSkinnedMeshIndicesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkinnedMeshIndices, true
}

// SetSkinnedMeshIndices sets field value
func (o *FileAnalysisAvatarStats) SetSkinnedMeshIndices(v int32) {
	o.SkinnedMeshIndices = v
}

// GetSkinnedMeshPolygons returns the SkinnedMeshPolygons field value
func (o *FileAnalysisAvatarStats) GetSkinnedMeshPolygons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SkinnedMeshPolygons
}

// GetSkinnedMeshPolygonsOk returns a tuple with the SkinnedMeshPolygons field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetSkinnedMeshPolygonsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkinnedMeshPolygons, true
}

// SetSkinnedMeshPolygons sets field value
func (o *FileAnalysisAvatarStats) SetSkinnedMeshPolygons(v int32) {
	o.SkinnedMeshPolygons = v
}

// GetSkinnedMeshVertices returns the SkinnedMeshVertices field value
func (o *FileAnalysisAvatarStats) GetSkinnedMeshVertices() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SkinnedMeshVertices
}

// GetSkinnedMeshVerticesOk returns a tuple with the SkinnedMeshVertices field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetSkinnedMeshVerticesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkinnedMeshVertices, true
}

// SetSkinnedMeshVertices sets field value
func (o *FileAnalysisAvatarStats) SetSkinnedMeshVertices(v int32) {
	o.SkinnedMeshVertices = v
}

// GetTotalClothVertices returns the TotalClothVertices field value
func (o *FileAnalysisAvatarStats) GetTotalClothVertices() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalClothVertices
}

// GetTotalClothVerticesOk returns a tuple with the TotalClothVertices field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetTotalClothVerticesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalClothVertices, true
}

// SetTotalClothVertices sets field value
func (o *FileAnalysisAvatarStats) SetTotalClothVertices(v int32) {
	o.TotalClothVertices = v
}

// GetTotalIndices returns the TotalIndices field value
func (o *FileAnalysisAvatarStats) GetTotalIndices() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalIndices
}

// GetTotalIndicesOk returns a tuple with the TotalIndices field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetTotalIndicesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalIndices, true
}

// SetTotalIndices sets field value
func (o *FileAnalysisAvatarStats) SetTotalIndices(v int32) {
	o.TotalIndices = v
}

// GetTotalMaxParticles returns the TotalMaxParticles field value
func (o *FileAnalysisAvatarStats) GetTotalMaxParticles() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalMaxParticles
}

// GetTotalMaxParticlesOk returns a tuple with the TotalMaxParticles field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetTotalMaxParticlesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMaxParticles, true
}

// SetTotalMaxParticles sets field value
func (o *FileAnalysisAvatarStats) SetTotalMaxParticles(v int32) {
	o.TotalMaxParticles = v
}

// GetTotalPolygons returns the TotalPolygons field value
func (o *FileAnalysisAvatarStats) GetTotalPolygons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPolygons
}

// GetTotalPolygonsOk returns a tuple with the TotalPolygons field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetTotalPolygonsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPolygons, true
}

// SetTotalPolygons sets field value
func (o *FileAnalysisAvatarStats) SetTotalPolygons(v int32) {
	o.TotalPolygons = v
}

// GetTotalTextureUsage returns the TotalTextureUsage field value
func (o *FileAnalysisAvatarStats) GetTotalTextureUsage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTextureUsage
}

// GetTotalTextureUsageOk returns a tuple with the TotalTextureUsage field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetTotalTextureUsageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTextureUsage, true
}

// SetTotalTextureUsage sets field value
func (o *FileAnalysisAvatarStats) SetTotalTextureUsage(v int32) {
	o.TotalTextureUsage = v
}

// GetTotalVertices returns the TotalVertices field value
func (o *FileAnalysisAvatarStats) GetTotalVertices() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalVertices
}

// GetTotalVerticesOk returns a tuple with the TotalVertices field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetTotalVerticesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalVertices, true
}

// SetTotalVertices sets field value
func (o *FileAnalysisAvatarStats) SetTotalVertices(v int32) {
	o.TotalVertices = v
}

// GetTrailRendererCount returns the TrailRendererCount field value
func (o *FileAnalysisAvatarStats) GetTrailRendererCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrailRendererCount
}

// GetTrailRendererCountOk returns a tuple with the TrailRendererCount field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetTrailRendererCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrailRendererCount, true
}

// SetTrailRendererCount sets field value
func (o *FileAnalysisAvatarStats) SetTrailRendererCount(v int32) {
	o.TrailRendererCount = v
}

// GetWriteDefaultsUsed returns the WriteDefaultsUsed field value
func (o *FileAnalysisAvatarStats) GetWriteDefaultsUsed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WriteDefaultsUsed
}

// GetWriteDefaultsUsedOk returns a tuple with the WriteDefaultsUsed field value
// and a boolean to check if the value has been set.
func (o *FileAnalysisAvatarStats) GetWriteDefaultsUsedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WriteDefaultsUsed, true
}

// SetWriteDefaultsUsed sets field value
func (o *FileAnalysisAvatarStats) SetWriteDefaultsUsed(v bool) {
	o.WriteDefaultsUsed = v
}

func (o FileAnalysisAvatarStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileAnalysisAvatarStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["animatorCount"] = o.AnimatorCount
	toSerialize["audioSourceCount"] = o.AudioSourceCount
	toSerialize["blendShapeCount"] = o.BlendShapeCount
	toSerialize["boneCount"] = o.BoneCount
	toSerialize["bounds"] = o.Bounds
	if !IsNil(o.CameraCount) {
		toSerialize["cameraCount"] = o.CameraCount
	}
	toSerialize["clothCount"] = o.ClothCount
	toSerialize["constraintCount"] = o.ConstraintCount
	toSerialize["constraintDepth"] = o.ConstraintDepth
	toSerialize["contactCount"] = o.ContactCount
	toSerialize["customExpressions"] = o.CustomExpressions
	toSerialize["customizeAnimationLayers"] = o.CustomizeAnimationLayers
	toSerialize["enableEyeLook"] = o.EnableEyeLook
	toSerialize["lightCount"] = o.LightCount
	toSerialize["lineRendererCount"] = o.LineRendererCount
	toSerialize["lipSync"] = o.LipSync
	toSerialize["materialCount"] = o.MaterialCount
	toSerialize["materialSlotsUsed"] = o.MaterialSlotsUsed
	toSerialize["meshCount"] = o.MeshCount
	toSerialize["meshIndices"] = o.MeshIndices
	toSerialize["meshParticleMaxPolygons"] = o.MeshParticleMaxPolygons
	toSerialize["meshPolygons"] = o.MeshPolygons
	toSerialize["meshVertices"] = o.MeshVertices
	toSerialize["particleCollisionEnabled"] = o.ParticleCollisionEnabled
	toSerialize["particleSystemCount"] = o.ParticleSystemCount
	toSerialize["particleTrailsEnabled"] = o.ParticleTrailsEnabled
	toSerialize["physBoneColliderCount"] = o.PhysBoneColliderCount
	toSerialize["physBoneCollisionCheckCount"] = o.PhysBoneCollisionCheckCount
	toSerialize["physBoneComponentCount"] = o.PhysBoneComponentCount
	toSerialize["physBoneTransformCount"] = o.PhysBoneTransformCount
	toSerialize["physicsColliders"] = o.PhysicsColliders
	toSerialize["physicsRigidbodies"] = o.PhysicsRigidbodies
	toSerialize["skinnedMeshCount"] = o.SkinnedMeshCount
	toSerialize["skinnedMeshIndices"] = o.SkinnedMeshIndices
	toSerialize["skinnedMeshPolygons"] = o.SkinnedMeshPolygons
	toSerialize["skinnedMeshVertices"] = o.SkinnedMeshVertices
	toSerialize["totalClothVertices"] = o.TotalClothVertices
	toSerialize["totalIndices"] = o.TotalIndices
	toSerialize["totalMaxParticles"] = o.TotalMaxParticles
	toSerialize["totalPolygons"] = o.TotalPolygons
	toSerialize["totalTextureUsage"] = o.TotalTextureUsage
	toSerialize["totalVertices"] = o.TotalVertices
	toSerialize["trailRendererCount"] = o.TrailRendererCount
	toSerialize["writeDefaultsUsed"] = o.WriteDefaultsUsed
	return toSerialize, nil
}

func (o *FileAnalysisAvatarStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"animatorCount",
		"audioSourceCount",
		"blendShapeCount",
		"boneCount",
		"bounds",
		"clothCount",
		"constraintCount",
		"constraintDepth",
		"contactCount",
		"customExpressions",
		"customizeAnimationLayers",
		"enableEyeLook",
		"lightCount",
		"lineRendererCount",
		"lipSync",
		"materialCount",
		"materialSlotsUsed",
		"meshCount",
		"meshIndices",
		"meshParticleMaxPolygons",
		"meshPolygons",
		"meshVertices",
		"particleCollisionEnabled",
		"particleSystemCount",
		"particleTrailsEnabled",
		"physBoneColliderCount",
		"physBoneCollisionCheckCount",
		"physBoneComponentCount",
		"physBoneTransformCount",
		"physicsColliders",
		"physicsRigidbodies",
		"skinnedMeshCount",
		"skinnedMeshIndices",
		"skinnedMeshPolygons",
		"skinnedMeshVertices",
		"totalClothVertices",
		"totalIndices",
		"totalMaxParticles",
		"totalPolygons",
		"totalTextureUsage",
		"totalVertices",
		"trailRendererCount",
		"writeDefaultsUsed",
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

	varFileAnalysisAvatarStats := _FileAnalysisAvatarStats{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileAnalysisAvatarStats)

	if err != nil {
		return err
	}

	*o = FileAnalysisAvatarStats(varFileAnalysisAvatarStats)

	return err
}

type NullableFileAnalysisAvatarStats struct {
	value *FileAnalysisAvatarStats
	isSet bool
}

func (v NullableFileAnalysisAvatarStats) Get() *FileAnalysisAvatarStats {
	return v.value
}

func (v *NullableFileAnalysisAvatarStats) Set(val *FileAnalysisAvatarStats) {
	v.value = val
	v.isSet = true
}

func (v NullableFileAnalysisAvatarStats) IsSet() bool {
	return v.isSet
}

func (v *NullableFileAnalysisAvatarStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileAnalysisAvatarStats(val *FileAnalysisAvatarStats) *NullableFileAnalysisAvatarStats {
	return &NullableFileAnalysisAvatarStats{value: val, isSet: true}
}

func (v NullableFileAnalysisAvatarStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileAnalysisAvatarStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


