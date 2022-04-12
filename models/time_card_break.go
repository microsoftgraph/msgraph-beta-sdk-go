package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeCardBreak 
type TimeCardBreak struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // ID of the timeCardBreak.
    breakId *string
    // The start event of the timeCardBreak.
    end TimeCardEventable
    // Notes about the timeCardBreak.
    notes ItemBodyable
    // The start property
    start TimeCardEventable
}
// NewTimeCardBreak instantiates a new timeCardBreak and sets the default values.
func NewTimeCardBreak()(*TimeCardBreak) {
    m := &TimeCardBreak{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTimeCardBreakFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeCardBreakFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeCardBreak(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeCardBreak) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBreakId gets the breakId property value. ID of the timeCardBreak.
func (m *TimeCardBreak) GetBreakId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.breakId
    }
}
// GetEnd gets the end property value. The start event of the timeCardBreak.
func (m *TimeCardBreak) GetEnd()(TimeCardEventable) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeCardBreak) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["breakId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBreakId(val)
        }
        return nil
    }
    res["end"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val.(TimeCardEventable))
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val.(ItemBodyable))
        }
        return nil
    }
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeCardEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val.(TimeCardEventable))
        }
        return nil
    }
    return res
}
// GetNotes gets the notes property value. Notes about the timeCardBreak.
func (m *TimeCardBreak) GetNotes()(ItemBodyable) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetStart gets the start property value. The start property
func (m *TimeCardBreak) GetStart()(TimeCardEventable) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
// Serialize serializes information the current object
func (m *TimeCardBreak) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("breakId", m.GetBreakId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("start", m.GetStart())
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
func (m *TimeCardBreak) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBreakId sets the breakId property value. ID of the timeCardBreak.
func (m *TimeCardBreak) SetBreakId(value *string)() {
    if m != nil {
        m.breakId = value
    }
}
// SetEnd sets the end property value. The start event of the timeCardBreak.
func (m *TimeCardBreak) SetEnd(value TimeCardEventable)() {
    if m != nil {
        m.end = value
    }
}
// SetNotes sets the notes property value. Notes about the timeCardBreak.
func (m *TimeCardBreak) SetNotes(value ItemBodyable)() {
    if m != nil {
        m.notes = value
    }
}
// SetStart sets the start property value. The start property
func (m *TimeCardBreak) SetStart(value TimeCardEventable)() {
    if m != nil {
        m.start = value
    }
}
