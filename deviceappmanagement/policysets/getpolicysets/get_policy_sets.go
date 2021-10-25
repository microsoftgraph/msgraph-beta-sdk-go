package getpolicysets

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GetPolicySets struct {
    additionalData map[string]interface{};
}
func NewGetPolicySets()(*GetPolicySets) {
    m := &GetPolicySets{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GetPolicySets) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GetPolicySets) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *GetPolicySets) IsNil()(bool) {
    return m == nil
}
func (m *GetPolicySets) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetPolicySets) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
