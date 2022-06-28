package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintTaskTrigger provides operations to manage the collection of accessReview entities.
type PrintTaskTrigger struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The definition property
    definition PrintTaskDefinitionable
    // The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table.
    event *PrintEvent
}
// NewPrintTaskTrigger instantiates a new printTaskTrigger and sets the default values.
func NewPrintTaskTrigger()(*PrintTaskTrigger) {
    m := &PrintTaskTrigger{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrintTaskTriggerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintTaskTriggerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintTaskTrigger(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintTaskTrigger) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefinition gets the definition property value. The definition property
func (m *PrintTaskTrigger) GetDefinition()(PrintTaskDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// GetEvent gets the event property value. The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table.
func (m *PrintTaskTrigger) GetEvent()(*PrintEvent) {
    if m == nil {
        return nil
    } else {
        return m.event
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintTaskTrigger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrintTaskDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinition(val.(PrintTaskDefinitionable))
        }
        return nil
    }
    res["event"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintEvent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val.(*PrintEvent))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PrintTaskTrigger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    if m.GetEvent() != nil {
        cast := (*m.GetEvent()).String()
        err = writer.WriteStringValue("event", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintTaskTrigger) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefinition sets the definition property value. The definition property
func (m *PrintTaskTrigger) SetDefinition(value PrintTaskDefinitionable)() {
    if m != nil {
        m.definition = value
    }
}
// SetEvent sets the event property value. The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table.
func (m *PrintTaskTrigger) SetEvent(value *PrintEvent)() {
    if m != nil {
        m.event = value
    }
}
