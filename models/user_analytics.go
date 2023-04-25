package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserAnalytics 
type UserAnalytics struct {
    Entity
}
// NewUserAnalytics instantiates a new UserAnalytics and sets the default values.
func NewUserAnalytics()(*UserAnalytics) {
    m := &UserAnalytics{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserAnalyticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserAnalyticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserAnalytics(), nil
}
// GetActivityStatistics gets the activityStatistics property value. The collection of work activities that a user spent time on during and outside of working hours. Read-only. Nullable.
func (m *UserAnalytics) GetActivityStatistics()([]ActivityStatisticsable) {
    val, err := m.GetBackingStore().Get("activityStatistics")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ActivityStatisticsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserAnalytics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityStatistics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActivityStatisticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActivityStatisticsable, len(val))
            for i, v := range val {
                res[i] = v.(ActivityStatisticsable)
            }
            m.SetActivityStatistics(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(Settingsable))
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. The current settings for a user to use the analytics API.
func (m *UserAnalytics) GetSettings()(Settingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Settingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserAnalytics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivityStatistics() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivityStatistics()))
        for i, v := range m.GetActivityStatistics() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("activityStatistics", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityStatistics sets the activityStatistics property value. The collection of work activities that a user spent time on during and outside of working hours. Read-only. Nullable.
func (m *UserAnalytics) SetActivityStatistics(value []ActivityStatisticsable)() {
    err := m.GetBackingStore().Set("activityStatistics", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. The current settings for a user to use the analytics API.
func (m *UserAnalytics) SetSettings(value Settingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// UserAnalyticsable 
type UserAnalyticsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityStatistics()([]ActivityStatisticsable)
    GetSettings()(Settingsable)
    SetActivityStatistics(value []ActivityStatisticsable)()
    SetSettings(value Settingsable)()
}
