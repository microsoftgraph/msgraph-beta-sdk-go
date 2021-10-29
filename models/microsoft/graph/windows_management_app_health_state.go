package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsManagementAppHealthState struct {
    Entity
    // Name of the device on which Windows management app is installed.
    deviceName *string;
    // Windows 10 OS version of the device on which Windows management app is installed.
    deviceOSVersion *string;
    // Windows management app health state. Possible values are: unknown, healthy, unhealthy.
    healthState *HealthState;
    // Windows management app installed version.
    installedVersion *string;
    // Windows management app last check-in time.
    lastCheckInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new windowsManagementAppHealthState and sets the default values.
func NewWindowsManagementAppHealthState()(*WindowsManagementAppHealthState) {
    m := &WindowsManagementAppHealthState{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceName property value. Name of the device on which Windows management app is installed.
func (m *WindowsManagementAppHealthState) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the deviceOSVersion property value. Windows 10 OS version of the device on which Windows management app is installed.
func (m *WindowsManagementAppHealthState) GetDeviceOSVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceOSVersion
    }
}
// Gets the healthState property value. Windows management app health state. Possible values are: unknown, healthy, unhealthy.
func (m *WindowsManagementAppHealthState) GetHealthState()(*HealthState) {
    if m == nil {
        return nil
    } else {
        return m.healthState
    }
}
// Gets the installedVersion property value. Windows management app installed version.
func (m *WindowsManagementAppHealthState) GetInstalledVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.installedVersion
    }
}
// Gets the lastCheckInDateTime property value. Windows management app last check-in time.
func (m *WindowsManagementAppHealthState) GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastCheckInDateTime
    }
}
// The deserialization information for the current model
func (m *WindowsManagementAppHealthState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["deviceOSVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceOSVersion(val)
        return nil
    }
    res["healthState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseHealthState)
        if err != nil {
            return err
        }
        cast := val.(HealthState)
        m.SetHealthState(&cast)
        return nil
    }
    res["installedVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInstalledVersion(val)
        return nil
    }
    res["lastCheckInDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastCheckInDateTime(val)
        return nil
    }
    return res
}
func (m *WindowsManagementAppHealthState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsManagementAppHealthState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceOSVersion", m.GetDeviceOSVersion())
        if err != nil {
            return err
        }
    }
    if m.GetHealthState() != nil {
        cast := m.GetHealthState().String()
        err = writer.WriteStringValue("healthState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("installedVersion", m.GetInstalledVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastCheckInDateTime", m.GetLastCheckInDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceName property value. Name of the device on which Windows management app is installed.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *WindowsManagementAppHealthState) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the deviceOSVersion property value. Windows 10 OS version of the device on which Windows management app is installed.
// Parameters:
//  - value : Value to set for the deviceOSVersion property.
func (m *WindowsManagementAppHealthState) SetDeviceOSVersion(value *string)() {
    m.deviceOSVersion = value
}
// Sets the healthState property value. Windows management app health state. Possible values are: unknown, healthy, unhealthy.
// Parameters:
//  - value : Value to set for the healthState property.
func (m *WindowsManagementAppHealthState) SetHealthState(value *HealthState)() {
    m.healthState = value
}
// Sets the installedVersion property value. Windows management app installed version.
// Parameters:
//  - value : Value to set for the installedVersion property.
func (m *WindowsManagementAppHealthState) SetInstalledVersion(value *string)() {
    m.installedVersion = value
}
// Sets the lastCheckInDateTime property value. Windows management app last check-in time.
// Parameters:
//  - value : Value to set for the lastCheckInDateTime property.
func (m *WindowsManagementAppHealthState) SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckInDateTime = value
}
