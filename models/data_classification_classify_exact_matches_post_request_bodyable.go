package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataClassificationClassifyExactMatchesPostRequestBodyable 
type DataClassificationClassifyExactMatchesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentClassifications()([]ContentClassificationable)
    GetSensitiveTypeIds()([]string)
    GetText()(*string)
    GetTimeoutInMs()(*string)
    SetContentClassifications(value []ContentClassificationable)()
    SetSensitiveTypeIds(value []string)()
    SetText(value *string)()
    SetTimeoutInMs(value *string)()
}
