package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BusinessFlow provides operations to manage the collection of approvalWorkflowProvider entities.
type BusinessFlow struct {
    Entity
    // 
    customData *string;
    // 
    deDuplicationId *string;
    // 
    description *string;
    // 
    displayName *string;
    // 
    policy GovernancePolicyable;
    // 
    policyTemplateId *string;
    // 
    recordVersion *string;
    // 
    schemaId *string;
    // 
    settings BusinessFlowSettingsable;
}
// NewBusinessFlow instantiates a new businessFlow and sets the default values.
func NewBusinessFlow()(*BusinessFlow) {
    m := &BusinessFlow{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBusinessFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessFlowFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBusinessFlow(), nil
}
// GetCustomData gets the customData property value. 
func (m *BusinessFlow) GetCustomData()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customData
    }
}
// GetDeDuplicationId gets the deDuplicationId property value. 
func (m *BusinessFlow) GetDeDuplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deDuplicationId
    }
}
// GetDescription gets the description property value. 
func (m *BusinessFlow) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. 
func (m *BusinessFlow) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessFlow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomData(val)
        }
        return nil
    }
    res["deDuplicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeDuplicationId(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["policy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val.(GovernancePolicyable))
        }
        return nil
    }
    res["policyTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyTemplateId(val)
        }
        return nil
    }
    res["recordVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordVersion(val)
        }
        return nil
    }
    res["schemaId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchemaId(val)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetPolicy gets the policy property value. 
func (m *BusinessFlow) GetPolicy()(GovernancePolicyable) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
// GetPolicyTemplateId gets the policyTemplateId property value. 
func (m *BusinessFlow) GetPolicyTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyTemplateId
    }
}
// GetRecordVersion gets the recordVersion property value. 
func (m *BusinessFlow) GetRecordVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordVersion
    }
}
// GetSchemaId gets the schemaId property value. 
func (m *BusinessFlow) GetSchemaId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schemaId
    }
}
// GetSettings gets the settings property value. 
func (m *BusinessFlow) GetSettings()(BusinessFlowSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *BusinessFlow) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BusinessFlow) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetCustomData sets the customData property value. 
func (m *BusinessFlow) SetCustomData(value *string)() {
    if m != nil {
        m.customData = value
    }
}
// SetDeDuplicationId sets the deDuplicationId property value. 
func (m *BusinessFlow) SetDeDuplicationId(value *string)() {
    if m != nil {
        m.deDuplicationId = value
    }
}
// SetDescription sets the description property value. 
func (m *BusinessFlow) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *BusinessFlow) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetPolicy sets the policy property value. 
func (m *BusinessFlow) SetPolicy(value GovernancePolicyable)() {
    if m != nil {
        m.policy = value
    }
}
// SetPolicyTemplateId sets the policyTemplateId property value. 
func (m *BusinessFlow) SetPolicyTemplateId(value *string)() {
    if m != nil {
        m.policyTemplateId = value
    }
}
// SetRecordVersion sets the recordVersion property value. 
func (m *BusinessFlow) SetRecordVersion(value *string)() {
    if m != nil {
        m.recordVersion = value
    }
}
// SetSchemaId sets the schemaId property value. 
func (m *BusinessFlow) SetSchemaId(value *string)() {
    if m != nil {
        m.schemaId = value
    }
}
// SetSettings sets the settings property value. 
func (m *BusinessFlow) SetSettings(value BusinessFlowSettingsable)() {
    if m != nil {
        m.settings = value
    }
}
