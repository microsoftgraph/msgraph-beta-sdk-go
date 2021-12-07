package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProviderTenantSetting 
type ProviderTenantSetting struct {
    Entity
    // 
    azureTenantId *string;
    // 
    enabled *bool;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    provider *string;
    // 
    vendor_escaped *string;
}
// NewProviderTenantSetting instantiates a new providerTenantSetting and sets the default values.
func NewProviderTenantSetting()(*ProviderTenantSetting) {
    m := &ProviderTenantSetting{
        Entity: *NewEntity(),
    }
    return m
}
// GetAzureTenantId gets the azureTenantId property value. 
func (m *ProviderTenantSetting) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// GetEnabled gets the enabled property value. 
func (m *ProviderTenantSetting) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *ProviderTenantSetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetProvider gets the provider property value. 
func (m *ProviderTenantSetting) GetProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provider
    }
}
// GetVendor gets the vendor property value. 
func (m *ProviderTenantSetting) GetVendor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProviderTenantSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
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
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
func (m *ProviderTenantSetting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProviderTenantSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provider", m.GetProvider())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendor", m.GetVendor())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureTenantId sets the azureTenantId property value. 
func (m *ProviderTenantSetting) SetAzureTenantId(value *string)() {
    if m != nil {
        m.azureTenantId = value
    }
}
// SetEnabled sets the enabled property value. 
func (m *ProviderTenantSetting) SetEnabled(value *bool)() {
    if m != nil {
        m.enabled = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *ProviderTenantSetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetProvider sets the provider property value. 
func (m *ProviderTenantSetting) SetProvider(value *string)() {
    if m != nil {
        m.provider = value
    }
}
// SetVendor sets the vendor property value. 
func (m *ProviderTenantSetting) SetVendor(value *string)() {
    if m != nil {
        m.vendor_escaped = value
    }
}
