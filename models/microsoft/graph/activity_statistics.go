package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ActivityStatistics struct {
    Entity
    // The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and meeting.
    activity *AnalyticsActivityType;
    // Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
    duration *string;
    // Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-03' that follows the YYYY-MM-DD format.
    endDate *string;
    // Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-04' that follows the YYYY-MM-DD format.
    startDate *string;
    // The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value could be 'Pacific Standard Time.'
    timeZoneUsed *string;
}
// Instantiates a new activityStatistics and sets the default values.
func NewActivityStatistics()(*ActivityStatistics) {
    m := &ActivityStatistics{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activity property value. The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and meeting.
func (m *ActivityStatistics) GetActivity()(*AnalyticsActivityType) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// Gets the duration property value. Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
func (m *ActivityStatistics) GetDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// Gets the endDate property value. Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-03' that follows the YYYY-MM-DD format.
func (m *ActivityStatistics) GetEndDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endDate
    }
}
// Gets the startDate property value. Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-04' that follows the YYYY-MM-DD format.
func (m *ActivityStatistics) GetStartDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startDate
    }
}
// Gets the timeZoneUsed property value. The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value could be 'Pacific Standard Time.'
func (m *ActivityStatistics) GetTimeZoneUsed()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZoneUsed
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the activity property value. The type of activity for which statistics are returned. The possible values are: call, chat, email, focus, and meeting.
// Parameters:
//  - value : Value to set for the activity property.
func (m *ActivityStatistics) SetActivity(value *AnalyticsActivityType)() {
    m.activity = value
}
// Sets the duration property value. Total hours spent on the activity. The value is represented in ISO 8601 format for durations.
// Parameters:
//  - value : Value to set for the duration property.
func (m *ActivityStatistics) SetDuration(value *string)() {
    m.duration = value
}
// Sets the endDate property value. Date when the activity ended, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-03' that follows the YYYY-MM-DD format.
// Parameters:
//  - value : Value to set for the endDate property.
func (m *ActivityStatistics) SetEndDate(value *string)() {
    m.endDate = value
}
// Sets the startDate property value. Date when the activity started, expressed in ISO 8601 format for calendar dates. For example, the property value could be '2019-07-04' that follows the YYYY-MM-DD format.
// Parameters:
//  - value : Value to set for the startDate property.
func (m *ActivityStatistics) SetStartDate(value *string)() {
    m.startDate = value
}
// Sets the timeZoneUsed property value. The time zone that the user sets in Microsoft Outlook is used for the computation. For example, the property value could be 'Pacific Standard Time.'
// Parameters:
//  - value : Value to set for the timeZoneUsed property.
func (m *ActivityStatistics) SetTimeZoneUsed(value *string)() {
    m.timeZoneUsed = value
}
