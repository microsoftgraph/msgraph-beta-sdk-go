package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemAddress 
type ItemAddress struct {
    ItemFacet
    // The detail property
    detail PhysicalAddressable
    // Friendly name the user has assigned to this address.
    displayName *string
    // The geocoordinates of the address.
    geoCoordinates GeoCoordinatesable
}
// NewItemAddress instantiates a new ItemAddress and sets the default values.
func NewItemAddress()(*ItemAddress) {
    m := &ItemAddress{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.itemAddress";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateItemAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemAddressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAddress(), nil
}
// GetDetail gets the detail property value. The detail property
func (m *ItemAddress) GetDetail()(PhysicalAddressable) {
    return m.detail
}
// GetDisplayName gets the displayName property value. Friendly name the user has assigned to this address.
func (m *ItemAddress) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemAddress) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["detail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePhysicalAddressFromDiscriminatorValue , m.SetDetail)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["geoCoordinates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGeoCoordinatesFromDiscriminatorValue , m.SetGeoCoordinates)
    return res
}
// GetGeoCoordinates gets the geoCoordinates property value. The geocoordinates of the address.
func (m *ItemAddress) GetGeoCoordinates()(GeoCoordinatesable) {
    return m.geoCoordinates
}
// Serialize serializes information the current object
func (m *ItemAddress) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("geoCoordinates", m.GetGeoCoordinates())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetail sets the detail property value. The detail property
func (m *ItemAddress) SetDetail(value PhysicalAddressable)() {
    m.detail = value
}
// SetDisplayName sets the displayName property value. Friendly name the user has assigned to this address.
func (m *ItemAddress) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGeoCoordinates sets the geoCoordinates property value. The geocoordinates of the address.
func (m *ItemAddress) SetGeoCoordinates(value GeoCoordinatesable)() {
    m.geoCoordinates = value
}
