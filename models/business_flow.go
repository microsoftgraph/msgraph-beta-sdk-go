package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessFlow 
type BusinessFlow struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewBusinessFlow instantiates a new businessFlow and sets the default values.
func NewBusinessFlow()(*BusinessFlow) {
    m := &BusinessFlow{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBusinessFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessFlow(), nil
}
// GetCustomData gets the customData property value. The customData property
func (m *BusinessFlow) GetCustomData()(*string) {
    val, err := m.GetBackingStore().Get("customData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeDuplicationId gets the deDuplicationId property value. The deDuplicationId property
func (m *BusinessFlow) GetDeDuplicationId()(*string) {
    val, err := m.GetBackingStore().Get("deDuplicationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The description property
func (m *BusinessFlow) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *BusinessFlow) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomData(val)
        }
        return nil
    }
    res["deDuplicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeDuplicationId(val)
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
    res["policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val.(GovernancePolicyable))
        }
        return nil
    }
    res["policyTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyTemplateId(val)
        }
        return nil
    }
    res["recordVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordVersion(val)
        }
        return nil
    }
    res["schemaId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaId(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBusinessFlowSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(BusinessFlowSettingsable))
        }
        return nil
    }
    return res
}
// GetPolicy gets the policy property value. The policy property
func (m *BusinessFlow) GetPolicy()(GovernancePolicyable) {
    val, err := m.GetBackingStore().Get("policy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GovernancePolicyable)
    }
    return nil
}
// GetPolicyTemplateId gets the policyTemplateId property value. The policyTemplateId property
func (m *BusinessFlow) GetPolicyTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("policyTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecordVersion gets the recordVersion property value. The recordVersion property
func (m *BusinessFlow) GetRecordVersion()(*string) {
    val, err := m.GetBackingStore().Get("recordVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSchemaId gets the schemaId property value. The schemaId property
func (m *BusinessFlow) GetSchemaId()(*string) {
    val, err := m.GetBackingStore().Get("schemaId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettings gets the settings property value. The settings property
func (m *BusinessFlow) GetSettings()(BusinessFlowSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(BusinessFlowSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BusinessFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("customData", m.GetCustomData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deDuplicationId", m.GetDeDuplicationId())
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
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyTemplateId", m.GetPolicyTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recordVersion", m.GetRecordVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schemaId", m.GetSchemaId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomData sets the customData property value. The customData property
func (m *BusinessFlow) SetCustomData(value *string)() {
    err := m.GetBackingStore().Set("customData", value)
    if err != nil {
        panic(err)
    }
}
// SetDeDuplicationId sets the deDuplicationId property value. The deDuplicationId property
func (m *BusinessFlow) SetDeDuplicationId(value *string)() {
    err := m.GetBackingStore().Set("deDuplicationId", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *BusinessFlow) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *BusinessFlow) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicy sets the policy property value. The policy property
func (m *BusinessFlow) SetPolicy(value GovernancePolicyable)() {
    err := m.GetBackingStore().Set("policy", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyTemplateId sets the policyTemplateId property value. The policyTemplateId property
func (m *BusinessFlow) SetPolicyTemplateId(value *string)() {
    err := m.GetBackingStore().Set("policyTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetRecordVersion sets the recordVersion property value. The recordVersion property
func (m *BusinessFlow) SetRecordVersion(value *string)() {
    err := m.GetBackingStore().Set("recordVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetSchemaId sets the schemaId property value. The schemaId property
func (m *BusinessFlow) SetSchemaId(value *string)() {
    err := m.GetBackingStore().Set("schemaId", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. The settings property
func (m *BusinessFlow) SetSettings(value BusinessFlowSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// BusinessFlowable 
type BusinessFlowable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomData()(*string)
    GetDeDuplicationId()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetPolicy()(GovernancePolicyable)
    GetPolicyTemplateId()(*string)
    GetRecordVersion()(*string)
    GetSchemaId()(*string)
    GetSettings()(BusinessFlowSettingsable)
    SetCustomData(value *string)()
    SetDeDuplicationId(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetPolicy(value GovernancePolicyable)()
    SetPolicyTemplateId(value *string)()
    SetRecordVersion(value *string)()
    SetSchemaId(value *string)()
    SetSettings(value BusinessFlowSettingsable)()
}
