package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemAddress provides operations to manage the compliance singleton.
type ItemAddress struct {
    ItemFacet
    // 
    detail PhysicalAddressable;
    // Friendly name the user has assigned to this address.
    displayName *string;
    // The geocoordinates of the address.
    geoCoordinates GeoCoordinatesable;
}
// NewItemAddress instantiates a new itemAddress and sets the default values.
func NewItemAddress()(*ItemAddress) {
    m := &ItemAddress{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// CreateItemAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemAddressFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewItemAddress(), nil
}
// GetDetail gets the detail property value. 
func (m *ItemAddress) GetDetail()(PhysicalAddressable) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// GetDisplayName gets the displayName property value. Friendly name the user has assigned to this address.
func (m *ItemAddress) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(PhysicalAddressable))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["geoCoordinates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    if m == nil {
        return nil
    } else {
        return m.geoCoordinates
    }
}
func (m *ItemAddress) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ItemAddress) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetDetail sets the detail property value. 
func (m *ItemAddress) SetDetail(value PhysicalAddressable)() {
    if m != nil {
        m.detail = value
    }
}
// SetDisplayName sets the displayName property value. Friendly name the user has assigned to this address.
func (m *ItemAddress) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetGeoCoordinates sets the geoCoordinates property value. The geocoordinates of the address.
func (m *ItemAddress) SetGeoCoordinates(value GeoCoordinatesable)() {
    if m != nil {
        m.geoCoordinates = value
    }
}
