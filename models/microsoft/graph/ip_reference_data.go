package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IpReferenceData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    asn *int64;
    // 
    city *string;
    // 
    countryOrRegionCode *string;
    // 
    organization *string;
    // 
    state *string;
    // 
    vendor_escaped *string;
}
// Instantiates a new ipReferenceData and sets the default values.
func NewIpReferenceData()(*IpReferenceData) {
    m := &IpReferenceData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IpReferenceData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the asn property value. 
func (m *IpReferenceData) GetAsn()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.asn
    }
}
// Gets the city property value. 
func (m *IpReferenceData) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// Gets the countryOrRegionCode property value. 
func (m *IpReferenceData) GetCountryOrRegionCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegionCode
    }
}
// Gets the organization property value. 
func (m *IpReferenceData) GetOrganization()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organization
    }
}
// Gets the state property value. 
func (m *IpReferenceData) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the vendor_escaped property value. 
func (m *IpReferenceData) GetVendor_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// The deserialization information for the current model
func (m *IpReferenceData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["asn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAsn(val)
        return nil
    }
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCity(val)
        return nil
    }
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
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetState(val)
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
func (m *IpReferenceData) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IpReferenceData) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("asn", m.GetAsn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteStringValue("state", m.GetState())
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
func (m *IpReferenceData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the asn property value. 
// Parameters:
//  - value : Value to set for the asn property.
func (m *IpReferenceData) SetAsn(value *int64)() {
    m.asn = value
}
// Sets the city property value. 
// Parameters:
//  - value : Value to set for the city property.
func (m *IpReferenceData) SetCity(value *string)() {
    m.city = value
}
// Sets the countryOrRegionCode property value. 
// Parameters:
//  - value : Value to set for the countryOrRegionCode property.
func (m *IpReferenceData) SetCountryOrRegionCode(value *string)() {
    m.countryOrRegionCode = value
}
// Sets the organization property value. 
// Parameters:
//  - value : Value to set for the organization property.
func (m *IpReferenceData) SetOrganization(value *string)() {
    m.organization = value
}
// Sets the state property value. 
// Parameters:
//  - value : Value to set for the state property.
func (m *IpReferenceData) SetState(value *string)() {
    m.state = value
}
// Sets the vendor_escaped property value. 
// Parameters:
//  - value : Value to set for the vendor_escaped property.
func (m *IpReferenceData) SetVendor_escaped(value *string)() {
    m.vendor_escaped = value
}
