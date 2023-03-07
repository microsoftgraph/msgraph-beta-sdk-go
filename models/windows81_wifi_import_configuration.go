package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81WifiImportConfiguration 
type Windows81WifiImportConfiguration struct {
    DeviceConfiguration
}
// NewWindows81WifiImportConfiguration instantiates a new Windows81WifiImportConfiguration and sets the default values.
func NewWindows81WifiImportConfiguration()(*Windows81WifiImportConfiguration) {
    m := &Windows81WifiImportConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows81WifiImportConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows81WifiImportConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows81WifiImportConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows81WifiImportConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows81WifiImportConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val)
        }
        return nil
    }
    res["payloadFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadFileName(val)
        }
        return nil
    }
    res["profileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileName(val)
        }
        return nil
    }
    return res
}
// GetPayload gets the payload property value. Payload. (UTF8 encoded byte array). This is the XML file saved on the device you used to connect to the Wi-Fi endpoint.
func (m *Windows81WifiImportConfiguration) GetPayload()([]byte) {
    val, err := m.GetBackingStore().Get("payload")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetPayloadFileName gets the payloadFileName property value. Payload file name (.xml).
func (m *Windows81WifiImportConfiguration) GetPayloadFileName()(*string) {
    val, err := m.GetBackingStore().Get("payloadFileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProfileName gets the profileName property value. Profile name displayed in the UI.
func (m *Windows81WifiImportConfiguration) GetProfileName()(*string) {
    val, err := m.GetBackingStore().Get("profileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Windows81WifiImportConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payloadFileName", m.GetPayloadFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profileName", m.GetProfileName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPayload sets the payload property value. Payload. (UTF8 encoded byte array). This is the XML file saved on the device you used to connect to the Wi-Fi endpoint.
func (m *Windows81WifiImportConfiguration) SetPayload(value []byte)() {
    err := m.GetBackingStore().Set("payload", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloadFileName sets the payloadFileName property value. Payload file name (.xml).
func (m *Windows81WifiImportConfiguration) SetPayloadFileName(value *string)() {
    err := m.GetBackingStore().Set("payloadFileName", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileName sets the profileName property value. Profile name displayed in the UI.
func (m *Windows81WifiImportConfiguration) SetProfileName(value *string)() {
    err := m.GetBackingStore().Set("profileName", value)
    if err != nil {
        panic(err)
    }
}
// Windows81WifiImportConfigurationable 
type Windows81WifiImportConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPayload()([]byte)
    GetPayloadFileName()(*string)
    GetProfileName()(*string)
    SetPayload(value []byte)()
    SetPayloadFileName(value *string)()
    SetProfileName(value *string)()
}
