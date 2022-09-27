package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsHealthMonitoringConfiguration 
type WindowsHealthMonitoringConfiguration struct {
    DeviceConfiguration
    // Possible values of a property
    allowDeviceHealthMonitoring *Enablement
    // Specifies custom set of events collected from the device where health monitoring is enabled
    configDeviceHealthMonitoringCustomScope *string
    // Device health monitoring scope
    configDeviceHealthMonitoringScope *WindowsHealthMonitoringScope
}
// NewWindowsHealthMonitoringConfiguration instantiates a new WindowsHealthMonitoringConfiguration and sets the default values.
func NewWindowsHealthMonitoringConfiguration()(*WindowsHealthMonitoringConfiguration) {
    m := &WindowsHealthMonitoringConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsHealthMonitoringConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsHealthMonitoringConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsHealthMonitoringConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsHealthMonitoringConfiguration(), nil
}
// GetAllowDeviceHealthMonitoring gets the allowDeviceHealthMonitoring property value. Possible values of a property
func (m *WindowsHealthMonitoringConfiguration) GetAllowDeviceHealthMonitoring()(*Enablement) {
    return m.allowDeviceHealthMonitoring
}
// GetConfigDeviceHealthMonitoringCustomScope gets the configDeviceHealthMonitoringCustomScope property value. Specifies custom set of events collected from the device where health monitoring is enabled
func (m *WindowsHealthMonitoringConfiguration) GetConfigDeviceHealthMonitoringCustomScope()(*string) {
    return m.configDeviceHealthMonitoringCustomScope
}
// GetConfigDeviceHealthMonitoringScope gets the configDeviceHealthMonitoringScope property value. Device health monitoring scope
func (m *WindowsHealthMonitoringConfiguration) GetConfigDeviceHealthMonitoringScope()(*WindowsHealthMonitoringScope) {
    return m.configDeviceHealthMonitoringScope
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsHealthMonitoringConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["allowDeviceHealthMonitoring"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEnablement , m.SetAllowDeviceHealthMonitoring)
    res["configDeviceHealthMonitoringCustomScope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetConfigDeviceHealthMonitoringCustomScope)
    res["configDeviceHealthMonitoringScope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWindowsHealthMonitoringScope , m.SetConfigDeviceHealthMonitoringScope)
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
    m.allowDeviceHealthMonitoring = value
}
// SetConfigDeviceHealthMonitoringCustomScope sets the configDeviceHealthMonitoringCustomScope property value. Specifies custom set of events collected from the device where health monitoring is enabled
func (m *WindowsHealthMonitoringConfiguration) SetConfigDeviceHealthMonitoringCustomScope(value *string)() {
    m.configDeviceHealthMonitoringCustomScope = value
}
// SetConfigDeviceHealthMonitoringScope sets the configDeviceHealthMonitoringScope property value. Device health monitoring scope
func (m *WindowsHealthMonitoringConfiguration) SetConfigDeviceHealthMonitoringScope(value *WindowsHealthMonitoringScope)() {
    m.configDeviceHealthMonitoringScope = value
}
