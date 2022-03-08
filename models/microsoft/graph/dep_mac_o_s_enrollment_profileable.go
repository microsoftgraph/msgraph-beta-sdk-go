package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DepMacOSEnrollmentProfileable 
type DepMacOSEnrollmentProfileable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    DepEnrollmentBaseProfileable
    GetAccessibilityScreenDisabled()(*bool)
    GetChooseYourLockScreenDisabled()(*bool)
    GetFileVaultDisabled()(*bool)
    GetICloudDiagnosticsDisabled()(*bool)
    GetICloudStorageDisabled()(*bool)
    GetPassCodeDisabled()(*bool)
    GetRegistrationDisabled()(*bool)
    GetZoomDisabled()(*bool)
    SetAccessibilityScreenDisabled(value *bool)()
    SetChooseYourLockScreenDisabled(value *bool)()
    SetFileVaultDisabled(value *bool)()
    SetICloudDiagnosticsDisabled(value *bool)()
    SetICloudStorageDisabled(value *bool)()
    SetPassCodeDisabled(value *bool)()
    SetRegistrationDisabled(value *bool)()
    SetZoomDisabled(value *bool)()
}
