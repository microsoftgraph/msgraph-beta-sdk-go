package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DetectedSensitiveContentBaseable 
type DetectedSensitiveContentBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfidence()(*int32)
    GetDisplayName()(*string)
    GetId()(*string)
    GetRecommendedConfidence()(*int32)
    GetType()(*string)
    GetUniqueCount()(*int32)
    SetConfidence(value *int32)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetRecommendedConfidence(value *int32)()
    SetType(value *string)()
    SetUniqueCount(value *int32)()
}
