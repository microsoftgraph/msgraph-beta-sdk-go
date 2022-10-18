package devicemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PortalNotification 
type PortalNotification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The associated alert impact.
    alertImpact AlertImpactable
    // The associated alert record ID.
    alertRecordId *string
    // The associated alert rule ID.
    alertRuleId *string
    // The associated alert rule name.
    alertRuleName *string
    // The associated alert rule template. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario, cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue.
    alertRuleTemplate *AlertRuleTemplate
    // The unique identifier for the portal notification.
    id *string
    // If true, the portal notification has already been sent for the user; otherwise, the portal notification hasn't been sent yet.
    isPortalNotificationSent *bool
    // The OdataType property
    odataType *string
    // The associated alert rule severity. The possible values are: unknown, informational, warning, critical, unknownFutureValue.
    severity *RuleSeverityType
}
// NewPortalNotification instantiates a new portalNotification and sets the default values.
func NewPortalNotification()(*PortalNotification) {
    m := &PortalNotification{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.deviceManagement.portalNotification";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePortalNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePortalNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPortalNotification(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PortalNotification) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAlertImpact gets the alertImpact property value. The associated alert impact.
func (m *PortalNotification) GetAlertImpact()(AlertImpactable) {
    return m.alertImpact
}
// GetAlertRecordId gets the alertRecordId property value. The associated alert record ID.
func (m *PortalNotification) GetAlertRecordId()(*string) {
    return m.alertRecordId
}
// GetAlertRuleId gets the alertRuleId property value. The associated alert rule ID.
func (m *PortalNotification) GetAlertRuleId()(*string) {
    return m.alertRuleId
}
// GetAlertRuleName gets the alertRuleName property value. The associated alert rule name.
func (m *PortalNotification) GetAlertRuleName()(*string) {
    return m.alertRuleName
}
// GetAlertRuleTemplate gets the alertRuleTemplate property value. The associated alert rule template. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario, cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue.
func (m *PortalNotification) GetAlertRuleTemplate()(*AlertRuleTemplate) {
    return m.alertRuleTemplate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PortalNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alertImpact"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAlertImpactFromDiscriminatorValue , m.SetAlertImpact)
    res["alertRecordId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAlertRecordId)
    res["alertRuleId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAlertRuleId)
    res["alertRuleName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAlertRuleName)
    res["alertRuleTemplate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAlertRuleTemplate , m.SetAlertRuleTemplate)
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    res["isPortalNotificationSent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsPortalNotificationSent)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["severity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRuleSeverityType , m.SetSeverity)
    return res
}
// GetId gets the id property value. The unique identifier for the portal notification.
func (m *PortalNotification) GetId()(*string) {
    return m.id
}
// GetIsPortalNotificationSent gets the isPortalNotificationSent property value. If true, the portal notification has already been sent for the user; otherwise, the portal notification hasn't been sent yet.
func (m *PortalNotification) GetIsPortalNotificationSent()(*bool) {
    return m.isPortalNotificationSent
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PortalNotification) GetOdataType()(*string) {
    return m.odataType
}
// GetSeverity gets the severity property value. The associated alert rule severity. The possible values are: unknown, informational, warning, critical, unknownFutureValue.
func (m *PortalNotification) GetSeverity()(*RuleSeverityType) {
    return m.severity
}
// Serialize serializes information the current object
func (m *PortalNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("alertImpact", m.GetAlertImpact())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("alertRecordId", m.GetAlertRecordId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("alertRuleId", m.GetAlertRuleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("alertRuleName", m.GetAlertRuleName())
        if err != nil {
            return err
        }
    }
    if m.GetAlertRuleTemplate() != nil {
        cast := (*m.GetAlertRuleTemplate()).String()
        err := writer.WriteStringValue("alertRuleTemplate", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPortalNotificationSent", m.GetIsPortalNotificationSent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err := writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PortalNotification) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAlertImpact sets the alertImpact property value. The associated alert impact.
func (m *PortalNotification) SetAlertImpact(value AlertImpactable)() {
    m.alertImpact = value
}
// SetAlertRecordId sets the alertRecordId property value. The associated alert record ID.
func (m *PortalNotification) SetAlertRecordId(value *string)() {
    m.alertRecordId = value
}
// SetAlertRuleId sets the alertRuleId property value. The associated alert rule ID.
func (m *PortalNotification) SetAlertRuleId(value *string)() {
    m.alertRuleId = value
}
// SetAlertRuleName sets the alertRuleName property value. The associated alert rule name.
func (m *PortalNotification) SetAlertRuleName(value *string)() {
    m.alertRuleName = value
}
// SetAlertRuleTemplate sets the alertRuleTemplate property value. The associated alert rule template. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario, cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue.
func (m *PortalNotification) SetAlertRuleTemplate(value *AlertRuleTemplate)() {
    m.alertRuleTemplate = value
}
// SetId sets the id property value. The unique identifier for the portal notification.
func (m *PortalNotification) SetId(value *string)() {
    m.id = value
}
// SetIsPortalNotificationSent sets the isPortalNotificationSent property value. If true, the portal notification has already been sent for the user; otherwise, the portal notification hasn't been sent yet.
func (m *PortalNotification) SetIsPortalNotificationSent(value *bool)() {
    m.isPortalNotificationSent = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PortalNotification) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSeverity sets the severity property value. The associated alert rule severity. The possible values are: unknown, informational, warning, critical, unknownFutureValue.
func (m *PortalNotification) SetSeverity(value *RuleSeverityType)() {
    m.severity = value
}
