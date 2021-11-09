package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ItemAddress struct {
    ItemFacet
    // 
    detail *PhysicalAddress;
    // Friendly name the user has assigned to this address.
    displayName *string;
    // The geocoordinates of the address.
    geoCoordinates *GeoCoordinates;
}
// Instantiates a new itemAddress and sets the default values.
func NewItemAddress()(*ItemAddress) {
    m := &ItemAddress{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the detail property value. 
func (m *ItemAddress) GetDetail()(*PhysicalAddress) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// Gets the displayName property value. Friendly name the user has assigned to this address.
func (m *ItemAddress) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the geoCoordinates property value. The geocoordinates of the address.
func (m *ItemAddress) GetGeoCoordinates()(*GeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.geoCoordinates
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the detail property value. 
// Parameters:
//  - value : Value to set for the detail property.
func (m *ItemAddress) SetDetail(value *PhysicalAddress)() {
    m.detail = value
}
// Sets the displayName property value. Friendly name the user has assigned to this address.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ItemAddress) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the geoCoordinates property value. The geocoordinates of the address.
// Parameters:
//  - value : Value to set for the geoCoordinates property.
func (m *ItemAddress) SetGeoCoordinates(value *GeoCoordinates)() {
    m.geoCoordinates = value
}
