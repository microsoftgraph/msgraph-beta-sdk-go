package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TargetPolicyEndpoints struct {
    additionalData map[string]interface{};
    platformTypes []string;
}
func NewTargetPolicyEndpoints()(*TargetPolicyEndpoints) {
    m := &TargetPolicyEndpoints{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TargetPolicyEndpoints) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TargetPolicyEndpoints) GetPlatformTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.platformTypes
    }
}
func (m *TargetPolicyEndpoints) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["platformTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetPlatformTypes(res)
        return nil
    }
    return res
}
func (m *TargetPolicyEndpoints) IsNil()(bool) {
    return m == nil
}
func (m *TargetPolicyEndpoints) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("platformTypes", m.GetPlatformTypes())
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
func (m *TargetPolicyEndpoints) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TargetPolicyEndpoints) SetPlatformTypes(value []string)() {
    m.platformTypes = value
}
