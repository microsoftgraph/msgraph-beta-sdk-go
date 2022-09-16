package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PortalNotification 
type PortalNotification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The alertImpact property
    alertImpact AlertImpactable
    // The alertRecordId property
    alertRecordId *string
    // The alertRuleId property
    alertRuleId *string
    // The alertRuleName property
    alertRuleName *string
    // The alertRuleTemplate property
    alertRuleTemplate *AlertRuleTemplate
    // The id property
    id *string
    // The isPortalNotificationSent property
    isPortalNotificationSent *bool
    // The OdataType property
    odataType *string
    // The severity property
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
// GetAlertImpact gets the alertImpact property value. The alertImpact property
func (m *PortalNotification) GetAlertImpact()(AlertImpactable) {
    return m.alertImpact
}
// GetAlertRecordId gets the alertRecordId property value. The alertRecordId property
func (m *PortalNotification) GetAlertRecordId()(*string) {
    return m.alertRecordId
}
// GetAlertRuleId gets the alertRuleId property value. The alertRuleId property
func (m *PortalNotification) GetAlertRuleId()(*string) {
    return m.alertRuleId
}
// GetAlertRuleName gets the alertRuleName property value. The alertRuleName property
func (m *PortalNotification) GetAlertRuleName()(*string) {
    return m.alertRuleName
}
// GetAlertRuleTemplate gets the alertRuleTemplate property value. The alertRuleTemplate property
func (m *PortalNotification) GetAlertRuleTemplate()(*AlertRuleTemplate) {
    return m.alertRuleTemplate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PortalNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alertImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAlertImpactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertImpact(val.(AlertImpactable))
        }
        return nil
    }
    res["alertRecordId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertRecordId(val)
        }
        return nil
    }
    res["alertRuleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertRuleId(val)
        }
        return nil
    }
    res["alertRuleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertRuleName(val)
        }
        return nil
    }
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isPortalNotificationSent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPortalNotificationSent(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
    return res
}
// GetId gets the id property value. The id property
func (m *PortalNotification) GetId()(*string) {
    return m.id
}
// GetIsPortalNotificationSent gets the isPortalNotificationSent property value. The isPortalNotificationSent property
func (m *PortalNotification) GetIsPortalNotificationSent()(*bool) {
    return m.isPortalNotificationSent
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PortalNotification) GetOdataType()(*string) {
    return m.odataType
}
// GetSeverity gets the severity property value. The severity property
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
// SetAlertImpact sets the alertImpact property value. The alertImpact property
func (m *PortalNotification) SetAlertImpact(value AlertImpactable)() {
    m.alertImpact = value
}
// SetAlertRecordId sets the alertRecordId property value. The alertRecordId property
func (m *PortalNotification) SetAlertRecordId(value *string)() {
    m.alertRecordId = value
}
// SetAlertRuleId sets the alertRuleId property value. The alertRuleId property
func (m *PortalNotification) SetAlertRuleId(value *string)() {
    m.alertRuleId = value
}
// SetAlertRuleName sets the alertRuleName property value. The alertRuleName property
func (m *PortalNotification) SetAlertRuleName(value *string)() {
    m.alertRuleName = value
}
// SetAlertRuleTemplate sets the alertRuleTemplate property value. The alertRuleTemplate property
func (m *PortalNotification) SetAlertRuleTemplate(value *AlertRuleTemplate)() {
    m.alertRuleTemplate = value
}
// SetId sets the id property value. The id property
func (m *PortalNotification) SetId(value *string)() {
    m.id = value
}
// SetIsPortalNotificationSent sets the isPortalNotificationSent property value. The isPortalNotificationSent property
func (m *PortalNotification) SetIsPortalNotificationSent(value *bool)() {
    m.isPortalNotificationSent = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PortalNotification) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSeverity sets the severity property value. The severity property
func (m *PortalNotification) SetSeverity(value *RuleSeverityType)() {
    m.severity = value
}
