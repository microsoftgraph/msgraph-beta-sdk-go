package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSoftwareUpdateAccountSummary macOS software update account summary report for a device and user
type MacOSSoftwareUpdateAccountSummary struct {
    Entity
}
// NewMacOSSoftwareUpdateAccountSummary instantiates a new MacOSSoftwareUpdateAccountSummary and sets the default values.
func NewMacOSSoftwareUpdateAccountSummary()(*MacOSSoftwareUpdateAccountSummary) {
    m := &MacOSSoftwareUpdateAccountSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSSoftwareUpdateAccountSummary(), nil
}
// GetCategorySummaries gets the categorySummaries property value. Summary of the updates by category.
// returns a []MacOSSoftwareUpdateCategorySummaryable when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetCategorySummaries()([]MacOSSoftwareUpdateCategorySummaryable) {
    val, err := m.GetBackingStore().Get("categorySummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSSoftwareUpdateCategorySummaryable)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The device ID.
// returns a *string when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The device name.
// returns a *string when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of the report
// returns a *string when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFailedUpdateCount gets the failedUpdateCount property value. Number of failed updates on the device.
// returns a *int32 when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetFailedUpdateCount()(*int32) {
    val, err := m.GetBackingStore().Get("failedUpdateCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categorySummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSSoftwareUpdateCategorySummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MacOSSoftwareUpdateCategorySummaryable)
                }
            }
            m.SetCategorySummaries(res)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["failedUpdateCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedUpdateCount(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["successfulUpdateCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulUpdateCount(val)
        }
        return nil
    }
    res["totalUpdateCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUpdateCount(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
// returns a *Time when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastUpdatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. The OS version.
// returns a *string when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSuccessfulUpdateCount gets the successfulUpdateCount property value. Number of successful updates on the device.
// returns a *int32 when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetSuccessfulUpdateCount()(*int32) {
    val, err := m.GetBackingStore().Get("successfulUpdateCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalUpdateCount gets the totalUpdateCount property value. Number of total updates on the device.
// returns a *int32 when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetTotalUpdateCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalUpdateCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUserId gets the userId property value. The user ID.
// returns a *string when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name
// returns a *string when successful
func (m *MacOSSoftwareUpdateAccountSummary) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSSoftwareUpdateAccountSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategorySummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategorySummaries()))
        for i, v := range m.GetCategorySummaries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
func (m *MacOSSoftwareUpdateAccountSummary) SetCategorySummaries(value []MacOSSoftwareUpdateCategorySummaryable)() {
    err := m.GetBackingStore().Set("categorySummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The device ID.
func (m *MacOSSoftwareUpdateAccountSummary) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The device name.
func (m *MacOSSoftwareUpdateAccountSummary) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of the report
func (m *MacOSSoftwareUpdateAccountSummary) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetFailedUpdateCount sets the failedUpdateCount property value. Number of failed updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) SetFailedUpdateCount(value *int32)() {
    err := m.GetBackingStore().Set("failedUpdateCount", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
func (m *MacOSSoftwareUpdateAccountSummary) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastUpdatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. The OS version.
func (m *MacOSSoftwareUpdateAccountSummary) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulUpdateCount sets the successfulUpdateCount property value. Number of successful updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) SetSuccessfulUpdateCount(value *int32)() {
    err := m.GetBackingStore().Set("successfulUpdateCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalUpdateCount sets the totalUpdateCount property value. Number of total updates on the device.
func (m *MacOSSoftwareUpdateAccountSummary) SetTotalUpdateCount(value *int32)() {
    err := m.GetBackingStore().Set("totalUpdateCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The user ID.
func (m *MacOSSoftwareUpdateAccountSummary) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name
func (m *MacOSSoftwareUpdateAccountSummary) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
type MacOSSoftwareUpdateAccountSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategorySummaries()([]MacOSSoftwareUpdateCategorySummaryable)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetDisplayName()(*string)
    GetFailedUpdateCount()(*int32)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOsVersion()(*string)
    GetSuccessfulUpdateCount()(*int32)
    GetTotalUpdateCount()(*int32)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetCategorySummaries(value []MacOSSoftwareUpdateCategorySummaryable)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetDisplayName(value *string)()
    SetFailedUpdateCount(value *int32)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOsVersion(value *string)()
    SetSuccessfulUpdateCount(value *int32)()
    SetTotalUpdateCount(value *int32)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
