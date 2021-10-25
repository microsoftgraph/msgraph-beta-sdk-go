package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ActivityStatistics struct {
    Entity
    activity *AnalyticsActivityType;
    duration *string;
    endDate *string;
    startDate *string;
    timeZoneUsed *string;
}
func NewActivityStatistics()(*ActivityStatistics) {
    m := &ActivityStatistics{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ActivityStatistics) GetActivity()(*AnalyticsActivityType) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
func (m *ActivityStatistics) GetDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
func (m *ActivityStatistics) GetEndDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endDate
    }
}
func (m *ActivityStatistics) GetStartDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startDate
    }
}
func (m *ActivityStatistics) GetTimeZoneUsed()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZoneUsed
    }
}
func (m *ActivityStatistics) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAnalyticsActivityType)
        if err != nil {
            return err
        }
        cast := val.(AnalyticsActivityType)
        m.SetActivity(&cast)
        return nil
    }
    res["duration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDuration(val)
        return nil
    }
    res["endDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEndDate(val)
        return nil
    }
    res["startDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStartDate(val)
        return nil
    }
    res["timeZoneUsed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTimeZoneUsed(val)
        return nil
    }
    return res
}
func (m *ActivityStatistics) IsNil()(bool) {
    return m == nil
}
func (m *ActivityStatistics) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivity() != nil {
        cast := m.GetActivity().String()
        err = writer.WriteStringValue("activity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timeZoneUsed", m.GetTimeZoneUsed())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ActivityStatistics) SetActivity(value *AnalyticsActivityType)() {
    m.activity = value
}
func (m *ActivityStatistics) SetDuration(value *string)() {
    m.duration = value
}
func (m *ActivityStatistics) SetEndDate(value *string)() {
    m.endDate = value
}
func (m *ActivityStatistics) SetStartDate(value *string)() {
    m.startDate = value
}
func (m *ActivityStatistics) SetTimeZoneUsed(value *string)() {
    m.timeZoneUsed = value
}
