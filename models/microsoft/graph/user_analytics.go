package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserAnalytics 
type UserAnalytics struct {
    Entity
    // The collection of work activities that a user spent time on during and outside of working hours. Read-only. Nullable.
    activityStatistics []ActivityStatistics;
    // The current settings for a user to use the analytics API.
    settings *Settings;
}
// NewUserAnalytics instantiates a new userAnalytics and sets the default values.
func NewUserAnalytics()(*UserAnalytics) {
    m := &UserAnalytics{
        Entity: *NewEntity(),
    }
    return m
}
// GetActivityStatistics gets the activityStatistics property value. The collection of work activities that a user spent time on during and outside of working hours. Read-only. Nullable.
func (m *UserAnalytics) GetActivityStatistics()([]ActivityStatistics) {
    if m == nil {
        return nil
    } else {
        return m.activityStatistics
    }
}
// GetSettings gets the settings property value. The current settings for a user to use the analytics API.
func (m *UserAnalytics) GetSettings()(*Settings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserAnalytics) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityStatistics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewActivityStatistics() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActivityStatistics, len(val))
            for i, v := range val {
                res[i] = *(v.(*ActivityStatistics))
            }
            m.SetActivityStatistics(res)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(*Settings))
        }
        return nil
    }
    return res
}
func (m *UserAnalytics) IsNil()(bool) {
    return m == nil
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
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
func (m *UserAnalytics) SetActivityStatistics(value []ActivityStatistics)() {
    if m != nil {
        m.activityStatistics = value
    }
}
// SetSettings sets the settings property value. The current settings for a user to use the analytics API.
func (m *UserAnalytics) SetSettings(value *Settings)() {
    if m != nil {
        m.settings = value
    }
}
