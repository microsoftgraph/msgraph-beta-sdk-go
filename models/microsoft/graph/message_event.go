package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MessageEvent struct {
    Entity
    // 
    dateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    description *string;
    // 
    eventType *MessageEventType;
}
// Instantiates a new messageEvent and sets the default values.
func NewMessageEvent()(*MessageEvent) {
    m := &MessageEvent{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the dateTime property value. 
func (m *MessageEvent) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dateTime
    }
}
// Gets the description property value. 
func (m *MessageEvent) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the eventType property value. 
func (m *MessageEvent) GetEventType()(*MessageEventType) {
    if m == nil {
        return nil
    } else {
        return m.eventType
    }
}
// The deserialization information for the current model
func (m *MessageEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["eventType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMessageEventType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MessageEventType)
            m.SetEventType(&cast)
        }
        return nil
    }
    return res
}
func (m *MessageEvent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MessageEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("dateTime", m.GetDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetEventType() != nil {
        cast := m.GetEventType().String()
        err = writer.WriteStringValue("eventType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the dateTime property value. 
// Parameters:
//  - value : Value to set for the dateTime property.
func (m *MessageEvent) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateTime = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *MessageEvent) SetDescription(value *string)() {
    m.description = value
}
// Sets the eventType property value. 
// Parameters:
//  - value : Value to set for the eventType property.
func (m *MessageEvent) SetEventType(value *MessageEventType)() {
    m.eventType = value
}
