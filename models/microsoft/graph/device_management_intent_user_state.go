package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementIntentUserState struct {
    Entity
    // Count of Devices that belongs to a user for an intent
    deviceCount *int32;
    // Last modified date time of an intent report
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // User state for an intent. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
    state *ComplianceStatus;
    // The user name that is being reported on a device
    userName *string;
    // The user principal name that is being reported on a device
    userPrincipalName *string;
}
// Instantiates a new deviceManagementIntentUserState and sets the default values.
func NewDeviceManagementIntentUserState()(*DeviceManagementIntentUserState) {
    m := &DeviceManagementIntentUserState{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceCount property value. Count of Devices that belongs to a user for an intent
func (m *DeviceManagementIntentUserState) GetDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// Gets the lastReportedDateTime property value. Last modified date time of an intent report
func (m *DeviceManagementIntentUserState) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportedDateTime
    }
}
// Gets the state property value. User state for an intent. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *DeviceManagementIntentUserState) GetState()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the userName property value. The user name that is being reported on a device
func (m *DeviceManagementIntentUserState) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// Gets the userPrincipalName property value. The user principal name that is being reported on a device
func (m *DeviceManagementIntentUserState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *DeviceManagementIntentUserState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceCount(val)
        return nil
    }
    res["lastReportedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastReportedDateTime(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        cast := val.(ComplianceStatus)
        m.SetState(&cast)
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserName(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *DeviceManagementIntentUserState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementIntentUserState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
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
        cast := m.GetState().String()
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
// Sets the deviceCount property value. Count of Devices that belongs to a user for an intent
// Parameters:
//  - value : Value to set for the deviceCount property.
func (m *DeviceManagementIntentUserState) SetDeviceCount(value *int32)() {
    m.deviceCount = value
}
// Sets the lastReportedDateTime property value. Last modified date time of an intent report
// Parameters:
//  - value : Value to set for the lastReportedDateTime property.
func (m *DeviceManagementIntentUserState) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// Sets the state property value. User state for an intent. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
// Parameters:
//  - value : Value to set for the state property.
func (m *DeviceManagementIntentUserState) SetState(value *ComplianceStatus)() {
    m.state = value
}
// Sets the userName property value. The user name that is being reported on a device
// Parameters:
//  - value : Value to set for the userName property.
func (m *DeviceManagementIntentUserState) SetUserName(value *string)() {
    m.userName = value
}
// Sets the userPrincipalName property value. The user principal name that is being reported on a device
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *DeviceManagementIntentUserState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
