package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PortalNotificationable 
type PortalNotificationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertImpact()(AlertImpactable)
    GetAlertRecordId()(*string)
    GetAlertRuleId()(*string)
    GetAlertRuleName()(*string)
    GetAlertRuleTemplate()(*AlertRuleTemplate)
    GetId()(*string)
    GetIsPortalNotificationSent()(*bool)
    GetOdataType()(*string)
    GetSeverity()(*RuleSeverityType)
    SetAlertImpact(value AlertImpactable)()
    SetAlertRecordId(value *string)()
    SetAlertRuleId(value *string)()
    SetAlertRuleName(value *string)()
    SetAlertRuleTemplate(value *AlertRuleTemplate)()
    SetId(value *string)()
    SetIsPortalNotificationSent(value *bool)()
    SetOdataType(value *string)()
    SetSeverity(value *RuleSeverityType)()
}
