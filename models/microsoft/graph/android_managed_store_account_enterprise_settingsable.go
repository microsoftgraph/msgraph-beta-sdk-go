package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidManagedStoreAccountEnterpriseSettingsable 
type AndroidManagedStoreAccountEnterpriseSettingsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAndroidDeviceOwnerFullyManagedEnrollmentEnabled()(*bool)
    GetBindStatus()(*AndroidManagedStoreAccountBindStatus)
    GetCompanyCodes()([]AndroidEnrollmentCompanyCodeable)
    GetDeviceOwnerManagementEnabled()(*bool)
    GetEnrollmentTarget()(*AndroidManagedStoreAccountEnrollmentTarget)
    GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastAppSyncStatus()(*AndroidManagedStoreAccountAppSyncStatus)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedGooglePlayInitialScopeTagIds()([]string)
    GetOwnerOrganizationName()(*string)
    GetOwnerUserPrincipalName()(*string)
    GetTargetGroupIds()([]string)
    SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled(value *bool)()
    SetBindStatus(value *AndroidManagedStoreAccountBindStatus)()
    SetCompanyCodes(value []AndroidEnrollmentCompanyCodeable)()
    SetDeviceOwnerManagementEnabled(value *bool)()
    SetEnrollmentTarget(value *AndroidManagedStoreAccountEnrollmentTarget)()
    SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastAppSyncStatus(value *AndroidManagedStoreAccountAppSyncStatus)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedGooglePlayInitialScopeTagIds(value []string)()
    SetOwnerOrganizationName(value *string)()
    SetOwnerUserPrincipalName(value *string)()
    SetTargetGroupIds(value []string)()
}
