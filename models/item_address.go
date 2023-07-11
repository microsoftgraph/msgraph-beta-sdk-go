package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemAddress 
type ItemAddress struct {
    ItemFacet
    // The OdataType property
    OdataType *string
}
// NewItemAddress instantiates a new itemAddress and sets the default values.
func NewItemAddress()(*ItemAddress) {
    m := &ItemAddress{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.itemAddress"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateItemAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemAddressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAddress(), nil
}
// GetDetail gets the detail property value. The detail property
func (m *ItemAddress) GetDetail()(PhysicalAddressable) {
    val, err := m.GetBackingStore().Get("detail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PhysicalAddressable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Friendly name the user has assigned to this address.
func (m *ItemAddress) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemAddress) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(PhysicalAddressable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["geoCoordinates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGeoCoordinatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeoCoordinates(val.(GeoCoordinatesable))
        }
        return nil
    }
    return res
}
// GetGeoCoordinates gets the geoCoordinates property value. The geocoordinates of the address.
func (m *ItemAddress) GetGeoCoordinates()(GeoCoordinatesable) {
    val, err := m.GetBackingStore().Get("geoCoordinates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GeoCoordinatesable)
    }
    return nil
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
    err := m.GetBackingStore().Set("detail", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Friendly name the user has assigned to this address.
func (m *ItemAddress) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetGeoCoordinates sets the geoCoordinates property value. The geocoordinates of the address.
func (m *ItemAddress) SetGeoCoordinates(value GeoCoordinatesable)() {
    err := m.GetBackingStore().Set("geoCoordinates", value)
    if err != nil {
        panic(err)
    }
}
// ItemAddressable 
type ItemAddressable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetail()(PhysicalAddressable)
    GetDisplayName()(*string)
    GetGeoCoordinates()(GeoCoordinatesable)
    SetDetail(value PhysicalAddressable)()
    SetDisplayName(value *string)()
    SetGeoCoordinates(value GeoCoordinatesable)()
}
