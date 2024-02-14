package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TeamworkPeripheral struct {
    Entity
}
// NewTeamworkPeripheral instantiates a new TeamworkPeripheral and sets the default values.
func NewTeamworkPeripheral()(*TeamworkPeripheral) {
    m := &TeamworkPeripheral{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkPeripheralFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamworkPeripheralFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkPeripheral(), nil
}
// GetDisplayName gets the displayName property value. Display name for the peripheral.
// returns a *string when successful
func (m *TeamworkPeripheral) GetDisplayName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamworkPeripheral) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["productId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductId(val)
        }
        return nil
    }
    res["vendorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetProductId gets the productId property value. The product ID of the device. Each product from a vendor has its own ID.
// returns a *string when successful
func (m *TeamworkPeripheral) GetProductId()(*string) {
    val, err := m.GetBackingStore().Get("productId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVendorId gets the vendorId property value. The unique identifier for the vendor of the device. Each vendor has a unique ID.
// returns a *string when successful
func (m *TeamworkPeripheral) GetVendorId()(*string) {
    val, err := m.GetBackingStore().Get("vendorId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkPeripheral) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetProductId sets the productId property value. The product ID of the device. Each product from a vendor has its own ID.
func (m *TeamworkPeripheral) SetProductId(value *string)() {
    err := m.GetBackingStore().Set("productId", value)
    if err != nil {
        panic(err)
    }
}
// SetVendorId sets the vendorId property value. The unique identifier for the vendor of the device. Each vendor has a unique ID.
func (m *TeamworkPeripheral) SetVendorId(value *string)() {
    err := m.GetBackingStore().Set("vendorId", value)
    if err != nil {
        panic(err)
    }
}
type TeamworkPeripheralable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetProductId()(*string)
    GetVendorId()(*string)
    SetDisplayName(value *string)()
    SetProductId(value *string)()
    SetVendorId(value *string)()
}
