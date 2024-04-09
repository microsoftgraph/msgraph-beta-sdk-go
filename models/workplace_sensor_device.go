package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkplaceSensorDevice struct {
    Entity
}
// NewWorkplaceSensorDevice instantiates a new WorkplaceSensorDevice and sets the default values.
func NewWorkplaceSensorDevice()(*WorkplaceSensorDevice) {
    m := &WorkplaceSensorDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkplaceSensorDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkplaceSensorDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkplaceSensorDevice(), nil
}
// GetDescription gets the description property value. The description of the device.
// returns a *string when successful
func (m *WorkplaceSensorDevice) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The user-defined unique identifier of the device provided at the time of creation.
// returns a *string when successful
func (m *WorkplaceSensorDevice) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the device.
// returns a *string when successful
func (m *WorkplaceSensorDevice) GetDisplayName()(*string) {
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
func (m *WorkplaceSensorDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
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
    res["ipV4Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpV4Address(val)
        }
        return nil
    }
    res["ipV6Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpV6Address(val)
        }
        return nil
    }
    res["macAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacAddress(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["placeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlaceId(val)
        }
        return nil
    }
    res["sensors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkplaceSensorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkplaceSensorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkplaceSensorable)
                }
            }
            m.SetSensors(res)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
// GetIpV4Address gets the ipV4Address property value. The IPv4 address of the device.
// returns a *string when successful
func (m *WorkplaceSensorDevice) GetIpV4Address()(*string) {
    val, err := m.GetBackingStore().Get("ipV4Address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIpV6Address gets the ipV6Address property value. The IPv6 address of the device.
// returns a *string when successful
func (m *WorkplaceSensorDevice) GetIpV6Address()(*string) {
    val, err := m.GetBackingStore().Get("ipV6Address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMacAddress gets the macAddress property value. The MAC address of the device.
// returns a *string when successful
func (m *WorkplaceSensorDevice) GetMacAddress()(*string) {
    val, err := m.GetBackingStore().Get("macAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. The manufacturer of the device.
// returns a *string when successful
func (m *WorkplaceSensorDevice) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlaceId gets the placeId property value. The unique identifier of the place where the device is located. If the device is installed in a room equipped with a mailbox, this property should match the ExternalDirectoryObjectId or Microsoft Entra object ID of the room mailbox.
// returns a *string when successful
func (m *WorkplaceSensorDevice) GetPlaceId()(*string) {
    val, err := m.GetBackingStore().Get("placeId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSensors gets the sensors property value. A list of sensors associated with the device that collect and report data about physical or environmental conditions, such as occupancy, people count, inferred occupancy, temperature, and more.
// returns a []WorkplaceSensorable when successful
func (m *WorkplaceSensorDevice) GetSensors()([]WorkplaceSensorable) {
    val, err := m.GetBackingStore().Get("sensors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WorkplaceSensorable)
    }
    return nil
}
// GetTags gets the tags property value. A list of custom tags associated with the device.
// returns a []string when successful
func (m *WorkplaceSensorDevice) GetTags()([]string) {
    val, err := m.GetBackingStore().Get("tags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkplaceSensorDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
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
        err = writer.WriteStringValue("ipV4Address", m.GetIpV4Address())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ipV6Address", m.GetIpV6Address())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("macAddress", m.GetMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("placeId", m.GetPlaceId())
        if err != nil {
            return err
        }
    }
    if m.GetSensors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSensors()))
        for i, v := range m.GetSensors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sensors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description of the device.
func (m *WorkplaceSensorDevice) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The user-defined unique identifier of the device provided at the time of creation.
func (m *WorkplaceSensorDevice) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the device.
func (m *WorkplaceSensorDevice) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIpV4Address sets the ipV4Address property value. The IPv4 address of the device.
func (m *WorkplaceSensorDevice) SetIpV4Address(value *string)() {
    err := m.GetBackingStore().Set("ipV4Address", value)
    if err != nil {
        panic(err)
    }
}
// SetIpV6Address sets the ipV6Address property value. The IPv6 address of the device.
func (m *WorkplaceSensorDevice) SetIpV6Address(value *string)() {
    err := m.GetBackingStore().Set("ipV6Address", value)
    if err != nil {
        panic(err)
    }
}
// SetMacAddress sets the macAddress property value. The MAC address of the device.
func (m *WorkplaceSensorDevice) SetMacAddress(value *string)() {
    err := m.GetBackingStore().Set("macAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer of the device.
func (m *WorkplaceSensorDevice) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetPlaceId sets the placeId property value. The unique identifier of the place where the device is located. If the device is installed in a room equipped with a mailbox, this property should match the ExternalDirectoryObjectId or Microsoft Entra object ID of the room mailbox.
func (m *WorkplaceSensorDevice) SetPlaceId(value *string)() {
    err := m.GetBackingStore().Set("placeId", value)
    if err != nil {
        panic(err)
    }
}
// SetSensors sets the sensors property value. A list of sensors associated with the device that collect and report data about physical or environmental conditions, such as occupancy, people count, inferred occupancy, temperature, and more.
func (m *WorkplaceSensorDevice) SetSensors(value []WorkplaceSensorable)() {
    err := m.GetBackingStore().Set("sensors", value)
    if err != nil {
        panic(err)
    }
}
// SetTags sets the tags property value. A list of custom tags associated with the device.
func (m *WorkplaceSensorDevice) SetTags(value []string)() {
    err := m.GetBackingStore().Set("tags", value)
    if err != nil {
        panic(err)
    }
}
type WorkplaceSensorDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDeviceId()(*string)
    GetDisplayName()(*string)
    GetIpV4Address()(*string)
    GetIpV6Address()(*string)
    GetMacAddress()(*string)
    GetManufacturer()(*string)
    GetPlaceId()(*string)
    GetSensors()([]WorkplaceSensorable)
    GetTags()([]string)
    SetDescription(value *string)()
    SetDeviceId(value *string)()
    SetDisplayName(value *string)()
    SetIpV4Address(value *string)()
    SetIpV6Address(value *string)()
    SetMacAddress(value *string)()
    SetManufacturer(value *string)()
    SetPlaceId(value *string)()
    SetSensors(value []WorkplaceSensorable)()
    SetTags(value []string)()
}
