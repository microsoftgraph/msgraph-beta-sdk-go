package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId the user experience analytics application performance entity contains app performance by app version device id.
type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId struct {
    Entity
    // The number of crashes for the app. Valid values -2147483648 to 2147483647
    appCrashCount *int32
    // The friendly name of the application.
    appDisplayName *string
    // The name of the application.
    appName *string
    // The publisher of the application.
    appPublisher *string
    // The version of the application.
    appVersion *string
    // The name of the device.
    deviceDisplayName *string
    // The id of the device.
    deviceId *string
    // The date and time when the statistics were last computed.
    processedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId instantiates a new userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId and sets the default values.
func NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId()(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) {
    m := &UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId(), nil
}
// GetAppCrashCount gets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetAppCrashCount()(*int32) {
    return m.appCrashCount
}
// GetAppDisplayName gets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppName gets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetAppName()(*string) {
    return m.appName
}
// GetAppPublisher gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetAppPublisher()(*string) {
    return m.appPublisher
}
// GetAppVersion gets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetAppVersion()(*string) {
    return m.appVersion
}
// GetDeviceDisplayName gets the deviceDisplayName property value. The name of the device.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetDeviceDisplayName()(*string) {
    return m.deviceDisplayName
}
// GetDeviceId gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appCrashCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetAppCrashCount)
    res["appDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppDisplayName)
    res["appName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppName)
    res["appPublisher"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppPublisher)
    res["appVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppVersion)
    res["deviceDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceDisplayName)
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["processedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetProcessedDateTime)
    return res
}
// GetProcessedDateTime gets the processedDateTime property value. The date and time when the statistics were last computed.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) GetProcessedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.processedDateTime
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("appCrashCount", m.GetAppCrashCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appName", m.GetAppName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appPublisher", m.GetAppPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appVersion", m.GetAppVersion())
        if err != nil {
            return err
        }
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
        err = writer.WriteTimeValue("processedDateTime", m.GetProcessedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppCrashCount sets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
// SetAppDisplayName sets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppName sets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetAppName(value *string)() {
    m.appName = value
}
// SetAppPublisher sets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// SetAppVersion sets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetAppVersion(value *string)() {
    m.appVersion = value
}
// SetDeviceDisplayName sets the deviceDisplayName property value. The name of the device.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetProcessedDateTime sets the processedDateTime property value. The date and time when the statistics were last computed.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) SetProcessedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.processedDateTime = value
}
