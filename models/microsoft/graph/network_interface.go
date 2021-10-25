package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type NetworkInterface struct {
    additionalData map[string]interface{};
    description *string;
    ipV4Address *string;
    ipV6Address *string;
    localIpV6Address *string;
    macAddress *string;
}
func NewNetworkInterface()(*NetworkInterface) {
    m := &NetworkInterface{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *NetworkInterface) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *NetworkInterface) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *NetworkInterface) GetIpV4Address()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipV4Address
    }
}
func (m *NetworkInterface) GetIpV6Address()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipV6Address
    }
}
func (m *NetworkInterface) GetLocalIpV6Address()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localIpV6Address
    }
}
func (m *NetworkInterface) GetMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.macAddress
    }
}
func (m *NetworkInterface) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["ipV4Address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIpV4Address(val)
        return nil
    }
    res["ipV6Address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIpV6Address(val)
        return nil
    }
    res["localIpV6Address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLocalIpV6Address(val)
        return nil
    }
    res["macAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMacAddress(val)
        return nil
    }
    return res
}
func (m *NetworkInterface) IsNil()(bool) {
    return m == nil
}
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
func (m *NetworkInterface) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *NetworkInterface) SetDescription(value *string)() {
    m.description = value
}
func (m *NetworkInterface) SetIpV4Address(value *string)() {
    m.ipV4Address = value
}
func (m *NetworkInterface) SetIpV6Address(value *string)() {
    m.ipV6Address = value
}
func (m *NetworkInterface) SetLocalIpV6Address(value *string)() {
    m.localIpV6Address = value
}
func (m *NetworkInterface) SetMacAddress(value *string)() {
    m.macAddress = value
}
