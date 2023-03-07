package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PortalNotification 
type PortalNotification struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPortalNotification instantiates a new portalNotification and sets the default values.
func NewPortalNotification()(*PortalNotification) {
    m := &PortalNotification{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePortalNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePortalNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPortalNotification(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PortalNotification) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAlertImpact gets the alertImpact property value. The associated alert impact.
func (m *PortalNotification) GetAlertImpact()(AlertImpactable) {
    val, err := m.GetBackingStore().Get("alertImpact")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AlertImpactable)
    }
    return nil
}
// GetAlertRecordId gets the alertRecordId property value. The associated alert record ID.
func (m *PortalNotification) GetAlertRecordId()(*string) {
    val, err := m.GetBackingStore().Get("alertRecordId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAlertRuleId gets the alertRuleId property value. The associated alert rule ID.
func (m *PortalNotification) GetAlertRuleId()(*string) {
    val, err := m.GetBackingStore().Get("alertRuleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAlertRuleName gets the alertRuleName property value. The associated alert rule name.
func (m *PortalNotification) GetAlertRuleName()(*string) {
    val, err := m.GetBackingStore().Get("alertRuleName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAlertRuleTemplate gets the alertRuleTemplate property value. The associated alert rule template. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario, cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue.
func (m *PortalNotification) GetAlertRuleTemplate()(*AlertRuleTemplate) {
    val, err := m.GetBackingStore().Get("alertRuleTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AlertRuleTemplate)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *PortalNotification) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
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
// GetId gets the id property value. The unique identifier for the portal notification.
func (m *PortalNotification) GetId()(*string) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsPortalNotificationSent gets the isPortalNotificationSent property value. true if the portal notification has already been sent to the user; false otherwise.
func (m *PortalNotification) GetIsPortalNotificationSent()(*bool) {
    val, err := m.GetBackingStore().Get("isPortalNotificationSent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PortalNotification) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSeverity gets the severity property value. The associated alert rule severity. The possible values are: unknown, informational, warning, critical, unknownFutureValue.
func (m *PortalNotification) GetSeverity()(*RuleSeverityType) {
    val, err := m.GetBackingStore().Get("severity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RuleSeverityType)
    }
    return nil
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
func (m *PortalNotification) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertImpact sets the alertImpact property value. The associated alert impact.
func (m *PortalNotification) SetAlertImpact(value AlertImpactable)() {
    err := m.GetBackingStore().Set("alertImpact", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertRecordId sets the alertRecordId property value. The associated alert record ID.
func (m *PortalNotification) SetAlertRecordId(value *string)() {
    err := m.GetBackingStore().Set("alertRecordId", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertRuleId sets the alertRuleId property value. The associated alert rule ID.
func (m *PortalNotification) SetAlertRuleId(value *string)() {
    err := m.GetBackingStore().Set("alertRuleId", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertRuleName sets the alertRuleName property value. The associated alert rule name.
func (m *PortalNotification) SetAlertRuleName(value *string)() {
    err := m.GetBackingStore().Set("alertRuleName", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertRuleTemplate sets the alertRuleTemplate property value. The associated alert rule template. The possible values are: cloudPcProvisionScenario, cloudPcImageUploadScenario, cloudPcOnPremiseNetworkConnectionCheckScenario, unknownFutureValue.
func (m *PortalNotification) SetAlertRuleTemplate(value *AlertRuleTemplate)() {
    err := m.GetBackingStore().Set("alertRuleTemplate", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *PortalNotification) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetId sets the id property value. The unique identifier for the portal notification.
func (m *PortalNotification) SetId(value *string)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPortalNotificationSent sets the isPortalNotificationSent property value. true if the portal notification has already been sent to the user; false otherwise.
func (m *PortalNotification) SetIsPortalNotificationSent(value *bool)() {
    err := m.GetBackingStore().Set("isPortalNotificationSent", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PortalNotification) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSeverity sets the severity property value. The associated alert rule severity. The possible values are: unknown, informational, warning, critical, unknownFutureValue.
func (m *PortalNotification) SetSeverity(value *RuleSeverityType)() {
    err := m.GetBackingStore().Set("severity", value)
    if err != nil {
        panic(err)
    }
}
// PortalNotificationable 
type PortalNotificationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertImpact()(AlertImpactable)
    GetAlertRecordId()(*string)
    GetAlertRuleId()(*string)
    GetAlertRuleName()(*string)
    GetAlertRuleTemplate()(*AlertRuleTemplate)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetId()(*string)
    GetIsPortalNotificationSent()(*bool)
    GetOdataType()(*string)
    GetSeverity()(*RuleSeverityType)
    SetAlertImpact(value AlertImpactable)()
    SetAlertRecordId(value *string)()
    SetAlertRuleId(value *string)()
    SetAlertRuleName(value *string)()
    SetAlertRuleTemplate(value *AlertRuleTemplate)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetId(value *string)()
    SetIsPortalNotificationSent(value *bool)()
    SetOdataType(value *string)()
    SetSeverity(value *RuleSeverityType)()
}
