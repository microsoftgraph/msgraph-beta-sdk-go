package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessFlow provides operations to manage the collection of approvalWorkflowProvider entities.
type BusinessFlow struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The customData property
    customData *string
    // The deDuplicationId property
    deDuplicationId *string
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The policy property
    policy GovernancePolicyable
    // The policyTemplateId property
    policyTemplateId *string
    // The recordVersion property
    recordVersion *string
    // The schemaId property
    schemaId *string
    // The settings property
    settings BusinessFlowSettingsable
}
// NewBusinessFlow instantiates a new businessFlow and sets the default values.
func NewBusinessFlow()(*BusinessFlow) {
    m := &BusinessFlow{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBusinessFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessFlow(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BusinessFlow) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCustomData gets the customData property value. The customData property
func (m *BusinessFlow) GetCustomData()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customData
    }
}
// GetDeDuplicationId gets the deDuplicationId property value. The deDuplicationId property
func (m *BusinessFlow) GetDeDuplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deDuplicationId
    }
}
// GetDescription gets the description property value. The description property
func (m *BusinessFlow) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *BusinessFlow) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
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
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
// GetPolicyTemplateId gets the policyTemplateId property value. The policyTemplateId property
func (m *BusinessFlow) GetPolicyTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyTemplateId
    }
}
// GetRecordVersion gets the recordVersion property value. The recordVersion property
func (m *BusinessFlow) GetRecordVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordVersion
    }
}
// GetSchemaId gets the schemaId property value. The schemaId property
func (m *BusinessFlow) GetSchemaId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schemaId
    }
}
// GetSettings gets the settings property value. The settings property
func (m *BusinessFlow) GetSettings()(BusinessFlowSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BusinessFlow) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCustomData sets the customData property value. The customData property
func (m *BusinessFlow) SetCustomData(value *string)() {
    if m != nil {
        m.customData = value
    }
}
// SetDeDuplicationId sets the deDuplicationId property value. The deDuplicationId property
func (m *BusinessFlow) SetDeDuplicationId(value *string)() {
    if m != nil {
        m.deDuplicationId = value
    }
}
// SetDescription sets the description property value. The description property
func (m *BusinessFlow) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *BusinessFlow) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetPolicy sets the policy property value. The policy property
func (m *BusinessFlow) SetPolicy(value GovernancePolicyable)() {
    if m != nil {
        m.policy = value
    }
}
// SetPolicyTemplateId sets the policyTemplateId property value. The policyTemplateId property
func (m *BusinessFlow) SetPolicyTemplateId(value *string)() {
    if m != nil {
        m.policyTemplateId = value
    }
}
// SetRecordVersion sets the recordVersion property value. The recordVersion property
func (m *BusinessFlow) SetRecordVersion(value *string)() {
    if m != nil {
        m.recordVersion = value
    }
}
// SetSchemaId sets the schemaId property value. The schemaId property
func (m *BusinessFlow) SetSchemaId(value *string)() {
    if m != nil {
        m.schemaId = value
    }
}
// SetSettings sets the settings property value. The settings property
func (m *BusinessFlow) SetSettings(value BusinessFlowSettingsable)() {
    if m != nil {
        m.settings = value
    }
}
