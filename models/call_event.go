package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallEvent 
type CallEvent struct {
    Entity
}
// NewCallEvent instantiates a new callEvent and sets the default values.
func NewCallEvent()(*CallEvent) {
    m := &CallEvent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCallEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallEvent(), nil
}
// GetCallEventType gets the callEventType property value. The callEventType property
func (m *CallEvent) GetCallEventType()(*CallEventType) {
    val, err := m.GetBackingStore().Get("callEventType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CallEventType)
    }
    return nil
}
// GetDirection gets the direction property value. The direction property
func (m *CallEvent) GetDirection()(*CallDirection) {
    val, err := m.GetBackingStore().Get("direction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CallDirection)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["callEventType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCallEventType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallEventType(val.(*CallEventType))
        }
        return nil
    }
    res["direction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCallDirection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirection(val.(*CallDirection))
        }
        return nil
    }
    res["joinCallUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinCallUrl(val)
        }
        return nil
    }
    return res
}
// GetJoinCallUrl gets the joinCallUrl property value. The joinCallUrl property
func (m *CallEvent) GetJoinCallUrl()(*string) {
    val, err := m.GetBackingStore().Get("joinCallUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CallEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCallEventType() != nil {
        cast := (*m.GetCallEventType()).String()
        err = writer.WriteStringValue("callEventType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDirection() != nil {
        cast := (*m.GetDirection()).String()
        err = writer.WriteStringValue("direction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joinCallUrl", m.GetJoinCallUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCallEventType sets the callEventType property value. The callEventType property
func (m *CallEvent) SetCallEventType(value *CallEventType)() {
    err := m.GetBackingStore().Set("callEventType", value)
    if err != nil {
        panic(err)
    }
}
// SetDirection sets the direction property value. The direction property
func (m *CallEvent) SetDirection(value *CallDirection)() {
    err := m.GetBackingStore().Set("direction", value)
    if err != nil {
        panic(err)
    }
}
// SetJoinCallUrl sets the joinCallUrl property value. The joinCallUrl property
func (m *CallEvent) SetJoinCallUrl(value *string)() {
    err := m.GetBackingStore().Set("joinCallUrl", value)
    if err != nil {
        panic(err)
    }
}
// CallEventable 
type CallEventable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallEventType()(*CallEventType)
    GetDirection()(*CallDirection)
    GetJoinCallUrl()(*string)
    SetCallEventType(value *CallEventType)()
    SetDirection(value *CallDirection)()
    SetJoinCallUrl(value *string)()
}
