package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
    policy *GovernancePolicy;
    // 
    policyTemplateId *string;
    // 
    recordVersion *string;
    // 
    schemaId *string;
    // 
    settings *BusinessFlowSettings;
}
// Instantiates a new businessFlow and sets the default values.
func NewBusinessFlow()(*BusinessFlow) {
    m := &BusinessFlow{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the customData property value. 
func (m *BusinessFlow) GetCustomData()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customData
    }
}
// Gets the deDuplicationId property value. 
func (m *BusinessFlow) GetDeDuplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deDuplicationId
    }
}
// Gets the description property value. 
func (m *BusinessFlow) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. 
func (m *BusinessFlow) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the policy property value. 
func (m *BusinessFlow) GetPolicy()(*GovernancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
// Gets the policyTemplateId property value. 
func (m *BusinessFlow) GetPolicyTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyTemplateId
    }
}
// Gets the recordVersion property value. 
func (m *BusinessFlow) GetRecordVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordVersion
    }
}
// Gets the schemaId property value. 
func (m *BusinessFlow) GetSchemaId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schemaId
    }
}
// Gets the settings property value. 
func (m *BusinessFlow) GetSettings()(*BusinessFlowSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// The deserialization information for the current model
func (m *BusinessFlow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomData(val)
        return nil
    }
    res["deDuplicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeDuplicationId(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["policy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernancePolicy() })
        if err != nil {
            return err
        }
        m.SetPolicy(val.(*GovernancePolicy))
        return nil
    }
    res["policyTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyTemplateId(val)
        return nil
    }
    res["recordVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecordVersion(val)
        return nil
    }
    res["schemaId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSchemaId(val)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBusinessFlowSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*BusinessFlowSettings))
        return nil
    }
    return res
}
func (m *BusinessFlow) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the customData property value. 
// Parameters:
//  - value : Value to set for the customData property.
func (m *BusinessFlow) SetCustomData(value *string)() {
    m.customData = value
}
// Sets the deDuplicationId property value. 
// Parameters:
//  - value : Value to set for the deDuplicationId property.
func (m *BusinessFlow) SetDeDuplicationId(value *string)() {
    m.deDuplicationId = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *BusinessFlow) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *BusinessFlow) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the policy property value. 
// Parameters:
//  - value : Value to set for the policy property.
func (m *BusinessFlow) SetPolicy(value *GovernancePolicy)() {
    m.policy = value
}
// Sets the policyTemplateId property value. 
// Parameters:
//  - value : Value to set for the policyTemplateId property.
func (m *BusinessFlow) SetPolicyTemplateId(value *string)() {
    m.policyTemplateId = value
}
// Sets the recordVersion property value. 
// Parameters:
//  - value : Value to set for the recordVersion property.
func (m *BusinessFlow) SetRecordVersion(value *string)() {
    m.recordVersion = value
}
// Sets the schemaId property value. 
// Parameters:
//  - value : Value to set for the schemaId property.
func (m *BusinessFlow) SetSchemaId(value *string)() {
    m.schemaId = value
}
// Sets the settings property value. 
// Parameters:
//  - value : Value to set for the settings property.
func (m *BusinessFlow) SetSettings(value *BusinessFlowSettings)() {
    m.settings = value
}
