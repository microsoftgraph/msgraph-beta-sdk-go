package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsManagementAppHealthState struct {
    Entity
    deviceName *string;
    deviceOSVersion *string;
    healthState *HealthState;
    installedVersion *string;
    lastCheckInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewWindowsManagementAppHealthState()(*WindowsManagementAppHealthState) {
    m := &WindowsManagementAppHealthState{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WindowsManagementAppHealthState) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *WindowsManagementAppHealthState) GetDeviceOSVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceOSVersion
    }
}
func (m *WindowsManagementAppHealthState) GetHealthState()(*HealthState) {
    if m == nil {
        return nil
    } else {
        return m.healthState
    }
}
func (m *WindowsManagementAppHealthState) GetInstalledVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.installedVersion
    }
}
func (m *WindowsManagementAppHealthState) GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastCheckInDateTime
    }
}
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
func (m *WindowsManagementAppHealthState) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *WindowsManagementAppHealthState) SetDeviceOSVersion(value *string)() {
    m.deviceOSVersion = value
}
func (m *WindowsManagementAppHealthState) SetHealthState(value *HealthState)() {
    m.healthState = value
}
func (m *WindowsManagementAppHealthState) SetInstalledVersion(value *string)() {
    m.installedVersion = value
}
func (m *WindowsManagementAppHealthState) SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckInDateTime = value
}
