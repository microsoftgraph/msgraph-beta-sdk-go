package models

import (
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DetectedSensitiveContentBaseable 
type DetectedSensitiveContentBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfidence()(*int32)
    GetDisplayName()(*string)
    GetId()(*UUID)
    GetOdataType()(*string)
    GetRecommendedConfidence()(*int32)
    GetUniqueCount()(*int32)
    SetConfidence(value *int32)()
    SetDisplayName(value *string)()
    SetId(value *UUID)()
    SetOdataType(value *string)()
    SetRecommendedConfidence(value *int32)()
    SetUniqueCount(value *int32)()
}
