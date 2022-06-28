package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XVpnConfiguration 
type Windows10XVpnConfiguration struct {
    DeviceManagementResourceAccessProfileBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // ID to the Authentication Certificate
    authenticationCertificateId *string
    // Custom XML commands that configures the VPN connection. (UTF8 byte encoding)
    customXml []byte
    // Custom Xml file name.
    customXmlFileName *string
}
// NewWindows10XVpnConfiguration instantiates a new Windows10XVpnConfiguration and sets the default values.
func NewWindows10XVpnConfiguration()(*Windows10XVpnConfiguration) {
    m := &Windows10XVpnConfiguration{
        DeviceManagementResourceAccessProfileBase: *NewDeviceManagementResourceAccessProfileBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindows10XVpnConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10XVpnConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10XVpnConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Windows10XVpnConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAuthenticationCertificateId gets the authenticationCertificateId property value. ID to the Authentication Certificate
func (m *Windows10XVpnConfiguration) GetAuthenticationCertificateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationCertificateId
    }
}
// GetCustomXml gets the customXml property value. Custom XML commands that configures the VPN connection. (UTF8 byte encoding)
func (m *Windows10XVpnConfiguration) GetCustomXml()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.customXml
    }
}
// GetCustomXmlFileName gets the customXmlFileName property value. Custom Xml file name.
func (m *Windows10XVpnConfiguration) GetCustomXmlFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customXmlFileName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10XVpnConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementResourceAccessProfileBase.GetFieldDeserializers()
    res["authenticationCertificateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationCertificateId(val)
        }
        return nil
    }
    res["customXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomXml(val)
        }
        return nil
    }
    res["customXmlFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomXmlFileName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Windows10XVpnConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Windows10XVpnConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAuthenticationCertificateId sets the authenticationCertificateId property value. ID to the Authentication Certificate
func (m *Windows10XVpnConfiguration) SetAuthenticationCertificateId(value *string)() {
    if m != nil {
        m.authenticationCertificateId = value
    }
}
// SetCustomXml sets the customXml property value. Custom XML commands that configures the VPN connection. (UTF8 byte encoding)
func (m *Windows10XVpnConfiguration) SetCustomXml(value []byte)() {
    if m != nil {
        m.customXml = value
    }
}
// SetCustomXmlFileName sets the customXmlFileName property value. Custom Xml file name.
func (m *Windows10XVpnConfiguration) SetCustomXmlFileName(value *string)() {
    if m != nil {
        m.customXmlFileName = value
    }
}
