package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion the user experience analytics application performance entity contains app performance details by app version.
type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion struct {
    Entity
}
// NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersion instantiates a new userExperienceAnalyticsAppHealthAppPerformanceByAppVersion and sets the default values.
func NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersion()(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) {
    m := &UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersion(), nil
}
// GetAppCrashCount gets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppCrashCount()(*int32) {
    val, err := m.GetBackingStore().Get("appCrashCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAppDisplayName gets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("appDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppName gets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppName()(*string) {
    val, err := m.GetBackingStore().Get("appName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppPublisher gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppPublisher()(*string) {
    val, err := m.GetBackingStore().Get("appPublisher")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppUsageDuration gets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppUsageDuration()(*int32) {
    val, err := m.GetBackingStore().Get("appUsageDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAppVersion gets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppVersion()(*string) {
    val, err := m.GetBackingStore().Get("appVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appCrashCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppCrashCount(val)
        }
        return nil
    }
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppName(val)
        }
        return nil
    }
    res["appPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppPublisher(val)
        }
        return nil
    }
    res["appUsageDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppUsageDuration(val)
        }
        return nil
    }
    res["appVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppVersion(val)
        }
        return nil
    }
    res["meanTimeToFailureInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeanTimeToFailureInMinutes(val)
        }
        return nil
    }
    return res
}
// GetMeanTimeToFailureInMinutes gets the meanTimeToFailureInMinutes property value. The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetMeanTimeToFailureInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("meanTimeToFailureInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("appUsageDuration", m.GetAppUsageDuration())
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
        err = writer.WriteInt32Value("meanTimeToFailureInMinutes", m.GetMeanTimeToFailureInMinutes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppCrashCount sets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppCrashCount(value *int32)() {
    err := m.GetBackingStore().Set("appCrashCount", value)
    if err != nil {
        panic(err)
    }
}
// SetAppDisplayName sets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppDisplayName(value *string)() {
    err := m.GetBackingStore().Set("appDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAppName sets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppName(value *string)() {
    err := m.GetBackingStore().Set("appName", value)
    if err != nil {
        panic(err)
    }
}
// SetAppPublisher sets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppPublisher(value *string)() {
    err := m.GetBackingStore().Set("appPublisher", value)
    if err != nil {
        panic(err)
    }
}
// SetAppUsageDuration sets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppUsageDuration(value *int32)() {
    err := m.GetBackingStore().Set("appUsageDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetAppVersion sets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppVersion(value *string)() {
    err := m.GetBackingStore().Set("appVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMeanTimeToFailureInMinutes sets the meanTimeToFailureInMinutes property value. The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetMeanTimeToFailureInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("meanTimeToFailureInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable 
type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppCrashCount()(*int32)
    GetAppDisplayName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetAppUsageDuration()(*int32)
    GetAppVersion()(*string)
    GetMeanTimeToFailureInMinutes()(*int32)
    SetAppCrashCount(value *int32)()
    SetAppDisplayName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetAppUsageDuration(value *int32)()
    SetAppVersion(value *string)()
    SetMeanTimeToFailureInMinutes(value *int32)()
}
