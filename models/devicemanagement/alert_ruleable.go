package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AlertRuleable 
type AlertRuleable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertRuleTemplate()(*AlertRuleTemplate)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEnabled()(*bool)
    GetIsSystemRule()(*bool)
    GetNotificationChannels()([]NotificationChannelable)
    GetSeverity()(*RuleSeverityType)
    GetThreshold()(RuleThresholdable)
    SetAlertRuleTemplate(value *AlertRuleTemplate)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEnabled(value *bool)()
    SetIsSystemRule(value *bool)()
    SetNotificationChannels(value []NotificationChannelable)()
    SetSeverity(value *RuleSeverityType)()
    SetThreshold(value RuleThresholdable)()
}
