package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsInformationProtectionWipeAction struct {
    Entity
    // Last checkin time of the device that was targeted by this wipe action.
    lastCheckInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Wipe action status. Possible values are: none, pending, canceled, active, done, failed, notSupported.
    status *ActionState;
    // Targeted device Mac address.
    targetedDeviceMacAddress *string;
    // Targeted device name.
    targetedDeviceName *string;
    // The DeviceRegistrationId being targeted by this wipe action.
    targetedDeviceRegistrationId *string;
    // The UserId being targeted by this wipe action.
    targetedUserId *string;
}
// Instantiates a new windowsInformationProtectionWipeAction and sets the default values.
func NewWindowsInformationProtectionWipeAction()(*WindowsInformationProtectionWipeAction) {
    m := &WindowsInformationProtectionWipeAction{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the lastCheckInDateTime property value. Last checkin time of the device that was targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastCheckInDateTime
    }
}
// Gets the status property value. Wipe action status. Possible values are: none, pending, canceled, active, done, failed, notSupported.
func (m *WindowsInformationProtectionWipeAction) GetStatus()(*ActionState) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the targetedDeviceMacAddress property value. Targeted device Mac address.
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetedDeviceMacAddress
    }
}
// Gets the targetedDeviceName property value. Targeted device name.
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetedDeviceName
    }
}
// Gets the targetedDeviceRegistrationId property value. The DeviceRegistrationId being targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceRegistrationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetedDeviceRegistrationId
    }
}
// Gets the targetedUserId property value. The UserId being targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) GetTargetedUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetedUserId
    }
}
// The deserialization information for the current model
func (m *WindowsInformationProtectionWipeAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastCheckInDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastCheckInDateTime(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionState)
        if err != nil {
            return err
        }
        cast := val.(ActionState)
        m.SetStatus(&cast)
        return nil
    }
    res["targetedDeviceMacAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetedDeviceMacAddress(val)
        return nil
    }
    res["targetedDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetedDeviceName(val)
        return nil
    }
    res["targetedDeviceRegistrationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetedDeviceRegistrationId(val)
        return nil
    }
    res["targetedUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetedUserId(val)
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionWipeAction) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsInformationProtectionWipeAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastCheckInDateTime", m.GetLastCheckInDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedDeviceMacAddress", m.GetTargetedDeviceMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedDeviceName", m.GetTargetedDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedDeviceRegistrationId", m.GetTargetedDeviceRegistrationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedUserId", m.GetTargetedUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the lastCheckInDateTime property value. Last checkin time of the device that was targeted by this wipe action.
// Parameters:
//  - value : Value to set for the lastCheckInDateTime property.
func (m *WindowsInformationProtectionWipeAction) SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckInDateTime = value
}
// Sets the status property value. Wipe action status. Possible values are: none, pending, canceled, active, done, failed, notSupported.
// Parameters:
//  - value : Value to set for the status property.
func (m *WindowsInformationProtectionWipeAction) SetStatus(value *ActionState)() {
    m.status = value
}
// Sets the targetedDeviceMacAddress property value. Targeted device Mac address.
// Parameters:
//  - value : Value to set for the targetedDeviceMacAddress property.
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceMacAddress(value *string)() {
    m.targetedDeviceMacAddress = value
}
// Sets the targetedDeviceName property value. Targeted device name.
// Parameters:
//  - value : Value to set for the targetedDeviceName property.
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceName(value *string)() {
    m.targetedDeviceName = value
}
// Sets the targetedDeviceRegistrationId property value. The DeviceRegistrationId being targeted by this wipe action.
// Parameters:
//  - value : Value to set for the targetedDeviceRegistrationId property.
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceRegistrationId(value *string)() {
    m.targetedDeviceRegistrationId = value
}
// Sets the targetedUserId property value. The UserId being targeted by this wipe action.
// Parameters:
//  - value : Value to set for the targetedUserId property.
func (m *WindowsInformationProtectionWipeAction) SetTargetedUserId(value *string)() {
    m.targetedUserId = value
}
