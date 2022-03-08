package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CredentialUserRegistrationsSummaryable 
type CredentialUserRegistrationsSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMfaAndSsprCapableUserCount()(*int32)
    GetMfaConditionalAccessPolicyState()(*string)
    GetMfaRegisteredUserCount()(*int32)
    GetSecurityDefaultsEnabled()(*bool)
    GetSsprEnabledUserCount()(*int32)
    GetSsprRegisteredUserCount()(*int32)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetTotalUserCount()(*int32)
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMfaAndSsprCapableUserCount(value *int32)()
    SetMfaConditionalAccessPolicyState(value *string)()
    SetMfaRegisteredUserCount(value *int32)()
    SetSecurityDefaultsEnabled(value *bool)()
    SetSsprEnabledUserCount(value *int32)()
    SetSsprRegisteredUserCount(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetTotalUserCount(value *int32)()
}
