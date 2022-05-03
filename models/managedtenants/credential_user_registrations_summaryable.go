package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CredentialUserRegistrationsSummaryable 
type CredentialUserRegistrationsSummaryable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMfaAndSsprCapableUserCount()(*int32)
    GetMfaConditionalAccessPolicyState()(*string)
    GetMfaExcludedUserCount()(*int32)
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
    SetMfaExcludedUserCount(value *int32)()
    SetMfaRegisteredUserCount(value *int32)()
    SetSecurityDefaultsEnabled(value *bool)()
    SetSsprEnabledUserCount(value *int32)()
    SetSsprRegisteredUserCount(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetTotalUserCount(value *int32)()
}
