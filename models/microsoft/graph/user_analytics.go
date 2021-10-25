package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserAnalytics struct {
    Entity
    activityStatistics []ActivityStatistics;
    settings *Settings;
}
func NewUserAnalytics()(*UserAnalytics) {
    m := &UserAnalytics{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserAnalytics) GetActivityStatistics()([]ActivityStatistics) {
    if m == nil {
        return nil
    } else {
        return m.activityStatistics
    }
}
func (m *UserAnalytics) GetSettings()(*Settings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *UserAnalytics) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityStatistics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewActivityStatistics() })
        if err != nil {
            return err
        }
        res := make([]ActivityStatistics, len(val))
        for i, v := range val {
            res[i] = *(v.(*ActivityStatistics))
        }
        m.SetActivityStatistics(res)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*Settings))
        return nil
    }
    return res
}
func (m *UserAnalytics) IsNil()(bool) {
    return m == nil
}
func (m *UserAnalytics) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
func (m *UserAnalytics) SetActivityStatistics(value []ActivityStatistics)() {
    m.activityStatistics = value
}
func (m *UserAnalytics) SetSettings(value *Settings)() {
    m.settings = value
}
