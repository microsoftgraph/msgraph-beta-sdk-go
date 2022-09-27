package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XWifiConfiguration 
type Windows10XWifiConfiguration struct {
    DeviceManagementResourceAccessProfileBase
    // ID to the Authentication Certificate
    authenticationCertificateId *string
    // Custom XML commands that configures the VPN connection. (UTF8 byte encoding)
    customXml []byte
    // Custom Xml file name.
    customXmlFileName *string
}
// NewWindows10XWifiConfiguration instantiates a new Windows10XWifiConfiguration and sets the default values.
func NewWindows10XWifiConfiguration()(*Windows10XWifiConfiguration) {
    m := &Windows10XWifiConfiguration{
        DeviceManagementResourceAccessProfileBase: *NewDeviceManagementResourceAccessProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.windows10XWifiConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10XWifiConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10XWifiConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10XWifiConfiguration(), nil
}
// GetAuthenticationCertificateId gets the authenticationCertificateId property value. ID to the Authentication Certificate
func (m *Windows10XWifiConfiguration) GetAuthenticationCertificateId()(*string) {
    return m.authenticationCertificateId
}
// GetCustomXml gets the customXml property value. Custom XML commands that configures the VPN connection. (UTF8 byte encoding)
func (m *Windows10XWifiConfiguration) GetCustomXml()([]byte) {
    return m.customXml
}
// GetCustomXmlFileName gets the customXmlFileName property value. Custom Xml file name.
func (m *Windows10XWifiConfiguration) GetCustomXmlFileName()(*string) {
    return m.customXmlFileName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10XWifiConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementResourceAccessProfileBase.GetFieldDeserializers()
    res["authenticationCertificateId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAuthenticationCertificateId)
    res["customXml"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetCustomXml)
    res["customXmlFileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomXmlFileName)
    return res
}
// Serialize serializes information the current object
func (m *Windows10XWifiConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementResourceAccessProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("authenticationCertificateId", m.GetAuthenticationCertificateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("customXml", m.GetCustomXml())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customXmlFileName", m.GetCustomXmlFileName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationCertificateId sets the authenticationCertificateId property value. ID to the Authentication Certificate
func (m *Windows10XWifiConfiguration) SetAuthenticationCertificateId(value *string)() {
    m.authenticationCertificateId = value
}
// SetCustomXml sets the customXml property value. Custom XML commands that configures the VPN connection. (UTF8 byte encoding)
func (m *Windows10XWifiConfiguration) SetCustomXml(value []byte)() {
    m.customXml = value
}
// SetCustomXmlFileName sets the customXmlFileName property value. Custom Xml file name.
func (m *Windows10XWifiConfiguration) SetCustomXmlFileName(value *string)() {
    m.customXmlFileName = value
}
