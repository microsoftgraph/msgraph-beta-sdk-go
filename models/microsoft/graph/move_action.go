package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MoveAction 
type MoveAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The name of the location the item was moved from.
    from *string;
    // The name of the location the item was moved to.
    to *string;
}
// NewMoveAction instantiates a new moveAction and sets the default values.
func NewMoveAction()(*MoveAction) {
    m := &MoveAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MoveAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFrom gets the from property value. The name of the location the item was moved from.
func (m *MoveAction) GetFrom()(*string) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
// GetTo gets the to property value. The name of the location the item was moved to.
func (m *MoveAction) GetTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.to
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MoveAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val)
        }
        return nil
    }
    res["to"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTo(val)
        }
        return nil
    }
    return res
}
func (m *MoveAction) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MoveAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("from", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("to", m.GetTo())
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
func (m *MoveAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFrom sets the from property value. The name of the location the item was moved from.
func (m *MoveAction) SetFrom(value *string)() {
    if m != nil {
        m.from = value
    }
}
// SetTo sets the to property value. The name of the location the item was moved to.
func (m *MoveAction) SetTo(value *string)() {
    if m != nil {
        m.to = value
    }
}
