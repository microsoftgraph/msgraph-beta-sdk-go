package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// NetworkInterface 
type NetworkInterface struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewNetworkInterface instantiates a new networkInterface and sets the default values.
func NewNetworkInterface()(*NetworkInterface) {
    m := &NetworkInterface{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNetworkInterfaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNetworkInterfaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNetworkInterface(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkInterface) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *NetworkInterface) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDescription gets the description property value. Description of the NIC (for example, Ethernet adapter, Wireless LAN adapter Local Area Connection, and so on).
func (m *NetworkInterface) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NetworkInterface) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["localIpV6Address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalIpV6Address(val)
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetIpV4Address gets the ipV4Address property value. Last IPv4 address associated with this NIC.
func (m *NetworkInterface) GetIpV4Address()(*string) {
    val, err := m.GetBackingStore().Get("ipV4Address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIpV6Address gets the ipV6Address property value. Last Public (also known as global) IPv6 address associated with this NIC.
func (m *NetworkInterface) GetIpV6Address()(*string) {
    val, err := m.GetBackingStore().Get("ipV6Address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLocalIpV6Address gets the localIpV6Address property value. Last local (link-local or site-local) IPv6 address associated with this NIC.
func (m *NetworkInterface) GetLocalIpV6Address()(*string) {
    val, err := m.GetBackingStore().Get("localIpV6Address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMacAddress gets the macAddress property value. MAC address of the NIC on this host.
func (m *NetworkInterface) GetMacAddress()(*string) {
    val, err := m.GetBackingStore().Get("macAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *NetworkInterface) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *NetworkInterface) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipV4Address", m.GetIpV4Address())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipV6Address", m.GetIpV6Address())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("localIpV6Address", m.GetLocalIpV6Address())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macAddress", m.GetMacAddress())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkInterface) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *NetworkInterface) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDescription sets the description property value. Description of the NIC (for example, Ethernet adapter, Wireless LAN adapter Local Area Connection, and so on).
func (m *NetworkInterface) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetIpV4Address sets the ipV4Address property value. Last IPv4 address associated with this NIC.
func (m *NetworkInterface) SetIpV4Address(value *string)() {
    err := m.GetBackingStore().Set("ipV4Address", value)
    if err != nil {
        panic(err)
    }
}
// SetIpV6Address sets the ipV6Address property value. Last Public (also known as global) IPv6 address associated with this NIC.
func (m *NetworkInterface) SetIpV6Address(value *string)() {
    err := m.GetBackingStore().Set("ipV6Address", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalIpV6Address sets the localIpV6Address property value. Last local (link-local or site-local) IPv6 address associated with this NIC.
func (m *NetworkInterface) SetLocalIpV6Address(value *string)() {
    err := m.GetBackingStore().Set("localIpV6Address", value)
    if err != nil {
        panic(err)
    }
}
// SetMacAddress sets the macAddress property value. MAC address of the NIC on this host.
func (m *NetworkInterface) SetMacAddress(value *string)() {
    err := m.GetBackingStore().Set("macAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NetworkInterface) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// NetworkInterfaceable 
type NetworkInterfaceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDescription()(*string)
    GetIpV4Address()(*string)
    GetIpV6Address()(*string)
    GetLocalIpV6Address()(*string)
    GetMacAddress()(*string)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDescription(value *string)()
    SetIpV4Address(value *string)()
    SetIpV6Address(value *string)()
    SetLocalIpV6Address(value *string)()
    SetMacAddress(value *string)()
    SetOdataType(value *string)()
}
