package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageEvent provides operations to manage the collection of accessReview entities.
type MessageEvent struct {
    Entity
    // The dateTime property
    dateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The eventType property
    eventType *MessageEventType
}
// NewMessageEvent instantiates a new messageEvent and sets the default values.
func NewMessageEvent()(*MessageEvent) {
    m := &MessageEvent{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.messageEvent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMessageEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageEvent(), nil
}
// GetDateTime gets the dateTime property value. The dateTime property
func (m *MessageEvent) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateTime
}
// GetDescription gets the description property value. The description property
func (m *MessageEvent) GetDescription()(*string) {
    return m.description
}
// GetEventType gets the eventType property value. The eventType property
func (m *MessageEvent) GetEventType()(*MessageEventType) {
    return m.eventType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["eventType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMessageEventType , m.SetEventType)
    return res
}
// Serialize serializes information the current object
func (m *MessageEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetEventType()).String()
        err = writer.WriteStringValue("eventType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDateTime sets the dateTime property value. The dateTime property
func (m *MessageEvent) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateTime = value
}
// SetDescription sets the description property value. The description property
func (m *MessageEvent) SetDescription(value *string)() {
    m.description = value
}
// SetEventType sets the eventType property value. The eventType property
func (m *MessageEvent) SetEventType(value *MessageEventType)() {
    m.eventType = value
}
