package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateAllowedCombinationsResultable 
type UpdateAllowedCombinationsResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalInformation()(*string)
    GetConditionalAccessReferences()([]string)
    GetCurrentCombinations()([]string)
    GetOdataType()(*string)
    GetPreviousCombinations()([]string)
    SetAdditionalInformation(value *string)()
    SetConditionalAccessReferences(value []string)()
    SetCurrentCombinations(value []string)()
    SetOdataType(value *string)()
    SetPreviousCombinations(value []string)()
}
