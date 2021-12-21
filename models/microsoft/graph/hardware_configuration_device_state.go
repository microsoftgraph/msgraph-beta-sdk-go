package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HardwareConfigurationDeviceState 
type HardwareConfigurationDeviceState struct {
    Entity
    // Error from the hardware configuration execution
    configurationError *string;
    // Output of the hardware configuration execution
    configurationOutput *string;
    // Configuration state from the lastest hardware configuration execution. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
    configurationState *RunState;
    // The name of the device
    deviceName *string;
    // The Policy internal version
    internalVersion *int32;
    // The last timestamp of when the hardware configuration executed
    lastStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Operating system version of the device.
    osVersion *string;
    // User Principal Name (UPN).
    upn *string;
}
// NewHardwareConfigurationDeviceState instantiates a new hardwareConfigurationDeviceState and sets the default values.
func NewHardwareConfigurationDeviceState()(*HardwareConfigurationDeviceState) {
    m := &HardwareConfigurationDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// GetConfigurationError gets the configurationError property value. Error from the hardware configuration execution
func (m *HardwareConfigurationDeviceState) GetConfigurationError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.configurationError
    }
}
// GetConfigurationOutput gets the configurationOutput property value. Output of the hardware configuration execution
func (m *HardwareConfigurationDeviceState) GetConfigurationOutput()(*string) {
    if m == nil {
        return nil
    } else {
        return m.configurationOutput
    }
}
// GetConfigurationState gets the configurationState property value. Configuration state from the lastest hardware configuration execution. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
func (m *HardwareConfigurationDeviceState) GetConfigurationState()(*RunState) {
    if m == nil {
        return nil
    } else {
        return m.configurationState
    }
}
// GetDeviceName gets the deviceName property value. The name of the device
func (m *HardwareConfigurationDeviceState) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetInternalVersion gets the internalVersion property value. The Policy internal version
func (m *HardwareConfigurationDeviceState) GetInternalVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.internalVersion
    }
}
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. The last timestamp of when the hardware configuration executed
func (m *HardwareConfigurationDeviceState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastStateUpdateDateTime
    }
}
// GetOsVersion gets the osVersion property value. Operating system version of the device.
func (m *HardwareConfigurationDeviceState) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetUpn gets the upn property value. User Principal Name (UPN).
func (m *HardwareConfigurationDeviceState) GetUpn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upn
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HardwareConfigurationDeviceState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationError(val)
        }
        return nil
    }
    res["configurationOutput"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationOutput(val)
        }
        return nil
    }
    res["configurationState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RunState)
            m.SetConfigurationState(&cast)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["internalVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalVersion(val)
        }
        return nil
    }
    res["lastStateUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastStateUpdateDateTime(val)
        }
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["upn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpn(val)
        }
        return nil
    }
    return res
}
func (m *HardwareConfigurationDeviceState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *HardwareConfigurationDeviceState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("configurationError", m.GetConfigurationError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("configurationOutput", m.GetConfigurationOutput())
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationState() != nil {
        cast := m.GetConfigurationState().String()
        err = writer.WriteStringValue("configurationState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("internalVersion", m.GetInternalVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastStateUpdateDateTime", m.GetLastStateUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("upn", m.GetUpn())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationError sets the configurationError property value. Error from the hardware configuration execution
func (m *HardwareConfigurationDeviceState) SetConfigurationError(value *string)() {
    if m != nil {
        m.configurationError = value
    }
}
// SetConfigurationOutput sets the configurationOutput property value. Output of the hardware configuration execution
func (m *HardwareConfigurationDeviceState) SetConfigurationOutput(value *string)() {
    if m != nil {
        m.configurationOutput = value
    }
}
// SetConfigurationState sets the configurationState property value. Configuration state from the lastest hardware configuration execution. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
func (m *HardwareConfigurationDeviceState) SetConfigurationState(value *RunState)() {
    if m != nil {
        m.configurationState = value
    }
}
// SetDeviceName sets the deviceName property value. The name of the device
func (m *HardwareConfigurationDeviceState) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetInternalVersion sets the internalVersion property value. The Policy internal version
func (m *HardwareConfigurationDeviceState) SetInternalVersion(value *int32)() {
    if m != nil {
        m.internalVersion = value
    }
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. The last timestamp of when the hardware configuration executed
func (m *HardwareConfigurationDeviceState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastStateUpdateDateTime = value
    }
}
// SetOsVersion sets the osVersion property value. Operating system version of the device.
func (m *HardwareConfigurationDeviceState) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetUpn sets the upn property value. User Principal Name (UPN).
func (m *HardwareConfigurationDeviceState) SetUpn(value *string)() {
    if m != nil {
        m.upn = value
    }
}
