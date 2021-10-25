package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BusinessFlow struct {
    Entity
    customData *string;
    deDuplicationId *string;
    description *string;
    displayName *string;
    policy *GovernancePolicy;
    policyTemplateId *string;
    recordVersion *string;
    schemaId *string;
    settings *BusinessFlowSettings;
}
func NewBusinessFlow()(*BusinessFlow) {
    m := &BusinessFlow{
        Entity: *NewEntity(),
    }
    return m
}
func (m *BusinessFlow) GetCustomData()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customData
    }
}
func (m *BusinessFlow) GetDeDuplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deDuplicationId
    }
}
func (m *BusinessFlow) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *BusinessFlow) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *BusinessFlow) GetPolicy()(*GovernancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
func (m *BusinessFlow) GetPolicyTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyTemplateId
    }
}
func (m *BusinessFlow) GetRecordVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recordVersion
    }
}
func (m *BusinessFlow) GetSchemaId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schemaId
    }
}
func (m *BusinessFlow) GetSettings()(*BusinessFlowSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
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
func (m *BusinessFlow) SetCustomData(value *string)() {
    m.customData = value
}
func (m *BusinessFlow) SetDeDuplicationId(value *string)() {
    m.deDuplicationId = value
}
func (m *BusinessFlow) SetDescription(value *string)() {
    m.description = value
}
func (m *BusinessFlow) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *BusinessFlow) SetPolicy(value *GovernancePolicy)() {
    m.policy = value
}
func (m *BusinessFlow) SetPolicyTemplateId(value *string)() {
    m.policyTemplateId = value
}
func (m *BusinessFlow) SetRecordVersion(value *string)() {
    m.recordVersion = value
}
func (m *BusinessFlow) SetSchemaId(value *string)() {
    m.schemaId = value
}
func (m *BusinessFlow) SetSettings(value *BusinessFlowSettings)() {
    m.settings = value
}
