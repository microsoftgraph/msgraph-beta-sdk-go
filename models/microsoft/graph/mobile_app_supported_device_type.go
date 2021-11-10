package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new mobileAppSupportedDeviceType and sets the default values.
func NewMobileAppSupportedDeviceType()(*MobileAppSupportedDeviceType) {
    m := &MobileAppSupportedDeviceType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppSupportedDeviceType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the maximumOperatingSystemVersion property value. Maximum OS version
func (m *MobileAppSupportedDeviceType) GetMaximumOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumOperatingSystemVersion
    }
}
// Gets the minimumOperatingSystemVersion property value. Minimum OS version
func (m *MobileAppSupportedDeviceType) GetMinimumOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumOperatingSystemVersion
    }
}
// Gets the type_escaped property value. Device type. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
func (m *MobileAppSupportedDeviceType) GetType_escaped()(*DeviceType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceType)
            m.SetType_escaped(&cast)
        }
        return nil
    }
    return res
}
func (m *MobileAppSupportedDeviceType) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
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
func (m *MobileAppSupportedDeviceType) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the maximumOperatingSystemVersion property value. Maximum OS version
// Parameters:
//  - value : Value to set for the maximumOperatingSystemVersion property.
func (m *MobileAppSupportedDeviceType) SetMaximumOperatingSystemVersion(value *string)() {
    m.maximumOperatingSystemVersion = value
}
// Sets the minimumOperatingSystemVersion property value. Minimum OS version
// Parameters:
//  - value : Value to set for the minimumOperatingSystemVersion property.
func (m *MobileAppSupportedDeviceType) SetMinimumOperatingSystemVersion(value *string)() {
    m.minimumOperatingSystemVersion = value
}
// Sets the type_escaped property value. Device type. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *MobileAppSupportedDeviceType) SetType_escaped(value *DeviceType)() {
    m.type_escaped = value
}
