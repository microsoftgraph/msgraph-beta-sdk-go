package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsInformationProtectionWipeAction struct {
    Entity
    lastCheckInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    status *ActionState;
    targetedDeviceMacAddress *string;
    targetedDeviceName *string;
    targetedDeviceRegistrationId *string;
    targetedUserId *string;
}
func NewWindowsInformationProtectionWipeAction()(*WindowsInformationProtectionWipeAction) {
    m := &WindowsInformationProtectionWipeAction{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WindowsInformationProtectionWipeAction) GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastCheckInDateTime
    }
}
func (m *WindowsInformationProtectionWipeAction) GetStatus()(*ActionState) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetedDeviceMacAddress
    }
}
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetedDeviceName
    }
}
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceRegistrationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetedDeviceRegistrationId
    }
}
func (m *WindowsInformationProtectionWipeAction) GetTargetedUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetedUserId
    }
}
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
func (m *WindowsInformationProtectionWipeAction) SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckInDateTime = value
}
func (m *WindowsInformationProtectionWipeAction) SetStatus(value *ActionState)() {
    m.status = value
}
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceMacAddress(value *string)() {
    m.targetedDeviceMacAddress = value
}
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceName(value *string)() {
    m.targetedDeviceName = value
}
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceRegistrationId(value *string)() {
    m.targetedDeviceRegistrationId = value
}
func (m *WindowsInformationProtectionWipeAction) SetTargetedUserId(value *string)() {
    m.targetedUserId = value
}
