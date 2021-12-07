package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// NetworkInterface 
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
// NewNetworkInterface instantiates a new networkInterface and sets the default values.
func NewNetworkInterface()(*NetworkInterface) {
    m := &NetworkInterface{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkInterface) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. Description of the NIC (e.g. Ethernet adapter, Wireless LAN adapter Local Area Connection <#>, etc.).
func (m *NetworkInterface) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetIpV4Address gets the ipV4Address property value. Last IPv4 address associated with this NIC.
func (m *NetworkInterface) GetIpV4Address()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipV4Address
    }
}
// GetIpV6Address gets the ipV6Address property value. Last Public (aka global) IPv6 address associated with this NIC.
func (m *NetworkInterface) GetIpV6Address()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipV6Address
    }
}
// GetLocalIpV6Address gets the localIpV6Address property value. Last local (link-local or site-local) IPv6 address associated with this NIC.
func (m *NetworkInterface) GetLocalIpV6Address()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localIpV6Address
    }
}
// GetMacAddress gets the macAddress property value. MAC address of the NIC on this host.
func (m *NetworkInterface) GetMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.macAddress
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkInterface) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. Description of the NIC (e.g. Ethernet adapter, Wireless LAN adapter Local Area Connection <#>, etc.).
func (m *NetworkInterface) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetIpV4Address sets the ipV4Address property value. Last IPv4 address associated with this NIC.
func (m *NetworkInterface) SetIpV4Address(value *string)() {
    if m != nil {
        m.ipV4Address = value
    }
}
// SetIpV6Address sets the ipV6Address property value. Last Public (aka global) IPv6 address associated with this NIC.
func (m *NetworkInterface) SetIpV6Address(value *string)() {
    if m != nil {
        m.ipV6Address = value
    }
}
// SetLocalIpV6Address sets the localIpV6Address property value. Last local (link-local or site-local) IPv6 address associated with this NIC.
func (m *NetworkInterface) SetLocalIpV6Address(value *string)() {
    if m != nil {
        m.localIpV6Address = value
    }
}
// SetMacAddress sets the macAddress property value. MAC address of the NIC on this host.
func (m *NetworkInterface) SetMacAddress(value *string)() {
    if m != nil {
        m.macAddress = value
    }
}
