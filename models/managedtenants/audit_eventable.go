package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AuditEventable 
type AuditEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivity()(*string)
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetActivityId()(*string)
    GetCategory()(*string)
    GetHttpVerb()(*string)
    GetInitiatedByAppId()(*string)
    GetInitiatedByUpn()(*string)
    GetInitiatedByUserId()(*string)
    GetIpAddress()(*string)
    GetRequestBody()(*string)
    GetRequestUrl()(*string)
    GetTenantIds()(*string)
    GetTenantNames()(*string)
    SetActivity(value *string)()
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetActivityId(value *string)()
    SetCategory(value *string)()
    SetHttpVerb(value *string)()
    SetInitiatedByAppId(value *string)()
    SetInitiatedByUpn(value *string)()
    SetInitiatedByUserId(value *string)()
    SetIpAddress(value *string)()
    SetRequestBody(value *string)()
    SetRequestUrl(value *string)()
    SetTenantIds(value *string)()
    SetTenantNames(value *string)()
}
