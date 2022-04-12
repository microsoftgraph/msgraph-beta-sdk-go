package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppSupportedDeviceType device properties
type MobileAppSupportedDeviceType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Maximum OS version
    maximumOperatingSystemVersion *string
    // Minimum OS version
    minimumOperatingSystemVersion *string
    // Device type. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
    type_escaped *DeviceType
}
// NewMobileAppSupportedDeviceType instantiates a new mobileAppSupportedDeviceType and sets the default values.
func NewMobileAppSupportedDeviceType()(*MobileAppSupportedDeviceType) {
    m := &MobileAppSupportedDeviceType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMobileAppSupportedDeviceTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppSupportedDeviceTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppSupportedDeviceType(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppSupportedDeviceType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppSupportedDeviceType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maximumOperatingSystemVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumOperatingSystemVersion(val)
        }
        return nil
    }
    res["minimumOperatingSystemVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumOperatingSystemVersion(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val.(*DeviceType))
        }
        return nil
    }
    return res
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
// Serialize serializes information the current object
func (m *MobileAppSupportedDeviceType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetType_escaped()).String()
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
