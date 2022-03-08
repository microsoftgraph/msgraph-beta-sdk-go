package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DomainRegistrant provides operations to manage the domainSecurityProfiles property of the microsoft.graph.security entity.
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
// NewDomainRegistrant instantiates a new domainRegistrant and sets the default values.
func NewDomainRegistrant()(*DomainRegistrant) {
    m := &DomainRegistrant{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDomainRegistrantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainRegistrantFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDomainRegistrant(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DomainRegistrant) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCountryOrRegionCode gets the countryOrRegionCode property value. 
func (m *DomainRegistrant) GetCountryOrRegionCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegionCode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainRegistrant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
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
// GetOrganization gets the organization property value. 
func (m *DomainRegistrant) GetOrganization()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organization
    }
}
// GetUrl gets the url property value. 
func (m *DomainRegistrant) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// GetVendor gets the vendor property value. 
func (m *DomainRegistrant) GetVendor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
func (m *DomainRegistrant) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
func (m *DomainRegistrant) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCountryOrRegionCode sets the countryOrRegionCode property value. 
func (m *DomainRegistrant) SetCountryOrRegionCode(value *string)() {
    if m != nil {
        m.countryOrRegionCode = value
    }
}
// SetOrganization sets the organization property value. 
func (m *DomainRegistrant) SetOrganization(value *string)() {
    if m != nil {
        m.organization = value
    }
}
// SetUrl sets the url property value. 
func (m *DomainRegistrant) SetUrl(value *string)() {
    if m != nil {
        m.url = value
    }
}
// SetVendor sets the vendor property value. 
func (m *DomainRegistrant) SetVendor(value *string)() {
    if m != nil {
        m.vendor_escaped = value
    }
}
