package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainRegistrant 
type DomainRegistrant struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The countryOrRegionCode property
    countryOrRegionCode *string
    // The organization property
    organization *string
    // The url property
    url *string
    // The vendor property
    vendor_escaped *string
}
// NewDomainRegistrant instantiates a new domainRegistrant and sets the default values.
func NewDomainRegistrant()(*DomainRegistrant) {
    m := &DomainRegistrant{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDomainRegistrantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainRegistrantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
// GetCountryOrRegionCode gets the countryOrRegionCode property value. The countryOrRegionCode property
func (m *DomainRegistrant) GetCountryOrRegionCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegionCode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainRegistrant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["countryOrRegionCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryOrRegionCode(val)
        }
        return nil
    }
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganization(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetOrganization gets the organization property value. The organization property
func (m *DomainRegistrant) GetOrganization()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organization
    }
}
// GetUrl gets the url property value. The url property
func (m *DomainRegistrant) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// GetVendor gets the vendor property value. The vendor property
func (m *DomainRegistrant) GetVendor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// Serialize serializes information the current object
func (m *DomainRegistrant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetCountryOrRegionCode sets the countryOrRegionCode property value. The countryOrRegionCode property
func (m *DomainRegistrant) SetCountryOrRegionCode(value *string)() {
    if m != nil {
        m.countryOrRegionCode = value
    }
}
// SetOrganization sets the organization property value. The organization property
func (m *DomainRegistrant) SetOrganization(value *string)() {
    if m != nil {
        m.organization = value
    }
}
// SetUrl sets the url property value. The url property
func (m *DomainRegistrant) SetUrl(value *string)() {
    if m != nil {
        m.url = value
    }
}
// SetVendor sets the vendor property value. The vendor property
func (m *DomainRegistrant) SetVendor(value *string)() {
    if m != nil {
        m.vendor_escaped = value
    }
}
