package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AlertRule 
type AlertRule struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The OdataType property
    OdataType *string
}
// NewAlertRule instantiates a new alertRule and sets the default values.
func NewAlertRule()(*AlertRule) {
    m := &AlertRule{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateAlertRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertRule(), nil
}
// GetAlertRuleTemplate gets the alertRuleTemplate property value. The rule template of the alert event. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario, cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue.
func (m *AlertRule) GetAlertRuleTemplate()(*AlertRuleTemplate) {
    val, err := m.GetBackingStore().Get("alertRuleTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AlertRuleTemplate)
    }
    return nil
}
// GetDescription gets the description property value. The rule description.
func (m *AlertRule) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the rule.
func (m *AlertRule) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnabled gets the enabled property value. The status of the rule that indicates whether the rule is enabled or disabled. If true, the rule is enabled; otherwise, the rule is disabled.
func (m *AlertRule) GetEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("enabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertRuleTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertRuleTemplate)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertRuleTemplate(val.(*AlertRuleTemplate))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["isSystemRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSystemRule(val)
        }
        return nil
    }
    res["notificationChannels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNotificationChannelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NotificationChannelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NotificationChannelable)
                }
            }
            m.SetNotificationChannels(res)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleSeverityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*RuleSeverityType))
        }
        return nil
    }
    res["threshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRuleThresholdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreshold(val.(RuleThresholdable))
        }
        return nil
    }
    return res
}
// GetIsSystemRule gets the isSystemRule property value. Indicates whether the rule is a system rule. If true, the rule is a system rule; otherwise, the rule is a custom defined rule and can be edited. System rules are built-in and only a few properties can be edited.
func (m *AlertRule) GetIsSystemRule()(*bool) {
    val, err := m.GetBackingStore().Get("isSystemRule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetNotificationChannels gets the notificationChannels property value. The notification channels of the rule selected by the user.
func (m *AlertRule) GetNotificationChannels()([]NotificationChannelable) {
    val, err := m.GetBackingStore().Get("notificationChannels")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NotificationChannelable)
    }
    return nil
}
// GetSeverity gets the severity property value. The severity of the rule. The possible values are: unknown, informational, warning, critical, unknownFutureValue.
func (m *AlertRule) GetSeverity()(*RuleSeverityType) {
    val, err := m.GetBackingStore().Get("severity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RuleSeverityType)
    }
    return nil
}
// GetThreshold gets the threshold property value. The conditions to send alerts. For example, send alert when provisioning has failed for greater than or equal to 6 Cloud PCs.
func (m *AlertRule) GetThreshold()(RuleThresholdable) {
    val, err := m.GetBackingStore().Get("threshold")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RuleThresholdable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AlertRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlertRuleTemplate() != nil {
        cast := (*m.GetAlertRuleTemplate()).String()
        err = writer.WriteStringValue("alertRuleTemplate", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSystemRule", m.GetIsSystemRule())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationChannels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotificationChannels()))
        for i, v := range m.GetNotificationChannels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("notificationChannels", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err = writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("threshold", m.GetThreshold())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertRuleTemplate sets the alertRuleTemplate property value. The rule template of the alert event. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario, cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue.
func (m *AlertRule) SetAlertRuleTemplate(value *AlertRuleTemplate)() {
    err := m.GetBackingStore().Set("alertRuleTemplate", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The rule description.
func (m *AlertRule) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the rule.
func (m *AlertRule) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEnabled sets the enabled property value. The status of the rule that indicates whether the rule is enabled or disabled. If true, the rule is enabled; otherwise, the rule is disabled.
func (m *AlertRule) SetEnabled(value *bool)() {
    err := m.GetBackingStore().Set("enabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSystemRule sets the isSystemRule property value. Indicates whether the rule is a system rule. If true, the rule is a system rule; otherwise, the rule is a custom defined rule and can be edited. System rules are built-in and only a few properties can be edited.
func (m *AlertRule) SetIsSystemRule(value *bool)() {
    err := m.GetBackingStore().Set("isSystemRule", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationChannels sets the notificationChannels property value. The notification channels of the rule selected by the user.
func (m *AlertRule) SetNotificationChannels(value []NotificationChannelable)() {
    err := m.GetBackingStore().Set("notificationChannels", value)
    if err != nil {
        panic(err)
    }
}
// SetSeverity sets the severity property value. The severity of the rule. The possible values are: unknown, informational, warning, critical, unknownFutureValue.
func (m *AlertRule) SetSeverity(value *RuleSeverityType)() {
    err := m.GetBackingStore().Set("severity", value)
    if err != nil {
        panic(err)
    }
}
// SetThreshold sets the threshold property value. The conditions to send alerts. For example, send alert when provisioning has failed for greater than or equal to 6 Cloud PCs.
func (m *AlertRule) SetThreshold(value RuleThresholdable)() {
    err := m.GetBackingStore().Set("threshold", value)
    if err != nil {
        panic(err)
    }
}
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
