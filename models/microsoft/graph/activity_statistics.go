package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ActivityStatistics 
type ActivityStatistics struct {
    Entity
    // The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and meeting.
    activity *AnalyticsActivityType;
    // Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
    duration *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-03' that follows the YYYY-MM-DD format.
    endDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-04' that follows the YYYY-MM-DD format.
    startDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value could be 'Pacific Standard Time.'
    timeZoneUsed *string;
}
// NewActivityStatistics instantiates a new activityStatistics and sets the default values.
func NewActivityStatistics()(*ActivityStatistics) {
    m := &ActivityStatistics{
        Entity: *NewEntity(),
    }
    return m
}
// GetActivity gets the activity property value. The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and meeting.
func (m *ActivityStatistics) GetActivity()(*AnalyticsActivityType) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetDuration gets the duration property value. Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
func (m *ActivityStatistics) GetDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// GetEndDate gets the endDate property value. Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-03' that follows the YYYY-MM-DD format.
func (m *ActivityStatistics) GetEndDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.endDate
    }
}
// GetStartDate gets the startDate property value. Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-04' that follows the YYYY-MM-DD format.
func (m *ActivityStatistics) GetStartDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.startDate
    }
}
// GetTimeZoneUsed gets the timeZoneUsed property value. The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value could be 'Pacific Standard Time.'
func (m *ActivityStatistics) GetTimeZoneUsed()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZoneUsed
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActivityStatistics) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAnalyticsActivityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(*AnalyticsActivityType))
        }
        return nil
    }
    res["duration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["endDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["startDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val)
        }
        return nil
    }
    res["timeZoneUsed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZoneUsed(val)
        }
        return nil
    }
    return res
}
func (m *ActivityStatistics) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ActivityStatistics) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivity() != nil {
        cast := (*m.GetActivity()).String()
        err = writer.WriteStringValue("activity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("startDate", m.GetStartDate())
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
// SetActivity sets the activity property value. The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and meeting.
func (m *ActivityStatistics) SetActivity(value *AnalyticsActivityType)() {
    if m != nil {
        m.activity = value
    }
}
// SetDuration sets the duration property value. Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
func (m *ActivityStatistics) SetDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.duration = value
    }
}
// SetEndDate sets the endDate property value. Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-03' that follows the YYYY-MM-DD format.
func (m *ActivityStatistics) SetEndDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.endDate = value
    }
}
// SetStartDate sets the startDate property value. Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-04' that follows the YYYY-MM-DD format.
func (m *ActivityStatistics) SetStartDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.startDate = value
    }
}
// SetTimeZoneUsed sets the timeZoneUsed property value. The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value could be 'Pacific Standard Time.'
func (m *ActivityStatistics) SetTimeZoneUsed(value *string)() {
    if m != nil {
        m.timeZoneUsed = value
    }
}
