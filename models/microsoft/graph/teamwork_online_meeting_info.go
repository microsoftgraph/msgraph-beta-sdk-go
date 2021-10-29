package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamworkOnlineMeetingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The identifier of the calendar event associated with the meeting.
    calendarEventId *string;
    // The URL which can be clicked on to join or uniquely identify the meeting.
    joinWebUrl *string;
    // The organizer of the meeting.
    organizer *TeamworkUserIdentity;
}
// Instantiates a new teamworkOnlineMeetingInfo and sets the default values.
func NewTeamworkOnlineMeetingInfo()(*TeamworkOnlineMeetingInfo) {
    m := &TeamworkOnlineMeetingInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkOnlineMeetingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the calendarEventId property value. The identifier of the calendar event associated with the meeting.
func (m *TeamworkOnlineMeetingInfo) GetCalendarEventId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.calendarEventId
    }
}
// Gets the joinWebUrl property value. The URL which can be clicked on to join or uniquely identify the meeting.
func (m *TeamworkOnlineMeetingInfo) GetJoinWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinWebUrl
    }
}
// Gets the organizer property value. The organizer of the meeting.
func (m *TeamworkOnlineMeetingInfo) GetOrganizer()(*TeamworkUserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.organizer
    }
}
// The deserialization information for the current model
func (m *TeamworkOnlineMeetingInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["calendarEventId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCalendarEventId(val)
        return nil
    }
    res["joinWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJoinWebUrl(val)
        return nil
    }
    res["organizer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkUserIdentity() })
        if err != nil {
            return err
        }
        m.SetOrganizer(val.(*TeamworkUserIdentity))
        return nil
    }
    return res
}
func (m *TeamworkOnlineMeetingInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TeamworkOnlineMeetingInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("calendarEventId", m.GetCalendarEventId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("organizer", m.GetOrganizer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TeamworkOnlineMeetingInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the calendarEventId property value. The identifier of the calendar event associated with the meeting.
// Parameters:
//  - value : Value to set for the calendarEventId property.
func (m *TeamworkOnlineMeetingInfo) SetCalendarEventId(value *string)() {
    m.calendarEventId = value
}
// Sets the joinWebUrl property value. The URL which can be clicked on to join or uniquely identify the meeting.
// Parameters:
//  - value : Value to set for the joinWebUrl property.
func (m *TeamworkOnlineMeetingInfo) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
// Sets the organizer property value. The organizer of the meeting.
// Parameters:
//  - value : Value to set for the organizer property.
func (m *TeamworkOnlineMeetingInfo) SetOrganizer(value *TeamworkUserIdentity)() {
    m.organizer = value
}
