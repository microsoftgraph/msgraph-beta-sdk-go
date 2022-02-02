package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MacOSSoftwareUpdateAccountSummary 
type MacOSSoftwareUpdateAccountSummary struct {
    Entity
    // Summary of the updates by category.
    categorySummaries []MacOSSoftwareUpdateCategorySummary;
    // The device ID.
    deviceId *string;
    // The device name.
    deviceName *string;
    // The name of the report
    displayName *string;
    // Number of failed updates on the device.
    failedUpdateCount *int32;
    // Last date time the report for this device was updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The OS version.
    osVersion *string;
    // Number of successful updates on the device.
    successfulUpdateCount *int32;
    // Number of total updates on the device.
    totalUpdateCount *int32;
    // The user ID.
    userId *string;
    // The user principal name
    userPrincipalName *string;
}
// NewMacOSSoftwareUpdateAccountSummary instantiates a new macOSSoftwareUpdateAccountSummary and sets the default values.
func NewMacOSSoftwareUpdateAccountSummary()(*MacOSSoftwareUpdateAccountSummary) {
    m := &MacOSSoftwareUpdateAccountSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetCategorySummaries gets the categorySummaries property value. Summary of the updates by category.
func (m *MacOSSoftwareUpdateAccountSummary) GetCategorySummaries()([]MacOSSoftwareUpdateCategorySummary) {
    if m == nil {
        return nil
    } else {
        return m.categorySummaries
    }
}
// GetDeviceId gets the deviceId property value. The device ID.
func (m *MacOSSoftwareUpdateAccountSummary) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDeviceName gets the deviceName property value. The device name.
func (m *MacOSSoftwareUpdateAccountSummary) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetDisplayName gets the displayName property value. The name of the report
func (m *MacOSSoftwareUpdateAccountSummary) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFailedUpdateCount gets the failedUpdateCount property value. Number of failed updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) GetFailedUpdateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedUpdateCount
    }
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
func (m *MacOSSoftwareUpdateAccountSummary) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// GetOsVersion gets the osVersion property value. The OS version.
func (m *MacOSSoftwareUpdateAccountSummary) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetSuccessfulUpdateCount gets the successfulUpdateCount property value. Number of successful updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) GetSuccessfulUpdateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successfulUpdateCount
    }
}
// GetTotalUpdateCount gets the totalUpdateCount property value. Number of total updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) GetTotalUpdateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalUpdateCount
    }
}
// GetUserId gets the userId property value. The user ID.
func (m *MacOSSoftwareUpdateAccountSummary) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name
func (m *MacOSSoftwareUpdateAccountSummary) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSoftwareUpdateAccountSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categorySummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMacOSSoftwareUpdateCategorySummary() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSSoftwareUpdateCategorySummary, len(val))
            for i, v := range val {
                res[i] = *(v.(*MacOSSoftwareUpdateCategorySummary))
            }
            m.SetCategorySummaries(res)
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["failedUpdateCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedUpdateCount(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
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
    res["successfulUpdateCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulUpdateCount(val)
        }
        return nil
    }
    res["totalUpdateCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUpdateCount(val)
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
func (m *MacOSSoftwareUpdateAccountSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MacOSSoftwareUpdateAccountSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategorySummaries() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCategorySummaries()))
        for i, v := range m.GetCategorySummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("categorySummaries", cast)
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
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedUpdateCount", m.GetFailedUpdateCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
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
        err = writer.WriteInt32Value("successfulUpdateCount", m.GetSuccessfulUpdateCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalUpdateCount", m.GetTotalUpdateCount())
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
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategorySummaries sets the categorySummaries property value. Summary of the updates by category.
func (m *MacOSSoftwareUpdateAccountSummary) SetCategorySummaries(value []MacOSSoftwareUpdateCategorySummary)() {
    if m != nil {
        m.categorySummaries = value
    }
}
// SetDeviceId sets the deviceId property value. The device ID.
func (m *MacOSSoftwareUpdateAccountSummary) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDeviceName sets the deviceName property value. The device name.
func (m *MacOSSoftwareUpdateAccountSummary) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetDisplayName sets the displayName property value. The name of the report
func (m *MacOSSoftwareUpdateAccountSummary) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFailedUpdateCount sets the failedUpdateCount property value. Number of failed updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) SetFailedUpdateCount(value *int32)() {
    if m != nil {
        m.failedUpdateCount = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
func (m *MacOSSoftwareUpdateAccountSummary) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetOsVersion sets the osVersion property value. The OS version.
func (m *MacOSSoftwareUpdateAccountSummary) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetSuccessfulUpdateCount sets the successfulUpdateCount property value. Number of successful updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) SetSuccessfulUpdateCount(value *int32)() {
    if m != nil {
        m.successfulUpdateCount = value
    }
}
// SetTotalUpdateCount sets the totalUpdateCount property value. Number of total updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) SetTotalUpdateCount(value *int32)() {
    if m != nil {
        m.totalUpdateCount = value
    }
}
// SetUserId sets the userId property value. The user ID.
func (m *MacOSSoftwareUpdateAccountSummary) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name
func (m *MacOSSoftwareUpdateAccountSummary) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
