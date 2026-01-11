# FileAnalysisAvatarStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnimatorCount** | **int32** |  | 
**AudioSourceCount** | **int32** |  | 
**BlendShapeCount** | **int32** |  | 
**BoneCount** | **int32** |  | 
**Bounds** | **[]float32** |  | 
**CameraCount** | Pointer to **int32** |  | [optional] 
**ClothCount** | **int32** |  | 
**ConstraintCount** | **int32** |  | 
**ConstraintDepth** | **int32** |  | 
**ContactCount** | **int32** |  | 
**CustomExpressions** | **bool** |  | 
**CustomizeAnimationLayers** | **bool** |  | 
**EnableEyeLook** | **bool** |  | 
**LightCount** | **int32** |  | 
**LineRendererCount** | **int32** |  | 
**LipSync** | **int32** |  | 
**MaterialCount** | **int32** |  | 
**MaterialSlotsUsed** | **int32** |  | 
**MeshCount** | **int32** |  | 
**MeshIndices** | **int32** |  | 
**MeshParticleMaxPolygons** | **int32** |  | 
**MeshPolygons** | **int32** |  | 
**MeshVertices** | **int32** |  | 
**ParticleCollisionEnabled** | **bool** |  | 
**ParticleSystemCount** | **int32** |  | 
**ParticleTrailsEnabled** | **bool** |  | 
**PhysBoneColliderCount** | **int32** |  | 
**PhysBoneCollisionCheckCount** | **int32** |  | 
**PhysBoneComponentCount** | **int32** |  | 
**PhysBoneTransformCount** | **int32** |  | 
**PhysicsColliders** | **int32** |  | 
**PhysicsRigidbodies** | **int32** |  | 
**SkinnedMeshCount** | **int32** |  | 
**SkinnedMeshIndices** | **int32** |  | 
**SkinnedMeshPolygons** | **int32** |  | 
**SkinnedMeshVertices** | **int32** |  | 
**TotalClothVertices** | **int32** |  | 
**TotalIndices** | **int32** |  | 
**TotalMaxParticles** | **int32** |  | 
**TotalPolygons** | **int32** |  | 
**TotalTextureUsage** | **int32** |  | 
**TotalVertices** | **int32** |  | 
**TrailRendererCount** | **int32** |  | 
**WriteDefaultsUsed** | **bool** |  | 

## Methods

### NewFileAnalysisAvatarStats

`func NewFileAnalysisAvatarStats(animatorCount int32, audioSourceCount int32, blendShapeCount int32, boneCount int32, bounds []float32, clothCount int32, constraintCount int32, constraintDepth int32, contactCount int32, customExpressions bool, customizeAnimationLayers bool, enableEyeLook bool, lightCount int32, lineRendererCount int32, lipSync int32, materialCount int32, materialSlotsUsed int32, meshCount int32, meshIndices int32, meshParticleMaxPolygons int32, meshPolygons int32, meshVertices int32, particleCollisionEnabled bool, particleSystemCount int32, particleTrailsEnabled bool, physBoneColliderCount int32, physBoneCollisionCheckCount int32, physBoneComponentCount int32, physBoneTransformCount int32, physicsColliders int32, physicsRigidbodies int32, skinnedMeshCount int32, skinnedMeshIndices int32, skinnedMeshPolygons int32, skinnedMeshVertices int32, totalClothVertices int32, totalIndices int32, totalMaxParticles int32, totalPolygons int32, totalTextureUsage int32, totalVertices int32, trailRendererCount int32, writeDefaultsUsed bool, ) *FileAnalysisAvatarStats`

NewFileAnalysisAvatarStats instantiates a new FileAnalysisAvatarStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileAnalysisAvatarStatsWithDefaults

`func NewFileAnalysisAvatarStatsWithDefaults() *FileAnalysisAvatarStats`

NewFileAnalysisAvatarStatsWithDefaults instantiates a new FileAnalysisAvatarStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnimatorCount

`func (o *FileAnalysisAvatarStats) GetAnimatorCount() int32`

GetAnimatorCount returns the AnimatorCount field if non-nil, zero value otherwise.

### GetAnimatorCountOk

`func (o *FileAnalysisAvatarStats) GetAnimatorCountOk() (*int32, bool)`

GetAnimatorCountOk returns a tuple with the AnimatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimatorCount

`func (o *FileAnalysisAvatarStats) SetAnimatorCount(v int32)`

SetAnimatorCount sets AnimatorCount field to given value.


### GetAudioSourceCount

`func (o *FileAnalysisAvatarStats) GetAudioSourceCount() int32`

GetAudioSourceCount returns the AudioSourceCount field if non-nil, zero value otherwise.

### GetAudioSourceCountOk

`func (o *FileAnalysisAvatarStats) GetAudioSourceCountOk() (*int32, bool)`

GetAudioSourceCountOk returns a tuple with the AudioSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioSourceCount

`func (o *FileAnalysisAvatarStats) SetAudioSourceCount(v int32)`

SetAudioSourceCount sets AudioSourceCount field to given value.


### GetBlendShapeCount

`func (o *FileAnalysisAvatarStats) GetBlendShapeCount() int32`

GetBlendShapeCount returns the BlendShapeCount field if non-nil, zero value otherwise.

### GetBlendShapeCountOk

`func (o *FileAnalysisAvatarStats) GetBlendShapeCountOk() (*int32, bool)`

GetBlendShapeCountOk returns a tuple with the BlendShapeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlendShapeCount

`func (o *FileAnalysisAvatarStats) SetBlendShapeCount(v int32)`

SetBlendShapeCount sets BlendShapeCount field to given value.


### GetBoneCount

`func (o *FileAnalysisAvatarStats) GetBoneCount() int32`

GetBoneCount returns the BoneCount field if non-nil, zero value otherwise.

### GetBoneCountOk

`func (o *FileAnalysisAvatarStats) GetBoneCountOk() (*int32, bool)`

GetBoneCountOk returns a tuple with the BoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoneCount

`func (o *FileAnalysisAvatarStats) SetBoneCount(v int32)`

SetBoneCount sets BoneCount field to given value.


### GetBounds

`func (o *FileAnalysisAvatarStats) GetBounds() []float32`

GetBounds returns the Bounds field if non-nil, zero value otherwise.

### GetBoundsOk

`func (o *FileAnalysisAvatarStats) GetBoundsOk() (*[]float32, bool)`

GetBoundsOk returns a tuple with the Bounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounds

`func (o *FileAnalysisAvatarStats) SetBounds(v []float32)`

SetBounds sets Bounds field to given value.


### GetCameraCount

`func (o *FileAnalysisAvatarStats) GetCameraCount() int32`

GetCameraCount returns the CameraCount field if non-nil, zero value otherwise.

### GetCameraCountOk

`func (o *FileAnalysisAvatarStats) GetCameraCountOk() (*int32, bool)`

GetCameraCountOk returns a tuple with the CameraCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraCount

`func (o *FileAnalysisAvatarStats) SetCameraCount(v int32)`

SetCameraCount sets CameraCount field to given value.

### HasCameraCount

`func (o *FileAnalysisAvatarStats) HasCameraCount() bool`

HasCameraCount returns a boolean if a field has been set.

### GetClothCount

`func (o *FileAnalysisAvatarStats) GetClothCount() int32`

GetClothCount returns the ClothCount field if non-nil, zero value otherwise.

### GetClothCountOk

`func (o *FileAnalysisAvatarStats) GetClothCountOk() (*int32, bool)`

GetClothCountOk returns a tuple with the ClothCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClothCount

`func (o *FileAnalysisAvatarStats) SetClothCount(v int32)`

SetClothCount sets ClothCount field to given value.


### GetConstraintCount

`func (o *FileAnalysisAvatarStats) GetConstraintCount() int32`

GetConstraintCount returns the ConstraintCount field if non-nil, zero value otherwise.

### GetConstraintCountOk

`func (o *FileAnalysisAvatarStats) GetConstraintCountOk() (*int32, bool)`

GetConstraintCountOk returns a tuple with the ConstraintCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintCount

`func (o *FileAnalysisAvatarStats) SetConstraintCount(v int32)`

SetConstraintCount sets ConstraintCount field to given value.


### GetConstraintDepth

`func (o *FileAnalysisAvatarStats) GetConstraintDepth() int32`

GetConstraintDepth returns the ConstraintDepth field if non-nil, zero value otherwise.

### GetConstraintDepthOk

`func (o *FileAnalysisAvatarStats) GetConstraintDepthOk() (*int32, bool)`

GetConstraintDepthOk returns a tuple with the ConstraintDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintDepth

`func (o *FileAnalysisAvatarStats) SetConstraintDepth(v int32)`

SetConstraintDepth sets ConstraintDepth field to given value.


### GetContactCount

`func (o *FileAnalysisAvatarStats) GetContactCount() int32`

GetContactCount returns the ContactCount field if non-nil, zero value otherwise.

### GetContactCountOk

`func (o *FileAnalysisAvatarStats) GetContactCountOk() (*int32, bool)`

GetContactCountOk returns a tuple with the ContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactCount

`func (o *FileAnalysisAvatarStats) SetContactCount(v int32)`

SetContactCount sets ContactCount field to given value.


### GetCustomExpressions

`func (o *FileAnalysisAvatarStats) GetCustomExpressions() bool`

GetCustomExpressions returns the CustomExpressions field if non-nil, zero value otherwise.

### GetCustomExpressionsOk

`func (o *FileAnalysisAvatarStats) GetCustomExpressionsOk() (*bool, bool)`

GetCustomExpressionsOk returns a tuple with the CustomExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomExpressions

`func (o *FileAnalysisAvatarStats) SetCustomExpressions(v bool)`

SetCustomExpressions sets CustomExpressions field to given value.


### GetCustomizeAnimationLayers

`func (o *FileAnalysisAvatarStats) GetCustomizeAnimationLayers() bool`

GetCustomizeAnimationLayers returns the CustomizeAnimationLayers field if non-nil, zero value otherwise.

### GetCustomizeAnimationLayersOk

`func (o *FileAnalysisAvatarStats) GetCustomizeAnimationLayersOk() (*bool, bool)`

GetCustomizeAnimationLayersOk returns a tuple with the CustomizeAnimationLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizeAnimationLayers

`func (o *FileAnalysisAvatarStats) SetCustomizeAnimationLayers(v bool)`

SetCustomizeAnimationLayers sets CustomizeAnimationLayers field to given value.


### GetEnableEyeLook

`func (o *FileAnalysisAvatarStats) GetEnableEyeLook() bool`

GetEnableEyeLook returns the EnableEyeLook field if non-nil, zero value otherwise.

### GetEnableEyeLookOk

`func (o *FileAnalysisAvatarStats) GetEnableEyeLookOk() (*bool, bool)`

GetEnableEyeLookOk returns a tuple with the EnableEyeLook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEyeLook

`func (o *FileAnalysisAvatarStats) SetEnableEyeLook(v bool)`

SetEnableEyeLook sets EnableEyeLook field to given value.


### GetLightCount

`func (o *FileAnalysisAvatarStats) GetLightCount() int32`

GetLightCount returns the LightCount field if non-nil, zero value otherwise.

### GetLightCountOk

`func (o *FileAnalysisAvatarStats) GetLightCountOk() (*int32, bool)`

GetLightCountOk returns a tuple with the LightCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightCount

`func (o *FileAnalysisAvatarStats) SetLightCount(v int32)`

SetLightCount sets LightCount field to given value.


### GetLineRendererCount

`func (o *FileAnalysisAvatarStats) GetLineRendererCount() int32`

GetLineRendererCount returns the LineRendererCount field if non-nil, zero value otherwise.

### GetLineRendererCountOk

`func (o *FileAnalysisAvatarStats) GetLineRendererCountOk() (*int32, bool)`

GetLineRendererCountOk returns a tuple with the LineRendererCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineRendererCount

`func (o *FileAnalysisAvatarStats) SetLineRendererCount(v int32)`

SetLineRendererCount sets LineRendererCount field to given value.


### GetLipSync

`func (o *FileAnalysisAvatarStats) GetLipSync() int32`

GetLipSync returns the LipSync field if non-nil, zero value otherwise.

### GetLipSyncOk

`func (o *FileAnalysisAvatarStats) GetLipSyncOk() (*int32, bool)`

GetLipSyncOk returns a tuple with the LipSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLipSync

`func (o *FileAnalysisAvatarStats) SetLipSync(v int32)`

SetLipSync sets LipSync field to given value.


### GetMaterialCount

`func (o *FileAnalysisAvatarStats) GetMaterialCount() int32`

GetMaterialCount returns the MaterialCount field if non-nil, zero value otherwise.

### GetMaterialCountOk

`func (o *FileAnalysisAvatarStats) GetMaterialCountOk() (*int32, bool)`

GetMaterialCountOk returns a tuple with the MaterialCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialCount

`func (o *FileAnalysisAvatarStats) SetMaterialCount(v int32)`

SetMaterialCount sets MaterialCount field to given value.


### GetMaterialSlotsUsed

`func (o *FileAnalysisAvatarStats) GetMaterialSlotsUsed() int32`

GetMaterialSlotsUsed returns the MaterialSlotsUsed field if non-nil, zero value otherwise.

### GetMaterialSlotsUsedOk

`func (o *FileAnalysisAvatarStats) GetMaterialSlotsUsedOk() (*int32, bool)`

GetMaterialSlotsUsedOk returns a tuple with the MaterialSlotsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialSlotsUsed

`func (o *FileAnalysisAvatarStats) SetMaterialSlotsUsed(v int32)`

SetMaterialSlotsUsed sets MaterialSlotsUsed field to given value.


### GetMeshCount

`func (o *FileAnalysisAvatarStats) GetMeshCount() int32`

GetMeshCount returns the MeshCount field if non-nil, zero value otherwise.

### GetMeshCountOk

`func (o *FileAnalysisAvatarStats) GetMeshCountOk() (*int32, bool)`

GetMeshCountOk returns a tuple with the MeshCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshCount

`func (o *FileAnalysisAvatarStats) SetMeshCount(v int32)`

SetMeshCount sets MeshCount field to given value.


### GetMeshIndices

`func (o *FileAnalysisAvatarStats) GetMeshIndices() int32`

GetMeshIndices returns the MeshIndices field if non-nil, zero value otherwise.

### GetMeshIndicesOk

`func (o *FileAnalysisAvatarStats) GetMeshIndicesOk() (*int32, bool)`

GetMeshIndicesOk returns a tuple with the MeshIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshIndices

`func (o *FileAnalysisAvatarStats) SetMeshIndices(v int32)`

SetMeshIndices sets MeshIndices field to given value.


### GetMeshParticleMaxPolygons

`func (o *FileAnalysisAvatarStats) GetMeshParticleMaxPolygons() int32`

GetMeshParticleMaxPolygons returns the MeshParticleMaxPolygons field if non-nil, zero value otherwise.

### GetMeshParticleMaxPolygonsOk

`func (o *FileAnalysisAvatarStats) GetMeshParticleMaxPolygonsOk() (*int32, bool)`

GetMeshParticleMaxPolygonsOk returns a tuple with the MeshParticleMaxPolygons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshParticleMaxPolygons

`func (o *FileAnalysisAvatarStats) SetMeshParticleMaxPolygons(v int32)`

SetMeshParticleMaxPolygons sets MeshParticleMaxPolygons field to given value.


### GetMeshPolygons

`func (o *FileAnalysisAvatarStats) GetMeshPolygons() int32`

GetMeshPolygons returns the MeshPolygons field if non-nil, zero value otherwise.

### GetMeshPolygonsOk

`func (o *FileAnalysisAvatarStats) GetMeshPolygonsOk() (*int32, bool)`

GetMeshPolygonsOk returns a tuple with the MeshPolygons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshPolygons

`func (o *FileAnalysisAvatarStats) SetMeshPolygons(v int32)`

SetMeshPolygons sets MeshPolygons field to given value.


### GetMeshVertices

`func (o *FileAnalysisAvatarStats) GetMeshVertices() int32`

GetMeshVertices returns the MeshVertices field if non-nil, zero value otherwise.

### GetMeshVerticesOk

`func (o *FileAnalysisAvatarStats) GetMeshVerticesOk() (*int32, bool)`

GetMeshVerticesOk returns a tuple with the MeshVertices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshVertices

`func (o *FileAnalysisAvatarStats) SetMeshVertices(v int32)`

SetMeshVertices sets MeshVertices field to given value.


### GetParticleCollisionEnabled

`func (o *FileAnalysisAvatarStats) GetParticleCollisionEnabled() bool`

GetParticleCollisionEnabled returns the ParticleCollisionEnabled field if non-nil, zero value otherwise.

### GetParticleCollisionEnabledOk

`func (o *FileAnalysisAvatarStats) GetParticleCollisionEnabledOk() (*bool, bool)`

GetParticleCollisionEnabledOk returns a tuple with the ParticleCollisionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticleCollisionEnabled

`func (o *FileAnalysisAvatarStats) SetParticleCollisionEnabled(v bool)`

SetParticleCollisionEnabled sets ParticleCollisionEnabled field to given value.


### GetParticleSystemCount

`func (o *FileAnalysisAvatarStats) GetParticleSystemCount() int32`

GetParticleSystemCount returns the ParticleSystemCount field if non-nil, zero value otherwise.

### GetParticleSystemCountOk

`func (o *FileAnalysisAvatarStats) GetParticleSystemCountOk() (*int32, bool)`

GetParticleSystemCountOk returns a tuple with the ParticleSystemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticleSystemCount

`func (o *FileAnalysisAvatarStats) SetParticleSystemCount(v int32)`

SetParticleSystemCount sets ParticleSystemCount field to given value.


### GetParticleTrailsEnabled

`func (o *FileAnalysisAvatarStats) GetParticleTrailsEnabled() bool`

GetParticleTrailsEnabled returns the ParticleTrailsEnabled field if non-nil, zero value otherwise.

### GetParticleTrailsEnabledOk

`func (o *FileAnalysisAvatarStats) GetParticleTrailsEnabledOk() (*bool, bool)`

GetParticleTrailsEnabledOk returns a tuple with the ParticleTrailsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticleTrailsEnabled

`func (o *FileAnalysisAvatarStats) SetParticleTrailsEnabled(v bool)`

SetParticleTrailsEnabled sets ParticleTrailsEnabled field to given value.


### GetPhysBoneColliderCount

`func (o *FileAnalysisAvatarStats) GetPhysBoneColliderCount() int32`

GetPhysBoneColliderCount returns the PhysBoneColliderCount field if non-nil, zero value otherwise.

### GetPhysBoneColliderCountOk

`func (o *FileAnalysisAvatarStats) GetPhysBoneColliderCountOk() (*int32, bool)`

GetPhysBoneColliderCountOk returns a tuple with the PhysBoneColliderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysBoneColliderCount

`func (o *FileAnalysisAvatarStats) SetPhysBoneColliderCount(v int32)`

SetPhysBoneColliderCount sets PhysBoneColliderCount field to given value.


### GetPhysBoneCollisionCheckCount

`func (o *FileAnalysisAvatarStats) GetPhysBoneCollisionCheckCount() int32`

GetPhysBoneCollisionCheckCount returns the PhysBoneCollisionCheckCount field if non-nil, zero value otherwise.

### GetPhysBoneCollisionCheckCountOk

`func (o *FileAnalysisAvatarStats) GetPhysBoneCollisionCheckCountOk() (*int32, bool)`

GetPhysBoneCollisionCheckCountOk returns a tuple with the PhysBoneCollisionCheckCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysBoneCollisionCheckCount

`func (o *FileAnalysisAvatarStats) SetPhysBoneCollisionCheckCount(v int32)`

SetPhysBoneCollisionCheckCount sets PhysBoneCollisionCheckCount field to given value.


### GetPhysBoneComponentCount

`func (o *FileAnalysisAvatarStats) GetPhysBoneComponentCount() int32`

GetPhysBoneComponentCount returns the PhysBoneComponentCount field if non-nil, zero value otherwise.

### GetPhysBoneComponentCountOk

`func (o *FileAnalysisAvatarStats) GetPhysBoneComponentCountOk() (*int32, bool)`

GetPhysBoneComponentCountOk returns a tuple with the PhysBoneComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysBoneComponentCount

`func (o *FileAnalysisAvatarStats) SetPhysBoneComponentCount(v int32)`

SetPhysBoneComponentCount sets PhysBoneComponentCount field to given value.


### GetPhysBoneTransformCount

`func (o *FileAnalysisAvatarStats) GetPhysBoneTransformCount() int32`

GetPhysBoneTransformCount returns the PhysBoneTransformCount field if non-nil, zero value otherwise.

### GetPhysBoneTransformCountOk

`func (o *FileAnalysisAvatarStats) GetPhysBoneTransformCountOk() (*int32, bool)`

GetPhysBoneTransformCountOk returns a tuple with the PhysBoneTransformCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysBoneTransformCount

`func (o *FileAnalysisAvatarStats) SetPhysBoneTransformCount(v int32)`

SetPhysBoneTransformCount sets PhysBoneTransformCount field to given value.


### GetPhysicsColliders

`func (o *FileAnalysisAvatarStats) GetPhysicsColliders() int32`

GetPhysicsColliders returns the PhysicsColliders field if non-nil, zero value otherwise.

### GetPhysicsCollidersOk

`func (o *FileAnalysisAvatarStats) GetPhysicsCollidersOk() (*int32, bool)`

GetPhysicsCollidersOk returns a tuple with the PhysicsColliders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicsColliders

`func (o *FileAnalysisAvatarStats) SetPhysicsColliders(v int32)`

SetPhysicsColliders sets PhysicsColliders field to given value.


### GetPhysicsRigidbodies

`func (o *FileAnalysisAvatarStats) GetPhysicsRigidbodies() int32`

GetPhysicsRigidbodies returns the PhysicsRigidbodies field if non-nil, zero value otherwise.

### GetPhysicsRigidbodiesOk

`func (o *FileAnalysisAvatarStats) GetPhysicsRigidbodiesOk() (*int32, bool)`

GetPhysicsRigidbodiesOk returns a tuple with the PhysicsRigidbodies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicsRigidbodies

`func (o *FileAnalysisAvatarStats) SetPhysicsRigidbodies(v int32)`

SetPhysicsRigidbodies sets PhysicsRigidbodies field to given value.


### GetSkinnedMeshCount

`func (o *FileAnalysisAvatarStats) GetSkinnedMeshCount() int32`

GetSkinnedMeshCount returns the SkinnedMeshCount field if non-nil, zero value otherwise.

### GetSkinnedMeshCountOk

`func (o *FileAnalysisAvatarStats) GetSkinnedMeshCountOk() (*int32, bool)`

GetSkinnedMeshCountOk returns a tuple with the SkinnedMeshCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkinnedMeshCount

`func (o *FileAnalysisAvatarStats) SetSkinnedMeshCount(v int32)`

SetSkinnedMeshCount sets SkinnedMeshCount field to given value.


### GetSkinnedMeshIndices

`func (o *FileAnalysisAvatarStats) GetSkinnedMeshIndices() int32`

GetSkinnedMeshIndices returns the SkinnedMeshIndices field if non-nil, zero value otherwise.

### GetSkinnedMeshIndicesOk

`func (o *FileAnalysisAvatarStats) GetSkinnedMeshIndicesOk() (*int32, bool)`

GetSkinnedMeshIndicesOk returns a tuple with the SkinnedMeshIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkinnedMeshIndices

`func (o *FileAnalysisAvatarStats) SetSkinnedMeshIndices(v int32)`

SetSkinnedMeshIndices sets SkinnedMeshIndices field to given value.


### GetSkinnedMeshPolygons

`func (o *FileAnalysisAvatarStats) GetSkinnedMeshPolygons() int32`

GetSkinnedMeshPolygons returns the SkinnedMeshPolygons field if non-nil, zero value otherwise.

### GetSkinnedMeshPolygonsOk

`func (o *FileAnalysisAvatarStats) GetSkinnedMeshPolygonsOk() (*int32, bool)`

GetSkinnedMeshPolygonsOk returns a tuple with the SkinnedMeshPolygons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkinnedMeshPolygons

`func (o *FileAnalysisAvatarStats) SetSkinnedMeshPolygons(v int32)`

SetSkinnedMeshPolygons sets SkinnedMeshPolygons field to given value.


### GetSkinnedMeshVertices

`func (o *FileAnalysisAvatarStats) GetSkinnedMeshVertices() int32`

GetSkinnedMeshVertices returns the SkinnedMeshVertices field if non-nil, zero value otherwise.

### GetSkinnedMeshVerticesOk

`func (o *FileAnalysisAvatarStats) GetSkinnedMeshVerticesOk() (*int32, bool)`

GetSkinnedMeshVerticesOk returns a tuple with the SkinnedMeshVertices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkinnedMeshVertices

`func (o *FileAnalysisAvatarStats) SetSkinnedMeshVertices(v int32)`

SetSkinnedMeshVertices sets SkinnedMeshVertices field to given value.


### GetTotalClothVertices

`func (o *FileAnalysisAvatarStats) GetTotalClothVertices() int32`

GetTotalClothVertices returns the TotalClothVertices field if non-nil, zero value otherwise.

### GetTotalClothVerticesOk

`func (o *FileAnalysisAvatarStats) GetTotalClothVerticesOk() (*int32, bool)`

GetTotalClothVerticesOk returns a tuple with the TotalClothVertices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClothVertices

`func (o *FileAnalysisAvatarStats) SetTotalClothVertices(v int32)`

SetTotalClothVertices sets TotalClothVertices field to given value.


### GetTotalIndices

`func (o *FileAnalysisAvatarStats) GetTotalIndices() int32`

GetTotalIndices returns the TotalIndices field if non-nil, zero value otherwise.

### GetTotalIndicesOk

`func (o *FileAnalysisAvatarStats) GetTotalIndicesOk() (*int32, bool)`

GetTotalIndicesOk returns a tuple with the TotalIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIndices

`func (o *FileAnalysisAvatarStats) SetTotalIndices(v int32)`

SetTotalIndices sets TotalIndices field to given value.


### GetTotalMaxParticles

`func (o *FileAnalysisAvatarStats) GetTotalMaxParticles() int32`

GetTotalMaxParticles returns the TotalMaxParticles field if non-nil, zero value otherwise.

### GetTotalMaxParticlesOk

`func (o *FileAnalysisAvatarStats) GetTotalMaxParticlesOk() (*int32, bool)`

GetTotalMaxParticlesOk returns a tuple with the TotalMaxParticles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMaxParticles

`func (o *FileAnalysisAvatarStats) SetTotalMaxParticles(v int32)`

SetTotalMaxParticles sets TotalMaxParticles field to given value.


### GetTotalPolygons

`func (o *FileAnalysisAvatarStats) GetTotalPolygons() int32`

GetTotalPolygons returns the TotalPolygons field if non-nil, zero value otherwise.

### GetTotalPolygonsOk

`func (o *FileAnalysisAvatarStats) GetTotalPolygonsOk() (*int32, bool)`

GetTotalPolygonsOk returns a tuple with the TotalPolygons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPolygons

`func (o *FileAnalysisAvatarStats) SetTotalPolygons(v int32)`

SetTotalPolygons sets TotalPolygons field to given value.


### GetTotalTextureUsage

`func (o *FileAnalysisAvatarStats) GetTotalTextureUsage() int32`

GetTotalTextureUsage returns the TotalTextureUsage field if non-nil, zero value otherwise.

### GetTotalTextureUsageOk

`func (o *FileAnalysisAvatarStats) GetTotalTextureUsageOk() (*int32, bool)`

GetTotalTextureUsageOk returns a tuple with the TotalTextureUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTextureUsage

`func (o *FileAnalysisAvatarStats) SetTotalTextureUsage(v int32)`

SetTotalTextureUsage sets TotalTextureUsage field to given value.


### GetTotalVertices

`func (o *FileAnalysisAvatarStats) GetTotalVertices() int32`

GetTotalVertices returns the TotalVertices field if non-nil, zero value otherwise.

### GetTotalVerticesOk

`func (o *FileAnalysisAvatarStats) GetTotalVerticesOk() (*int32, bool)`

GetTotalVerticesOk returns a tuple with the TotalVertices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVertices

`func (o *FileAnalysisAvatarStats) SetTotalVertices(v int32)`

SetTotalVertices sets TotalVertices field to given value.


### GetTrailRendererCount

`func (o *FileAnalysisAvatarStats) GetTrailRendererCount() int32`

GetTrailRendererCount returns the TrailRendererCount field if non-nil, zero value otherwise.

### GetTrailRendererCountOk

`func (o *FileAnalysisAvatarStats) GetTrailRendererCountOk() (*int32, bool)`

GetTrailRendererCountOk returns a tuple with the TrailRendererCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailRendererCount

`func (o *FileAnalysisAvatarStats) SetTrailRendererCount(v int32)`

SetTrailRendererCount sets TrailRendererCount field to given value.


### GetWriteDefaultsUsed

`func (o *FileAnalysisAvatarStats) GetWriteDefaultsUsed() bool`

GetWriteDefaultsUsed returns the WriteDefaultsUsed field if non-nil, zero value otherwise.

### GetWriteDefaultsUsedOk

`func (o *FileAnalysisAvatarStats) GetWriteDefaultsUsedOk() (*bool, bool)`

GetWriteDefaultsUsedOk returns a tuple with the WriteDefaultsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteDefaultsUsed

`func (o *FileAnalysisAvatarStats) SetWriteDefaultsUsed(v bool)`

SetWriteDefaultsUsed sets WriteDefaultsUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


