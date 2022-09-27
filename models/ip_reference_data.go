package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpReferenceData 
type IpReferenceData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The asn property
    asn *int64
    // The city property
    city *string
    // The countryOrRegionCode property
    countryOrRegionCode *string
    // The OdataType property
    odataType *string
    // The organization property
    organization *string
    // The state property
    state *string
    // The vendor property
    vendor_escaped *string
}
// NewIpReferenceData instantiates a new ipReferenceData and sets the default values.
func NewIpReferenceData()(*IpReferenceData) {
    m := &IpReferenceData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.ipReferenceData";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIpReferenceDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIpReferenceDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpReferenceData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IpReferenceData) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAsn gets the asn property value. The asn property
func (m *IpReferenceData) GetAsn()(*int64) {
    return m.asn
}
// GetCity gets the city property value. The city property
func (m *IpReferenceData) GetCity()(*string) {
    return m.city
}
// GetCountryOrRegionCode gets the countryOrRegionCode property value. The countryOrRegionCode property
func (m *IpReferenceData) GetCountryOrRegionCode()(*string) {
    return m.countryOrRegionCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IpReferenceData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["asn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetAsn)
    res["city"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCity)
    res["countryOrRegionCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCountryOrRegionCode)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["organization"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOrganization)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetState)
    res["vendor"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVendor)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IpReferenceData) GetOdataType()(*string) {
    return m.odataType
}
// GetOrganization gets the organization property value. The organization property
func (m *IpReferenceData) GetOrganization()(*string) {
    return m.organization
}
// GetState gets the state property value. The state property
func (m *IpReferenceData) GetState()(*string) {
    return m.state
}
// GetVendor gets the vendor property value. The vendor property
func (m *IpReferenceData) GetVendor()(*string) {
    return m.vendor_escaped
}
// Serialize serializes information the current object
func (m *IpReferenceData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    m.additionalData = value
}
// SetAsn sets the asn property value. The asn property
func (m *IpReferenceData) SetAsn(value *int64)() {
    m.asn = value
}
// SetCity sets the city property value. The city property
func (m *IpReferenceData) SetCity(value *string)() {
    m.city = value
}
// SetCountryOrRegionCode sets the countryOrRegionCode property value. The countryOrRegionCode property
func (m *IpReferenceData) SetCountryOrRegionCode(value *string)() {
    m.countryOrRegionCode = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IpReferenceData) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrganization sets the organization property value. The organization property
func (m *IpReferenceData) SetOrganization(value *string)() {
    m.organization = value
}
// SetState sets the state property value. The state property
func (m *IpReferenceData) SetState(value *string)() {
    m.state = value
}
// SetVendor sets the vendor property value. The vendor property
func (m *IpReferenceData) SetVendor(value *string)() {
    m.vendor_escaped = value
}
