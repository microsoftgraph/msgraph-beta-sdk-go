package getazureaduserfeatureusage

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GetAzureADUserFeatureUsage struct {
    additionalData map[string]interface{};
}
func NewGetAzureADUserFeatureUsage()(*GetAzureADUserFeatureUsage) {
    m := &GetAzureADUserFeatureUsage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GetAzureADUserFeatureUsage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GetAzureADUserFeatureUsage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *GetAzureADUserFeatureUsage) IsNil()(bool) {
    return m == nil
}
func (m *GetAzureADUserFeatureUsage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetAzureADUserFeatureUsage) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
