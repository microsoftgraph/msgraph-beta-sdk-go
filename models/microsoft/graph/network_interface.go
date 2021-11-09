package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type NetworkInterface struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Description of the NIC (e.g. Ethernet adapter, Wireless LAN adapter Local Area Connection <#>, etc.).
    description *string;
    // Last IPv4 address associated with this NIC.
    ipV4Address *string;
    // Last Public (aka global) IPv6 address associated with this NIC.
    ipV6Address *string;
    // Last local (link-local or site-local) IPv6 address associated with this NIC.
    localIpV6Address *string;
    // MAC address of the NIC on this host.
    macAddress *string;
}
// Instantiates a new networkInterface and sets the default values.
func NewNetworkInterface()(*NetworkInterface) {
    m := &NetworkInterface{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkInterface) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the description property value. Description of the NIC (e.g. Ethernet adapter, Wireless LAN adapter Local Area Connection <#>, etc.).
func (m *NetworkInterface) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the ipV4Address property value. Last IPv4 address associated with this NIC.
func (m *NetworkInterface) GetIpV4Address()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipV4Address
    }
}
// Gets the ipV6Address property value. Last Public (aka global) IPv6 address associated with this NIC.
func (m *NetworkInterface) GetIpV6Address()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipV6Address
    }
}
// Gets the localIpV6Address property value. Last local (link-local or site-local) IPv6 address associated with this NIC.
func (m *NetworkInterface) GetLocalIpV6Address()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localIpV6Address
    }
}
// Gets the macAddress property value. MAC address of the NIC on this host.
func (m *NetworkInterface) GetMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.macAddress
    }
}
// The deserialization information for the current model
func (m *NetworkInterface) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["ipV4Address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpV4Address(val)
        }
        return nil
    }
    res["ipV6Address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpV6Address(val)
        }
        return nil
    }
    res["localIpV6Address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalIpV6Address(val)
        }
        return nil
    }
    res["macAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacAddress(val)
        }
        return nil
    }
    return res
}
func (m *NetworkInterface) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *NetworkInterface) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *NetworkInterface) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the description property value. Description of the NIC (e.g. Ethernet adapter, Wireless LAN adapter Local Area Connection <#>, etc.).
// Parameters:
//  - value : Value to set for the description property.
func (m *NetworkInterface) SetDescription(value *string)() {
    m.description = value
}
// Sets the ipV4Address property value. Last IPv4 address associated with this NIC.
// Parameters:
//  - value : Value to set for the ipV4Address property.
func (m *NetworkInterface) SetIpV4Address(value *string)() {
    m.ipV4Address = value
}
// Sets the ipV6Address property value. Last Public (aka global) IPv6 address associated with this NIC.
// Parameters:
//  - value : Value to set for the ipV6Address property.
func (m *NetworkInterface) SetIpV6Address(value *string)() {
    m.ipV6Address = value
}
// Sets the localIpV6Address property value. Last local (link-local or site-local) IPv6 address associated with this NIC.
// Parameters:
//  - value : Value to set for the localIpV6Address property.
func (m *NetworkInterface) SetLocalIpV6Address(value *string)() {
    m.localIpV6Address = value
}
// Sets the macAddress property value. MAC address of the NIC on this host.
// Parameters:
//  - value : Value to set for the macAddress property.
func (m *NetworkInterface) SetMacAddress(value *string)() {
    m.macAddress = value
}
