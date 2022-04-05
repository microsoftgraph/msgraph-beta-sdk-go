package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VppTokenLicenseSummaryable 
type VppTokenLicenseSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppleId()(*string)
    GetAvailableLicenseCount()(*int32)
    GetOrganizationName()(*string)
    GetUsedLicenseCount()(*int32)
    GetVppTokenId()(*string)
    SetAppleId(value *string)()
    SetAvailableLicenseCount(value *int32)()
    SetOrganizationName(value *string)()
    SetUsedLicenseCount(value *int32)()
    SetVppTokenId(value *string)()
}
