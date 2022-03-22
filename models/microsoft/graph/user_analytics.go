package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserAnalytics 
type UserAnalytics struct {
    Entity
    // The collection of work activities that a user spent time on during and outside of working hours. Read-only. Nullable.
    activityStatistics []ActivityStatisticsable;
    // The current settings for a user to use the analytics API.
    settings Settingsable;
}
// NewUserAnalytics instantiates a new userAnalytics and sets the default values.
func NewUserAnalytics()(*UserAnalytics) {
    m := &UserAnalytics{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserAnalyticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserAnalyticsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserAnalytics(), nil
}
// GetActivityStatistics gets the activityStatistics property value. The collection of work activities that a user spent time on during and outside of working hours. Read-only. Nullable.
func (m *UserAnalytics) GetActivityStatistics()([]ActivityStatisticsable) {
    if m == nil {
        return nil
    } else {
        return m.activityStatistics
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserAnalytics) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityStatistics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Serialize serializes information the current object
func (m *UserAnalytics) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivityStatistics() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActivityStatistics()))
        for i, v := range m.GetActivityStatistics() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m != nil {
        m.activityStatistics = value
    }
}
// SetSettings sets the settings property value. The current settings for a user to use the analytics API.
func (m *UserAnalytics) SetSettings(value Settingsable)() {
    if m != nil {
        m.settings = value
    }
}
