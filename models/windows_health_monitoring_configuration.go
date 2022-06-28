package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsHealthMonitoringConfiguration 
type WindowsHealthMonitoringConfiguration struct {
    DeviceConfiguration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Enables device health monitoring on the device. Possible values are: notConfigured, enabled, disabled.
    allowDeviceHealthMonitoring *Enablement
    // Specifies custom set of events collected from the device where health monitoring is enabled
    configDeviceHealthMonitoringCustomScope *string
    // Specifies set of events collected from the device where health monitoring is enabled. Possible values are: undefined, healthMonitoring, bootPerformance, windowsUpdates.
    configDeviceHealthMonitoringScope *WindowsHealthMonitoringScope
}
// NewWindowsHealthMonitoringConfiguration instantiates a new WindowsHealthMonitoringConfiguration and sets the default values.
func NewWindowsHealthMonitoringConfiguration()(*WindowsHealthMonitoringConfiguration) {
    m := &WindowsHealthMonitoringConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsHealthMonitoringConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsHealthMonitoringConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsHealthMonitoringConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsHealthMonitoringConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowDeviceHealthMonitoring gets the allowDeviceHealthMonitoring property value. Enables device health monitoring on the device. Possible values are: notConfigured, enabled, disabled.
func (m *WindowsHealthMonitoringConfiguration) GetAllowDeviceHealthMonitoring()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.allowDeviceHealthMonitoring
    }
}
// GetConfigDeviceHealthMonitoringCustomScope gets the configDeviceHealthMonitoringCustomScope property value. Specifies custom set of events collected from the device where health monitoring is enabled
func (m *WindowsHealthMonitoringConfiguration) GetConfigDeviceHealthMonitoringCustomScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.configDeviceHealthMonitoringCustomScope
    }
}
// GetConfigDeviceHealthMonitoringScope gets the configDeviceHealthMonitoringScope property value. Specifies set of events collected from the device where health monitoring is enabled. Possible values are: undefined, healthMonitoring, bootPerformance, windowsUpdates.
func (m *WindowsHealthMonitoringConfiguration) GetConfigDeviceHealthMonitoringScope()(*WindowsHealthMonitoringScope) {
    if m == nil {
        return nil
    } else {
        return m.configDeviceHealthMonitoringScope
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsHealthMonitoringConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowDeviceHealthMonitoring sets the allowDeviceHealthMonitoring property value. Enables device health monitoring on the device. Possible values are: notConfigured, enabled, disabled.
func (m *WindowsHealthMonitoringConfiguration) SetAllowDeviceHealthMonitoring(value *Enablement)() {
    if m != nil {
        m.allowDeviceHealthMonitoring = value
    }
}
// SetConfigDeviceHealthMonitoringCustomScope sets the configDeviceHealthMonitoringCustomScope property value. Specifies custom set of events collected from the device where health monitoring is enabled
func (m *WindowsHealthMonitoringConfiguration) SetConfigDeviceHealthMonitoringCustomScope(value *string)() {
    if m != nil {
        m.configDeviceHealthMonitoringCustomScope = value
    }
}
// SetConfigDeviceHealthMonitoringScope sets the configDeviceHealthMonitoringScope property value. Specifies set of events collected from the device where health monitoring is enabled. Possible values are: undefined, healthMonitoring, bootPerformance, windowsUpdates.
func (m *WindowsHealthMonitoringConfiguration) SetConfigDeviceHealthMonitoringScope(value *WindowsHealthMonitoringScope)() {
    if m != nil {
        m.configDeviceHealthMonitoringScope = value
    }
}
