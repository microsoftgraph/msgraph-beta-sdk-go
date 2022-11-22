package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSoftwareUpdateAccountSummary macOS software update account summary report for a device and user
type MacOSSoftwareUpdateAccountSummary struct {
    Entity
    // Summary of the updates by category.
    categorySummaries []MacOSSoftwareUpdateCategorySummaryable
    // The device ID.
    deviceId *string
    // The device name.
    deviceName *string
    // The name of the report
    displayName *string
    // Number of failed updates on the device.
    failedUpdateCount *int32
    // Last date time the report for this device was updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OS version.
    osVersion *string
    // Number of successful updates on the device.
    successfulUpdateCount *int32
    // Number of total updates on the device.
    totalUpdateCount *int32
    // The user ID.
    userId *string
    // The user principal name
    userPrincipalName *string
}
// NewMacOSSoftwareUpdateAccountSummary instantiates a new macOSSoftwareUpdateAccountSummary and sets the default values.
func NewMacOSSoftwareUpdateAccountSummary()(*MacOSSoftwareUpdateAccountSummary) {
    m := &MacOSSoftwareUpdateAccountSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSSoftwareUpdateAccountSummary(), nil
}
// GetCategorySummaries gets the categorySummaries property value. Summary of the updates by category.
func (m *MacOSSoftwareUpdateAccountSummary) GetCategorySummaries()([]MacOSSoftwareUpdateCategorySummaryable) {
    return m.categorySummaries
}
// GetDeviceId gets the deviceId property value. The device ID.
func (m *MacOSSoftwareUpdateAccountSummary) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceName gets the deviceName property value. The device name.
func (m *MacOSSoftwareUpdateAccountSummary) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDisplayName gets the displayName property value. The name of the report
func (m *MacOSSoftwareUpdateAccountSummary) GetDisplayName()(*string) {
    return m.displayName
}
// GetFailedUpdateCount gets the failedUpdateCount property value. Number of failed updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) GetFailedUpdateCount()(*int32) {
    return m.failedUpdateCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSoftwareUpdateAccountSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categorySummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue , m.SetCategorySummaries)
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["failedUpdateCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetFailedUpdateCount)
    res["lastUpdatedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastUpdatedDateTime)
    res["osVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsVersion)
    res["successfulUpdateCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSuccessfulUpdateCount)
    res["totalUpdateCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalUpdateCount)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
func (m *MacOSSoftwareUpdateAccountSummary) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetOsVersion gets the osVersion property value. The OS version.
func (m *MacOSSoftwareUpdateAccountSummary) GetOsVersion()(*string) {
    return m.osVersion
}
// GetSuccessfulUpdateCount gets the successfulUpdateCount property value. Number of successful updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) GetSuccessfulUpdateCount()(*int32) {
    return m.successfulUpdateCount
}
// GetTotalUpdateCount gets the totalUpdateCount property value. Number of total updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) GetTotalUpdateCount()(*int32) {
    return m.totalUpdateCount
}
// GetUserId gets the userId property value. The user ID.
func (m *MacOSSoftwareUpdateAccountSummary) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name
func (m *MacOSSoftwareUpdateAccountSummary) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *MacOSSoftwareUpdateAccountSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategorySummaries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCategorySummaries())
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
func (m *MacOSSoftwareUpdateAccountSummary) SetCategorySummaries(value []MacOSSoftwareUpdateCategorySummaryable)() {
    m.categorySummaries = value
}
// SetDeviceId sets the deviceId property value. The device ID.
func (m *MacOSSoftwareUpdateAccountSummary) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceName sets the deviceName property value. The device name.
func (m *MacOSSoftwareUpdateAccountSummary) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDisplayName sets the displayName property value. The name of the report
func (m *MacOSSoftwareUpdateAccountSummary) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFailedUpdateCount sets the failedUpdateCount property value. Number of failed updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) SetFailedUpdateCount(value *int32)() {
    m.failedUpdateCount = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
func (m *MacOSSoftwareUpdateAccountSummary) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetOsVersion sets the osVersion property value. The OS version.
func (m *MacOSSoftwareUpdateAccountSummary) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetSuccessfulUpdateCount sets the successfulUpdateCount property value. Number of successful updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) SetSuccessfulUpdateCount(value *int32)() {
    m.successfulUpdateCount = value
}
// SetTotalUpdateCount sets the totalUpdateCount property value. Number of total updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) SetTotalUpdateCount(value *int32)() {
    m.totalUpdateCount = value
}
// SetUserId sets the userId property value. The user ID.
func (m *MacOSSoftwareUpdateAccountSummary) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name
func (m *MacOSSoftwareUpdateAccountSummary) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
