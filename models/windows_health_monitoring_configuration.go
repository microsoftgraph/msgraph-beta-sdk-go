package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsHealthMonitoringConfiguration 
type WindowsHealthMonitoringConfiguration struct {
    DeviceConfiguration
}
// NewWindowsHealthMonitoringConfiguration instantiates a new WindowsHealthMonitoringConfiguration and sets the default values.
func NewWindowsHealthMonitoringConfiguration()(*WindowsHealthMonitoringConfiguration) {
    m := &WindowsHealthMonitoringConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsHealthMonitoringConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsHealthMonitoringConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsHealthMonitoringConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsHealthMonitoringConfiguration(), nil
}
// GetAllowDeviceHealthMonitoring gets the allowDeviceHealthMonitoring property value. Possible values of a property
func (m *WindowsHealthMonitoringConfiguration) GetAllowDeviceHealthMonitoring()(*Enablement) {
    val, err := m.GetBackingStore().Get("allowDeviceHealthMonitoring")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetConfigDeviceHealthMonitoringCustomScope gets the configDeviceHealthMonitoringCustomScope property value. Specifies custom set of events collected from the device where health monitoring is enabled
func (m *WindowsHealthMonitoringConfiguration) GetConfigDeviceHealthMonitoringCustomScope()(*string) {
    val, err := m.GetBackingStore().Get("configDeviceHealthMonitoringCustomScope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConfigDeviceHealthMonitoringScope gets the configDeviceHealthMonitoringScope property value. Device health monitoring scope
func (m *WindowsHealthMonitoringConfiguration) GetConfigDeviceHealthMonitoringScope()(*WindowsHealthMonitoringScope) {
    val, err := m.GetBackingStore().Get("configDeviceHealthMonitoringScope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsHealthMonitoringScope)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsHealthMonitoringConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["allowDeviceHealthMonitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeviceHealthMonitoring(val.(*Enablement))
        }
        return nil
    }
    res["configDeviceHealthMonitoringCustomScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigDeviceHealthMonitoringCustomScope(val)
        }
        return nil
    }
    res["configDeviceHealthMonitoringScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsHealthMonitoringScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigDeviceHealthMonitoringScope(val.(*WindowsHealthMonitoringScope))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsHealthMonitoringConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowDeviceHealthMonitoring() != nil {
        cast := (*m.GetAllowDeviceHealthMonitoring()).String()
        err = writer.WriteStringValue("allowDeviceHealthMonitoring", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("configDeviceHealthMonitoringCustomScope", m.GetConfigDeviceHealthMonitoringCustomScope())
        if err != nil {
            return err
        }
    }
    if m.GetConfigDeviceHealthMonitoringScope() != nil {
        cast := (*m.GetConfigDeviceHealthMonitoringScope()).String()
        err = writer.WriteStringValue("configDeviceHealthMonitoringScope", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowDeviceHealthMonitoring sets the allowDeviceHealthMonitoring property value. Possible values of a property
func (m *WindowsHealthMonitoringConfiguration) SetAllowDeviceHealthMonitoring(value *Enablement)() {
    err := m.GetBackingStore().Set("allowDeviceHealthMonitoring", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigDeviceHealthMonitoringCustomScope sets the configDeviceHealthMonitoringCustomScope property value. Specifies custom set of events collected from the device where health monitoring is enabled
func (m *WindowsHealthMonitoringConfiguration) SetConfigDeviceHealthMonitoringCustomScope(value *string)() {
    err := m.GetBackingStore().Set("configDeviceHealthMonitoringCustomScope", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigDeviceHealthMonitoringScope sets the configDeviceHealthMonitoringScope property value. Device health monitoring scope
func (m *WindowsHealthMonitoringConfiguration) SetConfigDeviceHealthMonitoringScope(value *WindowsHealthMonitoringScope)() {
    err := m.GetBackingStore().Set("configDeviceHealthMonitoringScope", value)
    if err != nil {
        panic(err)
    }
}
// WindowsHealthMonitoringConfigurationable 
type WindowsHealthMonitoringConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDeviceHealthMonitoring()(*Enablement)
    GetConfigDeviceHealthMonitoringCustomScope()(*string)
    GetConfigDeviceHealthMonitoringScope()(*WindowsHealthMonitoringScope)
    SetAllowDeviceHealthMonitoring(value *Enablement)()
    SetConfigDeviceHealthMonitoringCustomScope(value *string)()
    SetConfigDeviceHealthMonitoringScope(value *WindowsHealthMonitoringScope)()
}
