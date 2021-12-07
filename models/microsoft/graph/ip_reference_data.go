package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IpReferenceData 
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
// NewIpReferenceData instantiates a new ipReferenceData and sets the default values.
func NewIpReferenceData()(*IpReferenceData) {
    m := &IpReferenceData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IpReferenceData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAsn gets the asn property value. 
func (m *IpReferenceData) GetAsn()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.asn
    }
}
// GetCity gets the city property value. 
func (m *IpReferenceData) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// GetCountryOrRegionCode gets the countryOrRegionCode property value. 
func (m *IpReferenceData) GetCountryOrRegionCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegionCode
    }
}
// GetOrganization gets the organization property value. 
func (m *IpReferenceData) GetOrganization()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organization
    }
}
// GetState gets the state property value. 
func (m *IpReferenceData) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetVendor gets the vendor property value. 
func (m *IpReferenceData) GetVendor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IpReferenceData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["asn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAsn(val)
        }
        return nil
    }
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["countryOrRegionCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryOrRegionCode(val)
        }
        return nil
    }
    res["organization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
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
func (m *IpReferenceData) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
func (m *IpReferenceData) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAsn sets the asn property value. 
func (m *IpReferenceData) SetAsn(value *int64)() {
    if m != nil {
        m.asn = value
    }
}
// SetCity sets the city property value. 
func (m *IpReferenceData) SetCity(value *string)() {
    if m != nil {
        m.city = value
    }
}
// SetCountryOrRegionCode sets the countryOrRegionCode property value. 
func (m *IpReferenceData) SetCountryOrRegionCode(value *string)() {
    if m != nil {
        m.countryOrRegionCode = value
    }
}
// SetOrganization sets the organization property value. 
func (m *IpReferenceData) SetOrganization(value *string)() {
    if m != nil {
        m.organization = value
    }
}
// SetState sets the state property value. 
func (m *IpReferenceData) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
// SetVendor sets the vendor property value. 
func (m *IpReferenceData) SetVendor(value *string)() {
    if m != nil {
        m.vendor_escaped = value
    }
}
