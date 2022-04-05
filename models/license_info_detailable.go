package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LicenseInfoDetailable 
type LicenseInfoDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLicenseType()(*AzureADLicenseType)
    GetTotalAssignedCount()(*int32)
    GetTotalLicenseCount()(*int32)
    GetTotalUsageCount()(*int32)
    SetLicenseType(value *AzureADLicenseType)()
    SetTotalAssignedCount(value *int32)()
    SetTotalLicenseCount(value *int32)()
    SetTotalUsageCount(value *int32)()
}
