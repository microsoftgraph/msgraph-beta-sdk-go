package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProvisionChannelEmailResult struct {
    additionalData map[string]interface{};
    email *string;
}
func NewProvisionChannelEmailResult()(*ProvisionChannelEmailResult) {
    m := &ProvisionChannelEmailResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ProvisionChannelEmailResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ProvisionChannelEmailResult) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
func (m *ProvisionChannelEmailResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmail(val)
        return nil
    }
    return res
}
func (m *ProvisionChannelEmailResult) IsNil()(bool) {
    return m == nil
}
func (m *ProvisionChannelEmailResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
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
func (m *ProvisionChannelEmailResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ProvisionChannelEmailResult) SetEmail(value *string)() {
    m.email = value
}
