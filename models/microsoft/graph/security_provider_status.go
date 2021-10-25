package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SecurityProviderStatus struct {
    additionalData map[string]interface{};
    enabled *bool;
    endpoint *string;
    provider *string;
    region *string;
    vendor *string;
}
func NewSecurityProviderStatus()(*SecurityProviderStatus) {
    m := &SecurityProviderStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SecurityProviderStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SecurityProviderStatus) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
func (m *SecurityProviderStatus) GetEndpoint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endpoint
    }
}
func (m *SecurityProviderStatus) GetProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provider
    }
}
func (m *SecurityProviderStatus) GetRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.region
    }
}
func (m *SecurityProviderStatus) GetVendor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor
    }
}
func (m *SecurityProviderStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnabled(val)
        return nil
    }
    res["endpoint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEndpoint(val)
        return nil
    }
    res["provider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProvider(val)
        return nil
    }
    res["region"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRegion(val)
        return nil
    }
    res["vendor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVendor(val)
        return nil
    }
    return res
}
func (m *SecurityProviderStatus) IsNil()(bool) {
    return m == nil
}
func (m *SecurityProviderStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("endpoint", m.GetEndpoint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("provider", m.GetProvider())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vendor", m.GetVendor())
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
func (m *SecurityProviderStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SecurityProviderStatus) SetEnabled(value *bool)() {
    m.enabled = value
}
func (m *SecurityProviderStatus) SetEndpoint(value *string)() {
    m.endpoint = value
}
func (m *SecurityProviderStatus) SetProvider(value *string)() {
    m.provider = value
}
func (m *SecurityProviderStatus) SetRegion(value *string)() {
    m.region = value
}
func (m *SecurityProviderStatus) SetVendor(value *string)() {
    m.vendor = value
}
