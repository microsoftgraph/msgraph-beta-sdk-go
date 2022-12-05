package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AlertRule provides operations to manage the deviceManagement singleton.
type AlertRule struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The rule template of the alert event. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario, cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue.
    alertRuleTemplate *AlertRuleTemplate
    // The rule description.
    description *string
    // The display name of the rule.
    displayName *string
    // The status of the rule that indicates whether the rule is enabled or disabled. If true, the rule is enabled; otherwise, the rule is disabled.
    enabled *bool
    // Indicates whether the rule is a system rule. If true, the rule is a system rule; otherwise, the rule is a custom defined rule and can be edited. System rules are built-in and only
    isSystemRule *bool
    // The notificationChannels property
    notificationChannels []NotificationChannelable
    // The severity property
    severity *RuleSeverityType
    // The threshold property
    threshold RuleThresholdable
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
    return m.alertRuleTemplate
}
// GetDescription gets the description property value. The rule description.
func (m *AlertRule) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the rule.
func (m *AlertRule) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnabled gets the enabled property value. The status of the rule that indicates whether the rule is enabled or disabled. If true, the rule is enabled; otherwise, the rule is disabled.
func (m *AlertRule) GetEnabled()(*bool) {
    return m.enabled
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
                res[i] = v.(NotificationChannelable)
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
// GetIsSystemRule gets the isSystemRule property value. Indicates whether the rule is a system rule. If true, the rule is a system rule; otherwise, the rule is a custom defined rule and can be edited. System rules are built-in and only
func (m *AlertRule) GetIsSystemRule()(*bool) {
    return m.isSystemRule
}
// GetNotificationChannels gets the notificationChannels property value. The notificationChannels property
func (m *AlertRule) GetNotificationChannels()([]NotificationChannelable) {
    return m.notificationChannels
}
// GetSeverity gets the severity property value. The severity property
func (m *AlertRule) GetSeverity()(*RuleSeverityType) {
    return m.severity
}
// GetThreshold gets the threshold property value. The threshold property
func (m *AlertRule) GetThreshold()(RuleThresholdable) {
    return m.threshold
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    m.alertRuleTemplate = value
}
// SetDescription sets the description property value. The rule description.
func (m *AlertRule) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the rule.
func (m *AlertRule) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnabled sets the enabled property value. The status of the rule that indicates whether the rule is enabled or disabled. If true, the rule is enabled; otherwise, the rule is disabled.
func (m *AlertRule) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetIsSystemRule sets the isSystemRule property value. Indicates whether the rule is a system rule. If true, the rule is a system rule; otherwise, the rule is a custom defined rule and can be edited. System rules are built-in and only
func (m *AlertRule) SetIsSystemRule(value *bool)() {
    m.isSystemRule = value
}
// SetNotificationChannels sets the notificationChannels property value. The notificationChannels property
func (m *AlertRule) SetNotificationChannels(value []NotificationChannelable)() {
    m.notificationChannels = value
}
// SetSeverity sets the severity property value. The severity property
func (m *AlertRule) SetSeverity(value *RuleSeverityType)() {
    m.severity = value
}
// SetThreshold sets the threshold property value. The threshold property
func (m *AlertRule) SetThreshold(value RuleThresholdable)() {
    m.threshold = value
}
