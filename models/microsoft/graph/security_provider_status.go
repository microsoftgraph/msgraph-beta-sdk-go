package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SecurityProviderStatus 
type SecurityProviderStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    enabled *bool;
    // 
    endpoint *string;
    // 
    provider *string;
    // 
    region *string;
    // 
    vendor_escaped *string;
}
// NewSecurityProviderStatus instantiates a new securityProviderStatus and sets the default values.
func NewSecurityProviderStatus()(*SecurityProviderStatus) {
    m := &SecurityProviderStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityProviderStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnabled gets the enabled property value. 
func (m *SecurityProviderStatus) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// GetEndpoint gets the endpoint property value. 
func (m *SecurityProviderStatus) GetEndpoint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endpoint
    }
}
// GetProvider gets the provider property value. 
func (m *SecurityProviderStatus) GetProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provider
    }
}
// GetRegion gets the region property value. 
func (m *SecurityProviderStatus) GetRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.region
    }
}
// GetVendor gets the vendor property value. 
func (m *SecurityProviderStatus) GetVendor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityProviderStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["endpoint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpoint(val)
        }
        return nil
    }
    res["provider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvider(val)
        }
        return nil
    }
    res["region"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["vendor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendor(val)
        }
        return nil
    }
    return res
}
func (m *SecurityProviderStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityProviderStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnabled sets the enabled property value. 
func (m *SecurityProviderStatus) SetEnabled(value *bool)() {
    if m != nil {
        m.enabled = value
    }
}
// SetEndpoint sets the endpoint property value. 
func (m *SecurityProviderStatus) SetEndpoint(value *string)() {
    if m != nil {
        m.endpoint = value
    }
}
// SetProvider sets the provider property value. 
func (m *SecurityProviderStatus) SetProvider(value *string)() {
    if m != nil {
        m.provider = value
    }
}
// SetRegion sets the region property value. 
func (m *SecurityProviderStatus) SetRegion(value *string)() {
    if m != nil {
        m.region = value
    }
}
// SetVendor sets the vendor property value. 
func (m *SecurityProviderStatus) SetVendor(value *string)() {
    if m != nil {
        m.vendor_escaped = value
    }
}
