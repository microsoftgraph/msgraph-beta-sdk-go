package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type NetworkLocationDetail struct {
    additionalData map[string]interface{};
    networkNames []string;
    networkType *NetworkType;
}
func NewNetworkLocationDetail()(*NetworkLocationDetail) {
    m := &NetworkLocationDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *NetworkLocationDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *NetworkLocationDetail) GetNetworkNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.networkNames
    }
}
func (m *NetworkLocationDetail) GetNetworkType()(*NetworkType) {
    if m == nil {
        return nil
    } else {
        return m.networkType
    }
}
func (m *NetworkLocationDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["networkNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetNetworkNames(res)
        return nil
    }
    res["networkType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkType)
        if err != nil {
            return err
        }
        cast := val.(NetworkType)
        m.SetNetworkType(&cast)
        return nil
    }
    return res
}
func (m *NetworkLocationDetail) IsNil()(bool) {
    return m == nil
}
func (m *NetworkLocationDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("networkNames", m.GetNetworkNames())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkType() != nil {
        cast := m.GetNetworkType().String()
        err := writer.WriteStringValue("networkType", &cast)
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
func (m *NetworkLocationDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *NetworkLocationDetail) SetNetworkNames(value []string)() {
    m.networkNames = value
}
func (m *NetworkLocationDetail) SetNetworkType(value *NetworkType)() {
    m.networkType = value
}
