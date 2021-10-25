package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type IpReferenceData struct {
    additionalData map[string]interface{};
    asn *int64;
    city *string;
    countryOrRegionCode *string;
    organization *string;
    state *string;
    vendor *string;
}
func NewIpReferenceData()(*IpReferenceData) {
    m := &IpReferenceData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *IpReferenceData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *IpReferenceData) GetAsn()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.asn
    }
}
func (m *IpReferenceData) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
func (m *IpReferenceData) GetCountryOrRegionCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegionCode
    }
}
func (m *IpReferenceData) GetOrganization()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organization
    }
}
func (m *IpReferenceData) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *IpReferenceData) GetVendor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor
    }
}
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
func (m *IpReferenceData) IsNil()(bool) {
    return m == nil
}
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
func (m *IpReferenceData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *IpReferenceData) SetAsn(value *int64)() {
    m.asn = value
}
func (m *IpReferenceData) SetCity(value *string)() {
    m.city = value
}
func (m *IpReferenceData) SetCountryOrRegionCode(value *string)() {
    m.countryOrRegionCode = value
}
func (m *IpReferenceData) SetOrganization(value *string)() {
    m.organization = value
}
func (m *IpReferenceData) SetState(value *string)() {
    m.state = value
}
func (m *IpReferenceData) SetVendor(value *string)() {
    m.vendor = value
}
