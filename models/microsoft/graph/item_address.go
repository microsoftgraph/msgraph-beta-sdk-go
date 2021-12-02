package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemAddress 
type ItemAddress struct {
    ItemFacet
    // 
    detail *PhysicalAddress;
    // Friendly name the user has assigned to this address.
    displayName *string;
    // The geocoordinates of the address.
    geoCoordinates *GeoCoordinates;
}
// NewItemAddress instantiates a new itemAddress and sets the default values.
func NewItemAddress()(*ItemAddress) {
    m := &ItemAddress{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// GetDetail gets the detail property value. 
func (m *ItemAddress) GetDetail()(*PhysicalAddress) {
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
// GetGeoCoordinates gets the geoCoordinates property value. The geocoordinates of the address.
func (m *ItemAddress) GetGeoCoordinates()(*GeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.geoCoordinates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhysicalAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(*PhysicalAddress))
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGeoCoordinates() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeoCoordinates(val.(*GeoCoordinates))
        }
        return nil
    }
    return res
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
func (m *ItemAddress) SetDetail(value *PhysicalAddress)() {
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
func (m *ItemAddress) SetGeoCoordinates(value *GeoCoordinates)() {
    if m != nil {
        m.geoCoordinates = value
    }
}
