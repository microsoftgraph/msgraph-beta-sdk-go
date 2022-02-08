package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementIntentDeviceState 
type DeviceManagementIntentDeviceState struct {
    Entity
    // Device name that is being reported
    deviceDisplayName *string;
    // Device id that is being reported
    deviceId *string;
    // Last modified date time of an intent report
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Device state for an intent. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
    state *ComplianceStatus;
    // The user name that is being reported on a device
    userName *string;
    // The user principal name that is being reported on a device
    userPrincipalName *string;
}
// NewDeviceManagementIntentDeviceState instantiates a new deviceManagementIntentDeviceState and sets the default values.
func NewDeviceManagementIntentDeviceState()(*DeviceManagementIntentDeviceState) {
    m := &DeviceManagementIntentDeviceState{
        Entity: *NewEntity(),
    }
    return m
}
// GetDeviceDisplayName gets the deviceDisplayName property value. Device name that is being reported
func (m *DeviceManagementIntentDeviceState) GetDeviceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDisplayName
    }
}
// GetDeviceId gets the deviceId property value. Device id that is being reported
func (m *DeviceManagementIntentDeviceState) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. Last modified date time of an intent report
func (m *DeviceManagementIntentDeviceState) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportedDateTime
    }
}
// GetState gets the state property value. Device state for an intent. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *DeviceManagementIntentDeviceState) GetState()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetUserName gets the userName property value. The user name that is being reported on a device
func (m *DeviceManagementIntentDeviceState) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name that is being reported on a device
func (m *DeviceManagementIntentDeviceState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentDeviceState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDisplayName(val)
        }
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["lastReportedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportedDateTime(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ComplianceStatus))
        }
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementIntentDeviceState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentDeviceState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceDisplayName", m.GetDeviceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastReportedDateTime", m.GetLastReportedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceDisplayName sets the deviceDisplayName property value. Device name that is being reported
func (m *DeviceManagementIntentDeviceState) SetDeviceDisplayName(value *string)() {
    if m != nil {
        m.deviceDisplayName = value
    }
}
// SetDeviceId sets the deviceId property value. Device id that is being reported
func (m *DeviceManagementIntentDeviceState) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. Last modified date time of an intent report
func (m *DeviceManagementIntentDeviceState) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastReportedDateTime = value
    }
}
// SetState sets the state property value. Device state for an intent. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *DeviceManagementIntentDeviceState) SetState(value *ComplianceStatus)() {
    if m != nil {
        m.state = value
    }
}
// SetUserName sets the userName property value. The user name that is being reported on a device
func (m *DeviceManagementIntentDeviceState) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name that is being reported on a device
func (m *DeviceManagementIntentDeviceState) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
