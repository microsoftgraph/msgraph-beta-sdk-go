package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBaseline the user experience analytics baseline entity contains baseline values against which to compare the user experience analytics scores.
type UserExperienceAnalyticsBaseline struct {
    Entity
    // The user experience analytics app health metrics.
    appHealthMetrics UserExperienceAnalyticsCategoryable
    // The user experience analytics battery health metrics.
    batteryHealthMetrics UserExperienceAnalyticsCategoryable
    // The user experience analytics best practices metrics.
    bestPracticesMetrics UserExperienceAnalyticsCategoryable
    // The date the custom baseline was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user experience analytics device boot performance metrics.
    deviceBootPerformanceMetrics UserExperienceAnalyticsCategoryable
    // The name of the user experience analytics baseline.
    displayName *string
    // Signifies if the current baseline is the commercial median baseline or a custom baseline.
    isBuiltIn *bool
    // The user experience analytics reboot analytics metrics.
    rebootAnalyticsMetrics UserExperienceAnalyticsCategoryable
    // The user experience analytics resource performance metrics.
    resourcePerformanceMetrics UserExperienceAnalyticsCategoryable
    // The user experience analytics work from anywhere metrics.
    workFromAnywhereMetrics UserExperienceAnalyticsCategoryable
}
// NewUserExperienceAnalyticsBaseline instantiates a new userExperienceAnalyticsBaseline and sets the default values.
func NewUserExperienceAnalyticsBaseline()(*UserExperienceAnalyticsBaseline) {
    m := &UserExperienceAnalyticsBaseline{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBaseline(), nil
}
// GetAppHealthMetrics gets the appHealthMetrics property value. The user experience analytics app health metrics.
func (m *UserExperienceAnalyticsBaseline) GetAppHealthMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.appHealthMetrics
}
// GetBatteryHealthMetrics gets the batteryHealthMetrics property value. The user experience analytics battery health metrics.
func (m *UserExperienceAnalyticsBaseline) GetBatteryHealthMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.batteryHealthMetrics
}
// GetBestPracticesMetrics gets the bestPracticesMetrics property value. The user experience analytics best practices metrics.
func (m *UserExperienceAnalyticsBaseline) GetBestPracticesMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.bestPracticesMetrics
}
// GetCreatedDateTime gets the createdDateTime property value. The date the custom baseline was created.
func (m *UserExperienceAnalyticsBaseline) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeviceBootPerformanceMetrics gets the deviceBootPerformanceMetrics property value. The user experience analytics device boot performance metrics.
func (m *UserExperienceAnalyticsBaseline) GetDeviceBootPerformanceMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.deviceBootPerformanceMetrics
}
// GetDisplayName gets the displayName property value. The name of the user experience analytics baseline.
func (m *UserExperienceAnalyticsBaseline) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBaseline) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appHealthMetrics"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue , m.SetAppHealthMetrics)
    res["batteryHealthMetrics"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue , m.SetBatteryHealthMetrics)
    res["bestPracticesMetrics"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue , m.SetBestPracticesMetrics)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["deviceBootPerformanceMetrics"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue , m.SetDeviceBootPerformanceMetrics)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["isBuiltIn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsBuiltIn)
    res["rebootAnalyticsMetrics"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue , m.SetRebootAnalyticsMetrics)
    res["resourcePerformanceMetrics"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue , m.SetResourcePerformanceMetrics)
    res["workFromAnywhereMetrics"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue , m.SetWorkFromAnywhereMetrics)
    return res
}
// GetIsBuiltIn gets the isBuiltIn property value. Signifies if the current baseline is the commercial median baseline or a custom baseline.
func (m *UserExperienceAnalyticsBaseline) GetIsBuiltIn()(*bool) {
    return m.isBuiltIn
}
// GetRebootAnalyticsMetrics gets the rebootAnalyticsMetrics property value. The user experience analytics reboot analytics metrics.
func (m *UserExperienceAnalyticsBaseline) GetRebootAnalyticsMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.rebootAnalyticsMetrics
}
// GetResourcePerformanceMetrics gets the resourcePerformanceMetrics property value. The user experience analytics resource performance metrics.
func (m *UserExperienceAnalyticsBaseline) GetResourcePerformanceMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.resourcePerformanceMetrics
}
// GetWorkFromAnywhereMetrics gets the workFromAnywhereMetrics property value. The user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsBaseline) GetWorkFromAnywhereMetrics()(UserExperienceAnalyticsCategoryable) {
    return m.workFromAnywhereMetrics
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBaseline) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appHealthMetrics", m.GetAppHealthMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("batteryHealthMetrics", m.GetBatteryHealthMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bestPracticesMetrics", m.GetBestPracticesMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceBootPerformanceMetrics", m.GetDeviceBootPerformanceMetrics())
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
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rebootAnalyticsMetrics", m.GetRebootAnalyticsMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resourcePerformanceMetrics", m.GetResourcePerformanceMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("workFromAnywhereMetrics", m.GetWorkFromAnywhereMetrics())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppHealthMetrics sets the appHealthMetrics property value. The user experience analytics app health metrics.
func (m *UserExperienceAnalyticsBaseline) SetAppHealthMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.appHealthMetrics = value
}
// SetBatteryHealthMetrics sets the batteryHealthMetrics property value. The user experience analytics battery health metrics.
func (m *UserExperienceAnalyticsBaseline) SetBatteryHealthMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.batteryHealthMetrics = value
}
// SetBestPracticesMetrics sets the bestPracticesMetrics property value. The user experience analytics best practices metrics.
func (m *UserExperienceAnalyticsBaseline) SetBestPracticesMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.bestPracticesMetrics = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date the custom baseline was created.
func (m *UserExperienceAnalyticsBaseline) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeviceBootPerformanceMetrics sets the deviceBootPerformanceMetrics property value. The user experience analytics device boot performance metrics.
func (m *UserExperienceAnalyticsBaseline) SetDeviceBootPerformanceMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.deviceBootPerformanceMetrics = value
}
// SetDisplayName sets the displayName property value. The name of the user experience analytics baseline.
func (m *UserExperienceAnalyticsBaseline) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsBuiltIn sets the isBuiltIn property value. Signifies if the current baseline is the commercial median baseline or a custom baseline.
func (m *UserExperienceAnalyticsBaseline) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
// SetRebootAnalyticsMetrics sets the rebootAnalyticsMetrics property value. The user experience analytics reboot analytics metrics.
func (m *UserExperienceAnalyticsBaseline) SetRebootAnalyticsMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.rebootAnalyticsMetrics = value
}
// SetResourcePerformanceMetrics sets the resourcePerformanceMetrics property value. The user experience analytics resource performance metrics.
func (m *UserExperienceAnalyticsBaseline) SetResourcePerformanceMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.resourcePerformanceMetrics = value
}
// SetWorkFromAnywhereMetrics sets the workFromAnywhereMetrics property value. The user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsBaseline) SetWorkFromAnywhereMetrics(value UserExperienceAnalyticsCategoryable)() {
    m.workFromAnywhereMetrics = value
}
