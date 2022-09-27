package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSoftwareUpdateCategorySummary macOS software update category summary report for a device and user
type MacOSSoftwareUpdateCategorySummary struct {
    Entity
    // The device ID.
    deviceId *string
    // The name of the report
    displayName *string
    // Number of failed updates on the device
    failedUpdateCount *int32
    // Last date time the report for this device was updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Number of successful updates on the device
    successfulUpdateCount *int32
    // Number of total updates on the device
    totalUpdateCount *int32
    // MacOS Software Update Category
    updateCategory *MacOSSoftwareUpdateCategory
    // Summary of the update states.
    updateStateSummaries []MacOSSoftwareUpdateStateSummaryable
    // The user ID.
    userId *string
}
// NewMacOSSoftwareUpdateCategorySummary instantiates a new macOSSoftwareUpdateCategorySummary and sets the default values.
func NewMacOSSoftwareUpdateCategorySummary()(*MacOSSoftwareUpdateCategorySummary) {
    m := &MacOSSoftwareUpdateCategorySummary{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.macOSSoftwareUpdateCategorySummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSSoftwareUpdateCategorySummary(), nil
}
// GetDeviceId gets the deviceId property value. The device ID.
func (m *MacOSSoftwareUpdateCategorySummary) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDisplayName gets the displayName property value. The name of the report
func (m *MacOSSoftwareUpdateCategorySummary) GetDisplayName()(*string) {
    return m.displayName
}
// GetFailedUpdateCount gets the failedUpdateCount property value. Number of failed updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) GetFailedUpdateCount()(*int32) {
    return m.failedUpdateCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSoftwareUpdateCategorySummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["failedUpdateCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetFailedUpdateCount)
    res["lastUpdatedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastUpdatedDateTime)
    res["successfulUpdateCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSuccessfulUpdateCount)
    res["totalUpdateCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalUpdateCount)
    res["updateCategory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMacOSSoftwareUpdateCategory , m.SetUpdateCategory)
    res["updateStateSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMacOSSoftwareUpdateStateSummaryFromDiscriminatorValue , m.SetUpdateStateSummaries)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
func (m *MacOSSoftwareUpdateCategorySummary) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetSuccessfulUpdateCount gets the successfulUpdateCount property value. Number of successful updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) GetSuccessfulUpdateCount()(*int32) {
    return m.successfulUpdateCount
}
// GetTotalUpdateCount gets the totalUpdateCount property value. Number of total updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) GetTotalUpdateCount()(*int32) {
    return m.totalUpdateCount
}
// GetUpdateCategory gets the updateCategory property value. MacOS Software Update Category
func (m *MacOSSoftwareUpdateCategorySummary) GetUpdateCategory()(*MacOSSoftwareUpdateCategory) {
    return m.updateCategory
}
// GetUpdateStateSummaries gets the updateStateSummaries property value. Summary of the update states.
func (m *MacOSSoftwareUpdateCategorySummary) GetUpdateStateSummaries()([]MacOSSoftwareUpdateStateSummaryable) {
    return m.updateStateSummaries
}
// GetUserId gets the userId property value. The user ID.
func (m *MacOSSoftwareUpdateCategorySummary) GetUserId()(*string) {
    return m.userId
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUpdateStateSummaries())
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
    m.deviceId = value
}
// SetDisplayName sets the displayName property value. The name of the report
func (m *MacOSSoftwareUpdateCategorySummary) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFailedUpdateCount sets the failedUpdateCount property value. Number of failed updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) SetFailedUpdateCount(value *int32)() {
    m.failedUpdateCount = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
func (m *MacOSSoftwareUpdateCategorySummary) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetSuccessfulUpdateCount sets the successfulUpdateCount property value. Number of successful updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) SetSuccessfulUpdateCount(value *int32)() {
    m.successfulUpdateCount = value
}
// SetTotalUpdateCount sets the totalUpdateCount property value. Number of total updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) SetTotalUpdateCount(value *int32)() {
    m.totalUpdateCount = value
}
// SetUpdateCategory sets the updateCategory property value. MacOS Software Update Category
func (m *MacOSSoftwareUpdateCategorySummary) SetUpdateCategory(value *MacOSSoftwareUpdateCategory)() {
    m.updateCategory = value
}
// SetUpdateStateSummaries sets the updateStateSummaries property value. Summary of the update states.
func (m *MacOSSoftwareUpdateCategorySummary) SetUpdateStateSummaries(value []MacOSSoftwareUpdateStateSummaryable)() {
    m.updateStateSummaries = value
}
// SetUserId sets the userId property value. The user ID.
func (m *MacOSSoftwareUpdateCategorySummary) SetUserId(value *string)() {
    m.userId = value
}
