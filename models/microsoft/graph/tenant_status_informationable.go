package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// TenantStatusInformationable 
type TenantStatusInformationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDelegatedPrivilegeStatus()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.DelegatedPrivilegeStatus)
    GetLastDelegatedPrivilegeRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOffboardedByUserId()(*string)
    GetOffboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnboardedByUserId()(*string)
    GetOnboardedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnboardingStatus()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingStatus)
    GetTenantOnboardingEligibilityReason()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingEligibilityReason)
    GetWorkloadStatuses()([]WorkloadStatusable)
    SetDelegatedPrivilegeStatus(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.DelegatedPrivilegeStatus)()
    SetLastDelegatedPrivilegeRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOffboardedByUserId(value *string)()
    SetOffboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnboardedByUserId(value *string)()
    SetOnboardedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnboardingStatus(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingStatus)()
    SetTenantOnboardingEligibilityReason(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.TenantOnboardingEligibilityReason)()
    SetWorkloadStatuses(value []WorkloadStatusable)()
}
