package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepMacOSEnrollmentProfileable 
type DepMacOSEnrollmentProfileable interface {
    DepEnrollmentBaseProfileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessibilityScreenDisabled()(*bool)
    GetAutoUnlockWithWatchDisabled()(*bool)
    GetChooseYourLockScreenDisabled()(*bool)
    GetFileVaultDisabled()(*bool)
    GetICloudDiagnosticsDisabled()(*bool)
    GetICloudStorageDisabled()(*bool)
    GetPassCodeDisabled()(*bool)
    GetRegistrationDisabled()(*bool)
    GetZoomDisabled()(*bool)
    SetAccessibilityScreenDisabled(value *bool)()
    SetAutoUnlockWithWatchDisabled(value *bool)()
    SetChooseYourLockScreenDisabled(value *bool)()
    SetFileVaultDisabled(value *bool)()
    SetICloudDiagnosticsDisabled(value *bool)()
    SetICloudStorageDisabled(value *bool)()
    SetPassCodeDisabled(value *bool)()
    SetRegistrationDisabled(value *bool)()
    SetZoomDisabled(value *bool)()
}
