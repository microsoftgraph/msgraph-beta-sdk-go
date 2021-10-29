package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DomainRegistrant struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    countryOrRegionCode *string;
    // 
    organization *string;
    // 
    url *string;
    // 
    vendor_escaped *string;
}
// Instantiates a new domainRegistrant and sets the default values.
func NewDomainRegistrant()(*DomainRegistrant) {
    m := &DomainRegistrant{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DomainRegistrant) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the countryOrRegionCode property value. 
func (m *DomainRegistrant) GetCountryOrRegionCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegionCode
    }
}
// Gets the organization property value. 
func (m *DomainRegistrant) GetOrganization()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organization
    }
}
// Gets the url property value. 
func (m *DomainRegistrant) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// Gets the vendor_escaped property value. 
func (m *DomainRegistrant) GetVendor_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// The deserialization information for the current model
func (m *DomainRegistrant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["countryOrRegionCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountryOrRegionCode(val)
        return nil
    }
    res["organization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrganization(val)
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUrl(val)
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
func (m *DomainRegistrant) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DomainRegistrant) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("countryOrRegionCode", m.GetCountryOrRegionCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organization", m.GetOrganization())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *DomainRegistrant) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the countryOrRegionCode property value. 
// Parameters:
//  - value : Value to set for the countryOrRegionCode property.
func (m *DomainRegistrant) SetCountryOrRegionCode(value *string)() {
    m.countryOrRegionCode = value
}
// Sets the organization property value. 
// Parameters:
//  - value : Value to set for the organization property.
func (m *DomainRegistrant) SetOrganization(value *string)() {
    m.organization = value
}
// Sets the url property value. 
// Parameters:
//  - value : Value to set for the url property.
func (m *DomainRegistrant) SetUrl(value *string)() {
    m.url = value
}
// Sets the vendor_escaped property value. 
// Parameters:
//  - value : Value to set for the vendor_escaped property.
func (m *DomainRegistrant) SetVendor_escaped(value *string)() {
    m.vendor_escaped = value
}
