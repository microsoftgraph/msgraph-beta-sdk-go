package classifyexactmatches

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ClassifyExactMatchesRequestBodyable 
type ClassifyExactMatchesRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentClassifications()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentClassificationable)
    GetSensitiveTypeIds()([]string)
    GetText()(*string)
    GetTimeoutInMs()(*string)
    SetContentClassifications(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentClassificationable)()
    SetSensitiveTypeIds(value []string)()
    SetText(value *string)()
    SetTimeoutInMs(value *string)()
}
