package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MobileAppSupportedDeviceType struct {
    additionalData map[string]interface{};
    maximumOperatingSystemVersion *string;
    minimumOperatingSystemVersion *string;
    type_escpaped *DeviceType;
}
func NewMobileAppSupportedDeviceType()(*MobileAppSupportedDeviceType) {
    m := &MobileAppSupportedDeviceType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MobileAppSupportedDeviceType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MobileAppSupportedDeviceType) GetMaximumOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumOperatingSystemVersion
    }
}
func (m *MobileAppSupportedDeviceType) GetMinimumOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumOperatingSystemVersion
    }
}
func (m *MobileAppSupportedDeviceType) GetType_escpaped()(*DeviceType) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *MobileAppSupportedDeviceType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["maximumOperatingSystemVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMaximumOperatingSystemVersion(val)
        return nil
    }
    res["minimumOperatingSystemVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumOperatingSystemVersion(val)
        return nil
    }
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        cast := val.(DeviceType)
        m.SetType_escpaped(&cast)
        return nil
    }
    return res
}
func (m *MobileAppSupportedDeviceType) IsNil()(bool) {
    return m == nil
}
func (m *MobileAppSupportedDeviceType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("maximumOperatingSystemVersion", m.GetMaximumOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("minimumOperatingSystemVersion", m.GetMinimumOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    if m.GetType_escpaped() != nil {
        cast := m.GetType_escpaped().String()
        err := writer.WriteStringValue("type_escpaped", &cast)
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
func (m *MobileAppSupportedDeviceType) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MobileAppSupportedDeviceType) SetMaximumOperatingSystemVersion(value *string)() {
    m.maximumOperatingSystemVersion = value
}
func (m *MobileAppSupportedDeviceType) SetMinimumOperatingSystemVersion(value *string)() {
    m.minimumOperatingSystemVersion = value
}
func (m *MobileAppSupportedDeviceType) SetType_escpaped(value *DeviceType)() {
    m.type_escpaped = value
}
