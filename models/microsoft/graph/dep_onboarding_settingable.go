package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DepOnboardingSettingable 
type DepOnboardingSettingable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppleIdentifier()(*string)
    GetDataSharingConsentGranted()(*bool)
    GetDefaultIosEnrollmentProfile()(DepIOSEnrollmentProfileable)
    GetDefaultMacOsEnrollmentProfile()(DepMacOSEnrollmentProfileable)
    GetEnrollmentProfiles()([]EnrollmentProfileable)
    GetImportedAppleDeviceIdentities()([]ImportedAppleDeviceIdentityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSyncErrorCode()(*int32)
    GetLastSyncTriggeredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetShareTokenWithSchoolDataSyncService()(*bool)
    GetSyncedDeviceCount()(*int32)
    GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTokenName()(*string)
    GetTokenType()(*DepTokenType)
    SetAppleIdentifier(value *string)()
    SetDataSharingConsentGranted(value *bool)()
    SetDefaultIosEnrollmentProfile(value DepIOSEnrollmentProfileable)()
    SetDefaultMacOsEnrollmentProfile(value DepMacOSEnrollmentProfileable)()
    SetEnrollmentProfiles(value []EnrollmentProfileable)()
    SetImportedAppleDeviceIdentities(value []ImportedAppleDeviceIdentityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSyncErrorCode(value *int32)()
    SetLastSyncTriggeredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetShareTokenWithSchoolDataSyncService(value *bool)()
    SetSyncedDeviceCount(value *int32)()
    SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTokenName(value *string)()
    SetTokenType(value *DepTokenType)()
}
