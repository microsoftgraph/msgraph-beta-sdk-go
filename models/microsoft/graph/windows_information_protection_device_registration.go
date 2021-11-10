package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsInformationProtectionDeviceRegistration struct {
    Entity
    // Device Mac address.
    deviceMacAddress *string;
    // Device name.
    deviceName *string;
    // Device identifier for this device registration record.
    deviceRegistrationId *string;
    // Device type, for example, Windows laptop VS Windows phone.
    deviceType *string;
    // Last checkin time of the device.
    lastCheckInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // UserId associated with this device registration record.
    userId *string;
}
// Instantiates a new windowsInformationProtectionDeviceRegistration and sets the default values.
func NewWindowsInformationProtectionDeviceRegistration()(*WindowsInformationProtectionDeviceRegistration) {
    m := &WindowsInformationProtectionDeviceRegistration{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceMacAddress property value. Device Mac address.
func (m *WindowsInformationProtectionDeviceRegistration) GetDeviceMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceMacAddress
    }
}
// Gets the deviceName property value. Device name.
func (m *WindowsInformationProtectionDeviceRegistration) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the deviceRegistrationId property value. Device identifier for this device registration record.
func (m *WindowsInformationProtectionDeviceRegistration) GetDeviceRegistrationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceRegistrationId
    }
}
// Gets the deviceType property value. Device type, for example, Windows laptop VS Windows phone.
func (m *WindowsInformationProtectionDeviceRegistration) GetDeviceType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
// Gets the lastCheckInDateTime property value. Last checkin time of the device.
func (m *WindowsInformationProtectionDeviceRegistration) GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastCheckInDateTime
    }
}
// Gets the userId property value. UserId associated with this device registration record.
func (m *WindowsInformationProtectionDeviceRegistration) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *WindowsInformationProtectionDeviceRegistration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceMacAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceMacAddress(val)
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
    res["deviceRegistrationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRegistrationId(val)
        }
        return nil
    }
    res["deviceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val)
        }
        return nil
    }
    res["lastCheckInDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckInDateTime(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionDeviceRegistration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsInformationProtectionDeviceRegistration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceMacAddress", m.GetDeviceMacAddress())
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
        err = writer.WriteStringValue("deviceRegistrationId", m.GetDeviceRegistrationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceType", m.GetDeviceType())
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
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceMacAddress property value. Device Mac address.
// Parameters:
//  - value : Value to set for the deviceMacAddress property.
func (m *WindowsInformationProtectionDeviceRegistration) SetDeviceMacAddress(value *string)() {
    m.deviceMacAddress = value
}
// Sets the deviceName property value. Device name.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *WindowsInformationProtectionDeviceRegistration) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the deviceRegistrationId property value. Device identifier for this device registration record.
// Parameters:
//  - value : Value to set for the deviceRegistrationId property.
func (m *WindowsInformationProtectionDeviceRegistration) SetDeviceRegistrationId(value *string)() {
    m.deviceRegistrationId = value
}
// Sets the deviceType property value. Device type, for example, Windows laptop VS Windows phone.
// Parameters:
//  - value : Value to set for the deviceType property.
func (m *WindowsInformationProtectionDeviceRegistration) SetDeviceType(value *string)() {
    m.deviceType = value
}
// Sets the lastCheckInDateTime property value. Last checkin time of the device.
// Parameters:
//  - value : Value to set for the lastCheckInDateTime property.
func (m *WindowsInformationProtectionDeviceRegistration) SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckInDateTime = value
}
// Sets the userId property value. UserId associated with this device registration record.
// Parameters:
//  - value : Value to set for the userId property.
func (m *WindowsInformationProtectionDeviceRegistration) SetUserId(value *string)() {
    m.userId = value
}
