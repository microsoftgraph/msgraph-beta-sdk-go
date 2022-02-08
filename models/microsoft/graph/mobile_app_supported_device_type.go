package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppSupportedDeviceType 
type MobileAppSupportedDeviceType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Maximum OS version
    maximumOperatingSystemVersion *string;
    // Minimum OS version
    minimumOperatingSystemVersion *string;
    // Device type. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
    type_escaped *DeviceType;
}
// NewMobileAppSupportedDeviceType instantiates a new mobileAppSupportedDeviceType and sets the default values.
func NewMobileAppSupportedDeviceType()(*MobileAppSupportedDeviceType) {
    m := &MobileAppSupportedDeviceType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppSupportedDeviceType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMaximumOperatingSystemVersion gets the maximumOperatingSystemVersion property value. Maximum OS version
func (m *MobileAppSupportedDeviceType) GetMaximumOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumOperatingSystemVersion
    }
}
// GetMinimumOperatingSystemVersion gets the minimumOperatingSystemVersion property value. Minimum OS version
func (m *MobileAppSupportedDeviceType) GetMinimumOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumOperatingSystemVersion
    }
}
// GetType gets the type property value. Device type. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
func (m *MobileAppSupportedDeviceType) GetType()(*DeviceType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppSupportedDeviceType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["maximumOperatingSystemVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumOperatingSystemVersion(val)
        }
        return nil
    }
    res["minimumOperatingSystemVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumOperatingSystemVersion(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*DeviceType))
        }
        return nil
    }
    return res
}
func (m *MobileAppSupportedDeviceType) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *MobileAppSupportedDeviceType) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaximumOperatingSystemVersion sets the maximumOperatingSystemVersion property value. Maximum OS version
func (m *MobileAppSupportedDeviceType) SetMaximumOperatingSystemVersion(value *string)() {
    if m != nil {
        m.maximumOperatingSystemVersion = value
    }
}
// SetMinimumOperatingSystemVersion sets the minimumOperatingSystemVersion property value. Minimum OS version
func (m *MobileAppSupportedDeviceType) SetMinimumOperatingSystemVersion(value *string)() {
    if m != nil {
        m.minimumOperatingSystemVersion = value
    }
}
// SetType sets the type property value. Device type. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
func (m *MobileAppSupportedDeviceType) SetType(value *DeviceType)() {
    if m != nil {
        m.type_escaped = value
    }
}
