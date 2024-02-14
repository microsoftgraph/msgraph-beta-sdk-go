package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSoftwareUpdateCategorySummary macOS software update category summary report for a device and user
type MacOSSoftwareUpdateCategorySummary struct {
    Entity
}
// NewMacOSSoftwareUpdateCategorySummary instantiates a new MacOSSoftwareUpdateCategorySummary and sets the default values.
func NewMacOSSoftwareUpdateCategorySummary()(*MacOSSoftwareUpdateCategorySummary) {
    m := &MacOSSoftwareUpdateCategorySummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSSoftwareUpdateCategorySummary(), nil
}
// GetDeviceId gets the deviceId property value. The device ID.
// returns a *string when successful
func (m *MacOSSoftwareUpdateCategorySummary) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
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
func (m *MacOSSoftwareUpdateCategorySummary) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFailedUpdateCount gets the failedUpdateCount property value. Number of failed updates on the device
// returns a *int32 when successful
func (m *MacOSSoftwareUpdateCategorySummary) GetFailedUpdateCount()(*int32) {
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
func (m *MacOSSoftwareUpdateCategorySummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["updateCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCategory(val.(*MacOSSoftwareUpdateCategory))
        }
        return nil
    }
    res["updateStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSSoftwareUpdateStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSSoftwareUpdateStateSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MacOSSoftwareUpdateStateSummaryable)
                }
            }
            m.SetUpdateStateSummaries(res)
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
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
// returns a *Time when successful
func (m *MacOSSoftwareUpdateCategorySummary) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastUpdatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSuccessfulUpdateCount gets the successfulUpdateCount property value. Number of successful updates on the device
// returns a *int32 when successful
func (m *MacOSSoftwareUpdateCategorySummary) GetSuccessfulUpdateCount()(*int32) {
    val, err := m.GetBackingStore().Get("successfulUpdateCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalUpdateCount gets the totalUpdateCount property value. Number of total updates on the device
// returns a *int32 when successful
func (m *MacOSSoftwareUpdateCategorySummary) GetTotalUpdateCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalUpdateCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUpdateCategory gets the updateCategory property value. MacOS Software Update Category
// returns a *MacOSSoftwareUpdateCategory when successful
func (m *MacOSSoftwareUpdateCategorySummary) GetUpdateCategory()(*MacOSSoftwareUpdateCategory) {
    val, err := m.GetBackingStore().Get("updateCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSSoftwareUpdateCategory)
    }
    return nil
}
// GetUpdateStateSummaries gets the updateStateSummaries property value. Summary of the update states.
// returns a []MacOSSoftwareUpdateStateSummaryable when successful
func (m *MacOSSoftwareUpdateCategorySummary) GetUpdateStateSummaries()([]MacOSSoftwareUpdateStateSummaryable) {
    val, err := m.GetBackingStore().Get("updateStateSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSSoftwareUpdateStateSummaryable)
    }
    return nil
}
// GetUserId gets the userId property value. The user ID.
// returns a *string when successful
func (m *MacOSSoftwareUpdateCategorySummary) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSSoftwareUpdateCategorySummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
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
    if m.GetUpdateCategory() != nil {
        cast := (*m.GetUpdateCategory()).String()
        err = writer.WriteStringValue("updateCategory", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdateStateSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUpdateStateSummaries()))
        for i, v := range m.GetUpdateStateSummaries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("updateStateSummaries", cast)
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
// SetDeviceId sets the deviceId property value. The device ID.
func (m *MacOSSoftwareUpdateCategorySummary) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of the report
func (m *MacOSSoftwareUpdateCategorySummary) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetFailedUpdateCount sets the failedUpdateCount property value. Number of failed updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) SetFailedUpdateCount(value *int32)() {
    err := m.GetBackingStore().Set("failedUpdateCount", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
func (m *MacOSSoftwareUpdateCategorySummary) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastUpdatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulUpdateCount sets the successfulUpdateCount property value. Number of successful updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) SetSuccessfulUpdateCount(value *int32)() {
    err := m.GetBackingStore().Set("successfulUpdateCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalUpdateCount sets the totalUpdateCount property value. Number of total updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) SetTotalUpdateCount(value *int32)() {
    err := m.GetBackingStore().Set("totalUpdateCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateCategory sets the updateCategory property value. MacOS Software Update Category
func (m *MacOSSoftwareUpdateCategorySummary) SetUpdateCategory(value *MacOSSoftwareUpdateCategory)() {
    err := m.GetBackingStore().Set("updateCategory", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateStateSummaries sets the updateStateSummaries property value. Summary of the update states.
func (m *MacOSSoftwareUpdateCategorySummary) SetUpdateStateSummaries(value []MacOSSoftwareUpdateStateSummaryable)() {
    err := m.GetBackingStore().Set("updateStateSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The user ID.
func (m *MacOSSoftwareUpdateCategorySummary) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
type MacOSSoftwareUpdateCategorySummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceId()(*string)
    GetDisplayName()(*string)
    GetFailedUpdateCount()(*int32)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSuccessfulUpdateCount()(*int32)
    GetTotalUpdateCount()(*int32)
    GetUpdateCategory()(*MacOSSoftwareUpdateCategory)
    GetUpdateStateSummaries()([]MacOSSoftwareUpdateStateSummaryable)
    GetUserId()(*string)
    SetDeviceId(value *string)()
    SetDisplayName(value *string)()
    SetFailedUpdateCount(value *int32)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSuccessfulUpdateCount(value *int32)()
    SetTotalUpdateCount(value *int32)()
    SetUpdateCategory(value *MacOSSoftwareUpdateCategory)()
    SetUpdateStateSummaries(value []MacOSSoftwareUpdateStateSummaryable)()
    SetUserId(value *string)()
}
