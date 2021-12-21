package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HardwareConfigurationRunSummary 
type HardwareConfigurationRunSummary struct {
    Entity
    // Number of devices for which hardware configuration state is error
    errorDeviceCount *int32;
    // Number of users for which hardware configuration state is error
    errorUserCount *int32;
    // Number of devices for which hardware configuration found an issue
    failedDeviceCount *int32;
    // Number of users for which hardware configuration found an issue
    failedUserCount *int32;
    // Last run time for the configuration across all devices
    lastRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Number of devices for which hardware configuration state is not applicable
    notApplicableDeviceCount *int32;
    // Number of users for which hardware configuration state is not applicable
    notApplicableUserCount *int32;
    // Number of devices for which hardware configuration is in pending state
    pendingDeviceCount *int32;
    // Number of users for which hardware configuration is in pending state
    pendingUserCount *int32;
    // Number of devices for which hardware configured without any issue
    successfulDeviceCount *int32;
    // Number of users for which hardware configured without any issue
    successfulUserCount *int32;
    // Number of devices for which hardware configuration state is unknown
    unknownDeviceCount *int32;
    // Number of users for which hardware configuration state is unknown
    unknownUserCount *int32;
}
// NewHardwareConfigurationRunSummary instantiates a new hardwareConfigurationRunSummary and sets the default values.
func NewHardwareConfigurationRunSummary()(*HardwareConfigurationRunSummary) {
    m := &HardwareConfigurationRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Number of devices for which hardware configuration state is error
func (m *HardwareConfigurationRunSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// GetErrorUserCount gets the errorUserCount property value. Number of users for which hardware configuration state is error
func (m *HardwareConfigurationRunSummary) GetErrorUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorUserCount
    }
}
// GetFailedDeviceCount gets the failedDeviceCount property value. Number of devices for which hardware configuration found an issue
func (m *HardwareConfigurationRunSummary) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// GetFailedUserCount gets the failedUserCount property value. Number of users for which hardware configuration found an issue
func (m *HardwareConfigurationRunSummary) GetFailedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedUserCount
    }
}
// GetLastRunDateTime gets the lastRunDateTime property value. Last run time for the configuration across all devices
func (m *HardwareConfigurationRunSummary) GetLastRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRunDateTime
    }
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Number of devices for which hardware configuration state is not applicable
func (m *HardwareConfigurationRunSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// GetNotApplicableUserCount gets the notApplicableUserCount property value. Number of users for which hardware configuration state is not applicable
func (m *HardwareConfigurationRunSummary) GetNotApplicableUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableUserCount
    }
}
// GetPendingDeviceCount gets the pendingDeviceCount property value. Number of devices for which hardware configuration is in pending state
func (m *HardwareConfigurationRunSummary) GetPendingDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingDeviceCount
    }
}
// GetPendingUserCount gets the pendingUserCount property value. Number of users for which hardware configuration is in pending state
func (m *HardwareConfigurationRunSummary) GetPendingUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingUserCount
    }
}
// GetSuccessfulDeviceCount gets the successfulDeviceCount property value. Number of devices for which hardware configured without any issue
func (m *HardwareConfigurationRunSummary) GetSuccessfulDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successfulDeviceCount
    }
}
// GetSuccessfulUserCount gets the successfulUserCount property value. Number of users for which hardware configured without any issue
func (m *HardwareConfigurationRunSummary) GetSuccessfulUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successfulUserCount
    }
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Number of devices for which hardware configuration state is unknown
func (m *HardwareConfigurationRunSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// GetUnknownUserCount gets the unknownUserCount property value. Number of users for which hardware configuration state is unknown
func (m *HardwareConfigurationRunSummary) GetUnknownUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownUserCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HardwareConfigurationRunSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["errorUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorUserCount(val)
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
    res["failedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedUserCount(val)
        }
        return nil
    }
    res["lastRunDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRunDateTime(val)
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
    res["notApplicableUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableUserCount(val)
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
    res["pendingUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingUserCount(val)
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
    res["successfulUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulUserCount(val)
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
    res["unknownUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownUserCount(val)
        }
        return nil
    }
    return res
}
func (m *HardwareConfigurationRunSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *HardwareConfigurationRunSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("errorUserCount", m.GetErrorUserCount())
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
        err = writer.WriteInt32Value("failedUserCount", m.GetFailedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRunDateTime", m.GetLastRunDateTime())
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
        err = writer.WriteInt32Value("notApplicableUserCount", m.GetNotApplicableUserCount())
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
        err = writer.WriteInt32Value("pendingUserCount", m.GetPendingUserCount())
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
        err = writer.WriteInt32Value("successfulUserCount", m.GetSuccessfulUserCount())
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
        err = writer.WriteInt32Value("unknownUserCount", m.GetUnknownUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Number of devices for which hardware configuration state is error
func (m *HardwareConfigurationRunSummary) SetErrorDeviceCount(value *int32)() {
    if m != nil {
        m.errorDeviceCount = value
    }
}
// SetErrorUserCount sets the errorUserCount property value. Number of users for which hardware configuration state is error
func (m *HardwareConfigurationRunSummary) SetErrorUserCount(value *int32)() {
    if m != nil {
        m.errorUserCount = value
    }
}
// SetFailedDeviceCount sets the failedDeviceCount property value. Number of devices for which hardware configuration found an issue
func (m *HardwareConfigurationRunSummary) SetFailedDeviceCount(value *int32)() {
    if m != nil {
        m.failedDeviceCount = value
    }
}
// SetFailedUserCount sets the failedUserCount property value. Number of users for which hardware configuration found an issue
func (m *HardwareConfigurationRunSummary) SetFailedUserCount(value *int32)() {
    if m != nil {
        m.failedUserCount = value
    }
}
// SetLastRunDateTime sets the lastRunDateTime property value. Last run time for the configuration across all devices
func (m *HardwareConfigurationRunSummary) SetLastRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastRunDateTime = value
    }
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Number of devices for which hardware configuration state is not applicable
func (m *HardwareConfigurationRunSummary) SetNotApplicableDeviceCount(value *int32)() {
    if m != nil {
        m.notApplicableDeviceCount = value
    }
}
// SetNotApplicableUserCount sets the notApplicableUserCount property value. Number of users for which hardware configuration state is not applicable
func (m *HardwareConfigurationRunSummary) SetNotApplicableUserCount(value *int32)() {
    if m != nil {
        m.notApplicableUserCount = value
    }
}
// SetPendingDeviceCount sets the pendingDeviceCount property value. Number of devices for which hardware configuration is in pending state
func (m *HardwareConfigurationRunSummary) SetPendingDeviceCount(value *int32)() {
    if m != nil {
        m.pendingDeviceCount = value
    }
}
// SetPendingUserCount sets the pendingUserCount property value. Number of users for which hardware configuration is in pending state
func (m *HardwareConfigurationRunSummary) SetPendingUserCount(value *int32)() {
    if m != nil {
        m.pendingUserCount = value
    }
}
// SetSuccessfulDeviceCount sets the successfulDeviceCount property value. Number of devices for which hardware configured without any issue
func (m *HardwareConfigurationRunSummary) SetSuccessfulDeviceCount(value *int32)() {
    if m != nil {
        m.successfulDeviceCount = value
    }
}
// SetSuccessfulUserCount sets the successfulUserCount property value. Number of users for which hardware configured without any issue
func (m *HardwareConfigurationRunSummary) SetSuccessfulUserCount(value *int32)() {
    if m != nil {
        m.successfulUserCount = value
    }
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Number of devices for which hardware configuration state is unknown
func (m *HardwareConfigurationRunSummary) SetUnknownDeviceCount(value *int32)() {
    if m != nil {
        m.unknownDeviceCount = value
    }
}
// SetUnknownUserCount sets the unknownUserCount property value. Number of users for which hardware configuration state is unknown
func (m *HardwareConfigurationRunSummary) SetUnknownUserCount(value *int32)() {
    if m != nil {
        m.unknownUserCount = value
    }
}
