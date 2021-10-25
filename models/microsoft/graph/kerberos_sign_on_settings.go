package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type KerberosSignOnSettings struct {
    additionalData map[string]interface{};
    kerberosServicePrincipalName *string;
    kerberosSignOnMappingAttributeType *KerberosSignOnMappingAttributeType;
}
func NewKerberosSignOnSettings()(*KerberosSignOnSettings) {
    m := &KerberosSignOnSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *KerberosSignOnSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *KerberosSignOnSettings) GetKerberosServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kerberosServicePrincipalName
    }
}
func (m *KerberosSignOnSettings) GetKerberosSignOnMappingAttributeType()(*KerberosSignOnMappingAttributeType) {
    if m == nil {
        return nil
    } else {
        return m.kerberosSignOnMappingAttributeType
    }
}
func (m *KerberosSignOnSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["kerberosServicePrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKerberosServicePrincipalName(val)
        return nil
    }
    res["kerberosSignOnMappingAttributeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseKerberosSignOnMappingAttributeType)
        if err != nil {
            return err
        }
        cast := val.(KerberosSignOnMappingAttributeType)
        m.SetKerberosSignOnMappingAttributeType(&cast)
        return nil
    }
    return res
}
func (m *KerberosSignOnSettings) IsNil()(bool) {
    return m == nil
}
func (m *KerberosSignOnSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("kerberosServicePrincipalName", m.GetKerberosServicePrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetKerberosSignOnMappingAttributeType() != nil {
        cast := m.GetKerberosSignOnMappingAttributeType().String()
        err := writer.WriteStringValue("kerberosSignOnMappingAttributeType", &cast)
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
func (m *KerberosSignOnSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *KerberosSignOnSettings) SetKerberosServicePrincipalName(value *string)() {
    m.kerberosServicePrincipalName = value
}
func (m *KerberosSignOnSettings) SetKerberosSignOnMappingAttributeType(value *KerberosSignOnMappingAttributeType)() {
    m.kerberosSignOnMappingAttributeType = value
}
