package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertTrigger 
type AlertTrigger struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of the property serving as a detection trigger.
    name *string
    // Type of the property in the key:value pair for interpretation. For example, String, Boolean etc.
    type_escaped *string
    // Value of the property serving as a detection trigger.
    value *string
}
// NewAlertTrigger instantiates a new alertTrigger and sets the default values.
func NewAlertTrigger()(*AlertTrigger) {
    m := &AlertTrigger{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAlertTriggerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertTriggerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertTrigger(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertTrigger) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertTrigger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the property serving as a detection trigger.
func (m *AlertTrigger) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetType gets the type property value. Type of the property in the key:value pair for interpretation. For example, String, Boolean etc.
func (m *AlertTrigger) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetValue gets the value property value. Value of the property serving as a detection trigger.
func (m *AlertTrigger) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *AlertTrigger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertTrigger) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetName sets the name property value. Name of the property serving as a detection trigger.
func (m *AlertTrigger) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetType sets the type property value. Type of the property in the key:value pair for interpretation. For example, String, Boolean etc.
func (m *AlertTrigger) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetValue sets the value property value. Value of the property serving as a detection trigger.
func (m *AlertTrigger) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
