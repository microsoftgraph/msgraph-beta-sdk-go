package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

type Setting struct {
    additionalData map[string]interface{};
    displayName *string;
    jsonValue *string;
    overwriteAllowed *bool;
    valueType *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType;
}
func NewSetting()(*Setting) {
    m := &Setting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Setting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Setting) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *Setting) GetJsonValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jsonValue
    }
}
func (m *Setting) GetOverwriteAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overwriteAllowed
    }
}
func (m *Setting) GetValueType()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
func (m *Setting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["jsonValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJsonValue(val)
        return nil
    }
    res["overwriteAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOverwriteAllowed(val)
        return nil
    }
    res["valueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementParameterValueType)
        if err != nil {
            return err
        }
        cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType)
        m.SetValueType(&cast)
        return nil
    }
    return res
}
func (m *Setting) IsNil()(bool) {
    return m == nil
}
func (m *Setting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jsonValue", m.GetJsonValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("overwriteAllowed", m.GetOverwriteAllowed())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := m.GetValueType().String()
        err := writer.WriteStringValue("valueType", &cast)
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
func (m *Setting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Setting) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *Setting) SetJsonValue(value *string)() {
    m.jsonValue = value
}
func (m *Setting) SetOverwriteAllowed(value *bool)() {
    m.overwriteAllowed = value
}
func (m *Setting) SetValueType(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType)() {
    m.valueType = value
}
