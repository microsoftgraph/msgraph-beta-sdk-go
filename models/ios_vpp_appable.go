package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppAppable 
type IosVppAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableDeviceType()(IosDeviceTypeable)
    GetAppStoreUrl()(*string)
    GetAssignedLicenses()([]IosVppAppAssignedLicenseable)
    GetBundleId()(*string)
    GetLicensingType()(VppLicensingTypeable)
    GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRevokeLicenseActionResults()([]IosVppAppRevokeLicensesActionResultable)
    GetTotalLicenseCount()(*int32)
    GetUsedLicenseCount()(*int32)
    GetVppTokenAccountType()(*VppTokenAccountType)
    GetVppTokenAppleId()(*string)
    GetVppTokenId()(*string)
    GetVppTokenOrganizationName()(*string)
    SetApplicableDeviceType(value IosDeviceTypeable)()
    SetAppStoreUrl(value *string)()
    SetAssignedLicenses(value []IosVppAppAssignedLicenseable)()
    SetBundleId(value *string)()
    SetLicensingType(value VppLicensingTypeable)()
    SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRevokeLicenseActionResults(value []IosVppAppRevokeLicensesActionResultable)()
    SetTotalLicenseCount(value *int32)()
    SetUsedLicenseCount(value *int32)()
    SetVppTokenAccountType(value *VppTokenAccountType)()
    SetVppTokenAppleId(value *string)()
    SetVppTokenId(value *string)()
    SetVppTokenOrganizationName(value *string)()
}
