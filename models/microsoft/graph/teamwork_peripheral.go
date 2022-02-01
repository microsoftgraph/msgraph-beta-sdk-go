package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkPeripheral 
type TeamworkPeripheral struct {
    Entity
    // Display name for the peripheral.
    displayName *string;
    // The product ID of the device. Each product from a vendor has its own ID.
    productId *string;
    // The unique identifier for the vendor of the device. Each vendor has a unique ID.
    vendorId *string;
}
// NewTeamworkPeripheral instantiates a new teamworkPeripheral and sets the default values.
func NewTeamworkPeripheral()(*TeamworkPeripheral) {
    m := &TeamworkPeripheral{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. Display name for the peripheral.
func (m *TeamworkPeripheral) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetProductId gets the productId property value. The product ID of the device. Each product from a vendor has its own ID.
func (m *TeamworkPeripheral) GetProductId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productId
    }
}
// GetVendorId gets the vendorId property value. The unique identifier for the vendor of the device. Each vendor has a unique ID.
func (m *TeamworkPeripheral) GetVendorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendorId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkPeripheral) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["productId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductId(val)
        }
        return nil
    }
    res["vendorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorId(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkPeripheral) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkPeripheral) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productId", m.GetProductId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendorId", m.GetVendorId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Display name for the peripheral.
func (m *TeamworkPeripheral) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetProductId sets the productId property value. The product ID of the device. Each product from a vendor has its own ID.
func (m *TeamworkPeripheral) SetProductId(value *string)() {
    if m != nil {
        m.productId = value
    }
}
// SetVendorId sets the vendorId property value. The unique identifier for the vendor of the device. Each vendor has a unique ID.
func (m *TeamworkPeripheral) SetVendorId(value *string)() {
    if m != nil {
        m.vendorId = value
    }
}
