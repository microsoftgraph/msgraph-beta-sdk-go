package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSCustomAppConfiguration 
type MacOSCustomAppConfiguration struct {
    DeviceConfiguration
}
// NewMacOSCustomAppConfiguration instantiates a new MacOSCustomAppConfiguration and sets the default values.
func NewMacOSCustomAppConfiguration()(*MacOSCustomAppConfiguration) {
    m := &MacOSCustomAppConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSCustomAppConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSCustomAppConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSCustomAppConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSCustomAppConfiguration(), nil
}
// GetBundleId gets the bundleId property value. Bundle id for targeting.
func (m *MacOSCustomAppConfiguration) GetBundleId()(*string) {
    val, err := m.GetBackingStore().Get("bundleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConfigurationXml gets the configurationXml property value. Configuration xml. (UTF8 encoded byte array)
func (m *MacOSCustomAppConfiguration) GetConfigurationXml()([]byte) {
    val, err := m.GetBackingStore().Get("configurationXml")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSCustomAppConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["bundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleId(val)
        }
        return nil
    }
    res["configurationXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationXml(val)
        }
        return nil
    }
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. Configuration file name (.plist
func (m *MacOSCustomAppConfiguration) GetFileName()(*string) {
    val, err := m.GetBackingStore().Get("fileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSCustomAppConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("bundleId", m.GetBundleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("configurationXml", m.GetConfigurationXml())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBundleId sets the bundleId property value. Bundle id for targeting.
func (m *MacOSCustomAppConfiguration) SetBundleId(value *string)() {
    err := m.GetBackingStore().Set("bundleId", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationXml sets the configurationXml property value. Configuration xml. (UTF8 encoded byte array)
func (m *MacOSCustomAppConfiguration) SetConfigurationXml(value []byte)() {
    err := m.GetBackingStore().Set("configurationXml", value)
    if err != nil {
        panic(err)
    }
}
// SetFileName sets the fileName property value. Configuration file name (.plist
func (m *MacOSCustomAppConfiguration) SetFileName(value *string)() {
    err := m.GetBackingStore().Set("fileName", value)
    if err != nil {
        panic(err)
    }
}
// MacOSCustomAppConfigurationable 
type MacOSCustomAppConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBundleId()(*string)
    GetConfigurationXml()([]byte)
    GetFileName()(*string)
    SetBundleId(value *string)()
    SetConfigurationXml(value []byte)()
    SetFileName(value *string)()
}
