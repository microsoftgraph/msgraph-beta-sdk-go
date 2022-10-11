package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedTenantAlertable 
type ManagedTenantAlertable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertData()(AlertDataable)
    GetAlertDataReferenceStrings()([]AlertDataReferenceStringable)
    GetAlertLogs()([]ManagedTenantAlertLogable)
    GetAlertRule()(ManagedTenantAlertRuleable)
    GetAlertRuleDisplayName()(*string)
    GetApiNotifications()([]ManagedTenantApiNotificationable)
    GetAssignedToUserId()(*string)
    GetCorrelationCount()(*int32)
    GetCorrelationId()(*string)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmailNotifications()([]ManagedTenantEmailNotificationable)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMessage()(*string)
    GetSeverity()(*AlertSeverity)
    GetStatus()(*AlertStatus)
    GetTenantId()(*string)
    GetTitle()(*string)
    SetAlertData(value AlertDataable)()
    SetAlertDataReferenceStrings(value []AlertDataReferenceStringable)()
    SetAlertLogs(value []ManagedTenantAlertLogable)()
    SetAlertRule(value ManagedTenantAlertRuleable)()
    SetAlertRuleDisplayName(value *string)()
    SetApiNotifications(value []ManagedTenantApiNotificationable)()
    SetAssignedToUserId(value *string)()
    SetCorrelationCount(value *int32)()
    SetCorrelationId(value *string)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmailNotifications(value []ManagedTenantEmailNotificationable)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMessage(value *string)()
    SetSeverity(value *AlertSeverity)()
    SetStatus(value *AlertStatus)()
    SetTenantId(value *string)()
    SetTitle(value *string)()
}
