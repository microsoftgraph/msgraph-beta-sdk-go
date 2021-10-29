package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new securityProviderStatus and sets the default values.
func NewSecurityProviderStatus()(*SecurityProviderStatus) {
    m := &SecurityProviderStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityProviderStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the enabled property value. 
func (m *SecurityProviderStatus) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// Gets the endpoint property value. 
func (m *SecurityProviderStatus) GetEndpoint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endpoint
    }
}
// Gets the provider property value. 
func (m *SecurityProviderStatus) GetProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provider
    }
}
// Gets the region property value. 
func (m *SecurityProviderStatus) GetRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.region
    }
}
// Gets the vendor_escaped property value. 
func (m *SecurityProviderStatus) GetVendor_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// The deserialization information for the current model
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
    res["vendor_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVendor_escaped(val)
        return nil
    }
    return res
}
func (m *SecurityProviderStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        err := writer.WriteStringValue("vendor_escaped", m.GetVendor_escaped())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SecurityProviderStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the enabled property value. 
// Parameters:
//  - value : Value to set for the enabled property.
func (m *SecurityProviderStatus) SetEnabled(value *bool)() {
    m.enabled = value
}
// Sets the endpoint property value. 
// Parameters:
//  - value : Value to set for the endpoint property.
func (m *SecurityProviderStatus) SetEndpoint(value *string)() {
    m.endpoint = value
}
// Sets the provider property value. 
// Parameters:
//  - value : Value to set for the provider property.
func (m *SecurityProviderStatus) SetProvider(value *string)() {
    m.provider = value
}
// Sets the region property value. 
// Parameters:
//  - value : Value to set for the region property.
func (m *SecurityProviderStatus) SetRegion(value *string)() {
    m.region = value
}
// Sets the vendor_escaped property value. 
// Parameters:
//  - value : Value to set for the vendor_escaped property.
func (m *SecurityProviderStatus) SetVendor_escaped(value *string)() {
    m.vendor_escaped = value
}
