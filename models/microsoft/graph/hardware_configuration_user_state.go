package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HardwareConfigurationUserState 
type HardwareConfigurationUserState struct {
    Entity
    // Error device count for specific user.
    errorDeviceCount *int32;
    // Failed device count for specific user
    failedDeviceCount *int32;
    // Last timestamp when the hardware configuration executed
    lastStateUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Not applicable device count for specific user.
    notApplicableDeviceCount *int32;
    // Pending device count for specific user.
    pendingDeviceCount *int32;
    // Success device count for specific user
    successfulDeviceCount *int32;
    // Unknown device count for specific user.
    unknownDeviceCount *int32;
    // User Principal Name (UPN)
    upn *string;
    // User Email address
    userEmail *string;
    // User name
    userName *string;
}
// NewHardwareConfigurationUserState instantiates a new hardwareConfigurationUserState and sets the default values.
func NewHardwareConfigurationUserState()(*HardwareConfigurationUserState) {
    m := &HardwareConfigurationUserState{
        Entity: *NewEntity(),
    }
    return m
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Error device count for specific user.
func (m *HardwareConfigurationUserState) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// GetFailedDeviceCount gets the failedDeviceCount property value. Failed device count for specific user
func (m *HardwareConfigurationUserState) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// GetLastStateUpdateDateTime gets the lastStateUpdateDateTime property value. Last timestamp when the hardware configuration executed
func (m *HardwareConfigurationUserState) GetLastStateUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastStateUpdateDateTime
    }
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Not applicable device count for specific user.
func (m *HardwareConfigurationUserState) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// GetPendingDeviceCount gets the pendingDeviceCount property value. Pending device count for specific user.
func (m *HardwareConfigurationUserState) GetPendingDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingDeviceCount
    }
}
// GetSuccessfulDeviceCount gets the successfulDeviceCount property value. Success device count for specific user
func (m *HardwareConfigurationUserState) GetSuccessfulDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successfulDeviceCount
    }
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Unknown device count for specific user.
func (m *HardwareConfigurationUserState) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// GetUpn gets the upn property value. User Principal Name (UPN)
func (m *HardwareConfigurationUserState) GetUpn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upn
    }
}
// GetUserEmail gets the userEmail property value. User Email address
func (m *HardwareConfigurationUserState) GetUserEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userEmail
    }
}
// GetUserName gets the userName property value. User name
func (m *HardwareConfigurationUserState) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HardwareConfigurationUserState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["errorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDeviceCount(val)
        }
        return nil
    }
    res["failedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedDeviceCount(val)
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
    res["notApplicableDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableDeviceCount(val)
        }
        return nil
    }
    res["pendingDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingDeviceCount(val)
        }
        return nil
    }
    res["successfulDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulDeviceCount(val)
        }
        return nil
    }
    res["unknownDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownDeviceCount(val)
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
    res["userEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmail(val)
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
    return res
}
func (m *HardwareConfigurationUserState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *HardwareConfigurationUserState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
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
        err = writer.WriteInt32Value("notApplicableDeviceCount", m.GetNotApplicableDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pendingDeviceCount", m.GetPendingDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successfulDeviceCount", m.GetSuccessfulDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unknownDeviceCount", m.GetUnknownDeviceCount())
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
    {
        err = writer.WriteStringValue("userEmail", m.GetUserEmail())
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
    return nil
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Error device count for specific user.
func (m *HardwareConfigurationUserState) SetErrorDeviceCount(value *int32)() {
    if m != nil {
        m.errorDeviceCount = value
    }
}
// SetFailedDeviceCount sets the failedDeviceCount property value. Failed device count for specific user
func (m *HardwareConfigurationUserState) SetFailedDeviceCount(value *int32)() {
    if m != nil {
        m.failedDeviceCount = value
    }
}
// SetLastStateUpdateDateTime sets the lastStateUpdateDateTime property value. Last timestamp when the hardware configuration executed
func (m *HardwareConfigurationUserState) SetLastStateUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastStateUpdateDateTime = value
    }
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Not applicable device count for specific user.
func (m *HardwareConfigurationUserState) SetNotApplicableDeviceCount(value *int32)() {
    if m != nil {
        m.notApplicableDeviceCount = value
    }
}
// SetPendingDeviceCount sets the pendingDeviceCount property value. Pending device count for specific user.
func (m *HardwareConfigurationUserState) SetPendingDeviceCount(value *int32)() {
    if m != nil {
        m.pendingDeviceCount = value
    }
}
// SetSuccessfulDeviceCount sets the successfulDeviceCount property value. Success device count for specific user
func (m *HardwareConfigurationUserState) SetSuccessfulDeviceCount(value *int32)() {
    if m != nil {
        m.successfulDeviceCount = value
    }
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Unknown device count for specific user.
func (m *HardwareConfigurationUserState) SetUnknownDeviceCount(value *int32)() {
    if m != nil {
        m.unknownDeviceCount = value
    }
}
// SetUpn sets the upn property value. User Principal Name (UPN)
func (m *HardwareConfigurationUserState) SetUpn(value *string)() {
    if m != nil {
        m.upn = value
    }
}
// SetUserEmail sets the userEmail property value. User Email address
func (m *HardwareConfigurationUserState) SetUserEmail(value *string)() {
    if m != nil {
        m.userEmail = value
    }
}
// SetUserName sets the userName property value. User name
func (m *HardwareConfigurationUserState) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
